// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Producer

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type VideoOrientationChangeNotificationT struct {
	Camera bool `json:"camera"`
	Flip bool `json:"flip"`
	Rotation uint16 `json:"rotation"`
}

func (t *VideoOrientationChangeNotificationT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	VideoOrientationChangeNotificationStart(builder)
	VideoOrientationChangeNotificationAddCamera(builder, t.Camera)
	VideoOrientationChangeNotificationAddFlip(builder, t.Flip)
	VideoOrientationChangeNotificationAddRotation(builder, t.Rotation)
	return VideoOrientationChangeNotificationEnd(builder)
}

func (rcv *VideoOrientationChangeNotification) UnPackTo(t *VideoOrientationChangeNotificationT) {
	t.Camera = rcv.Camera()
	t.Flip = rcv.Flip()
	t.Rotation = rcv.Rotation()
}

func (rcv *VideoOrientationChangeNotification) UnPack() *VideoOrientationChangeNotificationT {
	if rcv == nil { return nil }
	t := &VideoOrientationChangeNotificationT{}
	rcv.UnPackTo(t)
	return t
}

type VideoOrientationChangeNotification struct {
	_tab flatbuffers.Table
}

func GetRootAsVideoOrientationChangeNotification(buf []byte, offset flatbuffers.UOffsetT) *VideoOrientationChangeNotification {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &VideoOrientationChangeNotification{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsVideoOrientationChangeNotification(buf []byte, offset flatbuffers.UOffsetT) *VideoOrientationChangeNotification {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &VideoOrientationChangeNotification{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *VideoOrientationChangeNotification) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *VideoOrientationChangeNotification) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *VideoOrientationChangeNotification) Camera() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *VideoOrientationChangeNotification) MutateCamera(n bool) bool {
	return rcv._tab.MutateBoolSlot(4, n)
}

func (rcv *VideoOrientationChangeNotification) Flip() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *VideoOrientationChangeNotification) MutateFlip(n bool) bool {
	return rcv._tab.MutateBoolSlot(6, n)
}

func (rcv *VideoOrientationChangeNotification) Rotation() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *VideoOrientationChangeNotification) MutateRotation(n uint16) bool {
	return rcv._tab.MutateUint16Slot(8, n)
}

func VideoOrientationChangeNotificationStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func VideoOrientationChangeNotificationAddCamera(builder *flatbuffers.Builder, camera bool) {
	builder.PrependBoolSlot(0, camera, false)
}
func VideoOrientationChangeNotificationAddFlip(builder *flatbuffers.Builder, flip bool) {
	builder.PrependBoolSlot(1, flip, false)
}
func VideoOrientationChangeNotificationAddRotation(builder *flatbuffers.Builder, rotation uint16) {
	builder.PrependUint16Slot(2, rotation, 0)
}
func VideoOrientationChangeNotificationEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
