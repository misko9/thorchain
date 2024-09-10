// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: thorchain/v1/x/thorchain/types/msg_migrate.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type MsgMigrate struct {
	Tx          ObservedTx                                    `protobuf:"bytes,1,opt,name=tx,proto3" json:"tx"`
	BlockHeight int64                                         `protobuf:"varint,2,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	Signer      github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,3,opt,name=signer,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"signer,omitempty"`
}

func (m *MsgMigrate) Reset()         { *m = MsgMigrate{} }
func (m *MsgMigrate) String() string { return proto.CompactTextString(m) }
func (*MsgMigrate) ProtoMessage()    {}
func (*MsgMigrate) Descriptor() ([]byte, []int) {
	return fileDescriptor_632e7cb7718287df, []int{0}
}
func (m *MsgMigrate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgMigrate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgMigrate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgMigrate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgMigrate.Merge(m, src)
}
func (m *MsgMigrate) XXX_Size() int {
	return m.Size()
}
func (m *MsgMigrate) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgMigrate.DiscardUnknown(m)
}

var xxx_messageInfo_MsgMigrate proto.InternalMessageInfo

func (m *MsgMigrate) GetTx() ObservedTx {
	if m != nil {
		return m.Tx
	}
	return ObservedTx{}
}

func (m *MsgMigrate) GetBlockHeight() int64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *MsgMigrate) GetSigner() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Signer
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgMigrate)(nil), "types.MsgMigrate")
}

func init() {
	proto.RegisterFile("thorchain/v1/x/thorchain/types/msg_migrate.proto", fileDescriptor_632e7cb7718287df)
}

var fileDescriptor_632e7cb7718287df = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x28, 0xc9, 0xc8, 0x2f,
	0x4a, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x2f, 0x33, 0xd4, 0xaf, 0xd0, 0x47, 0x70, 0x4b, 0x2a, 0x0b,
	0x52, 0x8b, 0xf5, 0x73, 0x8b, 0xd3, 0xe3, 0x73, 0x33, 0xd3, 0x8b, 0x12, 0x4b, 0x52, 0xf5, 0x0a,
	0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x58, 0xc1, 0x12, 0x52, 0xa6, 0x04, 0x34, 0x82, 0xc8, 0xf8, 0xfc,
	0xa4, 0xe2, 0xd4, 0xa2, 0xb2, 0xd4, 0x94, 0xf8, 0x92, 0x0a, 0x88, 0x6e, 0x29, 0x91, 0xf4, 0xfc,
	0xf4, 0x7c, 0x30, 0x53, 0x1f, 0xc4, 0x82, 0x88, 0x2a, 0x2d, 0x66, 0xe4, 0xe2, 0xf2, 0x2d, 0x4e,
	0xf7, 0x85, 0x58, 0x24, 0xa4, 0xce, 0xc5, 0x54, 0x52, 0x21, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x6d,
	0x24, 0xa8, 0x07, 0x36, 0x4f, 0xcf, 0x1f, 0x6a, 0x54, 0x48, 0x85, 0x13, 0xcb, 0x89, 0x7b, 0xf2,
	0x0c, 0x41, 0x4c, 0x25, 0x15, 0x42, 0x8a, 0x5c, 0x3c, 0x49, 0x39, 0xf9, 0xc9, 0xd9, 0xf1, 0x19,
	0xa9, 0x99, 0xe9, 0x19, 0x25, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0xdc, 0x60, 0x31, 0x0f,
	0xb0, 0x90, 0x90, 0x27, 0x17, 0x5b, 0x71, 0x66, 0x7a, 0x5e, 0x6a, 0x91, 0x04, 0xb3, 0x02, 0xa3,
	0x06, 0x8f, 0x93, 0xe1, 0xaf, 0x7b, 0xf2, 0xba, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9,
	0xf9, 0xb9, 0xfa, 0xc9, 0xf9, 0xc5, 0xb9, 0xf9, 0xc5, 0x50, 0x4a, 0xb7, 0x38, 0x25, 0x1b, 0xe2,
	0x7e, 0x3d, 0xc7, 0xe4, 0x64, 0xc7, 0x94, 0x94, 0xa2, 0xd4, 0xe2, 0xe2, 0x20, 0xa8, 0x01, 0x4e,
	0x9e, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7,
	0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0xa5, 0x9f, 0x9e, 0x59, 0x92,
	0x93, 0x08, 0x31, 0x10, 0x29, 0x20, 0x32, 0xf2, 0x8b, 0xf2, 0xf2, 0x53, 0x52, 0x31, 0x43, 0x27,
	0x89, 0x0d, 0xec, 0x6f, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x59, 0xc9, 0xd8, 0xd2, 0x7f,
	0x01, 0x00, 0x00,
}

func (m *MsgMigrate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgMigrate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgMigrate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintMsgMigrate(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0x1a
	}
	if m.BlockHeight != 0 {
		i = encodeVarintMsgMigrate(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.Tx.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMsgMigrate(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintMsgMigrate(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsgMigrate(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgMigrate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Tx.Size()
	n += 1 + l + sovMsgMigrate(uint64(l))
	if m.BlockHeight != 0 {
		n += 1 + sovMsgMigrate(uint64(m.BlockHeight))
	}
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovMsgMigrate(uint64(l))
	}
	return n
}

func sovMsgMigrate(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsgMigrate(x uint64) (n int) {
	return sovMsgMigrate(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgMigrate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgMigrate
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
			return fmt.Errorf("proto: MsgMigrate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgMigrate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tx", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgMigrate
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
				return ErrInvalidLengthMsgMigrate
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsgMigrate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Tx.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgMigrate
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgMigrate
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
				return ErrInvalidLengthMsgMigrate
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgMigrate
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
			skippy, err := skipMsgMigrate(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgMigrate
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
func skipMsgMigrate(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsgMigrate
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
					return 0, ErrIntOverflowMsgMigrate
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
					return 0, ErrIntOverflowMsgMigrate
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
				return 0, ErrInvalidLengthMsgMigrate
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsgMigrate
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsgMigrate
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsgMigrate        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsgMigrate          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsgMigrate = fmt.Errorf("proto: unexpected end of group")
)
