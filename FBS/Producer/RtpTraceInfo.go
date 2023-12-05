// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Producer

import (
	flatbuffers "github.com/google/flatbuffers/go"

	FBS__RtpPacket "github.com/jiyeyuran/mediasoup-go/FBS/RtpPacket"
)

type RtpTraceInfoT struct {
	RtpPacket *FBS__RtpPacket.DumpT `json:"rtp_packet"`
	IsRtx bool `json:"is_rtx"`
}

func (t *RtpTraceInfoT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	rtpPacketOffset := t.RtpPacket.Pack(builder)
	RtpTraceInfoStart(builder)
	RtpTraceInfoAddRtpPacket(builder, rtpPacketOffset)
	RtpTraceInfoAddIsRtx(builder, t.IsRtx)
	return RtpTraceInfoEnd(builder)
}

func (rcv *RtpTraceInfo) UnPackTo(t *RtpTraceInfoT) {
	t.RtpPacket = rcv.RtpPacket(nil).UnPack()
	t.IsRtx = rcv.IsRtx()
}

func (rcv *RtpTraceInfo) UnPack() *RtpTraceInfoT {
	if rcv == nil { return nil }
	t := &RtpTraceInfoT{}
	rcv.UnPackTo(t)
	return t
}

type RtpTraceInfo struct {
	_tab flatbuffers.Table
}

func GetRootAsRtpTraceInfo(buf []byte, offset flatbuffers.UOffsetT) *RtpTraceInfo {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &RtpTraceInfo{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsRtpTraceInfo(buf []byte, offset flatbuffers.UOffsetT) *RtpTraceInfo {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &RtpTraceInfo{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *RtpTraceInfo) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *RtpTraceInfo) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *RtpTraceInfo) RtpPacket(obj *FBS__RtpPacket.Dump) *FBS__RtpPacket.Dump {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(FBS__RtpPacket.Dump)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *RtpTraceInfo) IsRtx() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *RtpTraceInfo) MutateIsRtx(n bool) bool {
	return rcv._tab.MutateBoolSlot(6, n)
}

func RtpTraceInfoStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func RtpTraceInfoAddRtpPacket(builder *flatbuffers.Builder, rtpPacket flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(rtpPacket), 0)
}
func RtpTraceInfoAddIsRtx(builder *flatbuffers.Builder, isRtx bool) {
	builder.PrependBoolSlot(1, isRtx, false)
}
func RtpTraceInfoEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}