// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: pb/scenario.proto

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

type CreateScenarioRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scenario *Scenario `protobuf:"bytes,1,opt,name=scenario,proto3" json:"scenario,omitempty"`
}

func (x *CreateScenarioRequest) Reset() {
	*x = CreateScenarioRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_scenario_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateScenarioRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateScenarioRequest) ProtoMessage() {}

func (x *CreateScenarioRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_scenario_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateScenarioRequest.ProtoReflect.Descriptor instead.
func (*CreateScenarioRequest) Descriptor() ([]byte, []int) {
	return file_pb_scenario_proto_rawDescGZIP(), []int{0}
}

func (x *CreateScenarioRequest) GetScenario() *Scenario {
	if x != nil {
		return x.Scenario
	}
	return nil
}

type CreateScenarioResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scenario *Scenario `protobuf:"bytes,1,opt,name=scenario,proto3" json:"scenario,omitempty"`
}

func (x *CreateScenarioResponse) Reset() {
	*x = CreateScenarioResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_scenario_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateScenarioResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateScenarioResponse) ProtoMessage() {}

func (x *CreateScenarioResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_scenario_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateScenarioResponse.ProtoReflect.Descriptor instead.
func (*CreateScenarioResponse) Descriptor() ([]byte, []int) {
	return file_pb_scenario_proto_rawDescGZIP(), []int{1}
}

func (x *CreateScenarioResponse) GetScenario() *Scenario {
	if x != nil {
		return x.Scenario
	}
	return nil
}

type ListScenariosRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListScenariosRequest) Reset() {
	*x = ListScenariosRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_scenario_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListScenariosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListScenariosRequest) ProtoMessage() {}

func (x *ListScenariosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_scenario_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListScenariosRequest.ProtoReflect.Descriptor instead.
func (*ListScenariosRequest) Descriptor() ([]byte, []int) {
	return file_pb_scenario_proto_rawDescGZIP(), []int{2}
}

type ListScenariosResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scenarios []*Scenario `protobuf:"bytes,1,rep,name=scenarios,proto3" json:"scenarios,omitempty"`
}

func (x *ListScenariosResponse) Reset() {
	*x = ListScenariosResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_scenario_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListScenariosResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListScenariosResponse) ProtoMessage() {}

func (x *ListScenariosResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_scenario_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListScenariosResponse.ProtoReflect.Descriptor instead.
func (*ListScenariosResponse) Descriptor() ([]byte, []int) {
	return file_pb_scenario_proto_rawDescGZIP(), []int{3}
}

func (x *ListScenariosResponse) GetScenarios() []*Scenario {
	if x != nil {
		return x.Scenarios
	}
	return nil
}

type Scenario struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScenarioId    int64  `protobuf:"varint,1,opt,name=scenario_id,json=scenarioId,proto3" json:"scenario_id,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	EnvironmentId int64  `protobuf:"varint,3,opt,name=environment_id,json=environmentId,proto3" json:"environment_id,omitempty"`
	RifleId       int64  `protobuf:"varint,4,opt,name=rifle_id,json=rifleId,proto3" json:"rifle_id,omitempty"`
	LoadId        int64  `protobuf:"varint,5,opt,name=load_id,json=loadId,proto3" json:"load_id,omitempty"`
}

func (x *Scenario) Reset() {
	*x = Scenario{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_scenario_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Scenario) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scenario) ProtoMessage() {}

func (x *Scenario) ProtoReflect() protoreflect.Message {
	mi := &file_pb_scenario_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scenario.ProtoReflect.Descriptor instead.
func (*Scenario) Descriptor() ([]byte, []int) {
	return file_pb_scenario_proto_rawDescGZIP(), []int{4}
}

func (x *Scenario) GetScenarioId() int64 {
	if x != nil {
		return x.ScenarioId
	}
	return 0
}

func (x *Scenario) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Scenario) GetEnvironmentId() int64 {
	if x != nil {
		return x.EnvironmentId
	}
	return 0
}

func (x *Scenario) GetRifleId() int64 {
	if x != nil {
		return x.RifleId
	}
	return 0
}

func (x *Scenario) GetLoadId() int64 {
	if x != nil {
		return x.LoadId
	}
	return 0
}

var File_pb_scenario_proto protoreflect.FileDescriptor

var file_pb_scenario_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x62, 0x2f, 0x73, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x41, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x28, 0x0a, 0x08, 0x73, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f,
	0x52, 0x08, 0x73, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x22, 0x42, 0x0a, 0x16, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x08, 0x73, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x63, 0x65, 0x6e,
	0x61, 0x72, 0x69, 0x6f, 0x52, 0x08, 0x73, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x22, 0x16,
	0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x43, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x63,
	0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2a, 0x0a, 0x09, 0x73, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f,
	0x52, 0x09, 0x73, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x73, 0x22, 0x9a, 0x01, 0x0a, 0x08,
	0x53, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x63, 0x65, 0x6e,
	0x61, 0x72, 0x69, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73,
	0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a,
	0x0e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x69, 0x66, 0x6c, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x69, 0x66, 0x6c, 0x65, 0x49, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x64, 0x42, 0x71, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x2e,
	0x70, 0x62, 0x42, 0x0d, 0x53, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6c, 0x69, 0x6e, 0x75, 0x78, 0x66, 0x72, 0x65, 0x61, 0x6b, 0x30, 0x30, 0x33, 0x2f, 0x62, 0x61,
	0x6c, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x63, 0x2d, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x6f, 0x72, 0x2f, 0x70, 0x62, 0xa2, 0x02, 0x03, 0x50, 0x58, 0x58, 0xaa, 0x02, 0x02, 0x50, 0x62,
	0xca, 0x02, 0x02, 0x50, 0x62, 0xe2, 0x02, 0x0e, 0x50, 0x62, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x02, 0x50, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_pb_scenario_proto_rawDescOnce sync.Once
	file_pb_scenario_proto_rawDescData = file_pb_scenario_proto_rawDesc
)

func file_pb_scenario_proto_rawDescGZIP() []byte {
	file_pb_scenario_proto_rawDescOnce.Do(func() {
		file_pb_scenario_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_scenario_proto_rawDescData)
	})
	return file_pb_scenario_proto_rawDescData
}

var file_pb_scenario_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pb_scenario_proto_goTypes = []interface{}{
	(*CreateScenarioRequest)(nil),  // 0: pb.CreateScenarioRequest
	(*CreateScenarioResponse)(nil), // 1: pb.CreateScenarioResponse
	(*ListScenariosRequest)(nil),   // 2: pb.ListScenariosRequest
	(*ListScenariosResponse)(nil),  // 3: pb.ListScenariosResponse
	(*Scenario)(nil),               // 4: pb.Scenario
}
var file_pb_scenario_proto_depIdxs = []int32{
	4, // 0: pb.CreateScenarioRequest.scenario:type_name -> pb.Scenario
	4, // 1: pb.CreateScenarioResponse.scenario:type_name -> pb.Scenario
	4, // 2: pb.ListScenariosResponse.scenarios:type_name -> pb.Scenario
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pb_scenario_proto_init() }
func file_pb_scenario_proto_init() {
	if File_pb_scenario_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_scenario_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateScenarioRequest); i {
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
		file_pb_scenario_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateScenarioResponse); i {
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
		file_pb_scenario_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListScenariosRequest); i {
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
		file_pb_scenario_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListScenariosResponse); i {
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
		file_pb_scenario_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Scenario); i {
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
			RawDescriptor: file_pb_scenario_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_scenario_proto_goTypes,
		DependencyIndexes: file_pb_scenario_proto_depIdxs,
		MessageInfos:      file_pb_scenario_proto_msgTypes,
	}.Build()
	File_pb_scenario_proto = out.File
	file_pb_scenario_proto_rawDesc = nil
	file_pb_scenario_proto_goTypes = nil
	file_pb_scenario_proto_depIdxs = nil
}
