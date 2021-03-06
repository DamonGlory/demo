// Code generated by protoc-gen-go. DO NOT EDIT.
// source: getToken.proto

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
type GetTokenRequest struct {
	//调用注册方法时需要携带的参数
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTokenRequest) Reset()         { *m = GetTokenRequest{} }
func (m *GetTokenRequest) String() string { return proto.CompactTextString(m) }
func (*GetTokenRequest) ProtoMessage()    {}
func (*GetTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3cae582c542d4d4, []int{0}
}

func (m *GetTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTokenRequest.Unmarshal(m, b)
}
func (m *GetTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTokenRequest.Marshal(b, m, deterministic)
}
func (m *GetTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTokenRequest.Merge(m, src)
}
func (m *GetTokenRequest) XXX_Size() int {
	return xxx_messageInfo_GetTokenRequest.Size(m)
}
func (m *GetTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTokenRequest proto.InternalMessageInfo

func (m *GetTokenRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Response 响应结构
type GetTokenResponse struct {
	Recode               int32    `protobuf:"varint,1,opt,name=recode,proto3" json:"recode,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Token                string   `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTokenResponse) Reset()         { *m = GetTokenResponse{} }
func (m *GetTokenResponse) String() string { return proto.CompactTextString(m) }
func (*GetTokenResponse) ProtoMessage()    {}
func (*GetTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3cae582c542d4d4, []int{1}
}

func (m *GetTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTokenResponse.Unmarshal(m, b)
}
func (m *GetTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTokenResponse.Marshal(b, m, deterministic)
}
func (m *GetTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTokenResponse.Merge(m, src)
}
func (m *GetTokenResponse) XXX_Size() int {
	return xxx_messageInfo_GetTokenResponse.Size(m)
}
func (m *GetTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTokenResponse proto.InternalMessageInfo

func (m *GetTokenResponse) GetRecode() int32 {
	if m != nil {
		return m.Recode
	}
	return 0
}

func (m *GetTokenResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *GetTokenResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*GetTokenRequest)(nil), "protobuf.GetTokenRequest")
	proto.RegisterType((*GetTokenResponse)(nil), "protobuf.GetTokenResponse")
}

func init() { proto.RegisterFile("getToken.proto", fileDescriptor_b3cae582c542d4d4) }

var fileDescriptor_b3cae582c542d4d4 = []byte{
	// 178 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0x4f, 0x2d, 0x09,
	0xc9, 0xcf, 0x4e, 0xcd, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0x49, 0xa5,
	0x69, 0x4a, 0xaa, 0x5c, 0xfc, 0xee, 0x50, 0xb9, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21,
	0x21, 0x2e, 0x96, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b,
	0x29, 0x8a, 0x4b, 0x00, 0xa1, 0xac, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x48, 0x8c, 0x8b, 0xad,
	0x28, 0x35, 0x39, 0x3f, 0x05, 0xa2, 0x92, 0x35, 0x08, 0xca, 0x13, 0x92, 0xe0, 0x62, 0xcf, 0x4d,
	0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x95, 0x60, 0x02, 0x1b, 0x01, 0xe3, 0x0a, 0x89, 0x70, 0xb1, 0x96,
	0x80, 0x8c, 0x90, 0x60, 0x06, 0x8b, 0x43, 0x38, 0x46, 0x61, 0x08, 0x27, 0x04, 0xa7, 0x16, 0x95,
	0x65, 0x26, 0xa7, 0x0a, 0x39, 0x73, 0x71, 0xc0, 0x84, 0x84, 0x24, 0xf5, 0x60, 0x8e, 0xd5, 0x43,
	0x73, 0xa9, 0x94, 0x14, 0x36, 0x29, 0x88, 0xeb, 0x94, 0x18, 0x9c, 0xb8, 0xa2, 0xe0, 0xde, 0x4c,
	0x62, 0x03, 0xb3, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa9, 0x3f, 0x6c, 0x6d, 0x09, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GetTokenServiceClient is the client API for GetTokenService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GetTokenServiceClient interface {
	// Register 测试方法
	GetToken(ctx context.Context, in *GetTokenRequest, opts ...grpc.CallOption) (*GetTokenResponse, error)
}

type getTokenServiceClient struct {
	cc *grpc.ClientConn
}

func NewGetTokenServiceClient(cc *grpc.ClientConn) GetTokenServiceClient {
	return &getTokenServiceClient{cc}
}

func (c *getTokenServiceClient) GetToken(ctx context.Context, in *GetTokenRequest, opts ...grpc.CallOption) (*GetTokenResponse, error) {
	out := new(GetTokenResponse)
	err := c.cc.Invoke(ctx, "/protobuf.GetTokenService/GetToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetTokenServiceServer is the server API for GetTokenService service.
type GetTokenServiceServer interface {
	// Register 测试方法
	GetToken(context.Context, *GetTokenRequest) (*GetTokenResponse, error)
}

// UnimplementedGetTokenServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGetTokenServiceServer struct {
}

func (*UnimplementedGetTokenServiceServer) GetToken(ctx context.Context, req *GetTokenRequest) (*GetTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetToken not implemented")
}

func RegisterGetTokenServiceServer(s *grpc.Server, srv GetTokenServiceServer) {
	s.RegisterService(&_GetTokenService_serviceDesc, srv)
}

func _GetTokenService_GetToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetTokenServiceServer).GetToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.GetTokenService/GetToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetTokenServiceServer).GetToken(ctx, req.(*GetTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GetTokenService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.GetTokenService",
	HandlerType: (*GetTokenServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetToken",
			Handler:    _GetTokenService_GetToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "getToken.proto",
}
