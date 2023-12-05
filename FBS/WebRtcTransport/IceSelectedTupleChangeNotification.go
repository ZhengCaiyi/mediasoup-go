// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package WebRtcTransport

import (
	flatbuffers "github.com/google/flatbuffers/go"

	FBS__Transport "github.com/jiyeyuran/mediasoup-go/FBS/Transport"
)

type IceSelectedTupleChangeNotificationT struct {
	Tuple *FBS__Transport.TupleT `json:"tuple"`
}

func (t *IceSelectedTupleChangeNotificationT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	tupleOffset := t.Tuple.Pack(builder)
	IceSelectedTupleChangeNotificationStart(builder)
	IceSelectedTupleChangeNotificationAddTuple(builder, tupleOffset)
	return IceSelectedTupleChangeNotificationEnd(builder)
}

func (rcv *IceSelectedTupleChangeNotification) UnPackTo(t *IceSelectedTupleChangeNotificationT) {
	t.Tuple = rcv.Tuple(nil).UnPack()
}

func (rcv *IceSelectedTupleChangeNotification) UnPack() *IceSelectedTupleChangeNotificationT {
	if rcv == nil { return nil }
	t := &IceSelectedTupleChangeNotificationT{}
	rcv.UnPackTo(t)
	return t
}

type IceSelectedTupleChangeNotification struct {
	_tab flatbuffers.Table
}

func GetRootAsIceSelectedTupleChangeNotification(buf []byte, offset flatbuffers.UOffsetT) *IceSelectedTupleChangeNotification {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &IceSelectedTupleChangeNotification{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsIceSelectedTupleChangeNotification(buf []byte, offset flatbuffers.UOffsetT) *IceSelectedTupleChangeNotification {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &IceSelectedTupleChangeNotification{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *IceSelectedTupleChangeNotification) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *IceSelectedTupleChangeNotification) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *IceSelectedTupleChangeNotification) Tuple(obj *FBS__Transport.Tuple) *FBS__Transport.Tuple {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(FBS__Transport.Tuple)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func IceSelectedTupleChangeNotificationStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func IceSelectedTupleChangeNotificationAddTuple(builder *flatbuffers.Builder, tuple flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(tuple), 0)
}
func IceSelectedTupleChangeNotificationEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}