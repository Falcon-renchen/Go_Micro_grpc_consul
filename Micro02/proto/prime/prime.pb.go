// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: Micro02/proto/prime/prime.proto

package prime

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

type PrimeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Input int64 `protobuf:"varint,1,opt,name=input,proto3" json:"input,omitempty"`
}

func (x *PrimeRequest) Reset() {
	*x = PrimeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Micro02_proto_prime_prime_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimeRequest) ProtoMessage() {}

func (x *PrimeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Micro02_proto_prime_prime_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimeRequest.ProtoReflect.Descriptor instead.
func (*PrimeRequest) Descriptor() ([]byte, []int) {
	return file_Micro02_proto_prime_prime_proto_rawDescGZIP(), []int{0}
}

func (x *PrimeRequest) GetInput() int64 {
	if x != nil {
		return x.Input
	}
	return 0
}

type PrimeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Output []int64 `protobuf:"varint,1,rep,packed,name=output,proto3" json:"output,omitempty"`
}

func (x *PrimeResponse) Reset() {
	*x = PrimeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Micro02_proto_prime_prime_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimeResponse) ProtoMessage() {}

func (x *PrimeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Micro02_proto_prime_prime_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimeResponse.ProtoReflect.Descriptor instead.
func (*PrimeResponse) Descriptor() ([]byte, []int) {
	return file_Micro02_proto_prime_prime_proto_rawDescGZIP(), []int{1}
}

func (x *PrimeResponse) GetOutput() []int64 {
	if x != nil {
		return x.Output
	}
	return nil
}

var File_Micro02_proto_prime_prime_proto protoreflect.FileDescriptor

var file_Micro02_proto_prime_prime_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x30, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x70, 0x72, 0x69, 0x6d, 0x65, 0x2f, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x24, 0x0a, 0x0c, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x27, 0x0a, 0x0d, 0x50, 0x72, 0x69, 0x6d, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x32, 0x34, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x69, 0x6d, 0x65, 0x12, 0x0d, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Micro02_proto_prime_prime_proto_rawDescOnce sync.Once
	file_Micro02_proto_prime_prime_proto_rawDescData = file_Micro02_proto_prime_prime_proto_rawDesc
)

func file_Micro02_proto_prime_prime_proto_rawDescGZIP() []byte {
	file_Micro02_proto_prime_prime_proto_rawDescOnce.Do(func() {
		file_Micro02_proto_prime_prime_proto_rawDescData = protoimpl.X.CompressGZIP(file_Micro02_proto_prime_prime_proto_rawDescData)
	})
	return file_Micro02_proto_prime_prime_proto_rawDescData
}

var file_Micro02_proto_prime_prime_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_Micro02_proto_prime_prime_proto_goTypes = []interface{}{
	(*PrimeRequest)(nil),  // 0: PrimeRequest
	(*PrimeResponse)(nil), // 1: PrimeResponse
}
var file_Micro02_proto_prime_prime_proto_depIdxs = []int32{
	0, // 0: Prime.GetPrime:input_type -> PrimeRequest
	1, // 1: Prime.GetPrime:output_type -> PrimeResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Micro02_proto_prime_prime_proto_init() }
func file_Micro02_proto_prime_prime_proto_init() {
	if File_Micro02_proto_prime_prime_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Micro02_proto_prime_prime_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimeRequest); i {
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
		file_Micro02_proto_prime_prime_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimeResponse); i {
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
			RawDescriptor: file_Micro02_proto_prime_prime_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Micro02_proto_prime_prime_proto_goTypes,
		DependencyIndexes: file_Micro02_proto_prime_prime_proto_depIdxs,
		MessageInfos:      file_Micro02_proto_prime_prime_proto_msgTypes,
	}.Build()
	File_Micro02_proto_prime_prime_proto = out.File
	file_Micro02_proto_prime_prime_proto_rawDesc = nil
	file_Micro02_proto_prime_prime_proto_goTypes = nil
	file_Micro02_proto_prime_prime_proto_depIdxs = nil
}
