// Code generated by protoc-gen-go.
// source: version.proto
// DO NOT EDIT!

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

type VersionListRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
}

func (m *VersionListRequest) Reset()                    { *m = VersionListRequest{} }
func (m *VersionListRequest) String() string            { return proto.CompactTextString(m) }
func (*VersionListRequest) ProtoMessage()               {}
func (*VersionListRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type VersionListResponse struct {
	Status   *dtypes.Status     `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Versions []*ArtifactVersion `protobuf:"bytes,2,rep,name=versions" json:"versions,omitempty"`
}

func (m *VersionListResponse) Reset()                    { *m = VersionListResponse{} }
func (m *VersionListResponse) String() string            { return proto.CompactTextString(m) }
func (*VersionListResponse) ProtoMessage()               {}
func (*VersionListResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *VersionListResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *VersionListResponse) GetVersions() []*ArtifactVersion {
	if m != nil {
		return m.Versions
	}
	return nil
}

type VersionDescribeRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Id   string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	Type string `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
}

func (m *VersionDescribeRequest) Reset()                    { *m = VersionDescribeRequest{} }
func (m *VersionDescribeRequest) String() string            { return proto.CompactTextString(m) }
func (*VersionDescribeRequest) ProtoMessage()               {}
func (*VersionDescribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

type VersionDescribeResponse struct {
	Status  *dtypes.Status   `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Version *ArtifactVersion `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
}

func (m *VersionDescribeResponse) Reset()                    { *m = VersionDescribeResponse{} }
func (m *VersionDescribeResponse) String() string            { return proto.CompactTextString(m) }
func (*VersionDescribeResponse) ProtoMessage()               {}
func (*VersionDescribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *VersionDescribeResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *VersionDescribeResponse) GetVersion() *ArtifactVersion {
	if m != nil {
		return m.Version
	}
	return nil
}

type ArtifactVersion struct {
	Id               string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name             string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Version          string `protobuf:"bytes,3,opt,name=version" json:"version,omitempty"`
	LastModifiedTime string `protobuf:"bytes,4,opt,name=last_modified_time,json=lastModifiedTime" json:"last_modified_time,omitempty"`
	Type             string `protobuf:"bytes,5,opt,name=type" json:"type,omitempty"`
	// Types that are valid to be assigned to Specs:
	//	*ArtifactVersion_Java
	//	*ArtifactVersion_Docker
	//	*ArtifactVersion_Php
	//	*ArtifactVersion_Npm
	Specs isArtifactVersion_Specs `protobuf_oneof:"specs"`
}

func (m *ArtifactVersion) Reset()                    { *m = ArtifactVersion{} }
func (m *ArtifactVersion) String() string            { return proto.CompactTextString(m) }
func (*ArtifactVersion) ProtoMessage()               {}
func (*ArtifactVersion) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

type isArtifactVersion_Specs interface {
	isArtifactVersion_Specs()
}

type ArtifactVersion_Java struct {
	Java *JavaSpec `protobuf:"bytes,6,opt,name=java,oneof"`
}
type ArtifactVersion_Docker struct {
	Docker *DockerSpec `protobuf:"bytes,7,opt,name=docker,oneof"`
}
type ArtifactVersion_Php struct {
	Php *PhpSpec `protobuf:"bytes,8,opt,name=php,oneof"`
}
type ArtifactVersion_Npm struct {
	Npm *NpmSpec `protobuf:"bytes,9,opt,name=npm,oneof"`
}

func (*ArtifactVersion_Java) isArtifactVersion_Specs()   {}
func (*ArtifactVersion_Docker) isArtifactVersion_Specs() {}
func (*ArtifactVersion_Php) isArtifactVersion_Specs()    {}
func (*ArtifactVersion_Npm) isArtifactVersion_Specs()    {}

func (m *ArtifactVersion) GetSpecs() isArtifactVersion_Specs {
	if m != nil {
		return m.Specs
	}
	return nil
}

func (m *ArtifactVersion) GetJava() *JavaSpec {
	if x, ok := m.GetSpecs().(*ArtifactVersion_Java); ok {
		return x.Java
	}
	return nil
}

func (m *ArtifactVersion) GetDocker() *DockerSpec {
	if x, ok := m.GetSpecs().(*ArtifactVersion_Docker); ok {
		return x.Docker
	}
	return nil
}

func (m *ArtifactVersion) GetPhp() *PhpSpec {
	if x, ok := m.GetSpecs().(*ArtifactVersion_Php); ok {
		return x.Php
	}
	return nil
}

func (m *ArtifactVersion) GetNpm() *NpmSpec {
	if x, ok := m.GetSpecs().(*ArtifactVersion_Npm); ok {
		return x.Npm
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ArtifactVersion) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ArtifactVersion_OneofMarshaler, _ArtifactVersion_OneofUnmarshaler, _ArtifactVersion_OneofSizer, []interface{}{
		(*ArtifactVersion_Java)(nil),
		(*ArtifactVersion_Docker)(nil),
		(*ArtifactVersion_Php)(nil),
		(*ArtifactVersion_Npm)(nil),
	}
}

func _ArtifactVersion_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ArtifactVersion)
	// specs
	switch x := m.Specs.(type) {
	case *ArtifactVersion_Java:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Java); err != nil {
			return err
		}
	case *ArtifactVersion_Docker:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Docker); err != nil {
			return err
		}
	case *ArtifactVersion_Php:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Php); err != nil {
			return err
		}
	case *ArtifactVersion_Npm:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Npm); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ArtifactVersion.Specs has unexpected type %T", x)
	}
	return nil
}

func _ArtifactVersion_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ArtifactVersion)
	switch tag {
	case 6: // specs.java
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(JavaSpec)
		err := b.DecodeMessage(msg)
		m.Specs = &ArtifactVersion_Java{msg}
		return true, err
	case 7: // specs.docker
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DockerSpec)
		err := b.DecodeMessage(msg)
		m.Specs = &ArtifactVersion_Docker{msg}
		return true, err
	case 8: // specs.php
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PhpSpec)
		err := b.DecodeMessage(msg)
		m.Specs = &ArtifactVersion_Php{msg}
		return true, err
	case 9: // specs.npm
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(NpmSpec)
		err := b.DecodeMessage(msg)
		m.Specs = &ArtifactVersion_Npm{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ArtifactVersion_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ArtifactVersion)
	// specs
	switch x := m.Specs.(type) {
	case *ArtifactVersion_Java:
		s := proto.Size(x.Java)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ArtifactVersion_Docker:
		s := proto.Size(x.Docker)
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ArtifactVersion_Php:
		s := proto.Size(x.Php)
		n += proto.SizeVarint(8<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ArtifactVersion_Npm:
		s := proto.Size(x.Npm)
		n += proto.SizeVarint(9<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type JavaSpec struct {
	GroupId     string `protobuf:"bytes,1,opt,name=group_id,json=groupId" json:"group_id,omitempty"`
	ArtifactId  string `protobuf:"bytes,2,opt,name=artifact_id,json=artifactId" json:"artifact_id,omitempty"`
	ArtifactUrl string `protobuf:"bytes,3,opt,name=artifact_url,json=artifactUrl" json:"artifact_url,omitempty"`
}

func (m *JavaSpec) Reset()                    { *m = JavaSpec{} }
func (m *JavaSpec) String() string            { return proto.CompactTextString(m) }
func (*JavaSpec) ProtoMessage()               {}
func (*JavaSpec) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

type DockerSpec struct {
	TarSums     []string `protobuf:"bytes,1,rep,name=tar_sums,json=tarSums" json:"tar_sums,omitempty"`
	ManifestUrl string   `protobuf:"bytes,2,opt,name=manifest_url,json=manifestUrl" json:"manifest_url,omitempty"`
}

func (m *DockerSpec) Reset()                    { *m = DockerSpec{} }
func (m *DockerSpec) String() string            { return proto.CompactTextString(m) }
func (*DockerSpec) ProtoMessage()               {}
func (*DockerSpec) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

type PhpSpec struct {
	DownloadUrl string `protobuf:"bytes,1,opt,name=download_url,json=downloadUrl" json:"download_url,omitempty"`
	Shasum      string `protobuf:"bytes,2,opt,name=shasum" json:"shasum,omitempty"`
}

func (m *PhpSpec) Reset()                    { *m = PhpSpec{} }
func (m *PhpSpec) String() string            { return proto.CompactTextString(m) }
func (*PhpSpec) ProtoMessage()               {}
func (*PhpSpec) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

type NpmSpec struct {
	Description string   `protobuf:"bytes,1,opt,name=description" json:"description,omitempty"`
	ArtifactId  string   `protobuf:"bytes,2,opt,name=artifact_id,json=artifactId" json:"artifact_id,omitempty"`
	Shasum      string   `protobuf:"bytes,3,opt,name=shasum" json:"shasum,omitempty"`
	Keywords    []string `protobuf:"bytes,4,rep,name=keywords" json:"keywords,omitempty"`
	ArtifactUrl string   `protobuf:"bytes,5,opt,name=artifact_url,json=artifactUrl" json:"artifact_url,omitempty"`
}

func (m *NpmSpec) Reset()                    { *m = NpmSpec{} }
func (m *NpmSpec) String() string            { return proto.CompactTextString(m) }
func (*NpmSpec) ProtoMessage()               {}
func (*NpmSpec) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func init() {
	proto.RegisterType((*VersionListRequest)(nil), "artifactory.v1beta1.VersionListRequest")
	proto.RegisterType((*VersionListResponse)(nil), "artifactory.v1beta1.VersionListResponse")
	proto.RegisterType((*VersionDescribeRequest)(nil), "artifactory.v1beta1.VersionDescribeRequest")
	proto.RegisterType((*VersionDescribeResponse)(nil), "artifactory.v1beta1.VersionDescribeResponse")
	proto.RegisterType((*ArtifactVersion)(nil), "artifactory.v1beta1.ArtifactVersion")
	proto.RegisterType((*JavaSpec)(nil), "artifactory.v1beta1.JavaSpec")
	proto.RegisterType((*DockerSpec)(nil), "artifactory.v1beta1.DockerSpec")
	proto.RegisterType((*PhpSpec)(nil), "artifactory.v1beta1.PhpSpec")
	proto.RegisterType((*NpmSpec)(nil), "artifactory.v1beta1.NpmSpec")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Versions service

type VersionsClient interface {
	List(ctx context.Context, in *VersionListRequest, opts ...grpc.CallOption) (*VersionListResponse, error)
	Describe(ctx context.Context, in *VersionDescribeRequest, opts ...grpc.CallOption) (*VersionDescribeResponse, error)
}

type versionsClient struct {
	cc *grpc.ClientConn
}

func NewVersionsClient(cc *grpc.ClientConn) VersionsClient {
	return &versionsClient{cc}
}

func (c *versionsClient) List(ctx context.Context, in *VersionListRequest, opts ...grpc.CallOption) (*VersionListResponse, error) {
	out := new(VersionListResponse)
	err := grpc.Invoke(ctx, "/artifactory.v1beta1.Versions/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *versionsClient) Describe(ctx context.Context, in *VersionDescribeRequest, opts ...grpc.CallOption) (*VersionDescribeResponse, error) {
	out := new(VersionDescribeResponse)
	err := grpc.Invoke(ctx, "/artifactory.v1beta1.Versions/Describe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Versions service

type VersionsServer interface {
	List(context.Context, *VersionListRequest) (*VersionListResponse, error)
	Describe(context.Context, *VersionDescribeRequest) (*VersionDescribeResponse, error)
}

func RegisterVersionsServer(s *grpc.Server, srv VersionsServer) {
	s.RegisterService(&_Versions_serviceDesc, srv)
}

func _Versions_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VersionsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/artifactory.v1beta1.Versions/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VersionsServer).List(ctx, req.(*VersionListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Versions_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionDescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VersionsServer).Describe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/artifactory.v1beta1.Versions/Describe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VersionsServer).Describe(ctx, req.(*VersionDescribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Versions_serviceDesc = grpc.ServiceDesc{
	ServiceName: "artifactory.v1beta1.Versions",
	HandlerType: (*VersionsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Versions_List_Handler,
		},
		{
			MethodName: "Describe",
			Handler:    _Versions_Describe_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor1,
}

func init() { proto.RegisterFile("version.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 683 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x55, 0xdd, 0x4e, 0x13, 0x41,
	0x14, 0xb6, 0xdb, 0xd2, 0x5d, 0x4e, 0x15, 0xcd, 0x90, 0xe0, 0xda, 0x60, 0xc0, 0x8d, 0x51, 0x12,
	0x49, 0x57, 0xe0, 0xca, 0xf8, 0x13, 0x45, 0x12, 0x85, 0xf8, 0x43, 0x8a, 0x78, 0xe1, 0x4d, 0x33,
	0x74, 0x06, 0x18, 0xe9, 0xee, 0x8c, 0x33, 0xb3, 0x25, 0xc4, 0x18, 0xa3, 0xaf, 0xe0, 0xbd, 0x37,
	0xbe, 0x83, 0x2f, 0xe2, 0x2b, 0x70, 0xed, 0x33, 0x98, 0x99, 0x9d, 0xdd, 0x56, 0xda, 0x00, 0xc6,
	0x9b, 0x66, 0xe7, 0x9b, 0xf3, 0x9d, 0xef, 0x7c, 0x67, 0x66, 0x4e, 0xe1, 0x52, 0x9f, 0x4a, 0xc5,
	0x78, 0xda, 0x12, 0x92, 0x6b, 0x8e, 0xa6, 0xb1, 0xd4, 0x6c, 0x17, 0x77, 0x35, 0x97, 0x47, 0xad,
	0xfe, 0xd2, 0x0e, 0xd5, 0x78, 0xa9, 0x39, 0xbb, 0xc7, 0xf9, 0x5e, 0x8f, 0xc6, 0x58, 0xb0, 0x18,
	0xa7, 0x29, 0xd7, 0x58, 0x33, 0x9e, 0xaa, 0x9c, 0xd2, 0x9c, 0x31, 0x30, 0xd1, 0x47, 0x82, 0xaa,
	0xd8, 0xfe, 0xe6, 0x78, 0xf4, 0x00, 0xd0, 0xdb, 0x3c, 0xf7, 0x0b, 0xa6, 0x74, 0x9b, 0x7e, 0xc8,
	0xa8, 0xd2, 0x08, 0x41, 0x2d, 0xc5, 0x09, 0x0d, 0x2b, 0xf3, 0x95, 0x85, 0xc9, 0xb6, 0xfd, 0x36,
	0x98, 0x21, 0x86, 0x5e, 0x8e, 0x99, 0xef, 0xe8, 0x33, 0x4c, 0xff, 0xc5, 0x56, 0x82, 0xa7, 0x8a,
	0xa2, 0x5b, 0x50, 0x57, 0x1a, 0xeb, 0x4c, 0xd9, 0x04, 0x8d, 0xe5, 0xa9, 0x56, 0xae, 0xdc, 0xda,
	0xb2, 0x68, 0xdb, 0xed, 0xa2, 0xc7, 0x10, 0x38, 0x63, 0x2a, 0xf4, 0xe6, 0xab, 0x0b, 0x8d, 0xe5,
	0x9b, 0xad, 0x31, 0xd6, 0x5a, 0x4f, 0x1c, 0xe6, 0xb4, 0xda, 0x25, 0x2b, 0xda, 0x84, 0x19, 0x07,
	0xae, 0x51, 0xd5, 0x95, 0x6c, 0x87, 0x9e, 0x66, 0x61, 0x0a, 0x3c, 0x46, 0x9c, 0x01, 0x8f, 0x91,
	0xd2, 0x52, 0x75, 0xc8, 0xd2, 0x97, 0x0a, 0x5c, 0x1d, 0x49, 0xf9, 0x8f, 0xbe, 0x1e, 0x81, 0xef,
	0x2a, 0xb4, 0x62, 0xe7, 0xb5, 0x55, 0x90, 0xa2, 0xdf, 0x1e, 0x5c, 0x3e, 0xb1, 0xe9, 0x6a, 0xaf,
	0x0c, 0xd7, 0x6e, 0xfd, 0x79, 0x43, 0xfe, 0xc2, 0x81, 0x6e, 0x6e, 0xa9, 0x58, 0xa2, 0x45, 0x40,
	0x3d, 0xac, 0x74, 0x27, 0xe1, 0x84, 0xed, 0x32, 0x4a, 0x3a, 0x9a, 0x25, 0x34, 0xac, 0xd9, 0xa0,
	0x2b, 0x66, 0xe7, 0xa5, 0xdb, 0x78, 0xc3, 0x86, 0x8e, 0x7a, 0x62, 0xd0, 0x17, 0xb4, 0x02, 0xb5,
	0xf7, 0xb8, 0x8f, 0xc3, 0xba, 0x35, 0x74, 0x7d, 0xac, 0xa1, 0x0d, 0xdc, 0xc7, 0x5b, 0x82, 0x76,
	0x9f, 0x5f, 0x68, 0xdb, 0x60, 0x74, 0x0f, 0xea, 0x84, 0x77, 0x0f, 0xa8, 0x0c, 0x7d, 0x4b, 0x9b,
	0x1b, 0x4b, 0x5b, 0xb3, 0x21, 0x8e, 0xe8, 0x08, 0xe8, 0x2e, 0x54, 0xc5, 0xbe, 0x08, 0x03, 0xcb,
	0x9b, 0x1d, 0xcb, 0xdb, 0xdc, 0x17, 0x8e, 0x64, 0x42, 0x0d, 0x23, 0x15, 0x49, 0x38, 0x79, 0x0a,
	0xe3, 0x95, 0x48, 0x0a, 0x46, 0x2a, 0x92, 0x55, 0x1f, 0x26, 0x94, 0xa0, 0x5d, 0x15, 0x31, 0x08,
	0x8a, 0xda, 0xd1, 0x35, 0x08, 0xf6, 0x24, 0xcf, 0x44, 0xa7, 0x6c, 0xb7, 0x6f, 0xd7, 0xeb, 0x04,
	0xcd, 0x41, 0xa3, 0xc8, 0xda, 0x29, 0x2f, 0x12, 0x14, 0xd0, 0x3a, 0x41, 0x37, 0xe0, 0x62, 0x19,
	0x90, 0xc9, 0x9e, 0x3b, 0x85, 0x92, 0xb4, 0x2d, 0x7b, 0xd1, 0x06, 0xc0, 0xc0, 0xaf, 0x11, 0xd3,
	0x58, 0x76, 0x54, 0x96, 0x98, 0x3b, 0x55, 0x35, 0x62, 0x1a, 0xcb, 0xad, 0x2c, 0x51, 0x26, 0x57,
	0x82, 0x53, 0xb6, 0x4b, 0x55, 0x9e, 0x2b, 0x57, 0x6b, 0x14, 0x98, 0xc9, 0xb5, 0x06, 0xbe, 0xeb,
	0x81, 0x89, 0x26, 0xfc, 0x30, 0xed, 0x71, 0x4c, 0x6c, 0x74, 0x5e, 0x79, 0xa3, 0xc0, 0xb6, 0x65,
	0x0f, 0xcd, 0x40, 0x5d, 0xed, 0x63, 0x95, 0x25, 0x2e, 0x95, 0x5b, 0x45, 0x3f, 0x2a, 0xe0, 0xbb,
	0xc6, 0xa0, 0x79, 0x68, 0x10, 0x7b, 0xeb, 0x85, 0x19, 0x1e, 0x65, 0x96, 0x01, 0x74, 0x76, 0x0f,
	0x06, 0x32, 0xd5, 0x61, 0x19, 0xd4, 0x84, 0xe0, 0x80, 0x1e, 0x1d, 0x72, 0x49, 0x54, 0x58, 0xb3,
	0x56, 0xcb, 0xf5, 0x48, 0xdf, 0x26, 0x46, 0xfa, 0xb6, 0x7c, 0xec, 0x41, 0xe0, 0xde, 0x82, 0x42,
	0xdf, 0x2b, 0x50, 0x33, 0x13, 0x07, 0xdd, 0x1e, 0x7b, 0xcc, 0xa3, 0x13, 0xad, 0xb9, 0x70, 0x76,
	0x60, 0xfe, 0xc8, 0xa3, 0xa7, 0x5f, 0x7f, 0x1d, 0x7f, 0xf3, 0x1e, 0xa2, 0xfb, 0x31, 0x16, 0x42,
	0x75, 0x39, 0x71, 0x23, 0x75, 0x40, 0x8f, 0x1d, 0xbd, 0xc4, 0x54, 0xfc, 0xd1, 0xbc, 0x91, 0x4f,
	0x71, 0x31, 0x97, 0xd0, 0xcf, 0x0a, 0x04, 0xc5, 0xf8, 0x40, 0x77, 0x4e, 0xd3, 0x3e, 0x31, 0xb7,
	0x9a, 0x8b, 0xe7, 0x0b, 0x76, 0xc5, 0xbe, 0xb6, 0xc5, 0xae, 0xa3, 0x67, 0xff, 0x51, 0x6c, 0x3c,
	0x74, 0xbc, 0xab, 0x93, 0xef, 0x7c, 0x47, 0xd8, 0xa9, 0xdb, 0x7f, 0x88, 0x95, 0x3f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xce, 0x70, 0x81, 0xe2, 0x7d, 0x06, 0x00, 0x00,
}