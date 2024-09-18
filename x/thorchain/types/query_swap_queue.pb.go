// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: types/query_swap_queue.proto

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

type QuerySwapQueueRequest struct {
}

func (m *QuerySwapQueueRequest) Reset()         { *m = QuerySwapQueueRequest{} }
func (m *QuerySwapQueueRequest) String() string { return proto.CompactTextString(m) }
func (*QuerySwapQueueRequest) ProtoMessage()    {}
func (*QuerySwapQueueRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8755d6524823a050, []int{0}
}
func (m *QuerySwapQueueRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySwapQueueRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySwapQueueRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySwapQueueRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySwapQueueRequest.Merge(m, src)
}
func (m *QuerySwapQueueRequest) XXX_Size() int {
	return m.Size()
}
func (m *QuerySwapQueueRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySwapQueueRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySwapQueueRequest proto.InternalMessageInfo

type QuerySwapQueueResponse struct {
	SwapQueue []*MsgSwap `protobuf:"bytes,1,rep,name=swap_queue,json=swapQueue,proto3" json:"swap_queue,omitempty"`
}

func (m *QuerySwapQueueResponse) Reset()         { *m = QuerySwapQueueResponse{} }
func (m *QuerySwapQueueResponse) String() string { return proto.CompactTextString(m) }
func (*QuerySwapQueueResponse) ProtoMessage()    {}
func (*QuerySwapQueueResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8755d6524823a050, []int{1}
}
func (m *QuerySwapQueueResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySwapQueueResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySwapQueueResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySwapQueueResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySwapQueueResponse.Merge(m, src)
}
func (m *QuerySwapQueueResponse) XXX_Size() int {
	return m.Size()
}
func (m *QuerySwapQueueResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySwapQueueResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySwapQueueResponse proto.InternalMessageInfo

func (m *QuerySwapQueueResponse) GetSwapQueue() []*MsgSwap {
	if m != nil {
		return m.SwapQueue
	}
	return nil
}

func init() {
	proto.RegisterType((*QuerySwapQueueRequest)(nil), "types.QuerySwapQueueRequest")
	proto.RegisterType((*QuerySwapQueueResponse)(nil), "types.QuerySwapQueueResponse")
}

func init() { proto.RegisterFile("types/query_swap_queue.proto", fileDescriptor_8755d6524823a050) }

var fileDescriptor_8755d6524823a050 = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x29, 0xa9, 0x2c, 0x48,
	0x2d, 0xd6, 0x2f, 0x2c, 0x4d, 0x2d, 0xaa, 0x8c, 0x2f, 0x2e, 0x4f, 0x2c, 0x88, 0x2f, 0x2c, 0x4d,
	0x2d, 0x4d, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0xcb, 0x4a, 0x89, 0xa4, 0xe7,
	0xa7, 0xe7, 0x83, 0x45, 0xf4, 0x41, 0x2c, 0x88, 0xa4, 0x94, 0x08, 0x44, 0x6b, 0x6e, 0x71, 0x3a,
	0x58, 0x23, 0x44, 0x54, 0x49, 0x9c, 0x4b, 0x34, 0x10, 0x64, 0x58, 0x70, 0x79, 0x62, 0x41, 0x20,
	0xc8, 0xa8, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x25, 0x77, 0x2e, 0x31, 0x74, 0x89, 0xe2,
	0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x5d, 0x2e, 0x2e, 0x84, 0xcd, 0x12, 0x8c, 0x0a, 0xcc, 0x1a,
	0xdc, 0x46, 0x7c, 0x7a, 0x60, 0xd3, 0xf5, 0x7c, 0x8b, 0xd3, 0x41, 0x1a, 0x82, 0x38, 0x8b, 0x61,
	0xda, 0x9c, 0x3c, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6,
	0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x4a, 0x3f, 0x3d,
	0xb3, 0x24, 0x27, 0x31, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0xbf, 0x24, 0x23, 0xbf, 0x28, 0x39, 0x23,
	0x31, 0x33, 0x0f, 0xcc, 0xca, 0xcb, 0x4f, 0x49, 0xd5, 0xaf, 0x40, 0x16, 0x04, 0x99, 0x9e, 0xc4,
	0x06, 0x76, 0xb3, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x58, 0x53, 0x0d, 0xed, 0x06, 0x01, 0x00,
	0x00,
}

func (m *QuerySwapQueueRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySwapQueueRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySwapQueueRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QuerySwapQueueResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySwapQueueResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySwapQueueResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SwapQueue) > 0 {
		for iNdEx := len(m.SwapQueue) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SwapQueue[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuerySwapQueue(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuerySwapQueue(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuerySwapQueue(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QuerySwapQueueRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QuerySwapQueueResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.SwapQueue) > 0 {
		for _, e := range m.SwapQueue {
			l = e.Size()
			n += 1 + l + sovQuerySwapQueue(uint64(l))
		}
	}
	return n
}

func sovQuerySwapQueue(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuerySwapQueue(x uint64) (n int) {
	return sovQuerySwapQueue(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QuerySwapQueueRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuerySwapQueue
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
			return fmt.Errorf("proto: QuerySwapQueueRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySwapQueueRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuerySwapQueue(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuerySwapQueue
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
func (m *QuerySwapQueueResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuerySwapQueue
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
			return fmt.Errorf("proto: QuerySwapQueueResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySwapQueueResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SwapQueue", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuerySwapQueue
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
				return ErrInvalidLengthQuerySwapQueue
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuerySwapQueue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SwapQueue = append(m.SwapQueue, &MsgSwap{})
			if err := m.SwapQueue[len(m.SwapQueue)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuerySwapQueue(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuerySwapQueue
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
func skipQuerySwapQueue(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuerySwapQueue
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
					return 0, ErrIntOverflowQuerySwapQueue
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
					return 0, ErrIntOverflowQuerySwapQueue
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
				return 0, ErrInvalidLengthQuerySwapQueue
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuerySwapQueue
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuerySwapQueue
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuerySwapQueue        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuerySwapQueue          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuerySwapQueue = fmt.Errorf("proto: unexpected end of group")
)
