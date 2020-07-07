// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/vpc/v1/security_group.proto

package vpc

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
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

type SecurityGroup_Status int32

const (
	SecurityGroup_STATUS_UNSPECIFIED SecurityGroup_Status = 0
	SecurityGroup_CREATING           SecurityGroup_Status = 1
	SecurityGroup_ACTIVE             SecurityGroup_Status = 2
	// updating is a long operation because we must update all instances in SG
	SecurityGroup_UPDATING SecurityGroup_Status = 3
	SecurityGroup_DELETING SecurityGroup_Status = 4
)

var SecurityGroup_Status_name = map[int32]string{
	0: "STATUS_UNSPECIFIED",
	1: "CREATING",
	2: "ACTIVE",
	3: "UPDATING",
	4: "DELETING",
}

var SecurityGroup_Status_value = map[string]int32{
	"STATUS_UNSPECIFIED": 0,
	"CREATING":           1,
	"ACTIVE":             2,
	"UPDATING":           3,
	"DELETING":           4,
}

func (x SecurityGroup_Status) String() string {
	return proto.EnumName(SecurityGroup_Status_name, int32(x))
}

func (SecurityGroup_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7d1de6b05354e7cc, []int{0, 0}
}

type SecurityGroupRule_Direction int32

const (
	SecurityGroupRule_DIRECTION_UNSPECIFIED SecurityGroupRule_Direction = 0
	SecurityGroupRule_INGRESS               SecurityGroupRule_Direction = 1
	SecurityGroupRule_EGRESS                SecurityGroupRule_Direction = 2
)

var SecurityGroupRule_Direction_name = map[int32]string{
	0: "DIRECTION_UNSPECIFIED",
	1: "INGRESS",
	2: "EGRESS",
}

var SecurityGroupRule_Direction_value = map[string]int32{
	"DIRECTION_UNSPECIFIED": 0,
	"INGRESS":               1,
	"EGRESS":                2,
}

func (x SecurityGroupRule_Direction) String() string {
	return proto.EnumName(SecurityGroupRule_Direction_name, int32(x))
}

func (SecurityGroupRule_Direction) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7d1de6b05354e7cc, []int{1, 0}
}

type SecurityGroup struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FolderId             string               `protobuf:"bytes,2,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Name                 string               `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Description          string               `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Labels               map[string]string    `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	NetworkId            string               `protobuf:"bytes,7,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
	Status               SecurityGroup_Status `protobuf:"varint,8,opt,name=status,proto3,enum=yandex.cloud.vpc.v1.SecurityGroup_Status" json:"status,omitempty"`
	Rules                []*SecurityGroupRule `protobuf:"bytes,9,rep,name=rules,proto3" json:"rules,omitempty"`
	DefaultForNetwork    bool                 `protobuf:"varint,10,opt,name=default_for_network,json=defaultForNetwork,proto3" json:"default_for_network,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SecurityGroup) Reset()         { *m = SecurityGroup{} }
func (m *SecurityGroup) String() string { return proto.CompactTextString(m) }
func (*SecurityGroup) ProtoMessage()    {}
func (*SecurityGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d1de6b05354e7cc, []int{0}
}

func (m *SecurityGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SecurityGroup.Unmarshal(m, b)
}
func (m *SecurityGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SecurityGroup.Marshal(b, m, deterministic)
}
func (m *SecurityGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecurityGroup.Merge(m, src)
}
func (m *SecurityGroup) XXX_Size() int {
	return xxx_messageInfo_SecurityGroup.Size(m)
}
func (m *SecurityGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_SecurityGroup.DiscardUnknown(m)
}

var xxx_messageInfo_SecurityGroup proto.InternalMessageInfo

func (m *SecurityGroup) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SecurityGroup) GetFolderId() string {
	if m != nil {
		return m.FolderId
	}
	return ""
}

func (m *SecurityGroup) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *SecurityGroup) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SecurityGroup) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *SecurityGroup) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *SecurityGroup) GetNetworkId() string {
	if m != nil {
		return m.NetworkId
	}
	return ""
}

func (m *SecurityGroup) GetStatus() SecurityGroup_Status {
	if m != nil {
		return m.Status
	}
	return SecurityGroup_STATUS_UNSPECIFIED
}

func (m *SecurityGroup) GetRules() []*SecurityGroupRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

func (m *SecurityGroup) GetDefaultForNetwork() bool {
	if m != nil {
		return m.DefaultForNetwork
	}
	return false
}

type SecurityGroupRule struct {
	Id          string                      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description string                      `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Labels      map[string]string           `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Direction   SecurityGroupRule_Direction `protobuf:"varint,4,opt,name=direction,proto3,enum=yandex.cloud.vpc.v1.SecurityGroupRule_Direction" json:"direction,omitempty"`
	Ports       *PortRange                  `protobuf:"bytes,5,opt,name=ports,proto3" json:"ports,omitempty"`
	// null value means any protocol
	// values from https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml
	ProtocolName   string `protobuf:"bytes,6,opt,name=protocol_name,json=protocolName,proto3" json:"protocol_name,omitempty"`
	ProtocolNumber int64  `protobuf:"varint,7,opt,name=protocol_number,json=protocolNumber,proto3" json:"protocol_number,omitempty"`
	// Types that are valid to be assigned to Target:
	//	*SecurityGroupRule_CidrBlocks
	//	*SecurityGroupRule_SecurityGroupId
	//	*SecurityGroupRule_PredefinedTarget
	Target               isSecurityGroupRule_Target `protobuf_oneof:"target"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *SecurityGroupRule) Reset()         { *m = SecurityGroupRule{} }
func (m *SecurityGroupRule) String() string { return proto.CompactTextString(m) }
func (*SecurityGroupRule) ProtoMessage()    {}
func (*SecurityGroupRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d1de6b05354e7cc, []int{1}
}

func (m *SecurityGroupRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SecurityGroupRule.Unmarshal(m, b)
}
func (m *SecurityGroupRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SecurityGroupRule.Marshal(b, m, deterministic)
}
func (m *SecurityGroupRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecurityGroupRule.Merge(m, src)
}
func (m *SecurityGroupRule) XXX_Size() int {
	return xxx_messageInfo_SecurityGroupRule.Size(m)
}
func (m *SecurityGroupRule) XXX_DiscardUnknown() {
	xxx_messageInfo_SecurityGroupRule.DiscardUnknown(m)
}

var xxx_messageInfo_SecurityGroupRule proto.InternalMessageInfo

func (m *SecurityGroupRule) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SecurityGroupRule) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *SecurityGroupRule) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *SecurityGroupRule) GetDirection() SecurityGroupRule_Direction {
	if m != nil {
		return m.Direction
	}
	return SecurityGroupRule_DIRECTION_UNSPECIFIED
}

func (m *SecurityGroupRule) GetPorts() *PortRange {
	if m != nil {
		return m.Ports
	}
	return nil
}

func (m *SecurityGroupRule) GetProtocolName() string {
	if m != nil {
		return m.ProtocolName
	}
	return ""
}

func (m *SecurityGroupRule) GetProtocolNumber() int64 {
	if m != nil {
		return m.ProtocolNumber
	}
	return 0
}

type isSecurityGroupRule_Target interface {
	isSecurityGroupRule_Target()
}

type SecurityGroupRule_CidrBlocks struct {
	CidrBlocks *CidrBlocks `protobuf:"bytes,8,opt,name=cidr_blocks,json=cidrBlocks,proto3,oneof"`
}

type SecurityGroupRule_SecurityGroupId struct {
	SecurityGroupId string `protobuf:"bytes,9,opt,name=security_group_id,json=securityGroupId,proto3,oneof"`
}

type SecurityGroupRule_PredefinedTarget struct {
	PredefinedTarget string `protobuf:"bytes,10,opt,name=predefined_target,json=predefinedTarget,proto3,oneof"`
}

func (*SecurityGroupRule_CidrBlocks) isSecurityGroupRule_Target() {}

func (*SecurityGroupRule_SecurityGroupId) isSecurityGroupRule_Target() {}

func (*SecurityGroupRule_PredefinedTarget) isSecurityGroupRule_Target() {}

func (m *SecurityGroupRule) GetTarget() isSecurityGroupRule_Target {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *SecurityGroupRule) GetCidrBlocks() *CidrBlocks {
	if x, ok := m.GetTarget().(*SecurityGroupRule_CidrBlocks); ok {
		return x.CidrBlocks
	}
	return nil
}

func (m *SecurityGroupRule) GetSecurityGroupId() string {
	if x, ok := m.GetTarget().(*SecurityGroupRule_SecurityGroupId); ok {
		return x.SecurityGroupId
	}
	return ""
}

func (m *SecurityGroupRule) GetPredefinedTarget() string {
	if x, ok := m.GetTarget().(*SecurityGroupRule_PredefinedTarget); ok {
		return x.PredefinedTarget
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SecurityGroupRule) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SecurityGroupRule_CidrBlocks)(nil),
		(*SecurityGroupRule_SecurityGroupId)(nil),
		(*SecurityGroupRule_PredefinedTarget)(nil),
	}
}

type PortRange struct {
	FromPort             int64    `protobuf:"varint,1,opt,name=from_port,json=fromPort,proto3" json:"from_port,omitempty"`
	ToPort               int64    `protobuf:"varint,2,opt,name=to_port,json=toPort,proto3" json:"to_port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PortRange) Reset()         { *m = PortRange{} }
func (m *PortRange) String() string { return proto.CompactTextString(m) }
func (*PortRange) ProtoMessage()    {}
func (*PortRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d1de6b05354e7cc, []int{2}
}

func (m *PortRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PortRange.Unmarshal(m, b)
}
func (m *PortRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PortRange.Marshal(b, m, deterministic)
}
func (m *PortRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PortRange.Merge(m, src)
}
func (m *PortRange) XXX_Size() int {
	return xxx_messageInfo_PortRange.Size(m)
}
func (m *PortRange) XXX_DiscardUnknown() {
	xxx_messageInfo_PortRange.DiscardUnknown(m)
}

var xxx_messageInfo_PortRange proto.InternalMessageInfo

func (m *PortRange) GetFromPort() int64 {
	if m != nil {
		return m.FromPort
	}
	return 0
}

func (m *PortRange) GetToPort() int64 {
	if m != nil {
		return m.ToPort
	}
	return 0
}

type CidrBlocks struct {
	V4CidrBlocks         []string `protobuf:"bytes,1,rep,name=v4_cidr_blocks,json=v4CidrBlocks,proto3" json:"v4_cidr_blocks,omitempty"`
	V6CidrBlocks         []string `protobuf:"bytes,2,rep,name=v6_cidr_blocks,json=v6CidrBlocks,proto3" json:"v6_cidr_blocks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CidrBlocks) Reset()         { *m = CidrBlocks{} }
func (m *CidrBlocks) String() string { return proto.CompactTextString(m) }
func (*CidrBlocks) ProtoMessage()    {}
func (*CidrBlocks) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d1de6b05354e7cc, []int{3}
}

func (m *CidrBlocks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CidrBlocks.Unmarshal(m, b)
}
func (m *CidrBlocks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CidrBlocks.Marshal(b, m, deterministic)
}
func (m *CidrBlocks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CidrBlocks.Merge(m, src)
}
func (m *CidrBlocks) XXX_Size() int {
	return xxx_messageInfo_CidrBlocks.Size(m)
}
func (m *CidrBlocks) XXX_DiscardUnknown() {
	xxx_messageInfo_CidrBlocks.DiscardUnknown(m)
}

var xxx_messageInfo_CidrBlocks proto.InternalMessageInfo

func (m *CidrBlocks) GetV4CidrBlocks() []string {
	if m != nil {
		return m.V4CidrBlocks
	}
	return nil
}

func (m *CidrBlocks) GetV6CidrBlocks() []string {
	if m != nil {
		return m.V6CidrBlocks
	}
	return nil
}

func init() {
	proto.RegisterEnum("yandex.cloud.vpc.v1.SecurityGroup_Status", SecurityGroup_Status_name, SecurityGroup_Status_value)
	proto.RegisterEnum("yandex.cloud.vpc.v1.SecurityGroupRule_Direction", SecurityGroupRule_Direction_name, SecurityGroupRule_Direction_value)
	proto.RegisterType((*SecurityGroup)(nil), "yandex.cloud.vpc.v1.SecurityGroup")
	proto.RegisterMapType((map[string]string)(nil), "yandex.cloud.vpc.v1.SecurityGroup.LabelsEntry")
	proto.RegisterType((*SecurityGroupRule)(nil), "yandex.cloud.vpc.v1.SecurityGroupRule")
	proto.RegisterMapType((map[string]string)(nil), "yandex.cloud.vpc.v1.SecurityGroupRule.LabelsEntry")
	proto.RegisterType((*PortRange)(nil), "yandex.cloud.vpc.v1.PortRange")
	proto.RegisterType((*CidrBlocks)(nil), "yandex.cloud.vpc.v1.CidrBlocks")
}

func init() {
	proto.RegisterFile("yandex/cloud/vpc/v1/security_group.proto", fileDescriptor_7d1de6b05354e7cc)
}

var fileDescriptor_7d1de6b05354e7cc = []byte{
	// 828 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0xae, 0xf3, 0xe3, 0xc6, 0xc7, 0xbb, 0xd9, 0x74, 0x96, 0x1f, 0x53, 0xb4, 0x6c, 0x14, 0x56,
	0x10, 0x24, 0xea, 0x6c, 0xb2, 0x6d, 0xc5, 0xb2, 0x48, 0x28, 0x3f, 0x6e, 0xd7, 0x68, 0x15, 0xaa,
	0x49, 0x5a, 0x21, 0xb8, 0xb0, 0x1c, 0xcf, 0x24, 0x58, 0x75, 0x32, 0xd6, 0x78, 0x1c, 0xe8, 0x9b,
	0xf0, 0x38, 0x70, 0xd5, 0x57, 0xe0, 0x11, 0x78, 0x06, 0xae, 0x90, 0x67, 0xdc, 0xa4, 0x69, 0x23,
	0x51, 0x69, 0xef, 0x66, 0xce, 0xf7, 0x9d, 0x33, 0xdf, 0x9c, 0xef, 0x8c, 0x0d, 0xcd, 0x2b, 0x7f,
	0x41, 0xe8, 0xef, 0xad, 0x20, 0x62, 0x29, 0x69, 0x2d, 0xe3, 0xa0, 0xb5, 0x6c, 0xb7, 0x12, 0x1a,
	0xa4, 0x3c, 0x14, 0x57, 0xde, 0x8c, 0xb3, 0x34, 0xb6, 0x63, 0xce, 0x04, 0x43, 0x4f, 0x15, 0xd3,
	0x96, 0x4c, 0x7b, 0x19, 0x07, 0xf6, 0xb2, 0xbd, 0xff, 0x7c, 0xc6, 0xd8, 0x2c, 0xa2, 0x2d, 0x49,
	0x99, 0xa4, 0xd3, 0x96, 0x08, 0xe7, 0x34, 0x11, 0xfe, 0x3c, 0xcf, 0xda, 0x7f, 0xb6, 0x59, 0xdf,
	0x8f, 0x42, 0xe2, 0x8b, 0x90, 0x2d, 0x14, 0xdc, 0xf8, 0xbb, 0x04, 0x8f, 0x47, 0xf9, 0x69, 0xa7,
	0xd9, 0x61, 0xa8, 0x0a, 0x85, 0x90, 0x58, 0x5a, 0x5d, 0x6b, 0x1a, 0xb8, 0x10, 0x12, 0xf4, 0x29,
	0x18, 0x53, 0x16, 0x11, 0xca, 0xbd, 0x90, 0x58, 0x05, 0x19, 0xae, 0xa8, 0x80, 0x4b, 0xd0, 0x6b,
	0x80, 0x80, 0x53, 0x5f, 0x50, 0xe2, 0xf9, 0xc2, 0x2a, 0xd6, 0xb5, 0xa6, 0xd9, 0xd9, 0xb7, 0x95,
	0x26, 0xfb, 0x46, 0x93, 0x3d, 0xbe, 0xd1, 0x84, 0x8d, 0x9c, 0xdd, 0x15, 0x08, 0x41, 0x69, 0xe1,
	0xcf, 0xa9, 0x55, 0x92, 0x25, 0xe5, 0x1a, 0xd5, 0xc1, 0x24, 0x34, 0x09, 0x78, 0x18, 0x67, 0x12,
	0xad, 0xb2, 0x84, 0x6e, 0x87, 0xd0, 0x09, 0xe8, 0x91, 0x3f, 0xa1, 0x51, 0x62, 0xe9, 0xf5, 0x62,
	0xd3, 0xec, 0xd8, 0xf6, 0x96, 0xae, 0xd8, 0x1b, 0x37, 0xb2, 0xdf, 0xc9, 0x04, 0x67, 0x21, 0xf8,
	0x15, 0xce, 0xb3, 0xd1, 0x33, 0x80, 0x05, 0x15, 0xbf, 0x31, 0x7e, 0x99, 0x5d, 0x6b, 0x57, 0x1e,
	0x64, 0xe4, 0x11, 0x97, 0xa0, 0x2e, 0xe8, 0x89, 0xf0, 0x45, 0x9a, 0x58, 0x95, 0xba, 0xd6, 0xac,
	0x76, 0xbe, 0x7a, 0xc0, 0x31, 0x23, 0x99, 0x80, 0xf3, 0x44, 0xf4, 0x1d, 0x94, 0x79, 0x1a, 0xd1,
	0xc4, 0x32, 0xa4, 0xd0, 0x2f, 0xfe, 0xbf, 0x02, 0x4e, 0x23, 0x8a, 0x55, 0x12, 0xb2, 0xe1, 0x29,
	0xa1, 0x53, 0x3f, 0x8d, 0x84, 0x37, 0x65, 0xdc, 0xcb, 0x95, 0x59, 0x50, 0xd7, 0x9a, 0x15, 0xbc,
	0x97, 0x43, 0x27, 0x8c, 0x0f, 0x15, 0xb0, 0xff, 0x1a, 0xcc, 0x5b, 0xd7, 0x44, 0x35, 0x28, 0x5e,
	0xd2, 0xab, 0xdc, 0xc5, 0x6c, 0x89, 0x3e, 0x80, 0xf2, 0xd2, 0x8f, 0x52, 0x9a, 0x5b, 0xa8, 0x36,
	0xdf, 0x16, 0xbe, 0xd1, 0x1a, 0x17, 0xa0, 0x2b, 0xe9, 0xe8, 0x23, 0x40, 0xa3, 0x71, 0x77, 0x7c,
	0x3e, 0xf2, 0xce, 0x87, 0xa3, 0x33, 0xa7, 0xef, 0x9e, 0xb8, 0xce, 0xa0, 0xb6, 0x83, 0x1e, 0x41,
	0xa5, 0x8f, 0x9d, 0xee, 0xd8, 0x1d, 0x9e, 0xd6, 0x34, 0x04, 0xa0, 0x77, 0xfb, 0x63, 0xf7, 0xc2,
	0xa9, 0x15, 0x32, 0xe4, 0xfc, 0x6c, 0xa0, 0x90, 0x62, 0xb6, 0x1b, 0x38, 0xef, 0x1c, 0xb9, 0x2b,
	0x35, 0xfe, 0x28, 0xc3, 0xde, 0xbd, 0xfb, 0xdd, 0x1b, 0xaf, 0x3b, 0x96, 0x17, 0xee, 0x5b, 0xfe,
	0xc3, 0xca, 0xf2, 0xa2, 0xec, 0x64, 0xe7, 0x61, 0x9d, 0xdc, 0x6a, 0xfb, 0x18, 0x0c, 0x12, 0x72,
	0x1a, 0xc8, 0xb3, 0x4a, 0xd2, 0xda, 0x97, 0x0f, 0x2c, 0x37, 0xb8, 0xc9, 0xeb, 0x95, 0xfe, 0xb9,
	0x6e, 0x6b, 0x78, 0x5d, 0x08, 0x1d, 0x42, 0x39, 0x66, 0x5c, 0x24, 0x72, 0x60, 0xcd, 0xce, 0x67,
	0x5b, 0x2b, 0x9e, 0x31, 0x2e, 0xb0, 0xbf, 0x98, 0x51, 0xac, 0xc8, 0xe8, 0x73, 0x78, 0x2c, 0x5f,
	0x48, 0xc0, 0x22, 0x4f, 0xbe, 0x04, 0x5d, 0xde, 0xfd, 0xd1, 0x4d, 0x70, 0x98, 0xbd, 0x88, 0x2f,
	0xe1, 0xc9, 0x9a, 0x94, 0xce, 0x27, 0x94, 0xcb, 0x61, 0x2d, 0xe2, 0xea, 0x8a, 0x26, 0xa3, 0xa8,
	0x07, 0x66, 0x10, 0x12, 0xee, 0x4d, 0x22, 0x16, 0x5c, 0xaa, 0xb1, 0x35, 0x3b, 0xcf, 0xb7, 0x2a,
	0xe9, 0x87, 0x84, 0xf7, 0x24, 0xed, 0xed, 0x0e, 0x86, 0x60, 0xb5, 0x43, 0x5f, 0xc3, 0xde, 0xe6,
	0x97, 0x27, 0x7b, 0x1b, 0x46, 0xa6, 0xea, 0xed, 0x0e, 0x7e, 0x92, 0xdc, 0x6e, 0x89, 0x4b, 0xd0,
	0x01, 0xec, 0xc5, 0x9c, 0x12, 0x3a, 0x0d, 0x17, 0x94, 0x78, 0xc2, 0xe7, 0x33, 0x2a, 0xe4, 0x80,
	0x66, 0xec, 0xda, 0x1a, 0x1a, 0x4b, 0xe4, 0x7d, 0x26, 0xf4, 0x7b, 0x30, 0x56, 0xdd, 0x47, 0x9f,
	0xc0, 0x87, 0x03, 0x17, 0x3b, 0xfd, 0xb1, 0xfb, 0xe3, 0xf0, 0xce, 0x9c, 0x9a, 0xb0, 0xeb, 0x0e,
	0x4f, 0xb1, 0x33, 0x1a, 0xa9, 0x31, 0x75, 0xd4, 0xba, 0xd0, 0xab, 0x82, 0xae, 0xf4, 0xa1, 0xd2,
	0x9f, 0x7f, 0xb5, 0xb5, 0xc6, 0x2f, 0x60, 0xac, 0xec, 0x40, 0x4d, 0x30, 0xa6, 0x9c, 0xcd, 0xbd,
	0xcc, 0x15, 0xa9, 0xa7, 0xd8, 0x33, 0xff, 0xbd, 0x6e, 0xef, 0xbe, 0x3c, 0x38, 0x3e, 0x3a, 0x7a,
	0x75, 0x84, 0x2b, 0x19, 0x9a, 0xd1, 0xd1, 0x0b, 0xd8, 0x15, 0x4c, 0xf1, 0x0a, 0xf7, 0x79, 0xba,
	0x60, 0x19, 0xab, 0xf1, 0x13, 0xc0, 0xba, 0xc3, 0xe8, 0x05, 0x54, 0x97, 0x87, 0xde, 0x6d, 0x6b,
	0xb4, 0x7a, 0x31, 0xb3, 0x79, 0x79, 0x78, 0x87, 0x75, 0xbc, 0xc1, 0x2a, 0xe4, 0xac, 0xe3, 0x35,
	0xab, 0x77, 0x01, 0x1f, 0x6f, 0xf8, 0xe9, 0xc7, 0x61, 0xee, 0xe9, 0xcf, 0x6f, 0x66, 0xa1, 0xf8,
	0x35, 0x9d, 0xd8, 0x01, 0x9b, 0xb7, 0x14, 0xe7, 0x40, 0x7d, 0xf1, 0x67, 0xec, 0x60, 0x46, 0x17,
	0x72, 0x58, 0x5a, 0x5b, 0x7e, 0x35, 0x6f, 0x96, 0x71, 0x30, 0xd1, 0x25, 0xfc, 0xea, 0xbf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x1c, 0x41, 0x3d, 0x06, 0x8c, 0x06, 0x00, 0x00,
}