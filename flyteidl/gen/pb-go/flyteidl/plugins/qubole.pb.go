// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/plugins/qubole.proto

package plugins

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// Defines a query to execute on a hive cluster.
type HiveQuery struct {
	Query                string   `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	TimeoutSec           uint32   `protobuf:"varint,2,opt,name=timeout_sec,json=timeoutSec,proto3" json:"timeout_sec,omitempty"`
	RetryCount           uint32   `protobuf:"varint,3,opt,name=retryCount,proto3" json:"retryCount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HiveQuery) Reset()         { *m = HiveQuery{} }
func (m *HiveQuery) String() string { return proto.CompactTextString(m) }
func (*HiveQuery) ProtoMessage()    {}
func (*HiveQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cb86d766c12ee2e, []int{0}
}

func (m *HiveQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HiveQuery.Unmarshal(m, b)
}
func (m *HiveQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HiveQuery.Marshal(b, m, deterministic)
}
func (m *HiveQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HiveQuery.Merge(m, src)
}
func (m *HiveQuery) XXX_Size() int {
	return xxx_messageInfo_HiveQuery.Size(m)
}
func (m *HiveQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_HiveQuery.DiscardUnknown(m)
}

var xxx_messageInfo_HiveQuery proto.InternalMessageInfo

func (m *HiveQuery) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *HiveQuery) GetTimeoutSec() uint32 {
	if m != nil {
		return m.TimeoutSec
	}
	return 0
}

func (m *HiveQuery) GetRetryCount() uint32 {
	if m != nil {
		return m.RetryCount
	}
	return 0
}

// Defines a collection of hive queries.
type HiveQueryCollection struct {
	Queries              []*HiveQuery `protobuf:"bytes,2,rep,name=queries,proto3" json:"queries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *HiveQueryCollection) Reset()         { *m = HiveQueryCollection{} }
func (m *HiveQueryCollection) String() string { return proto.CompactTextString(m) }
func (*HiveQueryCollection) ProtoMessage()    {}
func (*HiveQueryCollection) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cb86d766c12ee2e, []int{1}
}

func (m *HiveQueryCollection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HiveQueryCollection.Unmarshal(m, b)
}
func (m *HiveQueryCollection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HiveQueryCollection.Marshal(b, m, deterministic)
}
func (m *HiveQueryCollection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HiveQueryCollection.Merge(m, src)
}
func (m *HiveQueryCollection) XXX_Size() int {
	return xxx_messageInfo_HiveQueryCollection.Size(m)
}
func (m *HiveQueryCollection) XXX_DiscardUnknown() {
	xxx_messageInfo_HiveQueryCollection.DiscardUnknown(m)
}

var xxx_messageInfo_HiveQueryCollection proto.InternalMessageInfo

func (m *HiveQueryCollection) GetQueries() []*HiveQuery {
	if m != nil {
		return m.Queries
	}
	return nil
}

// This message works with the 'hive' task type in the SDK and is the object that will be in the 'custom' field
// of a hive task's TaskTemplate
type QuboleHiveJob struct {
	ClusterLabel         string               `protobuf:"bytes,1,opt,name=cluster_label,json=clusterLabel,proto3" json:"cluster_label,omitempty"`
	QueryCollection      *HiveQueryCollection `protobuf:"bytes,2,opt,name=query_collection,json=queryCollection,proto3" json:"query_collection,omitempty"` // Deprecated: Do not use.
	Tags                 []string             `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
	Query                *HiveQuery           `protobuf:"bytes,4,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *QuboleHiveJob) Reset()         { *m = QuboleHiveJob{} }
func (m *QuboleHiveJob) String() string { return proto.CompactTextString(m) }
func (*QuboleHiveJob) ProtoMessage()    {}
func (*QuboleHiveJob) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cb86d766c12ee2e, []int{2}
}

func (m *QuboleHiveJob) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuboleHiveJob.Unmarshal(m, b)
}
func (m *QuboleHiveJob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuboleHiveJob.Marshal(b, m, deterministic)
}
func (m *QuboleHiveJob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuboleHiveJob.Merge(m, src)
}
func (m *QuboleHiveJob) XXX_Size() int {
	return xxx_messageInfo_QuboleHiveJob.Size(m)
}
func (m *QuboleHiveJob) XXX_DiscardUnknown() {
	xxx_messageInfo_QuboleHiveJob.DiscardUnknown(m)
}

var xxx_messageInfo_QuboleHiveJob proto.InternalMessageInfo

func (m *QuboleHiveJob) GetClusterLabel() string {
	if m != nil {
		return m.ClusterLabel
	}
	return ""
}

// Deprecated: Do not use.
func (m *QuboleHiveJob) GetQueryCollection() *HiveQueryCollection {
	if m != nil {
		return m.QueryCollection
	}
	return nil
}

func (m *QuboleHiveJob) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *QuboleHiveJob) GetQuery() *HiveQuery {
	if m != nil {
		return m.Query
	}
	return nil
}

func init() {
	proto.RegisterType((*HiveQuery)(nil), "flyteidl.plugins.HiveQuery")
	proto.RegisterType((*HiveQueryCollection)(nil), "flyteidl.plugins.HiveQueryCollection")
	proto.RegisterType((*QuboleHiveJob)(nil), "flyteidl.plugins.QuboleHiveJob")
}

func init() { proto.RegisterFile("flyteidl/plugins/qubole.proto", fileDescriptor_7cb86d766c12ee2e) }

var fileDescriptor_7cb86d766c12ee2e = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0x41, 0x4b, 0xc3, 0x30,
	0x18, 0xa5, 0xeb, 0x54, 0xf6, 0xcd, 0xe1, 0x88, 0x1e, 0x0a, 0xa2, 0x96, 0x89, 0xd0, 0x8b, 0x0d,
	0x4e, 0x3c, 0x8a, 0xb0, 0x5d, 0x44, 0x76, 0x59, 0xf5, 0xe4, 0x65, 0x2c, 0xf1, 0x33, 0x06, 0xb2,
	0x66, 0x4b, 0x13, 0x61, 0x3f, 0xd3, 0x7f, 0x24, 0x4d, 0xbb, 0x0e, 0x7a, 0xd8, 0xed, 0xcb, 0x7b,
	0x5f, 0xf2, 0xde, 0xcb, 0x83, 0xab, 0x6f, 0xb5, 0xb5, 0x28, 0xbf, 0x14, 0x5d, 0x2b, 0x27, 0x64,
	0x5e, 0xd0, 0x8d, 0x63, 0x5a, 0x61, 0xba, 0x36, 0xda, 0x6a, 0x32, 0xdc, 0xd1, 0x69, 0x4d, 0x8f,
	0x18, 0xf4, 0x5e, 0xe5, 0x2f, 0xce, 0x1d, 0x9a, 0x2d, 0xb9, 0x80, 0xa3, 0x4d, 0x39, 0x44, 0x41,
	0x1c, 0x24, 0xbd, 0xac, 0x3a, 0x90, 0x1b, 0xe8, 0x5b, 0xb9, 0x42, 0xed, 0xec, 0xa2, 0x40, 0x1e,
	0x75, 0xe2, 0x20, 0x19, 0x64, 0x50, 0x43, 0xef, 0xc8, 0xc9, 0x35, 0x80, 0x41, 0x6b, 0xb6, 0x53,
	0xed, 0x72, 0x1b, 0x85, 0x15, 0xbf, 0x47, 0x46, 0x33, 0x38, 0x6f, 0x34, 0xa6, 0x5a, 0x29, 0xe4,
	0x56, 0xea, 0x9c, 0x3c, 0xc1, 0x49, 0x29, 0x20, 0xb1, 0x88, 0x3a, 0x71, 0x98, 0xf4, 0xc7, 0x97,
	0x69, 0xdb, 0x5e, 0xda, 0xdc, 0xcb, 0x76, 0xbb, 0xa3, 0xbf, 0x00, 0x06, 0x73, 0x1f, 0xaa, 0x24,
	0xdf, 0x34, 0x23, 0xb7, 0x30, 0xe0, 0xca, 0x15, 0x16, 0xcd, 0x42, 0x2d, 0x19, 0xaa, 0xda, 0xfe,
	0x69, 0x0d, 0xce, 0x4a, 0x8c, 0x7c, 0xc0, 0xd0, 0xc7, 0x59, 0xf0, 0xc6, 0x81, 0x8f, 0xd2, 0x1f,
	0xdf, 0x1d, 0x90, 0xdd, 0xdb, 0x9d, 0x74, 0xa2, 0x20, 0x3b, 0xdb, 0xb4, 0x32, 0x10, 0xe8, 0xda,
	0xa5, 0x28, 0xa2, 0x30, 0x0e, 0x93, 0x5e, 0xe6, 0x67, 0xf2, 0xb0, 0xfb, 0xc5, 0xae, 0x7f, 0xfe,
	0x60, 0xaa, 0x6a, 0x73, 0xf2, 0xf2, 0xf9, 0x2c, 0xa4, 0xfd, 0x71, 0x2c, 0xe5, 0x7a, 0x45, 0xfd,
	0xbe, 0x36, 0xa2, 0x1a, 0x68, 0x53, 0xa9, 0xc0, 0x9c, 0xae, 0xd9, 0xbd, 0xd0, 0xb4, 0xdd, 0x32,
	0x3b, 0xf6, 0xfd, 0x3e, 0xfe, 0x07, 0x00, 0x00, 0xff, 0xff, 0xee, 0x0b, 0x00, 0x6a, 0x00, 0x02,
	0x00, 0x00,
}
