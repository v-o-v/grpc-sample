// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/sample/sample.proto

package sample

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type HelloRequest struct {
	Abc                  string   `protobuf:"bytes,1,opt,name=abc,proto3" json:"abc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e0fbb3b12da20d9, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetAbc() string {
	if m != nil {
		return m.Abc
	}
	return ""
}

type HelloResponse struct {
	Xyz                  string   `protobuf:"bytes,1,opt,name=xyz,proto3" json:"xyz,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloResponse) Reset()         { *m = HelloResponse{} }
func (m *HelloResponse) String() string { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()    {}
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e0fbb3b12da20d9, []int{1}
}

func (m *HelloResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloResponse.Unmarshal(m, b)
}
func (m *HelloResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloResponse.Marshal(b, m, deterministic)
}
func (m *HelloResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloResponse.Merge(m, src)
}
func (m *HelloResponse) XXX_Size() int {
	return xxx_messageInfo_HelloResponse.Size(m)
}
func (m *HelloResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelloResponse proto.InternalMessageInfo

func (m *HelloResponse) GetXyz() string {
	if m != nil {
		return m.Xyz
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "sample.HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "sample.HelloResponse")
}

func init() { proto.RegisterFile("proto/sample/sample.proto", fileDescriptor_2e0fbb3b12da20d9) }

var fileDescriptor_2e0fbb3b12da20d9 = []byte{
	// 135 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x4e, 0xcc, 0x2d, 0xc8, 0x49, 0x85, 0x52, 0x7a, 0x60, 0x31, 0x21, 0x36, 0x08,
	0x4f, 0x49, 0x81, 0x8b, 0xc7, 0x23, 0x35, 0x27, 0x27, 0x3f, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8,
	0x44, 0x48, 0x80, 0x8b, 0x39, 0x31, 0x29, 0x59, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xc4,
	0x54, 0x52, 0xe4, 0xe2, 0x85, 0xaa, 0x28, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x05, 0x29, 0xa9, 0xa8,
	0xac, 0x82, 0x29, 0xa9, 0xa8, 0xac, 0x32, 0xb2, 0xe3, 0x62, 0xf1, 0xc8, 0x4f, 0x4f, 0x15, 0x32,
	0xe3, 0x62, 0x05, 0x2b, 0x15, 0x12, 0xd1, 0x83, 0x5a, 0x86, 0x6c, 0xb6, 0x94, 0x28, 0x9a, 0x28,
	0xc4, 0x3c, 0x25, 0x86, 0x24, 0x36, 0xb0, 0x9b, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xff,
	0x90, 0x80, 0x1c, 0xb0, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HogeClient is the client API for Hoge service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HogeClient interface {
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
}

type hogeClient struct {
	cc grpc.ClientConnInterface
}

func NewHogeClient(cc grpc.ClientConnInterface) HogeClient {
	return &hogeClient{cc}
}

func (c *hogeClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/sample.Hoge/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HogeServer is the server API for Hoge service.
type HogeServer interface {
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
}

// UnimplementedHogeServer can be embedded to have forward compatible implementations.
type UnimplementedHogeServer struct {
}

func (*UnimplementedHogeServer) Hello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}

func RegisterHogeServer(s *grpc.Server, srv HogeServer) {
	s.RegisterService(&_Hoge_serviceDesc, srv)
}

func _Hoge_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HogeServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sample.Hoge/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HogeServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Hoge_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sample.Hoge",
	HandlerType: (*HogeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _Hoge_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/sample/sample.proto",
}
