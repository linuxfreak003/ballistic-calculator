// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: pb/solve.proto

package pb

import (
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

type SolveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScenarioId       int64 `protobuf:"varint,1,opt,name=scenario_id,json=scenarioId,proto3" json:"scenario_id,omitempty"`
	IncludeSpinDrift bool  `protobuf:"varint,2,opt,name=include_spin_drift,json=includeSpinDrift,proto3" json:"include_spin_drift,omitempty"`
	IncludeCoriolis  bool  `protobuf:"varint,3,opt,name=include_coriolis,json=includeCoriolis,proto3" json:"include_coriolis,omitempty"`
	Range            int64 `protobuf:"varint,4,opt,name=range,proto3" json:"range,omitempty"`
}

func (x *SolveRequest) Reset() {
	*x = SolveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_solve_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SolveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SolveRequest) ProtoMessage() {}

func (x *SolveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_solve_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SolveRequest.ProtoReflect.Descriptor instead.
func (*SolveRequest) Descriptor() ([]byte, []int) {
	return file_pb_solve_proto_rawDescGZIP(), []int{0}
}

func (x *SolveRequest) GetScenarioId() int64 {
	if x != nil {
		return x.ScenarioId
	}
	return 0
}

func (x *SolveRequest) GetIncludeSpinDrift() bool {
	if x != nil {
		return x.IncludeSpinDrift
	}
	return false
}

func (x *SolveRequest) GetIncludeCoriolis() bool {
	if x != nil {
		return x.IncludeCoriolis
	}
	return false
}

func (x *SolveRequest) GetRange() int64 {
	if x != nil {
		return x.Range
	}
	return 0
}

type SolveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Solution *Solution `protobuf:"bytes,1,opt,name=solution,proto3" json:"solution,omitempty"`
}

func (x *SolveResponse) Reset() {
	*x = SolveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_solve_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SolveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SolveResponse) ProtoMessage() {}

func (x *SolveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_solve_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SolveResponse.ProtoReflect.Descriptor instead.
func (*SolveResponse) Descriptor() ([]byte, []int) {
	return file_pb_solve_proto_rawDescGZIP(), []int{1}
}

func (x *SolveResponse) GetSolution() *Solution {
	if x != nil {
		return x.Solution
	}
	return nil
}

type SolveTableRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScenarioId       int64 `protobuf:"varint,1,opt,name=scenario_id,json=scenarioId,proto3" json:"scenario_id,omitempty"`
	IncludeSpinDrift bool  `protobuf:"varint,2,opt,name=include_spin_drift,json=includeSpinDrift,proto3" json:"include_spin_drift,omitempty"`
	IncludeCoriolis  bool  `protobuf:"varint,3,opt,name=include_coriolis,json=includeCoriolis,proto3" json:"include_coriolis,omitempty"`
	MaxRange         int64 `protobuf:"varint,4,opt,name=max_range,json=maxRange,proto3" json:"max_range,omitempty"`
	Increment        int64 `protobuf:"varint,5,opt,name=increment,proto3" json:"increment,omitempty"`
}

func (x *SolveTableRequest) Reset() {
	*x = SolveTableRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_solve_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SolveTableRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SolveTableRequest) ProtoMessage() {}

func (x *SolveTableRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_solve_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SolveTableRequest.ProtoReflect.Descriptor instead.
func (*SolveTableRequest) Descriptor() ([]byte, []int) {
	return file_pb_solve_proto_rawDescGZIP(), []int{2}
}

func (x *SolveTableRequest) GetScenarioId() int64 {
	if x != nil {
		return x.ScenarioId
	}
	return 0
}

func (x *SolveTableRequest) GetIncludeSpinDrift() bool {
	if x != nil {
		return x.IncludeSpinDrift
	}
	return false
}

func (x *SolveTableRequest) GetIncludeCoriolis() bool {
	if x != nil {
		return x.IncludeCoriolis
	}
	return false
}

func (x *SolveTableRequest) GetMaxRange() int64 {
	if x != nil {
		return x.MaxRange
	}
	return 0
}

func (x *SolveTableRequest) GetIncrement() int64 {
	if x != nil {
		return x.Increment
	}
	return 0
}

type SolveTableResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Solutions []*Solution `protobuf:"bytes,1,rep,name=solutions,proto3" json:"solutions,omitempty"`
}

func (x *SolveTableResponse) Reset() {
	*x = SolveTableResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_solve_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SolveTableResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SolveTableResponse) ProtoMessage() {}

func (x *SolveTableResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_solve_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SolveTableResponse.ProtoReflect.Descriptor instead.
func (*SolveTableResponse) Descriptor() ([]byte, []int) {
	return file_pb_solve_proto_rawDescGZIP(), []int{3}
}

func (x *SolveTableResponse) GetSolutions() []*Solution {
	if x != nil {
		return x.Solutions
	}
	return nil
}

type Solution struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Range      int64   `protobuf:"varint,1,opt,name=range,proto3" json:"range,omitempty"`                              // Range in yards
	RawRange   float64 `protobuf:"fixed64,2,opt,name=raw_range,json=rawRange,proto3" json:"raw_range,omitempty"`       // Range in yards (for calculations)
	Drop       float64 `protobuf:"fixed64,3,opt,name=drop,proto3" json:"drop,omitempty"`                               // Bullet drop in inches
	DropMoa    float64 `protobuf:"fixed64,4,opt,name=drop_moa,json=dropMoa,proto3" json:"drop_moa,omitempty"`          // Bullet drop in MOA
	Time       float64 `protobuf:"fixed64,5,opt,name=time,proto3" json:"time,omitempty"`                               // Time of flight in seconds
	Windage    float64 `protobuf:"fixed64,6,opt,name=windage,proto3" json:"windage,omitempty"`                         // Windage compensation in inches
	WindageMoa float64 `protobuf:"fixed64,7,opt,name=windage_moa,json=windageMoa,proto3" json:"windage_moa,omitempty"` // Windage compensation in MOA
	Energy     float64 `protobuf:"fixed64,8,opt,name=energy,proto3" json:"energy,omitempty"`                           // Energy in ft*lbs
	Velocity   float64 `protobuf:"fixed64,9,opt,name=velocity,proto3" json:"velocity,omitempty"`                       // Velocity in fps
	VelocityX  float64 `protobuf:"fixed64,10,opt,name=velocity_x,json=velocityX,proto3" json:"velocity_x,omitempty"`   // Velocity X-component
	VelocityY  float64 `protobuf:"fixed64,11,opt,name=velocity_y,json=velocityY,proto3" json:"velocity_y,omitempty"`   // Velocity Y-component
}

func (x *Solution) Reset() {
	*x = Solution{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_solve_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Solution) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Solution) ProtoMessage() {}

func (x *Solution) ProtoReflect() protoreflect.Message {
	mi := &file_pb_solve_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Solution.ProtoReflect.Descriptor instead.
func (*Solution) Descriptor() ([]byte, []int) {
	return file_pb_solve_proto_rawDescGZIP(), []int{4}
}

func (x *Solution) GetRange() int64 {
	if x != nil {
		return x.Range
	}
	return 0
}

func (x *Solution) GetRawRange() float64 {
	if x != nil {
		return x.RawRange
	}
	return 0
}

func (x *Solution) GetDrop() float64 {
	if x != nil {
		return x.Drop
	}
	return 0
}

func (x *Solution) GetDropMoa() float64 {
	if x != nil {
		return x.DropMoa
	}
	return 0
}

func (x *Solution) GetTime() float64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *Solution) GetWindage() float64 {
	if x != nil {
		return x.Windage
	}
	return 0
}

func (x *Solution) GetWindageMoa() float64 {
	if x != nil {
		return x.WindageMoa
	}
	return 0
}

func (x *Solution) GetEnergy() float64 {
	if x != nil {
		return x.Energy
	}
	return 0
}

func (x *Solution) GetVelocity() float64 {
	if x != nil {
		return x.Velocity
	}
	return 0
}

func (x *Solution) GetVelocityX() float64 {
	if x != nil {
		return x.VelocityX
	}
	return 0
}

func (x *Solution) GetVelocityY() float64 {
	if x != nil {
		return x.VelocityY
	}
	return 0
}

var File_pb_solve_proto protoreflect.FileDescriptor

var file_pb_solve_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x62, 0x2f, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x22, 0x9e, 0x01, 0x0a, 0x0c, 0x53, 0x6f, 0x6c, 0x76, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69,
	0x6f, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x63, 0x65, 0x6e,
	0x61, 0x72, 0x69, 0x6f, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64,
	0x65, 0x5f, 0x73, 0x70, 0x69, 0x6e, 0x5f, 0x64, 0x72, 0x69, 0x66, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x10, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x53, 0x70, 0x69, 0x6e, 0x44,
	0x72, 0x69, 0x66, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f,
	0x63, 0x6f, 0x72, 0x69, 0x6f, 0x6c, 0x69, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f,
	0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x43, 0x6f, 0x72, 0x69, 0x6f, 0x6c, 0x69, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x72, 0x61, 0x6e, 0x67, 0x65, 0x22, 0x39, 0x0a, 0x0d, 0x53, 0x6f, 0x6c, 0x76, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x08, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x6f,
	0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0xc8, 0x01, 0x0a, 0x11, 0x53, 0x6f, 0x6c, 0x76, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x63, 0x65, 0x6e, 0x61, 0x72,
	0x69, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x63, 0x65,
	0x6e, 0x61, 0x72, 0x69, 0x6f, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x69, 0x6e, 0x63, 0x6c, 0x75,
	0x64, 0x65, 0x5f, 0x73, 0x70, 0x69, 0x6e, 0x5f, 0x64, 0x72, 0x69, 0x66, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x10, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x53, 0x70, 0x69, 0x6e,
	0x44, 0x72, 0x69, 0x66, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65,
	0x5f, 0x63, 0x6f, 0x72, 0x69, 0x6f, 0x6c, 0x69, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0f, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x43, 0x6f, 0x72, 0x69, 0x6f, 0x6c, 0x69, 0x73,
	0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x69, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x69, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x40, 0x0a, 0x12, 0x53,
	0x6f, 0x6c, 0x76, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2a, 0x0a, 0x09, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x09, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xad, 0x02,
	0x0a, 0x08, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x61,
	0x6e, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x72, 0x61, 0x6e, 0x67, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x72, 0x61, 0x77, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x08, 0x72, 0x61, 0x77, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x72, 0x6f, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x64, 0x72, 0x6f,
	0x70, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x72, 0x6f, 0x70, 0x5f, 0x6d, 0x6f, 0x61, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x07, 0x64, 0x72, 0x6f, 0x70, 0x4d, 0x6f, 0x61, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x77, 0x69, 0x6e, 0x64, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x07, 0x77, 0x69, 0x6e, 0x64, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x69,
	0x6e, 0x64, 0x61, 0x67, 0x65, 0x5f, 0x6d, 0x6f, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0a, 0x77, 0x69, 0x6e, 0x64, 0x61, 0x67, 0x65, 0x4d, 0x6f, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x65,
	0x6e, 0x65, 0x72, 0x67, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x65, 0x6e, 0x65,
	0x72, 0x67, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x76, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x12,
	0x1d, 0x0a, 0x0a, 0x76, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x78, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x09, 0x76, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x58, 0x12, 0x1d,
	0x0a, 0x0a, 0x76, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x79, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x09, 0x76, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x59, 0x42, 0x6e, 0x0a,
	0x06, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x62, 0x42, 0x0a, 0x53, 0x6f, 0x6c, 0x76, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x66, 0x72, 0x65, 0x61, 0x6b, 0x30, 0x30, 0x33, 0x2f,
	0x62, 0x61, 0x6c, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x63, 0x2d, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c,
	0x61, 0x74, 0x6f, 0x72, 0x2f, 0x70, 0x62, 0xa2, 0x02, 0x03, 0x50, 0x58, 0x58, 0xaa, 0x02, 0x02,
	0x50, 0x62, 0xca, 0x02, 0x02, 0x50, 0x62, 0xe2, 0x02, 0x0e, 0x50, 0x62, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x02, 0x50, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_solve_proto_rawDescOnce sync.Once
	file_pb_solve_proto_rawDescData = file_pb_solve_proto_rawDesc
)

func file_pb_solve_proto_rawDescGZIP() []byte {
	file_pb_solve_proto_rawDescOnce.Do(func() {
		file_pb_solve_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_solve_proto_rawDescData)
	})
	return file_pb_solve_proto_rawDescData
}

var file_pb_solve_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pb_solve_proto_goTypes = []interface{}{
	(*SolveRequest)(nil),       // 0: pb.SolveRequest
	(*SolveResponse)(nil),      // 1: pb.SolveResponse
	(*SolveTableRequest)(nil),  // 2: pb.SolveTableRequest
	(*SolveTableResponse)(nil), // 3: pb.SolveTableResponse
	(*Solution)(nil),           // 4: pb.Solution
}
var file_pb_solve_proto_depIdxs = []int32{
	4, // 0: pb.SolveResponse.solution:type_name -> pb.Solution
	4, // 1: pb.SolveTableResponse.solutions:type_name -> pb.Solution
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pb_solve_proto_init() }
func file_pb_solve_proto_init() {
	if File_pb_solve_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_solve_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SolveRequest); i {
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
		file_pb_solve_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SolveResponse); i {
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
		file_pb_solve_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SolveTableRequest); i {
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
		file_pb_solve_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SolveTableResponse); i {
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
		file_pb_solve_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Solution); i {
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
			RawDescriptor: file_pb_solve_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_solve_proto_goTypes,
		DependencyIndexes: file_pb_solve_proto_depIdxs,
		MessageInfos:      file_pb_solve_proto_msgTypes,
	}.Build()
	File_pb_solve_proto = out.File
	file_pb_solve_proto_rawDesc = nil
	file_pb_solve_proto_goTypes = nil
	file_pb_solve_proto_depIdxs = nil
}
