// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/proto/auth/auth.proto

package auth

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

type LoginReq struct {
	Username             string   `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginReq) Reset()         { *m = LoginReq{} }
func (m *LoginReq) String() string { return proto.CompactTextString(m) }
func (*LoginReq) ProtoMessage()    {}
func (*LoginReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fb64d842494440d, []int{0}
}

func (m *LoginReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginReq.Unmarshal(m, b)
}
func (m *LoginReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginReq.Marshal(b, m, deterministic)
}
func (m *LoginReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginReq.Merge(m, src)
}
func (m *LoginReq) XXX_Size() int {
	return xxx_messageInfo_LoginReq.Size(m)
}
func (m *LoginReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginReq.DiscardUnknown(m)
}

var xxx_messageInfo_LoginReq proto.InternalMessageInfo

func (m *LoginReq) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type Check struct {
	Cookie               string   `protobuf:"bytes,1,opt,name=cookie,proto3" json:"cookie,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Check) Reset()         { *m = Check{} }
func (m *Check) String() string { return proto.CompactTextString(m) }
func (*Check) ProtoMessage()    {}
func (*Check) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fb64d842494440d, []int{1}
}

func (m *Check) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Check.Unmarshal(m, b)
}
func (m *Check) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Check.Marshal(b, m, deterministic)
}
func (m *Check) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Check.Merge(m, src)
}
func (m *Check) XXX_Size() int {
	return xxx_messageInfo_Check.Size(m)
}
func (m *Check) XXX_DiscardUnknown() {
	xxx_messageInfo_Check.DiscardUnknown(m)
}

var xxx_messageInfo_Check proto.InternalMessageInfo

func (m *Check) GetCookie() string {
	if m != nil {
		return m.Cookie
	}
	return ""
}

type User struct {
	ID                   int32    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=Email,proto3" json:"Email,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=Password,proto3" json:"Password,omitempty"`
	Name                 string   `protobuf:"bytes,5,opt,name=Name,proto3" json:"Name,omitempty"`
	Surname              string   `protobuf:"bytes,6,opt,name=Surname,proto3" json:"Surname,omitempty"`
	Description          string   `protobuf:"bytes,7,opt,name=Description,proto3" json:"Description,omitempty"`
	Avatar               string   `protobuf:"bytes,8,opt,name=Avatar,proto3" json:"Avatar,omitempty"`
	Followers            int32    `protobuf:"varint,9,opt,name=Followers,proto3" json:"Followers,omitempty"`
	Following            int32    `protobuf:"varint,10,opt,name=Following,proto3" json:"Following,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fb64d842494440d, []int{2}
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

func (m *User) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetSurname() string {
	if m != nil {
		return m.Surname
	}
	return ""
}

func (m *User) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *User) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *User) GetFollowers() int32 {
	if m != nil {
		return m.Followers
	}
	return 0
}

func (m *User) GetFollowing() int32 {
	if m != nil {
		return m.Following
	}
	return 0
}

type Token struct {
	JwtT                 string   `protobuf:"bytes,1,opt,name=jwtT,proto3" json:"jwtT,omitempty"`
	CsrfT                string   `protobuf:"bytes,2,opt,name=csrfT,proto3" json:"csrfT,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fb64d842494440d, []int{3}
}

func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetJwtT() string {
	if m != nil {
		return m.JwtT
	}
	return ""
}

func (m *Token) GetCsrfT() string {
	if m != nil {
		return m.CsrfT
	}
	return ""
}

type LoginInfo struct {
	Valid                bool     `protobuf:"varint,1,opt,name=Valid,proto3" json:"Valid,omitempty"`
	Info                 *User    `protobuf:"bytes,2,opt,name=Info,proto3" json:"Info,omitempty"`
	Tokens               *Token   `protobuf:"bytes,3,opt,name=Tokens,proto3" json:"Tokens,omitempty"`
	Status               int32    `protobuf:"varint,4,opt,name=Status,proto3" json:"Status,omitempty"`
	Error                string   `protobuf:"bytes,5,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginInfo) Reset()         { *m = LoginInfo{} }
func (m *LoginInfo) String() string { return proto.CompactTextString(m) }
func (*LoginInfo) ProtoMessage()    {}
func (*LoginInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fb64d842494440d, []int{4}
}

func (m *LoginInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginInfo.Unmarshal(m, b)
}
func (m *LoginInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginInfo.Marshal(b, m, deterministic)
}
func (m *LoginInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginInfo.Merge(m, src)
}
func (m *LoginInfo) XXX_Size() int {
	return xxx_messageInfo_LoginInfo.Size(m)
}
func (m *LoginInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LoginInfo proto.InternalMessageInfo

func (m *LoginInfo) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *LoginInfo) GetInfo() *User {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *LoginInfo) GetTokens() *Token {
	if m != nil {
		return m.Tokens
	}
	return nil
}

func (m *LoginInfo) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *LoginInfo) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type UserID struct {
	Valid                bool     `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
	Id                   int32    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserID) Reset()         { *m = UserID{} }
func (m *UserID) String() string { return proto.CompactTextString(m) }
func (*UserID) ProtoMessage()    {}
func (*UserID) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fb64d842494440d, []int{5}
}

func (m *UserID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserID.Unmarshal(m, b)
}
func (m *UserID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserID.Marshal(b, m, deterministic)
}
func (m *UserID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserID.Merge(m, src)
}
func (m *UserID) XXX_Size() int {
	return xxx_messageInfo_UserID.Size(m)
}
func (m *UserID) XXX_DiscardUnknown() {
	xxx_messageInfo_UserID.DiscardUnknown(m)
}

var xxx_messageInfo_UserID proto.InternalMessageInfo

func (m *UserID) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *UserID) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*LoginReq)(nil), "auth.LoginReq")
	proto.RegisterType((*Check)(nil), "auth.Check")
	proto.RegisterType((*User)(nil), "auth.User")
	proto.RegisterType((*Token)(nil), "auth.Token")
	proto.RegisterType((*LoginInfo)(nil), "auth.LoginInfo")
	proto.RegisterType((*UserID)(nil), "auth.UserID")
}

func init() {
	proto.RegisterFile("pkg/proto/auth/auth.proto", fileDescriptor_0fb64d842494440d)
}

var fileDescriptor_0fb64d842494440d = []byte{
	// 420 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x52, 0xd1, 0xae, 0xd2, 0x40,
	0x10, 0x85, 0xda, 0x96, 0x76, 0x6a, 0x30, 0xd9, 0x18, 0xb3, 0x12, 0xa3, 0x64, 0x7d, 0x31, 0x3e,
	0x40, 0xc4, 0x2f, 0x40, 0xd0, 0x84, 0xc4, 0x18, 0x53, 0xd0, 0xf7, 0xb5, 0x2c, 0xb0, 0xb6, 0x74,
	0x71, 0xb7, 0xc0, 0x2f, 0xf8, 0x01, 0x7e, 0xb0, 0xdd, 0xd9, 0xf6, 0x52, 0xee, 0x4b, 0x33, 0xe7,
	0x9c, 0x9d, 0xce, 0xcc, 0x99, 0x81, 0x97, 0xa7, 0x7c, 0x3f, 0x3d, 0x69, 0x55, 0xa9, 0x29, 0x3f,
	0x57, 0x07, 0xfc, 0x4c, 0x10, 0x13, 0xdf, 0xc6, 0xec, 0x13, 0x44, 0x5f, 0xd5, 0x5e, 0x96, 0xa9,
	0xf8, 0x43, 0x46, 0x10, 0xfd, 0x30, 0x42, 0x97, 0xfc, 0x28, 0x68, 0x7f, 0xdc, 0x7f, 0x17, 0xa7,
	0x0f, 0xd8, 0x6a, 0xdf, 0xb9, 0x31, 0x57, 0xa5, 0xb7, 0xd4, 0x73, 0x5a, 0x8b, 0xd9, 0x1b, 0x08,
	0x16, 0x07, 0x91, 0xe5, 0xe4, 0x05, 0x84, 0x99, 0x52, 0xb9, 0x6c, 0xd3, 0x1b, 0xc4, 0xfe, 0x7a,
	0xe0, 0xdb, 0x3f, 0x91, 0x21, 0x78, 0xab, 0x25, 0x8a, 0x41, 0x5a, 0x47, 0x77, 0x15, 0xbd, 0x47,
	0x15, 0x9f, 0x43, 0xf0, 0xf9, 0xc8, 0x65, 0x41, 0x9f, 0xa0, 0xe0, 0xc0, 0x5d, 0x1f, 0xfe, 0x7d,
	0x1f, 0x84, 0x80, 0xff, 0xcd, 0xfe, 0x29, 0x40, 0x1e, 0x63, 0x42, 0x61, 0xb0, 0x3e, 0xbb, 0x02,
	0x21, 0xd2, 0x2d, 0x24, 0x63, 0x48, 0x96, 0xc2, 0x64, 0x5a, 0x9e, 0x2a, 0xa9, 0x4a, 0x3a, 0x40,
	0xb5, 0x4b, 0xd9, 0x71, 0xe6, 0x17, 0x5e, 0x71, 0x4d, 0x23, 0x37, 0x8e, 0x43, 0xe4, 0x15, 0xc4,
	0x5f, 0x54, 0x51, 0xa8, 0xab, 0xd0, 0x86, 0xc6, 0x38, 0xcc, 0x8d, 0xb8, 0xa9, 0xb2, 0xdc, 0x53,
	0xe8, 0xaa, 0x35, 0xc1, 0x3e, 0x40, 0xb0, 0x51, 0xb9, 0x28, 0x6d, 0xb3, 0xbf, 0xaf, 0xd5, 0xa6,
	0x71, 0x0a, 0x63, 0x3b, 0x72, 0x66, 0xf4, 0x6e, 0xd3, 0x78, 0xe1, 0x00, 0xfb, 0xd7, 0x87, 0x18,
	0x77, 0xb4, 0x2a, 0x77, 0xca, 0xbe, 0xf9, 0xc9, 0x0b, 0xb9, 0xc5, 0xc4, 0x28, 0x75, 0x80, 0xbc,
	0x06, 0xdf, 0xaa, 0x98, 0x98, 0xcc, 0x60, 0x82, 0x7b, 0xb6, 0x56, 0xa6, 0xc8, 0x93, 0xb7, 0x10,
	0x62, 0x59, 0x83, 0x6e, 0x26, 0xb3, 0xc4, 0xbd, 0x40, 0x2e, 0x6d, 0x24, 0x3b, 0xef, 0xba, 0xe2,
	0xd5, 0xd9, 0xa0, 0xb3, 0x41, 0xda, 0x20, 0xdc, 0x84, 0xd6, 0x4a, 0x37, 0xc6, 0x3a, 0xc0, 0x26,
	0x10, 0xda, 0x02, 0xf5, 0x16, 0x6b, 0xfd, 0xd2, 0x6d, 0x09, 0x81, 0xdd, 0xb5, 0x74, 0xb7, 0x52,
	0xef, 0x5a, 0x6e, 0x67, 0x02, 0x92, 0x79, 0x5d, 0x73, 0x2d, 0xf4, 0x45, 0x66, 0x82, 0xbc, 0x87,
	0x04, 0x8f, 0x66, 0x81, 0x27, 0x42, 0x9a, 0x86, 0x90, 0x1a, 0x3d, 0xbd, 0xf5, 0xbf, 0x5a, 0xb2,
	0x5e, 0xfd, 0x36, 0x40, 0x03, 0xc8, 0xd0, 0x09, 0xed, 0xc5, 0x8e, 0x9e, 0x75, 0xb0, 0x9d, 0x93,
	0xf5, 0x7e, 0x85, 0x78, 0xdd, 0x1f, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0xc3, 0x01, 0x57, 0xb4,
	0xfa, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthServiceClient interface {
	CheckCookie(ctx context.Context, in *Check, opts ...grpc.CallOption) (*UserID, error)
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginInfo, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) CheckCookie(ctx context.Context, in *Check, opts ...grpc.CallOption) (*UserID, error) {
	out := new(UserID)
	err := c.cc.Invoke(ctx, "/auth.AuthService/CheckCookie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginInfo, error) {
	out := new(LoginInfo)
	err := c.cc.Invoke(ctx, "/auth.AuthService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
type AuthServiceServer interface {
	CheckCookie(context.Context, *Check) (*UserID, error)
	Login(context.Context, *LoginReq) (*LoginInfo, error)
}

// UnimplementedAuthServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (*UnimplementedAuthServiceServer) CheckCookie(ctx context.Context, req *Check) (*UserID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckCookie not implemented")
}
func (*UnimplementedAuthServiceServer) Login(ctx context.Context, req *LoginReq) (*LoginInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}

func RegisterAuthServiceServer(s *grpc.Server, srv AuthServiceServer) {
	s.RegisterService(&_AuthService_serviceDesc, srv)
}

func _AuthService_CheckCookie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Check)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).CheckCookie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/CheckCookie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).CheckCookie(ctx, req.(*Check))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auth.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckCookie",
			Handler:    _AuthService_CheckCookie_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _AuthService_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/proto/auth/auth.proto",
}
