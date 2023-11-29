// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package DataConsumer

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type GetStatsResponseT struct {
	Timestamp uint64 `json:"timestamp"`
	Label string `json:"label"`
	Protocol string `json:"protocol"`
	MessagesSent uint64 `json:"messages_sent"`
	BytesSent uint64 `json:"bytes_sent"`
	BufferedAmount uint32 `json:"buffered_amount"`
}

func (t *GetStatsResponseT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	labelOffset := flatbuffers.UOffsetT(0)
	if t.Label != "" {
		labelOffset = builder.CreateString(t.Label)
	}
	protocolOffset := flatbuffers.UOffsetT(0)
	if t.Protocol != "" {
		protocolOffset = builder.CreateString(t.Protocol)
	}
	GetStatsResponseStart(builder)
	GetStatsResponseAddTimestamp(builder, t.Timestamp)
	GetStatsResponseAddLabel(builder, labelOffset)
	GetStatsResponseAddProtocol(builder, protocolOffset)
	GetStatsResponseAddMessagesSent(builder, t.MessagesSent)
	GetStatsResponseAddBytesSent(builder, t.BytesSent)
	GetStatsResponseAddBufferedAmount(builder, t.BufferedAmount)
	return GetStatsResponseEnd(builder)
}

func (rcv *GetStatsResponse) UnPackTo(t *GetStatsResponseT) {
	t.Timestamp = rcv.Timestamp()
	t.Label = string(rcv.Label())
	t.Protocol = string(rcv.Protocol())
	t.MessagesSent = rcv.MessagesSent()
	t.BytesSent = rcv.BytesSent()
	t.BufferedAmount = rcv.BufferedAmount()
}

func (rcv *GetStatsResponse) UnPack() *GetStatsResponseT {
	if rcv == nil { return nil }
	t := &GetStatsResponseT{}
	rcv.UnPackTo(t)
	return t
}

type GetStatsResponse struct {
	_tab flatbuffers.Table
}

func GetRootAsGetStatsResponse(buf []byte, offset flatbuffers.UOffsetT) *GetStatsResponse {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &GetStatsResponse{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsGetStatsResponse(buf []byte, offset flatbuffers.UOffsetT) *GetStatsResponse {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &GetStatsResponse{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *GetStatsResponse) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *GetStatsResponse) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *GetStatsResponse) Timestamp() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *GetStatsResponse) MutateTimestamp(n uint64) bool {
	return rcv._tab.MutateUint64Slot(4, n)
}

func (rcv *GetStatsResponse) Label() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *GetStatsResponse) Protocol() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *GetStatsResponse) MessagesSent() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *GetStatsResponse) MutateMessagesSent(n uint64) bool {
	return rcv._tab.MutateUint64Slot(10, n)
}

func (rcv *GetStatsResponse) BytesSent() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *GetStatsResponse) MutateBytesSent(n uint64) bool {
	return rcv._tab.MutateUint64Slot(12, n)
}

func (rcv *GetStatsResponse) BufferedAmount() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *GetStatsResponse) MutateBufferedAmount(n uint32) bool {
	return rcv._tab.MutateUint32Slot(14, n)
}

func GetStatsResponseStart(builder *flatbuffers.Builder) {
	builder.StartObject(6)
}
func GetStatsResponseAddTimestamp(builder *flatbuffers.Builder, timestamp uint64) {
	builder.PrependUint64Slot(0, timestamp, 0)
}
func GetStatsResponseAddLabel(builder *flatbuffers.Builder, label flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(label), 0)
}
func GetStatsResponseAddProtocol(builder *flatbuffers.Builder, protocol flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(protocol), 0)
}
func GetStatsResponseAddMessagesSent(builder *flatbuffers.Builder, messagesSent uint64) {
	builder.PrependUint64Slot(3, messagesSent, 0)
}
func GetStatsResponseAddBytesSent(builder *flatbuffers.Builder, bytesSent uint64) {
	builder.PrependUint64Slot(4, bytesSent, 0)
}
func GetStatsResponseAddBufferedAmount(builder *flatbuffers.Builder, bufferedAmount uint32) {
	builder.PrependUint32Slot(5, bufferedAmount, 0)
}
func GetStatsResponseEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
