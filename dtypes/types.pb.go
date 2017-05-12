// Code generated by protoc-gen-go. DO NOT EDIT.
// source: types.proto

/*
Package dtypes is a generated protocol buffer package.

It is generated from these files:
	types.proto

It has these top-level messages:
	ErrorDetails
	LongRunningResponse
	VoidRequest
	VoidResponse
	Uid
*/
package dtypes

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/any"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Basic Error details message to send in response. Application specific
// error messages can be provided.
type ErrorDetails struct {
	RequestedResource string `protobuf:"bytes,1,opt,name=requested_resource,json=requestedResource" json:"requested_resource,omitempty"`
	Stacktrace        string `protobuf:"bytes,2,opt,name=stacktrace" json:"stacktrace,omitempty"`
}

func (m *ErrorDetails) Reset()                    { *m = ErrorDetails{} }
func (m *ErrorDetails) String() string            { return proto.CompactTextString(m) }
func (*ErrorDetails) ProtoMessage()               {}
func (*ErrorDetails) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ErrorDetails) GetRequestedResource() string {
	if m != nil {
		return m.RequestedResource
	}
	return ""
}

func (m *ErrorDetails) GetStacktrace() string {
	if m != nil {
		return m.Stacktrace
	}
	return ""
}

// Types for long running operation. usually called as jobs.
// Next Id = 3
type LongRunningResponse struct {
	JobPhid string `protobuf:"bytes,1,opt,name=job_phid,json=jobPhid" json:"job_phid,omitempty"`
}

func (m *LongRunningResponse) Reset()                    { *m = LongRunningResponse{} }
func (m *LongRunningResponse) String() string            { return proto.CompactTextString(m) }
func (*LongRunningResponse) ProtoMessage()               {}
func (*LongRunningResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LongRunningResponse) GetJobPhid() string {
	if m != nil {
		return m.JobPhid
	}
	return ""
}

// Void Requests and response to use with other types.
type VoidRequest struct {
}

func (m *VoidRequest) Reset()                    { *m = VoidRequest{} }
func (m *VoidRequest) String() string            { return proto.CompactTextString(m) }
func (*VoidRequest) ProtoMessage()               {}
func (*VoidRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type VoidResponse struct {
}

func (m *VoidResponse) Reset()                    { *m = VoidResponse{} }
func (m *VoidResponse) String() string            { return proto.CompactTextString(m) }
func (*VoidResponse) ProtoMessage()               {}
func (*VoidResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// Response/Output only
type Uid struct {
	Phid string `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *Uid) Reset()                    { *m = Uid{} }
func (m *Uid) String() string            { return proto.CompactTextString(m) }
func (*Uid) ProtoMessage()               {}
func (*Uid) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Uid) GetPhid() string {
	if m != nil {
		return m.Phid
	}
	return ""
}

func (m *Uid) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*ErrorDetails)(nil), "appscode.dtypes.ErrorDetails")
	proto.RegisterType((*LongRunningResponse)(nil), "appscode.dtypes.LongRunningResponse")
	proto.RegisterType((*VoidRequest)(nil), "appscode.dtypes.VoidRequest")
	proto.RegisterType((*VoidResponse)(nil), "appscode.dtypes.VoidResponse")
	proto.RegisterType((*Uid)(nil), "appscode.dtypes.Uid")
}

func init() { proto.RegisterFile("types.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xc1, 0x4e, 0x03, 0x21,
	0x10, 0x86, 0xb3, 0x6a, 0xaa, 0x9d, 0x56, 0x8d, 0x78, 0xb0, 0xf5, 0x60, 0x0c, 0x27, 0x2f, 0xdd,
	0x9a, 0xf8, 0x06, 0x8d, 0xde, 0x3c, 0x6c, 0x88, 0x7a, 0x30, 0x31, 0x0d, 0x0b, 0xe3, 0x96, 0xda,
	0x32, 0x08, 0xec, 0x61, 0xdf, 0xde, 0x14, 0xd6, 0xa6, 0xb7, 0xe1, 0xfb, 0xe7, 0x07, 0xf2, 0xc1,
	0x28, 0x76, 0x0e, 0x43, 0xe9, 0x3c, 0x45, 0x62, 0x97, 0xd2, 0xb9, 0xa0, 0x48, 0x63, 0xa9, 0x13,
	0xbe, 0x9d, 0x36, 0x44, 0xcd, 0x06, 0xe7, 0x29, 0xae, 0xdb, 0xef, 0xb9, 0xb4, 0x5d, 0xde, 0xe5,
	0x5f, 0x30, 0x7e, 0xf1, 0x9e, 0xfc, 0x33, 0x46, 0x69, 0x36, 0x81, 0xcd, 0x80, 0x79, 0xfc, 0x6d,
	0x31, 0x44, 0xd4, 0x4b, 0x8f, 0x81, 0x5a, 0xaf, 0x70, 0x52, 0xdc, 0x17, 0x0f, 0x43, 0x71, 0xb5,
	0x4f, 0x44, 0x1f, 0xb0, 0x3b, 0x80, 0x10, 0xa5, 0xfa, 0x89, 0x5e, 0x2a, 0x9c, 0x1c, 0xa5, 0xb5,
	0x03, 0xc2, 0x1f, 0xe1, 0xfa, 0x95, 0x6c, 0x23, 0x5a, 0x6b, 0x8d, 0x6d, 0x04, 0x06, 0x47, 0x36,
	0x20, 0x9b, 0xc2, 0xd9, 0x9a, 0xea, 0xa5, 0x5b, 0x19, 0xdd, 0xdf, 0x7d, 0xba, 0xa6, 0xba, 0x5a,
	0x19, 0xcd, 0xcf, 0x61, 0xf4, 0x41, 0x46, 0x8b, 0xfc, 0x14, 0xbf, 0x80, 0x71, 0x3e, 0xe6, 0x26,
	0x9f, 0xc1, 0xf1, 0xbb, 0xd1, 0x8c, 0xc1, 0xc9, 0x41, 0x39, 0xcd, 0x3b, 0x66, 0xe5, 0xf6, 0xff,
	0x17, 0x69, 0x5e, 0x94, 0x70, 0xa3, 0x68, 0x5b, 0xee, 0x85, 0x48, 0x67, 0x7a, 0x29, 0x8b, 0xe1,
	0x5b, 0xe7, 0xb0, 0xda, 0x49, 0xa8, 0x8a, 0xcf, 0x41, 0x86, 0xf5, 0x20, 0x59, 0x79, 0xfa, 0x0b,
	0x00, 0x00, 0xff, 0xff, 0x02, 0x44, 0x3d, 0xb9, 0x50, 0x01, 0x00, 0x00,
}
