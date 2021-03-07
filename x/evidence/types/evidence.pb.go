// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: evidence.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Equivocation implements the Evidence interface and defines evidence of double
// signing misbehavior.
type Equivocation struct {
	Height           int64     `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Time             time.Time `protobuf:"bytes,2,opt,name=time,proto3,stdtime" json:"time"`
	Power            int64     `protobuf:"varint,3,opt,name=power,proto3" json:"power,omitempty"`
	ConsensusAddress string    `protobuf:"bytes,4,opt,name=consensus_address,json=consensusAddress,proto3" json:"consensus_address,omitempty" yaml:"consensus_address"`
}

func (m *Equivocation) Reset()      { *m = Equivocation{} }
func (*Equivocation) ProtoMessage() {}
func (*Equivocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b1d6725573e3e5a, []int{0}
}
func (m *Equivocation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Equivocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Equivocation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Equivocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Equivocation.Merge(m, src)
}
func (m *Equivocation) XXX_Size() int {
	return m.Size()
}
func (m *Equivocation) XXX_DiscardUnknown() {
	xxx_messageInfo_Equivocation.DiscardUnknown(m)
}

var xxx_messageInfo_Equivocation proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Equivocation)(nil), "kira.evidence.Equivocation")
}

func init() { proto.RegisterFile("evidence.proto", fileDescriptor_9b1d6725573e3e5a) }

var fileDescriptor_9b1d6725573e3e5a = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0x2d, 0xcb, 0x4c,
	0x49, 0xcd, 0x4b, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xcd, 0xce, 0x2c, 0x4a,
	0xd4, 0x83, 0x09, 0x4a, 0x89, 0xa4, 0xe7, 0xa7, 0xe7, 0x83, 0x65, 0xf4, 0x41, 0x2c, 0x88, 0x22,
	0x29, 0xf9, 0xf4, 0xfc, 0xfc, 0xf4, 0x9c, 0x54, 0x7d, 0x30, 0x2f, 0xa9, 0x34, 0x4d, 0xbf, 0x24,
	0x33, 0x37, 0xb5, 0xb8, 0x24, 0x31, 0xb7, 0x00, 0xa2, 0x40, 0xe9, 0x3c, 0x23, 0x17, 0x8f, 0x6b,
	0x61, 0x69, 0x66, 0x59, 0x7e, 0x72, 0x62, 0x49, 0x66, 0x7e, 0x9e, 0x90, 0x18, 0x17, 0x5b, 0x46,
	0x6a, 0x66, 0x7a, 0x46, 0x89, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x73, 0x10, 0x94, 0x27, 0x64, 0xc1,
	0xc5, 0x02, 0xd2, 0x2b, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x6d, 0x24, 0xa5, 0x07, 0x31, 0x58, 0x0f,
	0x66, 0xb0, 0x5e, 0x08, 0xcc, 0x60, 0x27, 0x8e, 0x13, 0xf7, 0xe4, 0x19, 0x26, 0xdc, 0x97, 0x67,
	0x0c, 0x02, 0xeb, 0x10, 0x12, 0xe1, 0x62, 0x2d, 0xc8, 0x2f, 0x4f, 0x2d, 0x92, 0x60, 0x06, 0x1b,
	0x08, 0xe1, 0x08, 0x79, 0x72, 0x09, 0x26, 0xe7, 0xe7, 0x15, 0xa7, 0xe6, 0x15, 0x97, 0x16, 0xc7,
	0x27, 0xa6, 0xa4, 0x14, 0xa5, 0x16, 0x17, 0x4b, 0xb0, 0x28, 0x30, 0x6a, 0x70, 0x3a, 0xc9, 0x7c,
	0xba, 0x27, 0x2f, 0x51, 0x99, 0x98, 0x9b, 0x63, 0xa5, 0x84, 0xa1, 0x44, 0x29, 0x48, 0x00, 0x2e,
	0xe6, 0x08, 0x11, 0xb2, 0xe2, 0xe9, 0x58, 0x20, 0xcf, 0x30, 0x63, 0x81, 0x3c, 0xc3, 0x8b, 0x05,
	0xf2, 0x0c, 0x4e, 0x1e, 0x2b, 0x1e, 0xc9, 0x31, 0x9e, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c,
	0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1,
	0x1c, 0x43, 0x94, 0x56, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0xbe, 0x77,
	0x66, 0x51, 0xa2, 0x73, 0x7e, 0x51, 0xaa, 0x7e, 0x71, 0x6a, 0x76, 0x62, 0xa6, 0x7e, 0x85, 0x3e,
	0x2c, 0x30, 0xf5, 0x4b, 0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0x9e, 0x33, 0x06, 0x04, 0x00,
	0x00, 0xff, 0xff, 0xc0, 0x87, 0xf8, 0x4d, 0x7a, 0x01, 0x00, 0x00,
}

func (m *Equivocation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Equivocation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Equivocation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ConsensusAddress) > 0 {
		i -= len(m.ConsensusAddress)
		copy(dAtA[i:], m.ConsensusAddress)
		i = encodeVarintEvidence(dAtA, i, uint64(len(m.ConsensusAddress)))
		i--
		dAtA[i] = 0x22
	}
	if m.Power != 0 {
		i = encodeVarintEvidence(dAtA, i, uint64(m.Power))
		i--
		dAtA[i] = 0x18
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Time, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Time):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintEvidence(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x12
	if m.Height != 0 {
		i = encodeVarintEvidence(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvidence(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvidence(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Equivocation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovEvidence(uint64(m.Height))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Time)
	n += 1 + l + sovEvidence(uint64(l))
	if m.Power != 0 {
		n += 1 + sovEvidence(uint64(m.Power))
	}
	l = len(m.ConsensusAddress)
	if l > 0 {
		n += 1 + l + sovEvidence(uint64(l))
	}
	return n
}

func sovEvidence(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvidence(x uint64) (n int) {
	return sovEvidence(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Equivocation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvidence
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
			return fmt.Errorf("proto: Equivocation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Equivocation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvidence
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvidence
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
				return ErrInvalidLengthEvidence
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvidence
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Time, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Power", wireType)
			}
			m.Power = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvidence
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Power |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsensusAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvidence
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
				return ErrInvalidLengthEvidence
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvidence
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConsensusAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvidence(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEvidence
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthEvidence
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
func skipEvidence(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvidence
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
					return 0, ErrIntOverflowEvidence
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
					return 0, ErrIntOverflowEvidence
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
				return 0, ErrInvalidLengthEvidence
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvidence
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvidence
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvidence        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvidence          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvidence = fmt.Errorf("proto: unexpected end of group")
)
