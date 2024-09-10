// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: thorchain/v1/x/thorchain/types/type_tss.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	gitlab_com_thorchain_thornode_common "gitlab.com/thorchain/thornode/common"
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

type TssVoter struct {
	ID                           string                                      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PoolPubKey                   gitlab_com_thorchain_thornode_common.PubKey `protobuf:"bytes,2,opt,name=pool_pub_key,json=poolPubKey,proto3,casttype=gitlab.com/thorchain/thornode/common.PubKey" json:"pool_pub_key,omitempty"`
	PubKeys                      []string                                    `protobuf:"bytes,3,rep,name=pub_keys,json=pubKeys,proto3" json:"pub_keys,omitempty"`
	BlockHeight                  int64                                       `protobuf:"varint,4,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	Chains                       []string                                    `protobuf:"bytes,5,rep,name=chains,proto3" json:"chains,omitempty"`
	Signers                      []string                                    `protobuf:"bytes,6,rep,name=signers,proto3" json:"signers,omitempty"`
	MajorityConsensusBlockHeight int64                                       `protobuf:"varint,7,opt,name=majority_consensus_block_height,json=majorityConsensusBlockHeight,proto3" json:"majority_consensus_block_height,omitempty"`
}

func (m *TssVoter) Reset()      { *m = TssVoter{} }
func (*TssVoter) ProtoMessage() {}
func (*TssVoter) Descriptor() ([]byte, []int) {
	return fileDescriptor_54da3cd5f1025677, []int{0}
}
func (m *TssVoter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TssVoter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TssVoter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TssVoter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TssVoter.Merge(m, src)
}
func (m *TssVoter) XXX_Size() int {
	return m.Size()
}
func (m *TssVoter) XXX_DiscardUnknown() {
	xxx_messageInfo_TssVoter.DiscardUnknown(m)
}

var xxx_messageInfo_TssVoter proto.InternalMessageInfo

func init() {
	proto.RegisterType((*TssVoter)(nil), "types.TssVoter")
}

func init() {
	proto.RegisterFile("thorchain/v1/x/thorchain/types/type_tss.proto", fileDescriptor_54da3cd5f1025677)
}

var fileDescriptor_54da3cd5f1025677 = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xc1, 0x4e, 0xf2, 0x40,
	0x18, 0x6c, 0xcb, 0x4f, 0x81, 0xfd, 0x39, 0x6d, 0x0c, 0xa9, 0xc6, 0x2c, 0xe8, 0x89, 0xc4, 0xc8,
	0xc6, 0xf8, 0x06, 0xa8, 0x89, 0xc6, 0x8b, 0x36, 0xea, 0xc1, 0x4b, 0x43, 0xcb, 0xa6, 0x5d, 0xa1,
	0xfd, 0x9a, 0x7e, 0x8b, 0xb1, 0x37, 0x1f, 0xc1, 0x07, 0xf1, 0x41, 0x38, 0x72, 0xe4, 0x44, 0xa4,
	0xbc, 0x85, 0x27, 0xc3, 0x02, 0x41, 0x63, 0xe2, 0x65, 0xf3, 0xcd, 0x7c, 0xb3, 0x33, 0x9b, 0x1d,
	0x72, 0xac, 0x22, 0xc8, 0x82, 0xa8, 0x27, 0x13, 0xfe, 0x7c, 0xc2, 0x5f, 0xf8, 0x16, 0xaa, 0x3c,
	0x15, 0xa8, 0x4f, 0x4f, 0x21, 0x76, 0xd2, 0x0c, 0x14, 0xd0, 0xb2, 0x66, 0xf7, 0x76, 0x42, 0x08,
	0x41, 0x33, 0x7c, 0x39, 0xad, 0x96, 0x87, 0xef, 0x16, 0xa9, 0xde, 0x21, 0x3e, 0x80, 0x12, 0x19,
	0x6d, 0x10, 0x4b, 0xf6, 0x1d, 0xb3, 0x65, 0xb6, 0x6b, 0x5d, 0xbb, 0x98, 0x35, 0xad, 0xab, 0x73,
	0xd7, 0x92, 0x7d, 0x7a, 0x4b, 0xea, 0x29, 0xc0, 0xd0, 0x4b, 0x47, 0xbe, 0x37, 0x10, 0xb9, 0x63,
	0x69, 0x05, 0xff, 0x9c, 0x35, 0x8f, 0x42, 0xa9, 0x86, 0x3d, 0xbf, 0x13, 0x40, 0xfc, 0xfd, 0x19,
	0x11, 0x64, 0x09, 0xf4, 0x05, 0x0f, 0x20, 0x8e, 0x21, 0xe9, 0xdc, 0x8c, 0xfc, 0x6b, 0x91, 0xbb,
	0x64, 0x69, 0xb2, 0x9a, 0xe9, 0x2e, 0xa9, 0xae, 0xdd, 0xd0, 0x29, 0xb5, 0x4a, 0xed, 0x9a, 0x5b,
	0x49, 0xf5, 0x06, 0xe9, 0x01, 0xa9, 0xfb, 0x43, 0x08, 0x06, 0x5e, 0x24, 0x64, 0x18, 0x29, 0xe7,
	0x5f, 0xcb, 0x6c, 0x97, 0xdc, 0xff, 0x9a, 0xbb, 0xd4, 0x14, 0x6d, 0x10, 0x5b, 0x27, 0xa1, 0x53,
	0xd6, 0x77, 0xd7, 0x88, 0x3a, 0xa4, 0x82, 0x32, 0x4c, 0x44, 0x86, 0x8e, 0xbd, 0x32, 0x5d, 0x43,
	0x7a, 0x41, 0x9a, 0x71, 0xef, 0x09, 0x32, 0xa9, 0x72, 0x2f, 0x80, 0x04, 0x45, 0x82, 0x23, 0xf4,
	0x7e, 0xe4, 0x54, 0x74, 0xce, 0xfe, 0x46, 0x76, 0xb6, 0x51, 0x75, 0xb7, 0xc1, 0xdd, 0xfb, 0xf1,
	0x9c, 0x19, 0xd3, 0x39, 0x33, 0x5e, 0x0b, 0x66, 0x8c, 0x0b, 0x66, 0x4e, 0x0a, 0x66, 0x7e, 0x14,
	0xcc, 0x7c, 0x5b, 0x30, 0x63, 0xb2, 0x60, 0xc6, 0x74, 0xc1, 0x8c, 0x47, 0xfe, 0xf7, 0xaf, 0xfc,
	0x6a, 0xcc, 0xb7, 0x75, 0x19, 0xa7, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x40, 0x9a, 0xe4, 0x9b,
	0xda, 0x01, 0x00, 0x00,
}

func (m *TssVoter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TssVoter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TssVoter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MajorityConsensusBlockHeight != 0 {
		i = encodeVarintTypeTss(dAtA, i, uint64(m.MajorityConsensusBlockHeight))
		i--
		dAtA[i] = 0x38
	}
	if len(m.Signers) > 0 {
		for iNdEx := len(m.Signers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Signers[iNdEx])
			copy(dAtA[i:], m.Signers[iNdEx])
			i = encodeVarintTypeTss(dAtA, i, uint64(len(m.Signers[iNdEx])))
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.Chains) > 0 {
		for iNdEx := len(m.Chains) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Chains[iNdEx])
			copy(dAtA[i:], m.Chains[iNdEx])
			i = encodeVarintTypeTss(dAtA, i, uint64(len(m.Chains[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.BlockHeight != 0 {
		i = encodeVarintTypeTss(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x20
	}
	if len(m.PubKeys) > 0 {
		for iNdEx := len(m.PubKeys) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.PubKeys[iNdEx])
			copy(dAtA[i:], m.PubKeys[iNdEx])
			i = encodeVarintTypeTss(dAtA, i, uint64(len(m.PubKeys[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.PoolPubKey) > 0 {
		i -= len(m.PoolPubKey)
		copy(dAtA[i:], m.PoolPubKey)
		i = encodeVarintTypeTss(dAtA, i, uint64(len(m.PoolPubKey)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ID) > 0 {
		i -= len(m.ID)
		copy(dAtA[i:], m.ID)
		i = encodeVarintTypeTss(dAtA, i, uint64(len(m.ID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypeTss(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypeTss(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TssVoter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovTypeTss(uint64(l))
	}
	l = len(m.PoolPubKey)
	if l > 0 {
		n += 1 + l + sovTypeTss(uint64(l))
	}
	if len(m.PubKeys) > 0 {
		for _, s := range m.PubKeys {
			l = len(s)
			n += 1 + l + sovTypeTss(uint64(l))
		}
	}
	if m.BlockHeight != 0 {
		n += 1 + sovTypeTss(uint64(m.BlockHeight))
	}
	if len(m.Chains) > 0 {
		for _, s := range m.Chains {
			l = len(s)
			n += 1 + l + sovTypeTss(uint64(l))
		}
	}
	if len(m.Signers) > 0 {
		for _, s := range m.Signers {
			l = len(s)
			n += 1 + l + sovTypeTss(uint64(l))
		}
	}
	if m.MajorityConsensusBlockHeight != 0 {
		n += 1 + sovTypeTss(uint64(m.MajorityConsensusBlockHeight))
	}
	return n
}

func sovTypeTss(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypeTss(x uint64) (n int) {
	return sovTypeTss(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TssVoter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypeTss
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
			return fmt.Errorf("proto: TssVoter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TssVoter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeTss
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
				return ErrInvalidLengthTypeTss
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypeTss
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolPubKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeTss
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
				return ErrInvalidLengthTypeTss
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypeTss
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PoolPubKey = gitlab_com_thorchain_thornode_common.PubKey(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKeys", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeTss
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
				return ErrInvalidLengthTypeTss
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypeTss
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PubKeys = append(m.PubKeys, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeTss
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chains", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeTss
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
				return ErrInvalidLengthTypeTss
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypeTss
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chains = append(m.Chains, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeTss
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
				return ErrInvalidLengthTypeTss
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypeTss
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signers = append(m.Signers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MajorityConsensusBlockHeight", wireType)
			}
			m.MajorityConsensusBlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeTss
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MajorityConsensusBlockHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypeTss(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypeTss
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
func skipTypeTss(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypeTss
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
					return 0, ErrIntOverflowTypeTss
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
					return 0, ErrIntOverflowTypeTss
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
				return 0, ErrInvalidLengthTypeTss
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypeTss
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypeTss
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypeTss        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypeTss          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypeTss = fmt.Errorf("proto: unexpected end of group")
)
