// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Transport

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type SendRtcpNotificationT struct {
	Data []byte `json:"data"`
}

func (t *SendRtcpNotificationT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	dataOffset := flatbuffers.UOffsetT(0)
	if t.Data != nil {
		dataOffset = builder.CreateByteString(t.Data)
	}
	SendRtcpNotificationStart(builder)
	SendRtcpNotificationAddData(builder, dataOffset)
	return SendRtcpNotificationEnd(builder)
}

func (rcv *SendRtcpNotification) UnPackTo(t *SendRtcpNotificationT) {
	t.Data = rcv.DataBytes()
}

func (rcv *SendRtcpNotification) UnPack() *SendRtcpNotificationT {
	if rcv == nil { return nil }
	t := &SendRtcpNotificationT{}
	rcv.UnPackTo(t)
	return t
}

type SendRtcpNotification struct {
	_tab flatbuffers.Table
}

func GetRootAsSendRtcpNotification(buf []byte, offset flatbuffers.UOffsetT) *SendRtcpNotification {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &SendRtcpNotification{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsSendRtcpNotification(buf []byte, offset flatbuffers.UOffsetT) *SendRtcpNotification {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &SendRtcpNotification{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *SendRtcpNotification) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SendRtcpNotification) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *SendRtcpNotification) Data(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *SendRtcpNotification) DataLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *SendRtcpNotification) DataBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *SendRtcpNotification) MutateData(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func SendRtcpNotificationStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func SendRtcpNotificationAddData(builder *flatbuffers.Builder, data flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(data), 0)
}
func SendRtcpNotificationStartDataVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func SendRtcpNotificationEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
