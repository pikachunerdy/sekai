// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tx.proto

package types

import (
	bytes "bytes"
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgProposalSoftwareUpgradeRequest struct {
	Resources            *Resource                                     `protobuf:"bytes,1,opt,name=resources,proto3" json:"resources,omitempty"`
	MinHaltTime          int64                                         `protobuf:"varint,2,opt,name=min_halt_time,json=minHaltTime,proto3" json:"min_halt_time,omitempty"`
	OldChainId           string                                        `protobuf:"bytes,3,opt,name=old_chain_id,json=oldChainId,proto3" json:"old_chain_id,omitempty"`
	NewChainId           string                                        `protobuf:"bytes,4,opt,name=new_chain_id,json=newChainId,proto3" json:"new_chain_id,omitempty"`
	RollbackChecksum     string                                        `protobuf:"bytes,5,opt,name=rollback_checksum,json=rollbackChecksum,proto3" json:"rollback_checksum,omitempty"`
	MaxEnrolmentDuration int64                                         `protobuf:"varint,6,opt,name=max_enrolment_duration,json=maxEnrolmentDuration,proto3" json:"max_enrolment_duration,omitempty"`
	Memo                 string                                        `protobuf:"bytes,7,opt,name=memo,proto3" json:"memo,omitempty"`
	Proposer             github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,8,opt,name=proposer,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"proposer,omitempty" yaml:"proposer"`
}

func (m *MsgProposalSoftwareUpgradeRequest) Reset()         { *m = MsgProposalSoftwareUpgradeRequest{} }
func (m *MsgProposalSoftwareUpgradeRequest) String() string { return proto.CompactTextString(m) }
func (*MsgProposalSoftwareUpgradeRequest) ProtoMessage()    {}
func (*MsgProposalSoftwareUpgradeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fd2153dc07d3b5c, []int{0}
}
func (m *MsgProposalSoftwareUpgradeRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgProposalSoftwareUpgradeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgProposalSoftwareUpgradeRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgProposalSoftwareUpgradeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgProposalSoftwareUpgradeRequest.Merge(m, src)
}
func (m *MsgProposalSoftwareUpgradeRequest) XXX_Size() int {
	return m.Size()
}
func (m *MsgProposalSoftwareUpgradeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgProposalSoftwareUpgradeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgProposalSoftwareUpgradeRequest proto.InternalMessageInfo

func (m *MsgProposalSoftwareUpgradeRequest) GetResources() *Resource {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *MsgProposalSoftwareUpgradeRequest) GetMinHaltTime() int64 {
	if m != nil {
		return m.MinHaltTime
	}
	return 0
}

func (m *MsgProposalSoftwareUpgradeRequest) GetOldChainId() string {
	if m != nil {
		return m.OldChainId
	}
	return ""
}

func (m *MsgProposalSoftwareUpgradeRequest) GetNewChainId() string {
	if m != nil {
		return m.NewChainId
	}
	return ""
}

func (m *MsgProposalSoftwareUpgradeRequest) GetRollbackChecksum() string {
	if m != nil {
		return m.RollbackChecksum
	}
	return ""
}

func (m *MsgProposalSoftwareUpgradeRequest) GetMaxEnrolmentDuration() int64 {
	if m != nil {
		return m.MaxEnrolmentDuration
	}
	return 0
}

func (m *MsgProposalSoftwareUpgradeRequest) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func (m *MsgProposalSoftwareUpgradeRequest) GetProposer() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Proposer
	}
	return nil
}

type MsgProposalSoftwareUpgradeResponse struct {
	ProposalID uint64 `protobuf:"varint,1,opt,name=proposalID,proto3" json:"proposalID,omitempty"`
}

func (m *MsgProposalSoftwareUpgradeResponse) Reset()         { *m = MsgProposalSoftwareUpgradeResponse{} }
func (m *MsgProposalSoftwareUpgradeResponse) String() string { return proto.CompactTextString(m) }
func (*MsgProposalSoftwareUpgradeResponse) ProtoMessage()    {}
func (*MsgProposalSoftwareUpgradeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fd2153dc07d3b5c, []int{1}
}
func (m *MsgProposalSoftwareUpgradeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgProposalSoftwareUpgradeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgProposalSoftwareUpgradeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgProposalSoftwareUpgradeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgProposalSoftwareUpgradeResponse.Merge(m, src)
}
func (m *MsgProposalSoftwareUpgradeResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgProposalSoftwareUpgradeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgProposalSoftwareUpgradeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgProposalSoftwareUpgradeResponse proto.InternalMessageInfo

func (m *MsgProposalSoftwareUpgradeResponse) GetProposalID() uint64 {
	if m != nil {
		return m.ProposalID
	}
	return 0
}

func init() {
	proto.RegisterType((*MsgProposalSoftwareUpgradeRequest)(nil), "kira.upgrade.MsgProposalSoftwareUpgradeRequest")
	proto.RegisterType((*MsgProposalSoftwareUpgradeResponse)(nil), "kira.upgrade.MsgProposalSoftwareUpgradeResponse")
}

func init() { proto.RegisterFile("tx.proto", fileDescriptor_0fd2153dc07d3b5c) }

var fileDescriptor_0fd2153dc07d3b5c = []byte{
	// 472 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4d, 0x6f, 0xd3, 0x4c,
	0x10, 0x80, 0xb3, 0x6f, 0xf2, 0x96, 0x74, 0x9b, 0x0a, 0x58, 0x55, 0xc5, 0xca, 0xc1, 0x09, 0x3e,
	0x05, 0xa1, 0xda, 0xa8, 0xf4, 0xd4, 0x5b, 0x9b, 0x20, 0x51, 0xa1, 0x4a, 0xc8, 0xc0, 0x85, 0x8b,
	0xd9, 0xd8, 0x83, 0xb3, 0xb2, 0xd7, 0x63, 0x76, 0xd7, 0x4a, 0x2a, 0x71, 0xe3, 0x0f, 0xf0, 0x13,
	0xfa, 0x73, 0x38, 0xf6, 0xc8, 0xa9, 0x42, 0xc9, 0x85, 0x33, 0x47, 0x4e, 0x28, 0xfe, 0xa0, 0xe1,
	0xc0, 0xc7, 0xc9, 0xeb, 0x79, 0x9e, 0x9d, 0x1d, 0xcd, 0x0c, 0xed, 0x9a, 0x85, 0x9b, 0x2b, 0x34,
	0xc8, 0x7a, 0x89, 0x50, 0xdc, 0x2d, 0xf2, 0x58, 0xf1, 0x08, 0xfa, 0x7b, 0x31, 0xc6, 0x58, 0x02,
	0x6f, 0x7d, 0xaa, 0x9c, 0xfe, 0x6e, 0x8d, 0xab, 0x5f, 0xe7, 0xb2, 0x4d, 0xef, 0x9f, 0xeb, 0xf8,
	0xb9, 0xc2, 0x1c, 0x35, 0x4f, 0x5f, 0xe0, 0x5b, 0x33, 0xe7, 0x0a, 0x5e, 0x55, 0x92, 0x0f, 0xef,
	0x0a, 0xd0, 0x86, 0x1d, 0xd1, 0x6d, 0x05, 0x1a, 0x0b, 0x15, 0x82, 0xb6, 0xc8, 0x90, 0x8c, 0x76,
	0x0e, 0xf7, 0xdd, 0xcd, 0xc7, 0xdc, 0x06, 0xfb, 0x37, 0x22, 0x73, 0xe8, 0xae, 0x14, 0x59, 0x30,
	0xe3, 0xa9, 0x09, 0x8c, 0x90, 0x60, 0xfd, 0x37, 0x24, 0xa3, 0xb6, 0xbf, 0x23, 0x45, 0xf6, 0x94,
	0xa7, 0xe6, 0xa5, 0x90, 0xc0, 0x86, 0xb4, 0x87, 0x69, 0x14, 0x84, 0x33, 0x2e, 0xb2, 0x40, 0x44,
	0x56, 0x7b, 0x48, 0x46, 0xdb, 0x3e, 0xc5, 0x34, 0x1a, 0xaf, 0x43, 0x67, 0xd1, 0xda, 0xc8, 0x60,
	0x7e, 0x63, 0x74, 0x2a, 0x23, 0x83, 0x79, 0x63, 0x3c, 0xa4, 0x77, 0x15, 0xa6, 0xe9, 0x94, 0x87,
	0x49, 0x10, 0xce, 0x20, 0x4c, 0x74, 0x21, 0xad, 0xff, 0x4b, 0xed, 0x4e, 0x03, 0xc6, 0x75, 0x9c,
	0x1d, 0xd1, 0x7d, 0xc9, 0x17, 0x01, 0x64, 0x0a, 0x53, 0x09, 0x99, 0x09, 0xa2, 0x42, 0x71, 0x23,
	0x30, 0xb3, 0xb6, 0xca, 0xea, 0xf6, 0x24, 0x5f, 0x3c, 0x69, 0xe0, 0xa4, 0x66, 0x8c, 0xd1, 0x8e,
	0x04, 0x89, 0xd6, 0xad, 0x32, 0x6b, 0x79, 0x66, 0x6f, 0x68, 0x37, 0x2f, 0xdb, 0x06, 0xca, 0xea,
	0x0e, 0xc9, 0xa8, 0x77, 0x3a, 0xf9, 0x76, 0x3d, 0xb8, 0x7d, 0xc1, 0x65, 0x7a, 0xec, 0x34, 0xc4,
	0xf9, 0x7e, 0x3d, 0x38, 0x88, 0x85, 0x99, 0x15, 0x53, 0x37, 0x44, 0xe9, 0x85, 0xa8, 0x25, 0xea,
	0xfa, 0x73, 0xa0, 0xa3, 0xc4, 0x33, 0x17, 0x39, 0x68, 0xf7, 0x24, 0x0c, 0x4f, 0xa2, 0x48, 0x81,
	0xd6, 0xfe, 0xcf, 0xac, 0xc7, 0x9d, 0xaf, 0x97, 0x03, 0xe2, 0x4c, 0xa8, 0xf3, 0xa7, 0x09, 0xe9,
	0x1c, 0x33, 0x0d, 0xcc, 0xa6, 0x34, 0xaf, 0x95, 0xb3, 0x49, 0x39, 0xa3, 0x8e, 0xbf, 0x11, 0x39,
	0xfc, 0x40, 0x68, 0xfb, 0x5c, 0xc7, 0xec, 0x3d, 0xbd, 0xf7, 0x9b, 0x54, 0xcc, 0xfb, 0x75, 0xa4,
	0x7f, 0x5d, 0x8b, 0xfe, 0xa3, 0x7f, 0xbf, 0x50, 0x55, 0x79, 0x3a, 0xfe, 0xb4, 0xb4, 0xc9, 0xd5,
	0xd2, 0x26, 0x5f, 0x96, 0x36, 0xf9, 0xb8, 0xb2, 0x5b, 0x57, 0x2b, 0xbb, 0xf5, 0x79, 0x65, 0xb7,
	0x5e, 0x3f, 0xd8, 0x68, 0xd2, 0x33, 0xa1, 0xf8, 0x18, 0x15, 0x78, 0x1a, 0x12, 0x2e, 0xbc, 0x85,
	0x57, 0xbf, 0x50, 0xf5, 0x6a, 0xba, 0x55, 0xae, 0xee, 0xe3, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x3b, 0xfe, 0xc1, 0x5b, 0xf9, 0x02, 0x00, 0x00,
}

func (this *MsgProposalSoftwareUpgradeRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgProposalSoftwareUpgradeRequest)
	if !ok {
		that2, ok := that.(MsgProposalSoftwareUpgradeRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Resources.Equal(that1.Resources) {
		return false
	}
	if this.MinHaltTime != that1.MinHaltTime {
		return false
	}
	if this.OldChainId != that1.OldChainId {
		return false
	}
	if this.NewChainId != that1.NewChainId {
		return false
	}
	if this.RollbackChecksum != that1.RollbackChecksum {
		return false
	}
	if this.MaxEnrolmentDuration != that1.MaxEnrolmentDuration {
		return false
	}
	if this.Memo != that1.Memo {
		return false
	}
	if !bytes.Equal(this.Proposer, that1.Proposer) {
		return false
	}
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	ProposalSoftwareUpgrade(ctx context.Context, in *MsgProposalSoftwareUpgradeRequest, opts ...grpc.CallOption) (*MsgProposalSoftwareUpgradeResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) ProposalSoftwareUpgrade(ctx context.Context, in *MsgProposalSoftwareUpgradeRequest, opts ...grpc.CallOption) (*MsgProposalSoftwareUpgradeResponse, error) {
	out := new(MsgProposalSoftwareUpgradeResponse)
	err := c.cc.Invoke(ctx, "/kira.upgrade.Msg/ProposalSoftwareUpgrade", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	ProposalSoftwareUpgrade(context.Context, *MsgProposalSoftwareUpgradeRequest) (*MsgProposalSoftwareUpgradeResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) ProposalSoftwareUpgrade(ctx context.Context, req *MsgProposalSoftwareUpgradeRequest) (*MsgProposalSoftwareUpgradeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProposalSoftwareUpgrade not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_ProposalSoftwareUpgrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgProposalSoftwareUpgradeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ProposalSoftwareUpgrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kira.upgrade.Msg/ProposalSoftwareUpgrade",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ProposalSoftwareUpgrade(ctx, req.(*MsgProposalSoftwareUpgradeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kira.upgrade.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProposalSoftwareUpgrade",
			Handler:    _Msg_ProposalSoftwareUpgrade_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tx.proto",
}

func (m *MsgProposalSoftwareUpgradeRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgProposalSoftwareUpgradeRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgProposalSoftwareUpgradeRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Proposer) > 0 {
		i -= len(m.Proposer)
		copy(dAtA[i:], m.Proposer)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Proposer)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Memo) > 0 {
		i -= len(m.Memo)
		copy(dAtA[i:], m.Memo)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Memo)))
		i--
		dAtA[i] = 0x3a
	}
	if m.MaxEnrolmentDuration != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.MaxEnrolmentDuration))
		i--
		dAtA[i] = 0x30
	}
	if len(m.RollbackChecksum) > 0 {
		i -= len(m.RollbackChecksum)
		copy(dAtA[i:], m.RollbackChecksum)
		i = encodeVarintTx(dAtA, i, uint64(len(m.RollbackChecksum)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.NewChainId) > 0 {
		i -= len(m.NewChainId)
		copy(dAtA[i:], m.NewChainId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.NewChainId)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.OldChainId) > 0 {
		i -= len(m.OldChainId)
		copy(dAtA[i:], m.OldChainId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.OldChainId)))
		i--
		dAtA[i] = 0x1a
	}
	if m.MinHaltTime != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.MinHaltTime))
		i--
		dAtA[i] = 0x10
	}
	if m.Resources != nil {
		{
			size, err := m.Resources.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgProposalSoftwareUpgradeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgProposalSoftwareUpgradeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgProposalSoftwareUpgradeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ProposalID != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.ProposalID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgProposalSoftwareUpgradeRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Resources != nil {
		l = m.Resources.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	if m.MinHaltTime != 0 {
		n += 1 + sovTx(uint64(m.MinHaltTime))
	}
	l = len(m.OldChainId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.NewChainId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.RollbackChecksum)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.MaxEnrolmentDuration != 0 {
		n += 1 + sovTx(uint64(m.MaxEnrolmentDuration))
	}
	l = len(m.Memo)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Proposer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgProposalSoftwareUpgradeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProposalID != 0 {
		n += 1 + sovTx(uint64(m.ProposalID))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgProposalSoftwareUpgradeRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgProposalSoftwareUpgradeRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgProposalSoftwareUpgradeRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Resources", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Resources == nil {
				m.Resources = &Resource{}
			}
			if err := m.Resources.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinHaltTime", wireType)
			}
			m.MinHaltTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinHaltTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OldChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OldChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NewChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RollbackChecksum", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RollbackChecksum = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxEnrolmentDuration", wireType)
			}
			m.MaxEnrolmentDuration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxEnrolmentDuration |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Memo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Memo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proposer", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Proposer = append(m.Proposer[:0], dAtA[iNdEx:postIndex]...)
			if m.Proposer == nil {
				m.Proposer = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgProposalSoftwareUpgradeResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgProposalSoftwareUpgradeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgProposalSoftwareUpgradeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposalID", wireType)
			}
			m.ProposalID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProposalID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
