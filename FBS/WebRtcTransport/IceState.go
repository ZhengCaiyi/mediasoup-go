// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package WebRtcTransport

import "strconv"

type IceState byte

const (
	IceStateNEW          IceState = 0
	IceStateCONNECTED    IceState = 1
	IceStateCOMPLETED    IceState = 2
	IceStateDISCONNECTED IceState = 3
)

var EnumNamesIceState = map[IceState]string{
	IceStateNEW:          "NEW",
	IceStateCONNECTED:    "CONNECTED",
	IceStateCOMPLETED:    "COMPLETED",
	IceStateDISCONNECTED: "DISCONNECTED",
}

var EnumValuesIceState = map[string]IceState{
	"NEW":          IceStateNEW,
	"CONNECTED":    IceStateCONNECTED,
	"COMPLETED":    IceStateCOMPLETED,
	"DISCONNECTED": IceStateDISCONNECTED,
}

func (v IceState) String() string {
	if s, ok := EnumNamesIceState[v]; ok {
		return s
	}
	return "IceState(" + strconv.FormatInt(int64(v), 10) + ")"
}
