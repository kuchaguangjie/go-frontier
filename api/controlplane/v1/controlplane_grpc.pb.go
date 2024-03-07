// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.9
// source: controlplane.proto

package v1

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
	ControlPlane_ListEdges_FullMethodName         = "/controlplane.ControlPlane/ListEdges"
	ControlPlane_GetEdge_FullMethodName           = "/controlplane.ControlPlane/GetEdge"
	ControlPlane_KickEdge_FullMethodName          = "/controlplane.ControlPlane/KickEdge"
	ControlPlane_ListEdgeRPCs_FullMethodName      = "/controlplane.ControlPlane/ListEdgeRPCs"
	ControlPlane_ListServices_FullMethodName      = "/controlplane.ControlPlane/ListServices"
	ControlPlane_GetService_FullMethodName        = "/controlplane.ControlPlane/GetService"
	ControlPlane_KickService_FullMethodName       = "/controlplane.ControlPlane/KickService"
	ControlPlane_ListServiceRPCs_FullMethodName   = "/controlplane.ControlPlane/ListServiceRPCs"
	ControlPlane_ListServiceTopics_FullMethodName = "/controlplane.ControlPlane/ListServiceTopics"
)

// ControlPlaneClient is the client API for ControlPlane service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ControlPlaneClient interface {
	// edge related
	ListEdges(ctx context.Context, in *ListEdgesRequest, opts ...grpc.CallOption) (*ListEdgesResponse, error)
	GetEdge(ctx context.Context, in *GetEdgeRequest, opts ...grpc.CallOption) (*Edge, error)
	KickEdge(ctx context.Context, in *KickEdgeRequest, opts ...grpc.CallOption) (*KickEdgeResponse, error)
	ListEdgeRPCs(ctx context.Context, in *ListEdgeRPCsRequest, opts ...grpc.CallOption) (*ListEdgeRPCsResponse, error)
	// service related
	ListServices(ctx context.Context, in *ListServicesRequest, opts ...grpc.CallOption) (*ListServicesResponse, error)
	GetService(ctx context.Context, in *GetServiceRequest, opts ...grpc.CallOption) (*Service, error)
	KickService(ctx context.Context, in *KickServiceRequest, opts ...grpc.CallOption) (*KickServiceResponse, error)
	ListServiceRPCs(ctx context.Context, in *ListServiceRPCsRequest, opts ...grpc.CallOption) (*ListServiceRPCsResponse, error)
	ListServiceTopics(ctx context.Context, in *ListServiceTopicsRequest, opts ...grpc.CallOption) (*ListServiceTopicsResponse, error)
}

type controlPlaneClient struct {
	cc grpc.ClientConnInterface
}

func NewControlPlaneClient(cc grpc.ClientConnInterface) ControlPlaneClient {
	return &controlPlaneClient{cc}
}

func (c *controlPlaneClient) ListEdges(ctx context.Context, in *ListEdgesRequest, opts ...grpc.CallOption) (*ListEdgesResponse, error) {
	out := new(ListEdgesResponse)
	err := c.cc.Invoke(ctx, ControlPlane_ListEdges_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlPlaneClient) GetEdge(ctx context.Context, in *GetEdgeRequest, opts ...grpc.CallOption) (*Edge, error) {
	out := new(Edge)
	err := c.cc.Invoke(ctx, ControlPlane_GetEdge_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlPlaneClient) KickEdge(ctx context.Context, in *KickEdgeRequest, opts ...grpc.CallOption) (*KickEdgeResponse, error) {
	out := new(KickEdgeResponse)
	err := c.cc.Invoke(ctx, ControlPlane_KickEdge_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlPlaneClient) ListEdgeRPCs(ctx context.Context, in *ListEdgeRPCsRequest, opts ...grpc.CallOption) (*ListEdgeRPCsResponse, error) {
	out := new(ListEdgeRPCsResponse)
	err := c.cc.Invoke(ctx, ControlPlane_ListEdgeRPCs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlPlaneClient) ListServices(ctx context.Context, in *ListServicesRequest, opts ...grpc.CallOption) (*ListServicesResponse, error) {
	out := new(ListServicesResponse)
	err := c.cc.Invoke(ctx, ControlPlane_ListServices_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlPlaneClient) GetService(ctx context.Context, in *GetServiceRequest, opts ...grpc.CallOption) (*Service, error) {
	out := new(Service)
	err := c.cc.Invoke(ctx, ControlPlane_GetService_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlPlaneClient) KickService(ctx context.Context, in *KickServiceRequest, opts ...grpc.CallOption) (*KickServiceResponse, error) {
	out := new(KickServiceResponse)
	err := c.cc.Invoke(ctx, ControlPlane_KickService_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlPlaneClient) ListServiceRPCs(ctx context.Context, in *ListServiceRPCsRequest, opts ...grpc.CallOption) (*ListServiceRPCsResponse, error) {
	out := new(ListServiceRPCsResponse)
	err := c.cc.Invoke(ctx, ControlPlane_ListServiceRPCs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlPlaneClient) ListServiceTopics(ctx context.Context, in *ListServiceTopicsRequest, opts ...grpc.CallOption) (*ListServiceTopicsResponse, error) {
	out := new(ListServiceTopicsResponse)
	err := c.cc.Invoke(ctx, ControlPlane_ListServiceTopics_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ControlPlaneServer is the server API for ControlPlane service.
// All implementations must embed UnimplementedControlPlaneServer
// for forward compatibility
type ControlPlaneServer interface {
	// edge related
	ListEdges(context.Context, *ListEdgesRequest) (*ListEdgesResponse, error)
	GetEdge(context.Context, *GetEdgeRequest) (*Edge, error)
	KickEdge(context.Context, *KickEdgeRequest) (*KickEdgeResponse, error)
	ListEdgeRPCs(context.Context, *ListEdgeRPCsRequest) (*ListEdgeRPCsResponse, error)
	// service related
	ListServices(context.Context, *ListServicesRequest) (*ListServicesResponse, error)
	GetService(context.Context, *GetServiceRequest) (*Service, error)
	KickService(context.Context, *KickServiceRequest) (*KickServiceResponse, error)
	ListServiceRPCs(context.Context, *ListServiceRPCsRequest) (*ListServiceRPCsResponse, error)
	ListServiceTopics(context.Context, *ListServiceTopicsRequest) (*ListServiceTopicsResponse, error)
	mustEmbedUnimplementedControlPlaneServer()
}

// UnimplementedControlPlaneServer must be embedded to have forward compatible implementations.
type UnimplementedControlPlaneServer struct {
}

func (UnimplementedControlPlaneServer) ListEdges(context.Context, *ListEdgesRequest) (*ListEdgesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEdges not implemented")
}
func (UnimplementedControlPlaneServer) GetEdge(context.Context, *GetEdgeRequest) (*Edge, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEdge not implemented")
}
func (UnimplementedControlPlaneServer) KickEdge(context.Context, *KickEdgeRequest) (*KickEdgeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KickEdge not implemented")
}
func (UnimplementedControlPlaneServer) ListEdgeRPCs(context.Context, *ListEdgeRPCsRequest) (*ListEdgeRPCsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEdgeRPCs not implemented")
}
func (UnimplementedControlPlaneServer) ListServices(context.Context, *ListServicesRequest) (*ListServicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListServices not implemented")
}
func (UnimplementedControlPlaneServer) GetService(context.Context, *GetServiceRequest) (*Service, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetService not implemented")
}
func (UnimplementedControlPlaneServer) KickService(context.Context, *KickServiceRequest) (*KickServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KickService not implemented")
}
func (UnimplementedControlPlaneServer) ListServiceRPCs(context.Context, *ListServiceRPCsRequest) (*ListServiceRPCsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListServiceRPCs not implemented")
}
func (UnimplementedControlPlaneServer) ListServiceTopics(context.Context, *ListServiceTopicsRequest) (*ListServiceTopicsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListServiceTopics not implemented")
}
func (UnimplementedControlPlaneServer) mustEmbedUnimplementedControlPlaneServer() {}

// UnsafeControlPlaneServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ControlPlaneServer will
// result in compilation errors.
type UnsafeControlPlaneServer interface {
	mustEmbedUnimplementedControlPlaneServer()
}

func RegisterControlPlaneServer(s grpc.ServiceRegistrar, srv ControlPlaneServer) {
	s.RegisterService(&ControlPlane_ServiceDesc, srv)
}

func _ControlPlane_ListEdges_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEdgesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlPlaneServer).ListEdges(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ControlPlane_ListEdges_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlPlaneServer).ListEdges(ctx, req.(*ListEdgesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControlPlane_GetEdge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEdgeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlPlaneServer).GetEdge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ControlPlane_GetEdge_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlPlaneServer).GetEdge(ctx, req.(*GetEdgeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControlPlane_KickEdge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KickEdgeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlPlaneServer).KickEdge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ControlPlane_KickEdge_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlPlaneServer).KickEdge(ctx, req.(*KickEdgeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControlPlane_ListEdgeRPCs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEdgeRPCsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlPlaneServer).ListEdgeRPCs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ControlPlane_ListEdgeRPCs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlPlaneServer).ListEdgeRPCs(ctx, req.(*ListEdgeRPCsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControlPlane_ListServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListServicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlPlaneServer).ListServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ControlPlane_ListServices_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlPlaneServer).ListServices(ctx, req.(*ListServicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControlPlane_GetService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlPlaneServer).GetService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ControlPlane_GetService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlPlaneServer).GetService(ctx, req.(*GetServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControlPlane_KickService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KickServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlPlaneServer).KickService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ControlPlane_KickService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlPlaneServer).KickService(ctx, req.(*KickServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControlPlane_ListServiceRPCs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListServiceRPCsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlPlaneServer).ListServiceRPCs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ControlPlane_ListServiceRPCs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlPlaneServer).ListServiceRPCs(ctx, req.(*ListServiceRPCsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControlPlane_ListServiceTopics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListServiceTopicsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlPlaneServer).ListServiceTopics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ControlPlane_ListServiceTopics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlPlaneServer).ListServiceTopics(ctx, req.(*ListServiceTopicsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ControlPlane_ServiceDesc is the grpc.ServiceDesc for ControlPlane service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ControlPlane_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "controlplane.ControlPlane",
	HandlerType: (*ControlPlaneServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListEdges",
			Handler:    _ControlPlane_ListEdges_Handler,
		},
		{
			MethodName: "GetEdge",
			Handler:    _ControlPlane_GetEdge_Handler,
		},
		{
			MethodName: "KickEdge",
			Handler:    _ControlPlane_KickEdge_Handler,
		},
		{
			MethodName: "ListEdgeRPCs",
			Handler:    _ControlPlane_ListEdgeRPCs_Handler,
		},
		{
			MethodName: "ListServices",
			Handler:    _ControlPlane_ListServices_Handler,
		},
		{
			MethodName: "GetService",
			Handler:    _ControlPlane_GetService_Handler,
		},
		{
			MethodName: "KickService",
			Handler:    _ControlPlane_KickService_Handler,
		},
		{
			MethodName: "ListServiceRPCs",
			Handler:    _ControlPlane_ListServiceRPCs_Handler,
		},
		{
			MethodName: "ListServiceTopics",
			Handler:    _ControlPlane_ListServiceTopics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "controlplane.proto",
}
