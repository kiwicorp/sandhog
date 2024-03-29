// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: selftechio/naas/naas.proto

package naas

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Information about a sandhog.
type SandhogInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid     string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Hostname string `protobuf:"bytes,2,opt,name=hostname,proto3" json:"hostname,omitempty"`
}

func (x *SandhogInfo) Reset() {
	*x = SandhogInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_selftechio_naas_naas_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SandhogInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SandhogInfo) ProtoMessage() {}

func (x *SandhogInfo) ProtoReflect() protoreflect.Message {
	mi := &file_selftechio_naas_naas_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SandhogInfo.ProtoReflect.Descriptor instead.
func (*SandhogInfo) Descriptor() ([]byte, []int) {
	return file_selftechio_naas_naas_proto_rawDescGZIP(), []int{0}
}

func (x *SandhogInfo) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *SandhogInfo) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

// A request made by a sandhog to register itself in the registry.
type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *SandhogInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_selftechio_naas_naas_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_selftechio_naas_naas_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_selftechio_naas_naas_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterRequest) GetInfo() *SandhogInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

// A response received by the sandhog from the registry upon registering.
type RegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_selftechio_naas_naas_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_selftechio_naas_naas_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_selftechio_naas_naas_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterResponse) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

var File_selftechio_naas_naas_proto protoreflect.FileDescriptor

var file_selftechio_naas_naas_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x65, 0x6c, 0x66, 0x74, 0x65, 0x63, 0x68, 0x69, 0x6f, 0x2f, 0x6e, 0x61, 0x61,
	0x73, 0x2f, 0x6e, 0x61, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x73, 0x65,
	0x6c, 0x66, 0x74, 0x65, 0x63, 0x68, 0x69, 0x6f, 0x2e, 0x6e, 0x61, 0x61, 0x73, 0x1a, 0x1c, 0x73,
	0x65, 0x6c, 0x66, 0x74, 0x65, 0x63, 0x68, 0x69, 0x6f, 0x2f, 0x6e, 0x61, 0x61, 0x73, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3d, 0x0a, 0x0b, 0x53,
	0x61, 0x6e, 0x64, 0x68, 0x6f, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x43, 0x0a, 0x0f, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x65,
	0x6c, 0x66, 0x74, 0x65, 0x63, 0x68, 0x69, 0x6f, 0x2e, 0x6e, 0x61, 0x61, 0x73, 0x2e, 0x53, 0x61,
	0x6e, 0x64, 0x68, 0x6f, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22,
	0x43, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x65, 0x6c, 0x66, 0x74, 0x65, 0x63, 0x68, 0x69, 0x6f,
	0x2e, 0x6e, 0x61, 0x61, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x32, 0x5d, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x12, 0x51, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x20, 0x2e, 0x73,
	0x65, 0x6c, 0x66, 0x74, 0x65, 0x63, 0x68, 0x69, 0x6f, 0x2e, 0x6e, 0x61, 0x61, 0x73, 0x2e, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21,
	0x2e, 0x73, 0x65, 0x6c, 0x66, 0x74, 0x65, 0x63, 0x68, 0x69, 0x6f, 0x2e, 0x6e, 0x61, 0x61, 0x73,
	0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x65, 0x6c, 0x66, 0x74, 0x65, 0x63, 0x68, 0x69, 0x6f, 0x2f, 0x73, 0x61, 0x6e,
	0x64, 0x68, 0x6f, 0x67, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x6e, 0x61, 0x61, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_selftechio_naas_naas_proto_rawDescOnce sync.Once
	file_selftechio_naas_naas_proto_rawDescData = file_selftechio_naas_naas_proto_rawDesc
)

func file_selftechio_naas_naas_proto_rawDescGZIP() []byte {
	file_selftechio_naas_naas_proto_rawDescOnce.Do(func() {
		file_selftechio_naas_naas_proto_rawDescData = protoimpl.X.CompressGZIP(file_selftechio_naas_naas_proto_rawDescData)
	})
	return file_selftechio_naas_naas_proto_rawDescData
}

var file_selftechio_naas_naas_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_selftechio_naas_naas_proto_goTypes = []interface{}{
	(*SandhogInfo)(nil),      // 0: selftechio.naas.SandhogInfo
	(*RegisterRequest)(nil),  // 1: selftechio.naas.RegisterRequest
	(*RegisterResponse)(nil), // 2: selftechio.naas.RegisterResponse
	(*Status)(nil),           // 3: selftechio.naas.Status
}
var file_selftechio_naas_naas_proto_depIdxs = []int32{
	0, // 0: selftechio.naas.RegisterRequest.info:type_name -> selftechio.naas.SandhogInfo
	3, // 1: selftechio.naas.RegisterResponse.status:type_name -> selftechio.naas.Status
	1, // 2: selftechio.naas.Registry.Register:input_type -> selftechio.naas.RegisterRequest
	2, // 3: selftechio.naas.Registry.Register:output_type -> selftechio.naas.RegisterResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_selftechio_naas_naas_proto_init() }
func file_selftechio_naas_naas_proto_init() {
	if File_selftechio_naas_naas_proto != nil {
		return
	}
	file_selftechio_naas_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_selftechio_naas_naas_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SandhogInfo); i {
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
		file_selftechio_naas_naas_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_selftechio_naas_naas_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResponse); i {
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
			RawDescriptor: file_selftechio_naas_naas_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_selftechio_naas_naas_proto_goTypes,
		DependencyIndexes: file_selftechio_naas_naas_proto_depIdxs,
		MessageInfos:      file_selftechio_naas_naas_proto_msgTypes,
	}.Build()
	File_selftechio_naas_naas_proto = out.File
	file_selftechio_naas_naas_proto_rawDesc = nil
	file_selftechio_naas_naas_proto_goTypes = nil
	file_selftechio_naas_naas_proto_depIdxs = nil
}
