// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Consumer

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type LayersChangeNotificationT struct {
	Layers *ConsumerLayersT `json:"layers"`
}

func (t *LayersChangeNotificationT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	layersOffset := t.Layers.Pack(builder)
	LayersChangeNotificationStart(builder)
	LayersChangeNotificationAddLayers(builder, layersOffset)
	return LayersChangeNotificationEnd(builder)
}

func (rcv *LayersChangeNotification) UnPackTo(t *LayersChangeNotificationT) {
	t.Layers = rcv.Layers(nil).UnPack()
}

func (rcv *LayersChangeNotification) UnPack() *LayersChangeNotificationT {
	if rcv == nil { return nil }
	t := &LayersChangeNotificationT{}
	rcv.UnPackTo(t)
	return t
}

type LayersChangeNotification struct {
	_tab flatbuffers.Table
}

func GetRootAsLayersChangeNotification(buf []byte, offset flatbuffers.UOffsetT) *LayersChangeNotification {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &LayersChangeNotification{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsLayersChangeNotification(buf []byte, offset flatbuffers.UOffsetT) *LayersChangeNotification {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &LayersChangeNotification{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *LayersChangeNotification) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *LayersChangeNotification) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *LayersChangeNotification) Layers(obj *ConsumerLayers) *ConsumerLayers {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(ConsumerLayers)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func LayersChangeNotificationStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func LayersChangeNotificationAddLayers(builder *flatbuffers.Builder, layers flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(layers), 0)
}
func LayersChangeNotificationEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
