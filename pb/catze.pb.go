// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: catze.proto

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

type Categorizer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Url      string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Category string `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
	Created  int64  `protobuf:"varint,4,opt,name=created,proto3" json:"created,omitempty"`
	Updated  int64  `protobuf:"varint,5,opt,name=updated,proto3" json:"updated,omitempty"`
	Revision string `protobuf:"bytes,6,opt,name=revision,proto3" json:"revision,omitempty"`
	Data     string `protobuf:"bytes,7,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Categorizer) Reset() {
	*x = Categorizer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catze_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Categorizer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Categorizer) ProtoMessage() {}

func (x *Categorizer) ProtoReflect() protoreflect.Message {
	mi := &file_catze_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Categorizer.ProtoReflect.Descriptor instead.
func (*Categorizer) Descriptor() ([]byte, []int) {
	return file_catze_proto_rawDescGZIP(), []int{0}
}

func (x *Categorizer) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Categorizer) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Categorizer) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *Categorizer) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *Categorizer) GetUpdated() int64 {
	if x != nil {
		return x.Updated
	}
	return 0
}

func (x *Categorizer) GetRevision() string {
	if x != nil {
		return x.Revision
	}
	return ""
}

func (x *Categorizer) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type CategorizeURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url  string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Data string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CategorizeURLRequest) Reset() {
	*x = CategorizeURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catze_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategorizeURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategorizeURLRequest) ProtoMessage() {}

func (x *CategorizeURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_catze_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategorizeURLRequest.ProtoReflect.Descriptor instead.
func (*CategorizeURLRequest) Descriptor() ([]byte, []int) {
	return file_catze_proto_rawDescGZIP(), []int{1}
}

func (x *CategorizeURLRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *CategorizeURLRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type CategorizeURLsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategorizeURLRequest []*CategorizeURLRequest `protobuf:"bytes,1,rep,name=categorizeURLRequest,proto3" json:"categorizeURLRequest,omitempty"`
}

func (x *CategorizeURLsRequest) Reset() {
	*x = CategorizeURLsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catze_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategorizeURLsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategorizeURLsRequest) ProtoMessage() {}

func (x *CategorizeURLsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_catze_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategorizeURLsRequest.ProtoReflect.Descriptor instead.
func (*CategorizeURLsRequest) Descriptor() ([]byte, []int) {
	return file_catze_proto_rawDescGZIP(), []int{2}
}

func (x *CategorizeURLsRequest) GetCategorizeURLRequest() []*CategorizeURLRequest {
	if x != nil {
		return x.CategorizeURLRequest
	}
	return nil
}

type CategorizeURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url      string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Category string `protobuf:"bytes,2,opt,name=category,proto3" json:"category,omitempty"`
}

func (x *CategorizeURLResponse) Reset() {
	*x = CategorizeURLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catze_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategorizeURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategorizeURLResponse) ProtoMessage() {}

func (x *CategorizeURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_catze_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategorizeURLResponse.ProtoReflect.Descriptor instead.
func (*CategorizeURLResponse) Descriptor() ([]byte, []int) {
	return file_catze_proto_rawDescGZIP(), []int{3}
}

func (x *CategorizeURLResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *CategorizeURLResponse) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

var File_catze_proto protoreflect.FileDescriptor

var file_catze_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x63, 0x61, 0x74, 0x7a, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x22, 0xaf, 0x01, 0x0a, 0x0b, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x7a, 0x65,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x3c, 0x0a, 0x14, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x7a,
	0x65, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x65, 0x0a, 0x15, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x55,
	0x52, 0x4c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4c, 0x0a, 0x14, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x14, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x55, 0x52,
	0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x45, 0x0a, 0x15, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x32,
	0x9e, 0x01, 0x0a, 0x0c, 0x43, 0x61, 0x74, 0x7a, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x44, 0x0a, 0x0d, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x55, 0x52,
	0x4c, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x7a,
	0x65, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x55, 0x52, 0x4c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x69, 0x7a, 0x65, 0x55, 0x52, 0x4c, 0x73, 0x12, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x55, 0x52, 0x4c, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x69, 0x7a, 0x65, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01,
	0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_catze_proto_rawDescOnce sync.Once
	file_catze_proto_rawDescData = file_catze_proto_rawDesc
)

func file_catze_proto_rawDescGZIP() []byte {
	file_catze_proto_rawDescOnce.Do(func() {
		file_catze_proto_rawDescData = protoimpl.X.CompressGZIP(file_catze_proto_rawDescData)
	})
	return file_catze_proto_rawDescData
}

var file_catze_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_catze_proto_goTypes = []interface{}{
	(*Categorizer)(nil),           // 0: pb.Categorizer
	(*CategorizeURLRequest)(nil),  // 1: pb.CategorizeURLRequest
	(*CategorizeURLsRequest)(nil), // 2: pb.CategorizeURLsRequest
	(*CategorizeURLResponse)(nil), // 3: pb.CategorizeURLResponse
}
var file_catze_proto_depIdxs = []int32{
	1, // 0: pb.CategorizeURLsRequest.categorizeURLRequest:type_name -> pb.CategorizeURLRequest
	1, // 1: pb.CatzeService.CategorizeURL:input_type -> pb.CategorizeURLRequest
	2, // 2: pb.CatzeService.CategorizeURLs:input_type -> pb.CategorizeURLsRequest
	3, // 3: pb.CatzeService.CategorizeURL:output_type -> pb.CategorizeURLResponse
	3, // 4: pb.CatzeService.CategorizeURLs:output_type -> pb.CategorizeURLResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_catze_proto_init() }
func file_catze_proto_init() {
	if File_catze_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_catze_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Categorizer); i {
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
		file_catze_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategorizeURLRequest); i {
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
		file_catze_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategorizeURLsRequest); i {
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
		file_catze_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategorizeURLResponse); i {
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
			RawDescriptor: file_catze_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_catze_proto_goTypes,
		DependencyIndexes: file_catze_proto_depIdxs,
		MessageInfos:      file_catze_proto_msgTypes,
	}.Build()
	File_catze_proto = out.File
	file_catze_proto_rawDesc = nil
	file_catze_proto_goTypes = nil
	file_catze_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CatzeServiceClient is the client API for CatzeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CatzeServiceClient interface {
	CategorizeURL(ctx context.Context, in *CategorizeURLRequest, opts ...grpc.CallOption) (*CategorizeURLResponse, error)
	CategorizeURLs(ctx context.Context, in *CategorizeURLsRequest, opts ...grpc.CallOption) (CatzeService_CategorizeURLsClient, error)
}

type catzeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCatzeServiceClient(cc grpc.ClientConnInterface) CatzeServiceClient {
	return &catzeServiceClient{cc}
}

func (c *catzeServiceClient) CategorizeURL(ctx context.Context, in *CategorizeURLRequest, opts ...grpc.CallOption) (*CategorizeURLResponse, error) {
	out := new(CategorizeURLResponse)
	err := c.cc.Invoke(ctx, "/pb.CatzeService/CategorizeURL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catzeServiceClient) CategorizeURLs(ctx context.Context, in *CategorizeURLsRequest, opts ...grpc.CallOption) (CatzeService_CategorizeURLsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CatzeService_serviceDesc.Streams[0], "/pb.CatzeService/CategorizeURLs", opts...)
	if err != nil {
		return nil, err
	}
	x := &catzeServiceCategorizeURLsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CatzeService_CategorizeURLsClient interface {
	Recv() (*CategorizeURLResponse, error)
	grpc.ClientStream
}

type catzeServiceCategorizeURLsClient struct {
	grpc.ClientStream
}

func (x *catzeServiceCategorizeURLsClient) Recv() (*CategorizeURLResponse, error) {
	m := new(CategorizeURLResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CatzeServiceServer is the server API for CatzeService service.
type CatzeServiceServer interface {
	CategorizeURL(context.Context, *CategorizeURLRequest) (*CategorizeURLResponse, error)
	CategorizeURLs(*CategorizeURLsRequest, CatzeService_CategorizeURLsServer) error
}

// UnimplementedCatzeServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCatzeServiceServer struct {
}

func (*UnimplementedCatzeServiceServer) CategorizeURL(context.Context, *CategorizeURLRequest) (*CategorizeURLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CategorizeURL not implemented")
}
func (*UnimplementedCatzeServiceServer) CategorizeURLs(*CategorizeURLsRequest, CatzeService_CategorizeURLsServer) error {
	return status.Errorf(codes.Unimplemented, "method CategorizeURLs not implemented")
}

func RegisterCatzeServiceServer(s *grpc.Server, srv CatzeServiceServer) {
	s.RegisterService(&_CatzeService_serviceDesc, srv)
}

func _CatzeService_CategorizeURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategorizeURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatzeServiceServer).CategorizeURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CatzeService/CategorizeURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatzeServiceServer).CategorizeURL(ctx, req.(*CategorizeURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatzeService_CategorizeURLs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CategorizeURLsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CatzeServiceServer).CategorizeURLs(m, &catzeServiceCategorizeURLsServer{stream})
}

type CatzeService_CategorizeURLsServer interface {
	Send(*CategorizeURLResponse) error
	grpc.ServerStream
}

type catzeServiceCategorizeURLsServer struct {
	grpc.ServerStream
}

func (x *catzeServiceCategorizeURLsServer) Send(m *CategorizeURLResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _CatzeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.CatzeService",
	HandlerType: (*CatzeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CategorizeURL",
			Handler:    _CatzeService_CategorizeURL_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CategorizeURLs",
			Handler:       _CatzeService_CategorizeURLs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "catze.proto",
}