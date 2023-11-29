// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package RtpParameters

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type RtcpFeedbackT struct {
	Type string `json:"type"`
	Parameter string `json:"parameter"`
}

func (t *RtcpFeedbackT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	type_Offset := flatbuffers.UOffsetT(0)
	if t.Type != "" {
		type_Offset = builder.CreateString(t.Type)
	}
	parameterOffset := flatbuffers.UOffsetT(0)
	if t.Parameter != "" {
		parameterOffset = builder.CreateString(t.Parameter)
	}
	RtcpFeedbackStart(builder)
	RtcpFeedbackAddType(builder, type_Offset)
	RtcpFeedbackAddParameter(builder, parameterOffset)
	return RtcpFeedbackEnd(builder)
}

func (rcv *RtcpFeedback) UnPackTo(t *RtcpFeedbackT) {
	t.Type = string(rcv.Type())
	t.Parameter = string(rcv.Parameter())
}

func (rcv *RtcpFeedback) UnPack() *RtcpFeedbackT {
	if rcv == nil { return nil }
	t := &RtcpFeedbackT{}
	rcv.UnPackTo(t)
	return t
}

type RtcpFeedback struct {
	_tab flatbuffers.Table
}

func GetRootAsRtcpFeedback(buf []byte, offset flatbuffers.UOffsetT) *RtcpFeedback {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &RtcpFeedback{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsRtcpFeedback(buf []byte, offset flatbuffers.UOffsetT) *RtcpFeedback {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &RtcpFeedback{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *RtcpFeedback) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *RtcpFeedback) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *RtcpFeedback) Type() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *RtcpFeedback) Parameter() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func RtcpFeedbackStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func RtcpFeedbackAddType(builder *flatbuffers.Builder, type_ flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(type_), 0)
}
func RtcpFeedbackAddParameter(builder *flatbuffers.Builder, parameter flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(parameter), 0)
}
func RtcpFeedbackEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
