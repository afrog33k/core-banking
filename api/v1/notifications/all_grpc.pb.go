// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: bnk.to/core/api/v1/notifications/all.proto

package notifications

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

// NotificationsServiceClient is the client API for NotificationsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotificationsServiceClient interface {
	// CreateNotification sends a new notification.
	CreateNotification(ctx context.Context, in *CreateNotificationRequest, opts ...grpc.CallOption) (*Notification, error)
	// ListNotifications lists notifications.
	ListNotifications(ctx context.Context, in *ListNotificationsRequest, opts ...grpc.CallOption) (*ListNotificationsResponse, error)
	// GetNotification retrieves the specified notification.
	GetNotification(ctx context.Context, in *GetNotificationRequest, opts ...grpc.CallOption) (*Notification, error)
	// ResendByDate resends all notifications within a time range.
	ResendByDate(ctx context.Context, in *ResendByDateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// ResendNotification resend the specified notification.
	ResendNotification(ctx context.Context, in *ResendNotificationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type notificationsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationsServiceClient(cc grpc.ClientConnInterface) NotificationsServiceClient {
	return &notificationsServiceClient{cc}
}

func (c *notificationsServiceClient) CreateNotification(ctx context.Context, in *CreateNotificationRequest, opts ...grpc.CallOption) (*Notification, error) {
	out := new(Notification)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.notifications.NotificationsService/CreateNotification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationsServiceClient) ListNotifications(ctx context.Context, in *ListNotificationsRequest, opts ...grpc.CallOption) (*ListNotificationsResponse, error) {
	out := new(ListNotificationsResponse)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.notifications.NotificationsService/ListNotifications", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationsServiceClient) GetNotification(ctx context.Context, in *GetNotificationRequest, opts ...grpc.CallOption) (*Notification, error) {
	out := new(Notification)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.notifications.NotificationsService/GetNotification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationsServiceClient) ResendByDate(ctx context.Context, in *ResendByDateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.notifications.NotificationsService/ResendByDate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationsServiceClient) ResendNotification(ctx context.Context, in *ResendNotificationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.notifications.NotificationsService/ResendNotification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationsServiceServer is the server API for NotificationsService service.
// All implementations must embed UnimplementedNotificationsServiceServer
// for forward compatibility
type NotificationsServiceServer interface {
	// CreateNotification sends a new notification.
	CreateNotification(context.Context, *CreateNotificationRequest) (*Notification, error)
	// ListNotifications lists notifications.
	ListNotifications(context.Context, *ListNotificationsRequest) (*ListNotificationsResponse, error)
	// GetNotification retrieves the specified notification.
	GetNotification(context.Context, *GetNotificationRequest) (*Notification, error)
	// ResendByDate resends all notifications within a time range.
	ResendByDate(context.Context, *ResendByDateRequest) (*emptypb.Empty, error)
	// ResendNotification resend the specified notification.
	ResendNotification(context.Context, *ResendNotificationRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedNotificationsServiceServer()
}

// UnimplementedNotificationsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNotificationsServiceServer struct{}

func (UnimplementedNotificationsServiceServer) CreateNotification(context.Context, *CreateNotificationRequest) (*Notification, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNotification not implemented")
}

func (UnimplementedNotificationsServiceServer) ListNotifications(context.Context, *ListNotificationsRequest) (*ListNotificationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNotifications not implemented")
}

func (UnimplementedNotificationsServiceServer) GetNotification(context.Context, *GetNotificationRequest) (*Notification, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNotification not implemented")
}

func (UnimplementedNotificationsServiceServer) ResendByDate(context.Context, *ResendByDateRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResendByDate not implemented")
}

func (UnimplementedNotificationsServiceServer) ResendNotification(context.Context, *ResendNotificationRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResendNotification not implemented")
}
func (UnimplementedNotificationsServiceServer) mustEmbedUnimplementedNotificationsServiceServer() {}

// UnsafeNotificationsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotificationsServiceServer will
// result in compilation errors.
type UnsafeNotificationsServiceServer interface {
	mustEmbedUnimplementedNotificationsServiceServer()
}

func RegisterNotificationsServiceServer(s grpc.ServiceRegistrar, srv NotificationsServiceServer) {
	s.RegisterService(&NotificationsService_ServiceDesc, srv)
}

func _NotificationsService_CreateNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServiceServer).CreateNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.notifications.NotificationsService/CreateNotification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServiceServer).CreateNotification(ctx, req.(*CreateNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationsService_ListNotifications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNotificationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServiceServer).ListNotifications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.notifications.NotificationsService/ListNotifications",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServiceServer).ListNotifications(ctx, req.(*ListNotificationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationsService_GetNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServiceServer).GetNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.notifications.NotificationsService/GetNotification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServiceServer).GetNotification(ctx, req.(*GetNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationsService_ResendByDate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResendByDateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServiceServer).ResendByDate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.notifications.NotificationsService/ResendByDate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServiceServer).ResendByDate(ctx, req.(*ResendByDateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationsService_ResendNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResendNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServiceServer).ResendNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.notifications.NotificationsService/ResendNotification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServiceServer).ResendNotification(ctx, req.(*ResendNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NotificationsService_ServiceDesc is the grpc.ServiceDesc for NotificationsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NotificationsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "openbank.core.v1.notifications.NotificationsService",
	HandlerType: (*NotificationsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNotification",
			Handler:    _NotificationsService_CreateNotification_Handler,
		},
		{
			MethodName: "ListNotifications",
			Handler:    _NotificationsService_ListNotifications_Handler,
		},
		{
			MethodName: "GetNotification",
			Handler:    _NotificationsService_GetNotification_Handler,
		},
		{
			MethodName: "ResendByDate",
			Handler:    _NotificationsService_ResendByDate_Handler,
		},
		{
			MethodName: "ResendNotification",
			Handler:    _NotificationsService_ResendNotification_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bnk.to/core/api/v1/notifications/all.proto",
}
