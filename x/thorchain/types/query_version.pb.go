// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: types/query_version.proto

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

type QueryVersionRequest struct {
}

func (m *QueryVersionRequest) Reset()         { *m = QueryVersionRequest{} }
func (m *QueryVersionRequest) String() string { return proto.CompactTextString(m) }
func (*QueryVersionRequest) ProtoMessage()    {}
func (*QueryVersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1349c76c9ec093e8, []int{0}
}
func (m *QueryVersionRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryVersionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryVersionRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryVersionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryVersionRequest.Merge(m, src)
}
func (m *QueryVersionRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryVersionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryVersionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryVersionRequest proto.InternalMessageInfo

type QueryVersionResponse struct {
	Current         string `protobuf:"bytes,1,opt,name=current,proto3" json:"current,omitempty"`
	Next            string `protobuf:"bytes,2,opt,name=next,proto3" json:"next,omitempty"`
	NextSinceHeight int64  `protobuf:"varint,3,opt,name=next_since_height,json=nextSinceHeight,proto3" json:"next_since_height,omitempty"`
	Querier         string `protobuf:"bytes,4,opt,name=querier,proto3" json:"querier,omitempty"`
}

func (m *QueryVersionResponse) Reset()         { *m = QueryVersionResponse{} }
func (m *QueryVersionResponse) String() string { return proto.CompactTextString(m) }
func (*QueryVersionResponse) ProtoMessage()    {}
func (*QueryVersionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1349c76c9ec093e8, []int{1}
}
func (m *QueryVersionResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryVersionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryVersionResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryVersionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryVersionResponse.Merge(m, src)
}
func (m *QueryVersionResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryVersionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryVersionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryVersionResponse proto.InternalMessageInfo

func (m *QueryVersionResponse) GetCurrent() string {
	if m != nil {
		return m.Current
	}
	return ""
}

func (m *QueryVersionResponse) GetNext() string {
	if m != nil {
		return m.Next
	}
	return ""
}

func (m *QueryVersionResponse) GetNextSinceHeight() int64 {
	if m != nil {
		return m.NextSinceHeight
	}
	return 0
}

func (m *QueryVersionResponse) GetQuerier() string {
	if m != nil {
		return m.Querier
	}
	return ""
}

func init() {
	proto.RegisterType((*QueryVersionRequest)(nil), "types.QueryVersionRequest")
	proto.RegisterType((*QueryVersionResponse)(nil), "types.QueryVersionResponse")
}

func init() { proto.RegisterFile("types/query_version.proto", fileDescriptor_1349c76c9ec093e8) }

var fileDescriptor_1349c76c9ec093e8 = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0xa9, 0x2c, 0x48,
	0x2d, 0xd6, 0x2f, 0x2c, 0x4d, 0x2d, 0xaa, 0x8c, 0x2f, 0x4b, 0x2d, 0x2a, 0xce, 0xcc, 0xcf, 0xd3,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x4b, 0x49, 0x89, 0xa4, 0xe7, 0xa7, 0xe7, 0x83,
	0x45, 0xf4, 0x41, 0x2c, 0x88, 0xa4, 0x92, 0x28, 0x97, 0x70, 0x20, 0x48, 0x4f, 0x18, 0x44, 0x4b,
	0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x52, 0x17, 0x23, 0x97, 0x08, 0xaa, 0x78, 0x71, 0x41,
	0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x04, 0x17, 0x7b, 0x72, 0x69, 0x51, 0x51, 0x6a, 0x5e, 0x89, 0x04,
	0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x8c, 0x2b, 0x24, 0xc4, 0xc5, 0x92, 0x97, 0x5a, 0x51, 0x22,
	0xc1, 0x04, 0x16, 0x06, 0xb3, 0x85, 0xb4, 0xb8, 0x04, 0x41, 0x74, 0x7c, 0x71, 0x66, 0x5e, 0x72,
	0x6a, 0x7c, 0x46, 0x6a, 0x66, 0x7a, 0x46, 0x89, 0x04, 0xb3, 0x02, 0xa3, 0x06, 0x73, 0x10, 0x3f,
	0x48, 0x22, 0x18, 0x24, 0xee, 0x01, 0x16, 0x06, 0x99, 0x0c, 0x72, 0x7d, 0x66, 0x6a, 0x91, 0x04,
	0x0b, 0xc4, 0x64, 0x28, 0xd7, 0xc9, 0xf3, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f,
	0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18,
	0xa2, 0xf4, 0xd3, 0x33, 0x4b, 0x72, 0x12, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x4b, 0x32, 0xf2,
	0x8b, 0x92, 0x33, 0x12, 0x33, 0xf3, 0xc0, 0xac, 0xbc, 0xfc, 0x94, 0x54, 0xfd, 0x0a, 0x64, 0x41,
	0x50, 0x20, 0x24, 0xb1, 0x81, 0x7d, 0x6d, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x33, 0xbe, 0x59,
	0x69, 0x2f, 0x01, 0x00, 0x00,
}

func (m *QueryVersionRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryVersionRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryVersionRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryVersionResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryVersionResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryVersionResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Querier) > 0 {
		i -= len(m.Querier)
		copy(dAtA[i:], m.Querier)
		i = encodeVarintQueryVersion(dAtA, i, uint64(len(m.Querier)))
		i--
		dAtA[i] = 0x22
	}
	if m.NextSinceHeight != 0 {
		i = encodeVarintQueryVersion(dAtA, i, uint64(m.NextSinceHeight))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Next) > 0 {
		i -= len(m.Next)
		copy(dAtA[i:], m.Next)
		i = encodeVarintQueryVersion(dAtA, i, uint64(len(m.Next)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Current) > 0 {
		i -= len(m.Current)
		copy(dAtA[i:], m.Current)
		i = encodeVarintQueryVersion(dAtA, i, uint64(len(m.Current)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQueryVersion(dAtA []byte, offset int, v uint64) int {
	offset -= sovQueryVersion(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryVersionRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryVersionResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Current)
	if l > 0 {
		n += 1 + l + sovQueryVersion(uint64(l))
	}
	l = len(m.Next)
	if l > 0 {
		n += 1 + l + sovQueryVersion(uint64(l))
	}
	if m.NextSinceHeight != 0 {
		n += 1 + sovQueryVersion(uint64(m.NextSinceHeight))
	}
	l = len(m.Querier)
	if l > 0 {
		n += 1 + l + sovQueryVersion(uint64(l))
	}
	return n
}

func sovQueryVersion(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQueryVersion(x uint64) (n int) {
	return sovQueryVersion(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryVersionRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryVersion
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
			return fmt.Errorf("proto: QueryVersionRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryVersionRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQueryVersion(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryVersion
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
func (m *QueryVersionResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryVersion
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
			return fmt.Errorf("proto: QueryVersionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryVersionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Current", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryVersion
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
				return ErrInvalidLengthQueryVersion
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryVersion
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Current = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Next", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryVersion
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
				return ErrInvalidLengthQueryVersion
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryVersion
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Next = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextSinceHeight", wireType)
			}
			m.NextSinceHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryVersion
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NextSinceHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Querier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryVersion
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
				return ErrInvalidLengthQueryVersion
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryVersion
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Querier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryVersion(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryVersion
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
func skipQueryVersion(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQueryVersion
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
					return 0, ErrIntOverflowQueryVersion
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
					return 0, ErrIntOverflowQueryVersion
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
				return 0, ErrInvalidLengthQueryVersion
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQueryVersion
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQueryVersion
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQueryVersion        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQueryVersion          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQueryVersion = fmt.Errorf("proto: unexpected end of group")
)
