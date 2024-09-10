// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: types/type_rune_provider.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
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

type RUNEProvider struct {
	RuneAddress        github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=rune_address,json=runeAddress,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"rune_address,omitempty"`
	DepositAmount      cosmossdk_io_math.Uint                        `protobuf:"bytes,2,opt,name=deposit_amount,json=depositAmount,proto3,customtype=cosmossdk.io/math.Uint" json:"deposit_amount"`
	WithdrawAmount     cosmossdk_io_math.Uint                        `protobuf:"bytes,3,opt,name=withdraw_amount,json=withdrawAmount,proto3,customtype=cosmossdk.io/math.Uint" json:"withdraw_amount"`
	Units              cosmossdk_io_math.Uint                        `protobuf:"bytes,4,opt,name=units,proto3,customtype=cosmossdk.io/math.Uint" json:"units"`
	LastDepositHeight  int64                                         `protobuf:"varint,5,opt,name=last_deposit_height,json=lastDepositHeight,proto3" json:"last_deposit_height,omitempty"`
	LastWithdrawHeight int64                                         `protobuf:"varint,6,opt,name=last_withdraw_height,json=lastWithdrawHeight,proto3" json:"last_withdraw_height,omitempty"`
}

func (m *RUNEProvider) Reset()         { *m = RUNEProvider{} }
func (m *RUNEProvider) String() string { return proto.CompactTextString(m) }
func (*RUNEProvider) ProtoMessage()    {}
func (*RUNEProvider) Descriptor() ([]byte, []int) {
	return fileDescriptor_52cac3b125100353, []int{0}
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

func init() { proto.RegisterFile("types/type_rune_provider.proto", fileDescriptor_52cac3b125100353) }

var fileDescriptor_52cac3b125100353 = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xb1, 0x4e, 0xeb, 0x30,
	0x14, 0x86, 0x93, 0xdb, 0xdb, 0x4a, 0xd7, 0xb7, 0x14, 0x11, 0x2a, 0x14, 0x31, 0xb8, 0x15, 0x53,
	0x97, 0x26, 0x20, 0x78, 0x81, 0x56, 0x54, 0xc0, 0x82, 0x50, 0x44, 0x85, 0xc4, 0x12, 0xb9, 0xb1,
	0x15, 0x5b, 0x6d, 0xe2, 0xc8, 0x76, 0x28, 0xbc, 0x03, 0x03, 0x8f, 0xd5, 0xb1, 0x23, 0x62, 0xa8,
	0x50, 0xfb, 0x16, 0x4c, 0x28, 0xb6, 0x8b, 0x18, 0xbb, 0x24, 0x27, 0xe7, 0x3b, 0xdf, 0x9f, 0xc4,
	0x07, 0x40, 0xf5, 0x52, 0x10, 0x19, 0x56, 0xd7, 0x58, 0x94, 0x39, 0x89, 0x0b, 0xc1, 0x9f, 0x18,
	0x26, 0x22, 0x28, 0x04, 0x57, 0xdc, 0xab, 0x6b, 0x7e, 0xdc, 0x4e, 0x79, 0xca, 0x75, 0x27, 0xac,
	0x2a, 0x03, 0x4f, 0x5e, 0x6b, 0xa0, 0x19, 0x8d, 0x6f, 0x47, 0x77, 0xd6, 0xf1, 0xee, 0x41, 0x53,
	0x87, 0x20, 0x8c, 0x05, 0x91, 0xd2, 0x77, 0xbb, 0x6e, 0xaf, 0x39, 0x3c, 0xfb, 0x5a, 0x75, 0xfa,
	0x29, 0x53, 0xb4, 0x9c, 0x04, 0x09, 0xcf, 0xc2, 0x84, 0xcb, 0x8c, 0x4b, 0x7b, 0xeb, 0x4b, 0x3c,
	0xd5, 0xaf, 0x97, 0xc1, 0x20, 0x49, 0x06, 0x46, 0x8c, 0xfe, 0x57, 0x31, 0xf6, 0xc1, 0x1b, 0x81,
	0x16, 0x26, 0x05, 0x97, 0x4c, 0xc5, 0x28, 0xe3, 0x65, 0xae, 0xfc, 0x3f, 0x5d, 0xb7, 0xf7, 0x6f,
	0x08, 0x17, 0xab, 0x8e, 0xf3, 0xb1, 0xea, 0x1c, 0x99, 0x24, 0x89, 0xa7, 0x01, 0xe3, 0x61, 0x86,
	0x14, 0x0d, 0xc6, 0x2c, 0x57, 0xd1, 0x9e, 0xb5, 0x06, 0x5a, 0xf2, 0xae, 0xc0, 0xfe, 0x9c, 0x29,
	0x8a, 0x05, 0x9a, 0x6f, 0x73, 0x6a, 0x3b, 0xe5, 0xb4, 0xb6, 0x9a, 0x0d, 0xba, 0x00, 0xf5, 0x32,
	0x67, 0x4a, 0xfa, 0x7f, 0x77, 0xd2, 0xcd, 0xb0, 0x17, 0x80, 0xc3, 0x19, 0x92, 0x2a, 0xde, 0xfe,
	0x0a, 0x25, 0x2c, 0xa5, 0xca, 0xaf, 0x77, 0xdd, 0x5e, 0x2d, 0x3a, 0xa8, 0xd0, 0xa5, 0x21, 0xd7,
	0x1a, 0x78, 0xa7, 0xa0, 0xad, 0xe7, 0x7f, 0xbe, 0xd9, 0x0a, 0x0d, 0x2d, 0x78, 0x15, 0x7b, 0xb0,
	0xc8, 0x18, 0xc3, 0x9b, 0xc5, 0x1a, 0xba, 0xcb, 0x35, 0x74, 0x3f, 0xd7, 0xd0, 0x7d, 0xdb, 0x40,
	0x67, 0xb9, 0x81, 0xce, 0xfb, 0x06, 0x3a, 0x8f, 0x61, 0xca, 0xd4, 0x0c, 0x99, 0xd3, 0x57, 0x94,
	0x8b, 0x84, 0x22, 0x96, 0xeb, 0x2a, 0xe7, 0x98, 0x84, 0xcf, 0xbf, 0x9b, 0xd5, 0x2a, 0x26, 0x0d,
	0xbd, 0xe0, 0xf3, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x45, 0xef, 0x19, 0x63, 0x1f, 0x02, 0x00,
	0x00,
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
			if (skippy < 0) || (iNdEx+skippy) < 0 {
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
