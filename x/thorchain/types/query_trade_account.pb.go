// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: types/query_trade_account.proto

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

type QueryTradeAccountRequest struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *QueryTradeAccountRequest) Reset()         { *m = QueryTradeAccountRequest{} }
func (m *QueryTradeAccountRequest) String() string { return proto.CompactTextString(m) }
func (*QueryTradeAccountRequest) ProtoMessage()    {}
func (*QueryTradeAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0892fac21581ae0, []int{0}
}
func (m *QueryTradeAccountRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryTradeAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryTradeAccountRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryTradeAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryTradeAccountRequest.Merge(m, src)
}
func (m *QueryTradeAccountRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryTradeAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryTradeAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryTradeAccountRequest proto.InternalMessageInfo

func (m *QueryTradeAccountRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type QueryTradeAccountResponse struct {
	Asset              string `protobuf:"bytes,1,opt,name=asset,proto3" json:"asset,omitempty"`
	Units              string `protobuf:"bytes,2,opt,name=units,proto3" json:"units,omitempty"`
	Owner              string `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	LastAddHeight      int64  `protobuf:"varint,4,opt,name=last_add_height,json=lastAddHeight,proto3" json:"last_add_height,omitempty"`
	LastWithdrawHeight int64  `protobuf:"varint,5,opt,name=last_withdraw_height,json=lastWithdrawHeight,proto3" json:"last_withdraw_height,omitempty"`
}

func (m *QueryTradeAccountResponse) Reset()         { *m = QueryTradeAccountResponse{} }
func (m *QueryTradeAccountResponse) String() string { return proto.CompactTextString(m) }
func (*QueryTradeAccountResponse) ProtoMessage()    {}
func (*QueryTradeAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0892fac21581ae0, []int{1}
}
func (m *QueryTradeAccountResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryTradeAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryTradeAccountResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryTradeAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryTradeAccountResponse.Merge(m, src)
}
func (m *QueryTradeAccountResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryTradeAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryTradeAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryTradeAccountResponse proto.InternalMessageInfo

func (m *QueryTradeAccountResponse) GetAsset() string {
	if m != nil {
		return m.Asset
	}
	return ""
}

func (m *QueryTradeAccountResponse) GetUnits() string {
	if m != nil {
		return m.Units
	}
	return ""
}

func (m *QueryTradeAccountResponse) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *QueryTradeAccountResponse) GetLastAddHeight() int64 {
	if m != nil {
		return m.LastAddHeight
	}
	return 0
}

func (m *QueryTradeAccountResponse) GetLastWithdrawHeight() int64 {
	if m != nil {
		return m.LastWithdrawHeight
	}
	return 0
}

type QueryTradeAccountsRequest struct {
	Asset string `protobuf:"bytes,1,opt,name=asset,proto3" json:"asset,omitempty"`
}

func (m *QueryTradeAccountsRequest) Reset()         { *m = QueryTradeAccountsRequest{} }
func (m *QueryTradeAccountsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryTradeAccountsRequest) ProtoMessage()    {}
func (*QueryTradeAccountsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0892fac21581ae0, []int{2}
}
func (m *QueryTradeAccountsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryTradeAccountsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryTradeAccountsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryTradeAccountsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryTradeAccountsRequest.Merge(m, src)
}
func (m *QueryTradeAccountsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryTradeAccountsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryTradeAccountsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryTradeAccountsRequest proto.InternalMessageInfo

func (m *QueryTradeAccountsRequest) GetAsset() string {
	if m != nil {
		return m.Asset
	}
	return ""
}

type QueryTradeAccountsResponse struct {
	TradeAccounts []*QueryTradeAccountResponse `protobuf:"bytes,1,rep,name=trade_accounts,json=tradeAccounts,proto3" json:"trade_accounts,omitempty"`
}

func (m *QueryTradeAccountsResponse) Reset()         { *m = QueryTradeAccountsResponse{} }
func (m *QueryTradeAccountsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryTradeAccountsResponse) ProtoMessage()    {}
func (*QueryTradeAccountsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0892fac21581ae0, []int{3}
}
func (m *QueryTradeAccountsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryTradeAccountsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryTradeAccountsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryTradeAccountsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryTradeAccountsResponse.Merge(m, src)
}
func (m *QueryTradeAccountsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryTradeAccountsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryTradeAccountsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryTradeAccountsResponse proto.InternalMessageInfo

func (m *QueryTradeAccountsResponse) GetTradeAccounts() []*QueryTradeAccountResponse {
	if m != nil {
		return m.TradeAccounts
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryTradeAccountRequest)(nil), "types.QueryTradeAccountRequest")
	proto.RegisterType((*QueryTradeAccountResponse)(nil), "types.QueryTradeAccountResponse")
	proto.RegisterType((*QueryTradeAccountsRequest)(nil), "types.QueryTradeAccountsRequest")
	proto.RegisterType((*QueryTradeAccountsResponse)(nil), "types.QueryTradeAccountsResponse")
}

func init() { proto.RegisterFile("types/query_trade_account.proto", fileDescriptor_c0892fac21581ae0) }

var fileDescriptor_c0892fac21581ae0 = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x85, 0x6b, 0x4a, 0x41, 0x18, 0x15, 0xa4, 0xa8, 0x43, 0xe8, 0x60, 0xaa, 0x0c, 0xa8, 0x53,
	0xc2, 0xdf, 0x0b, 0x94, 0x05, 0x18, 0x89, 0x90, 0x90, 0x58, 0x22, 0x37, 0xbe, 0x4a, 0x22, 0x15,
	0xbb, 0xb5, 0x6f, 0x54, 0xfa, 0x16, 0xbc, 0x0d, 0xaf, 0xc0, 0xd8, 0x91, 0x11, 0xb5, 0x2f, 0x82,
	0x6c, 0xb7, 0x52, 0x91, 0xc2, 0xe6, 0xf3, 0xdd, 0x73, 0x12, 0xdf, 0x63, 0x7a, 0x8e, 0x8b, 0x29,
	0x98, 0x64, 0x56, 0x83, 0x5e, 0x64, 0xa8, 0xb9, 0x80, 0x8c, 0xe7, 0xb9, 0xaa, 0x25, 0xc6, 0x53,
	0xad, 0x50, 0x05, 0x1d, 0x67, 0xe8, 0xf7, 0x0a, 0x55, 0x28, 0x47, 0x12, 0x7b, 0xf2, 0xc3, 0xe8,
	0x96, 0x86, 0x4f, 0x36, 0xf9, 0x6c, 0x83, 0x23, 0x9f, 0x4b, 0x61, 0x56, 0x83, 0xc1, 0x20, 0xa4,
	0x87, 0x5c, 0x08, 0x0d, 0xc6, 0x84, 0x64, 0x40, 0x86, 0x47, 0xe9, 0x56, 0x46, 0x9f, 0x84, 0x9e,
	0x35, 0xc4, 0xcc, 0x54, 0x49, 0x03, 0x41, 0x8f, 0x76, 0xb8, 0x31, 0x80, 0x9b, 0x94, 0x17, 0x96,
	0xd6, 0xb2, 0x42, 0x13, 0xee, 0x79, 0xea, 0x84, 0xa5, 0x6a, 0x2e, 0x41, 0x87, 0x6d, 0x4f, 0x9d,
	0x08, 0x2e, 0xe8, 0xe9, 0x84, 0x1b, 0xcc, 0xb8, 0x10, 0x59, 0x09, 0x55, 0x51, 0x62, 0xb8, 0x3f,
	0x20, 0xc3, 0x76, 0xda, 0xb5, 0x78, 0x24, 0xc4, 0x83, 0x83, 0xc1, 0x25, 0xed, 0x39, 0xdf, 0xbc,
	0xc2, 0x52, 0x68, 0x3e, 0xdf, 0x9a, 0x3b, 0xce, 0x1c, 0xd8, 0xd9, 0xcb, 0x66, 0xe4, 0x13, 0xd1,
	0x55, 0xc3, 0xc5, 0xcd, 0x76, 0xe1, 0xc6, 0x8b, 0x47, 0x40, 0xfb, 0x4d, 0x91, 0xcd, 0xb2, 0xf7,
	0xf4, 0xe4, 0x4f, 0xe9, 0xb6, 0xab, 0xf6, 0xf0, 0xf8, 0x7a, 0x10, 0xbb, 0xda, 0xe3, 0x7f, 0x6b,
	0x4a, 0xbb, 0xb8, 0xfb, 0xc1, 0xbb, 0xc7, 0xaf, 0x15, 0x23, 0xcb, 0x15, 0x23, 0x3f, 0x2b, 0x46,
	0x3e, 0xd6, 0xac, 0xb5, 0x5c, 0xb3, 0xd6, 0xf7, 0x9a, 0xb5, 0x5e, 0x93, 0xa2, 0xc2, 0x09, 0x1f,
	0xc7, 0xb9, 0x7a, 0x4b, 0xb0, 0x54, 0x3a, 0x2f, 0x79, 0x25, 0xdd, 0x49, 0x2a, 0x01, 0xc9, 0xfb,
	0x2e, 0xb4, 0xff, 0x1c, 0x1f, 0xb8, 0xb7, 0xbd, 0xf9, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x18, 0xf9,
	0xd7, 0x0c, 0x1b, 0x02, 0x00, 0x00,
}

func (m *QueryTradeAccountRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryTradeAccountRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryTradeAccountRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQueryTradeAccount(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryTradeAccountResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryTradeAccountResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryTradeAccountResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LastWithdrawHeight != 0 {
		i = encodeVarintQueryTradeAccount(dAtA, i, uint64(m.LastWithdrawHeight))
		i--
		dAtA[i] = 0x28
	}
	if m.LastAddHeight != 0 {
		i = encodeVarintQueryTradeAccount(dAtA, i, uint64(m.LastAddHeight))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintQueryTradeAccount(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Units) > 0 {
		i -= len(m.Units)
		copy(dAtA[i:], m.Units)
		i = encodeVarintQueryTradeAccount(dAtA, i, uint64(len(m.Units)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Asset) > 0 {
		i -= len(m.Asset)
		copy(dAtA[i:], m.Asset)
		i = encodeVarintQueryTradeAccount(dAtA, i, uint64(len(m.Asset)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryTradeAccountsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryTradeAccountsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryTradeAccountsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Asset) > 0 {
		i -= len(m.Asset)
		copy(dAtA[i:], m.Asset)
		i = encodeVarintQueryTradeAccount(dAtA, i, uint64(len(m.Asset)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryTradeAccountsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryTradeAccountsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryTradeAccountsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TradeAccounts) > 0 {
		for iNdEx := len(m.TradeAccounts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TradeAccounts[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQueryTradeAccount(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQueryTradeAccount(dAtA []byte, offset int, v uint64) int {
	offset -= sovQueryTradeAccount(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryTradeAccountRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQueryTradeAccount(uint64(l))
	}
	return n
}

func (m *QueryTradeAccountResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Asset)
	if l > 0 {
		n += 1 + l + sovQueryTradeAccount(uint64(l))
	}
	l = len(m.Units)
	if l > 0 {
		n += 1 + l + sovQueryTradeAccount(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovQueryTradeAccount(uint64(l))
	}
	if m.LastAddHeight != 0 {
		n += 1 + sovQueryTradeAccount(uint64(m.LastAddHeight))
	}
	if m.LastWithdrawHeight != 0 {
		n += 1 + sovQueryTradeAccount(uint64(m.LastWithdrawHeight))
	}
	return n
}

func (m *QueryTradeAccountsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Asset)
	if l > 0 {
		n += 1 + l + sovQueryTradeAccount(uint64(l))
	}
	return n
}

func (m *QueryTradeAccountsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.TradeAccounts) > 0 {
		for _, e := range m.TradeAccounts {
			l = e.Size()
			n += 1 + l + sovQueryTradeAccount(uint64(l))
		}
	}
	return n
}

func sovQueryTradeAccount(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQueryTradeAccount(x uint64) (n int) {
	return sovQueryTradeAccount(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryTradeAccountRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryTradeAccount
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
			return fmt.Errorf("proto: QueryTradeAccountRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryTradeAccountRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryTradeAccount
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
				return ErrInvalidLengthQueryTradeAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryTradeAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryTradeAccount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryTradeAccount
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
func (m *QueryTradeAccountResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryTradeAccount
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
			return fmt.Errorf("proto: QueryTradeAccountResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryTradeAccountResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Asset", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryTradeAccount
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
				return ErrInvalidLengthQueryTradeAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryTradeAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Asset = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Units", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryTradeAccount
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
				return ErrInvalidLengthQueryTradeAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryTradeAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Units = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryTradeAccount
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
				return ErrInvalidLengthQueryTradeAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryTradeAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastAddHeight", wireType)
			}
			m.LastAddHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryTradeAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastAddHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastWithdrawHeight", wireType)
			}
			m.LastWithdrawHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryTradeAccount
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
			skippy, err := skipQueryTradeAccount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryTradeAccount
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
func (m *QueryTradeAccountsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryTradeAccount
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
			return fmt.Errorf("proto: QueryTradeAccountsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryTradeAccountsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Asset", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryTradeAccount
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
				return ErrInvalidLengthQueryTradeAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryTradeAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Asset = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryTradeAccount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryTradeAccount
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
func (m *QueryTradeAccountsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryTradeAccount
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
			return fmt.Errorf("proto: QueryTradeAccountsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryTradeAccountsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TradeAccounts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryTradeAccount
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
				return ErrInvalidLengthQueryTradeAccount
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryTradeAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TradeAccounts = append(m.TradeAccounts, &QueryTradeAccountResponse{})
			if err := m.TradeAccounts[len(m.TradeAccounts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryTradeAccount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryTradeAccount
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
func skipQueryTradeAccount(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQueryTradeAccount
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
					return 0, ErrIntOverflowQueryTradeAccount
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
					return 0, ErrIntOverflowQueryTradeAccount
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
				return 0, ErrInvalidLengthQueryTradeAccount
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQueryTradeAccount
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQueryTradeAccount
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQueryTradeAccount        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQueryTradeAccount          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQueryTradeAccount = fmt.Errorf("proto: unexpected end of group")
)
