// Code generated by protoc-gen-go. DO NOT EDIT.
// source: register.proto

package protobuf

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

// Request 请求结构
type RegisterRequest struct {
	//调用注册方法时需要携带的参数
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1303fe8288f4efb6, []int{0}
}

func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (m *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(m, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RegisterRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

// Response 响应结构
type RegisterResponse struct {
	Recode               int32    `protobuf:"varint,1,opt,name=recode,proto3" json:"recode,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterResponse) Reset()         { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()    {}
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1303fe8288f4efb6, []int{1}
}

func (m *RegisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResponse.Unmarshal(m, b)
}
func (m *RegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResponse.Marshal(b, m, deterministic)
}
func (m *RegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResponse.Merge(m, src)
}
func (m *RegisterResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterResponse.Size(m)
}
func (m *RegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResponse proto.InternalMessageInfo

func (m *RegisterResponse) GetRecode() int32 {
	if m != nil {
		return m.Recode
	}
	return 0
}

func (m *RegisterResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*RegisterRequest)(nil), "protobuf.RegisterRequest")
	proto.RegisterType((*RegisterResponse)(nil), "protobuf.RegisterResponse")
}

func init() { proto.RegisterFile("register.proto", fileDescriptor_1303fe8288f4efb6) }

var fileDescriptor_1303fe8288f4efb6 = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x4a, 0x4d, 0xcf,
	0x2c, 0x2e, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0x49, 0xa5,
	0x69, 0x4a, 0x8e, 0x5c, 0xfc, 0x41, 0x50, 0xb9, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21,
	0x21, 0x2e, 0x96, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b,
	0x48, 0x8a, 0x8b, 0xa3, 0x20, 0xb1, 0xb8, 0xb8, 0x3c, 0xbf, 0x28, 0x45, 0x82, 0x09, 0x2c, 0x0e,
	0xe7, 0x2b, 0xb9, 0x70, 0x09, 0x20, 0x8c, 0x28, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0x12, 0xe3,
	0x62, 0x2b, 0x4a, 0x4d, 0xce, 0x4f, 0x81, 0x98, 0xc2, 0x1a, 0x04, 0xe5, 0x09, 0x49, 0x70, 0xb1,
	0xe7, 0xa6, 0x16, 0x17, 0x27, 0xa6, 0xa7, 0x42, 0x8d, 0x81, 0x71, 0x8d, 0xc2, 0x10, 0x0e, 0x09,
	0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x15, 0x72, 0xe6, 0xe2, 0x80, 0x09, 0x09, 0x49, 0xea, 0xc1,
	0x9c, 0xac, 0x87, 0xe6, 0x5e, 0x29, 0x29, 0x6c, 0x52, 0x10, 0x77, 0x28, 0x31, 0x38, 0x71, 0x45,
	0xc1, 0x3d, 0x9b, 0xc4, 0x06, 0x66, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x1b, 0xd1, 0x55,
	0x0f, 0x0f, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RegisterServiceClient is the client API for RegisterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RegisterServiceClient interface {
	// Register 测试方法
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
}

type registerServiceClient struct {
	cc *grpc.ClientConn
}

func NewRegisterServiceClient(cc *grpc.ClientConn) RegisterServiceClient {
	return &registerServiceClient{cc}
}

func (c *registerServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/protobuf.RegisterService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegisterServiceServer is the server API for RegisterService service.
type RegisterServiceServer interface {
	// Register 测试方法
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
}

// UnimplementedRegisterServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRegisterServiceServer struct {
}

func (*UnimplementedRegisterServiceServer) Register(ctx context.Context, req *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}

func RegisterRegisterServiceServer(s *grpc.Server, srv RegisterServiceServer) {
	s.RegisterService(&_RegisterService_serviceDesc, srv)
}

func _RegisterService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegisterServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.RegisterService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegisterServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RegisterService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.RegisterService",
	HandlerType: (*RegisterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _RegisterService_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "register.proto",
}
