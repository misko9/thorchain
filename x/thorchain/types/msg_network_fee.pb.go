// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: thorchain/v1/x/thorchain/types/msg_network_fee.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type MsgNetworkFee struct {
	BlockHeight        int64                                         `protobuf:"varint,1,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	Chain              gitlab_com_thorchain_thornode_common.Chain    `protobuf:"bytes,2,opt,name=chain,proto3,casttype=gitlab.com/thorchain/thornode/common.Chain" json:"chain,omitempty"`
	TransactionSize    uint64                                        `protobuf:"varint,3,opt,name=transaction_size,json=transactionSize,proto3" json:"transaction_size,omitempty"`
	TransactionFeeRate uint64                                        `protobuf:"varint,4,opt,name=transaction_fee_rate,json=transactionFeeRate,proto3" json:"transaction_fee_rate,omitempty"`
	Signer             github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,5,opt,name=signer,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"signer,omitempty"`
}

func (m *MsgNetworkFee) Reset()         { *m = MsgNetworkFee{} }
func (m *MsgNetworkFee) String() string { return proto.CompactTextString(m) }
func (*MsgNetworkFee) ProtoMessage()    {}
func (*MsgNetworkFee) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d204a455aa861be, []int{0}
}
func (m *MsgNetworkFee) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgNetworkFee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgNetworkFee.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgNetworkFee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgNetworkFee.Merge(m, src)
}
func (m *MsgNetworkFee) XXX_Size() int {
	return m.Size()
}
func (m *MsgNetworkFee) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgNetworkFee.DiscardUnknown(m)
}

var xxx_messageInfo_MsgNetworkFee proto.InternalMessageInfo

func (m *MsgNetworkFee) GetBlockHeight() int64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *MsgNetworkFee) GetChain() gitlab_com_thorchain_thornode_common.Chain {
	if m != nil {
		return m.Chain
	}
	return ""
}

func (m *MsgNetworkFee) GetTransactionSize() uint64 {
	if m != nil {
		return m.TransactionSize
	}
	return 0
}

func (m *MsgNetworkFee) GetTransactionFeeRate() uint64 {
	if m != nil {
		return m.TransactionFeeRate
	}
	return 0
}

func (m *MsgNetworkFee) GetSigner() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Signer
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgNetworkFee)(nil), "types.MsgNetworkFee")
}

func init() {
	proto.RegisterFile("thorchain/v1/x/thorchain/types/msg_network_fee.proto", fileDescriptor_9d204a455aa861be)
}

var fileDescriptor_9d204a455aa861be = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xbf, 0x4e, 0xeb, 0x30,
	0x14, 0xc6, 0xeb, 0xfe, 0x93, 0x6e, 0x6e, 0xaf, 0x2e, 0x8a, 0x3a, 0x44, 0x0c, 0x69, 0x60, 0x0a,
	0x48, 0x4d, 0xa8, 0xe0, 0x05, 0x5a, 0x50, 0x45, 0x07, 0x18, 0xc2, 0xc6, 0x12, 0xa5, 0xce, 0xa9,
	0x63, 0xb5, 0xf1, 0xa9, 0x6c, 0xf3, 0xaf, 0x4f, 0xc1, 0x23, 0xf0, 0x38, 0x8c, 0x1d, 0x99, 0x2a,
	0xd4, 0xbe, 0x45, 0x27, 0x54, 0xa7, 0x12, 0x91, 0x90, 0x98, 0x6c, 0xff, 0xbe, 0x73, 0xec, 0xcf,
	0xdf, 0xb1, 0x2e, 0x74, 0x86, 0x92, 0x66, 0x09, 0x17, 0xe1, 0x63, 0x2f, 0x7c, 0x0e, 0xbf, 0x8f,
	0xfa, 0x65, 0x0e, 0x2a, 0xcc, 0x15, 0x8b, 0x05, 0xe8, 0x27, 0x94, 0xd3, 0x78, 0x02, 0x10, 0xcc,
	0x25, 0x6a, 0xb4, 0x1b, 0x46, 0x3c, 0x6c, 0x33, 0x64, 0x68, 0x48, 0xb8, 0xdb, 0x15, 0xe2, 0xf1,
	0x5b, 0xd5, 0xfa, 0x77, 0xa3, 0xd8, 0x6d, 0xd1, 0x35, 0x04, 0xb0, 0x8f, 0xac, 0xd6, 0x78, 0x86,
	0x74, 0x1a, 0x67, 0xc0, 0x59, 0xa6, 0x1d, 0xe2, 0x11, 0xbf, 0x16, 0xfd, 0x35, 0xec, 0xda, 0x20,
	0xfb, 0xca, 0x6a, 0x98, 0x67, 0x9d, 0xaa, 0x47, 0xfc, 0x3f, 0x83, 0x60, 0xbb, 0xea, 0x9c, 0x32,
	0xae, 0x67, 0xc9, 0x38, 0xa0, 0x98, 0x97, 0x6d, 0x65, 0x28, 0x05, 0xa6, 0x10, 0x52, 0xcc, 0x73,
	0x14, 0xc1, 0xe5, 0x8e, 0x46, 0x45, 0xb3, 0x7d, 0x62, 0x1d, 0x68, 0x99, 0x08, 0x95, 0x50, 0xcd,
	0x51, 0xc4, 0x8a, 0x2f, 0xc0, 0xa9, 0x79, 0xc4, 0xaf, 0x47, 0xff, 0x4b, 0xfc, 0x8e, 0x2f, 0xc0,
	0x3e, 0xb3, 0xda, 0xe5, 0xd2, 0x09, 0x40, 0x2c, 0x13, 0x0d, 0x4e, 0xdd, 0x94, 0xdb, 0x25, 0x6d,
	0x08, 0x10, 0x25, 0x1a, 0xec, 0x91, 0xd5, 0x54, 0x9c, 0x09, 0x90, 0x4e, 0xc3, 0x23, 0x7e, 0x6b,
	0xd0, 0xdb, 0xae, 0x3a, 0x5d, 0xc6, 0x75, 0xf6, 0x50, 0x78, 0xa4, 0xa8, 0x72, 0x54, 0xfb, 0xa5,
	0xab, 0xd2, 0x69, 0x11, 0x61, 0xd0, 0xa7, 0xb4, 0x9f, 0xa6, 0x12, 0x94, 0x8a, 0xf6, 0x17, 0x0c,
	0x46, 0xef, 0x6b, 0x97, 0x2c, 0xd7, 0x2e, 0xf9, 0x5c, 0xbb, 0xe4, 0x75, 0xe3, 0x56, 0x96, 0x1b,
	0xb7, 0xf2, 0xb1, 0x71, 0x2b, 0xf7, 0xe1, 0xef, 0x9f, 0xfe, 0x31, 0xa0, 0x71, 0xd3, 0x84, 0x7e,
	0xfe, 0x15, 0x00, 0x00, 0xff, 0xff, 0x3c, 0xdd, 0xc4, 0xa3, 0xc9, 0x01, 0x00, 0x00,
}

func (m *MsgNetworkFee) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgNetworkFee) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgNetworkFee) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintMsgNetworkFee(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0x2a
	}
	if m.TransactionFeeRate != 0 {
		i = encodeVarintMsgNetworkFee(dAtA, i, uint64(m.TransactionFeeRate))
		i--
		dAtA[i] = 0x20
	}
	if m.TransactionSize != 0 {
		i = encodeVarintMsgNetworkFee(dAtA, i, uint64(m.TransactionSize))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintMsgNetworkFee(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0x12
	}
	if m.BlockHeight != 0 {
		i = encodeVarintMsgNetworkFee(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintMsgNetworkFee(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsgNetworkFee(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgNetworkFee) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BlockHeight != 0 {
		n += 1 + sovMsgNetworkFee(uint64(m.BlockHeight))
	}
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovMsgNetworkFee(uint64(l))
	}
	if m.TransactionSize != 0 {
		n += 1 + sovMsgNetworkFee(uint64(m.TransactionSize))
	}
	if m.TransactionFeeRate != 0 {
		n += 1 + sovMsgNetworkFee(uint64(m.TransactionFeeRate))
	}
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovMsgNetworkFee(uint64(l))
	}
	return n
}

func sovMsgNetworkFee(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsgNetworkFee(x uint64) (n int) {
	return sovMsgNetworkFee(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgNetworkFee) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgNetworkFee
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
			return fmt.Errorf("proto: MsgNetworkFee: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgNetworkFee: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgNetworkFee
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgNetworkFee
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
				return ErrInvalidLengthMsgNetworkFee
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgNetworkFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chain = gitlab_com_thorchain_thornode_common.Chain(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TransactionSize", wireType)
			}
			m.TransactionSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgNetworkFee
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TransactionSize |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TransactionFeeRate", wireType)
			}
			m.TransactionFeeRate = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgNetworkFee
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TransactionFeeRate |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgNetworkFee
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
				return ErrInvalidLengthMsgNetworkFee
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgNetworkFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signer = append(m.Signer[:0], dAtA[iNdEx:postIndex]...)
			if m.Signer == nil {
				m.Signer = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsgNetworkFee(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgNetworkFee
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
func skipMsgNetworkFee(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsgNetworkFee
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
					return 0, ErrIntOverflowMsgNetworkFee
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
					return 0, ErrIntOverflowMsgNetworkFee
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
				return 0, ErrInvalidLengthMsgNetworkFee
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsgNetworkFee
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsgNetworkFee
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsgNetworkFee        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsgNetworkFee          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsgNetworkFee = fmt.Errorf("proto: unexpected end of group")
)
