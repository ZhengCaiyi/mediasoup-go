// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Transport

import (
	flatbuffers "github.com/google/flatbuffers/go"

	FBS__SctpAssociation "github.com/jiyeyuran/mediasoup-go/FBS/SctpAssociation"
)

type StatsT struct {
	TransportId string `json:"transport_id"`
	Timestamp uint64 `json:"timestamp"`
	SctpState *FBS__SctpAssociation.SctpState `json:"sctp_state"`
	BytesReceived uint64 `json:"bytes_received"`
	RecvBitrate uint32 `json:"recv_bitrate"`
	BytesSent uint64 `json:"bytes_sent"`
	SendBitrate uint32 `json:"send_bitrate"`
	RtpBytesReceived uint64 `json:"rtp_bytes_received"`
	RtpRecvBitrate uint32 `json:"rtp_recv_bitrate"`
	RtpBytesSent uint64 `json:"rtp_bytes_sent"`
	RtpSendBitrate uint32 `json:"rtp_send_bitrate"`
	RtxBytesReceived uint64 `json:"rtx_bytes_received"`
	RtxRecvBitrate uint32 `json:"rtx_recv_bitrate"`
	RtxBytesSent uint64 `json:"rtx_bytes_sent"`
	RtxSendBitrate uint32 `json:"rtx_send_bitrate"`
	ProbationBytesSent uint64 `json:"probation_bytes_sent"`
	ProbationSendBitrate uint32 `json:"probation_send_bitrate"`
	AvailableOutgoingBitrate *uint32 `json:"available_outgoing_bitrate"`
	AvailableIncomingBitrate *uint32 `json:"available_incoming_bitrate"`
	MaxIncomingBitrate *uint32 `json:"max_incoming_bitrate"`
	MaxOutgoingBitrate *uint32 `json:"max_outgoing_bitrate"`
	MinOutgoingBitrate *uint32 `json:"min_outgoing_bitrate"`
	RtpPacketLossReceived *float64 `json:"rtp_packet_loss_received"`
	RtpPacketLossSent *float64 `json:"rtp_packet_loss_sent"`
}

func (t *StatsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	transportIdOffset := flatbuffers.UOffsetT(0)
	if t.TransportId != "" {
		transportIdOffset = builder.CreateString(t.TransportId)
	}
	StatsStart(builder)
	StatsAddTransportId(builder, transportIdOffset)
	StatsAddTimestamp(builder, t.Timestamp)
	if t.SctpState != nil {
		StatsAddSctpState(builder, *t.SctpState)
	}
	StatsAddBytesReceived(builder, t.BytesReceived)
	StatsAddRecvBitrate(builder, t.RecvBitrate)
	StatsAddBytesSent(builder, t.BytesSent)
	StatsAddSendBitrate(builder, t.SendBitrate)
	StatsAddRtpBytesReceived(builder, t.RtpBytesReceived)
	StatsAddRtpRecvBitrate(builder, t.RtpRecvBitrate)
	StatsAddRtpBytesSent(builder, t.RtpBytesSent)
	StatsAddRtpSendBitrate(builder, t.RtpSendBitrate)
	StatsAddRtxBytesReceived(builder, t.RtxBytesReceived)
	StatsAddRtxRecvBitrate(builder, t.RtxRecvBitrate)
	StatsAddRtxBytesSent(builder, t.RtxBytesSent)
	StatsAddRtxSendBitrate(builder, t.RtxSendBitrate)
	StatsAddProbationBytesSent(builder, t.ProbationBytesSent)
	StatsAddProbationSendBitrate(builder, t.ProbationSendBitrate)
	if t.AvailableOutgoingBitrate != nil {
		StatsAddAvailableOutgoingBitrate(builder, *t.AvailableOutgoingBitrate)
	}
	if t.AvailableIncomingBitrate != nil {
		StatsAddAvailableIncomingBitrate(builder, *t.AvailableIncomingBitrate)
	}
	if t.MaxIncomingBitrate != nil {
		StatsAddMaxIncomingBitrate(builder, *t.MaxIncomingBitrate)
	}
	if t.MaxOutgoingBitrate != nil {
		StatsAddMaxOutgoingBitrate(builder, *t.MaxOutgoingBitrate)
	}
	if t.MinOutgoingBitrate != nil {
		StatsAddMinOutgoingBitrate(builder, *t.MinOutgoingBitrate)
	}
	if t.RtpPacketLossReceived != nil {
		StatsAddRtpPacketLossReceived(builder, *t.RtpPacketLossReceived)
	}
	if t.RtpPacketLossSent != nil {
		StatsAddRtpPacketLossSent(builder, *t.RtpPacketLossSent)
	}
	return StatsEnd(builder)
}

func (rcv *Stats) UnPackTo(t *StatsT) {
	t.TransportId = string(rcv.TransportId())
	t.Timestamp = rcv.Timestamp()
	t.SctpState = rcv.SctpState()
	t.BytesReceived = rcv.BytesReceived()
	t.RecvBitrate = rcv.RecvBitrate()
	t.BytesSent = rcv.BytesSent()
	t.SendBitrate = rcv.SendBitrate()
	t.RtpBytesReceived = rcv.RtpBytesReceived()
	t.RtpRecvBitrate = rcv.RtpRecvBitrate()
	t.RtpBytesSent = rcv.RtpBytesSent()
	t.RtpSendBitrate = rcv.RtpSendBitrate()
	t.RtxBytesReceived = rcv.RtxBytesReceived()
	t.RtxRecvBitrate = rcv.RtxRecvBitrate()
	t.RtxBytesSent = rcv.RtxBytesSent()
	t.RtxSendBitrate = rcv.RtxSendBitrate()
	t.ProbationBytesSent = rcv.ProbationBytesSent()
	t.ProbationSendBitrate = rcv.ProbationSendBitrate()
	t.AvailableOutgoingBitrate = rcv.AvailableOutgoingBitrate()
	t.AvailableIncomingBitrate = rcv.AvailableIncomingBitrate()
	t.MaxIncomingBitrate = rcv.MaxIncomingBitrate()
	t.MaxOutgoingBitrate = rcv.MaxOutgoingBitrate()
	t.MinOutgoingBitrate = rcv.MinOutgoingBitrate()
	t.RtpPacketLossReceived = rcv.RtpPacketLossReceived()
	t.RtpPacketLossSent = rcv.RtpPacketLossSent()
}

func (rcv *Stats) UnPack() *StatsT {
	if rcv == nil { return nil }
	t := &StatsT{}
	rcv.UnPackTo(t)
	return t
}

type Stats struct {
	_tab flatbuffers.Table
}

func GetRootAsStats(buf []byte, offset flatbuffers.UOffsetT) *Stats {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Stats{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsStats(buf []byte, offset flatbuffers.UOffsetT) *Stats {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Stats{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Stats) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Stats) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Stats) TransportId() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Stats) Timestamp() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Stats) MutateTimestamp(n uint64) bool {
	return rcv._tab.MutateUint64Slot(6, n)
}

func (rcv *Stats) SctpState() *FBS__SctpAssociation.SctpState {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		v := FBS__SctpAssociation.SctpState(rcv._tab.GetByte(o + rcv._tab.Pos))
		return &v
	}
	return nil
}

func (rcv *Stats) MutateSctpState(n FBS__SctpAssociation.SctpState) bool {
	return rcv._tab.MutateByteSlot(8, byte(n))
}

func (rcv *Stats) BytesReceived() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Stats) MutateBytesReceived(n uint64) bool {
	return rcv._tab.MutateUint64Slot(10, n)
}

func (rcv *Stats) RecvBitrate() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Stats) MutateRecvBitrate(n uint32) bool {
	return rcv._tab.MutateUint32Slot(12, n)
}

func (rcv *Stats) BytesSent() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Stats) MutateBytesSent(n uint64) bool {
	return rcv._tab.MutateUint64Slot(14, n)
}

func (rcv *Stats) SendBitrate() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Stats) MutateSendBitrate(n uint32) bool {
	return rcv._tab.MutateUint32Slot(16, n)
}

func (rcv *Stats) RtpBytesReceived() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Stats) MutateRtpBytesReceived(n uint64) bool {
	return rcv._tab.MutateUint64Slot(18, n)
}

func (rcv *Stats) RtpRecvBitrate() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Stats) MutateRtpRecvBitrate(n uint32) bool {
	return rcv._tab.MutateUint32Slot(20, n)
}

func (rcv *Stats) RtpBytesSent() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Stats) MutateRtpBytesSent(n uint64) bool {
	return rcv._tab.MutateUint64Slot(22, n)
}

func (rcv *Stats) RtpSendBitrate() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Stats) MutateRtpSendBitrate(n uint32) bool {
	return rcv._tab.MutateUint32Slot(24, n)
}

func (rcv *Stats) RtxBytesReceived() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Stats) MutateRtxBytesReceived(n uint64) bool {
	return rcv._tab.MutateUint64Slot(26, n)
}

func (rcv *Stats) RtxRecvBitrate() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(28))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Stats) MutateRtxRecvBitrate(n uint32) bool {
	return rcv._tab.MutateUint32Slot(28, n)
}

func (rcv *Stats) RtxBytesSent() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(30))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Stats) MutateRtxBytesSent(n uint64) bool {
	return rcv._tab.MutateUint64Slot(30, n)
}

func (rcv *Stats) RtxSendBitrate() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(32))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Stats) MutateRtxSendBitrate(n uint32) bool {
	return rcv._tab.MutateUint32Slot(32, n)
}

func (rcv *Stats) ProbationBytesSent() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(34))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Stats) MutateProbationBytesSent(n uint64) bool {
	return rcv._tab.MutateUint64Slot(34, n)
}

func (rcv *Stats) ProbationSendBitrate() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(36))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Stats) MutateProbationSendBitrate(n uint32) bool {
	return rcv._tab.MutateUint32Slot(36, n)
}

func (rcv *Stats) AvailableOutgoingBitrate() *uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(38))
	if o != 0 {
		v := rcv._tab.GetUint32(o + rcv._tab.Pos)
		return &v
	}
	return nil
}

func (rcv *Stats) MutateAvailableOutgoingBitrate(n uint32) bool {
	return rcv._tab.MutateUint32Slot(38, n)
}

func (rcv *Stats) AvailableIncomingBitrate() *uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(40))
	if o != 0 {
		v := rcv._tab.GetUint32(o + rcv._tab.Pos)
		return &v
	}
	return nil
}

func (rcv *Stats) MutateAvailableIncomingBitrate(n uint32) bool {
	return rcv._tab.MutateUint32Slot(40, n)
}

func (rcv *Stats) MaxIncomingBitrate() *uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(42))
	if o != 0 {
		v := rcv._tab.GetUint32(o + rcv._tab.Pos)
		return &v
	}
	return nil
}

func (rcv *Stats) MutateMaxIncomingBitrate(n uint32) bool {
	return rcv._tab.MutateUint32Slot(42, n)
}

func (rcv *Stats) MaxOutgoingBitrate() *uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(44))
	if o != 0 {
		v := rcv._tab.GetUint32(o + rcv._tab.Pos)
		return &v
	}
	return nil
}

func (rcv *Stats) MutateMaxOutgoingBitrate(n uint32) bool {
	return rcv._tab.MutateUint32Slot(44, n)
}

func (rcv *Stats) MinOutgoingBitrate() *uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(46))
	if o != 0 {
		v := rcv._tab.GetUint32(o + rcv._tab.Pos)
		return &v
	}
	return nil
}

func (rcv *Stats) MutateMinOutgoingBitrate(n uint32) bool {
	return rcv._tab.MutateUint32Slot(46, n)
}

func (rcv *Stats) RtpPacketLossReceived() *float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(48))
	if o != 0 {
		v := rcv._tab.GetFloat64(o + rcv._tab.Pos)
		return &v
	}
	return nil
}

func (rcv *Stats) MutateRtpPacketLossReceived(n float64) bool {
	return rcv._tab.MutateFloat64Slot(48, n)
}

func (rcv *Stats) RtpPacketLossSent() *float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(50))
	if o != 0 {
		v := rcv._tab.GetFloat64(o + rcv._tab.Pos)
		return &v
	}
	return nil
}

func (rcv *Stats) MutateRtpPacketLossSent(n float64) bool {
	return rcv._tab.MutateFloat64Slot(50, n)
}

func StatsStart(builder *flatbuffers.Builder) {
	builder.StartObject(24)
}
func StatsAddTransportId(builder *flatbuffers.Builder, transportId flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(transportId), 0)
}
func StatsAddTimestamp(builder *flatbuffers.Builder, timestamp uint64) {
	builder.PrependUint64Slot(1, timestamp, 0)
}
func StatsAddSctpState(builder *flatbuffers.Builder, sctpState FBS__SctpAssociation.SctpState) {
	builder.PrependByte(byte(sctpState))
	builder.Slot(2)
}
func StatsAddBytesReceived(builder *flatbuffers.Builder, bytesReceived uint64) {
	builder.PrependUint64Slot(3, bytesReceived, 0)
}
func StatsAddRecvBitrate(builder *flatbuffers.Builder, recvBitrate uint32) {
	builder.PrependUint32Slot(4, recvBitrate, 0)
}
func StatsAddBytesSent(builder *flatbuffers.Builder, bytesSent uint64) {
	builder.PrependUint64Slot(5, bytesSent, 0)
}
func StatsAddSendBitrate(builder *flatbuffers.Builder, sendBitrate uint32) {
	builder.PrependUint32Slot(6, sendBitrate, 0)
}
func StatsAddRtpBytesReceived(builder *flatbuffers.Builder, rtpBytesReceived uint64) {
	builder.PrependUint64Slot(7, rtpBytesReceived, 0)
}
func StatsAddRtpRecvBitrate(builder *flatbuffers.Builder, rtpRecvBitrate uint32) {
	builder.PrependUint32Slot(8, rtpRecvBitrate, 0)
}
func StatsAddRtpBytesSent(builder *flatbuffers.Builder, rtpBytesSent uint64) {
	builder.PrependUint64Slot(9, rtpBytesSent, 0)
}
func StatsAddRtpSendBitrate(builder *flatbuffers.Builder, rtpSendBitrate uint32) {
	builder.PrependUint32Slot(10, rtpSendBitrate, 0)
}
func StatsAddRtxBytesReceived(builder *flatbuffers.Builder, rtxBytesReceived uint64) {
	builder.PrependUint64Slot(11, rtxBytesReceived, 0)
}
func StatsAddRtxRecvBitrate(builder *flatbuffers.Builder, rtxRecvBitrate uint32) {
	builder.PrependUint32Slot(12, rtxRecvBitrate, 0)
}
func StatsAddRtxBytesSent(builder *flatbuffers.Builder, rtxBytesSent uint64) {
	builder.PrependUint64Slot(13, rtxBytesSent, 0)
}
func StatsAddRtxSendBitrate(builder *flatbuffers.Builder, rtxSendBitrate uint32) {
	builder.PrependUint32Slot(14, rtxSendBitrate, 0)
}
func StatsAddProbationBytesSent(builder *flatbuffers.Builder, probationBytesSent uint64) {
	builder.PrependUint64Slot(15, probationBytesSent, 0)
}
func StatsAddProbationSendBitrate(builder *flatbuffers.Builder, probationSendBitrate uint32) {
	builder.PrependUint32Slot(16, probationSendBitrate, 0)
}
func StatsAddAvailableOutgoingBitrate(builder *flatbuffers.Builder, availableOutgoingBitrate uint32) {
	builder.PrependUint32(availableOutgoingBitrate)
	builder.Slot(17)
}
func StatsAddAvailableIncomingBitrate(builder *flatbuffers.Builder, availableIncomingBitrate uint32) {
	builder.PrependUint32(availableIncomingBitrate)
	builder.Slot(18)
}
func StatsAddMaxIncomingBitrate(builder *flatbuffers.Builder, maxIncomingBitrate uint32) {
	builder.PrependUint32(maxIncomingBitrate)
	builder.Slot(19)
}
func StatsAddMaxOutgoingBitrate(builder *flatbuffers.Builder, maxOutgoingBitrate uint32) {
	builder.PrependUint32(maxOutgoingBitrate)
	builder.Slot(20)
}
func StatsAddMinOutgoingBitrate(builder *flatbuffers.Builder, minOutgoingBitrate uint32) {
	builder.PrependUint32(minOutgoingBitrate)
	builder.Slot(21)
}
func StatsAddRtpPacketLossReceived(builder *flatbuffers.Builder, rtpPacketLossReceived float64) {
	builder.PrependFloat64(rtpPacketLossReceived)
	builder.Slot(22)
}
func StatsAddRtpPacketLossSent(builder *flatbuffers.Builder, rtpPacketLossSent float64) {
	builder.PrependFloat64(rtpPacketLossSent)
	builder.Slot(23)
}
func StatsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
