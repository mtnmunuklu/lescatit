// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: cat.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Category struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Url      string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Category string `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
	Created  int64  `protobuf:"varint,4,opt,name=created,proto3" json:"created,omitempty"`
	Updated  int64  `protobuf:"varint,5,opt,name=updated,proto3" json:"updated,omitempty"`
	Revision int64  `protobuf:"varint,6,opt,name=revision,proto3" json:"revision,omitempty"`
}

func (x *Category) Reset() {
	*x = Category{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Category) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Category) ProtoMessage() {}

func (x *Category) ProtoReflect() protoreflect.Message {
	mi := &file_cat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Category.ProtoReflect.Descriptor instead.
func (*Category) Descriptor() ([]byte, []int) {
	return file_cat_proto_rawDescGZIP(), []int{0}
}

func (x *Category) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Category) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Category) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *Category) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *Category) GetUpdated() int64 {
	if x != nil {
		return x.Updated
	}
	return 0
}

func (x *Category) GetRevision() int64 {
	if x != nil {
		return x.Revision
	}
	return 0
}

type GetCategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *GetCategoryRequest) Reset() {
	*x = GetCategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCategoryRequest) ProtoMessage() {}

func (x *GetCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCategoryRequest.ProtoReflect.Descriptor instead.
func (*GetCategoryRequest) Descriptor() ([]byte, []int) {
	return file_cat_proto_rawDescGZIP(), []int{1}
}

func (x *GetCategoryRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type ListUrlsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Category []string `protobuf:"bytes,1,rep,name=category,proto3" json:"category,omitempty"`
	Count    string   `protobuf:"bytes,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *ListUrlsRequest) Reset() {
	*x = ListUrlsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUrlsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUrlsRequest) ProtoMessage() {}

func (x *ListUrlsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUrlsRequest.ProtoReflect.Descriptor instead.
func (*ListUrlsRequest) Descriptor() ([]byte, []int) {
	return file_cat_proto_rawDescGZIP(), []int{2}
}

func (x *ListUrlsRequest) GetCategory() []string {
	if x != nil {
		return x.Category
	}
	return nil
}

func (x *ListUrlsRequest) GetCount() string {
	if x != nil {
		return x.Count
	}
	return ""
}

type UpdateCategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url      string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Category string `protobuf:"bytes,2,opt,name=category,proto3" json:"category,omitempty"`
}

func (x *UpdateCategoryRequest) Reset() {
	*x = UpdateCategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCategoryRequest) ProtoMessage() {}

func (x *UpdateCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCategoryRequest.ProtoReflect.Descriptor instead.
func (*UpdateCategoryRequest) Descriptor() ([]byte, []int) {
	return file_cat_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateCategoryRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *UpdateCategoryRequest) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

type DeleteUrlResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *DeleteUrlResponse) Reset() {
	*x = DeleteUrlResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUrlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUrlResponse) ProtoMessage() {}

func (x *DeleteUrlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cat_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUrlResponse.ProtoReflect.Descriptor instead.
func (*DeleteUrlResponse) Descriptor() ([]byte, []int) {
	return file_cat_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteUrlResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_cat_proto protoreflect.FileDescriptor

var file_cat_proto_rawDesc = []byte{
	0x0a, 0x09, 0x63, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22,
	0x98, 0x01, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x26, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x72, 0x6c, 0x22, 0x43, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x72, 0x6c, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x45, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x22, 0x25,
	0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x75, 0x72, 0x6c, 0x32, 0xd0, 0x02, 0x0a, 0x0a, 0x43, 0x61, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x3f, 0x0a, 0x17, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x4d, 0x69, 0x73, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x70,
	0x62, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x2f, 0x0a, 0x08, 0x4c, 0x69,
	0x73, 0x74, 0x55, 0x72, 0x6c, 0x73, 0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x55, 0x72, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x30, 0x01, 0x12, 0x39, 0x0a, 0x0e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x19, 0x2e,
	0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x24, 0x0a, 0x06, 0x41, 0x64, 0x64, 0x55, 0x72, 0x6c,
	0x12, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x1a, 0x0c,
	0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x3a, 0x0a, 0x09,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x72, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cat_proto_rawDescOnce sync.Once
	file_cat_proto_rawDescData = file_cat_proto_rawDesc
)

func file_cat_proto_rawDescGZIP() []byte {
	file_cat_proto_rawDescOnce.Do(func() {
		file_cat_proto_rawDescData = protoimpl.X.CompressGZIP(file_cat_proto_rawDescData)
	})
	return file_cat_proto_rawDescData
}

var file_cat_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_cat_proto_goTypes = []interface{}{
	(*Category)(nil),              // 0: pb.Category
	(*GetCategoryRequest)(nil),    // 1: pb.GetCategoryRequest
	(*ListUrlsRequest)(nil),       // 2: pb.ListUrlsRequest
	(*UpdateCategoryRequest)(nil), // 3: pb.UpdateCategoryRequest
	(*DeleteUrlResponse)(nil),     // 4: pb.DeleteUrlResponse
}
var file_cat_proto_depIdxs = []int32{
	1, // 0: pb.CatService.GetCategory:input_type -> pb.GetCategoryRequest
	1, // 1: pb.CatService.ReportMiscategorization:input_type -> pb.GetCategoryRequest
	2, // 2: pb.CatService.ListUrls:input_type -> pb.ListUrlsRequest
	3, // 3: pb.CatService.UpdateCategory:input_type -> pb.UpdateCategoryRequest
	0, // 4: pb.CatService.AddUrl:input_type -> pb.Category
	1, // 5: pb.CatService.DeleteUrl:input_type -> pb.GetCategoryRequest
	0, // 6: pb.CatService.GetCategory:output_type -> pb.Category
	0, // 7: pb.CatService.ReportMiscategorization:output_type -> pb.Category
	0, // 8: pb.CatService.ListUrls:output_type -> pb.Category
	0, // 9: pb.CatService.UpdateCategory:output_type -> pb.Category
	0, // 10: pb.CatService.AddUrl:output_type -> pb.Category
	4, // 11: pb.CatService.DeleteUrl:output_type -> pb.DeleteUrlResponse
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cat_proto_init() }
func file_cat_proto_init() {
	if File_cat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Category); i {
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
		file_cat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCategoryRequest); i {
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
		file_cat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUrlsRequest); i {
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
		file_cat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCategoryRequest); i {
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
		file_cat_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUrlResponse); i {
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
			RawDescriptor: file_cat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cat_proto_goTypes,
		DependencyIndexes: file_cat_proto_depIdxs,
		MessageInfos:      file_cat_proto_msgTypes,
	}.Build()
	File_cat_proto = out.File
	file_cat_proto_rawDesc = nil
	file_cat_proto_goTypes = nil
	file_cat_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CatServiceClient is the client API for CatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CatServiceClient interface {
	GetCategory(ctx context.Context, in *GetCategoryRequest, opts ...grpc.CallOption) (*Category, error)
	ReportMiscategorization(ctx context.Context, in *GetCategoryRequest, opts ...grpc.CallOption) (*Category, error)
	ListUrls(ctx context.Context, in *ListUrlsRequest, opts ...grpc.CallOption) (CatService_ListUrlsClient, error)
	UpdateCategory(ctx context.Context, in *UpdateCategoryRequest, opts ...grpc.CallOption) (*Category, error)
	AddUrl(ctx context.Context, in *Category, opts ...grpc.CallOption) (*Category, error)
	DeleteUrl(ctx context.Context, in *GetCategoryRequest, opts ...grpc.CallOption) (*DeleteUrlResponse, error)
}

type catServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCatServiceClient(cc grpc.ClientConnInterface) CatServiceClient {
	return &catServiceClient{cc}
}

func (c *catServiceClient) GetCategory(ctx context.Context, in *GetCategoryRequest, opts ...grpc.CallOption) (*Category, error) {
	out := new(Category)
	err := c.cc.Invoke(ctx, "/pb.CatService/GetCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catServiceClient) ReportMiscategorization(ctx context.Context, in *GetCategoryRequest, opts ...grpc.CallOption) (*Category, error) {
	out := new(Category)
	err := c.cc.Invoke(ctx, "/pb.CatService/ReportMiscategorization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catServiceClient) ListUrls(ctx context.Context, in *ListUrlsRequest, opts ...grpc.CallOption) (CatService_ListUrlsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CatService_serviceDesc.Streams[0], "/pb.CatService/ListUrls", opts...)
	if err != nil {
		return nil, err
	}
	x := &catServiceListUrlsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CatService_ListUrlsClient interface {
	Recv() (*Category, error)
	grpc.ClientStream
}

type catServiceListUrlsClient struct {
	grpc.ClientStream
}

func (x *catServiceListUrlsClient) Recv() (*Category, error) {
	m := new(Category)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *catServiceClient) UpdateCategory(ctx context.Context, in *UpdateCategoryRequest, opts ...grpc.CallOption) (*Category, error) {
	out := new(Category)
	err := c.cc.Invoke(ctx, "/pb.CatService/UpdateCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catServiceClient) AddUrl(ctx context.Context, in *Category, opts ...grpc.CallOption) (*Category, error) {
	out := new(Category)
	err := c.cc.Invoke(ctx, "/pb.CatService/AddUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catServiceClient) DeleteUrl(ctx context.Context, in *GetCategoryRequest, opts ...grpc.CallOption) (*DeleteUrlResponse, error) {
	out := new(DeleteUrlResponse)
	err := c.cc.Invoke(ctx, "/pb.CatService/DeleteUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CatServiceServer is the server API for CatService service.
type CatServiceServer interface {
	GetCategory(context.Context, *GetCategoryRequest) (*Category, error)
	ReportMiscategorization(context.Context, *GetCategoryRequest) (*Category, error)
	ListUrls(*ListUrlsRequest, CatService_ListUrlsServer) error
	UpdateCategory(context.Context, *UpdateCategoryRequest) (*Category, error)
	AddUrl(context.Context, *Category) (*Category, error)
	DeleteUrl(context.Context, *GetCategoryRequest) (*DeleteUrlResponse, error)
}

// UnimplementedCatServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCatServiceServer struct {
}

func (*UnimplementedCatServiceServer) GetCategory(context.Context, *GetCategoryRequest) (*Category, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategory not implemented")
}
func (*UnimplementedCatServiceServer) ReportMiscategorization(context.Context, *GetCategoryRequest) (*Category, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportMiscategorization not implemented")
}
func (*UnimplementedCatServiceServer) ListUrls(*ListUrlsRequest, CatService_ListUrlsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListUrls not implemented")
}
func (*UnimplementedCatServiceServer) UpdateCategory(context.Context, *UpdateCategoryRequest) (*Category, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCategory not implemented")
}
func (*UnimplementedCatServiceServer) AddUrl(context.Context, *Category) (*Category, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUrl not implemented")
}
func (*UnimplementedCatServiceServer) DeleteUrl(context.Context, *GetCategoryRequest) (*DeleteUrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUrl not implemented")
}

func RegisterCatServiceServer(s *grpc.Server, srv CatServiceServer) {
	s.RegisterService(&_CatService_serviceDesc, srv)
}

func _CatService_GetCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatServiceServer).GetCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CatService/GetCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatServiceServer).GetCategory(ctx, req.(*GetCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatService_ReportMiscategorization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatServiceServer).ReportMiscategorization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CatService/ReportMiscategorization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatServiceServer).ReportMiscategorization(ctx, req.(*GetCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatService_ListUrls_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListUrlsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CatServiceServer).ListUrls(m, &catServiceListUrlsServer{stream})
}

type CatService_ListUrlsServer interface {
	Send(*Category) error
	grpc.ServerStream
}

type catServiceListUrlsServer struct {
	grpc.ServerStream
}

func (x *catServiceListUrlsServer) Send(m *Category) error {
	return x.ServerStream.SendMsg(m)
}

func _CatService_UpdateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatServiceServer).UpdateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CatService/UpdateCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatServiceServer).UpdateCategory(ctx, req.(*UpdateCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatService_AddUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Category)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatServiceServer).AddUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CatService/AddUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatServiceServer).AddUrl(ctx, req.(*Category))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatService_DeleteUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatServiceServer).DeleteUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CatService/DeleteUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatServiceServer).DeleteUrl(ctx, req.(*GetCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CatService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.CatService",
	HandlerType: (*CatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCategory",
			Handler:    _CatService_GetCategory_Handler,
		},
		{
			MethodName: "ReportMiscategorization",
			Handler:    _CatService_ReportMiscategorization_Handler,
		},
		{
			MethodName: "UpdateCategory",
			Handler:    _CatService_UpdateCategory_Handler,
		},
		{
			MethodName: "AddUrl",
			Handler:    _CatService_AddUrl_Handler,
		},
		{
			MethodName: "DeleteUrl",
			Handler:    _CatService_DeleteUrl_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListUrls",
			Handler:       _CatService_ListUrls_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cat.proto",
}
