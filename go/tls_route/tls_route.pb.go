// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.0
// source: tls_route.proto

package tls_route

import (
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

// Describes how to match an `SNI` ClientHello extension.
type SniMatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Match:
	//
	//	*SniMatch_Exact
	//	*SniMatch_Suffix_
	Match isSniMatch_Match `protobuf_oneof:"match"`
}

func (x *SniMatch) Reset() {
	*x = SniMatch{}
	mi := &file_tls_route_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SniMatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SniMatch) ProtoMessage() {}

func (x *SniMatch) ProtoReflect() protoreflect.Message {
	mi := &file_tls_route_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SniMatch.ProtoReflect.Descriptor instead.
func (*SniMatch) Descriptor() ([]byte, []int) {
	return file_tls_route_proto_rawDescGZIP(), []int{0}
}

func (m *SniMatch) GetMatch() isSniMatch_Match {
	if m != nil {
		return m.Match
	}
	return nil
}

func (x *SniMatch) GetExact() string {
	if x, ok := x.GetMatch().(*SniMatch_Exact); ok {
		return x.Exact
	}
	return ""
}

func (x *SniMatch) GetSuffix() *SniMatch_Suffix {
	if x, ok := x.GetMatch().(*SniMatch_Suffix_); ok {
		return x.Suffix
	}
	return nil
}

type isSniMatch_Match interface {
	isSniMatch_Match()
}

type SniMatch_Exact struct {
	// Match an exact SNI, e.g. www.example.com.
	Exact string `protobuf:"bytes,1,opt,name=exact,proto3,oneof"`
}

type SniMatch_Suffix_ struct {
	// Match a SNI as a wildcard suffix, e.g. *.example.com.
	Suffix *SniMatch_Suffix `protobuf:"bytes,2,opt,name=suffix,proto3,oneof"`
}

func (*SniMatch_Exact) isSniMatch_Match() {}

func (*SniMatch_Suffix_) isSniMatch_Match() {}

// A match like `*.example.com` is encoded as [com, example].
type SniMatch_Suffix struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReverseLabels []string `protobuf:"bytes,1,rep,name=reverse_labels,json=reverseLabels,proto3" json:"reverse_labels,omitempty"`
}

func (x *SniMatch_Suffix) Reset() {
	*x = SniMatch_Suffix{}
	mi := &file_tls_route_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SniMatch_Suffix) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SniMatch_Suffix) ProtoMessage() {}

func (x *SniMatch_Suffix) ProtoReflect() protoreflect.Message {
	mi := &file_tls_route_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SniMatch_Suffix.ProtoReflect.Descriptor instead.
func (*SniMatch_Suffix) Descriptor() ([]byte, []int) {
	return file_tls_route_proto_rawDescGZIP(), []int{0, 0}
}

func (x *SniMatch_Suffix) GetReverseLabels() []string {
	if x != nil {
		return x.ReverseLabels
	}
	return nil
}

var File_tls_route_proto protoreflect.FileDescriptor

var file_tls_route_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x74, 0x6c, 0x73, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1a, 0x69, 0x6f, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2e, 0x74, 0x6c, 0x73, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x22, 0xa3, 0x01,
	0x0a, 0x08, 0x53, 0x6e, 0x69, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x16, 0x0a, 0x05, 0x65, 0x78,
	0x61, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x65, 0x78, 0x61,
	0x63, 0x74, 0x12, 0x45, 0x0a, 0x06, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x69, 0x6f, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x74, 0x6c, 0x73, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e,
	0x53, 0x6e, 0x69, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x53, 0x75, 0x66, 0x66, 0x69, 0x78, 0x48,
	0x00, 0x52, 0x06, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x1a, 0x2f, 0x0a, 0x06, 0x53, 0x75, 0x66,
	0x66, 0x69, 0x78, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x5f, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x76,
	0x65, 0x72, 0x73, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x42, 0x07, 0x0a, 0x05, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x64, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x72,
	0x64, 0x32, 0x2d, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f,
	0x74, 0x6c, 0x73, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_tls_route_proto_rawDescOnce sync.Once
	file_tls_route_proto_rawDescData = file_tls_route_proto_rawDesc
)

func file_tls_route_proto_rawDescGZIP() []byte {
	file_tls_route_proto_rawDescOnce.Do(func() {
		file_tls_route_proto_rawDescData = protoimpl.X.CompressGZIP(file_tls_route_proto_rawDescData)
	})
	return file_tls_route_proto_rawDescData
}

var file_tls_route_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tls_route_proto_goTypes = []any{
	(*SniMatch)(nil),        // 0: io.linkerd.proxy.tls_route.SniMatch
	(*SniMatch_Suffix)(nil), // 1: io.linkerd.proxy.tls_route.SniMatch.Suffix
}
var file_tls_route_proto_depIdxs = []int32{
	1, // 0: io.linkerd.proxy.tls_route.SniMatch.suffix:type_name -> io.linkerd.proxy.tls_route.SniMatch.Suffix
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tls_route_proto_init() }
func file_tls_route_proto_init() {
	if File_tls_route_proto != nil {
		return
	}
	file_tls_route_proto_msgTypes[0].OneofWrappers = []any{
		(*SniMatch_Exact)(nil),
		(*SniMatch_Suffix_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tls_route_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tls_route_proto_goTypes,
		DependencyIndexes: file_tls_route_proto_depIdxs,
		MessageInfos:      file_tls_route_proto_msgTypes,
	}.Build()
	File_tls_route_proto = out.File
	file_tls_route_proto_rawDesc = nil
	file_tls_route_proto_goTypes = nil
	file_tls_route_proto_depIdxs = nil
}
