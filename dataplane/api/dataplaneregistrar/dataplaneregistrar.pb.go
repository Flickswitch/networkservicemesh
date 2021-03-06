// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dataplaneregistrar.proto

package dataplaneregistrar

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

// DataplaneRegistrationRequest is sent by the dataplane to NSM
// to advertise itself and inform NSM about the location of the dataplane socket
// and its initially supported parameters.
type DataplaneRegistrationRequest struct {
	DataplaneName        string   `protobuf:"bytes,1,opt,name=dataplane_name,json=dataplaneName,proto3" json:"dataplane_name,omitempty"`
	DataplaneSocket      string   `protobuf:"bytes,2,opt,name=dataplane_socket,json=dataplaneSocket,proto3" json:"dataplane_socket,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataplaneRegistrationRequest) Reset()         { *m = DataplaneRegistrationRequest{} }
func (m *DataplaneRegistrationRequest) String() string { return proto.CompactTextString(m) }
func (*DataplaneRegistrationRequest) ProtoMessage()    {}
func (*DataplaneRegistrationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f4c86488a7f7eab, []int{0}
}

func (m *DataplaneRegistrationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataplaneRegistrationRequest.Unmarshal(m, b)
}
func (m *DataplaneRegistrationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataplaneRegistrationRequest.Marshal(b, m, deterministic)
}
func (m *DataplaneRegistrationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataplaneRegistrationRequest.Merge(m, src)
}
func (m *DataplaneRegistrationRequest) XXX_Size() int {
	return xxx_messageInfo_DataplaneRegistrationRequest.Size(m)
}
func (m *DataplaneRegistrationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DataplaneRegistrationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DataplaneRegistrationRequest proto.InternalMessageInfo

func (m *DataplaneRegistrationRequest) GetDataplaneName() string {
	if m != nil {
		return m.DataplaneName
	}
	return ""
}

func (m *DataplaneRegistrationRequest) GetDataplaneSocket() string {
	if m != nil {
		return m.DataplaneSocket
	}
	return ""
}

type DataplaneRegistrationReply struct {
	Registered           bool     `protobuf:"varint,1,opt,name=registered,proto3" json:"registered,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataplaneRegistrationReply) Reset()         { *m = DataplaneRegistrationReply{} }
func (m *DataplaneRegistrationReply) String() string { return proto.CompactTextString(m) }
func (*DataplaneRegistrationReply) ProtoMessage()    {}
func (*DataplaneRegistrationReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f4c86488a7f7eab, []int{1}
}

func (m *DataplaneRegistrationReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataplaneRegistrationReply.Unmarshal(m, b)
}
func (m *DataplaneRegistrationReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataplaneRegistrationReply.Marshal(b, m, deterministic)
}
func (m *DataplaneRegistrationReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataplaneRegistrationReply.Merge(m, src)
}
func (m *DataplaneRegistrationReply) XXX_Size() int {
	return xxx_messageInfo_DataplaneRegistrationReply.Size(m)
}
func (m *DataplaneRegistrationReply) XXX_DiscardUnknown() {
	xxx_messageInfo_DataplaneRegistrationReply.DiscardUnknown(m)
}

var xxx_messageInfo_DataplaneRegistrationReply proto.InternalMessageInfo

func (m *DataplaneRegistrationReply) GetRegistered() bool {
	if m != nil {
		return m.Registered
	}
	return false
}

// DataplaneUnRegistrationRequest is sent by the dataplane to NSM
// to remove itself from the list of available dataplanes.
type DataplaneUnRegistrationRequest struct {
	DataplaneName        string   `protobuf:"bytes,1,opt,name=dataplane_name,json=dataplaneName,proto3" json:"dataplane_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataplaneUnRegistrationRequest) Reset()         { *m = DataplaneUnRegistrationRequest{} }
func (m *DataplaneUnRegistrationRequest) String() string { return proto.CompactTextString(m) }
func (*DataplaneUnRegistrationRequest) ProtoMessage()    {}
func (*DataplaneUnRegistrationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f4c86488a7f7eab, []int{2}
}

func (m *DataplaneUnRegistrationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataplaneUnRegistrationRequest.Unmarshal(m, b)
}
func (m *DataplaneUnRegistrationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataplaneUnRegistrationRequest.Marshal(b, m, deterministic)
}
func (m *DataplaneUnRegistrationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataplaneUnRegistrationRequest.Merge(m, src)
}
func (m *DataplaneUnRegistrationRequest) XXX_Size() int {
	return xxx_messageInfo_DataplaneUnRegistrationRequest.Size(m)
}
func (m *DataplaneUnRegistrationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DataplaneUnRegistrationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DataplaneUnRegistrationRequest proto.InternalMessageInfo

func (m *DataplaneUnRegistrationRequest) GetDataplaneName() string {
	if m != nil {
		return m.DataplaneName
	}
	return ""
}

type DataplaneUnRegistrationReply struct {
	UnRegistered         bool     `protobuf:"varint,1,opt,name=un_registered,json=unRegistered,proto3" json:"un_registered,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataplaneUnRegistrationReply) Reset()         { *m = DataplaneUnRegistrationReply{} }
func (m *DataplaneUnRegistrationReply) String() string { return proto.CompactTextString(m) }
func (*DataplaneUnRegistrationReply) ProtoMessage()    {}
func (*DataplaneUnRegistrationReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f4c86488a7f7eab, []int{3}
}

func (m *DataplaneUnRegistrationReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataplaneUnRegistrationReply.Unmarshal(m, b)
}
func (m *DataplaneUnRegistrationReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataplaneUnRegistrationReply.Marshal(b, m, deterministic)
}
func (m *DataplaneUnRegistrationReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataplaneUnRegistrationReply.Merge(m, src)
}
func (m *DataplaneUnRegistrationReply) XXX_Size() int {
	return xxx_messageInfo_DataplaneUnRegistrationReply.Size(m)
}
func (m *DataplaneUnRegistrationReply) XXX_DiscardUnknown() {
	xxx_messageInfo_DataplaneUnRegistrationReply.DiscardUnknown(m)
}

var xxx_messageInfo_DataplaneUnRegistrationReply proto.InternalMessageInfo

func (m *DataplaneUnRegistrationReply) GetUnRegistered() bool {
	if m != nil {
		return m.UnRegistered
	}
	return false
}

func init() {
	proto.RegisterType((*DataplaneRegistrationRequest)(nil), "dataplaneregistrar.DataplaneRegistrationRequest")
	proto.RegisterType((*DataplaneRegistrationReply)(nil), "dataplaneregistrar.DataplaneRegistrationReply")
	proto.RegisterType((*DataplaneUnRegistrationRequest)(nil), "dataplaneregistrar.DataplaneUnRegistrationRequest")
	proto.RegisterType((*DataplaneUnRegistrationReply)(nil), "dataplaneregistrar.DataplaneUnRegistrationReply")
}

func init() { proto.RegisterFile("dataplaneregistrar.proto", fileDescriptor_7f4c86488a7f7eab) }

var fileDescriptor_7f4c86488a7f7eab = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x4f, 0x4b, 0xf3, 0x40,
	0x18, 0xc4, 0xd9, 0xf7, 0xf0, 0xa2, 0x0f, 0xd6, 0xca, 0x82, 0x1a, 0x42, 0x09, 0x12, 0x11, 0xea,
	0x65, 0x0d, 0xf1, 0xea, 0x4d, 0x8b, 0x17, 0xf1, 0x10, 0xf1, 0x5c, 0xb6, 0xf6, 0x31, 0x04, 0x93,
	0xdd, 0x75, 0x77, 0x23, 0xe4, 0xe6, 0xc9, 0x0f, 0xe2, 0x37, 0xf3, 0x9b, 0x48, 0xfe, 0x18, 0x4b,
	0x9a, 0x14, 0xea, 0x25, 0x87, 0xc9, 0x6f, 0x76, 0x67, 0x32, 0x01, 0x67, 0xc9, 0x2d, 0x57, 0x29,
	0x17, 0xa8, 0x31, 0x4e, 0x8c, 0xd5, 0x5c, 0x33, 0xa5, 0xa5, 0x95, 0x94, 0xae, 0xbf, 0x71, 0x1d,
	0x65, 0x0b, 0x85, 0xe6, 0x02, 0x33, 0x65, 0x8b, 0xfa, 0x59, 0xd3, 0xbe, 0x82, 0xc9, 0xcd, 0x0f,
	0x1f, 0x35, 0xbc, 0x4d, 0xa4, 0x88, 0xf0, 0x35, 0x47, 0x63, 0xe9, 0x19, 0xec, 0xb7, 0xe7, 0xcd,
	0x05, 0xcf, 0xd0, 0x21, 0x27, 0x64, 0xba, 0x1b, 0x8d, 0x5a, 0xf5, 0x9e, 0x67, 0x48, 0xcf, 0xe1,
	0xe0, 0x17, 0x33, 0xf2, 0xe9, 0x05, 0xad, 0xf3, 0xaf, 0x02, 0xc7, 0xad, 0xfe, 0x50, 0xc9, 0xfe,
	0x15, 0xb8, 0x03, 0x37, 0xaa, 0xb4, 0xa0, 0x1e, 0x40, 0x1d, 0x1b, 0x35, 0x2e, 0xab, 0xbb, 0x76,
	0xa2, 0x15, 0xc5, 0xbf, 0x05, 0xaf, 0x75, 0x3f, 0x8a, 0xbf, 0x27, 0xf6, 0xaf, 0x57, 0x8a, 0x77,
	0x0f, 0x2a, 0x83, 0x9c, 0xc2, 0x28, 0x17, 0xf3, 0xb5, 0x2c, 0x7b, 0x79, 0xc3, 0x96, 0x5a, 0xf8,
	0x45, 0xe0, 0xb0, 0xb7, 0x0c, 0x7d, 0x27, 0x30, 0x69, 0x12, 0xf5, 0x03, 0x01, 0xeb, 0x59, 0x70,
	0xd3, 0x14, 0x2e, 0xdb, 0xc2, 0x51, 0x36, 0x98, 0xc1, 0xb8, 0xb1, 0xde, 0x25, 0x6f, 0x28, 0xd0,
	0x18, 0x7a, 0xc4, 0x62, 0x29, 0xe3, 0x14, 0xeb, 0xf1, 0x17, 0xf9, 0x33, 0x9b, 0x95, 0xff, 0x82,
	0x3b, 0xa0, 0x4f, 0x49, 0x40, 0xc2, 0x4f, 0x02, 0xc7, 0x03, 0x5f, 0x8a, 0x7e, 0x10, 0xf0, 0xba,
	0x2d, 0x3b, 0x48, 0xb8, 0x31, 0x75, 0xef, 0x84, 0x6e, 0xb0, 0x95, 0x47, 0xa5, 0xc5, 0xe2, 0x7f,
	0x15, 0xfc, 0xf2, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x1c, 0xca, 0x68, 0xa3, 0x17, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DataplaneRegistrationClient is the client API for DataplaneRegistration service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DataplaneRegistrationClient interface {
	RequestDataplaneRegistration(ctx context.Context, in *DataplaneRegistrationRequest, opts ...grpc.CallOption) (*DataplaneRegistrationReply, error)
	// RequestLiveness is a stream initiated by NSM to inform the dataplane that NSM is still alive and
	// no re-registration is required. Detection a failure on this "channel" will mean
	// that NSM is gone and the dataplane needs to start re-registration logic.
	RequestLiveness(ctx context.Context, opts ...grpc.CallOption) (DataplaneRegistration_RequestLivenessClient, error)
}

type dataplaneRegistrationClient struct {
	cc *grpc.ClientConn
}

func NewDataplaneRegistrationClient(cc *grpc.ClientConn) DataplaneRegistrationClient {
	return &dataplaneRegistrationClient{cc}
}

func (c *dataplaneRegistrationClient) RequestDataplaneRegistration(ctx context.Context, in *DataplaneRegistrationRequest, opts ...grpc.CallOption) (*DataplaneRegistrationReply, error) {
	out := new(DataplaneRegistrationReply)
	err := c.cc.Invoke(ctx, "/dataplaneregistrar.DataplaneRegistration/RequestDataplaneRegistration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataplaneRegistrationClient) RequestLiveness(ctx context.Context, opts ...grpc.CallOption) (DataplaneRegistration_RequestLivenessClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DataplaneRegistration_serviceDesc.Streams[0], "/dataplaneregistrar.DataplaneRegistration/RequestLiveness", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataplaneRegistrationRequestLivenessClient{stream}
	return x, nil
}

type DataplaneRegistration_RequestLivenessClient interface {
	Send(*empty.Empty) error
	Recv() (*empty.Empty, error)
	grpc.ClientStream
}

type dataplaneRegistrationRequestLivenessClient struct {
	grpc.ClientStream
}

func (x *dataplaneRegistrationRequestLivenessClient) Send(m *empty.Empty) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dataplaneRegistrationRequestLivenessClient) Recv() (*empty.Empty, error) {
	m := new(empty.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DataplaneRegistrationServer is the server API for DataplaneRegistration service.
type DataplaneRegistrationServer interface {
	RequestDataplaneRegistration(context.Context, *DataplaneRegistrationRequest) (*DataplaneRegistrationReply, error)
	// RequestLiveness is a stream initiated by NSM to inform the dataplane that NSM is still alive and
	// no re-registration is required. Detection a failure on this "channel" will mean
	// that NSM is gone and the dataplane needs to start re-registration logic.
	RequestLiveness(DataplaneRegistration_RequestLivenessServer) error
}

// UnimplementedDataplaneRegistrationServer can be embedded to have forward compatible implementations.
type UnimplementedDataplaneRegistrationServer struct {
}

func (*UnimplementedDataplaneRegistrationServer) RequestDataplaneRegistration(ctx context.Context, req *DataplaneRegistrationRequest) (*DataplaneRegistrationReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestDataplaneRegistration not implemented")
}
func (*UnimplementedDataplaneRegistrationServer) RequestLiveness(srv DataplaneRegistration_RequestLivenessServer) error {
	return status.Errorf(codes.Unimplemented, "method RequestLiveness not implemented")
}

func RegisterDataplaneRegistrationServer(s *grpc.Server, srv DataplaneRegistrationServer) {
	s.RegisterService(&_DataplaneRegistration_serviceDesc, srv)
}

func _DataplaneRegistration_RequestDataplaneRegistration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataplaneRegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataplaneRegistrationServer).RequestDataplaneRegistration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dataplaneregistrar.DataplaneRegistration/RequestDataplaneRegistration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataplaneRegistrationServer).RequestDataplaneRegistration(ctx, req.(*DataplaneRegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataplaneRegistration_RequestLiveness_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DataplaneRegistrationServer).RequestLiveness(&dataplaneRegistrationRequestLivenessServer{stream})
}

type DataplaneRegistration_RequestLivenessServer interface {
	Send(*empty.Empty) error
	Recv() (*empty.Empty, error)
	grpc.ServerStream
}

type dataplaneRegistrationRequestLivenessServer struct {
	grpc.ServerStream
}

func (x *dataplaneRegistrationRequestLivenessServer) Send(m *empty.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dataplaneRegistrationRequestLivenessServer) Recv() (*empty.Empty, error) {
	m := new(empty.Empty)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _DataplaneRegistration_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dataplaneregistrar.DataplaneRegistration",
	HandlerType: (*DataplaneRegistrationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestDataplaneRegistration",
			Handler:    _DataplaneRegistration_RequestDataplaneRegistration_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RequestLiveness",
			Handler:       _DataplaneRegistration_RequestLiveness_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "dataplaneregistrar.proto",
}

// DataplaneUnRegistrationClient is the client API for DataplaneUnRegistration service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DataplaneUnRegistrationClient interface {
	RequestDataplaneUnRegistration(ctx context.Context, in *DataplaneUnRegistrationRequest, opts ...grpc.CallOption) (*DataplaneUnRegistrationReply, error)
}

type dataplaneUnRegistrationClient struct {
	cc *grpc.ClientConn
}

func NewDataplaneUnRegistrationClient(cc *grpc.ClientConn) DataplaneUnRegistrationClient {
	return &dataplaneUnRegistrationClient{cc}
}

func (c *dataplaneUnRegistrationClient) RequestDataplaneUnRegistration(ctx context.Context, in *DataplaneUnRegistrationRequest, opts ...grpc.CallOption) (*DataplaneUnRegistrationReply, error) {
	out := new(DataplaneUnRegistrationReply)
	err := c.cc.Invoke(ctx, "/dataplaneregistrar.DataplaneUnRegistration/RequestDataplaneUnRegistration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataplaneUnRegistrationServer is the server API for DataplaneUnRegistration service.
type DataplaneUnRegistrationServer interface {
	RequestDataplaneUnRegistration(context.Context, *DataplaneUnRegistrationRequest) (*DataplaneUnRegistrationReply, error)
}

// UnimplementedDataplaneUnRegistrationServer can be embedded to have forward compatible implementations.
type UnimplementedDataplaneUnRegistrationServer struct {
}

func (*UnimplementedDataplaneUnRegistrationServer) RequestDataplaneUnRegistration(ctx context.Context, req *DataplaneUnRegistrationRequest) (*DataplaneUnRegistrationReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestDataplaneUnRegistration not implemented")
}

func RegisterDataplaneUnRegistrationServer(s *grpc.Server, srv DataplaneUnRegistrationServer) {
	s.RegisterService(&_DataplaneUnRegistration_serviceDesc, srv)
}

func _DataplaneUnRegistration_RequestDataplaneUnRegistration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataplaneUnRegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataplaneUnRegistrationServer).RequestDataplaneUnRegistration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dataplaneregistrar.DataplaneUnRegistration/RequestDataplaneUnRegistration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataplaneUnRegistrationServer).RequestDataplaneUnRegistration(ctx, req.(*DataplaneUnRegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DataplaneUnRegistration_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dataplaneregistrar.DataplaneUnRegistration",
	HandlerType: (*DataplaneUnRegistrationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestDataplaneUnRegistration",
			Handler:    _DataplaneUnRegistration_RequestDataplaneUnRegistration_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dataplaneregistrar.proto",
}
