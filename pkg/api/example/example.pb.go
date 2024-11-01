// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: example.proto

package example

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	types "github.com/gogo/protobuf/types"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type FooRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FooRequest) Reset()         { *m = FooRequest{} }
func (m *FooRequest) String() string { return proto.CompactTextString(m) }
func (*FooRequest) ProtoMessage()    {}
func (*FooRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{0}
}
func (m *FooRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FooRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FooRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FooRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FooRequest.Merge(m, src)
}
func (m *FooRequest) XXX_Size() int {
	return m.Size()
}
func (m *FooRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FooRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FooRequest proto.InternalMessageInfo

func (m *FooRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (*FooRequest) XXX_MessageName() string {
	return "api.example.FooRequest"
}

type FooResponse struct {
	Created              *time.Time `protobuf:"bytes,1,opt,name=created,proto3,stdtime" json:"created,omitempty"`
	Id                   int64      `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string     `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *FooResponse) Reset()         { *m = FooResponse{} }
func (m *FooResponse) String() string { return proto.CompactTextString(m) }
func (*FooResponse) ProtoMessage()    {}
func (*FooResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{1}
}
func (m *FooResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FooResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FooResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FooResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FooResponse.Merge(m, src)
}
func (m *FooResponse) XXX_Size() int {
	return m.Size()
}
func (m *FooResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FooResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FooResponse proto.InternalMessageInfo

func (m *FooResponse) GetCreated() *time.Time {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *FooResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *FooResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (*FooResponse) XXX_MessageName() string {
	return "api.example.FooResponse"
}

type BarRequest struct {
	Data                 []*BarRequest_BarData `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *BarRequest) Reset()         { *m = BarRequest{} }
func (m *BarRequest) String() string { return proto.CompactTextString(m) }
func (*BarRequest) ProtoMessage()    {}
func (*BarRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{2}
}
func (m *BarRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BarRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BarRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BarRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BarRequest.Merge(m, src)
}
func (m *BarRequest) XXX_Size() int {
	return m.Size()
}
func (m *BarRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BarRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BarRequest proto.InternalMessageInfo

func (m *BarRequest) GetData() []*BarRequest_BarData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (*BarRequest) XXX_MessageName() string {
	return "api.example.BarRequest"
}

type BarRequest_BarData struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" validate:"required"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BarRequest_BarData) Reset()         { *m = BarRequest_BarData{} }
func (m *BarRequest_BarData) String() string { return proto.CompactTextString(m) }
func (*BarRequest_BarData) ProtoMessage()    {}
func (*BarRequest_BarData) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{2, 0}
}
func (m *BarRequest_BarData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BarRequest_BarData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BarRequest_BarData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BarRequest_BarData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BarRequest_BarData.Merge(m, src)
}
func (m *BarRequest_BarData) XXX_Size() int {
	return m.Size()
}
func (m *BarRequest_BarData) XXX_DiscardUnknown() {
	xxx_messageInfo_BarRequest_BarData.DiscardUnknown(m)
}

var xxx_messageInfo_BarRequest_BarData proto.InternalMessageInfo

func (m *BarRequest_BarData) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *BarRequest_BarData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (*BarRequest_BarData) XXX_MessageName() string {
	return "api.example.BarRequest.BarData"
}

type HealthRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthRequest) Reset()         { *m = HealthRequest{} }
func (m *HealthRequest) String() string { return proto.CompactTextString(m) }
func (*HealthRequest) ProtoMessage()    {}
func (*HealthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{3}
}
func (m *HealthRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HealthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HealthRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HealthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthRequest.Merge(m, src)
}
func (m *HealthRequest) XXX_Size() int {
	return m.Size()
}
func (m *HealthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HealthRequest proto.InternalMessageInfo

func (*HealthRequest) XXX_MessageName() string {
	return "api.example.HealthRequest"
}

type HealthResponse struct {
	Status               int64    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthResponse) Reset()         { *m = HealthResponse{} }
func (m *HealthResponse) String() string { return proto.CompactTextString(m) }
func (*HealthResponse) ProtoMessage()    {}
func (*HealthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{4}
}
func (m *HealthResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HealthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HealthResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HealthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthResponse.Merge(m, src)
}
func (m *HealthResponse) XXX_Size() int {
	return m.Size()
}
func (m *HealthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HealthResponse proto.InternalMessageInfo

func (m *HealthResponse) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (*HealthResponse) XXX_MessageName() string {
	return "api.example.HealthResponse"
}
func init() {
	proto.RegisterType((*FooRequest)(nil), "api.example.FooRequest")
	proto.RegisterType((*FooResponse)(nil), "api.example.FooResponse")
	proto.RegisterType((*BarRequest)(nil), "api.example.BarRequest")
	proto.RegisterType((*BarRequest_BarData)(nil), "api.example.BarRequest.BarData")
	proto.RegisterType((*HealthRequest)(nil), "api.example.HealthRequest")
	proto.RegisterType((*HealthResponse)(nil), "api.example.HealthResponse")
}

func init() { proto.RegisterFile("example.proto", fileDescriptor_15a1dc8d40dadaa6) }

var fileDescriptor_15a1dc8d40dadaa6 = []byte{
	// 472 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x65, 0xdd, 0x2a, 0xa6, 0x63, 0xa5, 0x45, 0x5b, 0x29, 0x8d, 0xdc, 0x62, 0x47, 0x7b, 0x21,
	0x42, 0xc2, 0x16, 0xe9, 0xad, 0x07, 0x24, 0x2c, 0x1a, 0xe0, 0x16, 0x59, 0x9c, 0xb8, 0x4d, 0xea,
	0xad, 0xb3, 0x52, 0xec, 0x75, 0xec, 0x75, 0x05, 0x57, 0x0e, 0x48, 0xdc, 0x90, 0xb8, 0xf0, 0x49,
	0x39, 0x22, 0x71, 0x2f, 0x28, 0xe5, 0x0b, 0xf8, 0x02, 0xd4, 0xf5, 0xba, 0x69, 0x5a, 0x6e, 0x33,
	0xf3, 0x9e, 0xdf, 0xbc, 0x79, 0x6b, 0xe8, 0xf2, 0x0f, 0x98, 0x15, 0x73, 0x1e, 0x14, 0xa5, 0x54,
	0x92, 0x3a, 0x58, 0x88, 0xc0, 0x8c, 0xdc, 0xa3, 0x54, 0xca, 0x74, 0xce, 0x43, 0x2c, 0x44, 0x88,
	0x79, 0x2e, 0x15, 0x2a, 0x21, 0xf3, 0xaa, 0xa1, 0xba, 0xcf, 0x52, 0xa1, 0x66, 0xf5, 0x34, 0x38,
	0x93, 0x59, 0x98, 0xca, 0x54, 0x86, 0x7a, 0x3c, 0xad, 0xcf, 0x75, 0xa7, 0x1b, 0x5d, 0x19, 0xba,
	0x6f, 0xc4, 0x6e, 0x58, 0x4a, 0x64, 0xbc, 0x52, 0x98, 0x15, 0x86, 0x70, 0x78, 0x97, 0xc0, 0xb3,
	0x42, 0x7d, 0x6c, 0x40, 0x76, 0x04, 0x30, 0x96, 0x32, 0xe6, 0x8b, 0x9a, 0x57, 0x8a, 0xee, 0x82,
	0x25, 0x92, 0x3e, 0x19, 0x90, 0xe1, 0x56, 0x6c, 0x89, 0x84, 0x2d, 0xc0, 0xd1, 0x68, 0x55, 0xc8,
	0xbc, 0xe2, 0xf4, 0x05, 0xd8, 0x67, 0x25, 0x47, 0xc5, 0x1b, 0x8e, 0x33, 0x72, 0x83, 0x46, 0x3b,
	0x68, 0xb5, 0x83, 0x77, 0xed, 0xf2, 0xe8, 0xe1, 0xf2, 0xd2, 0x27, 0x5f, 0x7f, 0xf9, 0x24, 0x6e,
	0x3f, 0x32, 0xf2, 0x56, 0x2b, 0x4f, 0x29, 0x6c, 0xe7, 0x98, 0xf1, 0xfe, 0xd6, 0x80, 0x0c, 0x77,
	0x62, 0x5d, 0xb3, 0x2f, 0x04, 0x20, 0xc2, 0xb2, 0x75, 0x74, 0x0c, 0xdb, 0x09, 0x2a, 0xec, 0x93,
	0xc1, 0xd6, 0xd0, 0x19, 0xf9, 0xc1, 0xad, 0x18, 0x83, 0x35, 0xed, 0xba, 0x7c, 0x85, 0x0a, 0x63,
	0x4d, 0x76, 0xc7, 0x60, 0x9b, 0x01, 0x7d, 0xb2, 0xbe, 0x28, 0x3a, 0xf8, 0x7b, 0xe9, 0xef, 0x5f,
	0xe0, 0x5c, 0x24, 0xa8, 0xf8, 0x09, 0x2b, 0xf9, 0xa2, 0x16, 0x25, 0x4f, 0xd8, 0x86, 0x17, 0xeb,
	0x96, 0x97, 0x3d, 0xe8, 0xbe, 0xe1, 0x38, 0x57, 0x33, 0xb3, 0x86, 0x0d, 0x61, 0xb7, 0x1d, 0x98,
	0x48, 0x7a, 0xd0, 0xa9, 0x14, 0xaa, 0xba, 0x32, 0xa9, 0x99, 0x6e, 0xf4, 0xd9, 0x02, 0x38, 0x6d,
	0x7c, 0xbe, 0x9c, 0xbc, 0xa5, 0x13, 0xe8, 0xbc, 0xe6, 0x6a, 0x2c, 0x25, 0x3d, 0xd8, 0x38, 0x61,
	0x9d, 0xbd, 0xdb, 0xbf, 0x0f, 0x34, 0x3b, 0xd8, 0xfe, 0xa7, 0x9f, 0x7f, 0xbe, 0x59, 0x5d, 0xea,
	0xe8, 0x1f, 0xe6, 0xe2, 0x79, 0x78, 0x2e, 0x25, 0x8d, 0xc1, 0x9e, 0xc8, 0x4a, 0x45, 0x58, 0xde,
	0x91, 0x5c, 0xa7, 0xe2, 0xf6, 0xee, 0x3d, 0xcf, 0xe9, 0xf5, 0xd3, 0xb3, 0x9e, 0x16, 0x7c, 0xc4,
	0x6e, 0x04, 0xa7, 0x58, 0x9e, 0x90, 0xa7, 0x34, 0x86, 0x4e, 0x73, 0x1e, 0x75, 0x37, 0x24, 0x37,
	0x42, 0x70, 0x0f, 0xff, 0x8b, 0x19, 0xaf, 0x7b, 0x5a, 0x7a, 0x87, 0xda, 0xe1, 0x4c, 0x03, 0xd1,
	0xe3, 0xe5, 0xca, 0x23, 0x3f, 0x56, 0x1e, 0xf9, 0xbd, 0xf2, 0xc8, 0xf7, 0x2b, 0xef, 0xc1, 0xf2,
	0xca, 0x23, 0xef, 0x6d, 0xf3, 0xf9, 0xb4, 0xa3, 0xad, 0x1d, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff,
	0x59, 0xf5, 0x34, 0xde, 0x2f, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ExampleAPIClient is the client API for ExampleAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExampleAPIClient interface {
	GetFoo(ctx context.Context, in *FooRequest, opts ...grpc.CallOption) (*FooResponse, error)
	PostBar(ctx context.Context, in *BarRequest, opts ...grpc.CallOption) (*types.Empty, error)
	Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error)
}

type exampleAPIClient struct {
	cc *grpc.ClientConn
}

func NewExampleAPIClient(cc *grpc.ClientConn) ExampleAPIClient {
	return &exampleAPIClient{cc}
}

func (c *exampleAPIClient) GetFoo(ctx context.Context, in *FooRequest, opts ...grpc.CallOption) (*FooResponse, error) {
	out := new(FooResponse)
	err := c.cc.Invoke(ctx, "/api.example.ExampleAPI/GetFoo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleAPIClient) PostBar(ctx context.Context, in *BarRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/api.example.ExampleAPI/PostBar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleAPIClient) Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error) {
	out := new(HealthResponse)
	err := c.cc.Invoke(ctx, "/api.example.ExampleAPI/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExampleAPIServer is the server API for ExampleAPI service.
type ExampleAPIServer interface {
	GetFoo(context.Context, *FooRequest) (*FooResponse, error)
	PostBar(context.Context, *BarRequest) (*types.Empty, error)
	Health(context.Context, *HealthRequest) (*HealthResponse, error)
}

// UnimplementedExampleAPIServer can be embedded to have forward compatible implementations.
type UnimplementedExampleAPIServer struct {
}

func (*UnimplementedExampleAPIServer) GetFoo(ctx context.Context, req *FooRequest) (*FooResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFoo not implemented")
}
func (*UnimplementedExampleAPIServer) PostBar(ctx context.Context, req *BarRequest) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostBar not implemented")
}
func (*UnimplementedExampleAPIServer) Health(ctx context.Context, req *HealthRequest) (*HealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}

func RegisterExampleAPIServer(s *grpc.Server, srv ExampleAPIServer) {
	s.RegisterService(&_ExampleAPI_serviceDesc, srv)
}

func _ExampleAPI_GetFoo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FooRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleAPIServer).GetFoo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.example.ExampleAPI/GetFoo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleAPIServer).GetFoo(ctx, req.(*FooRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExampleAPI_PostBar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleAPIServer).PostBar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.example.ExampleAPI/PostBar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleAPIServer).PostBar(ctx, req.(*BarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExampleAPI_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleAPIServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.example.ExampleAPI/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleAPIServer).Health(ctx, req.(*HealthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ExampleAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.example.ExampleAPI",
	HandlerType: (*ExampleAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFoo",
			Handler:    _ExampleAPI_GetFoo_Handler,
		},
		{
			MethodName: "PostBar",
			Handler:    _ExampleAPI_PostBar_Handler,
		},
		{
			MethodName: "Health",
			Handler:    _ExampleAPI_Health_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example.proto",
}

func (m *FooRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FooRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FooRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Id != 0 {
		i = encodeVarintExample(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *FooResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FooResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FooResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintExample(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Id != 0 {
		i = encodeVarintExample(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if m.Created != nil {
		n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.Created, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(*m.Created):])
		if err1 != nil {
			return 0, err1
		}
		i -= n1
		i = encodeVarintExample(dAtA, i, uint64(n1))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BarRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BarRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BarRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Data) > 0 {
		for iNdEx := len(m.Data) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Data[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintExample(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *BarRequest_BarData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BarRequest_BarData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BarRequest_BarData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintExample(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintExample(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *HealthRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HealthRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HealthRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func (m *HealthResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HealthResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HealthResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Status != 0 {
		i = encodeVarintExample(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintExample(dAtA []byte, offset int, v uint64) int {
	offset -= sovExample(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FooRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovExample(uint64(m.Id))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *FooResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Created != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.Created)
		n += 1 + l + sovExample(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovExample(uint64(m.Id))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovExample(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *BarRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Data) > 0 {
		for _, e := range m.Data {
			l = e.Size()
			n += 1 + l + sovExample(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *BarRequest_BarData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovExample(uint64(m.Id))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovExample(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *HealthRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *HealthResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Status != 0 {
		n += 1 + sovExample(uint64(m.Status))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovExample(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozExample(x uint64) (n int) {
	return sovExample(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FooRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExample
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FooRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FooRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExample
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipExample(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthExample
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *FooResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExample
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FooResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FooResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Created", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExample
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthExample
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthExample
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Created == nil {
				m.Created = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.Created, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExample
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExample
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthExample
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExample
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExample(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthExample
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BarRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExample
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BarRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BarRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExample
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthExample
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthExample
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data, &BarRequest_BarData{})
			if err := m.Data[len(m.Data)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExample(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthExample
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BarRequest_BarData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExample
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BarData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BarData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExample
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExample
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthExample
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExample
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExample(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthExample
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HealthRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExample
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HealthRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HealthRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipExample(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthExample
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HealthResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExample
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HealthResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HealthResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExample
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipExample(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthExample
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipExample(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowExample
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowExample
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowExample
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthExample
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupExample
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthExample
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthExample        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowExample          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupExample = fmt.Errorf("proto: unexpected end of group")
)
