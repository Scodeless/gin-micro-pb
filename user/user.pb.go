// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user/user.proto

package userPb

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

type Gender int32

const (
	Gender_unknown Gender = 0
	Gender_male    Gender = 1
	Gender_female  Gender = 2
)

var Gender_name = map[int32]string{
	0: "unknown",
	1: "male",
	2: "female",
}

var Gender_value = map[string]int32{
	"unknown": 0,
	"male":    1,
	"female":  2,
}

func (x Gender) String() string {
	return proto.EnumName(Gender_name, int32(x))
}

func (Gender) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ed89022014131a74, []int{0}
}

type LoginRequest struct {
	Mobile               string   `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89022014131a74, []int{0}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginResponse struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data                 *Data    `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89022014131a74, []int{1}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *LoginResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *LoginResponse) GetData() *Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type Data struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Gender               Gender   `protobuf:"varint,2,opt,name=gender,proto3,enum=userPb.Gender" json:"gender,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Uid                  string   `protobuf:"bytes,4,opt,name=uid,proto3" json:"uid,omitempty"`
	Mobile               uint64   `protobuf:"varint,5,opt,name=mobile,proto3" json:"mobile,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()    {}
func (*Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89022014131a74, []int{2}
}

func (m *Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Data.Unmarshal(m, b)
}
func (m *Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Data.Marshal(b, m, deterministic)
}
func (m *Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data.Merge(m, src)
}
func (m *Data) XXX_Size() int {
	return xxx_messageInfo_Data.Size(m)
}
func (m *Data) XXX_DiscardUnknown() {
	xxx_messageInfo_Data.DiscardUnknown(m)
}

var xxx_messageInfo_Data proto.InternalMessageInfo

func (m *Data) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Data) GetGender() Gender {
	if m != nil {
		return m.Gender
	}
	return Gender_unknown
}

func (m *Data) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Data) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *Data) GetMobile() uint64 {
	if m != nil {
		return m.Mobile
	}
	return 0
}

type RegisterRequest struct {
	Mobile               string   `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Gender               Gender   `protobuf:"varint,3,opt,name=gender,proto3,enum=userPb.Gender" json:"gender,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89022014131a74, []int{3}
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

func (m *RegisterRequest) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *RegisterRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *RegisterRequest) GetGender() Gender {
	if m != nil {
		return m.Gender
	}
	return Gender_unknown
}

func (m *RegisterRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetUserInfoByMobileRequest struct {
	Mobile               uint64   `protobuf:"varint,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserInfoByMobileRequest) Reset()         { *m = GetUserInfoByMobileRequest{} }
func (m *GetUserInfoByMobileRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserInfoByMobileRequest) ProtoMessage()    {}
func (*GetUserInfoByMobileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89022014131a74, []int{4}
}

func (m *GetUserInfoByMobileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserInfoByMobileRequest.Unmarshal(m, b)
}
func (m *GetUserInfoByMobileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserInfoByMobileRequest.Marshal(b, m, deterministic)
}
func (m *GetUserInfoByMobileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserInfoByMobileRequest.Merge(m, src)
}
func (m *GetUserInfoByMobileRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserInfoByMobileRequest.Size(m)
}
func (m *GetUserInfoByMobileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserInfoByMobileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserInfoByMobileRequest proto.InternalMessageInfo

func (m *GetUserInfoByMobileRequest) GetMobile() uint64 {
	if m != nil {
		return m.Mobile
	}
	return 0
}

type UserInfoResponse struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Gender               Gender   `protobuf:"varint,2,opt,name=gender,proto3,enum=userPb.Gender" json:"gender,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Uid                  string   `protobuf:"bytes,4,opt,name=uid,proto3" json:"uid,omitempty"`
	Phone                uint64   `protobuf:"varint,5,opt,name=phone,proto3" json:"phone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfoResponse) Reset()         { *m = UserInfoResponse{} }
func (m *UserInfoResponse) String() string { return proto.CompactTextString(m) }
func (*UserInfoResponse) ProtoMessage()    {}
func (*UserInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89022014131a74, []int{5}
}

func (m *UserInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfoResponse.Unmarshal(m, b)
}
func (m *UserInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfoResponse.Marshal(b, m, deterministic)
}
func (m *UserInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfoResponse.Merge(m, src)
}
func (m *UserInfoResponse) XXX_Size() int {
	return xxx_messageInfo_UserInfoResponse.Size(m)
}
func (m *UserInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfoResponse proto.InternalMessageInfo

func (m *UserInfoResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserInfoResponse) GetGender() Gender {
	if m != nil {
		return m.Gender
	}
	return Gender_unknown
}

func (m *UserInfoResponse) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserInfoResponse) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *UserInfoResponse) GetPhone() uint64 {
	if m != nil {
		return m.Phone
	}
	return 0
}

func init() {
	proto.RegisterEnum("userPb.Gender", Gender_name, Gender_value)
	proto.RegisterType((*LoginRequest)(nil), "userPb.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "userPb.LoginResponse")
	proto.RegisterType((*Data)(nil), "userPb.Data")
	proto.RegisterType((*RegisterRequest)(nil), "userPb.RegisterRequest")
	proto.RegisterType((*GetUserInfoByMobileRequest)(nil), "userPb.GetUserInfoByMobileRequest")
	proto.RegisterType((*UserInfoResponse)(nil), "userPb.UserInfoResponse")
}

func init() { proto.RegisterFile("user/user.proto", fileDescriptor_ed89022014131a74) }

var fileDescriptor_ed89022014131a74 = []byte{
	// 433 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x9d, 0x6d, 0x1a, 0xb7, 0xaf, 0x75, 0xb7, 0x3c, 0xab, 0x86, 0x1e, 0x24, 0xe4, 0x20,
	0x45, 0xd9, 0x16, 0xaa, 0x07, 0x4f, 0x1e, 0x8a, 0xb0, 0x08, 0x0a, 0x92, 0x45, 0x04, 0x6f, 0x93,
	0xe6, 0x6d, 0x76, 0xb0, 0x99, 0x89, 0x33, 0x89, 0x8b, 0x47, 0x3d, 0xfb, 0x11, 0xfd, 0x30, 0x32,
	0x93, 0x8c, 0x0d, 0x65, 0xd7, 0x8b, 0x78, 0x09, 0xef, 0x3f, 0x33, 0xef, 0x3f, 0xbf, 0xf7, 0xde,
	0x04, 0x4e, 0x1b, 0x43, 0x7a, 0x65, 0x3f, 0xcb, 0x4a, 0xab, 0x5a, 0x61, 0x68, 0xe3, 0xf7, 0x59,
	0xb2, 0x81, 0xc9, 0x5b, 0x55, 0x08, 0x99, 0xd2, 0x97, 0x86, 0x4c, 0x8d, 0x0f, 0x21, 0x2c, 0x55,
	0x26, 0x76, 0x14, 0xb1, 0x98, 0x2d, 0x46, 0x69, 0xa7, 0x70, 0x0e, 0xc7, 0x15, 0x37, 0xe6, 0x5a,
	0xe9, 0x3c, 0x3a, 0x72, 0x3b, 0x7f, 0x74, 0xf2, 0x11, 0xee, 0x75, 0x1e, 0xa6, 0x52, 0xd2, 0x10,
	0x22, 0x04, 0x5b, 0x95, 0xb7, 0x16, 0xc3, 0xd4, 0xc5, 0x38, 0x85, 0x41, 0x69, 0x8a, 0x2e, 0xd7,
	0x86, 0x18, 0x43, 0x90, 0xf3, 0x9a, 0x47, 0x83, 0x98, 0x2d, 0xc6, 0xeb, 0xc9, 0xb2, 0x25, 0x5a,
	0xbe, 0xe6, 0x35, 0x4f, 0xdd, 0x4e, 0xf2, 0x83, 0x41, 0x60, 0xa5, 0x35, 0x94, 0xbc, 0xf4, 0x4c,
	0x2e, 0xc6, 0x27, 0x10, 0x16, 0x24, 0x73, 0xd2, 0xce, 0xf3, 0x64, 0x7d, 0xe2, 0x0d, 0xce, 0xdd,
	0x6a, 0xda, 0xed, 0xe2, 0x0c, 0x86, 0x54, 0x72, 0xb1, 0x73, 0xf7, 0x8c, 0xd2, 0x56, 0x58, 0x9c,
	0x46, 0xe4, 0x51, 0xd0, 0xe2, 0x34, 0x22, 0xef, 0x55, 0x3e, 0x8c, 0xd9, 0x22, 0xf0, 0x95, 0x27,
	0xdf, 0x19, 0x9c, 0xa6, 0x54, 0x08, 0x53, 0x93, 0xfe, 0x87, 0x2e, 0xf5, 0x78, 0x07, 0x7f, 0xe5,
	0xf5, 0xb5, 0x06, 0xfb, 0x5a, 0x93, 0x17, 0x30, 0x3f, 0xa7, 0xfa, 0x83, 0x21, 0xfd, 0x46, 0x5e,
	0xaa, 0xcd, 0xb7, 0x77, 0xee, 0xba, 0x9b, 0x69, 0xf6, 0xe4, 0x3f, 0x19, 0x4c, 0x7d, 0x4e, 0x7f,
	0x36, 0xff, 0xbd, 0x95, 0x33, 0x18, 0x56, 0x57, 0x4a, 0xfa, 0x4e, 0xb6, 0xe2, 0xe9, 0x33, 0x08,
	0x5b, 0x3f, 0x1c, 0xc3, 0xdd, 0x46, 0x7e, 0x96, 0xea, 0x5a, 0x4e, 0xef, 0xe0, 0x31, 0x04, 0x25,
	0xdf, 0xd1, 0x94, 0x21, 0x40, 0x78, 0x49, 0x2e, 0x3e, 0x5a, 0xff, 0x62, 0x30, 0xb6, 0xec, 0x17,
	0xa4, 0xbf, 0x8a, 0x2d, 0xe1, 0x4b, 0x18, 0x59, 0xe9, 0xde, 0x19, 0xce, 0x3c, 0x5f, 0xff, 0xe9,
	0xce, 0x1f, 0x1c, 0xac, 0x76, 0x05, 0xbf, 0x82, 0x89, 0xcd, 0xf4, 0x23, 0xc4, 0x47, 0xfe, 0xd8,
	0xc1, 0x50, 0x6f, 0xcb, 0xbf, 0x80, 0xfb, 0x37, 0xf4, 0x1e, 0x93, 0x7d, 0x8f, 0x6e, 0x1b, 0xcc,
	0x3c, 0xf2, 0x67, 0x0e, 0xa7, 0xb0, 0x89, 0x3f, 0x3d, 0x2e, 0x44, 0x7d, 0xd5, 0x64, 0xcb, 0xad,
	0x2a, 0x57, 0x85, 0x90, 0x67, 0xa5, 0xd8, 0x6a, 0x75, 0x56, 0x65, 0xab, 0x36, 0x25, 0x0b, 0xdd,
	0x7f, 0xfa, 0xfc, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5d, 0x0e, 0x57, 0x9e, 0xba, 0x03, 0x00,
	0x00,
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
	//user login
	UserLogin(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	//user register
	UserRegister(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	//
	GetUserInfoByMobile(ctx context.Context, in *GetUserInfoByMobileRequest, opts ...grpc.CallOption) (*UserInfoResponse, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) UserLogin(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/userPb.UserService/UserLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserRegister(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/userPb.UserService/UserRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserInfoByMobile(ctx context.Context, in *GetUserInfoByMobileRequest, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	out := new(UserInfoResponse)
	err := c.cc.Invoke(ctx, "/userPb.UserService/GetUserInfoByMobile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	//user login
	UserLogin(context.Context, *LoginRequest) (*LoginResponse, error)
	//user register
	UserRegister(context.Context, *RegisterRequest) (*LoginResponse, error)
	//
	GetUserInfoByMobile(context.Context, *GetUserInfoByMobileRequest) (*UserInfoResponse, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) UserLogin(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogin not implemented")
}
func (*UnimplementedUserServiceServer) UserRegister(ctx context.Context, req *RegisterRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserRegister not implemented")
}
func (*UnimplementedUserServiceServer) GetUserInfoByMobile(ctx context.Context, req *GetUserInfoByMobileRequest) (*UserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfoByMobile not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_UserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userPb.UserService/UserLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserLogin(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userPb.UserService/UserRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserRegister(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserInfoByMobile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoByMobileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserInfoByMobile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userPb.UserService/GetUserInfoByMobile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserInfoByMobile(ctx, req.(*GetUserInfoByMobileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "userPb.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserLogin",
			Handler:    _UserService_UserLogin_Handler,
		},
		{
			MethodName: "UserRegister",
			Handler:    _UserService_UserRegister_Handler,
		},
		{
			MethodName: "GetUserInfoByMobile",
			Handler:    _UserService_GetUserInfoByMobile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/user.proto",
}
