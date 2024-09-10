// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: thorchain/v1/x/thorchain/types/msg_tx_outbound.proto

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

type MsgOutboundTx struct {
	Tx     ObservedTx                                    `protobuf:"bytes,1,opt,name=tx,proto3" json:"tx"`
	InTxID gitlab_com_thorchain_thornode_common.TxID     `protobuf:"bytes,2,opt,name=in_tx_id,json=inTxId,proto3,casttype=gitlab.com/thorchain/thornode/common.TxID" json:"in_tx_id,omitempty"`
	Signer github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,3,opt,name=signer,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"signer,omitempty"`
}

func (m *MsgOutboundTx) Reset()         { *m = MsgOutboundTx{} }
func (m *MsgOutboundTx) String() string { return proto.CompactTextString(m) }
func (*MsgOutboundTx) ProtoMessage()    {}
func (*MsgOutboundTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7cbbc7de1adc915, []int{0}
}
func (m *MsgOutboundTx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgOutboundTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgOutboundTx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgOutboundTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgOutboundTx.Merge(m, src)
}
func (m *MsgOutboundTx) XXX_Size() int {
	return m.Size()
}
func (m *MsgOutboundTx) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgOutboundTx.DiscardUnknown(m)
}

var xxx_messageInfo_MsgOutboundTx proto.InternalMessageInfo

func (m *MsgOutboundTx) GetTx() ObservedTx {
	if m != nil {
		return m.Tx
	}
	return ObservedTx{}
}

func (m *MsgOutboundTx) GetInTxID() gitlab_com_thorchain_thornode_common.TxID {
	if m != nil {
		return m.InTxID
	}
	return ""
}

func (m *MsgOutboundTx) GetSigner() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Signer
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgOutboundTx)(nil), "types.MsgOutboundTx")
}

func init() {
	proto.RegisterFile("thorchain/v1/x/thorchain/types/msg_tx_outbound.proto", fileDescriptor_a7cbbc7de1adc915)
}

var fileDescriptor_a7cbbc7de1adc915 = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xbb, 0x4e, 0xf3, 0x30,
	0x1c, 0xc5, 0xe3, 0x7e, 0x1f, 0x11, 0x04, 0x18, 0x88, 0x18, 0xaa, 0x0e, 0x49, 0xc5, 0x42, 0x19,
	0x1a, 0xab, 0x5c, 0xc4, 0xdc, 0x88, 0x25, 0x03, 0xaa, 0xa8, 0x3a, 0xb1, 0x44, 0x4d, 0x6c, 0x39,
	0x16, 0xc4, 0xff, 0xca, 0x76, 0x2b, 0xf3, 0x16, 0x3c, 0x56, 0xc7, 0x4e, 0x88, 0x29, 0x42, 0xe9,
	0x5b, 0x74, 0x42, 0xb9, 0x48, 0x20, 0x21, 0xc1, 0x62, 0x1f, 0x5f, 0xce, 0xcf, 0xc7, 0xc7, 0xb9,
	0xd6, 0x19, 0xc8, 0x34, 0x9b, 0x73, 0x81, 0x57, 0x23, 0x6c, 0xf0, 0xd7, 0x52, 0xbf, 0x2c, 0xa8,
	0xc2, 0xb9, 0x62, 0xb1, 0x36, 0x31, 0x2c, 0x75, 0x02, 0x4b, 0x41, 0x82, 0x85, 0x04, 0x0d, 0xee,
	0x5e, 0x7d, 0xd8, 0xbb, 0xf9, 0xc3, 0x5c, 0x8d, 0x31, 0x24, 0x8a, 0xca, 0x15, 0x25, 0xb1, 0x36,
	0x8d, 0xbb, 0x77, 0xca, 0x80, 0x41, 0x2d, 0x71, 0xa5, 0x9a, 0xdd, 0xb3, 0x37, 0xe4, 0x1c, 0xdf,
	0x2b, 0x36, 0x69, 0x5f, 0x9a, 0x19, 0xf7, 0xdc, 0xe9, 0x68, 0xd3, 0x45, 0x7d, 0x34, 0x38, 0xbc,
	0x3c, 0x09, 0x6a, 0x64, 0x30, 0x69, 0x69, 0x33, 0x13, 0xfe, 0x5f, 0x17, 0xbe, 0x35, 0xed, 0x68,
	0xe3, 0x3e, 0x38, 0xfb, 0x5c, 0x54, 0x31, 0x39, 0xe9, 0x76, 0xfa, 0x68, 0x70, 0x10, 0xde, 0x96,
	0x85, 0x6f, 0x47, 0x62, 0x66, 0xa2, 0xbb, 0x5d, 0xe1, 0x5f, 0x30, 0xae, 0x9f, 0xe7, 0x49, 0x90,
	0x42, 0xfe, 0x3d, 0x63, 0x06, 0x52, 0x00, 0xa1, 0x38, 0x85, 0x3c, 0x07, 0x11, 0x54, 0x97, 0xa7,
	0x36, 0xaf, 0x4c, 0xc4, 0x8d, 0x1c, 0x5b, 0x71, 0x26, 0xa8, 0xec, 0xfe, 0xeb, 0xa3, 0xc1, 0x51,
	0x38, 0xda, 0x15, 0xfe, 0x90, 0x71, 0x9d, 0x2d, 0x1b, 0x4c, 0x0a, 0x2a, 0x07, 0xd5, 0x4e, 0x43,
	0x45, 0x9e, 0x9a, 0x2f, 0x07, 0xe3, 0x34, 0x1d, 0x13, 0x22, 0xa9, 0x52, 0xd3, 0x16, 0x10, 0x46,
	0xeb, 0xd2, 0x43, 0x9b, 0xd2, 0x43, 0x1f, 0xa5, 0x87, 0x5e, 0xb7, 0x9e, 0xb5, 0xd9, 0x7a, 0xd6,
	0xfb, 0xd6, 0xb3, 0x1e, 0xf1, 0xef, 0xb9, 0x7e, 0x14, 0x9a, 0xd8, 0x75, 0x55, 0x57, 0x9f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xca, 0x04, 0xf7, 0x2b, 0xb6, 0x01, 0x00, 0x00,
}

func (m *MsgOutboundTx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgOutboundTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgOutboundTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintMsgTxOutbound(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.InTxID) > 0 {
		i -= len(m.InTxID)
		copy(dAtA[i:], m.InTxID)
		i = encodeVarintMsgTxOutbound(dAtA, i, uint64(len(m.InTxID)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Tx.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMsgTxOutbound(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintMsgTxOutbound(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsgTxOutbound(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgOutboundTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Tx.Size()
	n += 1 + l + sovMsgTxOutbound(uint64(l))
	l = len(m.InTxID)
	if l > 0 {
		n += 1 + l + sovMsgTxOutbound(uint64(l))
	}
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovMsgTxOutbound(uint64(l))
	}
	return n
}

func sovMsgTxOutbound(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsgTxOutbound(x uint64) (n int) {
	return sovMsgTxOutbound(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgOutboundTx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgTxOutbound
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
			return fmt.Errorf("proto: MsgOutboundTx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgOutboundTx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tx", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgTxOutbound
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
				return ErrInvalidLengthMsgTxOutbound
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsgTxOutbound
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Tx.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InTxID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgTxOutbound
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
				return ErrInvalidLengthMsgTxOutbound
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgTxOutbound
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InTxID = gitlab_com_thorchain_thornode_common.TxID(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgTxOutbound
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
				return ErrInvalidLengthMsgTxOutbound
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgTxOutbound
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
			skippy, err := skipMsgTxOutbound(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgTxOutbound
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
func skipMsgTxOutbound(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsgTxOutbound
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
					return 0, ErrIntOverflowMsgTxOutbound
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
					return 0, ErrIntOverflowMsgTxOutbound
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
				return 0, ErrInvalidLengthMsgTxOutbound
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsgTxOutbound
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsgTxOutbound
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsgTxOutbound        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsgTxOutbound          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsgTxOutbound = fmt.Errorf("proto: unexpected end of group")
)
