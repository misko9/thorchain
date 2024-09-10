// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: types/type_observed_network_fee.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

type ObservedNetworkFeeVoter struct {
	BlockHeight       int64                                      `protobuf:"varint,1,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	ReportBlockHeight int64                                      `protobuf:"varint,2,opt,name=report_block_height,json=reportBlockHeight,proto3" json:"report_block_height,omitempty"`
	Chain             gitlab_com_thorchain_thornode_common.Chain `protobuf:"bytes,3,opt,name=chain,proto3,casttype=gitlab.com/thorchain/thornode/common.Chain" json:"chain,omitempty"`
	Signers           []string                                   `protobuf:"bytes,4,rep,name=signers,proto3" json:"signers,omitempty"`
	FeeRate           int64                                      `protobuf:"varint,5,opt,name=fee_rate,json=feeRate,proto3" json:"fee_rate,omitempty"`
	TransactionSize   int64                                      `protobuf:"varint,6,opt,name=transaction_size,json=transactionSize,proto3" json:"transaction_size,omitempty"`
}

func (m *ObservedNetworkFeeVoter) Reset()      { *m = ObservedNetworkFeeVoter{} }
func (*ObservedNetworkFeeVoter) ProtoMessage() {}
func (*ObservedNetworkFeeVoter) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a0836138bb566bd, []int{0}
}
func (m *ObservedNetworkFeeVoter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ObservedNetworkFeeVoter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ObservedNetworkFeeVoter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ObservedNetworkFeeVoter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObservedNetworkFeeVoter.Merge(m, src)
}
func (m *ObservedNetworkFeeVoter) XXX_Size() int {
	return m.Size()
}
func (m *ObservedNetworkFeeVoter) XXX_DiscardUnknown() {
	xxx_messageInfo_ObservedNetworkFeeVoter.DiscardUnknown(m)
}

var xxx_messageInfo_ObservedNetworkFeeVoter proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ObservedNetworkFeeVoter)(nil), "types.ObservedNetworkFeeVoter")
}

func init() {
	proto.RegisterFile("types/type_observed_network_fee.proto", fileDescriptor_6a0836138bb566bd)
}

var fileDescriptor_6a0836138bb566bd = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x93, 0xd6, 0xb6, 0x76, 0x15, 0xd4, 0x28, 0x18, 0x3d, 0xac, 0x55, 0x10, 0xaa, 0x87,
	0xe4, 0xe0, 0x1b, 0x54, 0x11, 0x4f, 0x0a, 0x11, 0x3d, 0x78, 0x59, 0x92, 0x74, 0x9a, 0x2c, 0x6d,
	0x77, 0xca, 0x66, 0xf0, 0x4f, 0x4f, 0x3e, 0x82, 0xf8, 0x54, 0x3d, 0xf6, 0xd8, 0x93, 0xd8, 0xf4,
	0x2d, 0x3c, 0x49, 0x36, 0x0a, 0xf5, 0xe2, 0x65, 0x99, 0xfd, 0x7d, 0xbf, 0x0f, 0x86, 0x5d, 0x76,
	0x4c, 0x2f, 0x23, 0xc8, 0xfc, 0xe2, 0x14, 0x18, 0x65, 0xa0, 0x1f, 0xa1, 0x2b, 0x14, 0xd0, 0x13,
	0xea, 0xbe, 0xe8, 0x01, 0x78, 0x23, 0x8d, 0x84, 0x4e, 0xcd, 0x68, 0xfb, 0x3b, 0x09, 0x26, 0x68,
	0x88, 0x5f, 0x4c, 0x65, 0x78, 0xf4, 0x5e, 0x61, 0xbb, 0x37, 0x3f, 0xdd, 0xeb, 0xb2, 0x7a, 0x09,
	0x70, 0x8f, 0x04, 0xda, 0x39, 0x64, 0xeb, 0xd1, 0x00, 0xe3, 0xbe, 0x48, 0x41, 0x26, 0x29, 0xb9,
	0x76, 0xcb, 0x6e, 0x57, 0x83, 0x35, 0xc3, 0xae, 0x0c, 0x72, 0x3c, 0xb6, 0xad, 0x61, 0x84, 0x9a,
	0xc4, 0x1f, 0xb3, 0x62, 0xcc, 0xad, 0x32, 0xea, 0x2c, 0xf9, 0x17, 0xac, 0x16, 0xa7, 0xa1, 0x54,
	0x6e, 0xb5, 0x65, 0xb7, 0x9b, 0x1d, 0xef, 0xeb, 0xe3, 0xe0, 0x34, 0x91, 0x34, 0x08, 0x23, 0x2f,
	0xc6, 0xa1, 0x4f, 0x29, 0x6a, 0x93, 0x9b, 0x49, 0x61, 0x17, 0xfc, 0x18, 0x87, 0x43, 0x54, 0xde,
	0x79, 0x41, 0x83, 0xb2, 0xec, 0xb8, 0xac, 0x91, 0xc9, 0x44, 0x81, 0xce, 0xdc, 0x95, 0x56, 0xb5,
	0xdd, 0x0c, 0x7e, 0xaf, 0xce, 0x1e, 0x5b, 0xed, 0x01, 0x08, 0x1d, 0x12, 0xb8, 0x35, 0xb3, 0x44,
	0xa3, 0x07, 0x10, 0x84, 0x04, 0xce, 0x09, 0xdb, 0x24, 0x1d, 0xaa, 0x2c, 0x8c, 0x49, 0xa2, 0x12,
	0x99, 0x1c, 0x83, 0x5b, 0x37, 0xca, 0xc6, 0x12, 0xbf, 0x95, 0x63, 0xe8, 0xdc, 0x4d, 0xe6, 0xdc,
	0x9a, 0xcd, 0xb9, 0xf5, 0x9a, 0x73, 0x6b, 0x92, 0x73, 0x7b, 0x9a, 0x73, 0xfb, 0x33, 0xe7, 0xf6,
	0xdb, 0x82, 0x5b, 0xd3, 0x05, 0xb7, 0x66, 0x0b, 0x6e, 0x3d, 0xf8, 0xff, 0x2f, 0xfe, 0xbc, 0x0c,
	0x8b, 0x1f, 0x88, 0xea, 0xe6, 0xc9, 0xcf, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0xd6, 0x12, 0xd6,
	0x5b, 0xb8, 0x01, 0x00, 0x00,
}

func (m *ObservedNetworkFeeVoter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ObservedNetworkFeeVoter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ObservedNetworkFeeVoter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TransactionSize != 0 {
		i = encodeVarintTypeObservedNetworkFee(dAtA, i, uint64(m.TransactionSize))
		i--
		dAtA[i] = 0x30
	}
	if m.FeeRate != 0 {
		i = encodeVarintTypeObservedNetworkFee(dAtA, i, uint64(m.FeeRate))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Signers) > 0 {
		for iNdEx := len(m.Signers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Signers[iNdEx])
			copy(dAtA[i:], m.Signers[iNdEx])
			i = encodeVarintTypeObservedNetworkFee(dAtA, i, uint64(len(m.Signers[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintTypeObservedNetworkFee(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0x1a
	}
	if m.ReportBlockHeight != 0 {
		i = encodeVarintTypeObservedNetworkFee(dAtA, i, uint64(m.ReportBlockHeight))
		i--
		dAtA[i] = 0x10
	}
	if m.BlockHeight != 0 {
		i = encodeVarintTypeObservedNetworkFee(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypeObservedNetworkFee(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypeObservedNetworkFee(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ObservedNetworkFeeVoter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BlockHeight != 0 {
		n += 1 + sovTypeObservedNetworkFee(uint64(m.BlockHeight))
	}
	if m.ReportBlockHeight != 0 {
		n += 1 + sovTypeObservedNetworkFee(uint64(m.ReportBlockHeight))
	}
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovTypeObservedNetworkFee(uint64(l))
	}
	if len(m.Signers) > 0 {
		for _, s := range m.Signers {
			l = len(s)
			n += 1 + l + sovTypeObservedNetworkFee(uint64(l))
		}
	}
	if m.FeeRate != 0 {
		n += 1 + sovTypeObservedNetworkFee(uint64(m.FeeRate))
	}
	if m.TransactionSize != 0 {
		n += 1 + sovTypeObservedNetworkFee(uint64(m.TransactionSize))
	}
	return n
}

func sovTypeObservedNetworkFee(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypeObservedNetworkFee(x uint64) (n int) {
	return sovTypeObservedNetworkFee(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ObservedNetworkFeeVoter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypeObservedNetworkFee
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
			return fmt.Errorf("proto: ObservedNetworkFeeVoter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ObservedNetworkFeeVoter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeObservedNetworkFee
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
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReportBlockHeight", wireType)
			}
			m.ReportBlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeObservedNetworkFee
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReportBlockHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeObservedNetworkFee
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
				return ErrInvalidLengthTypeObservedNetworkFee
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypeObservedNetworkFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chain = gitlab_com_thorchain_thornode_common.Chain(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeObservedNetworkFee
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
				return ErrInvalidLengthTypeObservedNetworkFee
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypeObservedNetworkFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signers = append(m.Signers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeRate", wireType)
			}
			m.FeeRate = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeObservedNetworkFee
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FeeRate |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TransactionSize", wireType)
			}
			m.TransactionSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeObservedNetworkFee
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TransactionSize |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypeObservedNetworkFee(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypeObservedNetworkFee
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
func skipTypeObservedNetworkFee(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypeObservedNetworkFee
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
					return 0, ErrIntOverflowTypeObservedNetworkFee
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
					return 0, ErrIntOverflowTypeObservedNetworkFee
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
				return 0, ErrInvalidLengthTypeObservedNetworkFee
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypeObservedNetworkFee
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypeObservedNetworkFee
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypeObservedNetworkFee        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypeObservedNetworkFee          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypeObservedNetworkFee = fmt.Errorf("proto: unexpected end of group")
)
