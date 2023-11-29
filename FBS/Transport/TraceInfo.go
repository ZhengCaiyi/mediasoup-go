// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Transport

import (
	"strconv"
	flatbuffers "github.com/google/flatbuffers/go"
)

type TraceInfo byte

const (
	TraceInfoNONE         TraceInfo = 0
	TraceInfoBweTraceInfo TraceInfo = 1
)

var EnumNamesTraceInfo = map[TraceInfo]string{
	TraceInfoNONE:         "NONE",
	TraceInfoBweTraceInfo: "BweTraceInfo",
}

var EnumValuesTraceInfo = map[string]TraceInfo{
	"NONE":         TraceInfoNONE,
	"BweTraceInfo": TraceInfoBweTraceInfo,
}

func (v TraceInfo) String() string {
	if s, ok := EnumNamesTraceInfo[v]; ok {
		return s
	}
	return "TraceInfo(" + strconv.FormatInt(int64(v), 10) + ")"
}

type TraceInfoT struct {
	Type TraceInfo
	Value interface{}
}

func (t *TraceInfoT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	switch t.Type {
	case TraceInfoBweTraceInfo:
		return t.Value.(*BweTraceInfoT).Pack(builder)
	}
	return 0
}

func (rcv TraceInfo) UnPack(table flatbuffers.Table) *TraceInfoT {
	switch rcv {
	case TraceInfoBweTraceInfo:
		var x BweTraceInfo
		x.Init(table.Bytes, table.Pos)
		return &TraceInfoT{ Type: TraceInfoBweTraceInfo, Value: x.UnPack() }
	}
	return nil
}
