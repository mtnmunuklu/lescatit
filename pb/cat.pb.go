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

type DeleteUrlResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *DeleteUrlResponse) Reset() {
	*x = DeleteUrlResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUrlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUrlResponse) ProtoMessage() {}

func (x *DeleteUrlResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DeleteUrlResponse.ProtoReflect.Descriptor instead.
func (*DeleteUrlResponse) Descriptor() ([]byte, []int) {
	return file_cat_proto_rawDescGZIP(), []int{3}
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
	0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x25, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x32, 0xc8,
	0x02, 0x0a, 0x0f, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x33, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x3f, 0x0a, 0x17, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x4d, 0x69, 0x73, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x2f, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74,
	0x55, 0x72, 0x6c, 0x73, 0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x72,
	0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x30, 0x01, 0x12, 0x2c, 0x0a, 0x0e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x0c, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x24, 0x0a, 0x06, 0x41, 0x64, 0x64, 0x55, 0x72,
	0x6c, 0x12, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x1a,
	0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x3a, 0x0a,
	0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e,
	0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x72,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_cat_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_cat_proto_goTypes = []interface{}{
	(*Category)(nil),           // 0: pb.Category
	(*GetCategoryRequest)(nil), // 1: pb.GetCategoryRequest
	(*ListUrlsRequest)(nil),    // 2: pb.ListUrlsRequest
	(*DeleteUrlResponse)(nil),  // 3: pb.DeleteUrlResponse
}
var file_cat_proto_depIdxs = []int32{
	1, // 0: pb.CategoryService.GetCategory:input_type -> pb.GetCategoryRequest
	1, // 1: pb.CategoryService.ReportMiscategorization:input_type -> pb.GetCategoryRequest
	2, // 2: pb.CategoryService.ListUrls:input_type -> pb.ListUrlsRequest
	0, // 3: pb.CategoryService.UpdateCategory:input_type -> pb.Category
	0, // 4: pb.CategoryService.AddUrl:input_type -> pb.Category
	1, // 5: pb.CategoryService.DeleteUrl:input_type -> pb.GetCategoryRequest
	0, // 6: pb.CategoryService.GetCategory:output_type -> pb.Category
	0, // 7: pb.CategoryService.ReportMiscategorization:output_type -> pb.Category
	0, // 8: pb.CategoryService.ListUrls:output_type -> pb.Category
	0, // 9: pb.CategoryService.UpdateCategory:output_type -> pb.Category
	0, // 10: pb.CategoryService.AddUrl:output_type -> pb.Category
	3, // 11: pb.CategoryService.DeleteUrl:output_type -> pb.DeleteUrlResponse
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
			NumMessages:   4,
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

// CategoryServiceClient is the client API for CategoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CategoryServiceClient interface {
	GetCategory(ctx context.Context, in *GetCategoryRequest, opts ...grpc.CallOption) (*Category, error)
	ReportMiscategorization(ctx context.Context, in *GetCategoryRequest, opts ...grpc.CallOption) (*Category, error)
	ListUrls(ctx context.Context, in *ListUrlsRequest, opts ...grpc.CallOption) (CategoryService_ListUrlsClient, error)
	UpdateCategory(ctx context.Context, in *Category, opts ...grpc.CallOption) (*Category, error)
	AddUrl(ctx context.Context, in *Category, opts ...grpc.CallOption) (*Category, error)
	DeleteUrl(ctx context.Context, in *GetCategoryRequest, opts ...grpc.CallOption) (*DeleteUrlResponse, error)
}

type categoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCategoryServiceClient(cc grpc.ClientConnInterface) CategoryServiceClient {
	return &categoryServiceClient{cc}
}

func (c *categoryServiceClient) GetCategory(ctx context.Context, in *GetCategoryRequest, opts ...grpc.CallOption) (*Category, error) {
	out := new(Category)
	err := c.cc.Invoke(ctx, "/pb.CategoryService/GetCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) ReportMiscategorization(ctx context.Context, in *GetCategoryRequest, opts ...grpc.CallOption) (*Category, error) {
	out := new(Category)
	err := c.cc.Invoke(ctx, "/pb.CategoryService/ReportMiscategorization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) ListUrls(ctx context.Context, in *ListUrlsRequest, opts ...grpc.CallOption) (CategoryService_ListUrlsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CategoryService_serviceDesc.Streams[0], "/pb.CategoryService/ListUrls", opts...)
	if err != nil {
		return nil, err
	}
	x := &categoryServiceListUrlsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CategoryService_ListUrlsClient interface {
	Recv() (*Category, error)
	grpc.ClientStream
}

type categoryServiceListUrlsClient struct {
	grpc.ClientStream
}

func (x *categoryServiceListUrlsClient) Recv() (*Category, error) {
	m := new(Category)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *categoryServiceClient) UpdateCategory(ctx context.Context, in *Category, opts ...grpc.CallOption) (*Category, error) {
	out := new(Category)
	err := c.cc.Invoke(ctx, "/pb.CategoryService/UpdateCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) AddUrl(ctx context.Context, in *Category, opts ...grpc.CallOption) (*Category, error) {
	out := new(Category)
	err := c.cc.Invoke(ctx, "/pb.CategoryService/AddUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) DeleteUrl(ctx context.Context, in *GetCategoryRequest, opts ...grpc.CallOption) (*DeleteUrlResponse, error) {
	out := new(DeleteUrlResponse)
	err := c.cc.Invoke(ctx, "/pb.CategoryService/DeleteUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CategoryServiceServer is the server API for CategoryService service.
type CategoryServiceServer interface {
	GetCategory(context.Context, *GetCategoryRequest) (*Category, error)
	ReportMiscategorization(context.Context, *GetCategoryRequest) (*Category, error)
	ListUrls(*ListUrlsRequest, CategoryService_ListUrlsServer) error
	UpdateCategory(context.Context, *Category) (*Category, error)
	AddUrl(context.Context, *Category) (*Category, error)
	DeleteUrl(context.Context, *GetCategoryRequest) (*DeleteUrlResponse, error)
}

// UnimplementedCategoryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCategoryServiceServer struct {
}

func (*UnimplementedCategoryServiceServer) GetCategory(context.Context, *GetCategoryRequest) (*Category, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategory not implemented")
}
func (*UnimplementedCategoryServiceServer) ReportMiscategorization(context.Context, *GetCategoryRequest) (*Category, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportMiscategorization not implemented")
}
func (*UnimplementedCategoryServiceServer) ListUrls(*ListUrlsRequest, CategoryService_ListUrlsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListUrls not implemented")
}
func (*UnimplementedCategoryServiceServer) UpdateCategory(context.Context, *Category) (*Category, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCategory not implemented")
}
func (*UnimplementedCategoryServiceServer) AddUrl(context.Context, *Category) (*Category, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUrl not implemented")
}
func (*UnimplementedCategoryServiceServer) DeleteUrl(context.Context, *GetCategoryRequest) (*DeleteUrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUrl not implemented")
}

func RegisterCategoryServiceServer(s *grpc.Server, srv CategoryServiceServer) {
	s.RegisterService(&_CategoryService_serviceDesc, srv)
}

func _CategoryService_GetCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).GetCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CategoryService/GetCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).GetCategory(ctx, req.(*GetCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_ReportMiscategorization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).ReportMiscategorization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CategoryService/ReportMiscategorization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).ReportMiscategorization(ctx, req.(*GetCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_ListUrls_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListUrlsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CategoryServiceServer).ListUrls(m, &categoryServiceListUrlsServer{stream})
}

type CategoryService_ListUrlsServer interface {
	Send(*Category) error
	grpc.ServerStream
}

type categoryServiceListUrlsServer struct {
	grpc.ServerStream
}

func (x *categoryServiceListUrlsServer) Send(m *Category) error {
	return x.ServerStream.SendMsg(m)
}

func _CategoryService_UpdateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Category)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).UpdateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CategoryService/UpdateCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).UpdateCategory(ctx, req.(*Category))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_AddUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Category)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).AddUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CategoryService/AddUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).AddUrl(ctx, req.(*Category))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_DeleteUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).DeleteUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CategoryService/DeleteUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).DeleteUrl(ctx, req.(*GetCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CategoryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.CategoryService",
	HandlerType: (*CategoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCategory",
			Handler:    _CategoryService_GetCategory_Handler,
		},
		{
			MethodName: "ReportMiscategorization",
			Handler:    _CategoryService_ReportMiscategorization_Handler,
		},
		{
			MethodName: "UpdateCategory",
			Handler:    _CategoryService_UpdateCategory_Handler,
		},
		{
			MethodName: "AddUrl",
			Handler:    _CategoryService_AddUrl_Handler,
		},
		{
			MethodName: "DeleteUrl",
			Handler:    _CategoryService_DeleteUrl_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListUrls",
			Handler:       _CategoryService_ListUrls_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cat.proto",
}
