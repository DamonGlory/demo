// Code generated by protoc-gen-go. DO NOT EDIT.
// source: setToken.proto

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
type SetTokenRequest struct {
	//调用注册方法时需要携带的参数
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=Token,proto3" json:"Token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetTokenRequest) Reset()         { *m = SetTokenRequest{} }
func (m *SetTokenRequest) String() string { return proto.CompactTextString(m) }
func (*SetTokenRequest) ProtoMessage()    {}
func (*SetTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9323a60838b33f5, []int{0}
}

func (m *SetTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetTokenRequest.Unmarshal(m, b)
}
func (m *SetTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetTokenRequest.Marshal(b, m, deterministic)
}
func (m *SetTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetTokenRequest.Merge(m, src)
}
func (m *SetTokenRequest) XXX_Size() int {
	return xxx_messageInfo_SetTokenRequest.Size(m)
}
func (m *SetTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetTokenRequest proto.InternalMessageInfo

func (m *SetTokenRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SetTokenRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// Response 响应结构
type SetTokenResponse struct {
	Recode               int32    `protobuf:"varint,1,opt,name=recode,proto3" json:"recode,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetTokenResponse) Reset()         { *m = SetTokenResponse{} }
func (m *SetTokenResponse) String() string { return proto.CompactTextString(m) }
func (*SetTokenResponse) ProtoMessage()    {}
func (*SetTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9323a60838b33f5, []int{1}
}

func (m *SetTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetTokenResponse.Unmarshal(m, b)
}
func (m *SetTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetTokenResponse.Marshal(b, m, deterministic)
}
func (m *SetTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetTokenResponse.Merge(m, src)
}
func (m *SetTokenResponse) XXX_Size() int {
	return xxx_messageInfo_SetTokenResponse.Size(m)
}
func (m *SetTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetTokenResponse proto.InternalMessageInfo

func (m *SetTokenResponse) GetRecode() int32 {
	if m != nil {
		return m.Recode
	}
	return 0
}

func (m *SetTokenResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*SetTokenRequest)(nil), "protobuf.SetTokenRequest")
	proto.RegisterType((*SetTokenResponse)(nil), "protobuf.SetTokenResponse")
}

func init() { proto.RegisterFile("setToken.proto", fileDescriptor_a9323a60838b33f5) }

var fileDescriptor_a9323a60838b33f5 = []byte{
	// 175 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x4e, 0x2d, 0x09,
	0xc9, 0xcf, 0x4e, 0xcd, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0x49, 0xa5,
	0x69, 0x4a, 0xd6, 0x5c, 0xfc, 0xc1, 0x50, 0xb9, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21,
	0x21, 0x2e, 0x96, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b,
	0x48, 0x84, 0x8b, 0x15, 0xac, 0x46, 0x82, 0x09, 0x2c, 0x08, 0xe1, 0x28, 0xb9, 0x70, 0x09, 0x20,
	0x34, 0x17, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0x89, 0x71, 0xb1, 0x15, 0xa5, 0x26, 0xe7, 0xa7,
	0x40, 0xf4, 0xb3, 0x06, 0x41, 0x79, 0x42, 0x12, 0x5c, 0xec, 0xb9, 0xa9, 0xc5, 0xc5, 0x89, 0xe9,
	0xa9, 0x50, 0x33, 0x60, 0x5c, 0xa3, 0x30, 0x84, 0x13, 0x82, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53,
	0x85, 0x9c, 0xb9, 0x38, 0x60, 0x42, 0x42, 0x92, 0x7a, 0x30, 0xc7, 0xea, 0xa1, 0xb9, 0x54, 0x4a,
	0x0a, 0x9b, 0x14, 0xc4, 0x1d, 0x4a, 0x0c, 0x4e, 0x5c, 0x51, 0x70, 0x6f, 0x26, 0xb1, 0x81, 0x59,
	0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x97, 0xfd, 0x84, 0xe7, 0x09, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SetTokenServiceClient is the client API for SetTokenService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SetTokenServiceClient interface {
	// Register 测试方法
	SetToken(ctx context.Context, in *SetTokenRequest, opts ...grpc.CallOption) (*SetTokenResponse, error)
}

type setTokenServiceClient struct {
	cc *grpc.ClientConn
}

func NewSetTokenServiceClient(cc *grpc.ClientConn) SetTokenServiceClient {
	return &setTokenServiceClient{cc}
}

func (c *setTokenServiceClient) SetToken(ctx context.Context, in *SetTokenRequest, opts ...grpc.CallOption) (*SetTokenResponse, error) {
	out := new(SetTokenResponse)
	err := c.cc.Invoke(ctx, "/protobuf.SetTokenService/SetToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SetTokenServiceServer is the server API for SetTokenService service.
type SetTokenServiceServer interface {
	// Register 测试方法
	SetToken(context.Context, *SetTokenRequest) (*SetTokenResponse, error)
}

// UnimplementedSetTokenServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSetTokenServiceServer struct {
}

func (*UnimplementedSetTokenServiceServer) SetToken(ctx context.Context, req *SetTokenRequest) (*SetTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetToken not implemented")
}

func RegisterSetTokenServiceServer(s *grpc.Server, srv SetTokenServiceServer) {
	s.RegisterService(&_SetTokenService_serviceDesc, srv)
}

func _SetTokenService_SetToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SetTokenServiceServer).SetToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.SetTokenService/SetToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SetTokenServiceServer).SetToken(ctx, req.(*SetTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SetTokenService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.SetTokenService",
	HandlerType: (*SetTokenServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetToken",
			Handler:    _SetTokenService_SetToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "setToken.proto",
}
