package types

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kira/gov/genesis.proto

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"

	govtypes "github.com/KiraCore/sekai/x/gov/types"
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

type GenesisStateV01228 struct {
	// starting_proposal_id is the ID of the starting proposal.
	StartingProposalId uint64 `protobuf:"varint,1,opt,name=starting_proposal_id,json=startingProposalId,proto3" json:"starting_proposal_id,omitempty"`
	// permissions is the roles that are active from genesis.
	Permissions map[uint64]*govtypes.Permissions `protobuf:"bytes,2,rep,name=permissions,proto3" json:"permissions,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// NetworkActors are the actors that are saved from genesis.
	NetworkActors               []*govtypes.NetworkActor               `protobuf:"bytes,3,rep,name=network_actors,json=networkActors,proto3" json:"network_actors,omitempty"`
	NetworkProperties           *NetworkPropertiesV0228                `protobuf:"bytes,4,opt,name=network_properties,json=networkProperties,proto3" json:"network_properties,omitempty"`
	ExecutionFees               []*govtypes.ExecutionFee               `protobuf:"bytes,5,rep,name=execution_fees,json=executionFees,proto3" json:"execution_fees,omitempty"`
	PoorNetworkMessages         *govtypes.AllowedMessages              `protobuf:"bytes,6,opt,name=poor_network_messages,json=poorNetworkMessages,proto3" json:"poor_network_messages,omitempty"`
	Proposals                   []govtypes.Proposal                    `protobuf:"bytes,7,rep,name=proposals,proto3" json:"proposals"`
	Votes                       []govtypes.Vote                        `protobuf:"bytes,8,rep,name=votes,proto3" json:"votes"`
	DataRegistry                map[string]*govtypes.DataRegistryEntry `protobuf:"bytes,9,rep,name=data_registry,json=dataRegistry,proto3" json:"data_registry,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	IdentityRecords             []govtypes.IdentityRecord              `protobuf:"bytes,10,rep,name=identity_records,json=identityRecords,proto3" json:"identity_records"`
	LastIdentityRecordId        uint64                                 `protobuf:"varint,11,opt,name=last_identity_record_id,json=lastIdentityRecordId,proto3" json:"last_identity_record_id,omitempty"`
	IdRecordsVerifyRequests     []govtypes.IdentityRecordsVerify       `protobuf:"bytes,12,rep,name=id_records_verify_requests,json=idRecordsVerifyRequests,proto3" json:"id_records_verify_requests"`
	LastIdRecordVerifyRequestId uint64                                 `protobuf:"varint,13,opt,name=last_id_record_verify_request_id,json=lastIdRecordVerifyRequestId,proto3" json:"last_id_record_verify_request_id,omitempty"`
}

func (m *GenesisStateV01228) Reset()         { *m = GenesisStateV01228{} }
func (m *GenesisStateV01228) String() string { return proto.CompactTextString(m) }
func (*GenesisStateV01228) ProtoMessage()    {}
func (*GenesisStateV01228) Descriptor() ([]byte, []int) {
	return fileDescriptor_042721fb65a4ea2d, []int{0}
}
func (m *GenesisStateV01228) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisStateV01228) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
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
func (m *GenesisStateV01228) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisStateV01228) XXX_Size() int {
	return m.Size()
}
func (m *GenesisStateV01228) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisStateV01228) GetStartingProposalId() uint64 {
	if m != nil {
		return m.StartingProposalId
	}
	return 0
}

func (m *GenesisStateV01228) GetPermissions() map[uint64]*govtypes.Permissions {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func (m *GenesisStateV01228) GetNetworkActors() []*govtypes.NetworkActor {
	if m != nil {
		return m.NetworkActors
	}
	return nil
}

func (m *GenesisStateV01228) GetNetworkProperties() *NetworkPropertiesV0228 {
	if m != nil {
		return m.NetworkProperties
	}
	return nil
}

func (m *GenesisStateV01228) GetExecutionFees() []*govtypes.ExecutionFee {
	if m != nil {
		return m.ExecutionFees
	}
	return nil
}

func (m *GenesisStateV01228) GetPoorNetworkMessages() *govtypes.AllowedMessages {
	if m != nil {
		return m.PoorNetworkMessages
	}
	return nil
}

func (m *GenesisStateV01228) GetProposals() []govtypes.Proposal {
	if m != nil {
		return m.Proposals
	}
	return nil
}

func (m *GenesisStateV01228) GetVotes() []govtypes.Vote {
	if m != nil {
		return m.Votes
	}
	return nil
}

func (m *GenesisStateV01228) GetDataRegistry() map[string]*govtypes.DataRegistryEntry {
	if m != nil {
		return m.DataRegistry
	}
	return nil
}

func (m *GenesisStateV01228) GetIdentityRecords() []govtypes.IdentityRecord {
	if m != nil {
		return m.IdentityRecords
	}
	return nil
}

func (m *GenesisStateV01228) GetLastIdentityRecordId() uint64 {
	if m != nil {
		return m.LastIdentityRecordId
	}
	return 0
}

func (m *GenesisStateV01228) GetIdRecordsVerifyRequests() []govtypes.IdentityRecordsVerify {
	if m != nil {
		return m.IdRecordsVerifyRequests
	}
	return nil
}

func (m *GenesisStateV01228) GetLastIdRecordVerifyRequestId() uint64 {
	if m != nil {
		return m.LastIdRecordVerifyRequestId
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisStateV01228)(nil), "kira.gov.GenesisStateV01228")
}

func init() { proto.RegisterFile("kira/gov/genesis.proto", fileDescriptor_042721fb65a4ea2d) }

var fileDescriptor_042721fb65a4ea2d = []byte{
	// 666 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0x5d, 0x4f, 0x13, 0x41,
	0x14, 0xed, 0xf2, 0x25, 0x4c, 0x29, 0xc2, 0x50, 0x60, 0x2c, 0xa6, 0x54, 0x13, 0x63, 0xa3, 0x49,
	0xab, 0x18, 0x8d, 0x31, 0x31, 0x06, 0x14, 0x4d, 0x35, 0x18, 0xb2, 0x46, 0x1e, 0x8c, 0xc9, 0x66,
	0xe8, 0x5e, 0xd6, 0x49, 0xcb, 0xce, 0x3a, 0x33, 0x2d, 0xf4, 0x5f, 0xf8, 0x77, 0xfc, 0x07, 0x3c,
	0xf2, 0xe8, 0x93, 0x31, 0xf4, 0x8f, 0x98, 0x9d, 0x9d, 0xd9, 0x8f, 0x36, 0xbc, 0xed, 0x9e, 0x73,
	0xee, 0xd9, 0x33, 0xf7, 0xde, 0x59, 0xb4, 0xd9, 0x63, 0x82, 0xb6, 0x03, 0x3e, 0x6c, 0x07, 0x10,
	0x82, 0x64, 0xb2, 0x15, 0x09, 0xae, 0x38, 0x5e, 0x8c, 0xf1, 0x56, 0xc0, 0x87, 0xb5, 0x6a, 0xc0,
	0x03, 0xae, 0xc1, 0x76, 0xfc, 0x94, 0xf0, 0xb5, 0x6a, 0x5a, 0x47, 0xbb, 0x8a, 0x0b, 0x83, 0xae,
	0xa7, 0xa8, 0xe0, 0x7d, 0x30, 0xe0, 0x56, 0x0a, 0x46, 0x82, 0x47, 0x5c, 0xd2, 0xbe, 0x21, 0xee,
	0xa6, 0x84, 0x4f, 0x15, 0xf5, 0x04, 0x04, 0x4c, 0x2a, 0x31, 0x9a, 0x62, 0xe1, 0x02, 0xba, 0x03,
	0xc5, 0x78, 0xe8, 0x9d, 0x82, 0x35, 0xbd, 0x97, 0xb2, 0x21, 0xa8, 0x73, 0x2e, 0x7a, 0x5e, 0x6c,
	0x0e, 0x42, 0x31, 0x30, 0x47, 0xa8, 0xed, 0x64, 0x11, 0xfb, 0x7d, 0x7e, 0x0e, 0xbe, 0x77, 0x06,
	0x52, 0xd2, 0x20, 0x15, 0x64, 0x1e, 0xcc, 0x87, 0x50, 0x31, 0x35, 0xb2, 0x19, 0xa8, 0x39, 0xd0,
	0xfd, 0xdf, 0x8b, 0x68, 0xf9, 0x43, 0xd2, 0x98, 0x2f, 0x8a, 0x2a, 0xc0, 0x4f, 0x50, 0x55, 0x2a,
	0x2a, 0x14, 0x0b, 0x03, 0xcf, 0x1e, 0xc7, 0x63, 0x3e, 0x71, 0x1a, 0x4e, 0x73, 0xce, 0xc5, 0x96,
	0x3b, 0x32, 0x54, 0xc7, 0xc7, 0x1d, 0x54, 0x8e, 0x40, 0x9c, 0x31, 0x29, 0x19, 0x0f, 0x25, 0x99,
	0x69, 0xcc, 0x36, 0xcb, 0xbb, 0x0f, 0x5b, 0xb6, 0xbf, 0xad, 0xbc, 0x7d, 0xeb, 0x28, 0x53, 0x1e,
	0x84, 0x4a, 0x8c, 0xdc, 0x7c, 0x2d, 0x7e, 0x8d, 0x56, 0xec, 0x69, 0x75, 0xd7, 0x25, 0x99, 0xd5,
	0x6e, 0x9b, 0x99, 0xdb, 0xe7, 0x84, 0xdf, 0x8b, 0x69, 0xb7, 0x12, 0xe6, 0xde, 0x24, 0xfe, 0x88,
	0xf0, 0x74, 0xb3, 0xc8, 0x5c, 0xc3, 0x69, 0x96, 0x77, 0xb7, 0xa7, 0x2c, 0x8e, 0x52, 0x89, 0xbb,
	0x16, 0x4e, 0x42, 0x71, 0x94, 0xc2, 0x58, 0x24, 0x99, 0x9f, 0x8c, 0x72, 0x60, 0xf9, 0xf7, 0x00,
	0x6e, 0x05, 0x72, 0x6f, 0x12, 0x1f, 0xa2, 0x8d, 0x88, 0x73, 0xe1, 0xd9, 0x3c, 0x76, 0x32, 0x64,
	0x41, 0xa7, 0xb9, 0x93, 0xb9, 0xec, 0x25, 0xb3, 0x3b, 0x34, 0x02, 0x77, 0x3d, 0xae, 0x33, 0x11,
	0x2d, 0x88, 0x5f, 0xa0, 0x25, 0x3b, 0x0c, 0x49, 0x6e, 0xe9, 0x20, 0x38, 0xb3, 0xb0, 0xc3, 0xd8,
	0x9f, 0xbb, 0xfc, 0xbb, 0x53, 0x72, 0x33, 0x29, 0x7e, 0x84, 0xe6, 0x87, 0x5c, 0x81, 0x24, 0x8b,
	0xba, 0x66, 0x25, 0xab, 0x39, 0xe6, 0x0a, 0x8c, 0x3e, 0x91, 0xe0, 0x43, 0x54, 0x29, 0xac, 0x29,
	0x59, 0xd2, 0x35, 0xcd, 0x1b, 0x26, 0xf9, 0x8e, 0x2a, 0xea, 0x1a, 0x69, 0x32, 0xca, 0x65, 0x3f,
	0x07, 0xe1, 0x0e, 0x5a, 0xcd, 0x6d, 0x5d, 0x97, 0x0b, 0x5f, 0x12, 0xa4, 0x1d, 0x49, 0xe6, 0xd8,
	0x31, 0x0a, 0x57, 0x0b, 0x4c, 0x9e, 0xdb, 0xac, 0x80, 0x4a, 0xfc, 0x1c, 0x6d, 0xf5, 0xa9, 0x54,
	0xde, 0x84, 0x5f, 0xbc, 0x96, 0x65, 0xbd, 0x96, 0xd5, 0x98, 0x2e, 0x7a, 0x75, 0x7c, 0x7c, 0x82,
	0x6a, 0xcc, 0xb7, 0xdf, 0xf6, 0x86, 0x20, 0xd8, 0x69, 0x5c, 0xfa, 0x73, 0x00, 0x52, 0x49, 0xb2,
	0xac, 0xb3, 0xec, 0xdc, 0x94, 0x45, 0x1e, 0x6b, 0xbd, 0x89, 0xb4, 0xc5, 0xfc, 0x02, 0xec, 0x1a,
	0x17, 0x7c, 0x80, 0x1a, 0x26, 0x9a, 0x0d, 0x55, 0xfc, 0x4e, 0x9c, 0xb1, 0xa2, 0x33, 0x6e, 0x27,
	0x19, 0x13, 0x9b, 0x82, 0x4b, 0xc7, 0xaf, 0x7d, 0x45, 0xab, 0x93, 0x37, 0x03, 0xaf, 0xa2, 0xd9,
	0x1e, 0x8c, 0xcc, 0xc5, 0x8b, 0x1f, 0xf1, 0x63, 0x34, 0x3f, 0xa4, 0xfd, 0x01, 0x90, 0x19, 0xbd,
	0x44, 0x1b, 0xb9, 0x0d, 0xc8, 0x8a, 0xdd, 0x44, 0xf3, 0x6a, 0xe6, 0xa5, 0x53, 0xfb, 0x8e, 0xd6,
	0xa6, 0xc6, 0x94, 0xf7, 0x5d, 0x4a, 0x7c, 0x9f, 0x16, 0x7d, 0x73, 0x57, 0x65, 0x7a, 0xc8, 0x99,
	0xfb, 0xfe, 0x9b, 0xcb, 0xeb, 0xba, 0x73, 0x75, 0x5d, 0x77, 0xfe, 0x5d, 0xd7, 0x9d, 0x5f, 0xe3,
	0x7a, 0xe9, 0x6a, 0x5c, 0x2f, 0xfd, 0x19, 0xd7, 0x4b, 0xdf, 0x1e, 0x04, 0x4c, 0xfd, 0x18, 0x9c,
	0xb4, 0xba, 0xfc, 0xac, 0xfd, 0x89, 0x09, 0xfa, 0x96, 0x0b, 0x68, 0x4b, 0xe8, 0x51, 0xd6, 0xbe,
	0xd0, 0xff, 0x23, 0x35, 0x8a, 0x40, 0x9e, 0x2c, 0xe8, 0x7f, 0xd0, 0xb3, 0xff, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x03, 0x6d, 0xd7, 0x16, 0xa4, 0x05, 0x00, 0x00,
}

func (m *GenesisStateV01228) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisStateV01228) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisStateV01228) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LastIdRecordVerifyRequestId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.LastIdRecordVerifyRequestId))
		i--
		dAtA[i] = 0x68
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
			dAtA[i] = 0x62
		}
	}
	if m.LastIdentityRecordId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.LastIdentityRecordId))
		i--
		dAtA[i] = 0x58
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
			dAtA[i] = 0x52
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
			dAtA[i] = 0x4a
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
			dAtA[i] = 0x42
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
			dAtA[i] = 0x3a
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
		dAtA[i] = 0x32
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
			dAtA[i] = 0x2a
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
		dAtA[i] = 0x22
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
			dAtA[i] = 0x1a
		}
	}
	if len(m.Permissions) > 0 {
		for k := range m.Permissions {
			v := m.Permissions[k]
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
			dAtA[i] = 0x12
		}
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
func (m *GenesisStateV01228) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StartingProposalId != 0 {
		n += 1 + sovGenesis(uint64(m.StartingProposalId))
	}
	if len(m.Permissions) > 0 {
		for k, v := range m.Permissions {
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
func (m *GenesisStateV01228) Unmarshal(dAtA []byte) error {
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Permissions", wireType)
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
			if m.Permissions == nil {
				m.Permissions = make(map[uint64]*govtypes.Permissions)
			}
			var mapkey uint64
			var mapvalue *govtypes.Permissions
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
					mapvalue = &govtypes.Permissions{}
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
			m.Permissions[mapkey] = mapvalue
			iNdEx = postIndex
		case 3:
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
			m.NetworkActors = append(m.NetworkActors, &govtypes.NetworkActor{})
			if err := m.NetworkActors[len(m.NetworkActors)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
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
				m.NetworkProperties = &NetworkPropertiesV0228{}
			}
			if err := m.NetworkProperties.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
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
			m.ExecutionFees = append(m.ExecutionFees, &govtypes.ExecutionFee{})
			if err := m.ExecutionFees[len(m.ExecutionFees)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
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
				m.PoorNetworkMessages = &govtypes.AllowedMessages{}
			}
			if err := m.PoorNetworkMessages.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
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
			m.Proposals = append(m.Proposals, govtypes.Proposal{})
			if err := m.Proposals[len(m.Proposals)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
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
			m.Votes = append(m.Votes, govtypes.Vote{})
			if err := m.Votes[len(m.Votes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
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
				m.DataRegistry = make(map[string]*govtypes.DataRegistryEntry)
			}
			var mapkey string
			var mapvalue *govtypes.DataRegistryEntry
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
					mapvalue = &govtypes.DataRegistryEntry{}
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
		case 10:
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
			m.IdentityRecords = append(m.IdentityRecords, govtypes.IdentityRecord{})
			if err := m.IdentityRecords[len(m.IdentityRecords)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
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
		case 12:
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
			m.IdRecordsVerifyRequests = append(m.IdRecordsVerifyRequests, govtypes.IdentityRecordsVerify{})
			if err := m.IdRecordsVerifyRequests[len(m.IdRecordsVerifyRequests)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
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
