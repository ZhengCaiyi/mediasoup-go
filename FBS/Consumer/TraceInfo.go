// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Consumer

import (
	"strconv"
	flatbuffers "github.com/google/flatbuffers/go"
)

type TraceInfo byte

const (
	TraceInfoNONE              TraceInfo = 0
	TraceInfoKeyFrameTraceInfo TraceInfo = 1
	TraceInfoFirTraceInfo      TraceInfo = 2
	TraceInfoPliTraceInfo      TraceInfo = 3
	TraceInfoRtpTraceInfo      TraceInfo = 4
)

var EnumNamesTraceInfo = map[TraceInfo]string{
	TraceInfoNONE:              "NONE",
	TraceInfoKeyFrameTraceInfo: "KeyFrameTraceInfo",
	TraceInfoFirTraceInfo:      "FirTraceInfo",
	TraceInfoPliTraceInfo:      "PliTraceInfo",
	TraceInfoRtpTraceInfo:      "RtpTraceInfo",
}

var EnumValuesTraceInfo = map[string]TraceInfo{
	"NONE":              TraceInfoNONE,
	"KeyFrameTraceInfo": TraceInfoKeyFrameTraceInfo,
	"FirTraceInfo":      TraceInfoFirTraceInfo,
	"PliTraceInfo":      TraceInfoPliTraceInfo,
	"RtpTraceInfo":      TraceInfoRtpTraceInfo,
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
	case TraceInfoKeyFrameTraceInfo:
		return t.Value.(*KeyFrameTraceInfoT).Pack(builder)
	case TraceInfoFirTraceInfo:
		return t.Value.(*FirTraceInfoT).Pack(builder)
	case TraceInfoPliTraceInfo:
		return t.Value.(*PliTraceInfoT).Pack(builder)
	case TraceInfoRtpTraceInfo:
		return t.Value.(*RtpTraceInfoT).Pack(builder)
	}
	return 0
}

func (rcv TraceInfo) UnPack(table flatbuffers.Table) *TraceInfoT {
	switch rcv {
	case TraceInfoKeyFrameTraceInfo:
		var x KeyFrameTraceInfo
		x.Init(table.Bytes, table.Pos)
		return &TraceInfoT{ Type: TraceInfoKeyFrameTraceInfo, Value: x.UnPack() }
	case TraceInfoFirTraceInfo:
		var x FirTraceInfo
		x.Init(table.Bytes, table.Pos)
		return &TraceInfoT{ Type: TraceInfoFirTraceInfo, Value: x.UnPack() }
	case TraceInfoPliTraceInfo:
		var x PliTraceInfo
		x.Init(table.Bytes, table.Pos)
		return &TraceInfoT{ Type: TraceInfoPliTraceInfo, Value: x.UnPack() }
	case TraceInfoRtpTraceInfo:
		var x RtpTraceInfo
		x.Init(table.Bytes, table.Pos)
		return &TraceInfoT{ Type: TraceInfoRtpTraceInfo, Value: x.UnPack() }
	}
	return nil
}