// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Common

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type StringStringArrayT struct {
	Key string `json:"key"`
	Values []string `json:"values"`
}

func (t *StringStringArrayT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	keyOffset := flatbuffers.UOffsetT(0)
	if t.Key != "" {
		keyOffset = builder.CreateString(t.Key)
	}
	valuesOffset := flatbuffers.UOffsetT(0)
	if t.Values != nil {
		valuesLength := len(t.Values)
		valuesOffsets := make([]flatbuffers.UOffsetT, valuesLength)
		for j := 0; j < valuesLength; j++ {
			valuesOffsets[j] = builder.CreateString(t.Values[j])
		}
		StringStringArrayStartValuesVector(builder, valuesLength)
		for j := valuesLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(valuesOffsets[j])
		}
		valuesOffset = builder.EndVector(valuesLength)
	}
	StringStringArrayStart(builder)
	StringStringArrayAddKey(builder, keyOffset)
	StringStringArrayAddValues(builder, valuesOffset)
	return StringStringArrayEnd(builder)
}

func (rcv *StringStringArray) UnPackTo(t *StringStringArrayT) {
	t.Key = string(rcv.Key())
	valuesLength := rcv.ValuesLength()
	t.Values = make([]string, valuesLength)
	for j := 0; j < valuesLength; j++ {
		t.Values[j] = string(rcv.Values(j))
	}
}

func (rcv *StringStringArray) UnPack() *StringStringArrayT {
	if rcv == nil { return nil }
	t := &StringStringArrayT{}
	rcv.UnPackTo(t)
	return t
}

type StringStringArray struct {
	_tab flatbuffers.Table
}

func GetRootAsStringStringArray(buf []byte, offset flatbuffers.UOffsetT) *StringStringArray {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &StringStringArray{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsStringStringArray(buf []byte, offset flatbuffers.UOffsetT) *StringStringArray {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &StringStringArray{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *StringStringArray) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *StringStringArray) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *StringStringArray) Key() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *StringStringArray) Values(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *StringStringArray) ValuesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func StringStringArrayStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func StringStringArrayAddKey(builder *flatbuffers.Builder, key flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(key), 0)
}
func StringStringArrayAddValues(builder *flatbuffers.Builder, values flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(values), 0)
}
func StringStringArrayStartValuesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func StringStringArrayEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
