// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server/pubs.proto

package server

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PubRequest struct {
	Pid                  int32    `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PubRequest) Reset()         { *m = PubRequest{} }
func (m *PubRequest) String() string { return proto.CompactTextString(m) }
func (*PubRequest) ProtoMessage()    {}
func (*PubRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_96765780516460ab, []int{0}
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

func (m *PubRequest) GetPid() int32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

type Pubs struct {
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

func (m *Pubs) Reset()         { *m = Pubs{} }
func (m *Pubs) String() string { return proto.CompactTextString(m) }
func (*Pubs) ProtoMessage()    {}
func (*Pubs) Descriptor() ([]byte, []int) {
	return fileDescriptor_96765780516460ab, []int{1}
}

func (m *Pubs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pubs.Unmarshal(m, b)
}
func (m *Pubs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pubs.Marshal(b, m, deterministic)
}
func (m *Pubs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pubs.Merge(m, src)
}
func (m *Pubs) XXX_Size() int {
	return xxx_messageInfo_Pubs.Size(m)
}
func (m *Pubs) XXX_DiscardUnknown() {
	xxx_messageInfo_Pubs.DiscardUnknown(m)
}

var xxx_messageInfo_Pubs proto.InternalMessageInfo

func (m *Pubs) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Pubs) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Pubs) GetCite() string {
	if m != nil {
		return m.Cite
	}
	return ""
}

func (m *Pubs) GetLink() string {
	if m != nil {
		return m.Link
	}
	return ""
}

func (m *Pubs) GetSlides() string {
	if m != nil {
		return m.Slides
	}
	return ""
}

func (m *Pubs) GetAbstract() string {
	if m != nil {
		return m.Abstract
	}
	return ""
}

type PubsResponse struct {
	Pubs                 []*Pubs  `protobuf:"bytes,1,rep,name=pubs,proto3" json:"pubs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PubsResponse) Reset()         { *m = PubsResponse{} }
func (m *PubsResponse) String() string { return proto.CompactTextString(m) }
func (*PubsResponse) ProtoMessage()    {}
func (*PubsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_96765780516460ab, []int{2}
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

func (m *PubsResponse) GetPubs() []*Pubs {
	if m != nil {
		return m.Pubs
	}
	return nil
}

func init() {
	proto.RegisterType((*PubRequest)(nil), "server.PubRequest")
	proto.RegisterType((*Pubs)(nil), "server.Pubs")
	proto.RegisterType((*PubsResponse)(nil), "server.PubsResponse")
}

func init() { proto.RegisterFile("server/pubs.proto", fileDescriptor_96765780516460ab) }

var fileDescriptor_96765780516460ab = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x3f, 0x4f, 0xc3, 0x30,
	0x10, 0xc5, 0x71, 0xfe, 0x58, 0xe5, 0xa8, 0x10, 0x9c, 0x2a, 0x64, 0x75, 0x40, 0x51, 0xa6, 0x4c,
	0x01, 0x15, 0x3e, 0x41, 0x17, 0xd6, 0x28, 0x6c, 0x6c, 0x75, 0x72, 0x83, 0x45, 0xd4, 0x04, 0x9f,
	0xdd, 0xcf, 0xc0, 0xc7, 0x46, 0x76, 0xa2, 0x88, 0x6e, 0xef, 0xfd, 0xde, 0x59, 0xbe, 0x77, 0xf0,
	0xc8, 0x64, 0x2f, 0x64, 0x5f, 0x26, 0xaf, 0xb9, 0x9e, 0xec, 0xe8, 0x46, 0x94, 0x33, 0x2a, 0x9f,
	0x01, 0x1a, 0xaf, 0x5b, 0xfa, 0xf1, 0xc4, 0x0e, 0x1f, 0x20, 0x9d, 0x4c, 0xaf, 0x44, 0x21, 0xaa,
	0xbc, 0x0d, 0xb2, 0xfc, 0x15, 0x90, 0x35, 0x5e, 0x33, 0xde, 0x43, 0xb2, 0x26, 0x89, 0xe9, 0x71,
	0x07, 0xb9, 0x33, 0x6e, 0x20, 0x95, 0x14, 0xa2, 0xba, 0x6d, 0x67, 0x83, 0x08, 0x59, 0x67, 0x1c,
	0xa9, 0x34, 0xc2, 0xa8, 0x03, 0x1b, 0xcc, 0xf9, 0x5b, 0x65, 0x33, 0x0b, 0x1a, 0x9f, 0x40, 0xf2,
	0x60, 0x7a, 0x62, 0x95, 0x47, 0xba, 0x38, 0xdc, 0xc3, 0xe6, 0xa4, 0xd9, 0xd9, 0x53, 0xe7, 0x94,
	0x8c, 0xc9, 0xea, 0xcb, 0x57, 0xd8, 0x86, 0x4d, 0x5a, 0xe2, 0x69, 0x3c, 0x33, 0x61, 0x01, 0x59,
	0x28, 0xa4, 0x44, 0x91, 0x56, 0x77, 0x87, 0x6d, 0x3d, 0x37, 0xaa, 0xe3, 0x4c, 0x4c, 0x0e, 0xc7,
	0x58, 0xee, 0x93, 0xec, 0xc5, 0x74, 0x84, 0xef, 0x20, 0x3f, 0xc8, 0x35, 0x5e, 0x23, 0xfe, 0x9b,
	0x5d, 0xaa, 0xef, 0x77, 0x57, 0xef, 0x97, 0x3f, 0xca, 0x9b, 0xe3, 0xe6, 0x6b, 0x39, 0x95, 0x96,
	0xf1, 0x72, 0x6f, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x58, 0x96, 0xd2, 0x08, 0x4e, 0x01, 0x00,
	0x00,
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
	GetPub(ctx context.Context, in *PubRequest, opts ...grpc.CallOption) (*PubsResponse, error)
}

type pubServiceClient struct {
	cc *grpc.ClientConn
}

func NewPubServiceClient(cc *grpc.ClientConn) PubServiceClient {
	return &pubServiceClient{cc}
}

func (c *pubServiceClient) GetPub(ctx context.Context, in *PubRequest, opts ...grpc.CallOption) (*PubsResponse, error) {
	out := new(PubsResponse)
	err := c.cc.Invoke(ctx, "/server.PubService/GetPub", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PubServiceServer is the server API for PubService service.
type PubServiceServer interface {
	// Sends a greeting
	GetPub(context.Context, *PubRequest) (*PubsResponse, error)
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
		FullMethod: "/server.PubService/GetPub",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubServiceServer).GetPub(ctx, req.(*PubRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PubService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "server.PubService",
	HandlerType: (*PubServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPub",
			Handler:    _PubService_GetPub_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/pubs.proto",
}
