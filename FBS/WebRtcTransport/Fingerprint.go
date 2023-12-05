// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package WebRtcTransport

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type FingerprintT struct {
	Algorithm FingerprintAlgorithm `json:"algorithm"`
	Value string `json:"value"`
}

func (t *FingerprintT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	valueOffset := flatbuffers.UOffsetT(0)
	if t.Value != "" {
		valueOffset = builder.CreateString(t.Value)
	}
	FingerprintStart(builder)
	FingerprintAddAlgorithm(builder, t.Algorithm)
	FingerprintAddValue(builder, valueOffset)
	return FingerprintEnd(builder)
}

func (rcv *Fingerprint) UnPackTo(t *FingerprintT) {
	t.Algorithm = rcv.Algorithm()
	t.Value = string(rcv.Value())
}

func (rcv *Fingerprint) UnPack() *FingerprintT {
	if rcv == nil { return nil }
	t := &FingerprintT{}
	rcv.UnPackTo(t)
	return t
}

type Fingerprint struct {
	_tab flatbuffers.Table
}

func GetRootAsFingerprint(buf []byte, offset flatbuffers.UOffsetT) *Fingerprint {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Fingerprint{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsFingerprint(buf []byte, offset flatbuffers.UOffsetT) *Fingerprint {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Fingerprint{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Fingerprint) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Fingerprint) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Fingerprint) Algorithm() FingerprintAlgorithm {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return FingerprintAlgorithm(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *Fingerprint) MutateAlgorithm(n FingerprintAlgorithm) bool {
	return rcv._tab.MutateByteSlot(4, byte(n))
}

func (rcv *Fingerprint) Value() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func FingerprintStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func FingerprintAddAlgorithm(builder *flatbuffers.Builder, algorithm FingerprintAlgorithm) {
	builder.PrependByteSlot(0, byte(algorithm), 0)
}
func FingerprintAddValue(builder *flatbuffers.Builder, value flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(value), 0)
}
func FingerprintEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}