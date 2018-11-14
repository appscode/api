// Code generated by protoc-gen-go. DO NOT EDIT.
// source: authentication.proto

/*
Package v1alpha1 is a generated protocol buffer package.

It is generated from these files:
	authentication.proto
	conduit.proto
	project.proto

It has these top-level messages:
	LoginRequest
	LoginResponse
	CSRFTokenResponse
	ConduitWhoAmIResponse
	ConduitUsersResponse
	ConduitUser
	ProjectListRequest
	ProjectListResponse
	ProjectMemberListRequest
	ProjectMemberListResponse
	Project
	Member
*/
package v1alpha1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/appscodeapis/appscode/api"
import appscode_dtypes "appscode.com/api/dtypes"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type LoginRequest struct {
	Username   string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Password   string `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
	Token      string `protobuf:"bytes,4,opt,name=token" json:"token,omitempty"`
	IssueToken bool   `protobuf:"varint,5,opt,name=issue_token,json=issueToken" json:"issue_token,omitempty"`
}

func (m *LoginRequest) Reset()                    { *m = LoginRequest{} }
func (m *LoginRequest) String() string            { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()               {}
func (*LoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *LoginRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *LoginRequest) GetIssueToken() bool {
	if m != nil {
		return m.IssueToken
	}
	return false
}

type LoginResponse struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *LoginResponse) Reset()                    { *m = LoginResponse{} }
func (m *LoginResponse) String() string            { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()               {}
func (*LoginResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LoginResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type CSRFTokenResponse struct {
	CsrfToken string `protobuf:"bytes,1,opt,name=csrf_token,json=csrfToken" json:"csrf_token,omitempty"`
}

func (m *CSRFTokenResponse) Reset()                    { *m = CSRFTokenResponse{} }
func (m *CSRFTokenResponse) String() string            { return proto.CompactTextString(m) }
func (*CSRFTokenResponse) ProtoMessage()               {}
func (*CSRFTokenResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CSRFTokenResponse) GetCsrfToken() string {
	if m != nil {
		return m.CsrfToken
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "appscode.auth.v1alpha1.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "appscode.auth.v1alpha1.LoginResponse")
	proto.RegisterType((*CSRFTokenResponse)(nil), "appscode.auth.v1alpha1.CSRFTokenResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Authentication service

type AuthenticationClient interface {
	// This rpc is used to check a valid user from other applications.
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Logout(ctx context.Context, in *appscode_dtypes.VoidRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
	CSRFToken(ctx context.Context, in *appscode_dtypes.VoidRequest, opts ...grpc.CallOption) (*CSRFTokenResponse, error)
}

type authenticationClient struct {
	cc *grpc.ClientConn
}

func NewAuthenticationClient(cc *grpc.ClientConn) AuthenticationClient {
	return &authenticationClient{cc}
}

func (c *authenticationClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := grpc.Invoke(ctx, "/appscode.auth.v1alpha1.Authentication/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) Logout(ctx context.Context, in *appscode_dtypes.VoidRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.auth.v1alpha1.Authentication/Logout", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) CSRFToken(ctx context.Context, in *appscode_dtypes.VoidRequest, opts ...grpc.CallOption) (*CSRFTokenResponse, error) {
	out := new(CSRFTokenResponse)
	err := grpc.Invoke(ctx, "/appscode.auth.v1alpha1.Authentication/CSRFToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Authentication service

type AuthenticationServer interface {
	// This rpc is used to check a valid user from other applications.
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Logout(context.Context, *appscode_dtypes.VoidRequest) (*appscode_dtypes.VoidResponse, error)
	CSRFToken(context.Context, *appscode_dtypes.VoidRequest) (*CSRFTokenResponse, error)
}

func RegisterAuthenticationServer(s *grpc.Server, srv AuthenticationServer) {
	s.RegisterService(&_Authentication_serviceDesc, srv)
}

func _Authentication_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.auth.v1alpha1.Authentication/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(appscode_dtypes.VoidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.auth.v1alpha1.Authentication/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).Logout(ctx, req.(*appscode_dtypes.VoidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_CSRFToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(appscode_dtypes.VoidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).CSRFToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.auth.v1alpha1.Authentication/CSRFToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).CSRFToken(ctx, req.(*appscode_dtypes.VoidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Authentication_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.auth.v1alpha1.Authentication",
	HandlerType: (*AuthenticationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Authentication_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _Authentication_Logout_Handler,
		},
		{
			MethodName: "CSRFToken",
			Handler:    _Authentication_CSRFToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authentication.proto",
}

func init() { proto.RegisterFile("authentication.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 422 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcd, 0xce, 0xd2, 0x40,
	0x14, 0xcd, 0x80, 0x90, 0x32, 0xfe, 0x24, 0x8e, 0xc4, 0x34, 0x0d, 0x20, 0xa9, 0x12, 0xd1, 0x40,
	0x1b, 0x70, 0xa3, 0xec, 0xc0, 0xc4, 0x15, 0x0b, 0x52, 0x8d, 0x0b, 0x37, 0x64, 0x6c, 0xc7, 0x52,
	0x85, 0xb9, 0x63, 0x67, 0xaa, 0x71, 0x61, 0x62, 0x78, 0x00, 0x63, 0xe2, 0xb3, 0x18, 0x1f, 0xc4,
	0x57, 0xf0, 0x41, 0xbe, 0x74, 0x4a, 0x4b, 0x9b, 0xef, 0xe3, 0x63, 0x33, 0xc9, 0xbd, 0xe7, 0x9c,
	0x7b, 0x4f, 0xee, 0x19, 0xdc, 0xa6, 0x89, 0xda, 0x30, 0xae, 0x22, 0x9f, 0xaa, 0x08, 0xb8, 0x23,
	0x62, 0x50, 0x40, 0xee, 0x53, 0x21, 0xa4, 0x0f, 0x01, 0x73, 0x52, 0xd8, 0xf9, 0x32, 0xa1, 0x5b,
	0xb1, 0xa1, 0x13, 0xab, 0x13, 0x02, 0x84, 0x5b, 0xe6, 0x52, 0x11, 0xb9, 0x94, 0x73, 0x50, 0x5a,
	0x24, 0x33, 0x95, 0xd5, 0xcb, 0x55, 0x27, 0xf0, 0x87, 0xc5, 0x54, 0x1f, 0x76, 0x9a, 0x13, 0xa8,
	0x6f, 0x82, 0x49, 0x57, 0xbf, 0x19, 0xc9, 0xfe, 0x8e, 0x6f, 0x2d, 0x21, 0x8c, 0xb8, 0xc7, 0x3e,
	0x27, 0x4c, 0x2a, 0x62, 0x61, 0x23, 0x91, 0x2c, 0xe6, 0x74, 0xc7, 0xcc, 0x5a, 0x1f, 0x0d, 0x5b,
	0x5e, 0x51, 0xa7, 0x98, 0xa0, 0x52, 0x7e, 0x85, 0x38, 0x30, 0xeb, 0x19, 0x96, 0xd7, 0xa4, 0x8d,
	0x1b, 0x0a, 0x3e, 0x31, 0x6e, 0xde, 0xd0, 0x40, 0x56, 0x90, 0x07, 0xf8, 0x66, 0x24, 0x65, 0xc2,
	0xd6, 0x19, 0xd6, 0xe8, 0xa3, 0xa1, 0xe1, 0x61, 0xdd, 0x7a, 0x93, 0x76, 0xec, 0x01, 0xbe, 0x7d,
	0x58, 0x2f, 0x05, 0x70, 0xc9, 0x8e, 0x73, 0x50, 0x69, 0x8e, 0x3d, 0xc5, 0x77, 0x5f, 0xbe, 0xf6,
	0x5e, 0x69, 0x4d, 0x41, 0xed, 0x62, 0xec, 0xcb, 0xf8, 0xc3, 0xba, 0xcc, 0x6f, 0xa5, 0x1d, 0x4d,
	0x9b, 0xfe, 0xad, 0xe3, 0x3b, 0xf3, 0xca, 0xb5, 0xc9, 0x4f, 0x84, 0x1b, 0x7a, 0x1d, 0x79, 0xe4,
	0x5c, 0x7d, 0x72, 0xa7, 0x7c, 0x0c, 0x6b, 0x70, 0x86, 0x95, 0x19, 0xb1, 0x9f, 0xef, 0xff, 0x98,
	0x35, 0x03, 0xed, 0xff, 0xfd, 0xff, 0x5d, 0x1b, 0xd9, 0x8f, 0xdd, 0x75, 0x35, 0x98, 0x44, 0x6d,
	0xdc, 0x5c, 0xe9, 0x6e, 0x53, 0xa5, 0xfb, 0x51, 0x02, 0x9f, 0xa1, 0xa7, 0xe4, 0x07, 0xc2, 0xcd,
	0x25, 0x84, 0x90, 0x28, 0xd2, 0x39, 0xee, 0xca, 0x62, 0x72, 0xde, 0x42, 0x14, 0xe4, 0x4e, 0xba,
	0x27, 0xd0, 0x83, 0x83, 0x17, 0x25, 0x07, 0x63, 0x7b, 0x78, 0xce, 0x01, 0x24, 0xaa, 0xb0, 0xf0,
	0x0b, 0xe1, 0x56, 0x71, 0xdb, 0x33, 0x2e, 0x9e, 0x9c, 0xba, 0xc7, 0xa5, 0x70, 0xec, 0x59, 0xc9,
	0x91, 0x43, 0x46, 0xd7, 0x3a, 0x4a, 0x33, 0x1b, 0xeb, 0x14, 0xb5, 0xab, 0xc5, 0x1c, 0xf7, 0x7c,
	0xd8, 0x95, 0x76, 0x89, 0xa8, 0xba, 0x6f, 0x71, 0xaf, 0x1a, 0xec, 0x2a, 0xfd, 0xca, 0x2b, 0xf4,
	0xce, 0xc8, 0x09, 0xef, 0x9b, 0xfa, 0x77, 0x3f, 0xbb, 0x08, 0x00, 0x00, 0xff, 0xff, 0x85, 0xef,
	0xdb, 0x42, 0x70, 0x03, 0x00, 0x00,
}
