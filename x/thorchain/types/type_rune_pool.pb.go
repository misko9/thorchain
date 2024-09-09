// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: thorchain/v1/x/thorchain/types/type_rune_pool.proto

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

// RUNEPool represents ownership of currently active POL.
type RUNEPool struct {
	ReserveUnits  github_com_cosmos_cosmos_sdk_types.Uint `protobuf:"bytes,1,opt,name=reserve_units,json=reserveUnits,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Uint" json:"reserve_units"`
	PoolUnits     github_com_cosmos_cosmos_sdk_types.Uint `protobuf:"bytes,2,opt,name=pool_units,json=poolUnits,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Uint" json:"pool_units"`
	RuneDeposited github_com_cosmos_cosmos_sdk_types.Uint `protobuf:"bytes,3,opt,name=rune_deposited,json=runeDeposited,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Uint" json:"rune_deposited"`
	RuneWithdrawn github_com_cosmos_cosmos_sdk_types.Uint `protobuf:"bytes,4,opt,name=rune_withdrawn,json=runeWithdrawn,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Uint" json:"rune_withdrawn"`
}

func (m *RUNEPool) Reset()         { *m = RUNEPool{} }
func (m *RUNEPool) String() string { return proto.CompactTextString(m) }
func (*RUNEPool) ProtoMessage()    {}
func (*RUNEPool) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c2b9af96e410e76, []int{0}
}
func (m *RUNEPool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RUNEPool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RUNEPool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RUNEPool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RUNEPool.Merge(m, src)
}
func (m *RUNEPool) XXX_Size() int {
	return m.Size()
}
func (m *RUNEPool) XXX_DiscardUnknown() {
	xxx_messageInfo_RUNEPool.DiscardUnknown(m)
}

var xxx_messageInfo_RUNEPool proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RUNEPool)(nil), "types.RUNEPool")
}

func init() {
	proto.RegisterFile("thorchain/v1/x/thorchain/types/type_rune_pool.proto", fileDescriptor_3c2b9af96e410e76)
}

var fileDescriptor_3c2b9af96e410e76 = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x2e, 0xc9, 0xc8, 0x2f,
	0x4a, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x2f, 0x33, 0xd4, 0xaf, 0xd0, 0x47, 0x70, 0x4b, 0x2a, 0x0b,
	0x52, 0x8b, 0xc1, 0x64, 0x7c, 0x51, 0x69, 0x5e, 0x6a, 0x7c, 0x41, 0x7e, 0x7e, 0x8e, 0x5e, 0x41,
	0x51, 0x7e, 0x49, 0xbe, 0x10, 0x2b, 0x58, 0x4e, 0x4a, 0x24, 0x3d, 0x3f, 0x3d, 0x1f, 0x2c, 0xa2,
	0x0f, 0x62, 0x41, 0x24, 0x95, 0x1e, 0x30, 0x71, 0x71, 0x04, 0x85, 0xfa, 0xb9, 0x06, 0xe4, 0xe7,
	0xe7, 0x08, 0x85, 0x70, 0xf1, 0x16, 0xa5, 0x16, 0xa7, 0x16, 0x95, 0xa5, 0xc6, 0x97, 0xe6, 0x65,
	0x96, 0x14, 0x4b, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x3a, 0xe9, 0x9f, 0xb8, 0x27, 0xcf, 0x70, 0xeb,
	0x9e, 0xbc, 0x7a, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e, 0x72, 0x7e,
	0x71, 0x6e, 0x7e, 0x31, 0x94, 0xd2, 0x2d, 0x4e, 0xc9, 0x86, 0xb8, 0x40, 0x2f, 0x34, 0x33, 0xaf,
	0x24, 0x88, 0x07, 0x6a, 0x4a, 0x28, 0xc8, 0x10, 0x21, 0x3f, 0x2e, 0x2e, 0x90, 0x6b, 0xa0, 0x46,
	0x32, 0x91, 0x67, 0x24, 0x27, 0xc8, 0x08, 0x88, 0x79, 0x61, 0x5c, 0x7c, 0x60, 0x2f, 0xa6, 0xa4,
	0x16, 0xe4, 0x17, 0x67, 0x96, 0xa4, 0xa6, 0x48, 0x30, 0x93, 0x67, 0x26, 0x2f, 0xc8, 0x18, 0x17,
	0x98, 0x29, 0x70, 0x73, 0xcb, 0x33, 0x4b, 0x32, 0x52, 0x8a, 0x12, 0xcb, 0xf3, 0x24, 0x58, 0x28,
	0x30, 0x37, 0x1c, 0x66, 0x8a, 0x93, 0xe7, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e,
	0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31,
	0x44, 0xe9, 0xa7, 0x67, 0x96, 0xe4, 0x24, 0x42, 0x4c, 0x44, 0x8a, 0xca, 0x8c, 0xfc, 0xa2, 0xbc,
	0xfc, 0x94, 0x54, 0xcc, 0xf8, 0x4d, 0x62, 0x03, 0x47, 0x9a, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff,
	0x60, 0x9e, 0x3e, 0x29, 0x08, 0x02, 0x00, 0x00,
}

func (m *RUNEPool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RUNEPool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RUNEPool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.RuneWithdrawn.Size()
		i -= size
		if _, err := m.RuneWithdrawn.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypeRunePool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.RuneDeposited.Size()
		i -= size
		if _, err := m.RuneDeposited.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypeRunePool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.PoolUnits.Size()
		i -= size
		if _, err := m.PoolUnits.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypeRunePool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.ReserveUnits.Size()
		i -= size
		if _, err := m.ReserveUnits.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypeRunePool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintTypeRunePool(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypeRunePool(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RUNEPool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ReserveUnits.Size()
	n += 1 + l + sovTypeRunePool(uint64(l))
	l = m.PoolUnits.Size()
	n += 1 + l + sovTypeRunePool(uint64(l))
	l = m.RuneDeposited.Size()
	n += 1 + l + sovTypeRunePool(uint64(l))
	l = m.RuneWithdrawn.Size()
	n += 1 + l + sovTypeRunePool(uint64(l))
	return n
}

func sovTypeRunePool(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypeRunePool(x uint64) (n int) {
	return sovTypeRunePool(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RUNEPool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypeRunePool
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
			return fmt.Errorf("proto: RUNEPool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RUNEPool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReserveUnits", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeRunePool
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
				return ErrInvalidLengthTypeRunePool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypeRunePool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ReserveUnits.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolUnits", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeRunePool
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
				return ErrInvalidLengthTypeRunePool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypeRunePool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PoolUnits.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RuneDeposited", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeRunePool
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
				return ErrInvalidLengthTypeRunePool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypeRunePool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RuneDeposited.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RuneWithdrawn", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeRunePool
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
				return ErrInvalidLengthTypeRunePool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypeRunePool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RuneWithdrawn.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypeRunePool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypeRunePool
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypeRunePool
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
func skipTypeRunePool(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypeRunePool
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
					return 0, ErrIntOverflowTypeRunePool
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
					return 0, ErrIntOverflowTypeRunePool
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
				return 0, ErrInvalidLengthTypeRunePool
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypeRunePool
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypeRunePool
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypeRunePool        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypeRunePool          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypeRunePool = fmt.Errorf("proto: unexpected end of group")
)
