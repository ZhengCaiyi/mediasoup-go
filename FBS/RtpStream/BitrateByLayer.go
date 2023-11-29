// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package RtpStream

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type BitrateByLayerT struct {
	Layer string `json:"layer"`
	Bitrate uint32 `json:"bitrate"`
}

func (t *BitrateByLayerT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	layerOffset := flatbuffers.UOffsetT(0)
	if t.Layer != "" {
		layerOffset = builder.CreateString(t.Layer)
	}
	BitrateByLayerStart(builder)
	BitrateByLayerAddLayer(builder, layerOffset)
	BitrateByLayerAddBitrate(builder, t.Bitrate)
	return BitrateByLayerEnd(builder)
}

func (rcv *BitrateByLayer) UnPackTo(t *BitrateByLayerT) {
	t.Layer = string(rcv.Layer())
	t.Bitrate = rcv.Bitrate()
}

func (rcv *BitrateByLayer) UnPack() *BitrateByLayerT {
	if rcv == nil { return nil }
	t := &BitrateByLayerT{}
	rcv.UnPackTo(t)
	return t
}

type BitrateByLayer struct {
	_tab flatbuffers.Table
}

func GetRootAsBitrateByLayer(buf []byte, offset flatbuffers.UOffsetT) *BitrateByLayer {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &BitrateByLayer{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsBitrateByLayer(buf []byte, offset flatbuffers.UOffsetT) *BitrateByLayer {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &BitrateByLayer{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *BitrateByLayer) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *BitrateByLayer) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *BitrateByLayer) Layer() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *BitrateByLayer) Bitrate() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *BitrateByLayer) MutateBitrate(n uint32) bool {
	return rcv._tab.MutateUint32Slot(6, n)
}

func BitrateByLayerStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func BitrateByLayerAddLayer(builder *flatbuffers.Builder, layer flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(layer), 0)
}
func BitrateByLayerAddBitrate(builder *flatbuffers.Builder, bitrate uint32) {
	builder.PrependUint32Slot(1, bitrate, 0)
}
func BitrateByLayerEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
