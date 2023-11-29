// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package WebRtcServer

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type TupleHashT struct {
	TupleHash uint64 `json:"tuple_hash"`
	WebRtcTransportId string `json:"web_rtc_transport_id"`
}

func (t *TupleHashT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	webRtcTransportIdOffset := flatbuffers.UOffsetT(0)
	if t.WebRtcTransportId != "" {
		webRtcTransportIdOffset = builder.CreateString(t.WebRtcTransportId)
	}
	TupleHashStart(builder)
	TupleHashAddTupleHash(builder, t.TupleHash)
	TupleHashAddWebRtcTransportId(builder, webRtcTransportIdOffset)
	return TupleHashEnd(builder)
}

func (rcv *TupleHash) UnPackTo(t *TupleHashT) {
	t.TupleHash = rcv.TupleHash()
	t.WebRtcTransportId = string(rcv.WebRtcTransportId())
}

func (rcv *TupleHash) UnPack() *TupleHashT {
	if rcv == nil { return nil }
	t := &TupleHashT{}
	rcv.UnPackTo(t)
	return t
}

type TupleHash struct {
	_tab flatbuffers.Table
}

func GetRootAsTupleHash(buf []byte, offset flatbuffers.UOffsetT) *TupleHash {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &TupleHash{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsTupleHash(buf []byte, offset flatbuffers.UOffsetT) *TupleHash {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &TupleHash{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *TupleHash) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *TupleHash) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *TupleHash) TupleHash() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *TupleHash) MutateTupleHash(n uint64) bool {
	return rcv._tab.MutateUint64Slot(4, n)
}

func (rcv *TupleHash) WebRtcTransportId() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func TupleHashStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func TupleHashAddTupleHash(builder *flatbuffers.Builder, tupleHash uint64) {
	builder.PrependUint64Slot(0, tupleHash, 0)
}
func TupleHashAddWebRtcTransportId(builder *flatbuffers.Builder, webRtcTransportId flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(webRtcTransportId), 0)
}
func TupleHashEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
