// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: thorchain/v1/x/thorchain/types/type_rune_provider.proto

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

type RUNEProvider struct {
	RuneAddress        github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=rune_address,json=runeAddress,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"rune_address,omitempty"`
	DepositAmount      github_com_cosmos_cosmos_sdk_types.Uint       `protobuf:"bytes,2,opt,name=deposit_amount,json=depositAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Uint" json:"deposit_amount"`
	WithdrawAmount     github_com_cosmos_cosmos_sdk_types.Uint       `protobuf:"bytes,3,opt,name=withdraw_amount,json=withdrawAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Uint" json:"withdraw_amount"`
	Units              github_com_cosmos_cosmos_sdk_types.Uint       `protobuf:"bytes,4,opt,name=units,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Uint" json:"units"`
	LastDepositHeight  int64                                         `protobuf:"varint,5,opt,name=last_deposit_height,json=lastDepositHeight,proto3" json:"last_deposit_height,omitempty"`
	LastWithdrawHeight int64                                         `protobuf:"varint,6,opt,name=last_withdraw_height,json=lastWithdrawHeight,proto3" json:"last_withdraw_height,omitempty"`
}

func (m *RUNEProvider) Reset()         { *m = RUNEProvider{} }
func (m *RUNEProvider) String() string { return proto.CompactTextString(m) }
func (*RUNEProvider) ProtoMessage()    {}
func (*RUNEProvider) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5b76a6b4f418288, []int{0}
}
func (m *RUNEProvider) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RUNEProvider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RUNEProvider.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RUNEProvider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RUNEProvider.Merge(m, src)
}
func (m *RUNEProvider) XXX_Size() int {
	return m.Size()
}
func (m *RUNEProvider) XXX_DiscardUnknown() {
	xxx_messageInfo_RUNEProvider.DiscardUnknown(m)
}

var xxx_messageInfo_RUNEProvider proto.InternalMessageInfo

func (m *RUNEProvider) GetRuneAddress() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.RuneAddress
	}
	return nil
}

func (m *RUNEProvider) GetLastDepositHeight() int64 {
	if m != nil {
		return m.LastDepositHeight
	}
	return 0
}

func (m *RUNEProvider) GetLastWithdrawHeight() int64 {
	if m != nil {
		return m.LastWithdrawHeight
	}
	return 0
}

func init() {
	proto.RegisterType((*RUNEProvider)(nil), "types.RUNEProvider")
}

func init() {
	proto.RegisterFile("thorchain/v1/x/thorchain/types/type_rune_provider.proto", fileDescriptor_e5b76a6b4f418288)
}

var fileDescriptor_e5b76a6b4f418288 = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x41, 0x4f, 0xf2, 0x30,
	0x18, 0xc7, 0xb7, 0x97, 0x17, 0x12, 0x2b, 0x62, 0x9c, 0x1c, 0x16, 0x0f, 0x83, 0x78, 0x91, 0x0b,
	0xab, 0xc4, 0x83, 0x67, 0x88, 0x24, 0x7a, 0x31, 0x66, 0x11, 0x35, 0x5e, 0x96, 0xb1, 0x36, 0x6b,
	0x23, 0xac, 0x4b, 0xdb, 0x81, 0x7e, 0x0b, 0xbf, 0x91, 0x57, 0x8e, 0x1c, 0x8d, 0x07, 0x62, 0xe0,
	0x5b, 0x78, 0x32, 0x6b, 0x8b, 0x9a, 0x78, 0x31, 0x5c, 0xb6, 0x3e, 0xfd, 0xf5, 0xff, 0xcb, 0xf3,
	0x24, 0x0f, 0x38, 0x95, 0x84, 0xf1, 0x98, 0x44, 0x34, 0x85, 0x93, 0x0e, 0x7c, 0x84, 0xdf, 0xa5,
	0x7c, 0xca, 0xb0, 0x50, 0xdf, 0x90, 0xe7, 0x29, 0x0e, 0x33, 0xce, 0x26, 0x14, 0x61, 0xee, 0x67,
	0x9c, 0x49, 0xe6, 0x94, 0x15, 0x3f, 0xa8, 0x27, 0x2c, 0x61, 0xea, 0x06, 0x16, 0x27, 0x0d, 0x0f,
	0x5f, 0x4a, 0xa0, 0x1a, 0x0c, 0x2e, 0xfb, 0x57, 0x26, 0xe3, 0x5c, 0x83, 0xaa, 0x92, 0x44, 0x08,
	0x71, 0x2c, 0x84, 0x6b, 0x37, 0xed, 0x56, 0xb5, 0xd7, 0xf9, 0x58, 0x34, 0xda, 0x09, 0x95, 0x24,
	0x1f, 0xfa, 0x31, 0x1b, 0xc3, 0x98, 0x89, 0x31, 0x13, 0xe6, 0xd7, 0x16, 0xe8, 0x41, 0x37, 0xe1,
	0x77, 0xe3, 0xb8, 0xab, 0x83, 0xc1, 0x76, 0xa1, 0x31, 0x85, 0x73, 0x03, 0x6a, 0x08, 0x67, 0x4c,
	0x50, 0x19, 0x46, 0x63, 0x96, 0xa7, 0xd2, 0xfd, 0xd7, 0xb4, 0x5b, 0x5b, 0x3d, 0x38, 0x5b, 0x34,
	0xac, 0xb7, 0x45, 0xe3, 0xe8, 0x0f, 0xee, 0x01, 0x4d, 0x65, 0xb0, 0x63, 0x34, 0x5d, 0x65, 0x71,
	0xee, 0xc0, 0xee, 0x94, 0x4a, 0x82, 0x78, 0x34, 0x5d, 0x8b, 0x4b, 0x9b, 0x89, 0x6b, 0x6b, 0x8f,
	0x31, 0xf7, 0x41, 0x39, 0x4f, 0xa9, 0x14, 0xee, 0xff, 0xcd, 0x7c, 0x3a, 0xed, 0xf8, 0x60, 0x7f,
	0x14, 0x09, 0x19, 0xae, 0xa7, 0x27, 0x98, 0x26, 0x44, 0xba, 0xe5, 0xa6, 0xdd, 0x2a, 0x05, 0x7b,
	0x05, 0x3a, 0xd3, 0xe4, 0x5c, 0x01, 0xe7, 0x18, 0xd4, 0xd5, 0xfb, 0xaf, 0xa9, 0x4c, 0xa0, 0xa2,
	0x02, 0x4e, 0xc1, 0x6e, 0x0d, 0xd2, 0x89, 0xde, 0xc5, 0x6c, 0xe9, 0xd9, 0xf3, 0xa5, 0x67, 0xbf,
	0x2f, 0x3d, 0xfb, 0x79, 0xe5, 0x59, 0xf3, 0x95, 0x67, 0xbd, 0xae, 0x3c, 0xeb, 0x1e, 0x26, 0x54,
	0x8e, 0x22, 0xdd, 0xeb, 0x8f, 0x6d, 0x21, 0x8c, 0xa7, 0x0c, 0xe1, 0xdf, 0x2b, 0x34, 0xac, 0xa8,
	0x9d, 0x38, 0xf9, 0x0c, 0x00, 0x00, 0xff, 0xff, 0xb3, 0xdb, 0x88, 0x9c, 0x6b, 0x02, 0x00, 0x00,
}

func (m *RUNEProvider) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RUNEProvider) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RUNEProvider) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LastWithdrawHeight != 0 {
		i = encodeVarintTypeRuneProvider(dAtA, i, uint64(m.LastWithdrawHeight))
		i--
		dAtA[i] = 0x30
	}
	if m.LastDepositHeight != 0 {
		i = encodeVarintTypeRuneProvider(dAtA, i, uint64(m.LastDepositHeight))
		i--
		dAtA[i] = 0x28
	}
	{
		size := m.Units.Size()
		i -= size
		if _, err := m.Units.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypeRuneProvider(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.WithdrawAmount.Size()
		i -= size
		if _, err := m.WithdrawAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypeRuneProvider(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.DepositAmount.Size()
		i -= size
		if _, err := m.DepositAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypeRuneProvider(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.RuneAddress) > 0 {
		i -= len(m.RuneAddress)
		copy(dAtA[i:], m.RuneAddress)
		i = encodeVarintTypeRuneProvider(dAtA, i, uint64(len(m.RuneAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypeRuneProvider(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypeRuneProvider(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RUNEProvider) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RuneAddress)
	if l > 0 {
		n += 1 + l + sovTypeRuneProvider(uint64(l))
	}
	l = m.DepositAmount.Size()
	n += 1 + l + sovTypeRuneProvider(uint64(l))
	l = m.WithdrawAmount.Size()
	n += 1 + l + sovTypeRuneProvider(uint64(l))
	l = m.Units.Size()
	n += 1 + l + sovTypeRuneProvider(uint64(l))
	if m.LastDepositHeight != 0 {
		n += 1 + sovTypeRuneProvider(uint64(m.LastDepositHeight))
	}
	if m.LastWithdrawHeight != 0 {
		n += 1 + sovTypeRuneProvider(uint64(m.LastWithdrawHeight))
	}
	return n
}

func sovTypeRuneProvider(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypeRuneProvider(x uint64) (n int) {
	return sovTypeRuneProvider(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RUNEProvider) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypeRuneProvider
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
			return fmt.Errorf("proto: RUNEProvider: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RUNEProvider: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RuneAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeRuneProvider
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
				return ErrInvalidLengthTypeRuneProvider
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypeRuneProvider
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RuneAddress = append(m.RuneAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.RuneAddress == nil {
				m.RuneAddress = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeRuneProvider
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
				return ErrInvalidLengthTypeRuneProvider
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypeRuneProvider
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DepositAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeRuneProvider
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
				return ErrInvalidLengthTypeRuneProvider
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypeRuneProvider
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.WithdrawAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Units", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeRuneProvider
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
				return ErrInvalidLengthTypeRuneProvider
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypeRuneProvider
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Units.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastDepositHeight", wireType)
			}
			m.LastDepositHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeRuneProvider
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastDepositHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastWithdrawHeight", wireType)
			}
			m.LastWithdrawHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeRuneProvider
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastWithdrawHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypeRuneProvider(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypeRuneProvider
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypeRuneProvider
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
func skipTypeRuneProvider(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypeRuneProvider
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
					return 0, ErrIntOverflowTypeRuneProvider
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
					return 0, ErrIntOverflowTypeRuneProvider
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
				return 0, ErrInvalidLengthTypeRuneProvider
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypeRuneProvider
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypeRuneProvider
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypeRuneProvider        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypeRuneProvider          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypeRuneProvider = fmt.Errorf("proto: unexpected end of group")
)
