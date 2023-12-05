// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package DataConsumer

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type BufferedAmountLowNotificationT struct {
	BufferedAmount uint32 `json:"buffered_amount"`
}

func (t *BufferedAmountLowNotificationT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	BufferedAmountLowNotificationStart(builder)
	BufferedAmountLowNotificationAddBufferedAmount(builder, t.BufferedAmount)
	return BufferedAmountLowNotificationEnd(builder)
}

func (rcv *BufferedAmountLowNotification) UnPackTo(t *BufferedAmountLowNotificationT) {
	t.BufferedAmount = rcv.BufferedAmount()
}

func (rcv *BufferedAmountLowNotification) UnPack() *BufferedAmountLowNotificationT {
	if rcv == nil { return nil }
	t := &BufferedAmountLowNotificationT{}
	rcv.UnPackTo(t)
	return t
}

type BufferedAmountLowNotification struct {
	_tab flatbuffers.Table
}

func GetRootAsBufferedAmountLowNotification(buf []byte, offset flatbuffers.UOffsetT) *BufferedAmountLowNotification {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &BufferedAmountLowNotification{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsBufferedAmountLowNotification(buf []byte, offset flatbuffers.UOffsetT) *BufferedAmountLowNotification {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &BufferedAmountLowNotification{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *BufferedAmountLowNotification) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *BufferedAmountLowNotification) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *BufferedAmountLowNotification) BufferedAmount() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *BufferedAmountLowNotification) MutateBufferedAmount(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func BufferedAmountLowNotificationStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func BufferedAmountLowNotificationAddBufferedAmount(builder *flatbuffers.Builder, bufferedAmount uint32) {
	builder.PrependUint32Slot(0, bufferedAmount, 0)
}
func BufferedAmountLowNotificationEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}