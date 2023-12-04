package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/jiyeyuran/go-protoo"
	"github.com/jiyeyuran/mediasoup-go"
	"github.com/jiyeyuran/mediasoup-go/internal/proto"
	"github.com/rs/zerolog"
)

type Room struct {
	mediasoup.IEventEmitter
	logger             zerolog.Logger
	peerLockers        sync.Map
	config             Config
	roomId             string
	protooRoom         *protoo.Room
	mediasoupRouter    *mediasoup.Router
	webrtcServer       *mediasoup.WebRtcServer
	audioLevelObserver *mediasoup.AudioLevelObserver
	bot                *Bot
	networkThrottled   bool
	broadcasters       sync.Map
	closed             uint32
}

func CreateRoom(config Config, roomId string, worker *mediasoup.Worker, server *mediasoup.WebRtcServer) (room *Room, err error) {
	logger := NewLogger("Room")

	mediasoupRouter, err := worker.CreateRouter(config.Mediasoup.RouterOptions)
	if err != nil {
		logger.Err(err).Msg("create router")
		return
	}

	audioLevelObserver, err := mediasoupRouter.CreateAudioLevelObserver(func(o *mediasoup.AudioLevelObserverOptions) {
		o.MaxEntries = 1
		o.Threshold = -80
		o.Interval = 800
	})
	if err != nil {
		logger.Err(err).Msg("create audio level observer")
		return
	}

	bot, err := CreateBot(mediasoupRouter)
	if err != nil {
		return
	}

	room = &Room{
		IEventEmitter:      mediasoup.NewEventEmitter(),
		logger:             logger,
		config:             config,
		roomId:             roomId,
		protooRoom:         protoo.NewRoom(),
		mediasoupRouter:    mediasoupRouter,
		audioLevelObserver: audioLevelObserver,
		bot:                bot,
		webrtcServer:       server,
	}
	room.handleAudioLevelObserver()

	return
}

func (r *Room) Close() {
	if atomic.CompareAndSwapUint32(&r.closed, 0, 1) {
		r.logger.Debug().Msg("close()")

		r.protooRoom.Close()
		r.mediasoupRouter.Close()
		r.bot.Close()

		r.SafeEmit("close")
		r.RemoveAllListeners()

		if r.networkThrottled {
			//TODO: throttle stop
		}
	}
}

func (r *Room) Closed() bool {
	return atomic.LoadUint32(&r.closed) > 0
}

func (r *Room) LogStatus() {
	dump, err := r.mediasoupRouter.Dump()
	if err != nil {
		r.logger.Err(err).Msg("LogStatus()")
		return
	}

	r.logger.Info().
		Str("roomId", r.roomId).
		Int("peers", len(r.getJoinedPeers())).
		Int("transports", len(dump.TransportIds)).
		Msg("logStatus()")
}

func (r *Room) GetRouterRtpCapabilities() mediasoup.RtpCapabilities {
	return r.mediasoupRouter.RtpCapabilities()
}

func (r *Room) HandleProtooConnection(peerId string, transport protoo.Transport) (err error) {
	existingPeer := r.protooRoom.GetPeer(peerId)
	if existingPeer != nil {
		r.logger.Warn().
			Str("peerId", peerId).
			Msg("handleProtooConnection() | there is already a protoo Peer with same peerId, closing it")

		existingPeer.Close()
	}

	peerData := proto.NewPeerData()
	peer, err := r.protooRoom.CreatePeer(peerId, peerData, transport)
	if err != nil {
		r.logger.Err(err).Msg("protooRoom.createPeer() failed")
		return
	}

	locker := &sync.Mutex{}

	r.peerLockers.Store(peer.Id(), locker)

	peer.On("close", func() {
		if r.Closed() {
			return
		}
		r.logger.Debug().Str("peerId", peer.Id()).Msg(`protoo Peer "close" event`)

		data := peer.Data().(*proto.PeerData)

		// If the Peer was joined, notify all Peers.
		if data.Joined {
			for _, otherPeer := range r.getJoinedPeers(peer) {
				otherPeer.Notify("peerClosed", H{
					"peerId": peer.Id(),
				})
			}
		}

		for _, consumer := range data.Consumers() {
			r.removeConsumeData(consumer)
		}

		// Iterate and close all mediasoup Transport associated to this Peer, so all
		// its Producers and Consumers will also be closed.
		for _, transport := range data.Transports() {
			transport.Close()
		}

		// If this is the latest Peer in the room, close the room.
		if len(r.protooRoom.Peers()) == 0 {
			r.logger.Info().Str("roomId", r.roomId).Msg(`last Peer in the room left, closing the room`)
			r.Close()
		}

		// delay for a second to clean locker
		time.AfterFunc(time.Second, func() {
			r.peerLockers.Delete(peer.Id())
		})
	})

	peer.On("request", func(request protoo.Message, accept func(data interface{}), reject func(err error)) {
		r.logger.Debug().Str("method", request.Method).Str("peerId", peerId).Msg(`protoo Peer "request" event`)

		err := r.handleProtooRequest(peer, request, accept)
		if err != nil {
			reject(err)
		}
	})

	return
}

func (r *Room) handleAudioLevelObserver() {
	r.audioLevelObserver.OnVolumes(func(volumes []mediasoup.AudioLevelObserverVolume) {
		producer := volumes[0].Producer
		volume := volumes[0].Volume

		r.logger.Debug().
			Str("producerId", producer.Id()).Int("volume", volume).
			Msg(`audioLevelObserver "volumes" event`)

		for _, peer := range r.getJoinedPeers() {
			peer.Notify("activeSpeaker", map[string]interface{}{
				"peerId": producer.AppData().(H)["peerId"],
				"volume": volume,
			})
		}
	})

	r.audioLevelObserver.OnSilence(func() {
		r.logger.Debug().Msg(`audioLevelObserver "silence" event`)

		for _, peer := range r.getJoinedPeers() {
			peer.Notify("activeSpeaker", map[string]interface{}{
				"peerId": nil,
			})
		}
	})
}

func (r *Room) handleProtooRequest(peer *protoo.Peer, request protoo.Message, accept func(data interface{})) (err error) {
	peerData := peer.Data().(*proto.PeerData)

	switch request.Method {
	case "getRouterRtpCapabilities":
		accept(r.GetRouterRtpCapabilities())

	case "join":
		if peerData.Joined {
			err = errors.New("Peer already joined")
			return
		}

		requestData := proto.PeerData{}

		if err = json.Unmarshal(request.Data, &requestData); err != nil {
			return
		}

		// Store client data into the protoo Peer data object.
		peerData.Joined = true
		peerData.DisplayName = requestData.DisplayName
		peerData.Device = requestData.Device
		peerData.RtpCapabilities = requestData.RtpCapabilities
		peerData.SctpCapabilities = requestData.SctpCapabilities

		joinedPeers := []*protoo.Peer{}

		for _, peer := range r.getJoinedPeers() {
			joinedPeers = append(joinedPeers, peer)
		}

		r.broadcasters.Range(func(key, val interface{}) bool {
			peerInfo := val.(*proto.PeerInfo)
			joinedPeers = append(joinedPeers, protoo.NewPeer(peerInfo.Id, peerInfo.Data, nil))

			return true
		})

		peerInfos := []*proto.PeerInfo{}

		for _, joinedPeer := range joinedPeers {
			if joinedPeer.Id() == peer.Id() {
				continue
			}
			data := joinedPeer.Data().(*proto.PeerData)
			peerInfos = append(peerInfos, &proto.PeerInfo{
				Id:          joinedPeer.Id(),
				DisplayName: data.DisplayName,
				Device:      data.Device,
			})
		}

		accept(H{"peers": peerInfos})

		for _, joinedPeer := range joinedPeers {
			data := joinedPeer.Data().(*proto.PeerData)

			// Create Consumers for existing Producers.
			for _, producer := range data.Producers() {
				r.createConsumer(peer, joinedPeer.Id(), producer)
			}

			// Create DataConsumers for existing DataProducers.
			for _, dataProducer := range data.DataProducers() {
				r.createDataConsumer(peer, joinedPeer.Id(), dataProducer)
			}
		}

		// Create DataConsumers for bot DataProducer.
		//r.createDataConsumer(peer, "", r.bot.dataProducer)

		// // Notify the new Peer to all other Peers.
		for _, otherPeer := range r.getJoinedPeers(peer) {
			otherPeer.Notify("newPeer", &proto.PeerInfo{
				Id:          peer.Id(),
				DisplayName: peerData.DisplayName,
				Device:      peerData.Device,
			})
		}

	case "createWebRtcTransport":
		{
			var requestData struct {
				ForceTcp         bool
				Producing        bool
				Consuming        bool
				SctpCapabilities *mediasoup.SctpCapabilities
			}
			if err = json.Unmarshal(request.Data, &requestData); err != nil {
				return
			}

			webRtcTransportOptions := mediasoup.WebRtcTransportOptions{}
			Clone(&webRtcTransportOptions, r.config.Mediasoup.WebRtcTransportOptions)

			webRtcTransportOptions.EnableSctp = requestData.SctpCapabilities != nil

			if requestData.SctpCapabilities != nil {
				webRtcTransportOptions.NumSctpStreams = requestData.SctpCapabilities.NumStreams
			}

			webRtcTransportOptions.AppData = &proto.TransportData{
				Producing: requestData.Producing,
				Consuming: requestData.Consuming,
			}
			webRtcTransportOptions.WebRtcServer = r.webrtcServer
			if requestData.ForceTcp {
				webRtcTransportOptions.EnableUdp = NewBool(false)
				webRtcTransportOptions.EnableTcp = true
			}

			transport, err := r.mediasoupRouter.CreateWebRtcTransport(webRtcTransportOptions)
			if err != nil {
				return err
			}
			transport.OnSctpStateChange(func(sctpState mediasoup.SctpState) {
				r.logger.Debug().Str("sctpState", string(sctpState)).Msg(`WebRtcTransport "sctpstatechange" event`)
			})
			transport.OnDtlsStateChange(func(dtlsState mediasoup.DtlsState) {
				if dtlsState == "failed" || dtlsState == "closed" {
					r.logger.Warn().Str("dtlsState", string(dtlsState)).Msg(`WebRtcTransport "dtlsstatechange" event`)
				}
			})

			// NOTE: For testing.
			// transport.EnableTraceEvent("probation", "bwe")
			if err = transport.EnableTraceEvent("bwe"); err != nil {
				return err
			}

			transport.OnTrace(func(trace *mediasoup.TransportTraceEventData) {
				r.logger.Debug().
					Str("transportId", transport.Id()).
					Str("trace.type", string(trace.Type)).
					Interface("trace", trace).
					Msg(`"transport "trace" event`)

				if trace.Type == "bwe" && trace.Direction == "out" {
					peer.Notify("downlinkBwe", trace.Info)
				}
			})

			// Store the WebRtcTransport into the protoo Peer data Object.
			peerData.AddTransport(transport)

			accept(H{
				"id":             transport.Id(),
				"iceParameters":  transport.IceParameters(),
				"iceCandidates":  transport.IceCandidates(),
				"dtlsParameters": transport.DtlsParameters(),
				"sctpParameters": transport.SctpParameters(),
			})

			maxIncomingBitrate := r.config.Mediasoup.WebRtcTransportOptions.MaxIncomingBitrate

			if maxIncomingBitrate > 0 {
				transport.SetMaxIncomingBitrate(maxIncomingBitrate)
			}
		}

	case "connectWebRtcTransport":
		var requestData struct {
			TransportId    string                    `json:"transportId,omitempty"`
			DtlsParameters *mediasoup.DtlsParameters `json:"dtlsParameters,omitempty"`
		}
		if err = json.Unmarshal(request.Data, &requestData); err != nil {
			return
		}
		transport := peerData.GetTransport(requestData.TransportId)
		if transport == nil {
			err = fmt.Errorf(`transport with id "%s" not found`, requestData.TransportId)
			return
		}
		transport.Connect(mediasoup.TransportConnectOptions{
			DtlsParameters: requestData.DtlsParameters,
		})
		accept(nil)

	case "restartIce":
		var requestData struct {
			TransportId string `json:"transportId,omitempty"`
		}
		if err = json.Unmarshal(request.Data, &requestData); err != nil {
			return
		}
		transport := peerData.GetTransport(requestData.TransportId)
		if transport == nil {
			err = fmt.Errorf(`transport with id "%s" not found`, requestData.TransportId)
			return
		}
		iceParameters, err := transport.(*mediasoup.WebRtcTransport).RestartIce()
		if err != nil {
			return err
		}
		accept(iceParameters)

	case "produce":
		// Ensure the Peer is joined.
		if !peerData.Joined {
			err = errors.New("Peer not yet joined")
			return
		}
		var requestData struct {
			TransportId   string                  `json:"transportId,omitempty"`
			Kind          mediasoup.MediaKind     `json:"kind,omitempty"`
			RtpParameters mediasoup.RtpParameters `json:"rtpParameters,omitempty"`
			AppData       H                       `json:"appData,omitempty"`
		}
		if err = json.Unmarshal(request.Data, &requestData); err != nil {
			return
		}
		transport := peerData.GetTransport(requestData.TransportId)
		if transport == nil {
			err = fmt.Errorf(`transport with id "%s" not found`, requestData.TransportId)
			return
		}
		// // Add peerId into appData to later get the associated Peer during
		// // the "loudest" event of the audioLevelObserver.
		appData := requestData.AppData
		if appData == nil {
			appData = H{}
		}
		appData["peerId"] = peer.Id()

		producer, err := transport.Produce(mediasoup.ProducerOptions{
			Kind:          requestData.Kind,
			RtpParameters: requestData.RtpParameters,
			AppData:       appData,
			// KeyFrameRequestDelay: 5000,
		})
		if err != nil {
			return err
		}
		// Store the Producer into the protoo Peer data Object.
		peerData.AddProducer(producer)

		producer.OnScore(func(score []mediasoup.ProducerScore) {
			r.logger.Debug().Str("producerId", producer.Id()).Interface("score", score).Msg(`producer "score" event`)

			peer.Notify("producerScore", H{
				"producerId": producer.Id(),
				"score":      score,
			})
		})
		producer.OnVideoOrientationChange(func(videoOrientation *mediasoup.ProducerVideoOrientation) {
			r.logger.Debug().
				Str("producerId", producer.Id()).
				Interface("videoOrientation", videoOrientation).
				Msg(`producer "videoorientationchange" event`)
		})

		// NOTE: For testing.
		// producer.EnableTraceEvent("rtp", "keyframe", "nack", "pli", "fir");
		// producer.EnableTraceEvent("pli", "fir");
		// producer.EnableTraceEvent("keyframe");

		producer.OnTrace(func(trace *mediasoup.ProducerTraceEventData) {
			r.logger.Debug().
				Str("producerId", producer.Id()).
				Str("trace.type", string(trace.Type)).
				Interface("trace", trace).
				Msg(`producer "trace" event`)
		})

		accept(H{"id": producer.Id()})

		// Optimization: Create a server-side Consumer for each Peer.
		for _, otherPeer := range r.getJoinedPeers(peer) {
			r.createConsumer(otherPeer, peer.Id(), producer)
		}

		// // Add into the audioLevelObserver.
		if producer.Kind() == mediasoup.MediaKind_Audio {
			r.audioLevelObserver.AddProducer(producer.Id())
		}

	case "closeProducer":
		// Ensure the Peer is joined.
		if !peerData.Joined {
			err = errors.New("Peer not yet joined")
			return
		}
		var requestData struct {
			ProducerId string
		}
		if err = json.Unmarshal(request.Data, &requestData); err != nil {
			return
		}
		producer := peerData.GetProducer(requestData.ProducerId)
		if producer == nil {
			err = fmt.Errorf(`producer with id "%s" not found`, requestData.ProducerId)
			return
		}
		producer.Close()
		peerData.DeleteProducer(producer.Id())

		accept(nil)

	case "pauseProducer":
		// Ensure the Peer is joined.
		if !peerData.Joined {
			err = errors.New("Peer not yet joined")
			return
		}
		var requestData struct {
			ProducerId string
		}
		if err = json.Unmarshal(request.Data, &requestData); err != nil {
			return
		}
		producer := peerData.GetProducer(requestData.ProducerId)
		if producer == nil {
			err = fmt.Errorf(`producer with id "%s" not found`, requestData.ProducerId)
			return
		}
		if err = producer.Pause(); err != nil {
			return
		}

		accept(nil)

	case "resumeProducer":
		// Ensure the Peer is joined.
		if !peerData.Joined {
			err = errors.New("Peer not yet joined")
			return
		}
		var requestData struct {
			ProducerId string
		}
		if err = json.Unmarshal(request.Data, &requestData); err != nil {
			return
		}
		producer := peerData.GetProducer(requestData.ProducerId)
		if producer == nil {
			err = fmt.Errorf(`producer with id "%s" not found`, requestData.ProducerId)
			return
		}
		if err = producer.Resume(); err != nil {
			return
		}

		accept(nil)

	case "pauseConsumer":
		// Ensure the Peer is joined.
		if !peerData.Joined {
			err = errors.New("Peer not yet joined")
			return
		}
		var requestData struct {
			ConsumerId string
		}
		if err = json.Unmarshal(request.Data, &requestData); err != nil {
			return
		}
		consumer := peerData.GetConsumer(requestData.ConsumerId)
		if consumer == nil {
			err = fmt.Errorf(`consumer with id "%s" not found`, requestData.ConsumerId)
			return
		}
		if err = consumer.Pause(); err != nil {
			return
		}

		accept(nil)

	case "resumeConsumer":
		// Ensure the Peer is joined.
		if !peerData.Joined {
			err = errors.New("Peer not yet joined")
			return
		}
		var requestData struct {
			ConsumerId string
		}
		if err = json.Unmarshal(request.Data, &requestData); err != nil {
			return
		}
		consumer := peerData.GetConsumer(requestData.ConsumerId)
		if consumer == nil {
			err = fmt.Errorf(`consumer with id "%s" not found`, requestData.ConsumerId)
			return
		}
		if err = consumer.Resume(); err != nil {
			return
		}

		accept(nil)

	case "setConsumerPreferredLayers":
		// Ensure the Peer is joined.
		if !peerData.Joined {
			err = errors.New("Peer not yet joined")
			return
		}
		var requestData struct {
			mediasoup.ConsumerLayers
			ConsumerId string
		}
		if err = json.Unmarshal(request.Data, &requestData); err != nil {
			return
		}
		consumer := peerData.GetConsumer(requestData.ConsumerId)
		if consumer == nil {
			err = fmt.Errorf(`consumer with id "%s" not found`, requestData.ConsumerId)
			return
		}
		if err = consumer.SetPreferredLayers(requestData.ConsumerLayers); err != nil {
			return
		}

		accept(nil)
		r.addOrUpdateConsumeData(consumer, int(requestData.ConsumerLayers.SpatialLayer))
	case "setConsumerPriority":
		// Ensure the Peer is joined.
		if !peerData.Joined {
			err = errors.New("Peer not yet joined")
			return
		}
		var requestData struct {
			ConsumerId string
			Priority   uint32
		}
		if err = json.Unmarshal(request.Data, &requestData); err != nil {
			return
		}
		consumer := peerData.GetConsumer(requestData.ConsumerId)
		if consumer == nil {
			err = fmt.Errorf(`consumer with id "%s" not found`, requestData.ConsumerId)
			return
		}
		if err = consumer.SetPriority(requestData.Priority); err != nil {
			return
		}

		accept(nil)

	case "requestConsumerKeyFrame":
		// Ensure the Peer is joined.
		if !peerData.Joined {
			err = errors.New("Peer not yet joined")
			return
		}
		var requestData struct {
			ConsumerId string
		}
		if err = json.Unmarshal(request.Data, &requestData); err != nil {
			return
		}
		consumer := peerData.GetConsumer(requestData.ConsumerId)
		if consumer == nil {
			err = fmt.Errorf(`consumer with id "%s" not found`, requestData.ConsumerId)
			return
		}
		if err = consumer.RequestKeyFrame(); err != nil {
			return
		}

		accept(nil)

	case "produceData":
		// Ensure the Peer is joined.
		if !peerData.Joined {
			err = errors.New("Peer not yet joined")
			return
		}
		var requestData struct {
			TransportId          string                          `json:"transportId,omitempty"`
			SctpStreamParameters *mediasoup.SctpStreamParameters `json:"sctpStreamParameters,omitempty"`
			Label                string                          `json:"label,omitempty"`
			Protocol             string                          `json:"protocol,omitempty"`
			AppData              H                               `json:"appData,omitempty"`
		}
		if err = json.Unmarshal(request.Data, &requestData); err != nil {
			return
		}
		transport := peerData.GetTransport(requestData.TransportId)
		if transport == nil {
			err = fmt.Errorf(`transport with id "%s" not found`, requestData.TransportId)
			return
		}
		dataProducer, err := transport.ProduceData(mediasoup.DataProducerOptions{
			SctpStreamParameters: requestData.SctpStreamParameters,
			Label:                requestData.Label,
			Protocol:             requestData.Protocol,
			AppData:              requestData.AppData,
		})
		if err != nil {
			return err
		}
		peerData.AddDataProducer(dataProducer)

		accept(H{"id": dataProducer.Id()})

		switch dataProducer.Label() {
		case "chat":
			// Create a server-side DataConsumer for each Peer.
			for _, otherPeer := range r.getJoinedPeers(peer) {
				r.createDataConsumer(otherPeer, peer.Id(), dataProducer)
			}

		case "bot":
			// Pass it to the bot.
			r.bot.HandlePeerDataProducer(dataProducer.Id(), peer)
		}

	case "changeDisplayName":
		// Ensure the Peer is joined.
		if !peerData.Joined {
			err = errors.New("Peer not yet joined")
			return
		}
		var requestData struct {
			DisplayName string `json:"displayName,omitempty"`
		}
		if err = json.Unmarshal(request.Data, &requestData); err != nil {
			return
		}
		oldDisplayName := peerData.DisplayName
		peerData.DisplayName = requestData.DisplayName

		// Notify other joined Peers.
		for _, otherPeer := range r.getJoinedPeers(peer) {
			otherPeer.Notify("peerDisplayNameChanged", H{
				"peerId":         peer.Id(),
				"displayName":    requestData.DisplayName,
				"oldDisplayName": oldDisplayName,
			})
		}

		accept(nil)

	case "getTransportStats":
		// Ensure the Peer is joined.
		if !peerData.Joined {
			err = errors.New("Peer not yet joined")
			return
		}
		var requestData struct {
			TransportId string `json:"transportId,omitempty"`
		}
		if err = json.Unmarshal(request.Data, &requestData); err != nil {
			return
		}
		transport := peerData.GetTransport(requestData.TransportId)
		if transport == nil {
			err = fmt.Errorf(`transport with id "%s" not found`, requestData.TransportId)
			return
		}
		stats, err := transport.GetStats()
		if err != nil {
			return err
		}

		accept(stats)

	case "getProducerStats":
		var requestData struct {
			ProducerId string
		}
		if err = json.Unmarshal(request.Data, &requestData); err != nil {
			return
		}
		producer := peerData.GetProducer(requestData.ProducerId)
		if producer == nil {
			err = fmt.Errorf(`producer with id "%s" not found`, requestData.ProducerId)
			return
		}
		stats, err := producer.GetStats()
		if err != nil {
			return err
		}

		accept(stats)

	case "getConsumerStats":
		var requestData struct {
			ConsumerId string
		}
		if err = json.Unmarshal(request.Data, &requestData); err != nil {
			return
		}
		consumer := peerData.GetConsumer(requestData.ConsumerId)
		if consumer == nil {
			err = fmt.Errorf(`consumer with id "%s" not found`, requestData.ConsumerId)
			return
		}
		stats, err := consumer.GetStats()
		if err != nil {
			return err
		}

		accept(stats)

	case "getDataProducerStats":
		var requestData struct {
			DataProducerId string
		}
		if err = json.Unmarshal(request.Data, &requestData); err != nil {
			return
		}
		dataProducer := peerData.GetDataProducer(requestData.DataProducerId)
		if dataProducer == nil {
			err = fmt.Errorf(`dataProducer with id "%s" not found`, requestData.DataProducerId)
			return
		}
		stats, err := dataProducer.GetStats()
		if err != nil {
			return err
		}

		accept(stats)

	case "getDataConsumerStats":
		var requestData struct {
			DataConsumerId string
		}
		if err = json.Unmarshal(request.Data, &requestData); err != nil {
			return
		}
		dataConsumer := peerData.GetDataConsumer(requestData.DataConsumerId)
		if dataConsumer == nil {
			err = fmt.Errorf(`dataConsumer with id "%s" not found`, requestData.DataConsumerId)
			return
		}
		stats, err := dataConsumer.GetStats()
		if err != nil {
			return err
		}

		accept(stats)

	case "applyNetworkThrottle":
		//TODO: throttle.start

	case "resetNetworkThrottle":
		//TODO: throttle.stop

	}

	return
}
func (r *Room) addOrUpdateConsumeData(consumer *mediasoup.Consumer, layer int) {
	producerId := consumer.ProducerId()
	producerPeerId := consumer.AppData().(H)["peerId"].(string)
	producerPeer := r.protooRoom.GetPeer(producerPeerId)
	producerPeerData := producerPeer.Data().(*proto.PeerData)
	producerPeerData.AddOrUpdateConsumeStatus(consumer.ProducerId(), consumer.Id(), consumer.Paused(), layer)

	newPaused, newDesiredQ := producerPeerData.GetProducerConsumeData(producerId)

	producerPeer.Notify("producerConsumeStatusChange", H{
		"producerId": producerId,
		"paused":     newPaused,
		"desiredQ":   newDesiredQ,
	})
}

func (r *Room) removeConsumeData(consumer *mediasoup.Consumer) {
	producerId := consumer.ProducerId()
	producerPeerId := consumer.AppData().(H)["peerId"].(string)
	producerPeer := r.protooRoom.GetPeer(producerPeerId)
	producerPeerData := producerPeer.Data().(*proto.PeerData)
	producerPeerData.RemoveConsumeStatus(producerId, consumer.Id())

	newPaused, newDesiredQ := producerPeerData.GetProducerConsumeData(producerId)

	producerPeer.Notify("producerConsumeStatusChange", H{
		"producerId": producerId,
		"paused":     newPaused,
		"desiredQ":   newDesiredQ,
	})
}

func (r *Room) createConsumer(consumerPeer *protoo.Peer, producerPeerId string, producer *mediasoup.Producer) {
	// Optimization:
	// - Create the server-side Consumer in paused mode.
	// - Tell its Peer about it and wait for its response.
	// - Upon receipt of the response, resume the server-side Consumer.
	// - If video, this will mean a single key frame requested by the
	//   server-side Consumer (when resuming it).
	// - If audio (or video), it will avoid that RTP packets are received by the
	//   remote endpoint *before* the Consumer is locally created in the endpoint
	//   (and before the local SDP O/A procedure ends). If that happens (RTP
	//   packets are received before the SDP O/A is done) the PeerConnection may
	//   fail to associate the RTP stream.

	consumerPeerData := consumerPeer.Data().(*proto.PeerData)

	// NOTE: Don"t create the Consumer if the remote Peer cannot consume it.
	if consumerPeerData.RtpCapabilities == nil ||
		!r.mediasoupRouter.CanConsume(producer.Id(), *consumerPeerData.RtpCapabilities) {
		return
	}

	// Must take the Transport the remote Peer is using for consuming.
	var transport mediasoup.ITransport

	for _, t := range consumerPeerData.Transports() {
		if data, ok := t.AppData().(*proto.TransportData); ok && data.Consuming {
			transport = t
			break
		}
	}

	// This should not happen.
	if transport == nil {
		r.logger.Warn().Msg("createConsumer() | Transport for consuming not found")
		return
	}

	consumer, err := transport.Consume(mediasoup.ConsumerOptions{
		ProducerId:      producer.Id(),
		RtpCapabilities: *consumerPeerData.RtpCapabilities,
		Paused:          false,
		AppData:         producer.AppData(),
	})
	if err != nil {
		r.logger.Err(err).Msg("createConsumer() | transport.consume()")
		return
	}

	// Store the Consumer into the protoo consumerPeer data Object.
	consumerPeerData.AddConsumer(consumer)

	// Set Consumer events.
	consumer.OnTransportClose(func() {
		// Remove from its map.
		consumerPeerData.DeleteConsumer(consumer.Id())
	})
	consumer.OnProducerClose(func() {
		// Remove from its map.
		consumerPeerData.DeleteConsumer(consumer.Id())
		consumerPeer.Notify("consumerClosed", H{
			"consumerId": consumer.Id(),
		})
	})
	consumer.OnProducerPause(func() {
		consumerPeer.Notify("consumerPaused", H{
			"consumerId": consumer.Id(),
		})
	})
	consumer.OnProducerResume(func() {
		consumerPeer.Notify("consumerResumed", H{
			"consumerId": consumer.Id(),
		})
	})
	consumer.OnScore(func(score *mediasoup.ConsumerScore) {
		r.logger.Debug().
			Str("consumerId", consumer.Id()).
			Interface("score", score).Msg(`consumer "score" event`)

		consumerPeer.Notify("consumerScore", H{
			"consumerId": consumer.Id(),
			"score":      score,
		})
	})
	consumer.OnLayersChange(func(layers *mediasoup.ConsumerLayers) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in f", r)
			}
		}()

		notifyData := H{
			"consumerId": consumer.Id(),
		}
		notifyData["spatialLayer"] = layers.SpatialLayer
		notifyData["temporalLayer"] = layers.TemporalLayer
		consumerPeer.Notify("consumerLayersChanged", notifyData)
	})

	// NOTE: For testing.
	// consumer.EnableTraceEvent("rtp", "keyframe", "nack", "pli", "fir");
	// consumer.EnableTraceEvent("pli", "fir");
	// consumer.EnableTraceEvent("keyframe");

	consumer.OnTrace(func(trace *mediasoup.ConsumerTraceEventData) {
		r.logger.Debug().
			Str("consumerId", consumer.Id()).
			Str("trace.type", string(trace.Type)).
			Interface("trace", trace).
			Msg(`consumer "trace" event`)
	})

	go func() {
		// Send a protoo request to the remote Peer with Consumer parameters.
		rsp := consumerPeer.Request("newConsumer", H{
			"peerId":         producerPeerId,
			"producerId":     producer.Id(),
			"id":             consumer.Id(),
			"kind":           consumer.Kind(),
			"rtpParameters":  consumer.RtpParameters(),
			"type":           consumer.Type(),
			"appData":        consumer.AppData(),
			"producerPaused": consumer.ProducerPaused(),
		})
		if rsp.Err() != nil {
			r.logger.Warn().Err(rsp.Err()).Msg("createConsumer() | failed")
			return
		}

		// Now that we got the positive response from the remote endpoint, resume
		// the Consumer so the remote endpoint will receive the a first RTP packet
		// of this new stream once its PeerConnection is already ready to process
		// and associate it.
		if err = consumer.Resume(); err != nil {
			r.logger.Warn().Err(err).Msg("createConsumer() | failed")
			return
		}

		consumerPeer.Notify("consumerScore", H{
			"consumerId": consumer.Id(),
			"score":      consumer.Score(),
		})
	}()
}

func (r *Room) createDataConsumer(dataConsumerPeer *protoo.Peer, dataProducerPeerId string, dataProducer *mediasoup.DataProducer) {
	dataConsumerPeerData := dataConsumerPeer.Data().(*proto.PeerData)

	// NOTE: Don't create the DataConsumer if the remote Peer cannot consume it.
	if dataConsumerPeerData.SctpCapabilities == nil {
		return
	}

	// Must take the Transport the remote Peer is using for consuming.
	var transport mediasoup.ITransport

	for _, t := range dataConsumerPeerData.Transports() {
		if data, ok := t.AppData().(*proto.TransportData); ok && data.Consuming {
			transport = t
			break
		}
	}

	// This should not happen.
	if transport == nil {
		r.logger.Warn().Msg("createDataConsumer() | Transport for consuming not found")
		return
	}

	dataConsumer, err := transport.ConsumeData(mediasoup.DataConsumerOptions{
		DataProducerId: dataProducer.Id(),
	})
	if err != nil {
		r.logger.Err(err).Msg("createDataConsumer() | transport.consumeData()")
		return
	}

	// Store the Consumer into the protoo consumerPeer data Object.
	dataConsumerPeerData.AddDataConsumer(dataConsumer)

	// Set DataConsumer events.
	dataConsumer.OnTransportClose(func() {
		// Remove from its map.
		dataConsumerPeerData.DeleteDataConsumer(dataConsumer.Id())
	})
	dataConsumer.OnDataProducerClose(func() {
		// Remove from its map.
		dataConsumerPeerData.DeleteDataConsumer(dataConsumer.Id())
		dataConsumerPeer.Notify("dataConsumerClosed", H{
			"dataConsumerId": dataConsumer.Id(),
		})
	})

	go func() {
		// Send a protoo request to the remote Peer with Consumer parameters.
		rsp := dataConsumerPeer.Request("newDataConsumer", H{
			// This is null for bot DataProducer.
			"peerId":               dataProducerPeerId,
			"dataProducerId":       dataProducer.Id(),
			"id":                   dataConsumer.Id(),
			"sctpStreamParameters": dataConsumer.SctpStreamParameters(),
			"label":                dataConsumer.Label(),
			"protocol":             dataConsumer.Protocol(),
			"appData":              dataConsumer.AppData(),
		})
		if rsp.Err() != nil {
			r.logger.Warn().Err(rsp.Err()).Msg("createDataConsumer() | failed")
		}
	}()
}

func (r *Room) getJoinedPeers(excludePeers ...*protoo.Peer) (peers []*protoo.Peer) {
	for _, peer := range r.protooRoom.Peers() {
		found := false
		for _, excludePeer := range excludePeers {
			if peer == excludePeer {
				found = true
				break
			}
		}
		if !found {
			peers = append(peers, peer)
		}
	}
	return
}
