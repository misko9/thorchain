// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: types/query_swapper_clout.proto

package types

import (
	fmt "fmt"
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

type QuerySwapperCloutRequest struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *QuerySwapperCloutRequest) Reset()         { *m = QuerySwapperCloutRequest{} }
func (m *QuerySwapperCloutRequest) String() string { return proto.CompactTextString(m) }
func (*QuerySwapperCloutRequest) ProtoMessage()    {}
func (*QuerySwapperCloutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c754ff6fbc74d0df, []int{0}
}
func (m *QuerySwapperCloutRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySwapperCloutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySwapperCloutRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySwapperCloutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySwapperCloutRequest.Merge(m, src)
}
func (m *QuerySwapperCloutRequest) XXX_Size() int {
	return m.Size()
}
func (m *QuerySwapperCloutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySwapperCloutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySwapperCloutRequest proto.InternalMessageInfo

func (m *QuerySwapperCloutRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*QuerySwapperCloutRequest)(nil), "types.QuerySwapperCloutRequest")
}

func init() { proto.RegisterFile("types/query_swapper_clout.proto", fileDescriptor_c754ff6fbc74d0df) }

var fileDescriptor_c754ff6fbc74d0df = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2f, 0xa9, 0x2c, 0x48,
	0x2d, 0xd6, 0x2f, 0x2c, 0x4d, 0x2d, 0xaa, 0x8c, 0x2f, 0x2e, 0x4f, 0x2c, 0x28, 0x48, 0x2d, 0x8a,
	0x4f, 0xce, 0xc9, 0x2f, 0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x2b, 0x90,
	0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x8b, 0xe8, 0x83, 0x58, 0x10, 0x49, 0x25, 0x13, 0x2e, 0x89,
	0x40, 0x90, 0xce, 0x60, 0x88, 0x46, 0x67, 0x90, 0xbe, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12,
	0x21, 0x09, 0x2e, 0xf6, 0xc4, 0x94, 0x94, 0xa2, 0xd4, 0xe2, 0x62, 0x09, 0x46, 0x05, 0x46, 0x0d,
	0xce, 0x20, 0x18, 0xd7, 0xc9, 0xf3, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c,
	0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2,
	0xf4, 0xd3, 0x33, 0x4b, 0x72, 0x12, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x4b, 0x32, 0xf2, 0x8b,
	0x92, 0x33, 0x12, 0x33, 0xf3, 0xc0, 0xac, 0xbc, 0xfc, 0x94, 0x54, 0xfd, 0x0a, 0x64, 0x41, 0x90,
	0xb3, 0x92, 0xd8, 0xc0, 0xee, 0x30, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x25, 0x09, 0xf1, 0x91,
	0xc7, 0x00, 0x00, 0x00,
}

func (m *QuerySwapperCloutRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySwapperCloutRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySwapperCloutRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuerySwapperClout(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuerySwapperClout(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuerySwapperClout(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QuerySwapperCloutRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuerySwapperClout(uint64(l))
	}
	return n
}

func sovQuerySwapperClout(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuerySwapperClout(x uint64) (n int) {
	return sovQuerySwapperClout(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QuerySwapperCloutRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuerySwapperClout
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
			return fmt.Errorf("proto: QuerySwapperCloutRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySwapperCloutRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuerySwapperClout
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
				return ErrInvalidLengthQuerySwapperClout
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuerySwapperClout
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuerySwapperClout(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuerySwapperClout
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
func skipQuerySwapperClout(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuerySwapperClout
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
					return 0, ErrIntOverflowQuerySwapperClout
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
					return 0, ErrIntOverflowQuerySwapperClout
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
				return 0, ErrInvalidLengthQuerySwapperClout
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuerySwapperClout
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuerySwapperClout
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuerySwapperClout        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuerySwapperClout          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuerySwapperClout = fmt.Errorf("proto: unexpected end of group")
)
