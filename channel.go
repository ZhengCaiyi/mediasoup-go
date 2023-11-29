package mediasoup

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/go-logr/logr"
	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/jiyeyuran/mediasoup-go/netcodec"

	FBSMessage "github.com/jiyeyuran/mediasoup-go/FBS/Message"

	FBSNotifcication "github.com/jiyeyuran/mediasoup-go/FBS/Notification"

	FBSRequest "github.com/jiyeyuran/mediasoup-go/FBS/Request"

	FBSReponse "github.com/jiyeyuran/mediasoup-go/FBS/Response"
)

type channelSubscriber func(event string, data []byte)
type channelSubscriberFBS func(event FBSNotifcication.Event, data *FBSNotifcication.BodyT)

type Channel struct {
	logger          logr.Logger
	codec           netcodec.Codec
	closed          int32
	pid             int
	nextId          uint32
	sents           sync.Map
	sentChan        chan sentInfo
	closeCh         chan struct{}
	useHandlerID    bool
	oldCloseMethods map[string]string
	subscribers     sync.Map
}

func newChannel(codec netcodec.Codec, pid int, useHandlerID bool) *Channel {
	logger := NewLogger("Channel")

	logger.V(1).Info("constructor()", "useHandlerID", useHandlerID)

	channel := &Channel{
		logger:       logger,
		codec:        codec,
		pid:          pid,
		sentChan:     make(chan sentInfo),
		closeCh:      make(chan struct{}),
		useHandlerID: useHandlerID,
		oldCloseMethods: map[string]string{
			"worker.closeWebRtcServer":    "webRtcServer.close",
			"worker.closeRouter":          "router.close",
			"router.closeTransport":       "transport.close",
			"router.closeRtpObserver":     "rtpObserver.close",
			"transport.closeProducer":     "producer.close",
			"transport.closeConsumer":     "consumer.close",
			"transport.closeDataProducer": "dataProducer.close",
			"transport.closeDataConsumer": "dataConsumer.close",
		},
	}

	return channel
}

func (c *Channel) Start() {
	go c.runWriteLoop()
	go c.runReadLoop()
}

func (c *Channel) Close() error {
	if atomic.CompareAndSwapInt32(&c.closed, 0, 1) {
		c.logger.V(1).Info("close()")
		close(c.closeCh)
		return c.codec.Close()
	}
	return nil
}

func (c *Channel) Closed() bool {
	return atomic.LoadInt32(&c.closed) > 0
}

func (c *Channel) FFSRequest(method FBSRequest.Method, body *FBSRequest.BodyT, internal internalData) (rsp workerResponse) {

	if c.Closed() {
		rsp.err = NewInvalidStateError("Channel closed")
		return
	}

	id := atomic.AddUint32(&c.nextId, 1)

	builder := flatbuffers.NewBuilder(0)
	methodString := FBSRequest.EnumNamesMethod[method]
	handlerId := internal.HandlerID(methodString)
	req := new(FBSRequest.RequestT)
	req.Id = id
	req.Method = method
	req.HandlerId = handlerId
	req.Body = body

	msgBody := new(FBSMessage.BodyT)
	msgBody.Type = FBSMessage.BodyRequest
	msgBody.Value = req

	message := new(FBSMessage.MessageT)
	message.Data = msgBody

	offsset := message.Pack(builder)
	builder.Finish(offsset)
	request := builder.FinishedBytes()

	sent := sentInfo{
		method:  methodString,
		request: request,
		respCh:  make(chan workerResponse),
	}
	c.sents.Store(id, sent)
	defer c.sents.Delete(id)

	timer := time.NewTimer(time.Duration(3000) * time.Millisecond)
	defer timer.Stop()

	// send request
	select {
	case c.sentChan <- sent:
	case <-timer.C:
		rsp.err = fmt.Errorf("Channel request timeout, id: %d, method: %s", id, method)
	case <-c.closeCh:
		rsp.err = NewInvalidStateError("Channel closed, id: %d, method: %s", id, method)
	}
	if rsp.err != nil {
		return
	}

	// wait response
	select {
	case rsp = <-sent.respCh:
	case <-timer.C:
		rsp.err = fmt.Errorf("Channel response timeout, id: %d, method: %s", id, method)
	case <-c.closeCh:
		rsp.err = NewInvalidStateError("Channel closed, id: %d, method: %s", id, method)
	}
	return

}

func (c *Channel) Request(method string, internal internalData, data ...interface{}) (rsp workerResponse) {
	return
}

func (c *Channel) Subscribe(targetId string, handler channelSubscriber) {
	c.subscribers.Store(targetId, handler)
}

func (c *Channel) SubscribeFBS(targetId string, handler channelSubscriberFBS) {
	c.subscribers.Store(targetId, handler)
}

func (c *Channel) Unsubscribe(targetId string) {
	c.subscribers.Delete(targetId)
}

func (c *Channel) runWriteLoop() {
	defer c.Close()

	for {
		select {
		case sentInfo := <-c.sentChan:
			if err := c.codec.WritePayload(sentInfo.request); err != nil {
				sentInfo.respCh <- workerResponse{err: err}
				break
			}
		case <-c.closeCh:
			return
		}
	}
}

func (c *Channel) runReadLoop() {
	defer c.Close()

	for {
		payload, err := c.codec.ReadPayload()
		if err != nil {
			c.logger.Error(err, "read failed")
			break
		}
		c.processPayload(payload)
	}
}

func (c *Channel) processPayload(nsPayload []byte) {

	/*
		switch nsPayload[0] {
		case '{':
			c.processMessage(nsPayload)
		case 'D':
			c.logger.V(1).Info(string(nsPayload[1:]), "pid", c.pid)
		case 'W':
			c.logger.Info(string(nsPayload[1:]), "pid", c.pid, "warn", true)
		case 'E':
			c.logger.Error(nil, string(nsPayload[1:]), "pid", c.pid)
		case 'X':
			fmt.Printf("%s\n", nsPayload[1:])
		default:
			c.logger.Error(errors.New(string(nsPayload[1:])), "unexpected data", "pid", c.pid)
		}
	*/
	messageFBS := FBSMessage.GetSizePrefixedRootAsMessage(nsPayload, 0)
	msg := messageFBS.UnPack()
	msgData := msg.Data
	if msgData.Type == FBSMessage.BodyResponse {
		response := msgData.Value.(*FBSReponse.ResponseT)
		c.processResponse(response)
	} else if msgData.Type == FBSMessage.BodyNotification {
		notification := msgData.Value.(*FBSNotifcication.NotificationT)
		c.processNotification(notification)
	}
}

func (c *Channel) processResponse(msg *FBSReponse.ResponseT) {
	value, ok := c.sents.Load(msg.Id)

	if !ok {
		return
	}
	sent := value.(sentInfo)

	if msg.Accepted {
		c.logger.V(1).Info("request succeeded", "method", sent.method, "id", msg.Id)

		sent.respCh <- workerResponse{response: msg.Body}
	} else if len(msg.Error) > 0 {
		c.logger.Error(errors.New(msg.Reason), "request failed", "method", sent.method, "id", msg.Id)

		if msg.Error == "TypeError" {
			sent.respCh <- workerResponse{err: NewTypeError(msg.Reason)}
		} else {
			sent.respCh <- workerResponse{err: errors.New(msg.Reason)}
		}
	} else {
		c.logger.Error(nil, "received response is not accepted nor rejected", "method", sent.method, "id", msg.Id)
	}
}

func (c *Channel) processNotification(msg *FBSNotifcication.NotificationT) {

	if handler, ok := c.subscribers.Load(msg.HandlerId); ok {
		handler.(channelSubscriberFBS)(msg.Event, msg.Body)
		c.logger.V(1).Info("received a notification", "targetId", msg.HandlerId, "event", msg.Event.String())
	} else {
		c.logger.V(1).Info("received an unhandled notification", "targetId", msg.HandlerId, "event", msg.Event.String())
	}
}

func (c *Channel) processMessage(nsPayload []byte) {

	/*
		var msg struct {
			// response
			Id       int64  `json:"id,omitempty"`
			Accepted bool   `json:"accepted,omitempty"`
			Error    string `json:"error,omitempty"`
			Reason   string `json:"reason,omitempty"`
			// notification
			TargetId interface{} `json:"targetId,omitempty"`
			Event    string      `json:"event,omitempty"`
			// common data
			Data json.RawMessage `json:"data,omitempty"`
		}
		if err := json.Unmarshal(nsPayload, &msg); err != nil {
			c.logger.Error(err, "received response, failed to unmarshal to json", "payload", json.RawMessage(nsPayload))
			return
		}

		if msg.Id > 0 {
			value, ok := c.sents.Load(msg.Id)
			if !ok {
				c.logger.Error(nil, "received response does not match any sent request", "id", msg.Id)
				return
			}
			sent := value.(sentInfo)

			if msg.Accepted {
				c.logger.V(1).Info("request succeeded", "method", sent.method, "id", msg.Id)

				sent.respCh <- workerResponse{data: msg.Data}
			} else if len(msg.Error) > 0 {
				c.logger.Error(errors.New(msg.Reason), "request failed", "method", sent.method, "id", msg.Id)

				if msg.Error == "TypeError" {
					sent.respCh <- workerResponse{err: NewTypeError(msg.Reason)}
				} else {
					sent.respCh <- workerResponse{err: errors.New(msg.Reason)}
				}
			} else {
				c.logger.Error(nil, "received response is not accepted nor rejected", "method", sent.method, "id", msg.Id)
			}
		} else if msg.TargetId != nil && len(msg.Event) > 0 {
			var targetId string
			// The type of msg.TargetId should be string or float64
			switch v := msg.TargetId.(type) {
			case string:
				targetId = v
			case float64:
				targetId = fmt.Sprintf("%0.0f", v)
			default:
				targetId = fmt.Sprintf("%v", v)
			}

			if handler, ok := c.subscribers.Load(targetId); ok {
				handler.(channelSubscriber)(msg.Event, msg.Data)
				c.logger.V(1).Info("received a notification", "targetId", targetId, "event", msg.Event)
			} else {
				c.logger.V(1).Info("received an unhandled notification", "targetId", targetId, "event", msg.Event)
			}
		} else {
			c.logger.Error(nil, "received message is not a response nor a notification")
		}*/
}
