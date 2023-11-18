// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: pb/service.proto

package pb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DragFunction int32

const (
	DragFunction_DRAG_FUNCTION_UNSPECIFIED DragFunction = 0
	DragFunction_DRAG_FUNCTION_G1          DragFunction = 1
	DragFunction_DRAG_FUNCTION_G7          DragFunction = 7
)

// Enum value maps for DragFunction.
var (
	DragFunction_name = map[int32]string{
		0: "DRAG_FUNCTION_UNSPECIFIED",
		1: "DRAG_FUNCTION_G1",
		7: "DRAG_FUNCTION_G7",
	}
	DragFunction_value = map[string]int32{
		"DRAG_FUNCTION_UNSPECIFIED": 0,
		"DRAG_FUNCTION_G1":          1,
		"DRAG_FUNCTION_G7":          7,
	}
)

func (x DragFunction) Enum() *DragFunction {
	p := new(DragFunction)
	*p = x
	return p
}

func (x DragFunction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DragFunction) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_service_proto_enumTypes[0].Descriptor()
}

func (DragFunction) Type() protoreflect.EnumType {
	return &file_pb_service_proto_enumTypes[0]
}

func (x DragFunction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DragFunction.Descriptor instead.
func (DragFunction) EnumDescriptor() ([]byte, []int) {
	return file_pb_service_proto_rawDescGZIP(), []int{0}
}

type CreateLoadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Load *Load `protobuf:"bytes,1,opt,name=load,proto3" json:"load,omitempty"`
}

func (x *CreateLoadRequest) Reset() {
	*x = CreateLoadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLoadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLoadRequest) ProtoMessage() {}

func (x *CreateLoadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLoadRequest.ProtoReflect.Descriptor instead.
func (*CreateLoadRequest) Descriptor() ([]byte, []int) {
	return file_pb_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateLoadRequest) GetLoad() *Load {
	if x != nil {
		return x.Load
	}
	return nil
}

type CreateLoadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Load *Load `protobuf:"bytes,1,opt,name=load,proto3" json:"load,omitempty"`
}

func (x *CreateLoadResponse) Reset() {
	*x = CreateLoadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLoadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLoadResponse) ProtoMessage() {}

func (x *CreateLoadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLoadResponse.ProtoReflect.Descriptor instead.
func (*CreateLoadResponse) Descriptor() ([]byte, []int) {
	return file_pb_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateLoadResponse) GetLoad() *Load {
	if x != nil {
		return x.Load
	}
	return nil
}

type Load struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LoadId         int64   `protobuf:"varint,1,opt,name=load_id,json=loadId,proto3" json:"load_id,omitempty"`
	Bullet         *Bullet `protobuf:"bytes,2,opt,name=bullet,proto3" json:"bullet,omitempty"`
	MuzzleVelocity float32 `protobuf:"fixed32,3,opt,name=muzzle_velocity,json=muzzleVelocity,proto3" json:"muzzle_velocity,omitempty"`
	Powder         string  `protobuf:"bytes,4,opt,name=powder,proto3" json:"powder,omitempty"`                                   // optional
	PowderCharge   float32 `protobuf:"fixed32,5,opt,name=powder_charge,json=powderCharge,proto3" json:"powder_charge,omitempty"` // optional
}

func (x *Load) Reset() {
	*x = Load{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Load) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Load) ProtoMessage() {}

func (x *Load) ProtoReflect() protoreflect.Message {
	mi := &file_pb_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Load.ProtoReflect.Descriptor instead.
func (*Load) Descriptor() ([]byte, []int) {
	return file_pb_service_proto_rawDescGZIP(), []int{2}
}

func (x *Load) GetLoadId() int64 {
	if x != nil {
		return x.LoadId
	}
	return 0
}

func (x *Load) GetBullet() *Bullet {
	if x != nil {
		return x.Bullet
	}
	return nil
}

func (x *Load) GetMuzzleVelocity() float32 {
	if x != nil {
		return x.MuzzleVelocity
	}
	return 0
}

func (x *Load) GetPowder() string {
	if x != nil {
		return x.Powder
	}
	return ""
}

func (x *Load) GetPowderCharge() float32 {
	if x != nil {
		return x.PowderCharge
	}
	return 0
}

type Bullet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Caliber float32               `protobuf:"fixed32,1,opt,name=caliber,proto3" json:"caliber,omitempty"`
	Weight  float32               `protobuf:"fixed32,2,opt,name=weight,proto3" json:"weight,omitempty"`
	Bc      *BallisticCoefficient `protobuf:"bytes,3,opt,name=bc,proto3" json:"bc,omitempty"`
	Length  float32               `protobuf:"fixed32,4,opt,name=length,proto3" json:"length,omitempty"`
}

func (x *Bullet) Reset() {
	*x = Bullet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bullet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bullet) ProtoMessage() {}

func (x *Bullet) ProtoReflect() protoreflect.Message {
	mi := &file_pb_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bullet.ProtoReflect.Descriptor instead.
func (*Bullet) Descriptor() ([]byte, []int) {
	return file_pb_service_proto_rawDescGZIP(), []int{3}
}

func (x *Bullet) GetCaliber() float32 {
	if x != nil {
		return x.Caliber
	}
	return 0
}

func (x *Bullet) GetWeight() float32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *Bullet) GetBc() *BallisticCoefficient {
	if x != nil {
		return x.Bc
	}
	return nil
}

func (x *Bullet) GetLength() float32 {
	if x != nil {
		return x.Length
	}
	return 0
}

type BallisticCoefficient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value    float32      `protobuf:"fixed32,1,opt,name=value,proto3" json:"value,omitempty"`
	Function DragFunction `protobuf:"varint,2,opt,name=function,proto3,enum=DragFunction" json:"function,omitempty"`
}

func (x *BallisticCoefficient) Reset() {
	*x = BallisticCoefficient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BallisticCoefficient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BallisticCoefficient) ProtoMessage() {}

func (x *BallisticCoefficient) ProtoReflect() protoreflect.Message {
	mi := &file_pb_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BallisticCoefficient.ProtoReflect.Descriptor instead.
func (*BallisticCoefficient) Descriptor() ([]byte, []int) {
	return file_pb_service_proto_rawDescGZIP(), []int{4}
}

func (x *BallisticCoefficient) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *BallisticCoefficient) GetFunction() DragFunction {
	if x != nil {
		return x.Function
	}
	return DragFunction_DRAG_FUNCTION_UNSPECIFIED
}

var File_pb_service_proto protoreflect.FileDescriptor

var file_pb_service_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x62, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x2e, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x61, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x04, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x52, 0x04, 0x6c, 0x6f, 0x61, 0x64,
	0x22, 0x2f, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x61, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x04, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x52, 0x04, 0x6c, 0x6f, 0x61,
	0x64, 0x22, 0xa6, 0x01, 0x0a, 0x04, 0x4c, 0x6f, 0x61, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x6f,
	0x61, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6c, 0x6f, 0x61,
	0x64, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x06, 0x62, 0x75, 0x6c, 0x6c, 0x65, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x42, 0x75, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x06, 0x62, 0x75,
	0x6c, 0x6c, 0x65, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x6d, 0x75, 0x7a, 0x7a, 0x6c, 0x65, 0x5f, 0x76,
	0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0e, 0x6d,
	0x75, 0x7a, 0x7a, 0x6c, 0x65, 0x56, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a,
	0x06, 0x70, 0x6f, 0x77, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70,
	0x6f, 0x77, 0x64, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x6f, 0x77, 0x64, 0x65, 0x72, 0x5f,
	0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x70, 0x6f,
	0x77, 0x64, 0x65, 0x72, 0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x22, 0x79, 0x0a, 0x06, 0x42, 0x75,
	0x6c, 0x6c, 0x65, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x61, 0x6c, 0x69, 0x62, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x63, 0x61, 0x6c, 0x69, 0x62, 0x65, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06,
	0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x25, 0x0a, 0x02, 0x62, 0x63, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x42, 0x61, 0x6c, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x63, 0x43, 0x6f,
	0x65, 0x66, 0x66, 0x69, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x02, 0x62, 0x63, 0x12, 0x16, 0x0a,
	0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x6c,
	0x65, 0x6e, 0x67, 0x74, 0x68, 0x22, 0x57, 0x0a, 0x14, 0x42, 0x61, 0x6c, 0x6c, 0x69, 0x73, 0x74,
	0x69, 0x63, 0x43, 0x6f, 0x65, 0x66, 0x66, 0x69, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x29, 0x0a, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x44, 0x72, 0x61, 0x67, 0x46, 0x75, 0x6e, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2a, 0x59,
	0x0a, 0x0c, 0x44, 0x72, 0x61, 0x67, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d,
	0x0a, 0x19, 0x44, 0x52, 0x41, 0x47, 0x5f, 0x46, 0x55, 0x4e, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a,
	0x10, 0x44, 0x52, 0x41, 0x47, 0x5f, 0x46, 0x55, 0x4e, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x47,
	0x31, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x44, 0x52, 0x41, 0x47, 0x5f, 0x46, 0x55, 0x4e, 0x43,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x47, 0x37, 0x10, 0x07, 0x32, 0x6b, 0x0a, 0x10, 0x42, 0x61, 0x6c,
	0x6c, 0x69, 0x73, 0x74, 0x69, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x57, 0x0a,
	0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x61, 0x64, 0x12, 0x12, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x13, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x3a, 0x01, 0x2a, 0x22,
	0x15, 0x2f, 0x62, 0x61, 0x6c, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x63, 0x2f, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x42, 0x42, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x66, 0x72, 0x65, 0x61, 0x6b, 0x30,
	0x30, 0x33, 0x2f, 0x62, 0x61, 0x6c, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x63, 0x2d, 0x63, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_pb_service_proto_rawDescOnce sync.Once
	file_pb_service_proto_rawDescData = file_pb_service_proto_rawDesc
)

func file_pb_service_proto_rawDescGZIP() []byte {
	file_pb_service_proto_rawDescOnce.Do(func() {
		file_pb_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_service_proto_rawDescData)
	})
	return file_pb_service_proto_rawDescData
}

var file_pb_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pb_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pb_service_proto_goTypes = []interface{}{
	(DragFunction)(0),            // 0: DragFunction
	(*CreateLoadRequest)(nil),    // 1: CreateLoadRequest
	(*CreateLoadResponse)(nil),   // 2: CreateLoadResponse
	(*Load)(nil),                 // 3: Load
	(*Bullet)(nil),               // 4: Bullet
	(*BallisticCoefficient)(nil), // 5: BallisticCoefficient
}
var file_pb_service_proto_depIdxs = []int32{
	3, // 0: CreateLoadRequest.load:type_name -> Load
	3, // 1: CreateLoadResponse.load:type_name -> Load
	4, // 2: Load.bullet:type_name -> Bullet
	5, // 3: Bullet.bc:type_name -> BallisticCoefficient
	0, // 4: BallisticCoefficient.function:type_name -> DragFunction
	1, // 5: BallisticService.CreateLoad:input_type -> CreateLoadRequest
	2, // 6: BallisticService.CreateLoad:output_type -> CreateLoadResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_pb_service_proto_init() }
func file_pb_service_proto_init() {
	if File_pb_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLoadRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLoadResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Load); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bullet); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BallisticCoefficient); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_service_proto_goTypes,
		DependencyIndexes: file_pb_service_proto_depIdxs,
		EnumInfos:         file_pb_service_proto_enumTypes,
		MessageInfos:      file_pb_service_proto_msgTypes,
	}.Build()
	File_pb_service_proto = out.File
	file_pb_service_proto_rawDesc = nil
	file_pb_service_proto_goTypes = nil
	file_pb_service_proto_depIdxs = nil
}
