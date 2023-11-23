// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: proto/chatservice.proto

package proto

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
	ChatService_Message_FullMethodName = "/proto.ChatService/Message"
)

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatServiceClient interface {
	Message(ctx context.Context, opts ...grpc.CallOption) (ChatService_MessageClient, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) Message(ctx context.Context, opts ...grpc.CallOption) (ChatService_MessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChatService_ServiceDesc.Streams[0], ChatService_Message_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceMessageClient{stream}
	return x, nil
}

type ChatService_MessageClient interface {
	Send(*Request) error
	Recv() (*Reply, error)
	grpc.ClientStream
}

type chatServiceMessageClient struct {
	grpc.ClientStream
}

func (x *chatServiceMessageClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatServiceMessageClient) Recv() (*Reply, error) {
	m := new(Reply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatServiceServer is the server API for ChatService service.
// All implementations must embed UnimplementedChatServiceServer
// for forward compatibility
type ChatServiceServer interface {
	Message(ChatService_MessageServer) error
	mustEmbedUnimplementedChatServiceServer()
}

// UnimplementedChatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (UnimplementedChatServiceServer) Message(ChatService_MessageServer) error {
	return status.Errorf(codes.Unimplemented, "method Message not implemented")
}
func (UnimplementedChatServiceServer) mustEmbedUnimplementedChatServiceServer() {}

// UnsafeChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServiceServer will
// result in compilation errors.
type UnsafeChatServiceServer interface {
	mustEmbedUnimplementedChatServiceServer()
}

func RegisterChatServiceServer(s grpc.ServiceRegistrar, srv ChatServiceServer) {
	s.RegisterService(&ChatService_ServiceDesc, srv)
}

func _ChatService_Message_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServiceServer).Message(&chatServiceMessageServer{stream})
}

type ChatService_MessageServer interface {
	Send(*Reply) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type chatServiceMessageServer struct {
	grpc.ServerStream
}

func (x *chatServiceMessageServer) Send(m *Reply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatServiceMessageServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatService_ServiceDesc is the grpc.ServiceDesc for ChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Message",
			Handler:       _ChatService_Message_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/chatservice.proto",
}