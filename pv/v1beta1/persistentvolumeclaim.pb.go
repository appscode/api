// Code generated by protoc-gen-go.
// source: persistentvolumeclaim.proto
// DO NOT EDIT!

package v1beta1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/appscodeapis/appscode/api"
import appscode_dtypes "github.com/appscode/api/dtypes"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type PVCRegisterRequest struct {
	Cluster   string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	SizeGb    int64  `protobuf:"varint,3,opt,name=size_gb,json=sizeGb" json:"size_gb,omitempty"`
	Namespace string `protobuf:"bytes,4,opt,name=namespace" json:"namespace,omitempty"`
}

func (m *PVCRegisterRequest) Reset()                    { *m = PVCRegisterRequest{} }
func (m *PVCRegisterRequest) String() string            { return proto.CompactTextString(m) }
func (*PVCRegisterRequest) ProtoMessage()               {}
func (*PVCRegisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type PVCUnregisterRequest struct {
	Cluster   string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Namespace string `protobuf:"bytes,3,opt,name=namespace" json:"namespace,omitempty"`
}

func (m *PVCUnregisterRequest) Reset()                    { *m = PVCUnregisterRequest{} }
func (m *PVCUnregisterRequest) String() string            { return proto.CompactTextString(m) }
func (*PVCUnregisterRequest) ProtoMessage()               {}
func (*PVCUnregisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type PVCDescribeRequest struct {
	Cluster   string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Namespace string `protobuf:"bytes,3,opt,name=namespace" json:"namespace,omitempty"`
}

func (m *PVCDescribeRequest) Reset()                    { *m = PVCDescribeRequest{} }
func (m *PVCDescribeRequest) String() string            { return proto.CompactTextString(m) }
func (*PVCDescribeRequest) ProtoMessage()               {}
func (*PVCDescribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

type PVCInfo struct {
	Name        string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	SizeGb      int64    `protobuf:"varint,2,opt,name=size_gb,json=sizeGb" json:"size_gb,omitempty"`
	Namespace   string   `protobuf:"bytes,3,opt,name=namespace" json:"namespace,omitempty"`
	Status      string   `protobuf:"bytes,4,opt,name=status" json:"status,omitempty"`
	Volume      string   `protobuf:"bytes,5,opt,name=volume" json:"volume,omitempty"`
	AccessModes []string `protobuf:"bytes,6,rep,name=accessModes" json:"accessModes,omitempty"`
}

func (m *PVCInfo) Reset()                    { *m = PVCInfo{} }
func (m *PVCInfo) String() string            { return proto.CompactTextString(m) }
func (*PVCInfo) ProtoMessage()               {}
func (*PVCInfo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

type PVCDescribeResponse struct {
	Status *appscode_dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Pvc    *PVCInfo                `protobuf:"bytes,2,opt,name=pvc" json:"pvc,omitempty"`
}

func (m *PVCDescribeResponse) Reset()                    { *m = PVCDescribeResponse{} }
func (m *PVCDescribeResponse) String() string            { return proto.CompactTextString(m) }
func (*PVCDescribeResponse) ProtoMessage()               {}
func (*PVCDescribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *PVCDescribeResponse) GetStatus() *appscode_dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *PVCDescribeResponse) GetPvc() *PVCInfo {
	if m != nil {
		return m.Pvc
	}
	return nil
}

func init() {
	proto.RegisterType((*PVCRegisterRequest)(nil), "appscode.pv.v1beta1.PVCRegisterRequest")
	proto.RegisterType((*PVCUnregisterRequest)(nil), "appscode.pv.v1beta1.PVCUnregisterRequest")
	proto.RegisterType((*PVCDescribeRequest)(nil), "appscode.pv.v1beta1.PVCDescribeRequest")
	proto.RegisterType((*PVCInfo)(nil), "appscode.pv.v1beta1.PVCInfo")
	proto.RegisterType((*PVCDescribeResponse)(nil), "appscode.pv.v1beta1.PVCDescribeResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for PersistentVolumeClaims service

type PersistentVolumeClaimsClient interface {
	Describe(ctx context.Context, in *PVCDescribeRequest, opts ...grpc.CallOption) (*PVCDescribeResponse, error)
	Register(ctx context.Context, in *PVCRegisterRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
	Unregister(ctx context.Context, in *PVCUnregisterRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
}

type persistentVolumeClaimsClient struct {
	cc *grpc.ClientConn
}

func NewPersistentVolumeClaimsClient(cc *grpc.ClientConn) PersistentVolumeClaimsClient {
	return &persistentVolumeClaimsClient{cc}
}

func (c *persistentVolumeClaimsClient) Describe(ctx context.Context, in *PVCDescribeRequest, opts ...grpc.CallOption) (*PVCDescribeResponse, error) {
	out := new(PVCDescribeResponse)
	err := grpc.Invoke(ctx, "/appscode.pv.v1beta1.PersistentVolumeClaims/Describe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *persistentVolumeClaimsClient) Register(ctx context.Context, in *PVCRegisterRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.pv.v1beta1.PersistentVolumeClaims/Register", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *persistentVolumeClaimsClient) Unregister(ctx context.Context, in *PVCUnregisterRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.pv.v1beta1.PersistentVolumeClaims/Unregister", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PersistentVolumeClaims service

type PersistentVolumeClaimsServer interface {
	Describe(context.Context, *PVCDescribeRequest) (*PVCDescribeResponse, error)
	Register(context.Context, *PVCRegisterRequest) (*appscode_dtypes.VoidResponse, error)
	Unregister(context.Context, *PVCUnregisterRequest) (*appscode_dtypes.VoidResponse, error)
}

func RegisterPersistentVolumeClaimsServer(s *grpc.Server, srv PersistentVolumeClaimsServer) {
	s.RegisterService(&_PersistentVolumeClaims_serviceDesc, srv)
}

func _PersistentVolumeClaims_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PVCDescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersistentVolumeClaimsServer).Describe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.pv.v1beta1.PersistentVolumeClaims/Describe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersistentVolumeClaimsServer).Describe(ctx, req.(*PVCDescribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersistentVolumeClaims_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PVCRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersistentVolumeClaimsServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.pv.v1beta1.PersistentVolumeClaims/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersistentVolumeClaimsServer).Register(ctx, req.(*PVCRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersistentVolumeClaims_Unregister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PVCUnregisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersistentVolumeClaimsServer).Unregister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.pv.v1beta1.PersistentVolumeClaims/Unregister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersistentVolumeClaimsServer).Unregister(ctx, req.(*PVCUnregisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PersistentVolumeClaims_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.pv.v1beta1.PersistentVolumeClaims",
	HandlerType: (*PersistentVolumeClaimsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Describe",
			Handler:    _PersistentVolumeClaims_Describe_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _PersistentVolumeClaims_Register_Handler,
		},
		{
			MethodName: "Unregister",
			Handler:    _PersistentVolumeClaims_Unregister_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor1,
}

func init() { proto.RegisterFile("persistentvolumeclaim.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 533 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xd6, 0x26, 0x6d, 0x7e, 0x36, 0xb7, 0x2d, 0x6a, 0x2d, 0x37, 0x40, 0x94, 0x0b, 0xa1, 0x07,
	0xaf, 0x1a, 0x6e, 0x1c, 0x9b, 0x4a, 0x88, 0x03, 0x92, 0x65, 0x44, 0x0e, 0x5c, 0x60, 0xed, 0x0c,
	0xc1, 0x10, 0x7b, 0x17, 0xcf, 0xda, 0x52, 0xa8, 0x72, 0xe9, 0x1b, 0x20, 0xde, 0x81, 0x37, 0xe0,
	0x31, 0xe0, 0x82, 0x78, 0x03, 0x1e, 0x04, 0xf9, 0x27, 0x4e, 0x9c, 0xc6, 0xea, 0x01, 0xd1, 0x8b,
	0xb5, 0x33, 0x3b, 0xb3, 0x33, 0xdf, 0xcc, 0xf7, 0x99, 0x9e, 0x2a, 0x88, 0xd0, 0x47, 0x0d, 0xa1,
	0x4e, 0xe4, 0x22, 0x0e, 0xc0, 0x5b, 0x08, 0x3f, 0xb0, 0x54, 0x24, 0xb5, 0x64, 0x47, 0x42, 0x29,
	0xf4, 0xe4, 0x0c, 0x2c, 0x95, 0x58, 0xc9, 0xb9, 0x0b, 0x5a, 0x9c, 0x9b, 0xfd, 0xb9, 0x94, 0xf3,
	0x05, 0x70, 0xa1, 0x7c, 0x2e, 0xc2, 0x50, 0x6a, 0xa1, 0x7d, 0x19, 0x62, 0x9e, 0x62, 0x3e, 0x58,
	0xa7, 0xd4, 0xdc, 0x3f, 0xac, 0xdc, 0xcf, 0xf4, 0x52, 0x01, 0xf2, 0xec, 0x9b, 0x07, 0x0c, 0x97,
	0x94, 0xd9, 0xd3, 0x89, 0x03, 0xf3, 0xb4, 0xa9, 0xc8, 0x81, 0x4f, 0x31, 0xa0, 0x66, 0x06, 0x6d,
	0x7b, 0x8b, 0x38, 0xf5, 0x18, 0x64, 0x40, 0x46, 0x5d, 0x67, 0x6d, 0x32, 0x46, 0x0f, 0x42, 0x11,
	0x80, 0xd1, 0xc8, 0xdc, 0xd9, 0x99, 0x9d, 0xd0, 0x36, 0xfa, 0x9f, 0xe1, 0xcd, 0xdc, 0x35, 0x9a,
	0x03, 0x32, 0x6a, 0x3a, 0xad, 0xd4, 0x7c, 0xe6, 0xb2, 0x3e, 0xed, 0xa6, 0x01, 0xa8, 0x84, 0x07,
	0xc6, 0x41, 0x96, 0xb1, 0x71, 0x0c, 0x5d, 0x7a, 0xcf, 0x9e, 0x4e, 0x5e, 0x85, 0xd1, 0x3f, 0x15,
	0xaf, 0xd4, 0x68, 0xee, 0xd6, 0x78, 0x9b, 0xc1, 0xbb, 0x04, 0xf4, 0x22, 0xdf, 0x85, 0xff, 0x51,
	0xe1, 0x1b, 0xa1, 0x6d, 0x7b, 0x3a, 0x79, 0x1e, 0xbe, 0x93, 0x65, 0x36, 0xd9, 0x3f, 0x9c, 0x46,
	0xfd, 0x70, 0x76, 0x9f, 0x65, 0xc7, 0xb4, 0x85, 0x5a, 0xe8, 0x18, 0x8b, 0xb9, 0x15, 0x56, 0xea,
	0xcf, 0x89, 0x63, 0x1c, 0xe6, 0xfe, 0xdc, 0x62, 0x03, 0xda, 0x13, 0x9e, 0x07, 0x88, 0x2f, 0xe4,
	0x0c, 0xd0, 0x68, 0x0d, 0x9a, 0xa3, 0xae, 0xb3, 0xed, 0x1a, 0x26, 0xf4, 0xa8, 0x32, 0x0a, 0x54,
	0x32, 0x44, 0x60, 0xbc, 0x2c, 0x94, 0x76, 0xdd, 0x1b, 0x9f, 0x58, 0x25, 0x0b, 0x73, 0xba, 0x58,
	0x2f, 0xb3, 0xeb, 0xb2, 0x03, 0x8b, 0x36, 0x55, 0xe2, 0x65, 0x60, 0x7a, 0xe3, 0xbe, 0xb5, 0x87,
	0xb3, 0x56, 0x31, 0x0f, 0x27, 0x0d, 0x1c, 0x7f, 0x39, 0xa4, 0xc7, 0x76, 0xc9, 0xfa, 0x69, 0xd6,
	0xee, 0x24, 0x65, 0x3d, 0xb2, 0xdf, 0x84, 0x76, 0xd6, 0x0d, 0xb1, 0x47, 0x75, 0x4f, 0xed, 0x6c,
	0xcf, 0x1c, 0xdd, 0x1e, 0x98, 0x63, 0x1b, 0x26, 0xd7, 0xdf, 0x8d, 0x46, 0x87, 0x5c, 0xff, 0xfa,
	0xf3, 0xb5, 0xf1, 0x81, 0xbd, 0xe7, 0x15, 0x31, 0x7c, 0x8c, 0x5d, 0x88, 0x42, 0xd0, 0x80, 0xbc,
	0x78, 0x83, 0x17, 0x2c, 0x40, 0x7e, 0x55, 0x9c, 0x56, 0xbc, 0xdc, 0x08, 0xf2, 0xab, 0xf2, 0xbc,
	0xe2, 0x7b, 0x65, 0x5c, 0x84, 0xac, 0xd8, 0x0f, 0x42, 0x3b, 0x6b, 0x49, 0xd5, 0xe3, 0xda, 0x11,
	0x9d, 0x79, 0xff, 0xc6, 0xe4, 0xa7, 0xd2, 0x9f, 0x95, 0x60, 0x96, 0x5b, 0x60, 0x02, 0xf3, 0xce,
	0xc0, 0x3c, 0x25, 0x67, 0xec, 0x27, 0xa1, 0x74, 0xa3, 0x53, 0xf6, 0xb8, 0x0e, 0xd1, 0x0d, 0x2d,
	0xdf, 0x86, 0xa9, 0xb2, 0xa0, 0xb3, 0x3b, 0xc3, 0x74, 0x71, 0x49, 0x4f, 0x3d, 0x19, 0x6c, 0x7a,
	0x13, 0xca, 0xdf, 0x82, 0x72, 0x61, 0xee, 0xe5, 0xab, 0x9d, 0xfe, 0x30, 0x6d, 0xf2, 0xba, 0x5d,
	0x84, 0xb9, 0xad, 0xec, 0x17, 0xfa, 0xe4, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xda, 0x02, 0x7f,
	0xd0, 0xd5, 0x05, 0x00, 0x00,
}
