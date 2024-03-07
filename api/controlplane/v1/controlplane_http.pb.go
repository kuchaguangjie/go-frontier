// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.2
// - protoc             v3.21.9
// source: controlplane.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationControlPlaneGetEdge = "/controlplane.ControlPlane/GetEdge"
const OperationControlPlaneGetService = "/controlplane.ControlPlane/GetService"
const OperationControlPlaneKickEdge = "/controlplane.ControlPlane/KickEdge"
const OperationControlPlaneKickService = "/controlplane.ControlPlane/KickService"
const OperationControlPlaneListEdgeRPCs = "/controlplane.ControlPlane/ListEdgeRPCs"
const OperationControlPlaneListEdges = "/controlplane.ControlPlane/ListEdges"
const OperationControlPlaneListServiceRPCs = "/controlplane.ControlPlane/ListServiceRPCs"
const OperationControlPlaneListServiceTopics = "/controlplane.ControlPlane/ListServiceTopics"
const OperationControlPlaneListServices = "/controlplane.ControlPlane/ListServices"

type ControlPlaneHTTPServer interface {
	GetEdge(context.Context, *GetEdgeRequest) (*Edge, error)
	GetService(context.Context, *GetServiceRequest) (*Service, error)
	KickEdge(context.Context, *KickEdgeRequest) (*KickEdgeResponse, error)
	KickService(context.Context, *KickServiceRequest) (*KickServiceResponse, error)
	ListEdgeRPCs(context.Context, *ListEdgeRPCsRequest) (*ListEdgeRPCsResponse, error)
	// ListEdges edge related
	ListEdges(context.Context, *ListEdgesRequest) (*ListEdgesResponse, error)
	ListServiceRPCs(context.Context, *ListServiceRPCsRequest) (*ListServiceRPCsResponse, error)
	ListServiceTopics(context.Context, *ListServiceTopicsRequest) (*ListServiceTopicsResponse, error)
	// ListServices service related
	ListServices(context.Context, *ListServicesRequest) (*ListServicesResponse, error)
}

func RegisterControlPlaneHTTPServer(s *http.Server, srv ControlPlaneHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/edges", _ControlPlane_ListEdges0_HTTP_Handler(srv))
	r.GET("/v1/edges/{edge_id}", _ControlPlane_GetEdge0_HTTP_Handler(srv))
	r.DELETE("/v1/edges/{edge_id}", _ControlPlane_KickEdge0_HTTP_Handler(srv))
	r.GET("/v1/edges/rpcs", _ControlPlane_ListEdgeRPCs0_HTTP_Handler(srv))
	r.GET("/v1/services", _ControlPlane_ListServices0_HTTP_Handler(srv))
	r.GET("/v1/services/{service_id}", _ControlPlane_GetService0_HTTP_Handler(srv))
	r.DELETE("/v1/services/{service_id}", _ControlPlane_KickService0_HTTP_Handler(srv))
	r.GET("/v1/services/rpcs", _ControlPlane_ListServiceRPCs0_HTTP_Handler(srv))
	r.GET("/v1/services/topics", _ControlPlane_ListServiceTopics0_HTTP_Handler(srv))
}

func _ControlPlane_ListEdges0_HTTP_Handler(srv ControlPlaneHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListEdgesRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationControlPlaneListEdges)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListEdges(ctx, req.(*ListEdgesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListEdgesResponse)
		return ctx.Result(200, reply)
	}
}

func _ControlPlane_GetEdge0_HTTP_Handler(srv ControlPlaneHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetEdgeRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationControlPlaneGetEdge)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetEdge(ctx, req.(*GetEdgeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Edge)
		return ctx.Result(200, reply)
	}
}

func _ControlPlane_KickEdge0_HTTP_Handler(srv ControlPlaneHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in KickEdgeRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationControlPlaneKickEdge)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.KickEdge(ctx, req.(*KickEdgeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*KickEdgeResponse)
		return ctx.Result(200, reply)
	}
}

func _ControlPlane_ListEdgeRPCs0_HTTP_Handler(srv ControlPlaneHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListEdgeRPCsRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationControlPlaneListEdgeRPCs)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListEdgeRPCs(ctx, req.(*ListEdgeRPCsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListEdgeRPCsResponse)
		return ctx.Result(200, reply)
	}
}

func _ControlPlane_ListServices0_HTTP_Handler(srv ControlPlaneHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListServicesRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationControlPlaneListServices)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListServices(ctx, req.(*ListServicesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListServicesResponse)
		return ctx.Result(200, reply)
	}
}

func _ControlPlane_GetService0_HTTP_Handler(srv ControlPlaneHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetServiceRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationControlPlaneGetService)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetService(ctx, req.(*GetServiceRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Service)
		return ctx.Result(200, reply)
	}
}

func _ControlPlane_KickService0_HTTP_Handler(srv ControlPlaneHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in KickServiceRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationControlPlaneKickService)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.KickService(ctx, req.(*KickServiceRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*KickServiceResponse)
		return ctx.Result(200, reply)
	}
}

func _ControlPlane_ListServiceRPCs0_HTTP_Handler(srv ControlPlaneHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListServiceRPCsRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationControlPlaneListServiceRPCs)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListServiceRPCs(ctx, req.(*ListServiceRPCsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListServiceRPCsResponse)
		return ctx.Result(200, reply)
	}
}

func _ControlPlane_ListServiceTopics0_HTTP_Handler(srv ControlPlaneHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListServiceTopicsRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationControlPlaneListServiceTopics)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListServiceTopics(ctx, req.(*ListServiceTopicsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListServiceTopicsResponse)
		return ctx.Result(200, reply)
	}
}

type ControlPlaneHTTPClient interface {
	GetEdge(ctx context.Context, req *GetEdgeRequest, opts ...http.CallOption) (rsp *Edge, err error)
	GetService(ctx context.Context, req *GetServiceRequest, opts ...http.CallOption) (rsp *Service, err error)
	KickEdge(ctx context.Context, req *KickEdgeRequest, opts ...http.CallOption) (rsp *KickEdgeResponse, err error)
	KickService(ctx context.Context, req *KickServiceRequest, opts ...http.CallOption) (rsp *KickServiceResponse, err error)
	ListEdgeRPCs(ctx context.Context, req *ListEdgeRPCsRequest, opts ...http.CallOption) (rsp *ListEdgeRPCsResponse, err error)
	ListEdges(ctx context.Context, req *ListEdgesRequest, opts ...http.CallOption) (rsp *ListEdgesResponse, err error)
	ListServiceRPCs(ctx context.Context, req *ListServiceRPCsRequest, opts ...http.CallOption) (rsp *ListServiceRPCsResponse, err error)
	ListServiceTopics(ctx context.Context, req *ListServiceTopicsRequest, opts ...http.CallOption) (rsp *ListServiceTopicsResponse, err error)
	ListServices(ctx context.Context, req *ListServicesRequest, opts ...http.CallOption) (rsp *ListServicesResponse, err error)
}

type ControlPlaneHTTPClientImpl struct {
	cc *http.Client
}

func NewControlPlaneHTTPClient(client *http.Client) ControlPlaneHTTPClient {
	return &ControlPlaneHTTPClientImpl{client}
}

func (c *ControlPlaneHTTPClientImpl) GetEdge(ctx context.Context, in *GetEdgeRequest, opts ...http.CallOption) (*Edge, error) {
	var out Edge
	pattern := "/v1/edges/{edge_id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationControlPlaneGetEdge))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *ControlPlaneHTTPClientImpl) GetService(ctx context.Context, in *GetServiceRequest, opts ...http.CallOption) (*Service, error) {
	var out Service
	pattern := "/v1/services/{service_id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationControlPlaneGetService))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *ControlPlaneHTTPClientImpl) KickEdge(ctx context.Context, in *KickEdgeRequest, opts ...http.CallOption) (*KickEdgeResponse, error) {
	var out KickEdgeResponse
	pattern := "/v1/edges/{edge_id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationControlPlaneKickEdge))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *ControlPlaneHTTPClientImpl) KickService(ctx context.Context, in *KickServiceRequest, opts ...http.CallOption) (*KickServiceResponse, error) {
	var out KickServiceResponse
	pattern := "/v1/services/{service_id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationControlPlaneKickService))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *ControlPlaneHTTPClientImpl) ListEdgeRPCs(ctx context.Context, in *ListEdgeRPCsRequest, opts ...http.CallOption) (*ListEdgeRPCsResponse, error) {
	var out ListEdgeRPCsResponse
	pattern := "/v1/edges/rpcs"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationControlPlaneListEdgeRPCs))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *ControlPlaneHTTPClientImpl) ListEdges(ctx context.Context, in *ListEdgesRequest, opts ...http.CallOption) (*ListEdgesResponse, error) {
	var out ListEdgesResponse
	pattern := "/v1/edges"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationControlPlaneListEdges))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *ControlPlaneHTTPClientImpl) ListServiceRPCs(ctx context.Context, in *ListServiceRPCsRequest, opts ...http.CallOption) (*ListServiceRPCsResponse, error) {
	var out ListServiceRPCsResponse
	pattern := "/v1/services/rpcs"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationControlPlaneListServiceRPCs))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *ControlPlaneHTTPClientImpl) ListServiceTopics(ctx context.Context, in *ListServiceTopicsRequest, opts ...http.CallOption) (*ListServiceTopicsResponse, error) {
	var out ListServiceTopicsResponse
	pattern := "/v1/services/topics"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationControlPlaneListServiceTopics))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *ControlPlaneHTTPClientImpl) ListServices(ctx context.Context, in *ListServicesRequest, opts ...http.CallOption) (*ListServicesResponse, error) {
	var out ListServicesResponse
	pattern := "/v1/services"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationControlPlaneListServices))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
