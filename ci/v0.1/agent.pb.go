// Code generated by protoc-gen-go.
// source: agent.proto
// DO NOT EDIT!

/*
Package ci is a generated protocol buffer package.

It is generated from these files:
	agent.proto
	build.proto
	job.proto
	master.proto

It has these top-level messages:
	AgentListRequest
	AgentListResponse
	Agent
	AgentCreateRequest
	PortInfo
	AgentDescribeRequest
	AgentDescribeResponse
	AgentDeleteRequest
	AgentRestartRequest
	AgentRestartResponse
	BuildDescribeRequest
	BuildDescribeResponse
	BuildListRequest
	BuildListResponse
	Build
	JobListRequest
	JobListResponse
	JobBuildRequest
	JobDescribeRequest
	JobDescribeResponse
	JobHealth
	JobDeleteRequest
	JobCreateRequest
	JobCopyRequest
	Job
	MasterCreateRequest
	MasterDeleteRequest
*/
package ci

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

type AgentListRequest struct {
	// List of status to get the agent filterd on the status
	// values in
	//   PENDING
	//   FAILED
	//   ONLINE
	//   OFFLINE
	//   DELETED
	Status []string `protobuf:"bytes,1,rep,name=status" json:"status,omitempty"`
}

func (m *AgentListRequest) Reset()                    { *m = AgentListRequest{} }
func (m *AgentListRequest) String() string            { return proto.CompactTextString(m) }
func (*AgentListRequest) ProtoMessage()               {}
func (*AgentListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type AgentListResponse struct {
	Status *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Agents []*Agent       `protobuf:"bytes,2,rep,name=agents" json:"agents,omitempty"`
}

func (m *AgentListResponse) Reset()                    { *m = AgentListResponse{} }
func (m *AgentListResponse) String() string            { return proto.CompactTextString(m) }
func (*AgentListResponse) ProtoMessage()               {}
func (*AgentListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AgentListResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *AgentListResponse) GetAgents() []*Agent {
	if m != nil {
		return m.Agents
	}
	return nil
}

type Agent struct {
	Name        string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Status      string `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
	AgentStatus string `protobuf:"bytes,3,opt,name=agent_status,json=agentStatus" json:"agent_status,omitempty"`
	IsRefreshed int32  `protobuf:"varint,4,opt,name=is_refreshed,json=isRefreshed" json:"is_refreshed,omitempty"`
	CreatedAt   string `protobuf:"bytes,5,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	UpdatedAt   string `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
}

func (m *Agent) Reset()                    { *m = Agent{} }
func (m *Agent) String() string            { return proto.CompactTextString(m) }
func (*Agent) ProtoMessage()               {}
func (*Agent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type AgentCreateRequest struct {
	Sku               string      `protobuf:"bytes,1,opt,name=sku" json:"sku,omitempty"`
	Executors         int32       `protobuf:"varint,2,opt,name=executors" json:"executors,omitempty"`
	Labels            string      `protobuf:"bytes,3,opt,name=labels" json:"labels,omitempty"`
	UserStartupScript string      `protobuf:"bytes,4,opt,name=user_startup_script,json=userStartupScript" json:"user_startup_script,omitempty"`
	SaltbaseVersion   string      `protobuf:"bytes,5,opt,name=saltbase_version,json=saltbaseVersion" json:"saltbase_version,omitempty"`
	CiStarterVersion  string      `protobuf:"bytes,6,opt,name=ci_starter_version,json=ciStarterVersion" json:"ci_starter_version,omitempty"`
	Ports             []*PortInfo `protobuf:"bytes,7,rep,name=ports" json:"ports,omitempty"`
	Role              string      `protobuf:"bytes,8,opt,name=role" json:"role,omitempty"`
}

func (m *AgentCreateRequest) Reset()                    { *m = AgentCreateRequest{} }
func (m *AgentCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*AgentCreateRequest) ProtoMessage()               {}
func (*AgentCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AgentCreateRequest) GetPorts() []*PortInfo {
	if m != nil {
		return m.Ports
	}
	return nil
}

type PortInfo struct {
	Protocol  string `protobuf:"bytes,1,opt,name=protocol" json:"protocol,omitempty"`
	PortRange string `protobuf:"bytes,2,opt,name=port_range,json=portRange" json:"port_range,omitempty"`
}

func (m *PortInfo) Reset()                    { *m = PortInfo{} }
func (m *PortInfo) String() string            { return proto.CompactTextString(m) }
func (*PortInfo) ProtoMessage()               {}
func (*PortInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type AgentDescribeRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *AgentDescribeRequest) Reset()                    { *m = AgentDescribeRequest{} }
func (m *AgentDescribeRequest) String() string            { return proto.CompactTextString(m) }
func (*AgentDescribeRequest) ProtoMessage()               {}
func (*AgentDescribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type AgentDescribeResponse struct {
	Status           *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Name             string         `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Executors        int64          `protobuf:"varint,3,opt,name=executors" json:"executors,omitempty"`
	AgentStatus      string         `protobuf:"bytes,4,opt,name=agent_status,json=agentStatus" json:"agent_status,omitempty"`
	AgentStatusCause string         `protobuf:"bytes,5,opt,name=agent_status_cause,json=agentStatusCause" json:"agent_status_cause,omitempty"`
	Label            string         `protobuf:"bytes,6,opt,name=label" json:"label,omitempty"`
	Provider         string         `protobuf:"bytes,7,opt,name=provider" json:"provider,omitempty"`
	Sku              string         `protobuf:"bytes,8,opt,name=sku" json:"sku,omitempty"`
	StartupScript    string         `protobuf:"bytes,9,opt,name=startup_script,json=startupScript" json:"startup_script,omitempty"`
	CreatedAt        string         `protobuf:"bytes,10,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	UpdatedAt        string         `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
}

func (m *AgentDescribeResponse) Reset()                    { *m = AgentDescribeResponse{} }
func (m *AgentDescribeResponse) String() string            { return proto.CompactTextString(m) }
func (*AgentDescribeResponse) ProtoMessage()               {}
func (*AgentDescribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *AgentDescribeResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type AgentDeleteRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *AgentDeleteRequest) Reset()                    { *m = AgentDeleteRequest{} }
func (m *AgentDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*AgentDeleteRequest) ProtoMessage()               {}
func (*AgentDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type AgentRestartRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *AgentRestartRequest) Reset()                    { *m = AgentRestartRequest{} }
func (m *AgentRestartRequest) String() string            { return proto.CompactTextString(m) }
func (*AgentRestartRequest) ProtoMessage()               {}
func (*AgentRestartRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type AgentRestartResponse struct {
	Status *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *AgentRestartResponse) Reset()                    { *m = AgentRestartResponse{} }
func (m *AgentRestartResponse) String() string            { return proto.CompactTextString(m) }
func (*AgentRestartResponse) ProtoMessage()               {}
func (*AgentRestartResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *AgentRestartResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*AgentListRequest)(nil), "ci.AgentListRequest")
	proto.RegisterType((*AgentListResponse)(nil), "ci.AgentListResponse")
	proto.RegisterType((*Agent)(nil), "ci.Agent")
	proto.RegisterType((*AgentCreateRequest)(nil), "ci.AgentCreateRequest")
	proto.RegisterType((*PortInfo)(nil), "ci.PortInfo")
	proto.RegisterType((*AgentDescribeRequest)(nil), "ci.AgentDescribeRequest")
	proto.RegisterType((*AgentDescribeResponse)(nil), "ci.AgentDescribeResponse")
	proto.RegisterType((*AgentDeleteRequest)(nil), "ci.AgentDeleteRequest")
	proto.RegisterType((*AgentRestartRequest)(nil), "ci.AgentRestartRequest")
	proto.RegisterType((*AgentRestartResponse)(nil), "ci.AgentRestartResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Agents service

type AgentsClient interface {
	List(ctx context.Context, in *AgentListRequest, opts ...grpc.CallOption) (*AgentListResponse, error)
	Describe(ctx context.Context, in *AgentDescribeRequest, opts ...grpc.CallOption) (*AgentDescribeResponse, error)
	Create(ctx context.Context, in *AgentCreateRequest, opts ...grpc.CallOption) (*dtypes.LongRunningResponse, error)
	Delete(ctx context.Context, in *AgentDeleteRequest, opts ...grpc.CallOption) (*dtypes.LongRunningResponse, error)
	Restart(ctx context.Context, in *AgentRestartRequest, opts ...grpc.CallOption) (*AgentRestartResponse, error)
}

type agentsClient struct {
	cc *grpc.ClientConn
}

func NewAgentsClient(cc *grpc.ClientConn) AgentsClient {
	return &agentsClient{cc}
}

func (c *agentsClient) List(ctx context.Context, in *AgentListRequest, opts ...grpc.CallOption) (*AgentListResponse, error) {
	out := new(AgentListResponse)
	err := grpc.Invoke(ctx, "/ci.Agents/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentsClient) Describe(ctx context.Context, in *AgentDescribeRequest, opts ...grpc.CallOption) (*AgentDescribeResponse, error) {
	out := new(AgentDescribeResponse)
	err := grpc.Invoke(ctx, "/ci.Agents/Describe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentsClient) Create(ctx context.Context, in *AgentCreateRequest, opts ...grpc.CallOption) (*dtypes.LongRunningResponse, error) {
	out := new(dtypes.LongRunningResponse)
	err := grpc.Invoke(ctx, "/ci.Agents/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentsClient) Delete(ctx context.Context, in *AgentDeleteRequest, opts ...grpc.CallOption) (*dtypes.LongRunningResponse, error) {
	out := new(dtypes.LongRunningResponse)
	err := grpc.Invoke(ctx, "/ci.Agents/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentsClient) Restart(ctx context.Context, in *AgentRestartRequest, opts ...grpc.CallOption) (*AgentRestartResponse, error) {
	out := new(AgentRestartResponse)
	err := grpc.Invoke(ctx, "/ci.Agents/Restart", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Agents service

type AgentsServer interface {
	List(context.Context, *AgentListRequest) (*AgentListResponse, error)
	Describe(context.Context, *AgentDescribeRequest) (*AgentDescribeResponse, error)
	Create(context.Context, *AgentCreateRequest) (*dtypes.LongRunningResponse, error)
	Delete(context.Context, *AgentDeleteRequest) (*dtypes.LongRunningResponse, error)
	Restart(context.Context, *AgentRestartRequest) (*AgentRestartResponse, error)
}

func RegisterAgentsServer(s *grpc.Server, srv AgentsServer) {
	s.RegisterService(&_Agents_serviceDesc, srv)
}

func _Agents_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ci.Agents/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentsServer).List(ctx, req.(*AgentListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agents_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentDescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentsServer).Describe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ci.Agents/Describe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentsServer).Describe(ctx, req.(*AgentDescribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agents_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ci.Agents/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentsServer).Create(ctx, req.(*AgentCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agents_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ci.Agents/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentsServer).Delete(ctx, req.(*AgentDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agents_Restart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentRestartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentsServer).Restart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ci.Agents/Restart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentsServer).Restart(ctx, req.(*AgentRestartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Agents_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ci.Agents",
	HandlerType: (*AgentsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Agents_List_Handler,
		},
		{
			MethodName: "Describe",
			Handler:    _Agents_Describe_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Agents_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Agents_Delete_Handler,
		},
		{
			MethodName: "Restart",
			Handler:    _Agents_Restart_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("agent.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 758 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x55, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0x55, 0x7e, 0x1b, 0xdf, 0xf4, 0xeb, 0x97, 0x4e, 0x7f, 0x30, 0xa6, 0xa0, 0xd6, 0x50, 0x48,
	0x43, 0x15, 0x43, 0xd8, 0xb1, 0x40, 0xaa, 0x5a, 0x16, 0x48, 0x5d, 0x20, 0x47, 0x42, 0x62, 0x43,
	0xe4, 0x38, 0xb7, 0xc1, 0x60, 0x3c, 0x66, 0x66, 0x5c, 0x81, 0x10, 0x1b, 0x5e, 0x81, 0x05, 0x2b,
	0x9e, 0x83, 0x07, 0xe1, 0x15, 0xd8, 0xf1, 0x12, 0x68, 0x7e, 0xec, 0x38, 0x69, 0x15, 0xb5, 0x9b,
	0xc8, 0x73, 0xce, 0xf5, 0xbd, 0xbe, 0xf7, 0x9c, 0xb9, 0x81, 0x76, 0x30, 0xc5, 0x44, 0xf4, 0x53,
	0x46, 0x05, 0x25, 0xd5, 0x30, 0x72, 0x76, 0xa6, 0x94, 0x4e, 0x63, 0xf4, 0x82, 0x34, 0xf2, 0x82,
	0x24, 0xa1, 0x22, 0x10, 0x11, 0x4d, 0xb8, 0x8e, 0x70, 0xb6, 0x25, 0x3c, 0x11, 0x9f, 0x53, 0xe4,
	0x9e, 0xfa, 0xd5, 0xb8, 0xdb, 0x83, 0xce, 0x91, 0x4c, 0x74, 0x1a, 0x71, 0xe1, 0xe3, 0xc7, 0x0c,
	0xb9, 0x20, 0xdb, 0xd0, 0xe4, 0x22, 0x10, 0x19, 0xb7, 0x2b, 0xbb, 0xb5, 0xae, 0xe5, 0x9b, 0x93,
	0xfb, 0x06, 0xd6, 0x4b, 0xb1, 0x3c, 0xa5, 0x09, 0x47, 0x72, 0xbf, 0x14, 0x5c, 0xe9, 0xb6, 0x07,
	0x6b, 0x7d, 0x5d, 0xa5, 0x3f, 0x54, 0x68, 0xfe, 0x32, 0xd9, 0x83, 0xa6, 0xfa, 0x62, 0x6e, 0x57,
	0x77, 0x6b, 0xdd, 0xf6, 0xc0, 0xea, 0x87, 0x51, 0x5f, 0xa5, 0xf3, 0x0d, 0xe1, 0xfe, 0xaa, 0x40,
	0x43, 0x21, 0x84, 0x40, 0x3d, 0x09, 0x3e, 0xa0, 0x4a, 0x69, 0xf9, 0xea, 0xb9, 0xf4, 0x55, 0x55,
	0x85, 0xce, 0x12, 0xaf, 0xaa, 0xf7, 0x47, 0x86, 0xad, 0x29, 0x56, 0x8f, 0x67, 0x58, 0x84, 0x44,
	0x7c, 0xc4, 0xf0, 0x8c, 0x21, 0x7f, 0x8b, 0x13, 0xbb, 0xbe, 0x5b, 0xe9, 0x36, 0xfc, 0x76, 0xc4,
	0xfd, 0x1c, 0x22, 0xb7, 0x01, 0x42, 0x86, 0x81, 0xc0, 0xc9, 0x28, 0x10, 0x76, 0x43, 0xe5, 0xb0,
	0x0c, 0x72, 0x24, 0x24, 0x9d, 0xa5, 0x93, 0x9c, 0x6e, 0x6a, 0xda, 0x20, 0x47, 0xc2, 0xfd, 0x59,
	0x05, 0xa2, 0xbe, 0xfc, 0x58, 0xbd, 0x91, 0x0f, 0xb2, 0x03, 0x35, 0xfe, 0x3e, 0x33, 0x5d, 0xc8,
	0x47, 0xb2, 0x03, 0x16, 0x7e, 0xc2, 0x30, 0x13, 0x94, 0xe9, 0x3e, 0x1a, 0xfe, 0x0c, 0x90, 0x2d,
	0xc6, 0xc1, 0x18, 0xe3, 0xbc, 0x09, 0x73, 0x22, 0x7d, 0xd8, 0xc8, 0x38, 0x32, 0xd9, 0x21, 0x13,
	0x59, 0x3a, 0xe2, 0x21, 0x8b, 0x52, 0xa1, 0xda, 0xb0, 0xfc, 0x75, 0x49, 0x0d, 0x35, 0x33, 0x54,
	0x04, 0x39, 0x80, 0x0e, 0x0f, 0x62, 0x31, 0x0e, 0x38, 0x8e, 0xce, 0x91, 0xf1, 0x88, 0x26, 0xa6,
	0xa5, 0xff, 0x73, 0xfc, 0x95, 0x86, 0xc9, 0x21, 0x90, 0x30, 0xd2, 0x89, 0x91, 0x15, 0xc1, 0xba,
	0xc1, 0x4e, 0x18, 0x0d, 0x35, 0x91, 0x47, 0xbb, 0xd0, 0x48, 0x29, 0x13, 0xdc, 0x5e, 0x51, 0x1a,
	0xae, 0x4a, 0x0d, 0x5f, 0x52, 0x26, 0x5e, 0x24, 0x67, 0xd4, 0xd7, 0x94, 0xd4, 0x8e, 0xd1, 0x18,
	0xed, 0x96, 0xd6, 0x4e, 0x3e, 0xbb, 0xcf, 0xa1, 0x95, 0x87, 0x11, 0x07, 0x5a, 0xca, 0x7a, 0x21,
	0x8d, 0xcd, 0x64, 0x8a, 0xb3, 0x1c, 0xb3, 0x4c, 0x32, 0x62, 0x41, 0x32, 0x45, 0xa3, 0xb3, 0x25,
	0x11, 0x5f, 0x02, 0x6e, 0x0f, 0x36, 0xd5, 0x94, 0x4f, 0x50, 0x4e, 0x60, 0x5c, 0xcc, 0xf9, 0x12,
	0xbb, 0xb8, 0x7f, 0xab, 0xb0, 0xb5, 0x10, 0x7c, 0x4d, 0xc7, 0xe6, 0x59, 0xab, 0x25, 0x13, 0xce,
	0xe9, 0x27, 0x45, 0xaa, 0x95, 0xf5, 0x5b, 0xb4, 0x62, 0xfd, 0xa2, 0x15, 0x0f, 0x81, 0x94, 0x43,
	0x46, 0x61, 0x90, 0x71, 0x34, 0xe2, 0x74, 0x4a, 0x81, 0xc7, 0x12, 0x27, 0x9b, 0xd0, 0x50, 0x16,
	0x30, 0x82, 0xe8, 0x83, 0x99, 0xe0, 0x79, 0x34, 0x41, 0x66, 0xaf, 0x14, 0x13, 0x54, 0xe7, 0xdc,
	0x72, 0xad, 0x99, 0xe5, 0xf6, 0x61, 0x6d, 0xc1, 0x37, 0x96, 0x22, 0xff, 0xe3, 0x73, 0x9e, 0x99,
	0xbf, 0x00, 0xb0, 0xfc, 0x02, 0xb4, 0x17, 0x2f, 0x40, 0xd7, 0xf8, 0xff, 0x04, 0x63, 0x14, 0x4b,
	0x75, 0x39, 0x80, 0x0d, 0x7d, 0xeb, 0x51, 0xd5, 0x5f, 0x16, 0xfa, 0xcc, 0xc8, 0x5d, 0x84, 0x5e,
	0x4f, 0xc0, 0xc1, 0x8f, 0x3a, 0x34, 0x55, 0x02, 0x4e, 0x5e, 0x43, 0x5d, 0x6e, 0x2d, 0xb2, 0x59,
	0x6c, 0x9d, 0xd2, 0xc2, 0x73, 0xb6, 0x16, 0x50, 0x5d, 0xc7, 0xbd, 0xf7, 0xed, 0xf7, 0x9f, 0xef,
	0xd5, 0x3b, 0x64, 0xc7, 0x0b, 0xd2, 0x94, 0x87, 0x74, 0xa2, 0x97, 0x6b, 0x18, 0x79, 0xe7, 0x8f,
	0xfa, 0x8f, 0x3d, 0xbd, 0xb5, 0x48, 0x0c, 0xad, 0xdc, 0x62, 0xc4, 0x2e, 0x12, 0x2d, 0x58, 0xd4,
	0xb9, 0x79, 0x09, 0x63, 0xca, 0x3c, 0x54, 0x65, 0xf6, 0xc9, 0xdd, 0x65, 0x65, 0xbc, 0x2f, 0x72,
	0x24, 0x5f, 0xc9, 0x19, 0x34, 0xf5, 0x8e, 0x21, 0xdb, 0x45, 0xc6, 0xb9, 0xa5, 0xe3, 0xdc, 0xca,
	0xa7, 0x71, 0x4a, 0x93, 0xa9, 0x9f, 0x25, 0x49, 0x94, 0x4c, 0x8b, 0x5a, 0x0f, 0x54, 0xad, 0x3d,
	0x67, 0x69, 0x4b, 0x4f, 0x2b, 0x3d, 0xf2, 0x0e, 0x9a, 0x5a, 0xcb, 0x52, 0x9d, 0x39, 0x71, 0x97,
	0xd7, 0x31, 0x3d, 0xf5, 0xae, 0xd4, 0x53, 0x0a, 0x2b, 0x46, 0x62, 0x72, 0x63, 0xf6, 0xaf, 0x30,
	0xe7, 0x0f, 0xc7, 0xbe, 0x48, 0x98, 0x52, 0x03, 0x55, 0xea, 0x90, 0xf4, 0xae, 0x50, 0xca, 0x63,
	0x38, 0xa6, 0x54, 0x8c, 0x9b, 0x6a, 0xe3, 0x3c, 0xf9, 0x17, 0x00, 0x00, 0xff, 0xff, 0x77, 0x5d,
	0x41, 0x05, 0x45, 0x07, 0x00, 0x00,
}
