// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Producer

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type SendNotificationT struct {
	Data []byte `json:"data"`
}

func (t *SendNotificationT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	dataOffset := flatbuffers.UOffsetT(0)
	if t.Data != nil {
		dataOffset = builder.CreateByteString(t.Data)
	}
	SendNotificationStart(builder)
	SendNotificationAddData(builder, dataOffset)
	return SendNotificationEnd(builder)
}

func (rcv *SendNotification) UnPackTo(t *SendNotificationT) {
	t.Data = rcv.DataBytes()
}

func (rcv *SendNotification) UnPack() *SendNotificationT {
	if rcv == nil { return nil }
	t := &SendNotificationT{}
	rcv.UnPackTo(t)
	return t
}

type SendNotification struct {
	_tab flatbuffers.Table
}

func GetRootAsSendNotification(buf []byte, offset flatbuffers.UOffsetT) *SendNotification {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &SendNotification{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsSendNotification(buf []byte, offset flatbuffers.UOffsetT) *SendNotification {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &SendNotification{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *SendNotification) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SendNotification) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *SendNotification) Data(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *SendNotification) DataLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *SendNotification) DataBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *SendNotification) MutateData(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func SendNotificationStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func SendNotificationAddData(builder *flatbuffers.Builder, data flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(data), 0)
}
func SendNotificationStartDataVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func SendNotificationEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
