// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package RtpParameters

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type RtcpParametersT struct {
	Cname string `json:"cname"`
	ReducedSize bool `json:"reduced_size"`
}

func (t *RtcpParametersT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	cnameOffset := flatbuffers.UOffsetT(0)
	if t.Cname != "" {
		cnameOffset = builder.CreateString(t.Cname)
	}
	RtcpParametersStart(builder)
	RtcpParametersAddCname(builder, cnameOffset)
	RtcpParametersAddReducedSize(builder, t.ReducedSize)
	return RtcpParametersEnd(builder)
}

func (rcv *RtcpParameters) UnPackTo(t *RtcpParametersT) {
	t.Cname = string(rcv.Cname())
	t.ReducedSize = rcv.ReducedSize()
}

func (rcv *RtcpParameters) UnPack() *RtcpParametersT {
	if rcv == nil { return nil }
	t := &RtcpParametersT{}
	rcv.UnPackTo(t)
	return t
}

type RtcpParameters struct {
	_tab flatbuffers.Table
}

func GetRootAsRtcpParameters(buf []byte, offset flatbuffers.UOffsetT) *RtcpParameters {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &RtcpParameters{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsRtcpParameters(buf []byte, offset flatbuffers.UOffsetT) *RtcpParameters {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &RtcpParameters{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *RtcpParameters) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *RtcpParameters) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *RtcpParameters) Cname() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *RtcpParameters) ReducedSize() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return true
}

func (rcv *RtcpParameters) MutateReducedSize(n bool) bool {
	return rcv._tab.MutateBoolSlot(6, n)
}

func RtcpParametersStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func RtcpParametersAddCname(builder *flatbuffers.Builder, cname flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(cname), 0)
}
func RtcpParametersAddReducedSize(builder *flatbuffers.Builder, reducedSize bool) {
	builder.PrependBoolSlot(1, reducedSize, true)
}
func RtcpParametersEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
