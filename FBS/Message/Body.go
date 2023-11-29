// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Message

import (
	"strconv"
	flatbuffers "github.com/google/flatbuffers/go"

	FBS__Log "github.com/jiyeyuran/mediasoup-go/FBS/Log"
	FBS__Notification "github.com/jiyeyuran/mediasoup-go/FBS/Notification"
	FBS__Request "github.com/jiyeyuran/mediasoup-go/FBS/Request"
	FBS__Response "github.com/jiyeyuran/mediasoup-go/FBS/Response"
)

type Body byte

const (
	BodyNONE         Body = 0
	BodyRequest      Body = 1
	BodyResponse     Body = 2
	BodyNotification Body = 3
	BodyLog          Body = 4
)

var EnumNamesBody = map[Body]string{
	BodyNONE:         "NONE",
	BodyRequest:      "Request",
	BodyResponse:     "Response",
	BodyNotification: "Notification",
	BodyLog:          "Log",
}

var EnumValuesBody = map[string]Body{
	"NONE":         BodyNONE,
	"Request":      BodyRequest,
	"Response":     BodyResponse,
	"Notification": BodyNotification,
	"Log":          BodyLog,
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
	case BodyRequest:
		return t.Value.(*FBS__Request.RequestT).Pack(builder)
	case BodyResponse:
		return t.Value.(*FBS__Response.ResponseT).Pack(builder)
	case BodyNotification:
		return t.Value.(*FBS__Notification.NotificationT).Pack(builder)
	case BodyLog:
		return t.Value.(*FBS__Log.LogT).Pack(builder)
	}
	return 0
}

func (rcv Body) UnPack(table flatbuffers.Table) *BodyT {
	switch rcv {
	case BodyRequest:
		var x FBS__Request.Request
		x.Init(table.Bytes, table.Pos)
		return &BodyT{ Type: BodyRequest, Value: x.UnPack() }
	case BodyResponse:
		var x FBS__Response.Response
		x.Init(table.Bytes, table.Pos)
		return &BodyT{ Type: BodyResponse, Value: x.UnPack() }
	case BodyNotification:
		var x FBS__Notification.Notification
		x.Init(table.Bytes, table.Pos)
		return &BodyT{ Type: BodyNotification, Value: x.UnPack() }
	case BodyLog:
		var x FBS__Log.Log
		x.Init(table.Bytes, table.Pos)
		return &BodyT{ Type: BodyLog, Value: x.UnPack() }
	}
	return nil
}
