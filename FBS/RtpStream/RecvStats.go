// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package RtpStream

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type RecvStatsT struct {
	Base *StatsT `json:"base"`
	Jitter uint32 `json:"jitter"`
	PacketCount uint64 `json:"packet_count"`
	ByteCount uint64 `json:"byte_count"`
	Bitrate uint32 `json:"bitrate"`
	BitrateByLayer []*BitrateByLayerT `json:"bitrate_by_layer"`
}

func (t *RecvStatsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	baseOffset := t.Base.Pack(builder)
	bitrateByLayerOffset := flatbuffers.UOffsetT(0)
	if t.BitrateByLayer != nil {
		bitrateByLayerLength := len(t.BitrateByLayer)
		bitrateByLayerOffsets := make([]flatbuffers.UOffsetT, bitrateByLayerLength)
		for j := 0; j < bitrateByLayerLength; j++ {
			bitrateByLayerOffsets[j] = t.BitrateByLayer[j].Pack(builder)
		}
		RecvStatsStartBitrateByLayerVector(builder, bitrateByLayerLength)
		for j := bitrateByLayerLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(bitrateByLayerOffsets[j])
		}
		bitrateByLayerOffset = builder.EndVector(bitrateByLayerLength)
	}
	RecvStatsStart(builder)
	RecvStatsAddBase(builder, baseOffset)
	RecvStatsAddJitter(builder, t.Jitter)
	RecvStatsAddPacketCount(builder, t.PacketCount)
	RecvStatsAddByteCount(builder, t.ByteCount)
	RecvStatsAddBitrate(builder, t.Bitrate)
	RecvStatsAddBitrateByLayer(builder, bitrateByLayerOffset)
	return RecvStatsEnd(builder)
}

func (rcv *RecvStats) UnPackTo(t *RecvStatsT) {
	t.Base = rcv.Base(nil).UnPack()
	t.Jitter = rcv.Jitter()
	t.PacketCount = rcv.PacketCount()
	t.ByteCount = rcv.ByteCount()
	t.Bitrate = rcv.Bitrate()
	bitrateByLayerLength := rcv.BitrateByLayerLength()
	t.BitrateByLayer = make([]*BitrateByLayerT, bitrateByLayerLength)
	for j := 0; j < bitrateByLayerLength; j++ {
		x := BitrateByLayer{}
		rcv.BitrateByLayer(&x, j)
		t.BitrateByLayer[j] = x.UnPack()
	}
}

func (rcv *RecvStats) UnPack() *RecvStatsT {
	if rcv == nil { return nil }
	t := &RecvStatsT{}
	rcv.UnPackTo(t)
	return t
}

type RecvStats struct {
	_tab flatbuffers.Table
}

func GetRootAsRecvStats(buf []byte, offset flatbuffers.UOffsetT) *RecvStats {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &RecvStats{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsRecvStats(buf []byte, offset flatbuffers.UOffsetT) *RecvStats {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &RecvStats{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *RecvStats) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *RecvStats) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *RecvStats) Base(obj *Stats) *Stats {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Stats)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *RecvStats) Jitter() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *RecvStats) MutateJitter(n uint32) bool {
	return rcv._tab.MutateUint32Slot(6, n)
}

func (rcv *RecvStats) PacketCount() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *RecvStats) MutatePacketCount(n uint64) bool {
	return rcv._tab.MutateUint64Slot(8, n)
}

func (rcv *RecvStats) ByteCount() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *RecvStats) MutateByteCount(n uint64) bool {
	return rcv._tab.MutateUint64Slot(10, n)
}

func (rcv *RecvStats) Bitrate() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *RecvStats) MutateBitrate(n uint32) bool {
	return rcv._tab.MutateUint32Slot(12, n)
}

func (rcv *RecvStats) BitrateByLayer(obj *BitrateByLayer, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *RecvStats) BitrateByLayerLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func RecvStatsStart(builder *flatbuffers.Builder) {
	builder.StartObject(6)
}
func RecvStatsAddBase(builder *flatbuffers.Builder, base flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(base), 0)
}
func RecvStatsAddJitter(builder *flatbuffers.Builder, jitter uint32) {
	builder.PrependUint32Slot(1, jitter, 0)
}
func RecvStatsAddPacketCount(builder *flatbuffers.Builder, packetCount uint64) {
	builder.PrependUint64Slot(2, packetCount, 0)
}
func RecvStatsAddByteCount(builder *flatbuffers.Builder, byteCount uint64) {
	builder.PrependUint64Slot(3, byteCount, 0)
}
func RecvStatsAddBitrate(builder *flatbuffers.Builder, bitrate uint32) {
	builder.PrependUint32Slot(4, bitrate, 0)
}
func RecvStatsAddBitrateByLayer(builder *flatbuffers.Builder, bitrateByLayer flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(bitrateByLayer), 0)
}
func RecvStatsStartBitrateByLayerVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func RecvStatsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}