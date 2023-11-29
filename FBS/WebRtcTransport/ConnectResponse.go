// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package WebRtcTransport

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ConnectResponseT struct {
	DtlsLocalRole DtlsRole `json:"dtls_local_role"`
}

func (t *ConnectResponseT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	ConnectResponseStart(builder)
	ConnectResponseAddDtlsLocalRole(builder, t.DtlsLocalRole)
	return ConnectResponseEnd(builder)
}

func (rcv *ConnectResponse) UnPackTo(t *ConnectResponseT) {
	t.DtlsLocalRole = rcv.DtlsLocalRole()
}

func (rcv *ConnectResponse) UnPack() *ConnectResponseT {
	if rcv == nil { return nil }
	t := &ConnectResponseT{}
	rcv.UnPackTo(t)
	return t
}

type ConnectResponse struct {
	_tab flatbuffers.Table
}

func GetRootAsConnectResponse(buf []byte, offset flatbuffers.UOffsetT) *ConnectResponse {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ConnectResponse{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsConnectResponse(buf []byte, offset flatbuffers.UOffsetT) *ConnectResponse {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ConnectResponse{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ConnectResponse) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ConnectResponse) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ConnectResponse) DtlsLocalRole() DtlsRole {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return DtlsRole(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *ConnectResponse) MutateDtlsLocalRole(n DtlsRole) bool {
	return rcv._tab.MutateByteSlot(4, byte(n))
}

func ConnectResponseStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func ConnectResponseAddDtlsLocalRole(builder *flatbuffers.Builder, dtlsLocalRole DtlsRole) {
	builder.PrependByteSlot(0, byte(dtlsLocalRole), 0)
}
func ConnectResponseEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
