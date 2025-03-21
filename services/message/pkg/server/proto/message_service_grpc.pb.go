// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: message_service.proto

package message

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
	MessagingService_SendMessage_FullMethodName    = "/message_service.MessagingService/SendMessage"
	MessagingService_EditMessage_FullMethodName    = "/message_service.MessagingService/EditMessage"
	MessagingService_DeleteMessage_FullMethodName  = "/message_service.MessagingService/DeleteMessage"
	MessagingService_GetMessageByID_FullMethodName = "/message_service.MessagingService/GetMessageByID"
	MessagingService_AddReaction_FullMethodName    = "/message_service.MessagingService/AddReaction"
	MessagingService_RemoveReaction_FullMethodName = "/message_service.MessagingService/RemoveReaction"
	MessagingService_ReadMessage_FullMethodName    = "/message_service.MessagingService/ReadMessage"
)

// MessagingServiceClient is the client API for MessagingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// / MessagingService defines the gRPC API for messaging operations.
type MessagingServiceClient interface {
	// Sends a new message to a chat.
	SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error)
	// Edits an existing message.
	EditMessage(ctx context.Context, in *EditMessageRequest, opts ...grpc.CallOption) (*EditMessageResponse, error)
	// Deletes a message.
	DeleteMessage(ctx context.Context, in *DeleteMessageRequest, opts ...grpc.CallOption) (*DeleteMessageResponse, error)
	// Gets a message by ID.
	GetMessageByID(ctx context.Context, in *GetMessageByIDRequest, opts ...grpc.CallOption) (*GetMessageByIDResponse, error)
	// Adds an emoji reaction to a message.
	AddReaction(ctx context.Context, in *AddReactionRequest, opts ...grpc.CallOption) (*AddReactionResponse, error)
	// Removes an emoji reaction from a message.
	RemoveReaction(ctx context.Context, in *RemoveReactionRequest, opts ...grpc.CallOption) (*RemoveReactionResponse, error)
	// Retrieves new messages for a chat.
	ReadMessage(ctx context.Context, in *ReadMessageRequest, opts ...grpc.CallOption) (*ReadMessageResponse, error)
}

type messagingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMessagingServiceClient(cc grpc.ClientConnInterface) MessagingServiceClient {
	return &messagingServiceClient{cc}
}

func (c *messagingServiceClient) SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendMessageResponse)
	err := c.cc.Invoke(ctx, MessagingService_SendMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingServiceClient) EditMessage(ctx context.Context, in *EditMessageRequest, opts ...grpc.CallOption) (*EditMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EditMessageResponse)
	err := c.cc.Invoke(ctx, MessagingService_EditMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingServiceClient) DeleteMessage(ctx context.Context, in *DeleteMessageRequest, opts ...grpc.CallOption) (*DeleteMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteMessageResponse)
	err := c.cc.Invoke(ctx, MessagingService_DeleteMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingServiceClient) GetMessageByID(ctx context.Context, in *GetMessageByIDRequest, opts ...grpc.CallOption) (*GetMessageByIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMessageByIDResponse)
	err := c.cc.Invoke(ctx, MessagingService_GetMessageByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingServiceClient) AddReaction(ctx context.Context, in *AddReactionRequest, opts ...grpc.CallOption) (*AddReactionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddReactionResponse)
	err := c.cc.Invoke(ctx, MessagingService_AddReaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingServiceClient) RemoveReaction(ctx context.Context, in *RemoveReactionRequest, opts ...grpc.CallOption) (*RemoveReactionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemoveReactionResponse)
	err := c.cc.Invoke(ctx, MessagingService_RemoveReaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingServiceClient) ReadMessage(ctx context.Context, in *ReadMessageRequest, opts ...grpc.CallOption) (*ReadMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReadMessageResponse)
	err := c.cc.Invoke(ctx, MessagingService_ReadMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessagingServiceServer is the server API for MessagingService service.
// All implementations must embed UnimplementedMessagingServiceServer
// for forward compatibility.
//
// / MessagingService defines the gRPC API for messaging operations.
type MessagingServiceServer interface {
	// Sends a new message to a chat.
	SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error)
	// Edits an existing message.
	EditMessage(context.Context, *EditMessageRequest) (*EditMessageResponse, error)
	// Deletes a message.
	DeleteMessage(context.Context, *DeleteMessageRequest) (*DeleteMessageResponse, error)
	// Gets a message by ID.
	GetMessageByID(context.Context, *GetMessageByIDRequest) (*GetMessageByIDResponse, error)
	// Adds an emoji reaction to a message.
	AddReaction(context.Context, *AddReactionRequest) (*AddReactionResponse, error)
	// Removes an emoji reaction from a message.
	RemoveReaction(context.Context, *RemoveReactionRequest) (*RemoveReactionResponse, error)
	// Retrieves new messages for a chat.
	ReadMessage(context.Context, *ReadMessageRequest) (*ReadMessageResponse, error)
	mustEmbedUnimplementedMessagingServiceServer()
}

// UnimplementedMessagingServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMessagingServiceServer struct{}

func (UnimplementedMessagingServiceServer) SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedMessagingServiceServer) EditMessage(context.Context, *EditMessageRequest) (*EditMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditMessage not implemented")
}
func (UnimplementedMessagingServiceServer) DeleteMessage(context.Context, *DeleteMessageRequest) (*DeleteMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMessage not implemented")
}
func (UnimplementedMessagingServiceServer) GetMessageByID(context.Context, *GetMessageByIDRequest) (*GetMessageByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessageByID not implemented")
}
func (UnimplementedMessagingServiceServer) AddReaction(context.Context, *AddReactionRequest) (*AddReactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddReaction not implemented")
}
func (UnimplementedMessagingServiceServer) RemoveReaction(context.Context, *RemoveReactionRequest) (*RemoveReactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveReaction not implemented")
}
func (UnimplementedMessagingServiceServer) ReadMessage(context.Context, *ReadMessageRequest) (*ReadMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadMessage not implemented")
}
func (UnimplementedMessagingServiceServer) mustEmbedUnimplementedMessagingServiceServer() {}
func (UnimplementedMessagingServiceServer) testEmbeddedByValue()                          {}

// UnsafeMessagingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessagingServiceServer will
// result in compilation errors.
type UnsafeMessagingServiceServer interface {
	mustEmbedUnimplementedMessagingServiceServer()
}

func RegisterMessagingServiceServer(s grpc.ServiceRegistrar, srv MessagingServiceServer) {
	// If the following call pancis, it indicates UnimplementedMessagingServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MessagingService_ServiceDesc, srv)
}

func _MessagingService_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServiceServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessagingService_SendMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServiceServer).SendMessage(ctx, req.(*SendMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessagingService_EditMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServiceServer).EditMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessagingService_EditMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServiceServer).EditMessage(ctx, req.(*EditMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessagingService_DeleteMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServiceServer).DeleteMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessagingService_DeleteMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServiceServer).DeleteMessage(ctx, req.(*DeleteMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessagingService_GetMessageByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessageByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServiceServer).GetMessageByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessagingService_GetMessageByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServiceServer).GetMessageByID(ctx, req.(*GetMessageByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessagingService_AddReaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddReactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServiceServer).AddReaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessagingService_AddReaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServiceServer).AddReaction(ctx, req.(*AddReactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessagingService_RemoveReaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveReactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServiceServer).RemoveReaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessagingService_RemoveReaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServiceServer).RemoveReaction(ctx, req.(*RemoveReactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessagingService_ReadMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServiceServer).ReadMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessagingService_ReadMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServiceServer).ReadMessage(ctx, req.(*ReadMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MessagingService_ServiceDesc is the grpc.ServiceDesc for MessagingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessagingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "message_service.MessagingService",
	HandlerType: (*MessagingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessage",
			Handler:    _MessagingService_SendMessage_Handler,
		},
		{
			MethodName: "EditMessage",
			Handler:    _MessagingService_EditMessage_Handler,
		},
		{
			MethodName: "DeleteMessage",
			Handler:    _MessagingService_DeleteMessage_Handler,
		},
		{
			MethodName: "GetMessageByID",
			Handler:    _MessagingService_GetMessageByID_Handler,
		},
		{
			MethodName: "AddReaction",
			Handler:    _MessagingService_AddReaction_Handler,
		},
		{
			MethodName: "RemoveReaction",
			Handler:    _MessagingService_RemoveReaction_Handler,
		},
		{
			MethodName: "ReadMessage",
			Handler:    _MessagingService_ReadMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message_service.proto",
}
