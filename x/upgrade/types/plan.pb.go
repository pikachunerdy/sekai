// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: plan.proto

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

type Plan struct {
	MinHaltTime          int64  `protobuf:"varint,1,opt,name=min_halt_time,json=minHaltTime,proto3" json:"min_halt_time,omitempty"`
	RollbackChecksum     string `protobuf:"bytes,2,opt,name=rollback_checksum,json=rollbackChecksum,proto3" json:"rollback_checksum,omitempty"`
	MaxEnrolmentDuration int64  `protobuf:"varint,3,opt,name=max_enrolment_duration,json=maxEnrolmentDuration,proto3" json:"max_enrolment_duration,omitempty"`
	Name                 string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *Plan) Reset()         { *m = Plan{} }
func (m *Plan) String() string { return proto.CompactTextString(m) }
func (*Plan) ProtoMessage()    {}
func (*Plan) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d655ab2f7683c23, []int{0}
}
func (m *Plan) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Plan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Plan.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Plan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Plan.Merge(m, src)
}
func (m *Plan) XXX_Size() int {
	return m.Size()
}
func (m *Plan) XXX_DiscardUnknown() {
	xxx_messageInfo_Plan.DiscardUnknown(m)
}

var xxx_messageInfo_Plan proto.InternalMessageInfo

func (m *Plan) GetMinHaltTime() int64 {
	if m != nil {
		return m.MinHaltTime
	}
	return 0
}

func (m *Plan) GetRollbackChecksum() string {
	if m != nil {
		return m.RollbackChecksum
	}
	return ""
}

func (m *Plan) GetMaxEnrolmentDuration() int64 {
	if m != nil {
		return m.MaxEnrolmentDuration
	}
	return 0
}

func (m *Plan) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Plan)(nil), "kira.upgrade.Plan")
}

func init() { proto.RegisterFile("plan.proto", fileDescriptor_2d655ab2f7683c23) }

var fileDescriptor_2d655ab2f7683c23 = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8f, 0x41, 0x4a, 0xc4, 0x30,
	0x14, 0x86, 0x1b, 0xa7, 0x08, 0x46, 0x05, 0x0d, 0x22, 0x5d, 0x85, 0x61, 0x56, 0x23, 0x42, 0xbb,
	0xd0, 0x13, 0x58, 0x05, 0xc1, 0x8d, 0x0c, 0xae, 0xdc, 0x84, 0xd7, 0xce, 0x63, 0x1a, 0x9a, 0x97,
	0x94, 0x34, 0x85, 0x7a, 0x0b, 0xaf, 0xe0, 0x6d, 0x5c, 0xce, 0xd2, 0xa5, 0xb4, 0x17, 0x19, 0x28,
	0x9d, 0xdd, 0xcf, 0xff, 0xc1, 0x07, 0x1f, 0xe7, 0x8d, 0x01, 0x9b, 0x36, 0xde, 0x05, 0x27, 0x2e,
	0x6a, 0xed, 0x21, 0xed, 0x9a, 0x9d, 0x87, 0x2d, 0xae, 0x7e, 0x18, 0x8f, 0xdf, 0x0d, 0x58, 0xb1,
	0xe2, 0x97, 0xa4, 0xad, 0xaa, 0xc0, 0x04, 0x15, 0x34, 0x61, 0xc2, 0x96, 0x6c, 0xbd, 0xd8, 0x9c,
	0x93, 0xb6, 0xaf, 0x60, 0xc2, 0x87, 0x26, 0x14, 0xf7, 0xfc, 0xda, 0x3b, 0x63, 0x0a, 0x28, 0x6b,
	0x55, 0x56, 0x58, 0xd6, 0x6d, 0x47, 0xc9, 0xc9, 0x92, 0xad, 0xcf, 0x36, 0x57, 0x47, 0x90, 0xcf,
	0xbf, 0x78, 0xe4, 0xb7, 0x04, 0xbd, 0x42, 0xeb, 0x9d, 0x21, 0xb4, 0x41, 0x6d, 0x3b, 0x0f, 0x41,
	0x3b, 0x9b, 0x2c, 0x26, 0xf3, 0x0d, 0x41, 0xff, 0x72, 0x84, 0xcf, 0x33, 0x13, 0x82, 0xc7, 0x16,
	0x08, 0x93, 0x78, 0xb2, 0x4e, 0xfb, 0x29, 0xff, 0x1d, 0x24, 0xdb, 0x0f, 0x92, 0xfd, 0x0f, 0x92,
	0x7d, 0x8f, 0x32, 0xda, 0x8f, 0x32, 0xfa, 0x1b, 0x65, 0xf4, 0x79, 0xb7, 0xd3, 0xa1, 0xea, 0x8a,
	0xb4, 0x74, 0x94, 0xbd, 0x69, 0x0f, 0xb9, 0xf3, 0x98, 0xb5, 0x58, 0x83, 0xce, 0xfa, 0x6c, 0x4e,
	0xcc, 0xc2, 0x57, 0x83, 0x6d, 0x71, 0x3a, 0xd5, 0x3f, 0x1c, 0x02, 0x00, 0x00, 0xff, 0xff, 0x97,
	0xdb, 0xec, 0xa6, 0x0b, 0x01, 0x00, 0x00,
}

func (m *Plan) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Plan) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Plan) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintPlan(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x22
	}
	if m.MaxEnrolmentDuration != 0 {
		i = encodeVarintPlan(dAtA, i, uint64(m.MaxEnrolmentDuration))
		i--
		dAtA[i] = 0x18
	}
	if len(m.RollbackChecksum) > 0 {
		i -= len(m.RollbackChecksum)
		copy(dAtA[i:], m.RollbackChecksum)
		i = encodeVarintPlan(dAtA, i, uint64(len(m.RollbackChecksum)))
		i--
		dAtA[i] = 0x12
	}
	if m.MinHaltTime != 0 {
		i = encodeVarintPlan(dAtA, i, uint64(m.MinHaltTime))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPlan(dAtA []byte, offset int, v uint64) int {
	offset -= sovPlan(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Plan) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MinHaltTime != 0 {
		n += 1 + sovPlan(uint64(m.MinHaltTime))
	}
	l = len(m.RollbackChecksum)
	if l > 0 {
		n += 1 + l + sovPlan(uint64(l))
	}
	if m.MaxEnrolmentDuration != 0 {
		n += 1 + sovPlan(uint64(m.MaxEnrolmentDuration))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovPlan(uint64(l))
	}
	return n
}

func sovPlan(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPlan(x uint64) (n int) {
	return sovPlan(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Plan) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlan
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
			return fmt.Errorf("proto: Plan: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Plan: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinHaltTime", wireType)
			}
			m.MinHaltTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RollbackChecksum", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
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
				return ErrInvalidLengthPlan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPlan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RollbackChecksum = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxEnrolmentDuration", wireType)
			}
			m.MaxEnrolmentDuration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
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
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlan
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
				return ErrInvalidLengthPlan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPlan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPlan(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPlan
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPlan
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
func skipPlan(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPlan
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
					return 0, ErrIntOverflowPlan
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
					return 0, ErrIntOverflowPlan
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
				return 0, ErrInvalidLengthPlan
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPlan
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPlan
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPlan        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPlan          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPlan = fmt.Errorf("proto: unexpected end of group")
)
