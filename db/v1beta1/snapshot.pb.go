// Code generated by protoc-gen-go.
// source: snapshot.proto
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

// Next Id: 4
type DatabaseInfo struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Type    string `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
}

func (m *DatabaseInfo) Reset()                    { *m = DatabaseInfo{} }
func (m *DatabaseInfo) String() string            { return proto.CompactTextString(m) }
func (*DatabaseInfo) ProtoMessage()               {}
func (*DatabaseInfo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

// Next Id: 15
type Snapshot struct {
	Name              string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Phid              string `protobuf:"bytes,2,opt,name=phid" json:"phid,omitempty"`
	Provider          string `protobuf:"bytes,3,opt,name=provider" json:"provider,omitempty"`
	CloudCredential   string `protobuf:"bytes,4,opt,name=cloud_credential,json=cloudCredential" json:"cloud_credential,omitempty"`
	Bucket            string `protobuf:"bytes,5,opt,name=bucket" json:"bucket,omitempty"`
	Status            string `protobuf:"bytes,8,opt,name=status" json:"status,omitempty"`
	Process           string `protobuf:"bytes,9,opt,name=process" json:"process,omitempty"`
	Type              string `protobuf:"bytes,10,opt,name=type" json:"type,omitempty"`
	IsScheduledBackup int32  `protobuf:"varint,11,opt,name=is_scheduled_backup,json=isScheduledBackup" json:"is_scheduled_backup,omitempty"`
	SourceDatabase    string `protobuf:"bytes,12,opt,name=source_database,json=sourceDatabase" json:"source_database,omitempty"`
	SourceSnapshot    string `protobuf:"bytes,13,opt,name=source_snapshot,json=sourceSnapshot" json:"source_snapshot,omitempty"`
	CreatedAt         int64  `protobuf:"varint,14,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
}

func (m *Snapshot) Reset()                    { *m = Snapshot{} }
func (m *Snapshot) String() string            { return proto.CompactTextString(m) }
func (*Snapshot) ProtoMessage()               {}
func (*Snapshot) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

// Next Id: 3
type SnapshotListRequest struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Uid     string `protobuf:"bytes,2,opt,name=uid" json:"uid,omitempty"`
}

func (m *SnapshotListRequest) Reset()                    { *m = SnapshotListRequest{} }
func (m *SnapshotListRequest) String() string            { return proto.CompactTextString(m) }
func (*SnapshotListRequest) ProtoMessage()               {}
func (*SnapshotListRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

// Next Id: 4
type SnapshotListResponse struct {
	Status       *appscode_dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	DatabaseInfo *DatabaseInfo           `protobuf:"bytes,2,opt,name=database_info,json=databaseInfo" json:"database_info,omitempty"`
	List         []*Snapshot             `protobuf:"bytes,3,rep,name=list" json:"list,omitempty"`
}

func (m *SnapshotListResponse) Reset()                    { *m = SnapshotListResponse{} }
func (m *SnapshotListResponse) String() string            { return proto.CompactTextString(m) }
func (*SnapshotListResponse) ProtoMessage()               {}
func (*SnapshotListResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *SnapshotListResponse) GetStatus() *appscode_dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *SnapshotListResponse) GetDatabaseInfo() *DatabaseInfo {
	if m != nil {
		return m.DatabaseInfo
	}
	return nil
}

func (m *SnapshotListResponse) GetList() []*Snapshot {
	if m != nil {
		return m.List
	}
	return nil
}

// Next Id: 13
type BackupScheduleRequest struct {
	Cluster          string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Uid              string `protobuf:"bytes,2,opt,name=uid" json:"uid,omitempty"`
	CredentialUid    string `protobuf:"bytes,4,opt,name=credential_uid,json=credentialUid" json:"credential_uid,omitempty"`
	BucketName       string `protobuf:"bytes,5,opt,name=bucket_name,json=bucketName" json:"bucket_name,omitempty"`
	SnapshotName     string `protobuf:"bytes,6,opt,name=snapshot_name,json=snapshotName" json:"snapshot_name,omitempty"`
	AuthSecretName   string `protobuf:"bytes,8,opt,name=auth_secret_name,json=authSecretName" json:"auth_secret_name,omitempty"`
	ScheduleCronExpr string `protobuf:"bytes,9,opt,name=schedule_cron_expr,json=scheduleCronExpr" json:"schedule_cron_expr,omitempty"`
}

func (m *BackupScheduleRequest) Reset()                    { *m = BackupScheduleRequest{} }
func (m *BackupScheduleRequest) String() string            { return proto.CompactTextString(m) }
func (*BackupScheduleRequest) ProtoMessage()               {}
func (*BackupScheduleRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

// Next Id: 3
type BackupUnscheduleRequest struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Uid     string `protobuf:"bytes,2,opt,name=uid" json:"uid,omitempty"`
}

func (m *BackupUnscheduleRequest) Reset()                    { *m = BackupUnscheduleRequest{} }
func (m *BackupUnscheduleRequest) String() string            { return proto.CompactTextString(m) }
func (*BackupUnscheduleRequest) ProtoMessage()               {}
func (*BackupUnscheduleRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

// Next Id: 19
type SnapshotRestoreRequest struct {
	Cluster        string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Uid            string `protobuf:"bytes,2,opt,name=uid" json:"uid,omitempty"`
	SnapshotPhid   string `protobuf:"bytes,8,opt,name=snapshot_phid,json=snapshotPhid" json:"snapshot_phid,omitempty"`
	AuthSecretName string `protobuf:"bytes,13,opt,name=auth_secret_name,json=authSecretName" json:"auth_secret_name,omitempty"`
}

func (m *SnapshotRestoreRequest) Reset()                    { *m = SnapshotRestoreRequest{} }
func (m *SnapshotRestoreRequest) String() string            { return proto.CompactTextString(m) }
func (*SnapshotRestoreRequest) ProtoMessage()               {}
func (*SnapshotRestoreRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

// Next Id: 4
type SnapshotCheckRequest struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Phid    string `protobuf:"bytes,2,opt,name=phid" json:"phid,omitempty"`
}

func (m *SnapshotCheckRequest) Reset()                    { *m = SnapshotCheckRequest{} }
func (m *SnapshotCheckRequest) String() string            { return proto.CompactTextString(m) }
func (*SnapshotCheckRequest) ProtoMessage()               {}
func (*SnapshotCheckRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func init() {
	proto.RegisterType((*DatabaseInfo)(nil), "appscode.db.v1beta1.DatabaseInfo")
	proto.RegisterType((*Snapshot)(nil), "appscode.db.v1beta1.Snapshot")
	proto.RegisterType((*SnapshotListRequest)(nil), "appscode.db.v1beta1.SnapshotListRequest")
	proto.RegisterType((*SnapshotListResponse)(nil), "appscode.db.v1beta1.SnapshotListResponse")
	proto.RegisterType((*BackupScheduleRequest)(nil), "appscode.db.v1beta1.BackupScheduleRequest")
	proto.RegisterType((*BackupUnscheduleRequest)(nil), "appscode.db.v1beta1.BackupUnscheduleRequest")
	proto.RegisterType((*SnapshotRestoreRequest)(nil), "appscode.db.v1beta1.SnapshotRestoreRequest")
	proto.RegisterType((*SnapshotCheckRequest)(nil), "appscode.db.v1beta1.SnapshotCheckRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Snapshots service

type SnapshotsClient interface {
	List(ctx context.Context, in *SnapshotListRequest, opts ...grpc.CallOption) (*SnapshotListResponse, error)
	BackupSchedule(ctx context.Context, in *BackupScheduleRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
	BackupUnschedule(ctx context.Context, in *BackupUnscheduleRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
	Restore(ctx context.Context, in *SnapshotRestoreRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
}

type snapshotsClient struct {
	cc *grpc.ClientConn
}

func NewSnapshotsClient(cc *grpc.ClientConn) SnapshotsClient {
	return &snapshotsClient{cc}
}

func (c *snapshotsClient) List(ctx context.Context, in *SnapshotListRequest, opts ...grpc.CallOption) (*SnapshotListResponse, error) {
	out := new(SnapshotListResponse)
	err := grpc.Invoke(ctx, "/appscode.db.v1beta1.Snapshots/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snapshotsClient) BackupSchedule(ctx context.Context, in *BackupScheduleRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.db.v1beta1.Snapshots/BackupSchedule", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snapshotsClient) BackupUnschedule(ctx context.Context, in *BackupUnscheduleRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.db.v1beta1.Snapshots/BackupUnschedule", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snapshotsClient) Restore(ctx context.Context, in *SnapshotRestoreRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.db.v1beta1.Snapshots/Restore", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Snapshots service

type SnapshotsServer interface {
	List(context.Context, *SnapshotListRequest) (*SnapshotListResponse, error)
	BackupSchedule(context.Context, *BackupScheduleRequest) (*appscode_dtypes.VoidResponse, error)
	BackupUnschedule(context.Context, *BackupUnscheduleRequest) (*appscode_dtypes.VoidResponse, error)
	Restore(context.Context, *SnapshotRestoreRequest) (*appscode_dtypes.VoidResponse, error)
}

func RegisterSnapshotsServer(s *grpc.Server, srv SnapshotsServer) {
	s.RegisterService(&_Snapshots_serviceDesc, srv)
}

func _Snapshots_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SnapshotListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.db.v1beta1.Snapshots/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotsServer).List(ctx, req.(*SnapshotListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Snapshots_BackupSchedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BackupScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotsServer).BackupSchedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.db.v1beta1.Snapshots/BackupSchedule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotsServer).BackupSchedule(ctx, req.(*BackupScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Snapshots_BackupUnschedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BackupUnscheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotsServer).BackupUnschedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.db.v1beta1.Snapshots/BackupUnschedule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotsServer).BackupUnschedule(ctx, req.(*BackupUnscheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Snapshots_Restore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SnapshotRestoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotsServer).Restore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.db.v1beta1.Snapshots/Restore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotsServer).Restore(ctx, req.(*SnapshotRestoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Snapshots_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.db.v1beta1.Snapshots",
	HandlerType: (*SnapshotsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Snapshots_List_Handler,
		},
		{
			MethodName: "BackupSchedule",
			Handler:    _Snapshots_BackupSchedule_Handler,
		},
		{
			MethodName: "BackupUnschedule",
			Handler:    _Snapshots_BackupUnschedule_Handler,
		},
		{
			MethodName: "Restore",
			Handler:    _Snapshots_Restore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor1,
}

func init() { proto.RegisterFile("snapshot.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 837 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x56, 0xcd, 0x6e, 0x33, 0x35,
	0x14, 0xd5, 0x24, 0xf9, 0xd2, 0xe4, 0xe6, 0xe7, 0x0b, 0x2e, 0xf4, 0x1b, 0x05, 0x0a, 0x61, 0x10,
	0x22, 0x2d, 0x25, 0xa3, 0x86, 0x05, 0x52, 0x77, 0xfd, 0x43, 0x42, 0xe2, 0x27, 0x4a, 0x54, 0x84,
	0xa8, 0xd0, 0xe0, 0x8c, 0xdd, 0x66, 0x94, 0x74, 0x3c, 0xd8, 0x9e, 0xaa, 0xa8, 0xea, 0xa6, 0x2b,
	0xf6, 0x7d, 0x00, 0x9e, 0x02, 0x09, 0x89, 0x1d, 0x3b, 0x24, 0x76, 0xbc, 0x02, 0x0f, 0xc1, 0x12,
	0x8d, 0x3d, 0x9e, 0x4c, 0xda, 0x90, 0xa2, 0xaa, 0x9b, 0xca, 0x3e, 0xf7, 0xfa, 0xf6, 0x9e, 0xe3,
	0x73, 0x9d, 0x81, 0xa6, 0x08, 0x71, 0x24, 0x26, 0x4c, 0xf6, 0x22, 0xce, 0x24, 0x43, 0xeb, 0x38,
	0x8a, 0x84, 0xcf, 0x08, 0xed, 0x91, 0x71, 0xef, 0x72, 0x77, 0x4c, 0x25, 0xde, 0x6d, 0xbf, 0x75,
	0xce, 0xd8, 0xf9, 0x8c, 0xba, 0x38, 0x0a, 0x5c, 0x1c, 0x86, 0x4c, 0x62, 0x19, 0xb0, 0x50, 0xe8,
	0x23, 0xed, 0xb7, 0xcd, 0x91, 0xff, 0x88, 0x6f, 0x24, 0x30, 0x91, 0x3f, 0x46, 0x54, 0xb8, 0xea,
	0xaf, 0xc6, 0x9d, 0x01, 0xd4, 0x8f, 0xb0, 0xc4, 0x63, 0x2c, 0xe8, 0x67, 0xe1, 0x19, 0x43, 0x36,
	0xac, 0xf9, 0xb3, 0x58, 0x48, 0xca, 0x6d, 0xab, 0x63, 0x75, 0xab, 0x43, 0xb3, 0x45, 0x08, 0x4a,
	0x21, 0xbe, 0xa0, 0x76, 0x41, 0xc1, 0x6a, 0x9d, 0x60, 0x49, 0x31, 0xbb, 0xa8, 0xb1, 0x64, 0xed,
	0xfc, 0x53, 0x80, 0xca, 0x28, 0xe5, 0x93, 0x1d, 0xb2, 0x16, 0x0f, 0x45, 0x93, 0x80, 0x98, 0x42,
	0xc9, 0x1a, 0xb5, 0xa1, 0x12, 0x71, 0x76, 0x19, 0x10, 0xca, 0xd3, 0x62, 0xd9, 0x1e, 0x6d, 0x41,
	0xcb, 0x9f, 0xb1, 0x98, 0x78, 0x3e, 0xa7, 0x84, 0x86, 0x32, 0xc0, 0x33, 0xbb, 0xa4, 0x72, 0x5e,
	0x2a, 0xfc, 0x30, 0x83, 0xd1, 0x06, 0x94, 0xc7, 0xb1, 0x3f, 0xa5, 0xd2, 0x7e, 0xa1, 0x12, 0xd2,
	0x5d, 0x82, 0x0b, 0x89, 0x65, 0x2c, 0xec, 0x8a, 0xc6, 0xf5, 0x2e, 0x61, 0x1b, 0x71, 0xe6, 0x53,
	0x21, 0xec, 0xaa, 0x66, 0x9b, 0x6e, 0x33, 0x66, 0x30, 0x67, 0x86, 0x7a, 0xb0, 0x1e, 0x08, 0x4f,
	0xf8, 0x13, 0x4a, 0xe2, 0x19, 0x25, 0xde, 0x18, 0xfb, 0xd3, 0x38, 0xb2, 0x6b, 0x1d, 0xab, 0xfb,
	0x62, 0xf8, 0x5a, 0x20, 0x46, 0x26, 0x72, 0xa0, 0x02, 0xe8, 0x03, 0x78, 0x29, 0x58, 0xcc, 0x7d,
	0xea, 0x91, 0x54, 0x62, 0xbb, 0xae, 0xca, 0x35, 0x35, 0x6c, 0x84, 0xcf, 0x25, 0x1a, 0x23, 0xd8,
	0x8d, 0x7c, 0x62, 0x26, 0xe7, 0x26, 0x80, 0xcf, 0x29, 0x96, 0x94, 0x78, 0x58, 0xda, 0xcd, 0x8e,
	0xd5, 0x2d, 0x0e, 0xab, 0x29, 0xb2, 0x2f, 0x9d, 0x7d, 0x58, 0x37, 0xa9, 0x9f, 0x07, 0x42, 0x0e,
	0xe9, 0x0f, 0x31, 0x15, 0x72, 0xc5, 0x9d, 0xb6, 0xa0, 0x18, 0x67, 0x37, 0x91, 0x2c, 0x9d, 0xdf,
	0x2d, 0x78, 0x7d, 0xb1, 0x86, 0x88, 0x58, 0x28, 0x28, 0x72, 0x33, 0x09, 0x93, 0x1a, 0xb5, 0xfe,
	0xab, 0xde, 0xdc, 0xa4, 0xda, 0x50, 0x23, 0x15, 0xce, 0xb4, 0xfd, 0x14, 0x1a, 0x86, 0xb6, 0x17,
	0x84, 0x67, 0x4c, 0xfd, 0x97, 0x5a, 0xff, 0xdd, 0xde, 0x12, 0x73, 0xf7, 0xf2, 0x1e, 0x1c, 0xd6,
	0x49, 0xde, 0x91, 0xbb, 0x50, 0x9a, 0x05, 0x42, 0xda, 0xc5, 0x4e, 0xb1, 0x5b, 0xeb, 0x6f, 0x2e,
	0x3d, 0x6e, 0x3a, 0x1e, 0xaa, 0x54, 0xe7, 0xa7, 0x02, 0xbc, 0xa1, 0xef, 0xc0, 0x5c, 0xc9, 0x13,
	0xa4, 0x40, 0xef, 0x43, 0x73, 0xee, 0x38, 0x2f, 0x09, 0x6a, 0xd7, 0x35, 0xe6, 0xe8, 0x49, 0x40,
	0xd0, 0x3b, 0x50, 0xd3, 0x2e, 0xf3, 0x94, 0xd3, 0xb5, 0xf1, 0x40, 0x43, 0x5f, 0x26, 0x7e, 0x7f,
	0x0f, 0x1a, 0xe6, 0x5a, 0x75, 0x4a, 0x59, 0xa5, 0xd4, 0x0d, 0xa8, 0x92, 0xba, 0xd0, 0xc2, 0xb1,
	0x9c, 0x78, 0x82, 0xfa, 0xdc, 0x94, 0xd2, 0x5e, 0x6d, 0x26, 0xf8, 0x48, 0xc1, 0x2a, 0x73, 0x07,
	0x90, 0xb1, 0xa0, 0xe7, 0x73, 0x16, 0x7a, 0xf4, 0x2a, 0xe2, 0xa9, 0x7d, 0x5b, 0x26, 0x72, 0xc8,
	0x59, 0x78, 0x7c, 0x15, 0x71, 0xe7, 0x18, 0x5e, 0x69, 0x25, 0x4e, 0x42, 0xf1, 0x74, 0x2d, 0x9c,
	0x3b, 0x0b, 0x36, 0x32, 0x91, 0xa9, 0x90, 0x8c, 0x3f, 0x49, 0xd2, 0xbc, 0x14, 0xea, 0x0d, 0xa8,
	0x2c, 0x4a, 0x31, 0x48, 0xde, 0x82, 0x65, 0x52, 0x34, 0x96, 0x49, 0xe1, 0x1c, 0xcd, 0xbd, 0x7a,
	0x38, 0xa1, 0xfe, 0xf4, 0xf1, 0x96, 0x96, 0xbc, 0x3d, 0xfd, 0x9f, 0xcb, 0x50, 0x35, 0x65, 0x04,
	0xfa, 0xd5, 0x82, 0x52, 0x62, 0x7c, 0xd4, 0x5d, 0xe9, 0xb4, 0xdc, 0x7c, 0xb5, 0xb7, 0xfe, 0x47,
	0xa6, 0x9e, 0x22, 0xe7, 0xf4, 0xf6, 0x17, 0xbb, 0x50, 0xb1, 0x6e, 0xff, 0xfa, 0xfb, 0xae, 0xf0,
	0x15, 0xfa, 0xc2, 0x5d, 0x78, 0xb5, 0xa7, 0xf1, 0x98, 0xf2, 0x90, 0x4a, 0x2a, 0xdc, 0xb4, 0x88,
	0x9b, 0xf6, 0x2d, 0xdc, 0xeb, 0x74, 0x75, 0xe3, 0x9a, 0xe1, 0x10, 0xee, 0x75, 0x1c, 0x90, 0x1b,
	0x57, 0x64, 0xad, 0xff, 0x61, 0x41, 0x73, 0xd1, 0xf6, 0x68, 0x7b, 0x69, 0x6b, 0x4b, 0x67, 0xa3,
	0xbd, 0xf9, 0x60, 0xa2, 0xbf, 0x66, 0x01, 0xc9, 0x5a, 0x9f, 0xe5, 0x5a, 0xff, 0xde, 0x39, 0x7d,
	0x8e, 0xd6, 0xb1, 0xaf, 0x7e, 0xa2, 0x5c, 0xe3, 0xcc, 0x8f, 0xf4, 0x8b, 0xba, 0x67, 0x6d, 0xa3,
	0x3f, 0x2d, 0x68, 0xdd, 0x37, 0x2e, 0xda, 0x59, 0xc1, 0xe6, 0x81, 0xbf, 0x1f, 0xe3, 0x33, 0xcd,
	0xf1, 0xf1, 0xda, 0xdf, 0x3d, 0x27, 0x9f, 0x38, 0xbc, 0xc7, 0x08, 0xfd, 0x66, 0xc1, 0x5a, 0x3a,
	0x37, 0xe8, 0xc3, 0xd5, 0x4f, 0xd8, 0xc2, 0x74, 0x3d, 0x46, 0x82, 0xe4, 0x48, 0x7c, 0xe3, 0x8c,
	0x9e, 0x93, 0x04, 0xd7, 0x8d, 0xec, 0x59, 0xdb, 0x07, 0x9f, 0xc0, 0x9b, 0x3e, 0xbb, 0x98, 0x77,
	0x82, 0xa3, 0x20, 0xd7, 0xfa, 0x41, 0xc3, 0xf4, 0x3e, 0x48, 0x3e, 0x29, 0x06, 0xd6, 0xb7, 0x6b,
	0x69, 0x64, 0x5c, 0x56, 0x1f, 0x19, 0x1f, 0xff, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xd9, 0xf6, 0x91,
	0x38, 0xe1, 0x08, 0x00, 0x00,
}
