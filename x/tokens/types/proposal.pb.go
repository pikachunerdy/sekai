// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proposal.proto

package types

import (
	fmt "fmt"
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

type MsgProposalUpsertTokenAlias struct {
	Symbol   string   `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Name     string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Icon     string   `protobuf:"bytes,3,opt,name=icon,proto3" json:"icon,omitempty"`
	Decimals uint32   `protobuf:"varint,4,opt,name=decimals,proto3" json:"decimals,omitempty"`
	Denoms   []string `protobuf:"bytes,5,rep,name=denoms,proto3" json:"denoms,omitempty"`
}

func (m *MsgProposalUpsertTokenAlias) Reset()         { *m = MsgProposalUpsertTokenAlias{} }
func (m *MsgProposalUpsertTokenAlias) String() string { return proto.CompactTextString(m) }
func (*MsgProposalUpsertTokenAlias) ProtoMessage()    {}
func (*MsgProposalUpsertTokenAlias) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3ac5ce23bf32d05, []int{0}
}
func (m *MsgProposalUpsertTokenAlias) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgProposalUpsertTokenAlias) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgProposalUpsertTokenAlias.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgProposalUpsertTokenAlias) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgProposalUpsertTokenAlias.Merge(m, src)
}
func (m *MsgProposalUpsertTokenAlias) XXX_Size() int {
	return m.Size()
}
func (m *MsgProposalUpsertTokenAlias) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgProposalUpsertTokenAlias.DiscardUnknown(m)
}

var xxx_messageInfo_MsgProposalUpsertTokenAlias proto.InternalMessageInfo

func (m *MsgProposalUpsertTokenAlias) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *MsgProposalUpsertTokenAlias) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MsgProposalUpsertTokenAlias) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *MsgProposalUpsertTokenAlias) GetDecimals() uint32 {
	if m != nil {
		return m.Decimals
	}
	return 0
}

func (m *MsgProposalUpsertTokenAlias) GetDenoms() []string {
	if m != nil {
		return m.Denoms
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgProposalUpsertTokenAlias)(nil), "kira.tokens.MsgProposalUpsertTokenAlias")
}

func init() { proto.RegisterFile("proposal.proto", fileDescriptor_c3ac5ce23bf32d05) }

var fileDescriptor_c3ac5ce23bf32d05 = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x28, 0xca, 0x2f,
	0xc8, 0x2f, 0x4e, 0xcc, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xce, 0xce, 0x2c, 0x4a,
	0xd4, 0x2b, 0xc9, 0xcf, 0x4e, 0xcd, 0x2b, 0x56, 0x9a, 0xc8, 0xc8, 0x25, 0xed, 0x5b, 0x9c, 0x1e,
	0x00, 0x55, 0x12, 0x5a, 0x50, 0x9c, 0x5a, 0x54, 0x12, 0x02, 0x92, 0x73, 0xcc, 0xc9, 0x4c, 0x2c,
	0x16, 0x12, 0xe3, 0x62, 0x2b, 0xae, 0xcc, 0x4d, 0xca, 0xcf, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0,
	0x0c, 0x82, 0xf2, 0x84, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x98, 0xc0, 0xa2, 0x60,
	0x36, 0x48, 0x2c, 0x33, 0x39, 0x3f, 0x4f, 0x82, 0x19, 0x22, 0x06, 0x62, 0x0b, 0x49, 0x71, 0x71,
	0xa4, 0xa4, 0x26, 0x67, 0xe6, 0x26, 0xe6, 0x14, 0x4b, 0xb0, 0x28, 0x30, 0x6a, 0xf0, 0x06, 0xc1,
	0xf9, 0x20, 0xb3, 0x53, 0x52, 0xf3, 0xf2, 0x73, 0x8b, 0x25, 0x58, 0x15, 0x98, 0x41, 0x66, 0x43,
	0x78, 0x4e, 0x4e, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3,
	0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0xa5, 0x91, 0x9e,
	0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0xef, 0x9d, 0x59, 0x94, 0xe8, 0x9c, 0x5f,
	0x94, 0xaa, 0x5f, 0x9c, 0x9a, 0x9d, 0x98, 0xa9, 0x5f, 0xa1, 0x0f, 0xf1, 0x91, 0x7e, 0x49, 0x65,
	0x41, 0x6a, 0x71, 0x12, 0x1b, 0xd8, 0xaf, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd7, 0xb5,
	0x43, 0x8f, 0xfd, 0x00, 0x00, 0x00,
}

func (m *MsgProposalUpsertTokenAlias) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgProposalUpsertTokenAlias) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgProposalUpsertTokenAlias) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denoms) > 0 {
		for iNdEx := len(m.Denoms) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Denoms[iNdEx])
			copy(dAtA[i:], m.Denoms[iNdEx])
			i = encodeVarintProposal(dAtA, i, uint64(len(m.Denoms[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.Decimals != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.Decimals))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Icon) > 0 {
		i -= len(m.Icon)
		copy(dAtA[i:], m.Icon)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Icon)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProposal(dAtA []byte, offset int, v uint64) int {
	offset -= sovProposal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgProposalUpsertTokenAlias) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Icon)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	if m.Decimals != 0 {
		n += 1 + sovProposal(uint64(m.Decimals))
	}
	if len(m.Denoms) > 0 {
		for _, s := range m.Denoms {
			l = len(s)
			n += 1 + l + sovProposal(uint64(l))
		}
	}
	return n
}

func sovProposal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProposal(x uint64) (n int) {
	return sovProposal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgProposalUpsertTokenAlias) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: MsgProposalUpsertTokenAlias: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgProposalUpsertTokenAlias: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Icon", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Icon = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Decimals", wireType)
			}
			m.Decimals = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Decimals |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denoms", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denoms = append(m.Denoms, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProposal
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProposal
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
func skipProposal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
				return 0, ErrInvalidLengthProposal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProposal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProposal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProposal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProposal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProposal = fmt.Errorf("proto: unexpected end of group")
)
