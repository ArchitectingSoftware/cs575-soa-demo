// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pubs.proto

package pubs

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

type PubRequest struct {
	PubId                int32    `protobuf:"varint,1,opt,name=PubId,proto3" json:"PubId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PubRequest) Reset()         { *m = PubRequest{} }
func (m *PubRequest) String() string { return proto.CompactTextString(m) }
func (*PubRequest) ProtoMessage()    {}
func (*PubRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2c9ac9ec1a6cf4d, []int{0}
}

func (m *PubRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PubRequest.Unmarshal(m, b)
}
func (m *PubRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PubRequest.Marshal(b, m, deterministic)
}
func (m *PubRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PubRequest.Merge(m, src)
}
func (m *PubRequest) XXX_Size() int {
	return xxx_messageInfo_PubRequest.Size(m)
}
func (m *PubRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PubRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PubRequest proto.InternalMessageInfo

func (m *PubRequest) GetPubId() int32 {
	if m != nil {
		return m.PubId
	}
	return 0
}

type EmptyRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyRequest) Reset()         { *m = EmptyRequest{} }
func (m *EmptyRequest) String() string { return proto.CompactTextString(m) }
func (*EmptyRequest) ProtoMessage()    {}
func (*EmptyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2c9ac9ec1a6cf4d, []int{1}
}

func (m *EmptyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyRequest.Unmarshal(m, b)
}
func (m *EmptyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyRequest.Marshal(b, m, deterministic)
}
func (m *EmptyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyRequest.Merge(m, src)
}
func (m *EmptyRequest) XXX_Size() int {
	return xxx_messageInfo_EmptyRequest.Size(m)
}
func (m *EmptyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyRequest proto.InternalMessageInfo

type Pub struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Cite                 string   `protobuf:"bytes,3,opt,name=cite,proto3" json:"cite,omitempty"`
	Link                 string   `protobuf:"bytes,4,opt,name=link,proto3" json:"link,omitempty"`
	Slides               string   `protobuf:"bytes,5,opt,name=slides,proto3" json:"slides,omitempty"`
	Abstract             string   `protobuf:"bytes,6,opt,name=abstract,proto3" json:"abstract,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pub) Reset()         { *m = Pub{} }
func (m *Pub) String() string { return proto.CompactTextString(m) }
func (*Pub) ProtoMessage()    {}
func (*Pub) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2c9ac9ec1a6cf4d, []int{2}
}

func (m *Pub) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pub.Unmarshal(m, b)
}
func (m *Pub) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pub.Marshal(b, m, deterministic)
}
func (m *Pub) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pub.Merge(m, src)
}
func (m *Pub) XXX_Size() int {
	return xxx_messageInfo_Pub.Size(m)
}
func (m *Pub) XXX_DiscardUnknown() {
	xxx_messageInfo_Pub.DiscardUnknown(m)
}

var xxx_messageInfo_Pub proto.InternalMessageInfo

func (m *Pub) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Pub) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Pub) GetCite() string {
	if m != nil {
		return m.Cite
	}
	return ""
}

func (m *Pub) GetLink() string {
	if m != nil {
		return m.Link
	}
	return ""
}

func (m *Pub) GetSlides() string {
	if m != nil {
		return m.Slides
	}
	return ""
}

func (m *Pub) GetAbstract() string {
	if m != nil {
		return m.Abstract
	}
	return ""
}

type PubsResponse struct {
	Pubs                 []*Pub   `protobuf:"bytes,1,rep,name=Pubs,proto3" json:"Pubs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PubsResponse) Reset()         { *m = PubsResponse{} }
func (m *PubsResponse) String() string { return proto.CompactTextString(m) }
func (*PubsResponse) ProtoMessage()    {}
func (*PubsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2c9ac9ec1a6cf4d, []int{3}
}

func (m *PubsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PubsResponse.Unmarshal(m, b)
}
func (m *PubsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PubsResponse.Marshal(b, m, deterministic)
}
func (m *PubsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PubsResponse.Merge(m, src)
}
func (m *PubsResponse) XXX_Size() int {
	return xxx_messageInfo_PubsResponse.Size(m)
}
func (m *PubsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PubsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PubsResponse proto.InternalMessageInfo

func (m *PubsResponse) GetPubs() []*Pub {
	if m != nil {
		return m.Pubs
	}
	return nil
}

type PubResponse struct {
	Pub                  *Pub     `protobuf:"bytes,1,opt,name=Pub,proto3" json:"Pub,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PubResponse) Reset()         { *m = PubResponse{} }
func (m *PubResponse) String() string { return proto.CompactTextString(m) }
func (*PubResponse) ProtoMessage()    {}
func (*PubResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2c9ac9ec1a6cf4d, []int{4}
}

func (m *PubResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PubResponse.Unmarshal(m, b)
}
func (m *PubResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PubResponse.Marshal(b, m, deterministic)
}
func (m *PubResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PubResponse.Merge(m, src)
}
func (m *PubResponse) XXX_Size() int {
	return xxx_messageInfo_PubResponse.Size(m)
}
func (m *PubResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PubResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PubResponse proto.InternalMessageInfo

func (m *PubResponse) GetPub() *Pub {
	if m != nil {
		return m.Pub
	}
	return nil
}

func init() {
	proto.RegisterType((*PubRequest)(nil), "pubs.PubRequest")
	proto.RegisterType((*EmptyRequest)(nil), "pubs.EmptyRequest")
	proto.RegisterType((*Pub)(nil), "pubs.Pub")
	proto.RegisterType((*PubsResponse)(nil), "pubs.PubsResponse")
	proto.RegisterType((*PubResponse)(nil), "pubs.PubResponse")
}

func init() { proto.RegisterFile("pubs.proto", fileDescriptor_d2c9ac9ec1a6cf4d) }

var fileDescriptor_d2c9ac9ec1a6cf4d = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x41, 0x4b, 0xfb, 0x40,
	0x10, 0xc5, 0xbb, 0x6d, 0x1a, 0xfe, 0x9d, 0x96, 0xf2, 0x77, 0x10, 0x59, 0x2a, 0x42, 0xd9, 0x53,
	0x11, 0x6c, 0xa5, 0x82, 0x77, 0x05, 0x29, 0xde, 0x42, 0x7a, 0xf3, 0xd6, 0x6d, 0xe7, 0xb0, 0x98,
	0x36, 0x31, 0xbb, 0x2b, 0xf8, 0x09, 0xfc, 0x26, 0x7e, 0x4e, 0xd9, 0xd9, 0xd8, 0xc6, 0x83, 0xb7,
	0x79, 0xbf, 0xbc, 0x97, 0xe4, 0xcd, 0x00, 0x54, 0x5e, 0xdb, 0x79, 0x55, 0x97, 0xae, 0xc4, 0x24,
	0xcc, 0x4a, 0x01, 0x64, 0x5e, 0xe7, 0xf4, 0xe6, 0xc9, 0x3a, 0x3c, 0x87, 0x7e, 0xe6, 0xf5, 0xf3,
	0x4e, 0x8a, 0xa9, 0x98, 0xf5, 0xf3, 0x28, 0xd4, 0x18, 0x46, 0x4f, 0xfb, 0xca, 0x7d, 0x34, 0x2e,
	0xf5, 0x29, 0xa0, 0x97, 0x79, 0x8d, 0x63, 0xe8, 0x9a, 0x1f, 0x6b, 0xd7, 0xec, 0x42, 0xda, 0x19,
	0x57, 0x90, 0xec, 0x4e, 0xc5, 0x6c, 0x90, 0x47, 0x81, 0x08, 0xc9, 0xd6, 0x38, 0x92, 0x3d, 0x86,
	0x3c, 0x07, 0x56, 0x98, 0xc3, 0xab, 0x4c, 0x22, 0x0b, 0x33, 0x5e, 0x40, 0x6a, 0x0b, 0xb3, 0x23,
	0x2b, 0xfb, 0x4c, 0x1b, 0x85, 0x13, 0xf8, 0xb7, 0xd1, 0xd6, 0xd5, 0x9b, 0xad, 0x93, 0x29, 0x3f,
	0x39, 0x6a, 0x75, 0x03, 0xa3, 0xcc, 0x6b, 0x9b, 0x93, 0xad, 0xca, 0x83, 0x25, 0xbc, 0x82, 0x24,
	0x68, 0x29, 0xa6, 0xbd, 0xd9, 0x70, 0x39, 0x98, 0x73, 0xdd, 0xd0, 0x8f, 0xb1, 0xba, 0x86, 0x21,
	0x97, 0x6d, 0xdc, 0x97, 0x5c, 0x83, 0x0b, 0xfc, 0x32, 0x07, 0xba, 0xfc, 0x12, 0xbc, 0x99, 0x35,
	0xd5, 0xef, 0x66, 0x4b, 0xb8, 0x80, 0x74, 0x45, 0x2e, 0xb4, 0xfe, 0x7f, 0x32, 0xc6, 0x7d, 0x4c,
	0xce, 0x5a, 0x24, 0xbe, 0x5a, 0x75, 0xf0, 0x1e, 0x60, 0x45, 0xee, 0xa1, 0x28, 0xc2, 0x97, 0x11,
	0xa3, 0xa5, 0xbd, 0xc6, 0x09, 0x1e, 0x63, 0xb6, 0x95, 0x5b, 0x00, 0xac, 0x5d, 0x4d, 0x9b, 0xfd,
	0x9f, 0xb9, 0xd3, 0x9f, 0xaa, 0xce, 0xad, 0x78, 0x4c, 0x5f, 0xf8, 0x92, 0x3a, 0xe5, 0xb3, 0xde,
	0x7d, 0x07, 0x00, 0x00, 0xff, 0xff, 0xf4, 0xea, 0xc2, 0x83, 0xe4, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PubServiceClient is the client API for PubService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PubServiceClient interface {
	// Sends a greeting
	GetPub(ctx context.Context, in *PubRequest, opts ...grpc.CallOption) (*PubResponse, error)
	GetAllPubs(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*PubsResponse, error)
	StreamPubs(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (PubService_StreamPubsClient, error)
}

type pubServiceClient struct {
	cc *grpc.ClientConn
}

func NewPubServiceClient(cc *grpc.ClientConn) PubServiceClient {
	return &pubServiceClient{cc}
}

func (c *pubServiceClient) GetPub(ctx context.Context, in *PubRequest, opts ...grpc.CallOption) (*PubResponse, error) {
	out := new(PubResponse)
	err := c.cc.Invoke(ctx, "/pubs.PubService/GetPub", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubServiceClient) GetAllPubs(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*PubsResponse, error) {
	out := new(PubsResponse)
	err := c.cc.Invoke(ctx, "/pubs.PubService/GetAllPubs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubServiceClient) StreamPubs(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (PubService_StreamPubsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PubService_serviceDesc.Streams[0], "/pubs.PubService/StreamPubs", opts...)
	if err != nil {
		return nil, err
	}
	x := &pubServiceStreamPubsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PubService_StreamPubsClient interface {
	Recv() (*Pub, error)
	grpc.ClientStream
}

type pubServiceStreamPubsClient struct {
	grpc.ClientStream
}

func (x *pubServiceStreamPubsClient) Recv() (*Pub, error) {
	m := new(Pub)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PubServiceServer is the server API for PubService service.
type PubServiceServer interface {
	// Sends a greeting
	GetPub(context.Context, *PubRequest) (*PubResponse, error)
	GetAllPubs(context.Context, *EmptyRequest) (*PubsResponse, error)
	StreamPubs(*EmptyRequest, PubService_StreamPubsServer) error
}

// UnimplementedPubServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPubServiceServer struct {
}

func (*UnimplementedPubServiceServer) GetPub(ctx context.Context, req *PubRequest) (*PubResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPub not implemented")
}
func (*UnimplementedPubServiceServer) GetAllPubs(ctx context.Context, req *EmptyRequest) (*PubsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllPubs not implemented")
}
func (*UnimplementedPubServiceServer) StreamPubs(req *EmptyRequest, srv PubService_StreamPubsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamPubs not implemented")
}

func RegisterPubServiceServer(s *grpc.Server, srv PubServiceServer) {
	s.RegisterService(&_PubService_serviceDesc, srv)
}

func _PubService_GetPub_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PubRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubServiceServer).GetPub(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pubs.PubService/GetPub",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubServiceServer).GetPub(ctx, req.(*PubRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubService_GetAllPubs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubServiceServer).GetAllPubs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pubs.PubService/GetAllPubs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubServiceServer).GetAllPubs(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubService_StreamPubs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EmptyRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PubServiceServer).StreamPubs(m, &pubServiceStreamPubsServer{stream})
}

type PubService_StreamPubsServer interface {
	Send(*Pub) error
	grpc.ServerStream
}

type pubServiceStreamPubsServer struct {
	grpc.ServerStream
}

func (x *pubServiceStreamPubsServer) Send(m *Pub) error {
	return x.ServerStream.SendMsg(m)
}

var _PubService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pubs.PubService",
	HandlerType: (*PubServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPub",
			Handler:    _PubService_GetPub_Handler,
		},
		{
			MethodName: "GetAllPubs",
			Handler:    _PubService_GetAllPubs_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamPubs",
			Handler:       _PubService_StreamPubs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pubs.proto",
}
