// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: agent/v1/agent.proto

package agentv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/grafana/phlare/pkg/gen/agent/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// AgentServiceName is the fully-qualified name of the AgentService service.
	AgentServiceName = "agent.v1.AgentService"
)

// AgentServiceClient is a client for the agent.v1.AgentService service.
type AgentServiceClient interface {
	// Retrieve information about targets.
	GetTargets(context.Context, *connect_go.Request[v1.GetTargetsRequest]) (*connect_go.Response[v1.GetTargetsResponse], error)
}

// NewAgentServiceClient constructs a client for the agent.v1.AgentService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAgentServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) AgentServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &agentServiceClient{
		getTargets: connect_go.NewClient[v1.GetTargetsRequest, v1.GetTargetsResponse](
			httpClient,
			baseURL+"/agent.v1.AgentService/GetTargets",
			opts...,
		),
	}
}

// agentServiceClient implements AgentServiceClient.
type agentServiceClient struct {
	getTargets *connect_go.Client[v1.GetTargetsRequest, v1.GetTargetsResponse]
}

// GetTargets calls agent.v1.AgentService.GetTargets.
func (c *agentServiceClient) GetTargets(ctx context.Context, req *connect_go.Request[v1.GetTargetsRequest]) (*connect_go.Response[v1.GetTargetsResponse], error) {
	return c.getTargets.CallUnary(ctx, req)
}

// AgentServiceHandler is an implementation of the agent.v1.AgentService service.
type AgentServiceHandler interface {
	// Retrieve information about targets.
	GetTargets(context.Context, *connect_go.Request[v1.GetTargetsRequest]) (*connect_go.Response[v1.GetTargetsResponse], error)
}

// NewAgentServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewAgentServiceHandler(svc AgentServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/agent.v1.AgentService/GetTargets", connect_go.NewUnaryHandler(
		"/agent.v1.AgentService/GetTargets",
		svc.GetTargets,
		opts...,
	))
	return "/agent.v1.AgentService/", mux
}

// UnimplementedAgentServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedAgentServiceHandler struct{}

func (UnimplementedAgentServiceHandler) GetTargets(context.Context, *connect_go.Request[v1.GetTargetsRequest]) (*connect_go.Response[v1.GetTargetsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("agent.v1.AgentService.GetTargets is not implemented"))
}
