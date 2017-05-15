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
	Cause      string                   `protobuf:"bytes,1,opt,name=cause" json:"cause,omitempty"`
	StackTrace *ErrorDetails_StackTrace `protobuf:"bytes,2,opt,name=stack_trace,json=stackTrace" json:"stack_trace,omitempty"`
}

func (m *ErrorDetails) Reset()                    { *m = ErrorDetails{} }
func (m *ErrorDetails) String() string            { return proto.CompactTextString(m) }
func (*ErrorDetails) ProtoMessage()               {}
func (*ErrorDetails) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ErrorDetails) GetCause() string {
	if m != nil {
		return m.Cause
	}
	return ""
}

func (m *ErrorDetails) GetStackTrace() *ErrorDetails_StackTrace {
	if m != nil {
		return m.StackTrace
	}
	return nil
}

type ErrorDetails_StackTrace struct {
	Frames []string `protobuf:"bytes,1,rep,name=frames" json:"frames,omitempty"`
}

func (m *ErrorDetails_StackTrace) Reset()                    { *m = ErrorDetails_StackTrace{} }
func (m *ErrorDetails_StackTrace) String() string            { return proto.CompactTextString(m) }
func (*ErrorDetails_StackTrace) ProtoMessage()               {}
func (*ErrorDetails_StackTrace) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *ErrorDetails_StackTrace) GetFrames() []string {
	if m != nil {
		return m.Frames
	}
	return nil
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
	proto.RegisterType((*ErrorDetails_StackTrace)(nil), "appscode.dtypes.ErrorDetails.StackTrace")
	proto.RegisterType((*LongRunningResponse)(nil), "appscode.dtypes.LongRunningResponse")
	proto.RegisterType((*VoidRequest)(nil), "appscode.dtypes.VoidRequest")
	proto.RegisterType((*VoidResponse)(nil), "appscode.dtypes.VoidResponse")
	proto.RegisterType((*Uid)(nil), "appscode.dtypes.Uid")
}

func init() { proto.RegisterFile("types.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x50, 0xbf, 0x4e, 0xfb, 0x30,
	0x18, 0x94, 0x7f, 0xfd, 0x11, 0xc8, 0x97, 0x02, 0x92, 0x41, 0x10, 0x98, 0xa2, 0x88, 0x21, 0x0b,
	0x16, 0x82, 0x37, 0xa8, 0x60, 0x40, 0x62, 0x88, 0x4c, 0x61, 0x60, 0xa9, 0x9c, 0xf8, 0xa3, 0x75,
	0x21, 0xb6, 0xb1, 0x9d, 0xa1, 0x0f, 0xc2, 0xfb, 0xa2, 0xfc, 0x11, 0xad, 0xd8, 0xee, 0xce, 0xbe,
	0xd3, 0x77, 0x07, 0x49, 0xd8, 0x58, 0xf4, 0xcc, 0x3a, 0x13, 0x0c, 0x3d, 0x16, 0xd6, 0xfa, 0xda,
	0x48, 0x64, 0xb2, 0x97, 0xf3, 0x6f, 0x02, 0xd3, 0x07, 0xe7, 0x8c, 0xbb, 0xc7, 0x20, 0xd4, 0xa7,
	0xa7, 0xa7, 0xb0, 0x57, 0x8b, 0xd6, 0x63, 0x4a, 0x32, 0x52, 0xc4, 0x7c, 0x20, 0xf4, 0x11, 0x12,
	0x1f, 0x44, 0xfd, 0xb1, 0x08, 0x4e, 0xd4, 0x98, 0xfe, 0xcb, 0x48, 0x91, 0xdc, 0x16, 0xec, 0x4f,
	0x1a, 0xdb, 0x4d, 0x62, 0xcf, 0x9d, 0x61, 0xde, 0xfd, 0xe7, 0xe0, 0x7f, 0xf1, 0xe5, 0x15, 0xc0,
	0xf6, 0x85, 0x9e, 0x41, 0xf4, 0xee, 0x44, 0x83, 0x3e, 0x25, 0xd9, 0xa4, 0x88, 0xf9, 0xc8, 0xf2,
	0x1b, 0x38, 0x79, 0x32, 0x7a, 0xc9, 0x5b, 0xad, 0x95, 0x5e, 0x72, 0xf4, 0xd6, 0x68, 0x8f, 0xf4,
	0x02, 0x0e, 0xd6, 0xa6, 0x5a, 0xd8, 0x95, 0x92, 0xe3, 0x81, 0xfb, 0x6b, 0x53, 0x95, 0x2b, 0x25,
	0xf3, 0x43, 0x48, 0x5e, 0x8d, 0x92, 0x1c, 0xbf, 0x5a, 0xf4, 0x21, 0x3f, 0x82, 0xe9, 0x40, 0x07,
	0x67, 0x7e, 0x0d, 0x93, 0x17, 0x25, 0x29, 0x85, 0xff, 0x3b, 0xe6, 0x1e, 0x77, 0x9a, 0x16, 0xcd,
	0xd0, 0x2a, 0xe6, 0x3d, 0x9e, 0x31, 0x38, 0xaf, 0x4d, 0xb3, 0x2d, 0x28, 0xac, 0x1a, 0x4b, 0xce,
	0xe2, 0xf9, 0xc6, 0x62, 0xd9, 0xcd, 0x59, 0x92, 0xb7, 0x68, 0x10, 0xab, 0xa8, 0xdf, 0xf7, 0xee,
	0x27, 0x00, 0x00, 0xff, 0xff, 0x73, 0x2d, 0x5a, 0x18, 0x6e, 0x01, 0x00, 0x00,
}
