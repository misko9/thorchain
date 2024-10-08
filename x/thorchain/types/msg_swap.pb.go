// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: types/msg_swap.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

type OrderType int32

const (
	OrderType_market OrderType = 0
	OrderType_limit  OrderType = 1
)

var OrderType_name = map[int32]string{
	0: "market",
	1: "limit",
}

var OrderType_value = map[string]int32{
	"market": 0,
	"limit":  1,
}

func (x OrderType) String() string {
	return proto.EnumName(OrderType_name, int32(x))
}

func (OrderType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a59a5d8aa38a4a7a, []int{0}
}

type MsgSwap struct {
	Tx                      common.Tx                                     `protobuf:"bytes,1,opt,name=tx,proto3" json:"tx"`
	TargetAsset             common.Asset                                  `protobuf:"bytes,2,opt,name=target_asset,json=targetAsset,proto3" json:"target_asset"`
	Destination             gitlab_com_thorchain_thornode_common.Address  `protobuf:"bytes,3,opt,name=destination,proto3,casttype=gitlab.com/thorchain/thornode/common.Address" json:"destination,omitempty"`
	TradeTarget             cosmossdk_io_math.Uint                        `protobuf:"bytes,4,opt,name=trade_target,json=tradeTarget,proto3,customtype=cosmossdk.io/math.Uint" json:"trade_target"`
	AffiliateAddress        gitlab_com_thorchain_thornode_common.Address  `protobuf:"bytes,5,opt,name=affiliate_address,json=affiliateAddress,proto3,casttype=gitlab.com/thorchain/thornode/common.Address" json:"affiliate_address,omitempty"`
	AffiliateBasisPoints    cosmossdk_io_math.Uint                        `protobuf:"bytes,6,opt,name=affiliate_basis_points,json=affiliateBasisPoints,proto3,customtype=cosmossdk.io/math.Uint" json:"affiliate_basis_points"`
	Signer                  github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,7,opt,name=signer,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"signer,omitempty"`
	Aggregator              string                                        `protobuf:"bytes,8,opt,name=aggregator,proto3" json:"aggregator,omitempty"`
	AggregatorTargetAddress string                                        `protobuf:"bytes,9,opt,name=aggregator_target_address,json=aggregatorTargetAddress,proto3" json:"aggregator_target_address,omitempty"`
	AggregatorTargetLimit   *cosmossdk_io_math.Uint                       `protobuf:"bytes,10,opt,name=aggregator_target_limit,json=aggregatorTargetLimit,proto3,customtype=cosmossdk.io/math.Uint" json:"aggregator_target_limit,omitempty"`
	OrderType               OrderType                                     `protobuf:"varint,11,opt,name=order_type,json=orderType,proto3,enum=types.OrderType" json:"order_type,omitempty"`
	StreamQuantity          uint64                                        `protobuf:"varint,12,opt,name=stream_quantity,json=streamQuantity,proto3" json:"stream_quantity,omitempty"`
	StreamInterval          uint64                                        `protobuf:"varint,13,opt,name=stream_interval,json=streamInterval,proto3" json:"stream_interval,omitempty"`
}

func (m *MsgSwap) Reset()         { *m = MsgSwap{} }
func (m *MsgSwap) String() string { return proto.CompactTextString(m) }
func (*MsgSwap) ProtoMessage()    {}
func (*MsgSwap) Descriptor() ([]byte, []int) {
	return fileDescriptor_a59a5d8aa38a4a7a, []int{0}
}
func (m *MsgSwap) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSwap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSwap.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSwap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSwap.Merge(m, src)
}
func (m *MsgSwap) XXX_Size() int {
	return m.Size()
}
func (m *MsgSwap) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSwap.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSwap proto.InternalMessageInfo

func (m *MsgSwap) GetTx() common.Tx {
	if m != nil {
		return m.Tx
	}
	return common.Tx{}
}

func (m *MsgSwap) GetTargetAsset() common.Asset {
	if m != nil {
		return m.TargetAsset
	}
	return common.Asset{}
}

func (m *MsgSwap) GetDestination() gitlab_com_thorchain_thornode_common.Address {
	if m != nil {
		return m.Destination
	}
	return ""
}

func (m *MsgSwap) GetAffiliateAddress() gitlab_com_thorchain_thornode_common.Address {
	if m != nil {
		return m.AffiliateAddress
	}
	return ""
}

func (m *MsgSwap) GetSigner() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Signer
	}
	return nil
}

func (m *MsgSwap) GetAggregator() string {
	if m != nil {
		return m.Aggregator
	}
	return ""
}

func (m *MsgSwap) GetAggregatorTargetAddress() string {
	if m != nil {
		return m.AggregatorTargetAddress
	}
	return ""
}

func (m *MsgSwap) GetOrderType() OrderType {
	if m != nil {
		return m.OrderType
	}
	return OrderType_market
}

func (m *MsgSwap) GetStreamQuantity() uint64 {
	if m != nil {
		return m.StreamQuantity
	}
	return 0
}

func (m *MsgSwap) GetStreamInterval() uint64 {
	if m != nil {
		return m.StreamInterval
	}
	return 0
}

func init() {
	proto.RegisterEnum("types.OrderType", OrderType_name, OrderType_value)
	proto.RegisterType((*MsgSwap)(nil), "types.MsgSwap")
}

func init() { proto.RegisterFile("types/msg_swap.proto", fileDescriptor_a59a5d8aa38a4a7a) }

var fileDescriptor_a59a5d8aa38a4a7a = []byte{
	// 550 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x41, 0x6f, 0xd3, 0x3c,
	0x18, 0x6e, 0xfa, 0xb5, 0xdd, 0x57, 0xb7, 0x1b, 0xc5, 0x94, 0xcd, 0xec, 0x90, 0x46, 0xbb, 0x50,
	0x21, 0x96, 0xc0, 0x90, 0x38, 0x70, 0x6b, 0x6f, 0x95, 0x40, 0x40, 0x28, 0x1c, 0x90, 0x50, 0xe4,
	0x36, 0x5e, 0x6a, 0xb5, 0x89, 0x83, 0xfd, 0x8e, 0xb5, 0xff, 0x82, 0xbf, 0xc2, 0xbf, 0xd8, 0x71,
	0x47, 0xc4, 0xa1, 0x42, 0xed, 0xbf, 0xd8, 0x09, 0xc5, 0x4e, 0xda, 0x0a, 0xa4, 0x09, 0x71, 0xf2,
	0x9b, 0xc7, 0xcf, 0xf3, 0xbc, 0xef, 0xe3, 0xd8, 0xa8, 0x0d, 0x8b, 0x94, 0x29, 0x2f, 0x56, 0x51,
	0xa0, 0x2e, 0x69, 0xea, 0xa6, 0x52, 0x80, 0xc0, 0x55, 0x8d, 0x1e, 0xdf, 0x1b, 0x8b, 0x38, 0x16,
	0x89, 0x67, 0x16, 0xb3, 0x77, 0xdc, 0x8e, 0x44, 0x24, 0x74, 0xe9, 0x65, 0x95, 0x41, 0x4f, 0xbe,
	0xd5, 0xd0, 0xde, 0x2b, 0x15, 0xbd, 0xbb, 0xa4, 0x29, 0x76, 0x50, 0x19, 0xe6, 0xc4, 0x72, 0xac,
	0x6e, 0xe3, 0x0c, 0xb9, 0xb9, 0x78, 0x38, 0xef, 0x57, 0xae, 0x96, 0x9d, 0x92, 0x5f, 0x86, 0x39,
	0x7e, 0x8e, 0x9a, 0x40, 0x65, 0xc4, 0x20, 0xa0, 0x4a, 0x31, 0x20, 0x65, 0xcd, 0xdd, 0x2f, 0xb8,
	0xbd, 0x0c, 0xcc, 0xe9, 0x0d, 0x43, 0xd4, 0x10, 0xf6, 0x51, 0x23, 0x64, 0x0a, 0x78, 0x42, 0x81,
	0x8b, 0x84, 0xfc, 0xe7, 0x58, 0xdd, 0x7a, 0xff, 0xc9, 0xcd, 0xb2, 0xf3, 0x38, 0xe2, 0x30, 0xa3,
	0xa3, 0xcc, 0xc0, 0x83, 0x89, 0x90, 0xe3, 0x09, 0xe5, 0x89, 0xae, 0x12, 0x11, 0xb2, 0x22, 0x40,
	0x2f, 0x0c, 0x25, 0x53, 0xca, 0xdf, 0x35, 0xc1, 0x3d, 0xd4, 0x04, 0x49, 0x43, 0x16, 0x98, 0x46,
	0xa4, 0xa2, 0x4d, 0xed, 0xac, 0xf9, 0x8f, 0x65, 0xe7, 0x70, 0x2c, 0x54, 0x2c, 0x94, 0x0a, 0xa7,
	0x2e, 0x17, 0x5e, 0x4c, 0x61, 0xe2, 0xbe, 0xe7, 0x09, 0xf8, 0x0d, 0xad, 0x19, 0x6a, 0x09, 0xfe,
	0x84, 0xee, 0xd2, 0xf3, 0x73, 0x3e, 0xe3, 0x14, 0x58, 0x40, 0x4d, 0x13, 0x52, 0xfd, 0xc7, 0xe1,
	0x5a, 0x1b, 0xab, 0x1c, 0xc1, 0x43, 0x74, 0xb8, 0xb5, 0x1f, 0x51, 0xc5, 0x55, 0x90, 0x0a, 0x9e,
	0x80, 0x22, 0xb5, 0xbf, 0x9a, 0xb5, 0xbd, 0x51, 0xf7, 0x33, 0xf1, 0x1b, 0xad, 0xc5, 0x03, 0x54,
	0x53, 0x3c, 0x4a, 0x98, 0x24, 0x7b, 0x8e, 0xd5, 0x6d, 0xf6, 0x9f, 0xde, 0x2c, 0x3b, 0xa7, 0x11,
	0x87, 0xc9, 0x85, 0x99, 0xd4, 0x98, 0xe5, 0xcb, 0xa9, 0x0a, 0xa7, 0x9e, 0xbe, 0x14, 0x6e, 0x6f,
	0x3c, 0x2e, 0x46, 0xcd, 0x0d, 0xb0, 0x8d, 0x10, 0x8d, 0x22, 0xc9, 0x22, 0x0a, 0x42, 0x92, 0xff,
	0xb3, 0xa1, 0xfc, 0x1d, 0x04, 0xbf, 0x40, 0x0f, 0xb6, 0x5f, 0x41, 0xf1, 0xe7, 0xf3, 0x73, 0xaa,
	0x6b, 0xfa, 0xd1, 0x96, 0x60, 0x0e, 0xb5, 0x08, 0xff, 0x01, 0x1d, 0xfd, 0xa9, 0x9d, 0xf1, 0x98,
	0x03, 0x41, 0x9b, 0xf4, 0xd6, 0x2d, 0xe9, 0xef, 0xff, 0xee, 0xfc, 0x32, 0x13, 0x63, 0x0f, 0x21,
	0x21, 0x43, 0x26, 0x83, 0x2c, 0x15, 0x69, 0x38, 0x56, 0xf7, 0xe0, 0xac, 0xe5, 0x9a, 0x88, 0xaf,
	0xb3, 0x8d, 0xe1, 0x22, 0x65, 0x7e, 0x5d, 0x14, 0x25, 0x7e, 0x88, 0xee, 0x28, 0x90, 0x8c, 0xc6,
	0xc1, 0xe7, 0x0b, 0x9a, 0x00, 0x87, 0x05, 0x69, 0x3a, 0x56, 0xb7, 0xe2, 0x1f, 0x18, 0xf8, 0x6d,
	0x8e, 0xee, 0x10, 0x79, 0x02, 0x4c, 0x7e, 0xa1, 0x33, 0xb2, 0xbf, 0x4b, 0x1c, 0xe4, 0xe8, 0xa3,
	0x13, 0x54, 0xdf, 0x74, 0xc2, 0x08, 0xd5, 0x62, 0x2a, 0xa7, 0x0c, 0x5a, 0x25, 0x5c, 0x47, 0x55,
	0x9d, 0xb0, 0x65, 0xf5, 0x07, 0x57, 0x2b, 0xdb, 0xba, 0x5e, 0xd9, 0xd6, 0xcf, 0x95, 0x6d, 0x7d,
	0x5d, 0xdb, 0xa5, 0xeb, 0xb5, 0x5d, 0xfa, 0xbe, 0xb6, 0x4b, 0x1f, 0xbd, 0xdb, 0x6f, 0xd5, 0x7c,
	0x17, 0xcc, 0x52, 0x8d, 0x6a, 0xfa, 0xa5, 0x3e, 0xfb, 0x15, 0x00, 0x00, 0xff, 0xff, 0x26, 0x7f,
	0xd9, 0x3b, 0xf3, 0x03, 0x00, 0x00,
}

func (m *MsgSwap) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSwap) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSwap) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.StreamInterval != 0 {
		i = encodeVarintMsgSwap(dAtA, i, uint64(m.StreamInterval))
		i--
		dAtA[i] = 0x68
	}
	if m.StreamQuantity != 0 {
		i = encodeVarintMsgSwap(dAtA, i, uint64(m.StreamQuantity))
		i--
		dAtA[i] = 0x60
	}
	if m.OrderType != 0 {
		i = encodeVarintMsgSwap(dAtA, i, uint64(m.OrderType))
		i--
		dAtA[i] = 0x58
	}
	if m.AggregatorTargetLimit != nil {
		{
			size := m.AggregatorTargetLimit.Size()
			i -= size
			if _, err := m.AggregatorTargetLimit.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintMsgSwap(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x52
	}
	if len(m.AggregatorTargetAddress) > 0 {
		i -= len(m.AggregatorTargetAddress)
		copy(dAtA[i:], m.AggregatorTargetAddress)
		i = encodeVarintMsgSwap(dAtA, i, uint64(len(m.AggregatorTargetAddress)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.Aggregator) > 0 {
		i -= len(m.Aggregator)
		copy(dAtA[i:], m.Aggregator)
		i = encodeVarintMsgSwap(dAtA, i, uint64(len(m.Aggregator)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintMsgSwap(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0x3a
	}
	{
		size := m.AffiliateBasisPoints.Size()
		i -= size
		if _, err := m.AffiliateBasisPoints.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMsgSwap(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if len(m.AffiliateAddress) > 0 {
		i -= len(m.AffiliateAddress)
		copy(dAtA[i:], m.AffiliateAddress)
		i = encodeVarintMsgSwap(dAtA, i, uint64(len(m.AffiliateAddress)))
		i--
		dAtA[i] = 0x2a
	}
	{
		size := m.TradeTarget.Size()
		i -= size
		if _, err := m.TradeTarget.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMsgSwap(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.Destination) > 0 {
		i -= len(m.Destination)
		copy(dAtA[i:], m.Destination)
		i = encodeVarintMsgSwap(dAtA, i, uint64(len(m.Destination)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size, err := m.TargetAsset.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMsgSwap(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Tx.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMsgSwap(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintMsgSwap(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsgSwap(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSwap) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Tx.Size()
	n += 1 + l + sovMsgSwap(uint64(l))
	l = m.TargetAsset.Size()
	n += 1 + l + sovMsgSwap(uint64(l))
	l = len(m.Destination)
	if l > 0 {
		n += 1 + l + sovMsgSwap(uint64(l))
	}
	l = m.TradeTarget.Size()
	n += 1 + l + sovMsgSwap(uint64(l))
	l = len(m.AffiliateAddress)
	if l > 0 {
		n += 1 + l + sovMsgSwap(uint64(l))
	}
	l = m.AffiliateBasisPoints.Size()
	n += 1 + l + sovMsgSwap(uint64(l))
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovMsgSwap(uint64(l))
	}
	l = len(m.Aggregator)
	if l > 0 {
		n += 1 + l + sovMsgSwap(uint64(l))
	}
	l = len(m.AggregatorTargetAddress)
	if l > 0 {
		n += 1 + l + sovMsgSwap(uint64(l))
	}
	if m.AggregatorTargetLimit != nil {
		l = m.AggregatorTargetLimit.Size()
		n += 1 + l + sovMsgSwap(uint64(l))
	}
	if m.OrderType != 0 {
		n += 1 + sovMsgSwap(uint64(m.OrderType))
	}
	if m.StreamQuantity != 0 {
		n += 1 + sovMsgSwap(uint64(m.StreamQuantity))
	}
	if m.StreamInterval != 0 {
		n += 1 + sovMsgSwap(uint64(m.StreamInterval))
	}
	return n
}

func sovMsgSwap(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsgSwap(x uint64) (n int) {
	return sovMsgSwap(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSwap) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgSwap
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
			return fmt.Errorf("proto: MsgSwap: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSwap: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tx", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgSwap
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
				return ErrInvalidLengthMsgSwap
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsgSwap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Tx.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TargetAsset", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgSwap
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
				return ErrInvalidLengthMsgSwap
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsgSwap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TargetAsset.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Destination", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgSwap
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
				return ErrInvalidLengthMsgSwap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgSwap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Destination = gitlab_com_thorchain_thornode_common.Address(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TradeTarget", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgSwap
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
				return ErrInvalidLengthMsgSwap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgSwap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TradeTarget.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AffiliateAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgSwap
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
				return ErrInvalidLengthMsgSwap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgSwap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AffiliateAddress = gitlab_com_thorchain_thornode_common.Address(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AffiliateBasisPoints", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgSwap
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
				return ErrInvalidLengthMsgSwap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgSwap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AffiliateBasisPoints.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgSwap
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
				return ErrInvalidLengthMsgSwap
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgSwap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signer = append(m.Signer[:0], dAtA[iNdEx:postIndex]...)
			if m.Signer == nil {
				m.Signer = []byte{}
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Aggregator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgSwap
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
				return ErrInvalidLengthMsgSwap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgSwap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Aggregator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AggregatorTargetAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgSwap
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
				return ErrInvalidLengthMsgSwap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgSwap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AggregatorTargetAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AggregatorTargetLimit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgSwap
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
				return ErrInvalidLengthMsgSwap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgSwap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v cosmossdk_io_math.Uint
			m.AggregatorTargetLimit = &v
			if err := m.AggregatorTargetLimit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderType", wireType)
			}
			m.OrderType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgSwap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OrderType |= OrderType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StreamQuantity", wireType)
			}
			m.StreamQuantity = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgSwap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StreamQuantity |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StreamInterval", wireType)
			}
			m.StreamInterval = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgSwap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StreamInterval |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMsgSwap(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgSwap
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
func skipMsgSwap(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsgSwap
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
					return 0, ErrIntOverflowMsgSwap
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
					return 0, ErrIntOverflowMsgSwap
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
				return 0, ErrInvalidLengthMsgSwap
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsgSwap
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsgSwap
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsgSwap        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsgSwap          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsgSwap = fmt.Errorf("proto: unexpected end of group")
)
