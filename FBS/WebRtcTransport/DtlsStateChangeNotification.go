// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package WebRtcTransport

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type DtlsStateChangeNotificationT struct {
	DtlsState DtlsState `json:"dtls_state"`
	RemoteCert string `json:"remote_cert"`
}

func (t *DtlsStateChangeNotificationT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	remoteCertOffset := flatbuffers.UOffsetT(0)
	if t.RemoteCert != "" {
		remoteCertOffset = builder.CreateString(t.RemoteCert)
	}
	DtlsStateChangeNotificationStart(builder)
	DtlsStateChangeNotificationAddDtlsState(builder, t.DtlsState)
	DtlsStateChangeNotificationAddRemoteCert(builder, remoteCertOffset)
	return DtlsStateChangeNotificationEnd(builder)
}

func (rcv *DtlsStateChangeNotification) UnPackTo(t *DtlsStateChangeNotificationT) {
	t.DtlsState = rcv.DtlsState()
	t.RemoteCert = string(rcv.RemoteCert())
}

func (rcv *DtlsStateChangeNotification) UnPack() *DtlsStateChangeNotificationT {
	if rcv == nil { return nil }
	t := &DtlsStateChangeNotificationT{}
	rcv.UnPackTo(t)
	return t
}

type DtlsStateChangeNotification struct {
	_tab flatbuffers.Table
}

func GetRootAsDtlsStateChangeNotification(buf []byte, offset flatbuffers.UOffsetT) *DtlsStateChangeNotification {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &DtlsStateChangeNotification{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsDtlsStateChangeNotification(buf []byte, offset flatbuffers.UOffsetT) *DtlsStateChangeNotification {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &DtlsStateChangeNotification{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *DtlsStateChangeNotification) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *DtlsStateChangeNotification) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *DtlsStateChangeNotification) DtlsState() DtlsState {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return DtlsState(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *DtlsStateChangeNotification) MutateDtlsState(n DtlsState) bool {
	return rcv._tab.MutateByteSlot(4, byte(n))
}

func (rcv *DtlsStateChangeNotification) RemoteCert() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func DtlsStateChangeNotificationStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func DtlsStateChangeNotificationAddDtlsState(builder *flatbuffers.Builder, dtlsState DtlsState) {
	builder.PrependByteSlot(0, byte(dtlsState), 0)
}
func DtlsStateChangeNotificationAddRemoteCert(builder *flatbuffers.Builder, remoteCert flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(remoteCert), 0)
}
func DtlsStateChangeNotificationEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
