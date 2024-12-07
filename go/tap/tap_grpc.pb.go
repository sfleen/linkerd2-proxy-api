// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.0
// source: tap.proto

package tap

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
	Tap_Observe_FullMethodName = "/io.linkerd.proxy.tap.Tap/Observe"
)

// TapClient is the client API for Tap service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// A service exposed by proxy instances to setup
type TapClient interface {
	Observe(ctx context.Context, in *ObserveRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[TapEvent], error)
}

type tapClient struct {
	cc grpc.ClientConnInterface
}

func NewTapClient(cc grpc.ClientConnInterface) TapClient {
	return &tapClient{cc}
}

func (c *tapClient) Observe(ctx context.Context, in *ObserveRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[TapEvent], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Tap_ServiceDesc.Streams[0], Tap_Observe_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ObserveRequest, TapEvent]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Tap_ObserveClient = grpc.ServerStreamingClient[TapEvent]

// TapServer is the server API for Tap service.
// All implementations must embed UnimplementedTapServer
// for forward compatibility.
//
// A service exposed by proxy instances to setup
type TapServer interface {
	Observe(*ObserveRequest, grpc.ServerStreamingServer[TapEvent]) error
	mustEmbedUnimplementedTapServer()
}

// UnimplementedTapServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTapServer struct{}

func (UnimplementedTapServer) Observe(*ObserveRequest, grpc.ServerStreamingServer[TapEvent]) error {
	return status.Errorf(codes.Unimplemented, "method Observe not implemented")
}
func (UnimplementedTapServer) mustEmbedUnimplementedTapServer() {}
func (UnimplementedTapServer) testEmbeddedByValue()             {}

// UnsafeTapServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TapServer will
// result in compilation errors.
type UnsafeTapServer interface {
	mustEmbedUnimplementedTapServer()
}

func RegisterTapServer(s grpc.ServiceRegistrar, srv TapServer) {
	// If the following call pancis, it indicates UnimplementedTapServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Tap_ServiceDesc, srv)
}

func _Tap_Observe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ObserveRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TapServer).Observe(m, &grpc.GenericServerStream[ObserveRequest, TapEvent]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Tap_ObserveServer = grpc.ServerStreamingServer[TapEvent]

// Tap_ServiceDesc is the grpc.ServiceDesc for Tap service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Tap_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "io.linkerd.proxy.tap.Tap",
	HandlerType: (*TapServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Observe",
			Handler:       _Tap_Observe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "tap.proto",
}
