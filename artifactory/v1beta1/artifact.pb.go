// Code generated by protoc-gen-go.
// source: artifact.proto
// DO NOT EDIT!

/*
Package v1beta1 is a generated protocol buffer package.

It is generated from these files:
	artifact.proto
	version.proto

It has these top-level messages:
	ArtifactSearchRequest
	ArtifactSearchResponse
	ArtifactListRequest
	ArtifactListResponse
	Artifact
	VersionListRequest
	VersionListResponse
	VersionDescribeRequest
	VersionDescribeResponse
	ArtifactVersion
	JavaSpec
	DockerSpec
	PhpSpec
	NpmSpec
*/
package v1beta1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"
import dtypes "github.com/appscode/api/dtypes"

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

type ArtifactSearchRequest struct {
	Query string `protobuf:"bytes,1,opt,name=query" json:"query,omitempty"`
	Type  string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
}

func (m *ArtifactSearchRequest) Reset()                    { *m = ArtifactSearchRequest{} }
func (m *ArtifactSearchRequest) String() string            { return proto.CompactTextString(m) }
func (*ArtifactSearchRequest) ProtoMessage()               {}
func (*ArtifactSearchRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ArtifactSearchResponse struct {
	Status    *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Artifacts []*Artifact    `protobuf:"bytes,2,rep,name=artifacts" json:"artifacts,omitempty"`
}

func (m *ArtifactSearchResponse) Reset()                    { *m = ArtifactSearchResponse{} }
func (m *ArtifactSearchResponse) String() string            { return proto.CompactTextString(m) }
func (*ArtifactSearchResponse) ProtoMessage()               {}
func (*ArtifactSearchResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ArtifactSearchResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ArtifactSearchResponse) GetArtifacts() []*Artifact {
	if m != nil {
		return m.Artifacts
	}
	return nil
}

type ArtifactListRequest struct {
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
}

func (m *ArtifactListRequest) Reset()                    { *m = ArtifactListRequest{} }
func (m *ArtifactListRequest) String() string            { return proto.CompactTextString(m) }
func (*ArtifactListRequest) ProtoMessage()               {}
func (*ArtifactListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type ArtifactListResponse struct {
	Status    *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Artifacts []*Artifact    `protobuf:"bytes,2,rep,name=artifacts" json:"artifacts,omitempty"`
}

func (m *ArtifactListResponse) Reset()                    { *m = ArtifactListResponse{} }
func (m *ArtifactListResponse) String() string            { return proto.CompactTextString(m) }
func (*ArtifactListResponse) ProtoMessage()               {}
func (*ArtifactListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ArtifactListResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ArtifactListResponse) GetArtifacts() []*Artifact {
	if m != nil {
		return m.Artifacts
	}
	return nil
}

type Artifact struct {
	Name         string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Type         string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	LastModified int64  `protobuf:"varint,3,opt,name=last_modified,json=lastModified" json:"last_modified,omitempty"`
}

func (m *Artifact) Reset()                    { *m = Artifact{} }
func (m *Artifact) String() string            { return proto.CompactTextString(m) }
func (*Artifact) ProtoMessage()               {}
func (*Artifact) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func init() {
	proto.RegisterType((*ArtifactSearchRequest)(nil), "artifactory.v1beta1.ArtifactSearchRequest")
	proto.RegisterType((*ArtifactSearchResponse)(nil), "artifactory.v1beta1.ArtifactSearchResponse")
	proto.RegisterType((*ArtifactListRequest)(nil), "artifactory.v1beta1.ArtifactListRequest")
	proto.RegisterType((*ArtifactListResponse)(nil), "artifactory.v1beta1.ArtifactListResponse")
	proto.RegisterType((*Artifact)(nil), "artifactory.v1beta1.Artifact")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Artifacts service

type ArtifactsClient interface {
	Search(ctx context.Context, in *ArtifactSearchRequest, opts ...grpc.CallOption) (*ArtifactSearchResponse, error)
	List(ctx context.Context, in *ArtifactListRequest, opts ...grpc.CallOption) (*ArtifactListResponse, error)
}

type artifactsClient struct {
	cc *grpc.ClientConn
}

func NewArtifactsClient(cc *grpc.ClientConn) ArtifactsClient {
	return &artifactsClient{cc}
}

func (c *artifactsClient) Search(ctx context.Context, in *ArtifactSearchRequest, opts ...grpc.CallOption) (*ArtifactSearchResponse, error) {
	out := new(ArtifactSearchResponse)
	err := grpc.Invoke(ctx, "/artifactory.v1beta1.Artifacts/Search", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artifactsClient) List(ctx context.Context, in *ArtifactListRequest, opts ...grpc.CallOption) (*ArtifactListResponse, error) {
	out := new(ArtifactListResponse)
	err := grpc.Invoke(ctx, "/artifactory.v1beta1.Artifacts/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Artifacts service

type ArtifactsServer interface {
	Search(context.Context, *ArtifactSearchRequest) (*ArtifactSearchResponse, error)
	List(context.Context, *ArtifactListRequest) (*ArtifactListResponse, error)
}

func RegisterArtifactsServer(s *grpc.Server, srv ArtifactsServer) {
	s.RegisterService(&_Artifacts_serviceDesc, srv)
}

func _Artifacts_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArtifactSearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtifactsServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/artifactory.v1beta1.Artifacts/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtifactsServer).Search(ctx, req.(*ArtifactSearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Artifacts_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArtifactListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtifactsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/artifactory.v1beta1.Artifacts/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtifactsServer).List(ctx, req.(*ArtifactListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Artifacts_serviceDesc = grpc.ServiceDesc{
	ServiceName: "artifactory.v1beta1.Artifacts",
	HandlerType: (*ArtifactsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _Artifacts_Search_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Artifacts_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("artifact.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 385 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x93, 0xcf, 0x4e, 0xea, 0x40,
	0x14, 0xc6, 0xd3, 0xc2, 0xe5, 0xde, 0x1e, 0xee, 0x65, 0x31, 0x70, 0x49, 0xd3, 0x68, 0x42, 0x6a,
	0x62, 0x0a, 0x26, 0xad, 0x54, 0x57, 0xba, 0xc2, 0xb5, 0x6e, 0xca, 0xc2, 0xc4, 0x8d, 0x19, 0xe8,
	0x80, 0x4d, 0xa0, 0x53, 0x3a, 0x83, 0x09, 0x21, 0x6c, 0x7c, 0x05, 0x16, 0x3e, 0x8c, 0x8f, 0xe1,
	0x2b, 0xf8, 0x20, 0xa6, 0x9d, 0x19, 0xfe, 0x18, 0x82, 0xac, 0xdc, 0x34, 0x33, 0xdf, 0x9c, 0xf9,
	0xe6, 0x37, 0xe7, 0xeb, 0x40, 0x05, 0xa7, 0x3c, 0x1a, 0xe0, 0x3e, 0x77, 0x93, 0x94, 0x72, 0x8a,
	0xaa, 0x6a, 0x4e, 0xd3, 0x99, 0xfb, 0xdc, 0xee, 0x11, 0x8e, 0xdb, 0xd6, 0xd1, 0x90, 0xd2, 0xe1,
	0x88, 0x78, 0x38, 0x89, 0x3c, 0x1c, 0xc7, 0x94, 0x63, 0x1e, 0xd1, 0x98, 0x89, 0x2d, 0x56, 0x3d,
	0x93, 0x43, 0x3e, 0x4b, 0x08, 0xf3, 0xf2, 0xaf, 0xd0, 0xed, 0x0e, 0xfc, 0xef, 0x48, 0xb3, 0x2e,
	0xc1, 0x69, 0xff, 0x29, 0x20, 0x93, 0x29, 0x61, 0x1c, 0xd5, 0xe0, 0xd7, 0x64, 0x4a, 0xd2, 0x99,
	0xa9, 0x35, 0x34, 0xc7, 0x08, 0xc4, 0x04, 0x21, 0x28, 0x66, 0xbb, 0x4d, 0x3d, 0x17, 0xf3, 0xb1,
	0xbd, 0x80, 0xfa, 0x57, 0x0b, 0x96, 0xd0, 0x98, 0x11, 0x74, 0x0a, 0x25, 0xc6, 0x31, 0x9f, 0xb2,
	0xdc, 0xa4, 0xec, 0x57, 0x5c, 0x41, 0xe0, 0x76, 0x73, 0x35, 0x90, 0xab, 0xe8, 0x1a, 0x0c, 0x75,
	0x23, 0x66, 0xea, 0x8d, 0x82, 0x53, 0xf6, 0x8f, 0xdd, 0x1d, 0x77, 0x74, 0xd5, 0x39, 0xc1, 0xba,
	0xde, 0x6e, 0x42, 0x55, 0xc9, 0xb7, 0x11, 0xe3, 0x8a, 0x5f, 0x91, 0x6a, 0x1b, 0xa4, 0x73, 0xa8,
	0x6d, 0x97, 0xfe, 0x24, 0xe7, 0x3d, 0xfc, 0x51, 0x72, 0x06, 0x17, 0xe3, 0xf1, 0x0a, 0x2e, 0x1b,
	0xef, 0x6a, 0x2d, 0x3a, 0x81, 0x7f, 0x23, 0xcc, 0xf8, 0xe3, 0x98, 0x86, 0xd1, 0x20, 0x22, 0xa1,
	0x59, 0x68, 0x68, 0x4e, 0x21, 0xf8, 0x9b, 0x89, 0x77, 0x52, 0xf3, 0xdf, 0x74, 0x30, 0x94, 0x33,
	0x43, 0x4b, 0x0d, 0x4a, 0x22, 0x06, 0xd4, 0xda, 0xcb, 0xb6, 0x15, 0xb7, 0x75, 0x76, 0x50, 0xad,
	0xe8, 0x97, 0x7d, 0xfe, 0xf2, 0xfe, 0xb1, 0xd4, 0x5b, 0xc8, 0xf1, 0x70, 0x92, 0xb0, 0x3e, 0x0d,
	0xe5, 0x5f, 0xb7, 0x76, 0xf0, 0xa4, 0x83, 0xc7, 0x04, 0xca, 0xab, 0x06, 0xc5, 0xac, 0xe5, 0xc8,
	0xd9, 0x7b, 0xce, 0x46, 0x80, 0x56, 0xf3, 0x80, 0x4a, 0xc9, 0x73, 0x95, 0xf3, 0x5c, 0x22, 0xff,
	0x7b, 0x9e, 0x55, 0x1e, 0xde, 0x3c, 0xeb, 0xf0, 0xe2, 0xc6, 0x78, 0xf8, 0x2d, 0xd7, 0x7a, 0xa5,
	0xfc, 0x49, 0x5c, 0x7c, 0x06, 0x00, 0x00, 0xff, 0xff, 0x05, 0x02, 0xef, 0x3e, 0x6f, 0x03, 0x00,
	0x00,
}
