// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user_services.proto

package pb

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

type User struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	BirthDate            uint64   `protobuf:"varint,3,opt,name=birth_date,json=birthDate,proto3" json:"birth_date,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9e3ed7bbdd7919b, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetBirthDate() uint64 {
	if m != nil {
		return m.BirthDate
	}
	return 0
}

type UserResquest struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserResquest) Reset()         { *m = UserResquest{} }
func (m *UserResquest) String() string { return proto.CompactTextString(m) }
func (*UserResquest) ProtoMessage()    {}
func (*UserResquest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9e3ed7bbdd7919b, []int{1}
}

func (m *UserResquest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResquest.Unmarshal(m, b)
}
func (m *UserResquest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResquest.Marshal(b, m, deterministic)
}
func (m *UserResquest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResquest.Merge(m, src)
}
func (m *UserResquest) XXX_Size() int {
	return xxx_messageInfo_UserResquest.Size(m)
}
func (m *UserResquest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResquest.DiscardUnknown(m)
}

var xxx_messageInfo_UserResquest proto.InternalMessageInfo

func (m *UserResquest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type UserResponse struct {
	Age                  uint64   `protobuf:"varint,1,opt,name=age,proto3" json:"age,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserResponse) Reset()         { *m = UserResponse{} }
func (m *UserResponse) String() string { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()    {}
func (*UserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9e3ed7bbdd7919b, []int{2}
}

func (m *UserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResponse.Unmarshal(m, b)
}
func (m *UserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResponse.Marshal(b, m, deterministic)
}
func (m *UserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResponse.Merge(m, src)
}
func (m *UserResponse) XXX_Size() int {
	return xxx_messageInfo_UserResponse.Size(m)
}
func (m *UserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserResponse proto.InternalMessageInfo

func (m *UserResponse) GetAge() uint64 {
	if m != nil {
		return m.Age
	}
	return 0
}

func init() {
	proto.RegisterType((*User)(nil), "golang_grpc.User")
	proto.RegisterType((*UserResquest)(nil), "golang_grpc.UserResquest")
	proto.RegisterType((*UserResponse)(nil), "golang_grpc.UserResponse")
}

func init() { proto.RegisterFile("user_services.proto", fileDescriptor_d9e3ed7bbdd7919b) }

var fileDescriptor_d9e3ed7bbdd7919b = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x50, 0x31, 0x4b, 0xc5, 0x30,
	0x18, 0xa4, 0xef, 0xc5, 0x07, 0xef, 0xab, 0x83, 0x7e, 0x3a, 0x54, 0x41, 0x28, 0x05, 0xa1, 0x53,
	0x87, 0x8a, 0xab, 0x83, 0x88, 0x9b, 0x08, 0x11, 0x17, 0x97, 0x92, 0xd6, 0x8f, 0x58, 0x68, 0x93,
	0x98, 0xa4, 0xfe, 0x7e, 0xc9, 0x57, 0x11, 0xc1, 0xb7, 0x5d, 0xee, 0x2e, 0x77, 0x97, 0xc0, 0xd9,
	0x12, 0xc8, 0x77, 0x81, 0xfc, 0xd7, 0x38, 0x50, 0x68, 0x9c, 0xb7, 0xd1, 0x62, 0xae, 0xed, 0xa4,
	0x8c, 0xee, 0xb4, 0x77, 0x43, 0xf5, 0x0c, 0xe2, 0x35, 0x90, 0x47, 0x04, 0x61, 0xd4, 0x4c, 0x45,
	0x56, 0x66, 0xf5, 0x5e, 0x32, 0xc6, 0x73, 0x38, 0xa2, 0x59, 0x8d, 0x53, 0xb1, 0x61, 0x72, 0x3d,
	0xe0, 0x15, 0x40, 0x3f, 0xfa, 0xf8, 0xd1, 0xbd, 0xab, 0x48, 0xc5, 0xb6, 0xcc, 0x6a, 0x21, 0xf7,
	0xcc, 0x3c, 0xa8, 0x48, 0xd5, 0x2d, 0x1c, 0xa7, 0x40, 0x49, 0xe1, 0x73, 0xa1, 0x10, 0xf1, 0x1a,
	0x44, 0x1a, 0xc1, 0xc1, 0x79, 0x7b, 0xda, 0xfc, 0x29, 0x6f, 0xd8, 0xc8, 0x72, 0x55, 0xfe, 0x5e,
	0x73, 0xd6, 0x04, 0xc2, 0x13, 0xd8, 0x2a, 0xbd, 0xce, 0x11, 0x32, 0xc1, 0xf6, 0x09, 0xf2, 0xe4,
	0x78, 0x59, 0x1f, 0x83, 0x77, 0xb0, 0x7b, 0x34, 0x3c, 0xfd, 0xe2, 0x7f, 0xe6, 0x4f, 0xf9, 0xe5,
	0x41, 0x89, 0x0b, 0xee, 0xc5, 0xdb, 0xc6, 0xf5, 0xfd, 0x8e, 0xbf, 0xe4, 0xe6, 0x3b, 0x00, 0x00,
	0xff, 0xff, 0x4a, 0xfa, 0x7b, 0xc1, 0x29, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	FnUser(ctx context.Context, in *UserResquest, opts ...grpc.CallOption) (*UserResponse, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) FnUser(ctx context.Context, in *UserResquest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/golang_grpc.UserService/FnUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	FnUser(context.Context, *UserResquest) (*UserResponse, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) FnUser(ctx context.Context, req *UserResquest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FnUser not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_FnUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserResquest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).FnUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/golang_grpc.UserService/FnUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).FnUser(ctx, req.(*UserResquest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "golang_grpc.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FnUser",
			Handler:    _UserService_FnUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_services.proto",
}
