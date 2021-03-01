// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: Services/protos/ProdService.proto

//import "Services/protos/Models.proto";

package Services

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

//这个因为import失败，所以直接复制Models.proto
type ProdsModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProdID   int32  `protobuf:"varint,1,opt,name=ProdID,proto3" json:"ProdID,omitempty"`
	ProdName string `protobuf:"bytes,2,opt,name=ProdName,proto3" json:"ProdName,omitempty"`
}

func (x *ProdsModel) Reset() {
	*x = ProdsModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Services_protos_ProdService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProdsModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProdsModel) ProtoMessage() {}

func (x *ProdsModel) ProtoReflect() protoreflect.Message {
	mi := &file_Services_protos_ProdService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProdsModel.ProtoReflect.Descriptor instead.
func (*ProdsModel) Descriptor() ([]byte, []int) {
	return file_Services_protos_ProdService_proto_rawDescGZIP(), []int{0}
}

func (x *ProdsModel) GetProdID() int32 {
	if x != nil {
		return x.ProdID
	}
	return 0
}

func (x *ProdsModel) GetProdName() string {
	if x != nil {
		return x.ProdName
	}
	return ""
}

type ProdsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//@inject_tag: json:"size", form:"size"
	Size int32 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	//@inject_tag: uri:"pid"
	ProdId int32 `protobuf:"varint,2,opt,name=prod_id,json=prodId,proto3" json:"prod_id,omitempty"`
}

func (x *ProdsRequest) Reset() {
	*x = ProdsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Services_protos_ProdService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProdsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProdsRequest) ProtoMessage() {}

func (x *ProdsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Services_protos_ProdService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProdsRequest.ProtoReflect.Descriptor instead.
func (*ProdsRequest) Descriptor() ([]byte, []int) {
	return file_Services_protos_ProdService_proto_rawDescGZIP(), []int{1}
}

func (x *ProdsRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ProdsRequest) GetProdId() int32 {
	if x != nil {
		return x.ProdId
	}
	return 0
}

type ProdListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*ProdsModel `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ProdListResponse) Reset() {
	*x = ProdListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Services_protos_ProdService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProdListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProdListResponse) ProtoMessage() {}

func (x *ProdListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Services_protos_ProdService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProdListResponse.ProtoReflect.Descriptor instead.
func (*ProdListResponse) Descriptor() ([]byte, []int) {
	return file_Services_protos_ProdService_proto_rawDescGZIP(), []int{2}
}

func (x *ProdListResponse) GetData() []*ProdsModel {
	if x != nil {
		return x.Data
	}
	return nil
}

type ProdDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *ProdsModel `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ProdDetailResponse) Reset() {
	*x = ProdDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Services_protos_ProdService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProdDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProdDetailResponse) ProtoMessage() {}

func (x *ProdDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Services_protos_ProdService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProdDetailResponse.ProtoReflect.Descriptor instead.
func (*ProdDetailResponse) Descriptor() ([]byte, []int) {
	return file_Services_protos_ProdService_proto_rawDescGZIP(), []int{3}
}

func (x *ProdDetailResponse) GetData() *ProdsModel {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_Services_protos_ProdService_proto protoreflect.FileDescriptor

var file_Services_protos_ProdService_proto_rawDesc = []byte{
	0x0a, 0x21, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x50, 0x72, 0x6f, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x40, 0x0a,
	0x0a, 0x50, 0x72, 0x6f, 0x64, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x50,
	0x72, 0x6f, 0x64, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x50, 0x72, 0x6f,
	0x64, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x72, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0x3b, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x72, 0x6f, 0x64, 0x49, 0x64, 0x22, 0x3c, 0x0a, 0x10,
	0x50, 0x72, 0x6f, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x73, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3e, 0x0a, 0x12, 0x50, 0x72,
	0x6f, 0x64, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x73, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x98, 0x01, 0x0a, 0x0b, 0x50,
	0x72, 0x6f, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x0c, 0x47, 0x65,
	0x74, 0x50, 0x72, 0x6f, 0x64, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72,
	0x6f, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12,
	0x16, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Services_protos_ProdService_proto_rawDescOnce sync.Once
	file_Services_protos_ProdService_proto_rawDescData = file_Services_protos_ProdService_proto_rawDesc
)

func file_Services_protos_ProdService_proto_rawDescGZIP() []byte {
	file_Services_protos_ProdService_proto_rawDescOnce.Do(func() {
		file_Services_protos_ProdService_proto_rawDescData = protoimpl.X.CompressGZIP(file_Services_protos_ProdService_proto_rawDescData)
	})
	return file_Services_protos_ProdService_proto_rawDescData
}

var file_Services_protos_ProdService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_Services_protos_ProdService_proto_goTypes = []interface{}{
	(*ProdsModel)(nil),         // 0: Services.ProdsModel
	(*ProdsRequest)(nil),       // 1: Services.ProdsRequest
	(*ProdListResponse)(nil),   // 2: Services.ProdListResponse
	(*ProdDetailResponse)(nil), // 3: Services.ProdDetailResponse
}
var file_Services_protos_ProdService_proto_depIdxs = []int32{
	0, // 0: Services.ProdListResponse.data:type_name -> Services.ProdsModel
	0, // 1: Services.ProdDetailResponse.data:type_name -> Services.ProdsModel
	1, // 2: Services.ProdService.GetProdsList:input_type -> Services.ProdsRequest
	1, // 3: Services.ProdService.GetProdDetail:input_type -> Services.ProdsRequest
	2, // 4: Services.ProdService.GetProdsList:output_type -> Services.ProdListResponse
	3, // 5: Services.ProdService.GetProdDetail:output_type -> Services.ProdDetailResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_Services_protos_ProdService_proto_init() }
func file_Services_protos_ProdService_proto_init() {
	if File_Services_protos_ProdService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Services_protos_ProdService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProdsModel); i {
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
		file_Services_protos_ProdService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProdsRequest); i {
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
		file_Services_protos_ProdService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProdListResponse); i {
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
		file_Services_protos_ProdService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProdDetailResponse); i {
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
			RawDescriptor: file_Services_protos_ProdService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Services_protos_ProdService_proto_goTypes,
		DependencyIndexes: file_Services_protos_ProdService_proto_depIdxs,
		MessageInfos:      file_Services_protos_ProdService_proto_msgTypes,
	}.Build()
	File_Services_protos_ProdService_proto = out.File
	file_Services_protos_ProdService_proto_rawDesc = nil
	file_Services_protos_ProdService_proto_goTypes = nil
	file_Services_protos_ProdService_proto_depIdxs = nil
}
