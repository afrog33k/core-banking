// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: bnk.to/core/api/v1/centres/all.proto

package centres

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CentresServiceClient is the client API for CentresService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CentresServiceClient interface {
	// CreateCentre creates a new centre.
	CreateCentre(ctx context.Context, in *CreateCentreRequest, opts ...grpc.CallOption) (*Centre, error)
	// ListCentres lists centres.
	ListCentres(ctx context.Context, in *ListCentresRequest, opts ...grpc.CallOption) (*ListCentresResponse, error)
	// GetCentre retrieves the specified centre.
	GetCentre(ctx context.Context, in *GetCentreRequest, opts ...grpc.CallOption) (*Centre, error)
	// UpdateCentre updates the configuration of a centre.
	UpdateCentre(ctx context.Context, in *UpdateCentreRequest, opts ...grpc.CallOption) (*Centre, error)
	// DeleteCentre deletes the specified centre.
	DeleteCentre(ctx context.Context, in *DeleteCentreRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type centresServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCentresServiceClient(cc grpc.ClientConnInterface) CentresServiceClient {
	return &centresServiceClient{cc}
}

func (c *centresServiceClient) CreateCentre(ctx context.Context, in *CreateCentreRequest, opts ...grpc.CallOption) (*Centre, error) {
	out := new(Centre)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.centres.CentresService/CreateCentre", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *centresServiceClient) ListCentres(ctx context.Context, in *ListCentresRequest, opts ...grpc.CallOption) (*ListCentresResponse, error) {
	out := new(ListCentresResponse)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.centres.CentresService/ListCentres", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *centresServiceClient) GetCentre(ctx context.Context, in *GetCentreRequest, opts ...grpc.CallOption) (*Centre, error) {
	out := new(Centre)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.centres.CentresService/GetCentre", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *centresServiceClient) UpdateCentre(ctx context.Context, in *UpdateCentreRequest, opts ...grpc.CallOption) (*Centre, error) {
	out := new(Centre)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.centres.CentresService/UpdateCentre", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *centresServiceClient) DeleteCentre(ctx context.Context, in *DeleteCentreRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.centres.CentresService/DeleteCentre", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CentresServiceServer is the server API for CentresService service.
// All implementations must embed UnimplementedCentresServiceServer
// for forward compatibility
type CentresServiceServer interface {
	// CreateCentre creates a new centre.
	CreateCentre(context.Context, *CreateCentreRequest) (*Centre, error)
	// ListCentres lists centres.
	ListCentres(context.Context, *ListCentresRequest) (*ListCentresResponse, error)
	// GetCentre retrieves the specified centre.
	GetCentre(context.Context, *GetCentreRequest) (*Centre, error)
	// UpdateCentre updates the configuration of a centre.
	UpdateCentre(context.Context, *UpdateCentreRequest) (*Centre, error)
	// DeleteCentre deletes the specified centre.
	DeleteCentre(context.Context, *DeleteCentreRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedCentresServiceServer()
}

// UnimplementedCentresServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCentresServiceServer struct{}

func (UnimplementedCentresServiceServer) CreateCentre(context.Context, *CreateCentreRequest) (*Centre, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCentre not implemented")
}

func (UnimplementedCentresServiceServer) ListCentres(context.Context, *ListCentresRequest) (*ListCentresResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCentres not implemented")
}

func (UnimplementedCentresServiceServer) GetCentre(context.Context, *GetCentreRequest) (*Centre, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCentre not implemented")
}

func (UnimplementedCentresServiceServer) UpdateCentre(context.Context, *UpdateCentreRequest) (*Centre, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCentre not implemented")
}

func (UnimplementedCentresServiceServer) DeleteCentre(context.Context, *DeleteCentreRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCentre not implemented")
}
func (UnimplementedCentresServiceServer) mustEmbedUnimplementedCentresServiceServer() {}

// UnsafeCentresServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CentresServiceServer will
// result in compilation errors.
type UnsafeCentresServiceServer interface {
	mustEmbedUnimplementedCentresServiceServer()
}

func RegisterCentresServiceServer(s grpc.ServiceRegistrar, srv CentresServiceServer) {
	s.RegisterService(&CentresService_ServiceDesc, srv)
}

func _CentresService_CreateCentre_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCentreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CentresServiceServer).CreateCentre(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.centres.CentresService/CreateCentre",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CentresServiceServer).CreateCentre(ctx, req.(*CreateCentreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CentresService_ListCentres_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCentresRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CentresServiceServer).ListCentres(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.centres.CentresService/ListCentres",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CentresServiceServer).ListCentres(ctx, req.(*ListCentresRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CentresService_GetCentre_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCentreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CentresServiceServer).GetCentre(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.centres.CentresService/GetCentre",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CentresServiceServer).GetCentre(ctx, req.(*GetCentreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CentresService_UpdateCentre_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCentreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CentresServiceServer).UpdateCentre(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.centres.CentresService/UpdateCentre",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CentresServiceServer).UpdateCentre(ctx, req.(*UpdateCentreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CentresService_DeleteCentre_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCentreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CentresServiceServer).DeleteCentre(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.centres.CentresService/DeleteCentre",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CentresServiceServer).DeleteCentre(ctx, req.(*DeleteCentreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CentresService_ServiceDesc is the grpc.ServiceDesc for CentresService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CentresService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "openbank.core.v1.centres.CentresService",
	HandlerType: (*CentresServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCentre",
			Handler:    _CentresService_CreateCentre_Handler,
		},
		{
			MethodName: "ListCentres",
			Handler:    _CentresService_ListCentres_Handler,
		},
		{
			MethodName: "GetCentre",
			Handler:    _CentresService_GetCentre_Handler,
		},
		{
			MethodName: "UpdateCentre",
			Handler:    _CentresService_UpdateCentre_Handler,
		},
		{
			MethodName: "DeleteCentre",
			Handler:    _CentresService_DeleteCentre_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bnk.to/core/api/v1/centres/all.proto",
}
