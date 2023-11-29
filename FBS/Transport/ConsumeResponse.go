// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Transport

import (
	flatbuffers "github.com/google/flatbuffers/go"

	FBS__Consumer "github.com/jiyeyuran/mediasoup-go/FBS/Consumer"
)

type ConsumeResponseT struct {
	Paused bool `json:"paused"`
	ProducerPaused bool `json:"producer_paused"`
	Score *FBS__Consumer.ConsumerScoreT `json:"score"`
	PreferredLayers *FBS__Consumer.ConsumerLayersT `json:"preferred_layers"`
}

func (t *ConsumeResponseT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	scoreOffset := t.Score.Pack(builder)
	preferredLayersOffset := t.PreferredLayers.Pack(builder)
	ConsumeResponseStart(builder)
	ConsumeResponseAddPaused(builder, t.Paused)
	ConsumeResponseAddProducerPaused(builder, t.ProducerPaused)
	ConsumeResponseAddScore(builder, scoreOffset)
	ConsumeResponseAddPreferredLayers(builder, preferredLayersOffset)
	return ConsumeResponseEnd(builder)
}

func (rcv *ConsumeResponse) UnPackTo(t *ConsumeResponseT) {
	t.Paused = rcv.Paused()
	t.ProducerPaused = rcv.ProducerPaused()
	t.Score = rcv.Score(nil).UnPack()
	t.PreferredLayers = rcv.PreferredLayers(nil).UnPack()
}

func (rcv *ConsumeResponse) UnPack() *ConsumeResponseT {
	if rcv == nil { return nil }
	t := &ConsumeResponseT{}
	rcv.UnPackTo(t)
	return t
}

type ConsumeResponse struct {
	_tab flatbuffers.Table
}

func GetRootAsConsumeResponse(buf []byte, offset flatbuffers.UOffsetT) *ConsumeResponse {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ConsumeResponse{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsConsumeResponse(buf []byte, offset flatbuffers.UOffsetT) *ConsumeResponse {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ConsumeResponse{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ConsumeResponse) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ConsumeResponse) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ConsumeResponse) Paused() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *ConsumeResponse) MutatePaused(n bool) bool {
	return rcv._tab.MutateBoolSlot(4, n)
}

func (rcv *ConsumeResponse) ProducerPaused() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *ConsumeResponse) MutateProducerPaused(n bool) bool {
	return rcv._tab.MutateBoolSlot(6, n)
}

func (rcv *ConsumeResponse) Score(obj *FBS__Consumer.ConsumerScore) *FBS__Consumer.ConsumerScore {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(FBS__Consumer.ConsumerScore)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *ConsumeResponse) PreferredLayers(obj *FBS__Consumer.ConsumerLayers) *FBS__Consumer.ConsumerLayers {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(FBS__Consumer.ConsumerLayers)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func ConsumeResponseStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func ConsumeResponseAddPaused(builder *flatbuffers.Builder, paused bool) {
	builder.PrependBoolSlot(0, paused, false)
}
func ConsumeResponseAddProducerPaused(builder *flatbuffers.Builder, producerPaused bool) {
	builder.PrependBoolSlot(1, producerPaused, false)
}
func ConsumeResponseAddScore(builder *flatbuffers.Builder, score flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(score), 0)
}
func ConsumeResponseAddPreferredLayers(builder *flatbuffers.Builder, preferredLayers flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(preferredLayers), 0)
}
func ConsumeResponseEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
