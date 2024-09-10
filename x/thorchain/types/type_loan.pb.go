// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: thorchain/v1/x/thorchain/types/type_loan.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	common "gitlab.com/thorchain/thornode/common"
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

type Loan struct {
	Owner               gitlab_com_thorchain_thornode_common.Address `protobuf:"bytes,1,opt,name=owner,proto3,casttype=gitlab.com/thorchain/thornode/common.Address" json:"owner,omitempty"`
	Asset               common.Asset                                 `protobuf:"bytes,2,opt,name=asset,proto3" json:"asset"`
	DebtIssued          cosmossdk_io_math.Uint                       `protobuf:"bytes,3,opt,name=debt_issued,json=debtIssued,proto3,customtype=cosmossdk.io/math.Uint" json:"debt_issued"`
	DebtRepaid          cosmossdk_io_math.Uint                       `protobuf:"bytes,4,opt,name=debt_repaid,json=debtRepaid,proto3,customtype=cosmossdk.io/math.Uint" json:"debt_repaid"`
	CollateralDeposited cosmossdk_io_math.Uint                       `protobuf:"bytes,5,opt,name=collateral_deposited,json=collateralDeposited,proto3,customtype=cosmossdk.io/math.Uint" json:"collateral_deposited"`
	CollateralWithdrawn cosmossdk_io_math.Uint                       `protobuf:"bytes,6,opt,name=collateral_withdrawn,json=collateralWithdrawn,proto3,customtype=cosmossdk.io/math.Uint" json:"collateral_withdrawn"`
	LastOpenHeight      int64                                        `protobuf:"varint,9,opt,name=last_open_height,json=lastOpenHeight,proto3" json:"last_open_height,omitempty"`
	LastRepayHeight     int64                                        `protobuf:"varint,10,opt,name=last_repay_height,json=lastRepayHeight,proto3" json:"last_repay_height,omitempty"`
}

func (m *Loan) Reset()         { *m = Loan{} }
func (m *Loan) String() string { return proto.CompactTextString(m) }
func (*Loan) ProtoMessage()    {}
func (*Loan) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2061da80bf428b4, []int{0}
}
func (m *Loan) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Loan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Loan.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Loan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Loan.Merge(m, src)
}
func (m *Loan) XXX_Size() int {
	return m.Size()
}
func (m *Loan) XXX_DiscardUnknown() {
	xxx_messageInfo_Loan.DiscardUnknown(m)
}

var xxx_messageInfo_Loan proto.InternalMessageInfo

func (m *Loan) GetOwner() gitlab_com_thorchain_thornode_common.Address {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Loan) GetAsset() common.Asset {
	if m != nil {
		return m.Asset
	}
	return common.Asset{}
}

func (m *Loan) GetLastOpenHeight() int64 {
	if m != nil {
		return m.LastOpenHeight
	}
	return 0
}

func (m *Loan) GetLastRepayHeight() int64 {
	if m != nil {
		return m.LastRepayHeight
	}
	return 0
}

func init() {
	proto.RegisterType((*Loan)(nil), "types.Loan")
}

func init() {
	proto.RegisterFile("thorchain/v1/x/thorchain/types/type_loan.proto", fileDescriptor_c2061da80bf428b4)
}

var fileDescriptor_c2061da80bf428b4 = []byte{
	// 394 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcb, 0x6a, 0xdb, 0x40,
	0x14, 0x86, 0xa5, 0xda, 0x32, 0x78, 0x4c, 0x6f, 0xaa, 0x29, 0xc2, 0x0b, 0x59, 0x74, 0xa5, 0x96,
	0x22, 0xf5, 0xf2, 0x00, 0xa5, 0xa6, 0x94, 0x1a, 0x0a, 0xa5, 0x82, 0x52, 0xe8, 0x46, 0x8c, 0x35,
	0x83, 0x34, 0x54, 0x9a, 0x23, 0x66, 0x26, 0x71, 0xfc, 0x16, 0x79, 0x2c, 0x2f, 0xbd, 0x0c, 0x59,
	0x98, 0x60, 0x3f, 0x44, 0x20, 0xab, 0x30, 0xa3, 0x28, 0xce, 0x05, 0x82, 0xb3, 0x91, 0x46, 0xbf,
	0xbe, 0xef, 0xd7, 0x65, 0x0e, 0x8a, 0x54, 0x01, 0x22, 0x2b, 0x30, 0xe3, 0xf1, 0xe1, 0xc7, 0xf8,
	0x28, 0xde, 0x5d, 0xaa, 0x45, 0x4d, 0xa5, 0x39, 0xa6, 0x25, 0x60, 0x1e, 0xd5, 0x02, 0x14, 0xb8,
	0x8e, 0x89, 0x47, 0xc1, 0x2d, 0x2d, 0x83, 0xaa, 0x02, 0x7e, 0x75, 0x6a, 0xc0, 0xd1, 0x30, 0x87,
	0x1c, 0xcc, 0x32, 0xd6, 0xab, 0x26, 0x7d, 0x73, 0xde, 0x41, 0xdd, 0x9f, 0x80, 0xb9, 0xfb, 0x1d,
	0x39, 0x30, 0xe7, 0x54, 0x78, 0x76, 0x60, 0x87, 0xfd, 0xc9, 0x87, 0x8b, 0xf5, 0xf8, 0x7d, 0xce,
	0x54, 0x89, 0x67, 0x51, 0x06, 0xd5, 0xcd, 0xd7, 0x28, 0x40, 0x70, 0x20, 0xb4, 0x6d, 0xff, 0x4a,
	0x88, 0xa0, 0x52, 0x26, 0x8d, 0xee, 0xbe, 0x45, 0x0e, 0x96, 0x92, 0x2a, 0xef, 0x49, 0x60, 0x87,
	0x83, 0x4f, 0x4f, 0xa3, 0x16, 0xd3, 0xe1, 0xa4, 0xbb, 0x5c, 0x8f, 0xad, 0xa4, 0x21, 0xdc, 0x2f,
	0x68, 0x40, 0xe8, 0x4c, 0xa5, 0x4c, 0xca, 0x03, 0x4a, 0xbc, 0x8e, 0x79, 0xb0, 0xaf, 0x89, 0xd3,
	0xf5, 0xf8, 0x75, 0x06, 0xb2, 0x02, 0x29, 0xc9, 0xff, 0x88, 0x41, 0x5c, 0x61, 0x55, 0x44, 0x7f,
	0x18, 0x57, 0x09, 0xd2, 0xca, 0xd4, 0x18, 0xd7, 0x05, 0x82, 0xd6, 0x98, 0x11, 0xaf, 0xbb, 0x7f,
	0x41, 0x62, 0x0c, 0xf7, 0x37, 0x1a, 0x66, 0x50, 0x96, 0x58, 0x51, 0x81, 0xcb, 0x94, 0xd0, 0x1a,
	0x24, 0x53, 0x94, 0x78, 0xce, 0x5e, 0x4d, 0xaf, 0x76, 0xee, 0xb7, 0x56, 0xbd, 0x53, 0x39, 0x67,
	0xaa, 0x20, 0x02, 0xcf, 0xb9, 0xd7, 0x7b, 0x6c, 0xe5, 0xdf, 0x56, 0x75, 0x43, 0xf4, 0xa2, 0xc4,
	0x52, 0xa5, 0x50, 0x53, 0x9e, 0x16, 0x94, 0xe5, 0x85, 0xf2, 0xfa, 0x81, 0x1d, 0x76, 0x92, 0x67,
	0x3a, 0xff, 0x55, 0x53, 0xfe, 0xc3, 0xa4, 0xee, 0x3b, 0xf4, 0xd2, 0x90, 0xfa, 0x87, 0x2c, 0x5a,
	0x14, 0x19, 0xf4, 0xb9, 0xbe, 0xa1, 0x3f, 0x7b, 0xd1, 0xb0, 0x93, 0xe9, 0x72, 0xe3, 0xdb, 0xab,
	0x8d, 0x6f, 0x9f, 0x6d, 0x7c, 0xfb, 0x78, 0xeb, 0x5b, 0xab, 0xad, 0x6f, 0x9d, 0x6c, 0x7d, 0xeb,
	0x5f, 0xfc, 0xf0, 0xbe, 0xdf, 0x9b, 0xc9, 0x59, 0xcf, 0xcc, 0xd2, 0xe7, 0xcb, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x70, 0xeb, 0xaf, 0x51, 0xbc, 0x02, 0x00, 0x00,
}

func (m *Loan) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Loan) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Loan) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LastRepayHeight != 0 {
		i = encodeVarintTypeLoan(dAtA, i, uint64(m.LastRepayHeight))
		i--
		dAtA[i] = 0x50
	}
	if m.LastOpenHeight != 0 {
		i = encodeVarintTypeLoan(dAtA, i, uint64(m.LastOpenHeight))
		i--
		dAtA[i] = 0x48
	}
	if len(m.CollateralWithdrawn) > 0 {
		i -= len(m.CollateralWithdrawn)
		copy(dAtA[i:], m.CollateralWithdrawn)
		i = encodeVarintTypeLoan(dAtA, i, uint64(len(m.CollateralWithdrawn)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.CollateralDeposited) > 0 {
		i -= len(m.CollateralDeposited)
		copy(dAtA[i:], m.CollateralDeposited)
		i = encodeVarintTypeLoan(dAtA, i, uint64(len(m.CollateralDeposited)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.DebtRepaid) > 0 {
		i -= len(m.DebtRepaid)
		copy(dAtA[i:], m.DebtRepaid)
		i = encodeVarintTypeLoan(dAtA, i, uint64(len(m.DebtRepaid)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.DebtIssued) > 0 {
		i -= len(m.DebtIssued)
		copy(dAtA[i:], m.DebtIssued)
		i = encodeVarintTypeLoan(dAtA, i, uint64(len(m.DebtIssued)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size, err := m.Asset.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypeLoan(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintTypeLoan(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypeLoan(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypeLoan(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Loan) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovTypeLoan(uint64(l))
	}
	l = m.Asset.Size()
	n += 1 + l + sovTypeLoan(uint64(l))
	l = len(m.DebtIssued)
	if l > 0 {
		n += 1 + l + sovTypeLoan(uint64(l))
	}
	l = len(m.DebtRepaid)
	if l > 0 {
		n += 1 + l + sovTypeLoan(uint64(l))
	}
	l = len(m.CollateralDeposited)
	if l > 0 {
		n += 1 + l + sovTypeLoan(uint64(l))
	}
	l = len(m.CollateralWithdrawn)
	if l > 0 {
		n += 1 + l + sovTypeLoan(uint64(l))
	}
	if m.LastOpenHeight != 0 {
		n += 1 + sovTypeLoan(uint64(m.LastOpenHeight))
	}
	if m.LastRepayHeight != 0 {
		n += 1 + sovTypeLoan(uint64(m.LastRepayHeight))
	}
	return n
}

func sovTypeLoan(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypeLoan(x uint64) (n int) {
	return sovTypeLoan(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Loan) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypeLoan
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
			return fmt.Errorf("proto: Loan: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Loan: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeLoan
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
				return ErrInvalidLengthTypeLoan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypeLoan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = gitlab_com_thorchain_thornode_common.Address(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Asset", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeLoan
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
				return ErrInvalidLengthTypeLoan
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypeLoan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Asset.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DebtIssued", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeLoan
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
				return ErrInvalidLengthTypeLoan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypeLoan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DebtIssued = cosmossdk_io_math.Uint(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DebtRepaid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeLoan
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
				return ErrInvalidLengthTypeLoan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypeLoan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DebtRepaid = cosmossdk_io_math.Uint(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollateralDeposited", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeLoan
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
				return ErrInvalidLengthTypeLoan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypeLoan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CollateralDeposited = cosmossdk_io_math.Uint(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollateralWithdrawn", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeLoan
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
				return ErrInvalidLengthTypeLoan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypeLoan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CollateralWithdrawn = cosmossdk_io_math.Uint(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastOpenHeight", wireType)
			}
			m.LastOpenHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeLoan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastOpenHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastRepayHeight", wireType)
			}
			m.LastRepayHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeLoan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastRepayHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypeLoan(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypeLoan
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
func skipTypeLoan(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypeLoan
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
					return 0, ErrIntOverflowTypeLoan
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
					return 0, ErrIntOverflowTypeLoan
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
				return 0, ErrInvalidLengthTypeLoan
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypeLoan
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypeLoan
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypeLoan        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypeLoan          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypeLoan = fmt.Errorf("proto: unexpected end of group")
)
