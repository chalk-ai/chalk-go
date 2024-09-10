// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: chalk/server/v1/deploy.proto

package serverv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	DeployService_DeployBranch_FullMethodName      = "/chalk.server.v1.DeployService/DeployBranch"
	DeployService_GetDeployment_FullMethodName     = "/chalk.server.v1.DeployService/GetDeployment"
	DeployService_ListDeployments_FullMethodName   = "/chalk.server.v1.DeployService/ListDeployments"
	DeployService_SuspendDeployment_FullMethodName = "/chalk.server.v1.DeployService/SuspendDeployment"
	DeployService_ScaleDeployment_FullMethodName   = "/chalk.server.v1.DeployService/ScaleDeployment"
	DeployService_TagDeployment_FullMethodName     = "/chalk.server.v1.DeployService/TagDeployment"
)

// DeployServiceClient is the client API for DeployService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeployServiceClient interface {
	DeployBranch(ctx context.Context, in *DeployBranchRequest, opts ...grpc.CallOption) (*DeployBranchResponse, error)
	GetDeployment(ctx context.Context, in *GetDeploymentRequest, opts ...grpc.CallOption) (*GetDeploymentResponse, error)
	ListDeployments(ctx context.Context, in *ListDeploymentsRequest, opts ...grpc.CallOption) (*ListDeploymentsResponse, error)
	SuspendDeployment(ctx context.Context, in *SuspendDeploymentRequest, opts ...grpc.CallOption) (*SuspendDeploymentResponse, error)
	ScaleDeployment(ctx context.Context, in *ScaleDeploymentRequest, opts ...grpc.CallOption) (*ScaleDeploymentResponse, error)
	TagDeployment(ctx context.Context, in *TagDeploymentRequest, opts ...grpc.CallOption) (*TagDeploymentResponse, error)
}

type deployServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDeployServiceClient(cc grpc.ClientConnInterface) DeployServiceClient {
	return &deployServiceClient{cc}
}

func (c *deployServiceClient) DeployBranch(ctx context.Context, in *DeployBranchRequest, opts ...grpc.CallOption) (*DeployBranchResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeployBranchResponse)
	err := c.cc.Invoke(ctx, DeployService_DeployBranch_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deployServiceClient) GetDeployment(ctx context.Context, in *GetDeploymentRequest, opts ...grpc.CallOption) (*GetDeploymentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDeploymentResponse)
	err := c.cc.Invoke(ctx, DeployService_GetDeployment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deployServiceClient) ListDeployments(ctx context.Context, in *ListDeploymentsRequest, opts ...grpc.CallOption) (*ListDeploymentsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListDeploymentsResponse)
	err := c.cc.Invoke(ctx, DeployService_ListDeployments_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deployServiceClient) SuspendDeployment(ctx context.Context, in *SuspendDeploymentRequest, opts ...grpc.CallOption) (*SuspendDeploymentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SuspendDeploymentResponse)
	err := c.cc.Invoke(ctx, DeployService_SuspendDeployment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deployServiceClient) ScaleDeployment(ctx context.Context, in *ScaleDeploymentRequest, opts ...grpc.CallOption) (*ScaleDeploymentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ScaleDeploymentResponse)
	err := c.cc.Invoke(ctx, DeployService_ScaleDeployment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deployServiceClient) TagDeployment(ctx context.Context, in *TagDeploymentRequest, opts ...grpc.CallOption) (*TagDeploymentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TagDeploymentResponse)
	err := c.cc.Invoke(ctx, DeployService_TagDeployment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeployServiceServer is the server API for DeployService service.
// All implementations must embed UnimplementedDeployServiceServer
// for forward compatibility.
type DeployServiceServer interface {
	DeployBranch(context.Context, *DeployBranchRequest) (*DeployBranchResponse, error)
	GetDeployment(context.Context, *GetDeploymentRequest) (*GetDeploymentResponse, error)
	ListDeployments(context.Context, *ListDeploymentsRequest) (*ListDeploymentsResponse, error)
	SuspendDeployment(context.Context, *SuspendDeploymentRequest) (*SuspendDeploymentResponse, error)
	ScaleDeployment(context.Context, *ScaleDeploymentRequest) (*ScaleDeploymentResponse, error)
	TagDeployment(context.Context, *TagDeploymentRequest) (*TagDeploymentResponse, error)
	mustEmbedUnimplementedDeployServiceServer()
}

// UnimplementedDeployServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDeployServiceServer struct{}

func (UnimplementedDeployServiceServer) DeployBranch(context.Context, *DeployBranchRequest) (*DeployBranchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeployBranch not implemented")
}
func (UnimplementedDeployServiceServer) GetDeployment(context.Context, *GetDeploymentRequest) (*GetDeploymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeployment not implemented")
}
func (UnimplementedDeployServiceServer) ListDeployments(context.Context, *ListDeploymentsRequest) (*ListDeploymentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDeployments not implemented")
}
func (UnimplementedDeployServiceServer) SuspendDeployment(context.Context, *SuspendDeploymentRequest) (*SuspendDeploymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SuspendDeployment not implemented")
}
func (UnimplementedDeployServiceServer) ScaleDeployment(context.Context, *ScaleDeploymentRequest) (*ScaleDeploymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ScaleDeployment not implemented")
}
func (UnimplementedDeployServiceServer) TagDeployment(context.Context, *TagDeploymentRequest) (*TagDeploymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TagDeployment not implemented")
}
func (UnimplementedDeployServiceServer) mustEmbedUnimplementedDeployServiceServer() {}
func (UnimplementedDeployServiceServer) testEmbeddedByValue()                       {}

// UnsafeDeployServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeployServiceServer will
// result in compilation errors.
type UnsafeDeployServiceServer interface {
	mustEmbedUnimplementedDeployServiceServer()
}

func RegisterDeployServiceServer(s grpc.ServiceRegistrar, srv DeployServiceServer) {
	// If the following call pancis, it indicates UnimplementedDeployServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DeployService_ServiceDesc, srv)
}

func _DeployService_DeployBranch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeployBranchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeployServiceServer).DeployBranch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeployService_DeployBranch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeployServiceServer).DeployBranch(ctx, req.(*DeployBranchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeployService_GetDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeploymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeployServiceServer).GetDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeployService_GetDeployment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeployServiceServer).GetDeployment(ctx, req.(*GetDeploymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeployService_ListDeployments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDeploymentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeployServiceServer).ListDeployments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeployService_ListDeployments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeployServiceServer).ListDeployments(ctx, req.(*ListDeploymentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeployService_SuspendDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SuspendDeploymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeployServiceServer).SuspendDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeployService_SuspendDeployment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeployServiceServer).SuspendDeployment(ctx, req.(*SuspendDeploymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeployService_ScaleDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScaleDeploymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeployServiceServer).ScaleDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeployService_ScaleDeployment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeployServiceServer).ScaleDeployment(ctx, req.(*ScaleDeploymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeployService_TagDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TagDeploymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeployServiceServer).TagDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeployService_TagDeployment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeployServiceServer).TagDeployment(ctx, req.(*TagDeploymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DeployService_ServiceDesc is the grpc.ServiceDesc for DeployService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeployService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chalk.server.v1.DeployService",
	HandlerType: (*DeployServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeployBranch",
			Handler:    _DeployService_DeployBranch_Handler,
		},
		{
			MethodName: "GetDeployment",
			Handler:    _DeployService_GetDeployment_Handler,
		},
		{
			MethodName: "ListDeployments",
			Handler:    _DeployService_ListDeployments_Handler,
		},
		{
			MethodName: "SuspendDeployment",
			Handler:    _DeployService_SuspendDeployment_Handler,
		},
		{
			MethodName: "ScaleDeployment",
			Handler:    _DeployService_ScaleDeployment_Handler,
		},
		{
			MethodName: "TagDeployment",
			Handler:    _DeployService_TagDeployment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chalk/server/v1/deploy.proto",
}
