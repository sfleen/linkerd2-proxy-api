// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.20.3
// source: grpc_route.proto

package grpc_route

import (
	http_route "github.com/linkerd/linkerd2-proxy-api/go/http_route"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GrpcRouteMatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rpc *GrpcRpcMatch `protobuf:"bytes,1,opt,name=rpc,proto3" json:"rpc,omitempty"`
	// A set of header value matches that must be satisified. This match is not
	// comprehensive, so requests may include headers that are not covered by this
	// match.
	Headers []*http_route.HeaderMatch `protobuf:"bytes,2,rep,name=headers,proto3" json:"headers,omitempty"`
}

func (x *GrpcRouteMatch) Reset() {
	*x = GrpcRouteMatch{}
	mi := &file_grpc_route_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GrpcRouteMatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrpcRouteMatch) ProtoMessage() {}

func (x *GrpcRouteMatch) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_route_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrpcRouteMatch.ProtoReflect.Descriptor instead.
func (*GrpcRouteMatch) Descriptor() ([]byte, []int) {
	return file_grpc_route_proto_rawDescGZIP(), []int{0}
}

func (x *GrpcRouteMatch) GetRpc() *GrpcRpcMatch {
	if x != nil {
		return x.Rpc
	}
	return nil
}

func (x *GrpcRouteMatch) GetHeaders() []*http_route.HeaderMatch {
	if x != nil {
		return x.Headers
	}
	return nil
}

type GrpcRpcMatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Method  string `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
}

func (x *GrpcRpcMatch) Reset() {
	*x = GrpcRpcMatch{}
	mi := &file_grpc_route_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GrpcRpcMatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrpcRpcMatch) ProtoMessage() {}

func (x *GrpcRpcMatch) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_route_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrpcRpcMatch.ProtoReflect.Descriptor instead.
func (*GrpcRpcMatch) Descriptor() ([]byte, []int) {
	return file_grpc_route_proto_rawDescGZIP(), []int{1}
}

func (x *GrpcRpcMatch) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *GrpcRpcMatch) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

// Configures a route to respond with a fixed response.
type GrpcFailureInjector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The status code to use in the `grpc-status` response. Must be specified.
	Code uint32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// An error message to log and include in the `grpc-message` header.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// If specified, the rate of requests that should be failed. If not specified,
	// ALL requests are failed.
	Ratio *http_route.Ratio `protobuf:"bytes,3,opt,name=ratio,proto3" json:"ratio,omitempty"`
}

func (x *GrpcFailureInjector) Reset() {
	*x = GrpcFailureInjector{}
	mi := &file_grpc_route_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GrpcFailureInjector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrpcFailureInjector) ProtoMessage() {}

func (x *GrpcFailureInjector) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_route_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrpcFailureInjector.ProtoReflect.Descriptor instead.
func (*GrpcFailureInjector) Descriptor() ([]byte, []int) {
	return file_grpc_route_proto_rawDescGZIP(), []int{2}
}

func (x *GrpcFailureInjector) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GrpcFailureInjector) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GrpcFailureInjector) GetRatio() *http_route.Ratio {
	if x != nil {
		return x.Ratio
	}
	return nil
}

var File_grpc_route_proto protoreflect.FileDescriptor

var file_grpc_route_proto_rawDesc = []byte{
	0x0a, 0x10, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1b, 0x69, 0x6f, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x1a,
	0x10, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x91, 0x01, 0x0a, 0x0e, 0x47, 0x72, 0x70, 0x63, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x12, 0x3b, 0x0a, 0x03, 0x72, 0x70, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x29, 0x2e, 0x69, 0x6f, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e,
	0x47, 0x72, 0x70, 0x63, 0x52, 0x70, 0x63, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x03, 0x72, 0x70,
	0x63, 0x12, 0x42, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x28, 0x2e, 0x69, 0x6f, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x07, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x73, 0x22, 0x40, 0x0a, 0x0c, 0x47, 0x72, 0x70, 0x63, 0x52, 0x70, 0x63,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x22, 0x7d, 0x0a, 0x13, 0x47, 0x72, 0x70, 0x63, 0x46,
	0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x49, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x38, 0x0a, 0x05,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6f,
	0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x68,
	0x74, 0x74, 0x70, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x52,
	0x05, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x64, 0x2f, 0x6c, 0x69, 0x6e,
	0x6b, 0x65, 0x72, 0x64, 0x32, 0x2d, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2d, 0x61, 0x70, 0x69, 0x2f,
	0x67, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_route_proto_rawDescOnce sync.Once
	file_grpc_route_proto_rawDescData = file_grpc_route_proto_rawDesc
)

func file_grpc_route_proto_rawDescGZIP() []byte {
	file_grpc_route_proto_rawDescOnce.Do(func() {
		file_grpc_route_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_route_proto_rawDescData)
	})
	return file_grpc_route_proto_rawDescData
}

var file_grpc_route_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_grpc_route_proto_goTypes = []any{
	(*GrpcRouteMatch)(nil),         // 0: io.linkerd.proxy.grpc_route.GrpcRouteMatch
	(*GrpcRpcMatch)(nil),           // 1: io.linkerd.proxy.grpc_route.GrpcRpcMatch
	(*GrpcFailureInjector)(nil),    // 2: io.linkerd.proxy.grpc_route.GrpcFailureInjector
	(*http_route.HeaderMatch)(nil), // 3: io.linkerd.proxy.http_route.HeaderMatch
	(*http_route.Ratio)(nil),       // 4: io.linkerd.proxy.http_route.Ratio
}
var file_grpc_route_proto_depIdxs = []int32{
	1, // 0: io.linkerd.proxy.grpc_route.GrpcRouteMatch.rpc:type_name -> io.linkerd.proxy.grpc_route.GrpcRpcMatch
	3, // 1: io.linkerd.proxy.grpc_route.GrpcRouteMatch.headers:type_name -> io.linkerd.proxy.http_route.HeaderMatch
	4, // 2: io.linkerd.proxy.grpc_route.GrpcFailureInjector.ratio:type_name -> io.linkerd.proxy.http_route.Ratio
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_grpc_route_proto_init() }
func file_grpc_route_proto_init() {
	if File_grpc_route_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_grpc_route_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_grpc_route_proto_goTypes,
		DependencyIndexes: file_grpc_route_proto_depIdxs,
		MessageInfos:      file_grpc_route_proto_msgTypes,
	}.Build()
	File_grpc_route_proto = out.File
	file_grpc_route_proto_rawDesc = nil
	file_grpc_route_proto_goTypes = nil
	file_grpc_route_proto_depIdxs = nil
}
