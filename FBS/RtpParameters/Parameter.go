// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package RtpParameters

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ParameterT struct {
	Name string `json:"name"`
	Value *ValueT `json:"value"`
}

func (t *ParameterT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	nameOffset := flatbuffers.UOffsetT(0)
	if t.Name != "" {
		nameOffset = builder.CreateString(t.Name)
	}
	valueOffset := t.Value.Pack(builder)
	
	ParameterStart(builder)
	ParameterAddName(builder, nameOffset)
	if t.Value != nil {
		ParameterAddValueType(builder, t.Value.Type)
	}
	ParameterAddValue(builder, valueOffset)
	return ParameterEnd(builder)
}

func (rcv *Parameter) UnPackTo(t *ParameterT) {
	t.Name = string(rcv.Name())
	valueTable := flatbuffers.Table{}
	if rcv.Value(&valueTable) {
		t.Value = rcv.ValueType().UnPack(valueTable)
	}
}

func (rcv *Parameter) UnPack() *ParameterT {
	if rcv == nil { return nil }
	t := &ParameterT{}
	rcv.UnPackTo(t)
	return t
}

type Parameter struct {
	_tab flatbuffers.Table
}

func GetRootAsParameter(buf []byte, offset flatbuffers.UOffsetT) *Parameter {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Parameter{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsParameter(buf []byte, offset flatbuffers.UOffsetT) *Parameter {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Parameter{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Parameter) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Parameter) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Parameter) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Parameter) ValueType() Value {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return Value(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *Parameter) MutateValueType(n Value) bool {
	return rcv._tab.MutateByteSlot(6, byte(n))
}

func (rcv *Parameter) Value(obj *flatbuffers.Table) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		rcv._tab.Union(obj, o)
		return true
	}
	return false
}

func ParameterStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func ParameterAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(name), 0)
}
func ParameterAddValueType(builder *flatbuffers.Builder, valueType Value) {
	builder.PrependByteSlot(1, byte(valueType), 0)
}
func ParameterAddValue(builder *flatbuffers.Builder, value flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(value), 0)
}
func ParameterEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
