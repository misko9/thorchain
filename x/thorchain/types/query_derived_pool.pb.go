// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: types/query_derived_pool.proto

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

type QueryDerivedPoolRequest struct {
	Asset string `protobuf:"bytes,1,opt,name=asset,proto3" json:"asset,omitempty"`
}

func (m *QueryDerivedPoolRequest) Reset()         { *m = QueryDerivedPoolRequest{} }
func (m *QueryDerivedPoolRequest) String() string { return proto.CompactTextString(m) }
func (*QueryDerivedPoolRequest) ProtoMessage()    {}
func (*QueryDerivedPoolRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00e27877577b0667, []int{0}
}
func (m *QueryDerivedPoolRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDerivedPoolRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDerivedPoolRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDerivedPoolRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDerivedPoolRequest.Merge(m, src)
}
func (m *QueryDerivedPoolRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryDerivedPoolRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDerivedPoolRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDerivedPoolRequest proto.InternalMessageInfo

func (m *QueryDerivedPoolRequest) GetAsset() string {
	if m != nil {
		return m.Asset
	}
	return ""
}

type QueryDerivedPoolResponse struct {
	Asset           string `protobuf:"bytes,1,opt,name=asset,proto3" json:"asset,omitempty"`
	Status          string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Decimals        int64  `protobuf:"varint,3,opt,name=decimals,proto3" json:"decimals,omitempty"`
	BalanceAsset    string `protobuf:"bytes,4,opt,name=balance_asset,json=balanceAsset,proto3" json:"balance_asset,omitempty"`
	BalanceRune     string `protobuf:"bytes,5,opt,name=balance_rune,json=balanceRune,proto3" json:"balance_rune,omitempty"`
	DerivedDepthBps string `protobuf:"bytes,6,opt,name=derived_depth_bps,json=derivedDepthBps,proto3" json:"derived_depth_bps,omitempty"`
}

func (m *QueryDerivedPoolResponse) Reset()         { *m = QueryDerivedPoolResponse{} }
func (m *QueryDerivedPoolResponse) String() string { return proto.CompactTextString(m) }
func (*QueryDerivedPoolResponse) ProtoMessage()    {}
func (*QueryDerivedPoolResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00e27877577b0667, []int{1}
}
func (m *QueryDerivedPoolResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDerivedPoolResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDerivedPoolResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDerivedPoolResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDerivedPoolResponse.Merge(m, src)
}
func (m *QueryDerivedPoolResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryDerivedPoolResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDerivedPoolResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDerivedPoolResponse proto.InternalMessageInfo

func (m *QueryDerivedPoolResponse) GetAsset() string {
	if m != nil {
		return m.Asset
	}
	return ""
}

func (m *QueryDerivedPoolResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *QueryDerivedPoolResponse) GetDecimals() int64 {
	if m != nil {
		return m.Decimals
	}
	return 0
}

func (m *QueryDerivedPoolResponse) GetBalanceAsset() string {
	if m != nil {
		return m.BalanceAsset
	}
	return ""
}

func (m *QueryDerivedPoolResponse) GetBalanceRune() string {
	if m != nil {
		return m.BalanceRune
	}
	return ""
}

func (m *QueryDerivedPoolResponse) GetDerivedDepthBps() string {
	if m != nil {
		return m.DerivedDepthBps
	}
	return ""
}

type QueryDerivedPoolsRequest struct {
}

func (m *QueryDerivedPoolsRequest) Reset()         { *m = QueryDerivedPoolsRequest{} }
func (m *QueryDerivedPoolsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryDerivedPoolsRequest) ProtoMessage()    {}
func (*QueryDerivedPoolsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00e27877577b0667, []int{2}
}
func (m *QueryDerivedPoolsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDerivedPoolsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDerivedPoolsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDerivedPoolsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDerivedPoolsRequest.Merge(m, src)
}
func (m *QueryDerivedPoolsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryDerivedPoolsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDerivedPoolsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDerivedPoolsRequest proto.InternalMessageInfo

type QueryDerivedPoolsResponse struct {
	Pools []*QueryDerivedPoolResponse `protobuf:"bytes,1,rep,name=pools,proto3" json:"pools,omitempty"`
}

func (m *QueryDerivedPoolsResponse) Reset()         { *m = QueryDerivedPoolsResponse{} }
func (m *QueryDerivedPoolsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryDerivedPoolsResponse) ProtoMessage()    {}
func (*QueryDerivedPoolsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00e27877577b0667, []int{3}
}
func (m *QueryDerivedPoolsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDerivedPoolsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDerivedPoolsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDerivedPoolsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDerivedPoolsResponse.Merge(m, src)
}
func (m *QueryDerivedPoolsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryDerivedPoolsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDerivedPoolsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDerivedPoolsResponse proto.InternalMessageInfo

func (m *QueryDerivedPoolsResponse) GetPools() []*QueryDerivedPoolResponse {
	if m != nil {
		return m.Pools
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryDerivedPoolRequest)(nil), "types.QueryDerivedPoolRequest")
	proto.RegisterType((*QueryDerivedPoolResponse)(nil), "types.QueryDerivedPoolResponse")
	proto.RegisterType((*QueryDerivedPoolsRequest)(nil), "types.QueryDerivedPoolsRequest")
	proto.RegisterType((*QueryDerivedPoolsResponse)(nil), "types.QueryDerivedPoolsResponse")
}

func init() { proto.RegisterFile("types/query_derived_pool.proto", fileDescriptor_00e27877577b0667) }

var fileDescriptor_00e27877577b0667 = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x51, 0xbd, 0x4e, 0xc3, 0x30,
	0x10, 0xae, 0x29, 0xa9, 0xc0, 0x05, 0x21, 0xa2, 0x0a, 0x4c, 0x07, 0x53, 0xc2, 0x52, 0x31, 0x24,
	0x12, 0x88, 0x07, 0xa0, 0xea, 0xc2, 0x06, 0x19, 0x59, 0x22, 0x27, 0x39, 0xb5, 0x91, 0xd2, 0xd8,
	0xcd, 0x39, 0x88, 0xbe, 0x05, 0x8f, 0xc5, 0xd8, 0xb1, 0x23, 0x6a, 0x5f, 0x04, 0xc5, 0x71, 0x11,
	0x52, 0xcb, 0x76, 0xdf, 0xdf, 0x59, 0xf7, 0x99, 0x72, 0xbd, 0x50, 0x80, 0xc1, 0xbc, 0x82, 0x72,
	0x11, 0xa5, 0x50, 0x66, 0xef, 0x90, 0x46, 0x4a, 0xca, 0xdc, 0x57, 0xa5, 0xd4, 0xd2, 0x75, 0x8c,
	0xde, 0xef, 0x4d, 0xe4, 0x44, 0x1a, 0x26, 0xa8, 0xa7, 0x46, 0xf4, 0x02, 0x7a, 0xf9, 0x5a, 0x07,
	0xc7, 0x4d, 0xee, 0x45, 0xca, 0x3c, 0x84, 0x79, 0x05, 0xa8, 0xdd, 0x1e, 0x75, 0x04, 0x22, 0x68,
	0x46, 0x06, 0x64, 0x78, 0x1c, 0x36, 0xc0, 0x5b, 0x11, 0xca, 0x76, 0x13, 0xa8, 0x64, 0x81, 0xb0,
	0x3f, 0xe2, 0x5e, 0xd0, 0x0e, 0x6a, 0xa1, 0x2b, 0x64, 0x07, 0x86, 0xb6, 0xc8, 0xed, 0xd3, 0xa3,
	0x14, 0x92, 0x6c, 0x26, 0x72, 0x64, 0xed, 0x01, 0x19, 0xb6, 0xc3, 0x5f, 0xec, 0xde, 0xd2, 0xd3,
	0x58, 0xe4, 0xa2, 0x48, 0x20, 0x6a, 0x36, 0x1e, 0x9a, 0xe8, 0x89, 0x25, 0x9f, 0xcc, 0xe2, 0x1b,
	0xba, 0xc5, 0x51, 0x59, 0x15, 0xc0, 0x1c, 0xe3, 0xe9, 0x5a, 0x2e, 0xac, 0x0a, 0x70, 0xef, 0xe8,
	0xf9, 0xb6, 0x92, 0x14, 0x94, 0x9e, 0x46, 0xb1, 0x42, 0xd6, 0x31, 0xbe, 0x33, 0x2b, 0x8c, 0x6b,
	0x7e, 0xa4, 0xd0, 0xeb, 0xef, 0x5e, 0x86, 0xb6, 0x0c, 0x2f, 0xa4, 0x57, 0x7b, 0x34, 0x7b, 0xf6,
	0x23, 0x75, 0xea, 0xbe, 0x91, 0x91, 0x41, 0x7b, 0xd8, 0xbd, 0xbf, 0xf6, 0x4d, 0xe3, 0xfe, 0x7f,
	0x35, 0x85, 0x8d, 0x7b, 0xf4, 0xfc, 0xb5, 0xe6, 0x64, 0xb9, 0xe6, 0xe4, 0x7b, 0xcd, 0xc9, 0xe7,
	0x86, 0xb7, 0x96, 0x1b, 0xde, 0x5a, 0x6d, 0x78, 0xeb, 0x2d, 0x98, 0x64, 0x3a, 0x17, 0xb1, 0x9f,
	0xc8, 0x59, 0xa0, 0xa7, 0xb2, 0x4c, 0xa6, 0x22, 0x2b, 0xcc, 0x54, 0xc8, 0x14, 0x82, 0x8f, 0xbf,
	0x64, 0xfd, 0x54, 0xdc, 0x31, 0xbf, 0xf9, 0xf0, 0x13, 0x00, 0x00, 0xff, 0xff, 0xee, 0x78, 0x50,
	0xcf, 0x0c, 0x02, 0x00, 0x00,
}

func (m *QueryDerivedPoolRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDerivedPoolRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDerivedPoolRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Asset) > 0 {
		i -= len(m.Asset)
		copy(dAtA[i:], m.Asset)
		i = encodeVarintQueryDerivedPool(dAtA, i, uint64(len(m.Asset)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryDerivedPoolResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDerivedPoolResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDerivedPoolResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DerivedDepthBps) > 0 {
		i -= len(m.DerivedDepthBps)
		copy(dAtA[i:], m.DerivedDepthBps)
		i = encodeVarintQueryDerivedPool(dAtA, i, uint64(len(m.DerivedDepthBps)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.BalanceRune) > 0 {
		i -= len(m.BalanceRune)
		copy(dAtA[i:], m.BalanceRune)
		i = encodeVarintQueryDerivedPool(dAtA, i, uint64(len(m.BalanceRune)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.BalanceAsset) > 0 {
		i -= len(m.BalanceAsset)
		copy(dAtA[i:], m.BalanceAsset)
		i = encodeVarintQueryDerivedPool(dAtA, i, uint64(len(m.BalanceAsset)))
		i--
		dAtA[i] = 0x22
	}
	if m.Decimals != 0 {
		i = encodeVarintQueryDerivedPool(dAtA, i, uint64(m.Decimals))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintQueryDerivedPool(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Asset) > 0 {
		i -= len(m.Asset)
		copy(dAtA[i:], m.Asset)
		i = encodeVarintQueryDerivedPool(dAtA, i, uint64(len(m.Asset)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryDerivedPoolsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDerivedPoolsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDerivedPoolsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryDerivedPoolsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDerivedPoolsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDerivedPoolsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Pools) > 0 {
		for iNdEx := len(m.Pools) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Pools[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQueryDerivedPool(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQueryDerivedPool(dAtA []byte, offset int, v uint64) int {
	offset -= sovQueryDerivedPool(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryDerivedPoolRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Asset)
	if l > 0 {
		n += 1 + l + sovQueryDerivedPool(uint64(l))
	}
	return n
}

func (m *QueryDerivedPoolResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Asset)
	if l > 0 {
		n += 1 + l + sovQueryDerivedPool(uint64(l))
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovQueryDerivedPool(uint64(l))
	}
	if m.Decimals != 0 {
		n += 1 + sovQueryDerivedPool(uint64(m.Decimals))
	}
	l = len(m.BalanceAsset)
	if l > 0 {
		n += 1 + l + sovQueryDerivedPool(uint64(l))
	}
	l = len(m.BalanceRune)
	if l > 0 {
		n += 1 + l + sovQueryDerivedPool(uint64(l))
	}
	l = len(m.DerivedDepthBps)
	if l > 0 {
		n += 1 + l + sovQueryDerivedPool(uint64(l))
	}
	return n
}

func (m *QueryDerivedPoolsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryDerivedPoolsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Pools) > 0 {
		for _, e := range m.Pools {
			l = e.Size()
			n += 1 + l + sovQueryDerivedPool(uint64(l))
		}
	}
	return n
}

func sovQueryDerivedPool(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQueryDerivedPool(x uint64) (n int) {
	return sovQueryDerivedPool(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryDerivedPoolRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryDerivedPool
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
			return fmt.Errorf("proto: QueryDerivedPoolRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDerivedPoolRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Asset", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryDerivedPool
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
				return ErrInvalidLengthQueryDerivedPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryDerivedPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Asset = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryDerivedPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryDerivedPool
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
func (m *QueryDerivedPoolResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryDerivedPool
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
			return fmt.Errorf("proto: QueryDerivedPoolResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDerivedPoolResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Asset", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryDerivedPool
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
				return ErrInvalidLengthQueryDerivedPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryDerivedPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Asset = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryDerivedPool
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
				return ErrInvalidLengthQueryDerivedPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryDerivedPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Decimals", wireType)
			}
			m.Decimals = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryDerivedPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Decimals |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BalanceAsset", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryDerivedPool
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
				return ErrInvalidLengthQueryDerivedPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryDerivedPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BalanceAsset = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BalanceRune", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryDerivedPool
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
				return ErrInvalidLengthQueryDerivedPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryDerivedPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BalanceRune = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DerivedDepthBps", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryDerivedPool
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
				return ErrInvalidLengthQueryDerivedPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryDerivedPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DerivedDepthBps = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryDerivedPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryDerivedPool
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
func (m *QueryDerivedPoolsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryDerivedPool
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
			return fmt.Errorf("proto: QueryDerivedPoolsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDerivedPoolsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQueryDerivedPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryDerivedPool
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
func (m *QueryDerivedPoolsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryDerivedPool
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
			return fmt.Errorf("proto: QueryDerivedPoolsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDerivedPoolsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pools", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryDerivedPool
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
				return ErrInvalidLengthQueryDerivedPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryDerivedPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pools = append(m.Pools, &QueryDerivedPoolResponse{})
			if err := m.Pools[len(m.Pools)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryDerivedPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryDerivedPool
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
func skipQueryDerivedPool(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQueryDerivedPool
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
					return 0, ErrIntOverflowQueryDerivedPool
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
					return 0, ErrIntOverflowQueryDerivedPool
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
				return 0, ErrInvalidLengthQueryDerivedPool
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQueryDerivedPool
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQueryDerivedPool
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQueryDerivedPool        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQueryDerivedPool          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQueryDerivedPool = fmt.Errorf("proto: unexpected end of group")
)
