// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Producer

import "strconv"

type TraceEventType byte

const (
	TraceEventTypeKEYFRAME TraceEventType = 0
	TraceEventTypeFIR      TraceEventType = 1
	TraceEventTypeNACK     TraceEventType = 2
	TraceEventTypePLI      TraceEventType = 3
	TraceEventTypeRTP      TraceEventType = 4
)

var EnumNamesTraceEventType = map[TraceEventType]string{
	TraceEventTypeKEYFRAME: "KEYFRAME",
	TraceEventTypeFIR:      "FIR",
	TraceEventTypeNACK:     "NACK",
	TraceEventTypePLI:      "PLI",
	TraceEventTypeRTP:      "RTP",
}

var EnumValuesTraceEventType = map[string]TraceEventType{
	"KEYFRAME": TraceEventTypeKEYFRAME,
	"FIR":      TraceEventTypeFIR,
	"NACK":     TraceEventTypeNACK,
	"PLI":      TraceEventTypePLI,
	"RTP":      TraceEventTypeRTP,
}

func (v TraceEventType) String() string {
	if s, ok := EnumNamesTraceEventType[v]; ok {
		return s
	}
	return "TraceEventType(" + strconv.FormatInt(int64(v), 10) + ")"
}
