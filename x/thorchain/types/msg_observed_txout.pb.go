// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: types/msg_observed_txout.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

type MsgObservedTxOut struct {
	Txs    ObservedTxs                                   `protobuf:"bytes,1,rep,name=txs,proto3,castrepeated=ObservedTxs" json:"txs"`
	Signer github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=signer,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"signer,omitempty"`
}

func (m *MsgObservedTxOut) Reset()         { *m = MsgObservedTxOut{} }
func (m *MsgObservedTxOut) String() string { return proto.CompactTextString(m) }
func (*MsgObservedTxOut) ProtoMessage()    {}
func (*MsgObservedTxOut) Descriptor() ([]byte, []int) {
	return fileDescriptor_e42fe80ef60f50ae, []int{0}
}
func (m *MsgObservedTxOut) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgObservedTxOut) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgObservedTxOut.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgObservedTxOut) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgObservedTxOut.Merge(m, src)
}
func (m *MsgObservedTxOut) XXX_Size() int {
	return m.Size()
}
func (m *MsgObservedTxOut) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgObservedTxOut.DiscardUnknown(m)
}

var xxx_messageInfo_MsgObservedTxOut proto.InternalMessageInfo

func (m *MsgObservedTxOut) GetTxs() ObservedTxs {
	if m != nil {
		return m.Txs
	}
	return nil
}

func (m *MsgObservedTxOut) GetSigner() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Signer
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgObservedTxOut)(nil), "types.MsgObservedTxOut")
}

func init() { proto.RegisterFile("types/msg_observed_txout.proto", fileDescriptor_e42fe80ef60f50ae) }

var fileDescriptor_e42fe80ef60f50ae = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2b, 0xa9, 0x2c, 0x48,
	0x2d, 0xd6, 0xcf, 0x2d, 0x4e, 0x8f, 0xcf, 0x4f, 0x2a, 0x4e, 0x2d, 0x2a, 0x4b, 0x4d, 0x89, 0x2f,
	0xa9, 0xc8, 0x2f, 0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0xcb, 0x4b, 0xc9,
	0x40, 0x94, 0x81, 0x48, 0x64, 0x75, 0x10, 0x45, 0x52, 0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0xa6,
	0x3e, 0x88, 0x05, 0x11, 0x55, 0x9a, 0xce, 0xc8, 0x25, 0xe0, 0x5b, 0x9c, 0xee, 0x0f, 0x55, 0x1e,
	0x52, 0xe1, 0x5f, 0x5a, 0x22, 0x64, 0xc1, 0xc5, 0x5c, 0x52, 0x51, 0x2c, 0xc1, 0xa8, 0xc0, 0xac,
	0xc1, 0x6d, 0x24, 0xa8, 0x07, 0x36, 0x56, 0x0f, 0xa1, 0xc4, 0x49, 0xf8, 0xc4, 0x3d, 0x79, 0x86,
	0x55, 0xf7, 0xe5, 0xb9, 0x11, 0x62, 0xc5, 0x41, 0x20, 0x2d, 0x42, 0x9e, 0x5c, 0x6c, 0xc5, 0x99,
	0xe9, 0x79, 0xa9, 0x45, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x3c, 0x4e, 0x86, 0xbf, 0xee, 0xc9, 0xeb,
	0xa6, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x27, 0xe7, 0x17, 0xe7, 0xe6,
	0x17, 0x43, 0x29, 0xdd, 0xe2, 0x94, 0x6c, 0x7d, 0x88, 0xe1, 0x8e, 0xc9, 0xc9, 0x8e, 0x29, 0x29,
	0x45, 0xa9, 0xc5, 0xc5, 0x41, 0x50, 0x03, 0x9c, 0x3c, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48,
	0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1,
	0x58, 0x8e, 0x21, 0x4a, 0x3f, 0x3d, 0xb3, 0x24, 0x27, 0x11, 0x62, 0x60, 0x49, 0x46, 0x7e, 0x51,
	0x72, 0x46, 0x62, 0x66, 0x1e, 0x98, 0x95, 0x97, 0x9f, 0x92, 0xaa, 0x5f, 0x81, 0x2c, 0x08, 0x32,
	0x3d, 0x89, 0x0d, 0xec, 0x57, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf9, 0x53, 0xe6, 0xec,
	0x48, 0x01, 0x00, 0x00,
}

func (m *MsgObservedTxOut) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgObservedTxOut) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgObservedTxOut) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintMsgObservedTxout(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Txs) > 0 {
		for iNdEx := len(m.Txs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Txs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMsgObservedTxout(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintMsgObservedTxout(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsgObservedTxout(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgObservedTxOut) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Txs) > 0 {
		for _, e := range m.Txs {
			l = e.Size()
			n += 1 + l + sovMsgObservedTxout(uint64(l))
		}
	}
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovMsgObservedTxout(uint64(l))
	}
	return n
}

func sovMsgObservedTxout(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsgObservedTxout(x uint64) (n int) {
	return sovMsgObservedTxout(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgObservedTxOut) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgObservedTxout
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
			return fmt.Errorf("proto: MsgObservedTxOut: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgObservedTxOut: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Txs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgObservedTxout
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
				return ErrInvalidLengthMsgObservedTxout
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsgObservedTxout
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Txs = append(m.Txs, ObservedTx{})
			if err := m.Txs[len(m.Txs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgObservedTxout
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
				return ErrInvalidLengthMsgObservedTxout
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgObservedTxout
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
			skippy, err := skipMsgObservedTxout(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgObservedTxout
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
func skipMsgObservedTxout(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsgObservedTxout
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
					return 0, ErrIntOverflowMsgObservedTxout
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
					return 0, ErrIntOverflowMsgObservedTxout
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
				return 0, ErrInvalidLengthMsgObservedTxout
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsgObservedTxout
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsgObservedTxout
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsgObservedTxout        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsgObservedTxout          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsgObservedTxout = fmt.Errorf("proto: unexpected end of group")
)
