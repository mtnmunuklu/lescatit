// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.17.3
// source: crawl.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	CrawlService_GetURLData_FullMethodName  = "/messages.CrawlService/GetURLData"
	CrawlService_GetURLsData_FullMethodName = "/messages.CrawlService/GetURLsData"
	CrawlService_CrawlURL_FullMethodName    = "/messages.CrawlService/CrawlURL"
	CrawlService_CrawlURLs_FullMethodName   = "/messages.CrawlService/CrawlURLs"
)

// CrawlServiceClient is the client API for CrawlService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CrawlServiceClient interface {
	GetURLData(ctx context.Context, in *GetURLDataRequest, opts ...grpc.CallOption) (*GetURLDataResponse, error)
	GetURLsData(ctx context.Context, in *GetURLsDataRequest, opts ...grpc.CallOption) (CrawlService_GetURLsDataClient, error)
	CrawlURL(ctx context.Context, in *CrawlURLRequest, opts ...grpc.CallOption) (*CrawlURLResponse, error)
	CrawlURLs(ctx context.Context, in *CrawlURLsRequest, opts ...grpc.CallOption) (CrawlService_CrawlURLsClient, error)
}

type crawlServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCrawlServiceClient(cc grpc.ClientConnInterface) CrawlServiceClient {
	return &crawlServiceClient{cc}
}

func (c *crawlServiceClient) GetURLData(ctx context.Context, in *GetURLDataRequest, opts ...grpc.CallOption) (*GetURLDataResponse, error) {
	out := new(GetURLDataResponse)
	err := c.cc.Invoke(ctx, CrawlService_GetURLData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crawlServiceClient) GetURLsData(ctx context.Context, in *GetURLsDataRequest, opts ...grpc.CallOption) (CrawlService_GetURLsDataClient, error) {
	stream, err := c.cc.NewStream(ctx, &CrawlService_ServiceDesc.Streams[0], CrawlService_GetURLsData_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &crawlServiceGetURLsDataClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CrawlService_GetURLsDataClient interface {
	Recv() (*GetURLDataResponse, error)
	grpc.ClientStream
}

type crawlServiceGetURLsDataClient struct {
	grpc.ClientStream
}

func (x *crawlServiceGetURLsDataClient) Recv() (*GetURLDataResponse, error) {
	m := new(GetURLDataResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *crawlServiceClient) CrawlURL(ctx context.Context, in *CrawlURLRequest, opts ...grpc.CallOption) (*CrawlURLResponse, error) {
	out := new(CrawlURLResponse)
	err := c.cc.Invoke(ctx, CrawlService_CrawlURL_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crawlServiceClient) CrawlURLs(ctx context.Context, in *CrawlURLsRequest, opts ...grpc.CallOption) (CrawlService_CrawlURLsClient, error) {
	stream, err := c.cc.NewStream(ctx, &CrawlService_ServiceDesc.Streams[1], CrawlService_CrawlURLs_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &crawlServiceCrawlURLsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CrawlService_CrawlURLsClient interface {
	Recv() (*CrawlURLResponse, error)
	grpc.ClientStream
}

type crawlServiceCrawlURLsClient struct {
	grpc.ClientStream
}

func (x *crawlServiceCrawlURLsClient) Recv() (*CrawlURLResponse, error) {
	m := new(CrawlURLResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CrawlServiceServer is the server API for CrawlService service.
// All implementations should embed UnimplementedCrawlServiceServer
// for forward compatibility
type CrawlServiceServer interface {
	GetURLData(context.Context, *GetURLDataRequest) (*GetURLDataResponse, error)
	GetURLsData(*GetURLsDataRequest, CrawlService_GetURLsDataServer) error
	CrawlURL(context.Context, *CrawlURLRequest) (*CrawlURLResponse, error)
	CrawlURLs(*CrawlURLsRequest, CrawlService_CrawlURLsServer) error
}

// UnimplementedCrawlServiceServer should be embedded to have forward compatible implementations.
type UnimplementedCrawlServiceServer struct {
}

func (UnimplementedCrawlServiceServer) GetURLData(context.Context, *GetURLDataRequest) (*GetURLDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetURLData not implemented")
}
func (UnimplementedCrawlServiceServer) GetURLsData(*GetURLsDataRequest, CrawlService_GetURLsDataServer) error {
	return status.Errorf(codes.Unimplemented, "method GetURLsData not implemented")
}
func (UnimplementedCrawlServiceServer) CrawlURL(context.Context, *CrawlURLRequest) (*CrawlURLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CrawlURL not implemented")
}
func (UnimplementedCrawlServiceServer) CrawlURLs(*CrawlURLsRequest, CrawlService_CrawlURLsServer) error {
	return status.Errorf(codes.Unimplemented, "method CrawlURLs not implemented")
}

// UnsafeCrawlServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CrawlServiceServer will
// result in compilation errors.
type UnsafeCrawlServiceServer interface {
	mustEmbedUnimplementedCrawlServiceServer()
}

func RegisterCrawlServiceServer(s grpc.ServiceRegistrar, srv CrawlServiceServer) {
	s.RegisterService(&CrawlService_ServiceDesc, srv)
}

func _CrawlService_GetURLData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetURLDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrawlServiceServer).GetURLData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CrawlService_GetURLData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrawlServiceServer).GetURLData(ctx, req.(*GetURLDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CrawlService_GetURLsData_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetURLsDataRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CrawlServiceServer).GetURLsData(m, &crawlServiceGetURLsDataServer{stream})
}

type CrawlService_GetURLsDataServer interface {
	Send(*GetURLDataResponse) error
	grpc.ServerStream
}

type crawlServiceGetURLsDataServer struct {
	grpc.ServerStream
}

func (x *crawlServiceGetURLsDataServer) Send(m *GetURLDataResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CrawlService_CrawlURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CrawlURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrawlServiceServer).CrawlURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CrawlService_CrawlURL_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrawlServiceServer).CrawlURL(ctx, req.(*CrawlURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CrawlService_CrawlURLs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CrawlURLsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CrawlServiceServer).CrawlURLs(m, &crawlServiceCrawlURLsServer{stream})
}

type CrawlService_CrawlURLsServer interface {
	Send(*CrawlURLResponse) error
	grpc.ServerStream
}

type crawlServiceCrawlURLsServer struct {
	grpc.ServerStream
}

func (x *crawlServiceCrawlURLsServer) Send(m *CrawlURLResponse) error {
	return x.ServerStream.SendMsg(m)
}

// CrawlService_ServiceDesc is the grpc.ServiceDesc for CrawlService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CrawlService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "messages.CrawlService",
	HandlerType: (*CrawlServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetURLData",
			Handler:    _CrawlService_GetURLData_Handler,
		},
		{
			MethodName: "CrawlURL",
			Handler:    _CrawlService_CrawlURL_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetURLsData",
			Handler:       _CrawlService_GetURLsData_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "CrawlURLs",
			Handler:       _CrawlService_CrawlURLs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "crawl.proto",
}
