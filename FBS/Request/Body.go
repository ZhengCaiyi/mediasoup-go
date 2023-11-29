// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Request

import (
	"strconv"
	flatbuffers "github.com/google/flatbuffers/go"

	FBS__Consumer "github.com/jiyeyuran/mediasoup-go/FBS/Consumer"
	FBS__DataConsumer "github.com/jiyeyuran/mediasoup-go/FBS/DataConsumer"
	FBS__PipeTransport "github.com/jiyeyuran/mediasoup-go/FBS/PipeTransport"
	FBS__PlainTransport "github.com/jiyeyuran/mediasoup-go/FBS/PlainTransport"
	FBS__Producer "github.com/jiyeyuran/mediasoup-go/FBS/Producer"
	FBS__Router "github.com/jiyeyuran/mediasoup-go/FBS/Router"
	FBS__RtpObserver "github.com/jiyeyuran/mediasoup-go/FBS/RtpObserver"
	FBS__Transport "github.com/jiyeyuran/mediasoup-go/FBS/Transport"
	FBS__WebRtcTransport "github.com/jiyeyuran/mediasoup-go/FBS/WebRtcTransport"
	FBS__Worker "github.com/jiyeyuran/mediasoup-go/FBS/Worker"
)

type Body byte

const (
	BodyNONE                                              Body = 0
	BodyWorker_UpdateSettingsRequest                      Body = 1
	BodyWorker_CreateWebRtcServerRequest                  Body = 2
	BodyWorker_CloseWebRtcServerRequest                   Body = 3
	BodyWorker_CreateRouterRequest                        Body = 4
	BodyWorker_CloseRouterRequest                         Body = 5
	BodyRouter_CreateWebRtcTransportRequest               Body = 6
	BodyRouter_CreatePlainTransportRequest                Body = 7
	BodyRouter_CreatePipeTransportRequest                 Body = 8
	BodyRouter_CreateDirectTransportRequest               Body = 9
	BodyRouter_CreateActiveSpeakerObserverRequest         Body = 10
	BodyRouter_CreateAudioLevelObserverRequest            Body = 11
	BodyRouter_CloseTransportRequest                      Body = 12
	BodyRouter_CloseRtpObserverRequest                    Body = 13
	BodyTransport_SetMaxIncomingBitrateRequest            Body = 14
	BodyTransport_SetMaxOutgoingBitrateRequest            Body = 15
	BodyTransport_SetMinOutgoingBitrateRequest            Body = 16
	BodyTransport_ProduceRequest                          Body = 17
	BodyTransport_ConsumeRequest                          Body = 18
	BodyTransport_ProduceDataRequest                      Body = 19
	BodyTransport_ConsumeDataRequest                      Body = 20
	BodyTransport_EnableTraceEventRequest                 Body = 21
	BodyTransport_CloseProducerRequest                    Body = 22
	BodyTransport_CloseConsumerRequest                    Body = 23
	BodyTransport_CloseDataProducerRequest                Body = 24
	BodyTransport_CloseDataConsumerRequest                Body = 25
	BodyPlainTransport_ConnectRequest                     Body = 26
	BodyPipeTransport_ConnectRequest                      Body = 27
	BodyWebRtcTransport_ConnectRequest                    Body = 28
	BodyProducer_EnableTraceEventRequest                  Body = 29
	BodyConsumer_SetPreferredLayersRequest                Body = 30
	BodyConsumer_SetPriorityRequest                       Body = 31
	BodyConsumer_EnableTraceEventRequest                  Body = 32
	BodyDataConsumer_SetBufferedAmountLowThresholdRequest Body = 33
	BodyDataConsumer_SendRequest                          Body = 34
	BodyDataConsumer_SetSubchannelsRequest                Body = 35
	BodyRtpObserver_AddProducerRequest                    Body = 36
	BodyRtpObserver_RemoveProducerRequest                 Body = 37
)

var EnumNamesBody = map[Body]string{
	BodyNONE:                                              "NONE",
	BodyWorker_UpdateSettingsRequest:                      "Worker_UpdateSettingsRequest",
	BodyWorker_CreateWebRtcServerRequest:                  "Worker_CreateWebRtcServerRequest",
	BodyWorker_CloseWebRtcServerRequest:                   "Worker_CloseWebRtcServerRequest",
	BodyWorker_CreateRouterRequest:                        "Worker_CreateRouterRequest",
	BodyWorker_CloseRouterRequest:                         "Worker_CloseRouterRequest",
	BodyRouter_CreateWebRtcTransportRequest:               "Router_CreateWebRtcTransportRequest",
	BodyRouter_CreatePlainTransportRequest:                "Router_CreatePlainTransportRequest",
	BodyRouter_CreatePipeTransportRequest:                 "Router_CreatePipeTransportRequest",
	BodyRouter_CreateDirectTransportRequest:               "Router_CreateDirectTransportRequest",
	BodyRouter_CreateActiveSpeakerObserverRequest:         "Router_CreateActiveSpeakerObserverRequest",
	BodyRouter_CreateAudioLevelObserverRequest:            "Router_CreateAudioLevelObserverRequest",
	BodyRouter_CloseTransportRequest:                      "Router_CloseTransportRequest",
	BodyRouter_CloseRtpObserverRequest:                    "Router_CloseRtpObserverRequest",
	BodyTransport_SetMaxIncomingBitrateRequest:            "Transport_SetMaxIncomingBitrateRequest",
	BodyTransport_SetMaxOutgoingBitrateRequest:            "Transport_SetMaxOutgoingBitrateRequest",
	BodyTransport_SetMinOutgoingBitrateRequest:            "Transport_SetMinOutgoingBitrateRequest",
	BodyTransport_ProduceRequest:                          "Transport_ProduceRequest",
	BodyTransport_ConsumeRequest:                          "Transport_ConsumeRequest",
	BodyTransport_ProduceDataRequest:                      "Transport_ProduceDataRequest",
	BodyTransport_ConsumeDataRequest:                      "Transport_ConsumeDataRequest",
	BodyTransport_EnableTraceEventRequest:                 "Transport_EnableTraceEventRequest",
	BodyTransport_CloseProducerRequest:                    "Transport_CloseProducerRequest",
	BodyTransport_CloseConsumerRequest:                    "Transport_CloseConsumerRequest",
	BodyTransport_CloseDataProducerRequest:                "Transport_CloseDataProducerRequest",
	BodyTransport_CloseDataConsumerRequest:                "Transport_CloseDataConsumerRequest",
	BodyPlainTransport_ConnectRequest:                     "PlainTransport_ConnectRequest",
	BodyPipeTransport_ConnectRequest:                      "PipeTransport_ConnectRequest",
	BodyWebRtcTransport_ConnectRequest:                    "WebRtcTransport_ConnectRequest",
	BodyProducer_EnableTraceEventRequest:                  "Producer_EnableTraceEventRequest",
	BodyConsumer_SetPreferredLayersRequest:                "Consumer_SetPreferredLayersRequest",
	BodyConsumer_SetPriorityRequest:                       "Consumer_SetPriorityRequest",
	BodyConsumer_EnableTraceEventRequest:                  "Consumer_EnableTraceEventRequest",
	BodyDataConsumer_SetBufferedAmountLowThresholdRequest: "DataConsumer_SetBufferedAmountLowThresholdRequest",
	BodyDataConsumer_SendRequest:                          "DataConsumer_SendRequest",
	BodyDataConsumer_SetSubchannelsRequest:                "DataConsumer_SetSubchannelsRequest",
	BodyRtpObserver_AddProducerRequest:                    "RtpObserver_AddProducerRequest",
	BodyRtpObserver_RemoveProducerRequest:                 "RtpObserver_RemoveProducerRequest",
}

var EnumValuesBody = map[string]Body{
	"NONE":                                              BodyNONE,
	"Worker_UpdateSettingsRequest":                      BodyWorker_UpdateSettingsRequest,
	"Worker_CreateWebRtcServerRequest":                  BodyWorker_CreateWebRtcServerRequest,
	"Worker_CloseWebRtcServerRequest":                   BodyWorker_CloseWebRtcServerRequest,
	"Worker_CreateRouterRequest":                        BodyWorker_CreateRouterRequest,
	"Worker_CloseRouterRequest":                         BodyWorker_CloseRouterRequest,
	"Router_CreateWebRtcTransportRequest":               BodyRouter_CreateWebRtcTransportRequest,
	"Router_CreatePlainTransportRequest":                BodyRouter_CreatePlainTransportRequest,
	"Router_CreatePipeTransportRequest":                 BodyRouter_CreatePipeTransportRequest,
	"Router_CreateDirectTransportRequest":               BodyRouter_CreateDirectTransportRequest,
	"Router_CreateActiveSpeakerObserverRequest":         BodyRouter_CreateActiveSpeakerObserverRequest,
	"Router_CreateAudioLevelObserverRequest":            BodyRouter_CreateAudioLevelObserverRequest,
	"Router_CloseTransportRequest":                      BodyRouter_CloseTransportRequest,
	"Router_CloseRtpObserverRequest":                    BodyRouter_CloseRtpObserverRequest,
	"Transport_SetMaxIncomingBitrateRequest":            BodyTransport_SetMaxIncomingBitrateRequest,
	"Transport_SetMaxOutgoingBitrateRequest":            BodyTransport_SetMaxOutgoingBitrateRequest,
	"Transport_SetMinOutgoingBitrateRequest":            BodyTransport_SetMinOutgoingBitrateRequest,
	"Transport_ProduceRequest":                          BodyTransport_ProduceRequest,
	"Transport_ConsumeRequest":                          BodyTransport_ConsumeRequest,
	"Transport_ProduceDataRequest":                      BodyTransport_ProduceDataRequest,
	"Transport_ConsumeDataRequest":                      BodyTransport_ConsumeDataRequest,
	"Transport_EnableTraceEventRequest":                 BodyTransport_EnableTraceEventRequest,
	"Transport_CloseProducerRequest":                    BodyTransport_CloseProducerRequest,
	"Transport_CloseConsumerRequest":                    BodyTransport_CloseConsumerRequest,
	"Transport_CloseDataProducerRequest":                BodyTransport_CloseDataProducerRequest,
	"Transport_CloseDataConsumerRequest":                BodyTransport_CloseDataConsumerRequest,
	"PlainTransport_ConnectRequest":                     BodyPlainTransport_ConnectRequest,
	"PipeTransport_ConnectRequest":                      BodyPipeTransport_ConnectRequest,
	"WebRtcTransport_ConnectRequest":                    BodyWebRtcTransport_ConnectRequest,
	"Producer_EnableTraceEventRequest":                  BodyProducer_EnableTraceEventRequest,
	"Consumer_SetPreferredLayersRequest":                BodyConsumer_SetPreferredLayersRequest,
	"Consumer_SetPriorityRequest":                       BodyConsumer_SetPriorityRequest,
	"Consumer_EnableTraceEventRequest":                  BodyConsumer_EnableTraceEventRequest,
	"DataConsumer_SetBufferedAmountLowThresholdRequest": BodyDataConsumer_SetBufferedAmountLowThresholdRequest,
	"DataConsumer_SendRequest":                          BodyDataConsumer_SendRequest,
	"DataConsumer_SetSubchannelsRequest":                BodyDataConsumer_SetSubchannelsRequest,
	"RtpObserver_AddProducerRequest":                    BodyRtpObserver_AddProducerRequest,
	"RtpObserver_RemoveProducerRequest":                 BodyRtpObserver_RemoveProducerRequest,
}

func (v Body) String() string {
	if s, ok := EnumNamesBody[v]; ok {
		return s
	}
	return "Body(" + strconv.FormatInt(int64(v), 10) + ")"
}

type BodyT struct {
	Type Body
	Value interface{}
}

func (t *BodyT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	switch t.Type {
	case BodyWorker_UpdateSettingsRequest:
		return t.Value.(*FBS__Worker.UpdateSettingsRequestT).Pack(builder)
	case BodyWorker_CreateWebRtcServerRequest:
		return t.Value.(*FBS__Worker.CreateWebRtcServerRequestT).Pack(builder)
	case BodyWorker_CloseWebRtcServerRequest:
		return t.Value.(*FBS__Worker.CloseWebRtcServerRequestT).Pack(builder)
	case BodyWorker_CreateRouterRequest:
		return t.Value.(*FBS__Worker.CreateRouterRequestT).Pack(builder)
	case BodyWorker_CloseRouterRequest:
		return t.Value.(*FBS__Worker.CloseRouterRequestT).Pack(builder)
	case BodyRouter_CreateWebRtcTransportRequest:
		return t.Value.(*FBS__Router.CreateWebRtcTransportRequestT).Pack(builder)
	case BodyRouter_CreatePlainTransportRequest:
		return t.Value.(*FBS__Router.CreatePlainTransportRequestT).Pack(builder)
	case BodyRouter_CreatePipeTransportRequest:
		return t.Value.(*FBS__Router.CreatePipeTransportRequestT).Pack(builder)
	case BodyRouter_CreateDirectTransportRequest:
		return t.Value.(*FBS__Router.CreateDirectTransportRequestT).Pack(builder)
	case BodyRouter_CreateActiveSpeakerObserverRequest:
		return t.Value.(*FBS__Router.CreateActiveSpeakerObserverRequestT).Pack(builder)
	case BodyRouter_CreateAudioLevelObserverRequest:
		return t.Value.(*FBS__Router.CreateAudioLevelObserverRequestT).Pack(builder)
	case BodyRouter_CloseTransportRequest:
		return t.Value.(*FBS__Router.CloseTransportRequestT).Pack(builder)
	case BodyRouter_CloseRtpObserverRequest:
		return t.Value.(*FBS__Router.CloseRtpObserverRequestT).Pack(builder)
	case BodyTransport_SetMaxIncomingBitrateRequest:
		return t.Value.(*FBS__Transport.SetMaxIncomingBitrateRequestT).Pack(builder)
	case BodyTransport_SetMaxOutgoingBitrateRequest:
		return t.Value.(*FBS__Transport.SetMaxOutgoingBitrateRequestT).Pack(builder)
	case BodyTransport_SetMinOutgoingBitrateRequest:
		return t.Value.(*FBS__Transport.SetMinOutgoingBitrateRequestT).Pack(builder)
	case BodyTransport_ProduceRequest:
		return t.Value.(*FBS__Transport.ProduceRequestT).Pack(builder)
	case BodyTransport_ConsumeRequest:
		return t.Value.(*FBS__Transport.ConsumeRequestT).Pack(builder)
	case BodyTransport_ProduceDataRequest:
		return t.Value.(*FBS__Transport.ProduceDataRequestT).Pack(builder)
	case BodyTransport_ConsumeDataRequest:
		return t.Value.(*FBS__Transport.ConsumeDataRequestT).Pack(builder)
	case BodyTransport_EnableTraceEventRequest:
		return t.Value.(*FBS__Transport.EnableTraceEventRequestT).Pack(builder)
	case BodyTransport_CloseProducerRequest:
		return t.Value.(*FBS__Transport.CloseProducerRequestT).Pack(builder)
	case BodyTransport_CloseConsumerRequest:
		return t.Value.(*FBS__Transport.CloseConsumerRequestT).Pack(builder)
	case BodyTransport_CloseDataProducerRequest:
		return t.Value.(*FBS__Transport.CloseDataProducerRequestT).Pack(builder)
	case BodyTransport_CloseDataConsumerRequest:
		return t.Value.(*FBS__Transport.CloseDataConsumerRequestT).Pack(builder)
	case BodyPlainTransport_ConnectRequest:
		return t.Value.(*FBS__PlainTransport.ConnectRequestT).Pack(builder)
	case BodyPipeTransport_ConnectRequest:
		return t.Value.(*FBS__PipeTransport.ConnectRequestT).Pack(builder)
	case BodyWebRtcTransport_ConnectRequest:
		return t.Value.(*FBS__WebRtcTransport.ConnectRequestT).Pack(builder)
	case BodyProducer_EnableTraceEventRequest:
		return t.Value.(*FBS__Producer.EnableTraceEventRequestT).Pack(builder)
	case BodyConsumer_SetPreferredLayersRequest:
		return t.Value.(*FBS__Consumer.SetPreferredLayersRequestT).Pack(builder)
	case BodyConsumer_SetPriorityRequest:
		return t.Value.(*FBS__Consumer.SetPriorityRequestT).Pack(builder)
	case BodyConsumer_EnableTraceEventRequest:
		return t.Value.(*FBS__Consumer.EnableTraceEventRequestT).Pack(builder)
	case BodyDataConsumer_SetBufferedAmountLowThresholdRequest:
		return t.Value.(*FBS__DataConsumer.SetBufferedAmountLowThresholdRequestT).Pack(builder)
	case BodyDataConsumer_SendRequest:
		return t.Value.(*FBS__DataConsumer.SendRequestT).Pack(builder)
	case BodyDataConsumer_SetSubchannelsRequest:
		return t.Value.(*FBS__DataConsumer.SetSubchannelsRequestT).Pack(builder)
	case BodyRtpObserver_AddProducerRequest:
		return t.Value.(*FBS__RtpObserver.AddProducerRequestT).Pack(builder)
	case BodyRtpObserver_RemoveProducerRequest:
		return t.Value.(*FBS__RtpObserver.RemoveProducerRequestT).Pack(builder)
	}
	return 0
}

func (rcv Body) UnPack(table flatbuffers.Table) *BodyT {
	switch rcv {
	case BodyWorker_UpdateSettingsRequest:
		var x FBS__Worker.UpdateSettingsRequest
		x.Init(table.Bytes, table.Pos)
		return &BodyT{ Type: BodyWorker_UpdateSettingsRequest, Value: x.UnPack() }
	case BodyWorker_CreateWebRtcServerRequest:
		var x FBS__Worker.CreateWebRtcServerRequest
		x.Init(table.Bytes, table.Pos)
		return &BodyT{ Type: BodyWorker_CreateWebRtcServerRequest, Value: x.UnPack() }
	case BodyWorker_CloseWebRtcServerRequest:
		var x FBS__Worker.CloseWebRtcServerRequest
		x.Init(table.Bytes, table.Pos)
		return &BodyT{ Type: BodyWorker_CloseWebRtcServerRequest, Value: x.UnPack() }
	case BodyWorker_CreateRouterRequest:
		var x FBS__Worker.CreateRouterRequest
		x.Init(table.Bytes, table.Pos)
		return &BodyT{ Type: BodyWorker_CreateRouterRequest, Value: x.UnPack() }
	case BodyWorker_CloseRouterRequest:
		var x FBS__Worker.CloseRouterRequest
		x.Init(table.Bytes, table.Pos)
		return &BodyT{ Type: BodyWorker_CloseRouterRequest, Value: x.UnPack() }
	case BodyRouter_CreateWebRtcTransportRequest:
		var x FBS__Router.CreateWebRtcTransportRequest
		x.Init(table.Bytes, table.Pos)
		return &BodyT{ Type: BodyRouter_CreateWebRtcTransportRequest, Value: x.UnPack() }
	case BodyRouter_CreatePlainTransportRequest:
		var x FBS__Router.CreatePlainTransportRequest
		x.Init(table.Bytes, table.Pos)
		return &BodyT{ Type: BodyRouter_CreatePlainTransportRequest, Value: x.UnPack() }
	case BodyRouter_CreatePipeTransportRequest:
		var x FBS__Router.CreatePipeTransportRequest
		x.Init(table.Bytes, table.Pos)
		return &BodyT{ Type: BodyRouter_CreatePipeTransportRequest, Value: x.UnPack() }
	case BodyRouter_CreateDirectTransportRequest:
		var x FBS__Router.CreateDirectTransportRequest
		x.Init(table.Bytes, table.Pos)
		return &BodyT{ Type: BodyRouter_CreateDirectTransportRequest, Value: x.UnPack() }
	case BodyRouter_CreateActiveSpeakerObserverRequest:
		var x FBS__Router.CreateActiveSpeakerObserverRequest
		x.Init(table.Bytes, table.Pos)
		return &BodyT{ Type: BodyRouter_CreateActiveSpeakerObserverRequest, Value: x.UnPack() }
	case BodyRouter_CreateAudioLevelObserverRequest:
		var x FBS__Router.CreateAudioLevelObserverRequest
		x.Init(table.Bytes, table.Pos)
		return &BodyT{ Type: BodyRouter_CreateAudioLevelObserverRequest, Value: x.UnPack() }
	case BodyRouter_CloseTransportRequest:
		var x FBS__Router.CloseTransportRequest
		x.Init(table.Bytes, table.Pos)
		return &BodyT{ Type: BodyRouter_CloseTransportRequest, Value: x.UnPack() }
	case BodyRouter_CloseRtpObserverRequest:
		var x FBS__Router.CloseRtpObserverRequest
		x.Init(table.Bytes, table.Pos)
		return &BodyT{ Type: BodyRouter_CloseRtpObserverRequest, Value: x.UnPack() }
	case BodyTransport_SetMaxIncomingBitrateRequest:
		var x FBS__Transport.SetMaxIncomingBitrateRequest
		x.Init(table.Bytes, table.Pos)
		return &BodyT{ Type: BodyTransport_SetMaxIncomingBitrateRequest, Value: x.UnPack() }
	case BodyTransport_SetMaxOutgoingBitrateRequest:
		var x FBS__Transport.SetMaxOutgoingBitrateRequest
		x.Init(table.Bytes, table.Pos)
		return &BodyT{ Type: BodyTransport_SetMaxOutgoingBitrateRequest, Value: x.UnPack() }
	case BodyTransport_SetMinOutgoingBitrateRequest:
		var x FBS__Transport.SetMinOutgoingBitrateRequest
		x.Init(table.Bytes, table.Pos)
		return &BodyT{ Type: BodyTransport_SetMinOutgoingBitrateRequest, Value: x.UnPack() }
	case BodyTransport_ProduceRequest:
		var x FBS__Transport.ProduceRequest
		x.Init(table.Bytes, table.Pos)
		return &BodyT{ Type: BodyTransport_ProduceRequest, Value: x.UnPack() }
	case BodyTransport_ConsumeRequest:
		var x FBS__Transport.ConsumeRequest
		x.Init(table.Bytes, table.Pos)
		return &BodyT{ Type: BodyTransport_ConsumeRequest, Value: x.UnPack() }
	case BodyTransport_ProduceDataRequest:
		var x FBS__Transport.ProduceDataRequest
		x.Init(table.Bytes, table.Pos)
		return &BodyT{ Type: BodyTransport_ProduceDataRequest, Value: x.UnPack() }
	case BodyTransport_ConsumeDataRequest:
		var x FBS__Transport.ConsumeDataRequest
		x.Init(table.Bytes, table.Pos)
		return &BodyT{ Type: BodyTransport_ConsumeDataRequest, Value: x.UnPack() }
	case BodyTransport_EnableTraceEventRequest:
		var x FBS__Transport.EnableTraceEventRequest
		x.Init(table.Bytes, table.Pos)
		return &BodyT{ Type: BodyTransport_EnableTraceEventRequest, Value: x.UnPack() }
	case BodyTransport_CloseProducerRequest:
		var x FBS__Transport.CloseProducerRequest
		x.Init(table.Bytes, table.Pos)
		return &BodyT{ Type: BodyTransport_CloseProducerRequest, Value: x.UnPack() }
	case BodyTransport_CloseConsumerRequest:
		var x FBS__Transport.CloseConsumerRequest
		x.Init(table.Bytes, table.Pos)
		return &BodyT{ Type: BodyTransport_CloseConsumerRequest, Value: x.UnPack() }
	case BodyTransport_CloseDataProducerRequest:
		var x FBS__Transport.CloseDataProducerRequest
		x.Init(table.Bytes, table.Pos)
		return &BodyT{ Type: BodyTransport_CloseDataProducerRequest, Value: x.UnPack() }
	case BodyTransport_CloseDataConsumerRequest:
		var x FBS__Transport.CloseDataConsumerRequest
		x.Init(table.Bytes, table.Pos)
		return &BodyT{ Type: BodyTransport_CloseDataConsumerRequest, Value: x.UnPack() }
	case BodyPlainTransport_ConnectRequest:
		var x FBS__PlainTransport.ConnectRequest
		x.Init(table.Bytes, table.Pos)
		return &BodyT{ Type: BodyPlainTransport_ConnectRequest, Value: x.UnPack() }
	case BodyPipeTransport_ConnectRequest:
		var x FBS__PipeTransport.ConnectRequest
		x.Init(table.Bytes, table.Pos)
		return &BodyT{ Type: BodyPipeTransport_ConnectRequest, Value: x.UnPack() }
	case BodyWebRtcTransport_ConnectRequest:
		var x FBS__WebRtcTransport.ConnectRequest
		x.Init(table.Bytes, table.Pos)
		return &BodyT{ Type: BodyWebRtcTransport_ConnectRequest, Value: x.UnPack() }
	case BodyProducer_EnableTraceEventRequest:
		var x FBS__Producer.EnableTraceEventRequest
		x.Init(table.Bytes, table.Pos)
		return &BodyT{ Type: BodyProducer_EnableTraceEventRequest, Value: x.UnPack() }
	case BodyConsumer_SetPreferredLayersRequest:
		var x FBS__Consumer.SetPreferredLayersRequest
		x.Init(table.Bytes, table.Pos)
		return &BodyT{ Type: BodyConsumer_SetPreferredLayersRequest, Value: x.UnPack() }
	case BodyConsumer_SetPriorityRequest:
		var x FBS__Consumer.SetPriorityRequest
		x.Init(table.Bytes, table.Pos)
		return &BodyT{ Type: BodyConsumer_SetPriorityRequest, Value: x.UnPack() }
	case BodyConsumer_EnableTraceEventRequest:
		var x FBS__Consumer.EnableTraceEventRequest
		x.Init(table.Bytes, table.Pos)
		return &BodyT{ Type: BodyConsumer_EnableTraceEventRequest, Value: x.UnPack() }
	case BodyDataConsumer_SetBufferedAmountLowThresholdRequest:
		var x FBS__DataConsumer.SetBufferedAmountLowThresholdRequest
		x.Init(table.Bytes, table.Pos)
		return &BodyT{ Type: BodyDataConsumer_SetBufferedAmountLowThresholdRequest, Value: x.UnPack() }
	case BodyDataConsumer_SendRequest:
		var x FBS__DataConsumer.SendRequest
		x.Init(table.Bytes, table.Pos)
		return &BodyT{ Type: BodyDataConsumer_SendRequest, Value: x.UnPack() }
	case BodyDataConsumer_SetSubchannelsRequest:
		var x FBS__DataConsumer.SetSubchannelsRequest
		x.Init(table.Bytes, table.Pos)
		return &BodyT{ Type: BodyDataConsumer_SetSubchannelsRequest, Value: x.UnPack() }
	case BodyRtpObserver_AddProducerRequest:
		var x FBS__RtpObserver.AddProducerRequest
		x.Init(table.Bytes, table.Pos)
		return &BodyT{ Type: BodyRtpObserver_AddProducerRequest, Value: x.UnPack() }
	case BodyRtpObserver_RemoveProducerRequest:
		var x FBS__RtpObserver.RemoveProducerRequest
		x.Init(table.Bytes, table.Pos)
		return &BodyT{ Type: BodyRtpObserver_RemoveProducerRequest, Value: x.UnPack() }
	}
	return nil
}
