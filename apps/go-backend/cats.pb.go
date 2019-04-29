// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cats.proto

package cats

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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_38f0b8057a110604, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type Cat struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Author               string   `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cat) Reset()         { *m = Cat{} }
func (m *Cat) String() string { return proto.CompactTextString(m) }
func (*Cat) ProtoMessage()    {}
func (*Cat) Descriptor() ([]byte, []int) {
	return fileDescriptor_38f0b8057a110604, []int{1}
}

func (m *Cat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cat.Unmarshal(m, b)
}
func (m *Cat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cat.Marshal(b, m, deterministic)
}
func (m *Cat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cat.Merge(m, src)
}
func (m *Cat) XXX_Size() int {
	return xxx_messageInfo_Cat.Size(m)
}
func (m *Cat) XXX_DiscardUnknown() {
	xxx_messageInfo_Cat.DiscardUnknown(m)
}

var xxx_messageInfo_Cat proto.InternalMessageInfo

func (m *Cat) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Cat) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Cat) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

type CatList struct {
	Cats                 []*Cat   `protobuf:"bytes,1,rep,name=cats,proto3" json:"cats,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CatList) Reset()         { *m = CatList{} }
func (m *CatList) String() string { return proto.CompactTextString(m) }
func (*CatList) ProtoMessage()    {}
func (*CatList) Descriptor() ([]byte, []int) {
	return fileDescriptor_38f0b8057a110604, []int{2}
}

func (m *CatList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CatList.Unmarshal(m, b)
}
func (m *CatList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CatList.Marshal(b, m, deterministic)
}
func (m *CatList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CatList.Merge(m, src)
}
func (m *CatList) XXX_Size() int {
	return xxx_messageInfo_CatList.Size(m)
}
func (m *CatList) XXX_DiscardUnknown() {
	xxx_messageInfo_CatList.DiscardUnknown(m)
}

var xxx_messageInfo_CatList proto.InternalMessageInfo

func (m *CatList) GetCats() []*Cat {
	if m != nil {
		return m.Cats
	}
	return nil
}

type CatIdRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CatIdRequest) Reset()         { *m = CatIdRequest{} }
func (m *CatIdRequest) String() string { return proto.CompactTextString(m) }
func (*CatIdRequest) ProtoMessage()    {}
func (*CatIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_38f0b8057a110604, []int{3}
}

func (m *CatIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CatIdRequest.Unmarshal(m, b)
}
func (m *CatIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CatIdRequest.Marshal(b, m, deterministic)
}
func (m *CatIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CatIdRequest.Merge(m, src)
}
func (m *CatIdRequest) XXX_Size() int {
	return xxx_messageInfo_CatIdRequest.Size(m)
}
func (m *CatIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CatIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CatIdRequest proto.InternalMessageInfo

func (m *CatIdRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*Empty)(nil), "cats.Empty")
	proto.RegisterType((*Cat)(nil), "cats.Cat")
	proto.RegisterType((*CatList)(nil), "cats.CatList")
	proto.RegisterType((*CatIdRequest)(nil), "cats.CatIdRequest")
}

func init() { proto.RegisterFile("cats.proto", fileDescriptor_38f0b8057a110604) }

var fileDescriptor_38f0b8057a110604 = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xcd, 0x4a, 0xc3, 0x40,
	0x14, 0x85, 0x33, 0x49, 0x93, 0xd2, 0x53, 0x75, 0x71, 0x11, 0x09, 0x01, 0x25, 0x8c, 0x22, 0x01,
	0xa1, 0x48, 0x7d, 0x84, 0x51, 0xa4, 0xe0, 0x2a, 0x2e, 0x5c, 0x8f, 0xc9, 0x85, 0x0e, 0x54, 0x53,
	0x93, 0x5b, 0xc1, 0xb7, 0xf4, 0x91, 0xa4, 0x13, 0x49, 0xeb, 0xcf, 0x6e, 0xee, 0xcc, 0x3d, 0xdf,
	0x77, 0x18, 0xa0, 0xb2, 0xd2, 0xcd, 0xd6, 0x6d, 0x23, 0x0d, 0x8d, 0xb6, 0x67, 0x3d, 0x46, 0x7c,
	0xf7, 0xb2, 0x96, 0x0f, 0x6d, 0x10, 0x19, 0x2b, 0x74, 0x84, 0xd0, 0xd5, 0xa9, 0xca, 0x55, 0x11,
	0x97, 0xa1, 0xab, 0xe9, 0x18, 0xb1, 0x38, 0x59, 0x71, 0x1a, 0xe6, 0xaa, 0x98, 0x94, 0xfd, 0x40,
	0x27, 0x48, 0xec, 0x46, 0x96, 0x4d, 0x9b, 0x46, 0xfe, 0xfa, 0x7b, 0xd2, 0x05, 0xc6, 0xc6, 0xca,
	0x83, 0xeb, 0x84, 0x4e, 0xe1, 0x05, 0xa9, 0xca, 0xa3, 0x62, 0x3a, 0x9f, 0xcc, 0xbc, 0xd9, 0x58,
	0x29, 0x7b, 0xef, 0x19, 0x0e, 0x8c, 0x95, 0x45, 0x5d, 0xf2, 0xdb, 0x86, 0xbb, 0x3f, 0xde, 0xf9,
	0xa7, 0x02, 0x8c, 0x95, 0x47, 0x6e, 0xdf, 0x5d, 0xc5, 0x74, 0x81, 0x91, 0xa7, 0x4e, 0x7b, 0x8e,
	0xaf, 0x9c, 0x1d, 0x0e, 0xd0, 0xed, 0x9b, 0x0e, 0x48, 0x23, 0x59, 0xbc, 0x76, 0xdc, 0x0a, 0xed,
	0x7c, 0xd9, 0x7e, 0x44, 0x07, 0x74, 0x89, 0xe8, 0x9e, 0x85, 0x68, 0x58, 0x18, 0x3a, 0x64, 0xbb,
	0x90, 0x0e, 0xe8, 0x0a, 0xc9, 0x2d, 0xaf, 0x58, 0xf8, 0xdf, 0xd5, 0x5f, 0xd0, 0x73, 0xc4, 0x4f,
	0x56, 0xaa, 0xe5, 0xcf, 0x7e, 0xfb, 0xbc, 0x6b, 0xf5, 0x9c, 0xf8, 0x7f, 0xbf, 0xf9, 0x0a, 0x00,
	0x00, 0xff, 0xff, 0x21, 0xa0, 0x96, 0x93, 0x85, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CatServiceClient is the client API for CatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CatServiceClient interface {
	List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*CatList, error)
	Insert(ctx context.Context, in *Cat, opts ...grpc.CallOption) (*Empty, error)
	Get(ctx context.Context, in *CatIdRequest, opts ...grpc.CallOption) (*Cat, error)
	Delete(ctx context.Context, in *CatIdRequest, opts ...grpc.CallOption) (*Empty, error)
	Watch(ctx context.Context, in *Empty, opts ...grpc.CallOption) (CatService_WatchClient, error)
}

type catServiceClient struct {
	cc *grpc.ClientConn
}

func NewCatServiceClient(cc *grpc.ClientConn) CatServiceClient {
	return &catServiceClient{cc}
}

func (c *catServiceClient) List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*CatList, error) {
	out := new(CatList)
	err := c.cc.Invoke(ctx, "/cats.CatService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catServiceClient) Insert(ctx context.Context, in *Cat, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/cats.CatService/Insert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catServiceClient) Get(ctx context.Context, in *CatIdRequest, opts ...grpc.CallOption) (*Cat, error) {
	out := new(Cat)
	err := c.cc.Invoke(ctx, "/cats.CatService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catServiceClient) Delete(ctx context.Context, in *CatIdRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/cats.CatService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catServiceClient) Watch(ctx context.Context, in *Empty, opts ...grpc.CallOption) (CatService_WatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CatService_serviceDesc.Streams[0], "/cats.CatService/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &catServiceWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CatService_WatchClient interface {
	Recv() (*Cat, error)
	grpc.ClientStream
}

type catServiceWatchClient struct {
	grpc.ClientStream
}

func (x *catServiceWatchClient) Recv() (*Cat, error) {
	m := new(Cat)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CatServiceServer is the server API for CatService service.
type CatServiceServer interface {
	List(context.Context, *Empty) (*CatList, error)
	Insert(context.Context, *Cat) (*Empty, error)
	Get(context.Context, *CatIdRequest) (*Cat, error)
	Delete(context.Context, *CatIdRequest) (*Empty, error)
	Watch(*Empty, CatService_WatchServer) error
}

// UnimplementedCatServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCatServiceServer struct {
}

func (*UnimplementedCatServiceServer) List(ctx context.Context, req *Empty) (*CatList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedCatServiceServer) Insert(ctx context.Context, req *Cat) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Insert not implemented")
}
func (*UnimplementedCatServiceServer) Get(ctx context.Context, req *CatIdRequest) (*Cat, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedCatServiceServer) Delete(ctx context.Context, req *CatIdRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedCatServiceServer) Watch(req *Empty, srv CatService_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "method Watch not implemented")
}

func RegisterCatServiceServer(s *grpc.Server, srv CatServiceServer) {
	s.RegisterService(&_CatService_serviceDesc, srv)
}

func _CatService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cats.CatService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatServiceServer).List(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatService_Insert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Cat)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatServiceServer).Insert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cats.CatService/Insert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatServiceServer).Insert(ctx, req.(*Cat))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CatIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cats.CatService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatServiceServer).Get(ctx, req.(*CatIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CatIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cats.CatService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatServiceServer).Delete(ctx, req.(*CatIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatService_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CatServiceServer).Watch(m, &catServiceWatchServer{stream})
}

type CatService_WatchServer interface {
	Send(*Cat) error
	grpc.ServerStream
}

type catServiceWatchServer struct {
	grpc.ServerStream
}

func (x *catServiceWatchServer) Send(m *Cat) error {
	return x.ServerStream.SendMsg(m)
}

var _CatService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cats.CatService",
	HandlerType: (*CatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _CatService_List_Handler,
		},
		{
			MethodName: "Insert",
			Handler:    _CatService_Insert_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _CatService_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CatService_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _CatService_Watch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cats.proto",
}