// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: types/msg_tss_pool.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

type MsgTssPool struct {
	ID              string                                        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PoolPubKey      gitlab_com_thorchain_thornode_common.PubKey   `protobuf:"bytes,2,opt,name=pool_pub_key,json=poolPubKey,proto3,casttype=gitlab.com/thorchain/thornode/common.PubKey" json:"pool_pub_key,omitempty"`
	KeygenType      KeygenType                                    `protobuf:"varint,3,opt,name=keygen_type,json=keygenType,proto3,enum=types.KeygenType,casttype=KeygenType" json:"keygen_type,omitempty"`
	PubKeys         []string                                      `protobuf:"bytes,4,rep,name=pub_keys,json=pubKeys,proto3" json:"pub_keys,omitempty"`
	Height          int64                                         `protobuf:"varint,5,opt,name=height,proto3" json:"height,omitempty"`
	Blame           Blame                                         `protobuf:"bytes,6,opt,name=blame,proto3" json:"blame"`
	Chains          []string                                      `protobuf:"bytes,7,rep,name=chains,proto3" json:"chains,omitempty"`
	Signer          github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,8,opt,name=signer,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"signer,omitempty"`
	KeygenTime      int64                                         `protobuf:"varint,9,opt,name=keygen_time,json=keygenTime,proto3" json:"keygen_time,omitempty"`
	KeysharesBackup []byte                                        `protobuf:"bytes,10,opt,name=keyshares_backup,json=keysharesBackup,proto3" json:"keyshares_backup,omitempty"`
}

func (m *MsgTssPool) Reset()         { *m = MsgTssPool{} }
func (m *MsgTssPool) String() string { return proto.CompactTextString(m) }
func (*MsgTssPool) ProtoMessage()    {}
func (*MsgTssPool) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9c41698893502d1, []int{0}
}
func (m *MsgTssPool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgTssPool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgTssPool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgTssPool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgTssPool.Merge(m, src)
}
func (m *MsgTssPool) XXX_Size() int {
	return m.Size()
}
func (m *MsgTssPool) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgTssPool.DiscardUnknown(m)
}

var xxx_messageInfo_MsgTssPool proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgTssPool)(nil), "types.MsgTssPool")
}

func init() { proto.RegisterFile("types/msg_tss_pool.proto", fileDescriptor_a9c41698893502d1) }

var fileDescriptor_a9c41698893502d1 = []byte{
	// 438 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x4d, 0x6f, 0x9b, 0x40,
	0x10, 0x65, 0xed, 0x98, 0xc4, 0x1b, 0x2b, 0x6d, 0x57, 0x95, 0xbb, 0xcd, 0x01, 0x50, 0x4f, 0x54,
	0x55, 0x40, 0x4d, 0x7f, 0x41, 0x50, 0x2f, 0x51, 0x14, 0x29, 0x45, 0x39, 0xf5, 0x82, 0xf8, 0x58,
	0x2d, 0x2b, 0x1b, 0x16, 0x31, 0x58, 0xaa, 0xff, 0x45, 0xff, 0x55, 0x7d, 0xcc, 0xb1, 0x27, 0xd4,
	0xe2, 0x7f, 0xe1, 0x53, 0xb5, 0xbb, 0x34, 0xc9, 0x29, 0x17, 0x78, 0xf3, 0xde, 0xe8, 0xcd, 0x9b,
	0x01, 0x4c, 0xbb, 0x6d, 0xc3, 0x20, 0xac, 0x80, 0x27, 0x1d, 0x40, 0xd2, 0x48, 0xb9, 0x0e, 0x9a,
	0x56, 0x76, 0x92, 0xcc, 0xb4, 0x72, 0xbe, 0x34, 0x0d, 0xea, 0x99, 0x64, 0xeb, 0xb4, 0x62, 0x46,
	0x3e, 0x7f, 0xf7, 0x8c, 0x5f, 0xb1, 0x2d, 0x67, 0xf5, 0x28, 0xbc, 0xe5, 0x92, 0x4b, 0x0d, 0x43,
	0x85, 0x0c, 0xfb, 0xe1, 0xd7, 0x14, 0xe3, 0x5b, 0xe0, 0xf7, 0x00, 0x77, 0x52, 0xae, 0xc9, 0x12,
	0x4f, 0x44, 0x41, 0x91, 0x87, 0xfc, 0x79, 0x64, 0x0f, 0xbd, 0x3b, 0xb9, 0xfe, 0x1a, 0x4f, 0x44,
	0x41, 0xbe, 0xe1, 0x85, 0x8a, 0x90, 0x34, 0x9b, 0x4c, 0xb9, 0xd2, 0x89, 0xee, 0x08, 0x0f, 0xbd,
	0xfb, 0x89, 0x8b, 0x6e, 0x9d, 0x66, 0x41, 0x2e, 0xab, 0xb0, 0x2b, 0x65, 0x9b, 0x97, 0xa9, 0xa8,
	0x35, 0xaa, 0x65, 0xc1, 0xc2, 0x5c, 0x56, 0x95, 0xac, 0x83, 0xbb, 0x4d, 0x76, 0xc3, 0xb6, 0x31,
	0x56, 0x26, 0x06, 0x93, 0x08, 0x9f, 0x9a, 0x7c, 0x89, 0xca, 0x4a, 0xa7, 0x1e, 0xf2, 0xcf, 0x2e,
	0xdf, 0x04, 0x3a, 0x7e, 0x70, 0xa3, 0x95, 0xfb, 0x6d, 0xc3, 0xa2, 0xb3, 0x43, 0xef, 0xe2, 0xa7,
	0x3a, 0xc6, 0xab, 0x47, 0x4c, 0xde, 0xe3, 0x93, 0x31, 0x11, 0xd0, 0x23, 0x6f, 0xea, 0xcf, 0xe3,
	0xe3, 0x46, 0xbb, 0x03, 0x59, 0x62, 0xbb, 0x64, 0x82, 0x97, 0x1d, 0x9d, 0x79, 0xc8, 0x9f, 0xc6,
	0x63, 0x45, 0x7c, 0x3c, 0xd3, 0xe7, 0xa2, 0xb6, 0x87, 0xfc, 0xd3, 0xcb, 0xc5, 0x38, 0x30, 0x52,
	0x5c, 0x74, 0xb4, 0xeb, 0x5d, 0x2b, 0x36, 0x0d, 0xca, 0x41, 0x2f, 0x03, 0xf4, 0x58, 0x5b, 0x8f,
	0x15, 0xb9, 0xc6, 0x36, 0x08, 0x5e, 0xb3, 0x96, 0x9e, 0x78, 0xc8, 0x5f, 0x44, 0x9f, 0x0f, 0xbd,
	0x7b, 0xc1, 0x45, 0x57, 0x6e, 0xcc, 0x15, 0x72, 0x09, 0x95, 0x84, 0xf1, 0x75, 0x01, 0xc5, 0x2a,
	0x34, 0x23, 0xae, 0xf2, 0xfc, 0xaa, 0x28, 0x5a, 0x06, 0x10, 0x8f, 0x06, 0xc4, 0x7d, 0xba, 0x81,
	0xa8, 0x18, 0x9d, 0xeb, 0xa4, 0xff, 0x17, 0x14, 0x15, 0x23, 0x1f, 0xf1, 0x6b, 0xb5, 0x5c, 0x99,
	0xb6, 0x0c, 0x92, 0x2c, 0xcd, 0x57, 0x9b, 0x86, 0x62, 0x35, 0x35, 0x7e, 0xf5, 0xc8, 0x47, 0x9a,
	0x8e, 0x6e, 0x77, 0x7f, 0x1d, 0x6b, 0x37, 0x38, 0xe8, 0x61, 0x70, 0xd0, 0x9f, 0xc1, 0x41, 0x3f,
	0xf7, 0x8e, 0xf5, 0xb0, 0x77, 0xac, 0xdf, 0x7b, 0xc7, 0xfa, 0x1e, 0xbe, 0xfc, 0x99, 0x7e, 0x3c,
	0x27, 0x55, 0xda, 0xcc, 0xd6, 0xff, 0xc7, 0x97, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xdf, 0xd5,
	0xbf, 0x52, 0x89, 0x02, 0x00, 0x00,
}

func (m *MsgTssPool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgTssPool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgTssPool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.KeysharesBackup) > 0 {
		i -= len(m.KeysharesBackup)
		copy(dAtA[i:], m.KeysharesBackup)
		i = encodeVarintMsgTssPool(dAtA, i, uint64(len(m.KeysharesBackup)))
		i--
		dAtA[i] = 0x52
	}
	if m.KeygenTime != 0 {
		i = encodeVarintMsgTssPool(dAtA, i, uint64(m.KeygenTime))
		i--
		dAtA[i] = 0x48
	}
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintMsgTssPool(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Chains) > 0 {
		for iNdEx := len(m.Chains) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Chains[iNdEx])
			copy(dAtA[i:], m.Chains[iNdEx])
			i = encodeVarintMsgTssPool(dAtA, i, uint64(len(m.Chains[iNdEx])))
			i--
			dAtA[i] = 0x3a
		}
	}
	{
		size, err := m.Blame.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMsgTssPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if m.Height != 0 {
		i = encodeVarintMsgTssPool(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x28
	}
	if len(m.PubKeys) > 0 {
		for iNdEx := len(m.PubKeys) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.PubKeys[iNdEx])
			copy(dAtA[i:], m.PubKeys[iNdEx])
			i = encodeVarintMsgTssPool(dAtA, i, uint64(len(m.PubKeys[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if m.KeygenType != 0 {
		i = encodeVarintMsgTssPool(dAtA, i, uint64(m.KeygenType))
		i--
		dAtA[i] = 0x18
	}
	if len(m.PoolPubKey) > 0 {
		i -= len(m.PoolPubKey)
		copy(dAtA[i:], m.PoolPubKey)
		i = encodeVarintMsgTssPool(dAtA, i, uint64(len(m.PoolPubKey)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ID) > 0 {
		i -= len(m.ID)
		copy(dAtA[i:], m.ID)
		i = encodeVarintMsgTssPool(dAtA, i, uint64(len(m.ID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMsgTssPool(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsgTssPool(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgTssPool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovMsgTssPool(uint64(l))
	}
	l = len(m.PoolPubKey)
	if l > 0 {
		n += 1 + l + sovMsgTssPool(uint64(l))
	}
	if m.KeygenType != 0 {
		n += 1 + sovMsgTssPool(uint64(m.KeygenType))
	}
	if len(m.PubKeys) > 0 {
		for _, s := range m.PubKeys {
			l = len(s)
			n += 1 + l + sovMsgTssPool(uint64(l))
		}
	}
	if m.Height != 0 {
		n += 1 + sovMsgTssPool(uint64(m.Height))
	}
	l = m.Blame.Size()
	n += 1 + l + sovMsgTssPool(uint64(l))
	if len(m.Chains) > 0 {
		for _, s := range m.Chains {
			l = len(s)
			n += 1 + l + sovMsgTssPool(uint64(l))
		}
	}
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovMsgTssPool(uint64(l))
	}
	if m.KeygenTime != 0 {
		n += 1 + sovMsgTssPool(uint64(m.KeygenTime))
	}
	l = len(m.KeysharesBackup)
	if l > 0 {
		n += 1 + l + sovMsgTssPool(uint64(l))
	}
	return n
}

func sovMsgTssPool(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsgTssPool(x uint64) (n int) {
	return sovMsgTssPool(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgTssPool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgTssPool
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
			return fmt.Errorf("proto: MsgTssPool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgTssPool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgTssPool
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
				return ErrInvalidLengthMsgTssPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgTssPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolPubKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgTssPool
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
				return ErrInvalidLengthMsgTssPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgTssPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PoolPubKey = gitlab_com_thorchain_thornode_common.PubKey(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeygenType", wireType)
			}
			m.KeygenType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgTssPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.KeygenType |= KeygenType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKeys", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgTssPool
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
				return ErrInvalidLengthMsgTssPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgTssPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PubKeys = append(m.PubKeys, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgTssPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Blame", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgTssPool
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
				return ErrInvalidLengthMsgTssPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsgTssPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Blame.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chains", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgTssPool
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
				return ErrInvalidLengthMsgTssPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgTssPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chains = append(m.Chains, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgTssPool
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
				return ErrInvalidLengthMsgTssPool
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgTssPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signer = append(m.Signer[:0], dAtA[iNdEx:postIndex]...)
			if m.Signer == nil {
				m.Signer = []byte{}
			}
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeygenTime", wireType)
			}
			m.KeygenTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgTssPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.KeygenTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeysharesBackup", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgTssPool
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
				return ErrInvalidLengthMsgTssPool
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgTssPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KeysharesBackup = append(m.KeysharesBackup[:0], dAtA[iNdEx:postIndex]...)
			if m.KeysharesBackup == nil {
				m.KeysharesBackup = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsgTssPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgTssPool
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
func skipMsgTssPool(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsgTssPool
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
					return 0, ErrIntOverflowMsgTssPool
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
					return 0, ErrIntOverflowMsgTssPool
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
				return 0, ErrInvalidLengthMsgTssPool
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsgTssPool
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsgTssPool
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsgTssPool        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsgTssPool          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsgTssPool = fmt.Errorf("proto: unexpected end of group")
)
