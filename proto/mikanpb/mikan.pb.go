// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mikanpb/mikan.proto

package mikanpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Mikan struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Kind                 string   `protobuf:"bytes,2,opt,name=Kind,proto3" json:"Kind,omitempty"`
	Quality              int64    `protobuf:"varint,3,opt,name=Quality,proto3" json:"Quality,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Mikan) Reset()         { *m = Mikan{} }
func (m *Mikan) String() string { return proto.CompactTextString(m) }
func (*Mikan) ProtoMessage()    {}
func (*Mikan) Descriptor() ([]byte, []int) {
	return fileDescriptor_d670e046051ee7bf, []int{0}
}

func (m *Mikan) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Mikan.Unmarshal(m, b)
}
func (m *Mikan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Mikan.Marshal(b, m, deterministic)
}
func (m *Mikan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mikan.Merge(m, src)
}
func (m *Mikan) XXX_Size() int {
	return xxx_messageInfo_Mikan.Size(m)
}
func (m *Mikan) XXX_DiscardUnknown() {
	xxx_messageInfo_Mikan.DiscardUnknown(m)
}

var xxx_messageInfo_Mikan proto.InternalMessageInfo

func (m *Mikan) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Mikan) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *Mikan) GetQuality() int64 {
	if m != nil {
		return m.Quality
	}
	return 0
}

type MikanRequest struct {
	Mikan                *Mikan   `protobuf:"bytes,1,opt,name=mikan,proto3" json:"mikan,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MikanRequest) Reset()         { *m = MikanRequest{} }
func (m *MikanRequest) String() string { return proto.CompactTextString(m) }
func (*MikanRequest) ProtoMessage()    {}
func (*MikanRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d670e046051ee7bf, []int{1}
}

func (m *MikanRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MikanRequest.Unmarshal(m, b)
}
func (m *MikanRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MikanRequest.Marshal(b, m, deterministic)
}
func (m *MikanRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MikanRequest.Merge(m, src)
}
func (m *MikanRequest) XXX_Size() int {
	return xxx_messageInfo_MikanRequest.Size(m)
}
func (m *MikanRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MikanRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MikanRequest proto.InternalMessageInfo

func (m *MikanRequest) GetMikan() *Mikan {
	if m != nil {
		return m.Mikan
	}
	return nil
}

type MikanResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MikanResponse) Reset()         { *m = MikanResponse{} }
func (m *MikanResponse) String() string { return proto.CompactTextString(m) }
func (*MikanResponse) ProtoMessage()    {}
func (*MikanResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d670e046051ee7bf, []int{2}
}

func (m *MikanResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MikanResponse.Unmarshal(m, b)
}
func (m *MikanResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MikanResponse.Marshal(b, m, deterministic)
}
func (m *MikanResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MikanResponse.Merge(m, src)
}
func (m *MikanResponse) XXX_Size() int {
	return xxx_messageInfo_MikanResponse.Size(m)
}
func (m *MikanResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MikanResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MikanResponse proto.InternalMessageInfo

func (m *MikanResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type RegisterMikanRequest struct {
	Mikan                *Mikan   `protobuf:"bytes,1,opt,name=mikan,proto3" json:"mikan,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterMikanRequest) Reset()         { *m = RegisterMikanRequest{} }
func (m *RegisterMikanRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterMikanRequest) ProtoMessage()    {}
func (*RegisterMikanRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d670e046051ee7bf, []int{3}
}

func (m *RegisterMikanRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterMikanRequest.Unmarshal(m, b)
}
func (m *RegisterMikanRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterMikanRequest.Marshal(b, m, deterministic)
}
func (m *RegisterMikanRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterMikanRequest.Merge(m, src)
}
func (m *RegisterMikanRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterMikanRequest.Size(m)
}
func (m *RegisterMikanRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterMikanRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterMikanRequest proto.InternalMessageInfo

func (m *RegisterMikanRequest) GetMikan() *Mikan {
	if m != nil {
		return m.Mikan
	}
	return nil
}

type RegisterMikanResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterMikanResponse) Reset()         { *m = RegisterMikanResponse{} }
func (m *RegisterMikanResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterMikanResponse) ProtoMessage()    {}
func (*RegisterMikanResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d670e046051ee7bf, []int{4}
}

func (m *RegisterMikanResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterMikanResponse.Unmarshal(m, b)
}
func (m *RegisterMikanResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterMikanResponse.Marshal(b, m, deterministic)
}
func (m *RegisterMikanResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterMikanResponse.Merge(m, src)
}
func (m *RegisterMikanResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterMikanResponse.Size(m)
}
func (m *RegisterMikanResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterMikanResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterMikanResponse proto.InternalMessageInfo

func (m *RegisterMikanResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*Mikan)(nil), "mikan.Mikan")
	proto.RegisterType((*MikanRequest)(nil), "mikan.MikanRequest")
	proto.RegisterType((*MikanResponse)(nil), "mikan.MikanResponse")
	proto.RegisterType((*RegisterMikanRequest)(nil), "mikan.RegisterMikanRequest")
	proto.RegisterType((*RegisterMikanResponse)(nil), "mikan.RegisterMikanResponse")
}

func init() { proto.RegisterFile("mikanpb/mikan.proto", fileDescriptor_d670e046051ee7bf) }

var fileDescriptor_d670e046051ee7bf = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xcf, 0x4e, 0xc2, 0x30,
	0x18, 0x4f, 0x41, 0x20, 0x14, 0x48, 0xb4, 0x20, 0x2e, 0xc8, 0x81, 0xf4, 0x22, 0xe1, 0x40, 0x23,
	0xde, 0x38, 0x7a, 0x33, 0x44, 0x13, 0xe7, 0xcd, 0x5b, 0x87, 0x5f, 0x66, 0xc3, 0xd6, 0xce, 0xb5,
	0xc3, 0x78, 0xf5, 0x15, 0x7c, 0x20, 0x1f, 0xc2, 0x57, 0xf0, 0x41, 0xcc, 0xda, 0xee, 0x30, 0x43,
	0x4c, 0x3c, 0xf5, 0xfb, 0xf7, 0xfb, 0xf3, 0xb5, 0xc5, 0xc3, 0x54, 0xec, 0xb8, 0xcc, 0x22, 0x66,
	0xcf, 0x65, 0x96, 0x2b, 0xa3, 0x48, 0xcb, 0x26, 0x93, 0x69, 0xac, 0x54, 0x9c, 0x00, 0xe3, 0x99,
	0x60, 0x5c, 0x4a, 0x65, 0xb8, 0x11, 0x4a, 0x6a, 0x37, 0x44, 0x6f, 0x70, 0xeb, 0xb6, 0x1c, 0x23,
	0x04, 0x1f, 0xdd, 0xf1, 0x14, 0x02, 0x34, 0x43, 0xf3, 0x6e, 0x68, 0xe3, 0xb2, 0xb6, 0x11, 0xf2,
	0x29, 0x68, 0xb8, 0x5a, 0x19, 0x93, 0x00, 0x77, 0xee, 0x0b, 0x9e, 0x08, 0xf3, 0x16, 0x34, 0x67,
	0x68, 0xde, 0x0c, 0xab, 0x94, 0xae, 0x70, 0xdf, 0x52, 0x85, 0xf0, 0x52, 0x80, 0x36, 0x84, 0x62,
	0xe7, 0xc0, 0x52, 0xf6, 0x56, 0xfd, 0xa5, 0x33, 0xe7, 0x66, 0x5c, 0x8b, 0x5e, 0xe0, 0x81, 0xc7,
	0xe8, 0x4c, 0x49, 0x0d, 0x64, 0x8c, 0xdb, 0x39, 0xe8, 0x22, 0x31, 0xde, 0x88, 0xcf, 0xe8, 0x1a,
	0x8f, 0x42, 0x88, 0x85, 0x36, 0x90, 0xff, 0x5b, 0x84, 0xe1, 0xd3, 0x5f, 0xd8, 0xbf, 0xc5, 0x56,
	0x9f, 0xc8, 0xaf, 0xf2, 0x00, 0xf9, 0x5e, 0x6c, 0x81, 0x6c, 0xaa, 0x5b, 0x1a, 0xd6, 0xf8, 0x9d,
	0x87, 0xc9, 0xa8, 0x5e, 0x74, 0xe4, 0xf4, 0xec, 0xfd, 0xeb, 0xfb, 0xa3, 0x71, 0x42, 0xfb, 0x6c,
	0x7f, 0xc9, 0xd2, 0x9d, 0x64, 0xb0, 0x7d, 0x56, 0x6b, 0xb4, 0x20, 0x11, 0x1e, 0xd4, 0xec, 0x90,
	0x73, 0x8f, 0x3f, 0xb4, 0xe0, 0x64, 0x7a, 0xb8, 0xe9, 0x45, 0xc6, 0x56, 0xe4, 0x98, 0xf6, 0x2a,
	0x11, 0x09, 0xaf, 0x6b, 0xb4, 0xb8, 0xee, 0x3e, 0x76, 0xfc, 0x97, 0x88, 0xda, 0xf6, 0xa1, 0xaf,
	0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0xae, 0xff, 0xb8, 0xca, 0x24, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MikanServiceClient is the client API for MikanService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MikanServiceClient interface {
	// rpc Mikan(MikanRequest) returns (MikanResponse) {};
	Mikan(ctx context.Context, in *MikanRequest, opts ...grpc.CallOption) (*MikanResponse, error)
	// comment here
	RegisterMikan(ctx context.Context, in *RegisterMikanRequest, opts ...grpc.CallOption) (*RegisterMikanResponse, error)
}

type mikanServiceClient struct {
	cc *grpc.ClientConn
}

func NewMikanServiceClient(cc *grpc.ClientConn) MikanServiceClient {
	return &mikanServiceClient{cc}
}

func (c *mikanServiceClient) Mikan(ctx context.Context, in *MikanRequest, opts ...grpc.CallOption) (*MikanResponse, error) {
	out := new(MikanResponse)
	err := c.cc.Invoke(ctx, "/mikan.MikanService/Mikan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mikanServiceClient) RegisterMikan(ctx context.Context, in *RegisterMikanRequest, opts ...grpc.CallOption) (*RegisterMikanResponse, error) {
	out := new(RegisterMikanResponse)
	err := c.cc.Invoke(ctx, "/mikan.MikanService/RegisterMikan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MikanServiceServer is the server API for MikanService service.
type MikanServiceServer interface {
	// rpc Mikan(MikanRequest) returns (MikanResponse) {};
	Mikan(context.Context, *MikanRequest) (*MikanResponse, error)
	// comment here
	RegisterMikan(context.Context, *RegisterMikanRequest) (*RegisterMikanResponse, error)
}

// UnimplementedMikanServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMikanServiceServer struct {
}

func (*UnimplementedMikanServiceServer) Mikan(ctx context.Context, req *MikanRequest) (*MikanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Mikan not implemented")
}
func (*UnimplementedMikanServiceServer) RegisterMikan(ctx context.Context, req *RegisterMikanRequest) (*RegisterMikanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterMikan not implemented")
}

func RegisterMikanServiceServer(s *grpc.Server, srv MikanServiceServer) {
	s.RegisterService(&_MikanService_serviceDesc, srv)
}

func _MikanService_Mikan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MikanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MikanServiceServer).Mikan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mikan.MikanService/Mikan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MikanServiceServer).Mikan(ctx, req.(*MikanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MikanService_RegisterMikan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterMikanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MikanServiceServer).RegisterMikan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mikan.MikanService/RegisterMikan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MikanServiceServer).RegisterMikan(ctx, req.(*RegisterMikanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MikanService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mikan.MikanService",
	HandlerType: (*MikanServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Mikan",
			Handler:    _MikanService_Mikan_Handler,
		},
		{
			MethodName: "RegisterMikan",
			Handler:    _MikanService_RegisterMikan_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mikanpb/mikan.proto",
}