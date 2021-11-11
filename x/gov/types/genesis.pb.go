// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kira/gov/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type GenesisState struct {
	// starting_proposal_id is the ID of the starting proposal.
	StartingProposalId uint64  `protobuf:"varint,1,opt,name=starting_proposal_id,json=startingProposalId,proto3" json:"starting_proposal_id,omitempty"`
	NextRoleId         uint64  `protobuf:"varint,2,opt,name=next_role_id,json=nextRoleId,proto3" json:"next_role_id,omitempty"`
	Roles              []*Role `protobuf:"bytes,3,rep,name=roles,proto3" json:"roles,omitempty"`
	// role_permissions is the roles that are active from genesis.
	RolePermissions map[uint64]*Permissions `protobuf:"bytes,4,rep,name=role_permissions,json=rolePermissions,proto3" json:"role_permissions,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// NetworkActors are the actors that are saved from genesis.
	NetworkActors               []*NetworkActor               `protobuf:"bytes,5,rep,name=network_actors,json=networkActors,proto3" json:"network_actors,omitempty"`
	NetworkProperties           *NetworkProperties            `protobuf:"bytes,6,opt,name=network_properties,json=networkProperties,proto3" json:"network_properties,omitempty"`
	ExecutionFees               []*ExecutionFee               `protobuf:"bytes,7,rep,name=execution_fees,json=executionFees,proto3" json:"execution_fees,omitempty"`
	PoorNetworkMessages         *AllowedMessages              `protobuf:"bytes,8,opt,name=poor_network_messages,json=poorNetworkMessages,proto3" json:"poor_network_messages,omitempty"`
	Proposals                   []Proposal                    `protobuf:"bytes,9,rep,name=proposals,proto3" json:"proposals"`
	Votes                       []Vote                        `protobuf:"bytes,10,rep,name=votes,proto3" json:"votes"`
	DataRegistry                map[string]*DataRegistryEntry `protobuf:"bytes,11,rep,name=data_registry,json=dataRegistry,proto3" json:"data_registry,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	IdentityRecords             []IdentityRecord              `protobuf:"bytes,12,rep,name=identity_records,json=identityRecords,proto3" json:"identity_records"`
	LastIdentityRecordId        uint64                        `protobuf:"varint,13,opt,name=last_identity_record_id,json=lastIdentityRecordId,proto3" json:"last_identity_record_id,omitempty"`
	IdRecordsVerifyRequests     []IdentityRecordsVerify       `protobuf:"bytes,14,rep,name=id_records_verify_requests,json=idRecordsVerifyRequests,proto3" json:"id_records_verify_requests"`
	LastIdRecordVerifyRequestId uint64                        `protobuf:"varint,15,opt,name=last_id_record_verify_request_id,json=lastIdRecordVerifyRequestId,proto3" json:"last_id_record_verify_request_id,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_042721fb65a4ea2d, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetStartingProposalId() uint64 {
	if m != nil {
		return m.StartingProposalId
	}
	return 0
}

func (m *GenesisState) GetNextRoleId() uint64 {
	if m != nil {
		return m.NextRoleId
	}
	return 0
}

func (m *GenesisState) GetRoles() []*Role {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *GenesisState) GetRolePermissions() map[uint64]*Permissions {
	if m != nil {
		return m.RolePermissions
	}
	return nil
}

func (m *GenesisState) GetNetworkActors() []*NetworkActor {
	if m != nil {
		return m.NetworkActors
	}
	return nil
}

func (m *GenesisState) GetNetworkProperties() *NetworkProperties {
	if m != nil {
		return m.NetworkProperties
	}
	return nil
}

func (m *GenesisState) GetExecutionFees() []*ExecutionFee {
	if m != nil {
		return m.ExecutionFees
	}
	return nil
}

func (m *GenesisState) GetPoorNetworkMessages() *AllowedMessages {
	if m != nil {
		return m.PoorNetworkMessages
	}
	return nil
}

func (m *GenesisState) GetProposals() []Proposal {
	if m != nil {
		return m.Proposals
	}
	return nil
}

func (m *GenesisState) GetVotes() []Vote {
	if m != nil {
		return m.Votes
	}
	return nil
}

func (m *GenesisState) GetDataRegistry() map[string]*DataRegistryEntry {
	if m != nil {
		return m.DataRegistry
	}
	return nil
}

func (m *GenesisState) GetIdentityRecords() []IdentityRecord {
	if m != nil {
		return m.IdentityRecords
	}
	return nil
}

func (m *GenesisState) GetLastIdentityRecordId() uint64 {
	if m != nil {
		return m.LastIdentityRecordId
	}
	return 0
}

func (m *GenesisState) GetIdRecordsVerifyRequests() []IdentityRecordsVerify {
	if m != nil {
		return m.IdRecordsVerifyRequests
	}
	return nil
}

func (m *GenesisState) GetLastIdRecordVerifyRequestId() uint64 {
	if m != nil {
		return m.LastIdRecordVerifyRequestId
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "kira.gov.GenesisState")
	proto.RegisterMapType((map[string]*DataRegistryEntry)(nil), "kira.gov.GenesisState.DataRegistryEntry")
	proto.RegisterMapType((map[uint64]*Permissions)(nil), "kira.gov.GenesisState.RolePermissionsEntry")
}

func init() { proto.RegisterFile("kira/gov/genesis.proto", fileDescriptor_042721fb65a4ea2d) }

var fileDescriptor_042721fb65a4ea2d = []byte{
	// 705 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0x5f, 0x6b, 0x13, 0x4b,
	0x1c, 0x4d, 0xda, 0xa6, 0xb7, 0x99, 0x24, 0xfd, 0x33, 0x4d, 0xdb, 0xbd, 0xe9, 0x25, 0xcd, 0xbd,
	0x5c, 0x21, 0x58, 0x48, 0xb4, 0xa2, 0x88, 0x20, 0xd2, 0x6a, 0x95, 0x28, 0x95, 0xb2, 0x42, 0x41,
	0x11, 0x96, 0x69, 0xf6, 0xd7, 0x75, 0xc8, 0x76, 0x27, 0xce, 0x4c, 0xd2, 0xe6, 0x5b, 0x88, 0x9f,
	0xaa, 0x8f, 0x7d, 0xf4, 0x49, 0xa4, 0xfd, 0x22, 0x32, 0xb3, 0x33, 0xfb, 0x27, 0xb1, 0x6f, 0xbb,
	0xe7, 0x9c, 0xdf, 0xd9, 0x33, 0x73, 0x66, 0x16, 0x6d, 0x0e, 0x28, 0x27, 0xdd, 0x80, 0x8d, 0xbb,
	0x01, 0x44, 0x20, 0xa8, 0xe8, 0x0c, 0x39, 0x93, 0x0c, 0x2f, 0x29, 0xbc, 0x13, 0xb0, 0x71, 0xa3,
	0x1e, 0xb0, 0x80, 0x69, 0xb0, 0xab, 0x9e, 0x62, 0xbe, 0x51, 0x4f, 0xe6, 0x48, 0x5f, 0x32, 0x6e,
	0xd0, 0xf5, 0x04, 0xe5, 0x2c, 0x04, 0x03, 0x6e, 0x25, 0xe0, 0x90, 0xb3, 0x21, 0x13, 0x24, 0x34,
	0xc4, 0x3f, 0x09, 0xe1, 0x13, 0x49, 0x3c, 0x0e, 0x01, 0x15, 0x92, 0x4f, 0x66, 0x58, 0xb8, 0x84,
	0xfe, 0x48, 0x52, 0x16, 0x79, 0x67, 0x60, 0x4d, 0xff, 0x4d, 0xd8, 0x08, 0xe4, 0x05, 0xe3, 0x03,
	0x4f, 0x99, 0x03, 0x97, 0x14, 0xcc, 0x12, 0x1a, 0x3b, 0x69, 0xc4, 0x30, 0x64, 0x17, 0xe0, 0x7b,
	0xe7, 0x20, 0x04, 0x09, 0x12, 0x41, 0xea, 0x41, 0x7d, 0x88, 0x24, 0x95, 0x13, 0x9b, 0x81, 0x98,
	0x05, 0xfd, 0xf7, 0xbd, 0x8c, 0xaa, 0x6f, 0xe2, 0x8d, 0xf9, 0x20, 0x89, 0x04, 0xfc, 0x00, 0xd5,
	0x85, 0x24, 0x5c, 0xd2, 0x28, 0xf0, 0xec, 0x72, 0x3c, 0xea, 0x3b, 0xc5, 0x56, 0xb1, 0xbd, 0xe0,
	0x62, 0xcb, 0x1d, 0x1b, 0xaa, 0xe7, 0xe3, 0x16, 0xaa, 0x46, 0x70, 0x29, 0x3d, 0xb5, 0x23, 0x4a,
	0x39, 0xa7, 0x95, 0x48, 0x61, 0x2e, 0x0b, 0xa1, 0xe7, 0xe3, 0xff, 0x51, 0x49, 0x91, 0xc2, 0x99,
	0x6f, 0xcd, 0xb7, 0x2b, 0x7b, 0xcb, 0x1d, 0xbb, 0xf7, 0x1d, 0x25, 0x70, 0x63, 0x12, 0x9f, 0xa0,
	0x55, 0x6d, 0x31, 0x04, 0x7e, 0x4e, 0x85, 0xa0, 0x2c, 0x12, 0xce, 0x82, 0x1e, 0xd8, 0x4d, 0x07,
	0xb2, 0x59, 0xf5, 0xf4, 0x71, 0xaa, 0x3e, 0x8c, 0x24, 0x9f, 0xb8, 0x2b, 0x3c, 0x8f, 0xe2, 0xe7,
	0x68, 0xd9, 0x6e, 0xa1, 0xae, 0x52, 0x38, 0x25, 0xed, 0xba, 0x99, 0xba, 0xbe, 0x8f, 0xf9, 0x7d,
	0x45, 0xbb, 0xb5, 0x28, 0xf3, 0x26, 0xf0, 0x5b, 0x84, 0x67, 0x1b, 0x70, 0x16, 0x5b, 0xc5, 0x76,
	0x65, 0x6f, 0x7b, 0xc6, 0xe2, 0x38, 0x91, 0xb8, 0x6b, 0xd1, 0x34, 0xa4, 0xa2, 0xe4, 0xba, 0x16,
	0xce, 0x5f, 0xd3, 0x51, 0x0e, 0x2d, 0xff, 0x1a, 0xc0, 0xad, 0x41, 0xe6, 0x4d, 0xe0, 0x23, 0xb4,
	0x31, 0x64, 0x8c, 0x7b, 0x36, 0x8f, 0xad, 0xdb, 0x59, 0xd2, 0x69, 0xfe, 0x4e, 0x5d, 0xf6, 0xe3,
	0x03, 0x71, 0x64, 0x04, 0xee, 0xba, 0x9a, 0x33, 0x11, 0x2d, 0x88, 0x9f, 0xa0, 0xb2, 0x6d, 0x58,
	0x38, 0x65, 0x1d, 0x04, 0xa7, 0x16, 0xb6, 0xe1, 0x83, 0x85, 0xab, 0x9f, 0x3b, 0x05, 0x37, 0x95,
	0xe2, 0xfb, 0xa8, 0x34, 0x66, 0x12, 0x84, 0x83, 0xa6, 0xeb, 0x3c, 0x61, 0x12, 0x8c, 0x3e, 0x96,
	0xe0, 0x23, 0x54, 0xcb, 0x9d, 0x7d, 0xa7, 0xa2, 0x67, 0xda, 0x77, 0x34, 0xfa, 0x8a, 0x48, 0xe2,
	0x1a, 0x69, 0x5c, 0x67, 0xd5, 0xcf, 0x40, 0xb8, 0x87, 0x56, 0x33, 0x47, 0xb9, 0xcf, 0xb8, 0x2f,
	0x9c, 0xaa, 0x76, 0x74, 0x52, 0xc7, 0x9e, 0x51, 0xb8, 0x5a, 0x60, 0xf2, 0xac, 0xd0, 0x1c, 0x2a,
	0xf0, 0x63, 0xb4, 0x15, 0x12, 0x21, 0xbd, 0x29, 0x3f, 0x75, 0x82, 0x6b, 0xfa, 0x04, 0xd7, 0x15,
	0x9d, 0xf7, 0xea, 0xf9, 0xf8, 0x14, 0x35, 0xa8, 0x6f, 0xbf, 0xed, 0x8d, 0x81, 0xd3, 0x33, 0x35,
	0xfa, 0x75, 0x04, 0x42, 0x0a, 0x67, 0x59, 0x67, 0xd9, 0xb9, 0x2b, 0x8b, 0x38, 0xd1, 0x7a, 0x13,
	0x69, 0x8b, 0xfa, 0x39, 0xd8, 0x35, 0x2e, 0xf8, 0x10, 0xb5, 0x4c, 0x34, 0x1b, 0x2a, 0xff, 0x1d,
	0x95, 0x71, 0x45, 0x67, 0xdc, 0x8e, 0x33, 0xc6, 0x36, 0x39, 0x97, 0x9e, 0xdf, 0xf8, 0x88, 0xea,
	0x7f, 0xba, 0x21, 0x78, 0x15, 0xcd, 0x0f, 0x60, 0x62, 0x6e, 0xb4, 0x7a, 0xc4, 0xbb, 0xa8, 0x34,
	0x26, 0xe1, 0x08, 0xf4, 0xdd, 0xad, 0xec, 0x6d, 0x64, 0x4e, 0x41, 0x3a, 0xec, 0xc6, 0x9a, 0x67,
	0x73, 0x4f, 0x8b, 0x8d, 0xcf, 0x68, 0x6d, 0xa6, 0xaa, 0xac, 0x6f, 0x39, 0xf6, 0x7d, 0x98, 0xf7,
	0xcd, 0x5c, 0x97, 0xd9, 0xa2, 0x53, 0xf7, 0x83, 0x17, 0x57, 0x37, 0xcd, 0xe2, 0xf5, 0x4d, 0xb3,
	0xf8, 0xeb, 0xa6, 0x59, 0xfc, 0x76, 0xdb, 0x2c, 0x5c, 0xdf, 0x36, 0x0b, 0x3f, 0x6e, 0x9b, 0x85,
	0x4f, 0xf7, 0x02, 0x2a, 0xbf, 0x8c, 0x4e, 0x3b, 0x7d, 0x76, 0xde, 0x7d, 0x47, 0x39, 0x79, 0xc9,
	0x38, 0x74, 0x05, 0x0c, 0x08, 0xed, 0x5e, 0xea, 0x1f, 0x9d, 0x9c, 0x0c, 0x41, 0x9c, 0x2e, 0xea,
	0x9f, 0xdb, 0xa3, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8d, 0x1a, 0xc4, 0xab, 0xfd, 0x05, 0x00,
	0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LastIdRecordVerifyRequestId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.LastIdRecordVerifyRequestId))
		i--
		dAtA[i] = 0x78
	}
	if len(m.IdRecordsVerifyRequests) > 0 {
		for iNdEx := len(m.IdRecordsVerifyRequests) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.IdRecordsVerifyRequests[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x72
		}
	}
	if m.LastIdentityRecordId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.LastIdentityRecordId))
		i--
		dAtA[i] = 0x68
	}
	if len(m.IdentityRecords) > 0 {
		for iNdEx := len(m.IdentityRecords) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.IdentityRecords[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x62
		}
	}
	if len(m.DataRegistry) > 0 {
		for k := range m.DataRegistry {
			v := m.DataRegistry[k]
			baseI := i
			if v != nil {
				{
					size, err := v.MarshalToSizedBuffer(dAtA[:i])
					if err != nil {
						return 0, err
					}
					i -= size
					i = encodeVarintGenesis(dAtA, i, uint64(size))
				}
				i--
				dAtA[i] = 0x12
			}
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintGenesis(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintGenesis(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x5a
		}
	}
	if len(m.Votes) > 0 {
		for iNdEx := len(m.Votes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Votes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	if len(m.Proposals) > 0 {
		for iNdEx := len(m.Proposals) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Proposals[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if m.PoorNetworkMessages != nil {
		{
			size, err := m.PoorNetworkMessages.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if len(m.ExecutionFees) > 0 {
		for iNdEx := len(m.ExecutionFees) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ExecutionFees[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if m.NetworkProperties != nil {
		{
			size, err := m.NetworkProperties.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if len(m.NetworkActors) > 0 {
		for iNdEx := len(m.NetworkActors) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.NetworkActors[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.RolePermissions) > 0 {
		for k := range m.RolePermissions {
			v := m.RolePermissions[k]
			baseI := i
			if v != nil {
				{
					size, err := v.MarshalToSizedBuffer(dAtA[:i])
					if err != nil {
						return 0, err
					}
					i -= size
					i = encodeVarintGenesis(dAtA, i, uint64(size))
				}
				i--
				dAtA[i] = 0x12
			}
			i = encodeVarintGenesis(dAtA, i, uint64(k))
			i--
			dAtA[i] = 0x8
			i = encodeVarintGenesis(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Roles) > 0 {
		for iNdEx := len(m.Roles) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Roles[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.NextRoleId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.NextRoleId))
		i--
		dAtA[i] = 0x10
	}
	if m.StartingProposalId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.StartingProposalId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StartingProposalId != 0 {
		n += 1 + sovGenesis(uint64(m.StartingProposalId))
	}
	if m.NextRoleId != 0 {
		n += 1 + sovGenesis(uint64(m.NextRoleId))
	}
	if len(m.Roles) > 0 {
		for _, e := range m.Roles {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.RolePermissions) > 0 {
		for k, v := range m.RolePermissions {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovGenesis(uint64(l))
			}
			mapEntrySize := 1 + sovGenesis(uint64(k)) + l
			n += mapEntrySize + 1 + sovGenesis(uint64(mapEntrySize))
		}
	}
	if len(m.NetworkActors) > 0 {
		for _, e := range m.NetworkActors {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.NetworkProperties != nil {
		l = m.NetworkProperties.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.ExecutionFees) > 0 {
		for _, e := range m.ExecutionFees {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.PoorNetworkMessages != nil {
		l = m.PoorNetworkMessages.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.Proposals) > 0 {
		for _, e := range m.Proposals {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Votes) > 0 {
		for _, e := range m.Votes {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.DataRegistry) > 0 {
		for k, v := range m.DataRegistry {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovGenesis(uint64(l))
			}
			mapEntrySize := 1 + len(k) + sovGenesis(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovGenesis(uint64(mapEntrySize))
		}
	}
	if len(m.IdentityRecords) > 0 {
		for _, e := range m.IdentityRecords {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.LastIdentityRecordId != 0 {
		n += 1 + sovGenesis(uint64(m.LastIdentityRecordId))
	}
	if len(m.IdRecordsVerifyRequests) > 0 {
		for _, e := range m.IdRecordsVerifyRequests {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.LastIdRecordVerifyRequestId != 0 {
		n += 1 + sovGenesis(uint64(m.LastIdRecordVerifyRequestId))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartingProposalId", wireType)
			}
			m.StartingProposalId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartingProposalId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextRoleId", wireType)
			}
			m.NextRoleId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NextRoleId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Roles", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Roles = append(m.Roles, &Role{})
			if err := m.Roles[len(m.Roles)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RolePermissions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RolePermissions == nil {
				m.RolePermissions = make(map[uint64]*Permissions)
			}
			var mapkey uint64
			var mapvalue *Permissions
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenesis
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
				if fieldNum == 1 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenesis
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenesis
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthGenesis
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthGenesis
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &Permissions{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipGenesis(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthGenesis
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.RolePermissions[mapkey] = mapvalue
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NetworkActors", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NetworkActors = append(m.NetworkActors, &NetworkActor{})
			if err := m.NetworkActors[len(m.NetworkActors)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NetworkProperties", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.NetworkProperties == nil {
				m.NetworkProperties = &NetworkProperties{}
			}
			if err := m.NetworkProperties.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExecutionFees", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExecutionFees = append(m.ExecutionFees, &ExecutionFee{})
			if err := m.ExecutionFees[len(m.ExecutionFees)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoorNetworkMessages", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PoorNetworkMessages == nil {
				m.PoorNetworkMessages = &AllowedMessages{}
			}
			if err := m.PoorNetworkMessages.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proposals", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Proposals = append(m.Proposals, Proposal{})
			if err := m.Proposals[len(m.Proposals)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Votes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Votes = append(m.Votes, Vote{})
			if err := m.Votes[len(m.Votes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataRegistry", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DataRegistry == nil {
				m.DataRegistry = make(map[string]*DataRegistryEntry)
			}
			var mapkey string
			var mapvalue *DataRegistryEntry
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenesis
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenesis
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthGenesis
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthGenesis
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenesis
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthGenesis
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthGenesis
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &DataRegistryEntry{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipGenesis(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthGenesis
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.DataRegistry[mapkey] = mapvalue
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IdentityRecords", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IdentityRecords = append(m.IdentityRecords, IdentityRecord{})
			if err := m.IdentityRecords[len(m.IdentityRecords)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastIdentityRecordId", wireType)
			}
			m.LastIdentityRecordId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastIdentityRecordId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IdRecordsVerifyRequests", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IdRecordsVerifyRequests = append(m.IdRecordsVerifyRequests, IdentityRecordsVerify{})
			if err := m.IdRecordsVerifyRequests[len(m.IdRecordsVerifyRequests)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 15:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastIdRecordVerifyRequestId", wireType)
			}
			m.LastIdRecordVerifyRequestId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastIdRecordVerifyRequestId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
