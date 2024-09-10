// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package types

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/cosmos/gogoproto/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_MsgNetworkFee                      protoreflect.MessageDescriptor
	fd_MsgNetworkFee_block_height         protoreflect.FieldDescriptor
	fd_MsgNetworkFee_chain                protoreflect.FieldDescriptor
	fd_MsgNetworkFee_transaction_size     protoreflect.FieldDescriptor
	fd_MsgNetworkFee_transaction_fee_rate protoreflect.FieldDescriptor
	fd_MsgNetworkFee_signer               protoreflect.FieldDescriptor
)

func init() {
	file_types_msg_network_fee_proto_init()
	md_MsgNetworkFee = File_types_msg_network_fee_proto.Messages().ByName("MsgNetworkFee")
	fd_MsgNetworkFee_block_height = md_MsgNetworkFee.Fields().ByName("block_height")
	fd_MsgNetworkFee_chain = md_MsgNetworkFee.Fields().ByName("chain")
	fd_MsgNetworkFee_transaction_size = md_MsgNetworkFee.Fields().ByName("transaction_size")
	fd_MsgNetworkFee_transaction_fee_rate = md_MsgNetworkFee.Fields().ByName("transaction_fee_rate")
	fd_MsgNetworkFee_signer = md_MsgNetworkFee.Fields().ByName("signer")
}

var _ protoreflect.Message = (*fastReflection_MsgNetworkFee)(nil)

type fastReflection_MsgNetworkFee MsgNetworkFee

func (x *MsgNetworkFee) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgNetworkFee)(x)
}

func (x *MsgNetworkFee) slowProtoReflect() protoreflect.Message {
	mi := &file_types_msg_network_fee_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgNetworkFee_messageType fastReflection_MsgNetworkFee_messageType
var _ protoreflect.MessageType = fastReflection_MsgNetworkFee_messageType{}

type fastReflection_MsgNetworkFee_messageType struct{}

func (x fastReflection_MsgNetworkFee_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgNetworkFee)(nil)
}
func (x fastReflection_MsgNetworkFee_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgNetworkFee)
}
func (x fastReflection_MsgNetworkFee_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgNetworkFee
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgNetworkFee) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgNetworkFee
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgNetworkFee) Type() protoreflect.MessageType {
	return _fastReflection_MsgNetworkFee_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgNetworkFee) New() protoreflect.Message {
	return new(fastReflection_MsgNetworkFee)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgNetworkFee) Interface() protoreflect.ProtoMessage {
	return (*MsgNetworkFee)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgNetworkFee) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.BlockHeight != int64(0) {
		value := protoreflect.ValueOfInt64(x.BlockHeight)
		if !f(fd_MsgNetworkFee_block_height, value) {
			return
		}
	}
	if x.Chain != "" {
		value := protoreflect.ValueOfString(x.Chain)
		if !f(fd_MsgNetworkFee_chain, value) {
			return
		}
	}
	if x.TransactionSize != uint64(0) {
		value := protoreflect.ValueOfUint64(x.TransactionSize)
		if !f(fd_MsgNetworkFee_transaction_size, value) {
			return
		}
	}
	if x.TransactionFeeRate != uint64(0) {
		value := protoreflect.ValueOfUint64(x.TransactionFeeRate)
		if !f(fd_MsgNetworkFee_transaction_fee_rate, value) {
			return
		}
	}
	if len(x.Signer) != 0 {
		value := protoreflect.ValueOfBytes(x.Signer)
		if !f(fd_MsgNetworkFee_signer, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_MsgNetworkFee) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "types.MsgNetworkFee.block_height":
		return x.BlockHeight != int64(0)
	case "types.MsgNetworkFee.chain":
		return x.Chain != ""
	case "types.MsgNetworkFee.transaction_size":
		return x.TransactionSize != uint64(0)
	case "types.MsgNetworkFee.transaction_fee_rate":
		return x.TransactionFeeRate != uint64(0)
	case "types.MsgNetworkFee.signer":
		return len(x.Signer) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: types.MsgNetworkFee"))
		}
		panic(fmt.Errorf("message types.MsgNetworkFee does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgNetworkFee) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "types.MsgNetworkFee.block_height":
		x.BlockHeight = int64(0)
	case "types.MsgNetworkFee.chain":
		x.Chain = ""
	case "types.MsgNetworkFee.transaction_size":
		x.TransactionSize = uint64(0)
	case "types.MsgNetworkFee.transaction_fee_rate":
		x.TransactionFeeRate = uint64(0)
	case "types.MsgNetworkFee.signer":
		x.Signer = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: types.MsgNetworkFee"))
		}
		panic(fmt.Errorf("message types.MsgNetworkFee does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgNetworkFee) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "types.MsgNetworkFee.block_height":
		value := x.BlockHeight
		return protoreflect.ValueOfInt64(value)
	case "types.MsgNetworkFee.chain":
		value := x.Chain
		return protoreflect.ValueOfString(value)
	case "types.MsgNetworkFee.transaction_size":
		value := x.TransactionSize
		return protoreflect.ValueOfUint64(value)
	case "types.MsgNetworkFee.transaction_fee_rate":
		value := x.TransactionFeeRate
		return protoreflect.ValueOfUint64(value)
	case "types.MsgNetworkFee.signer":
		value := x.Signer
		return protoreflect.ValueOfBytes(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: types.MsgNetworkFee"))
		}
		panic(fmt.Errorf("message types.MsgNetworkFee does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgNetworkFee) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "types.MsgNetworkFee.block_height":
		x.BlockHeight = value.Int()
	case "types.MsgNetworkFee.chain":
		x.Chain = value.Interface().(string)
	case "types.MsgNetworkFee.transaction_size":
		x.TransactionSize = value.Uint()
	case "types.MsgNetworkFee.transaction_fee_rate":
		x.TransactionFeeRate = value.Uint()
	case "types.MsgNetworkFee.signer":
		x.Signer = value.Bytes()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: types.MsgNetworkFee"))
		}
		panic(fmt.Errorf("message types.MsgNetworkFee does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgNetworkFee) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "types.MsgNetworkFee.block_height":
		panic(fmt.Errorf("field block_height of message types.MsgNetworkFee is not mutable"))
	case "types.MsgNetworkFee.chain":
		panic(fmt.Errorf("field chain of message types.MsgNetworkFee is not mutable"))
	case "types.MsgNetworkFee.transaction_size":
		panic(fmt.Errorf("field transaction_size of message types.MsgNetworkFee is not mutable"))
	case "types.MsgNetworkFee.transaction_fee_rate":
		panic(fmt.Errorf("field transaction_fee_rate of message types.MsgNetworkFee is not mutable"))
	case "types.MsgNetworkFee.signer":
		panic(fmt.Errorf("field signer of message types.MsgNetworkFee is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: types.MsgNetworkFee"))
		}
		panic(fmt.Errorf("message types.MsgNetworkFee does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgNetworkFee) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "types.MsgNetworkFee.block_height":
		return protoreflect.ValueOfInt64(int64(0))
	case "types.MsgNetworkFee.chain":
		return protoreflect.ValueOfString("")
	case "types.MsgNetworkFee.transaction_size":
		return protoreflect.ValueOfUint64(uint64(0))
	case "types.MsgNetworkFee.transaction_fee_rate":
		return protoreflect.ValueOfUint64(uint64(0))
	case "types.MsgNetworkFee.signer":
		return protoreflect.ValueOfBytes(nil)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: types.MsgNetworkFee"))
		}
		panic(fmt.Errorf("message types.MsgNetworkFee does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgNetworkFee) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in types.MsgNetworkFee", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgNetworkFee) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgNetworkFee) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_MsgNetworkFee) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgNetworkFee) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgNetworkFee)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.BlockHeight != 0 {
			n += 1 + runtime.Sov(uint64(x.BlockHeight))
		}
		l = len(x.Chain)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.TransactionSize != 0 {
			n += 1 + runtime.Sov(uint64(x.TransactionSize))
		}
		if x.TransactionFeeRate != 0 {
			n += 1 + runtime.Sov(uint64(x.TransactionFeeRate))
		}
		l = len(x.Signer)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgNetworkFee)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Signer) > 0 {
			i -= len(x.Signer)
			copy(dAtA[i:], x.Signer)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Signer)))
			i--
			dAtA[i] = 0x2a
		}
		if x.TransactionFeeRate != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.TransactionFeeRate))
			i--
			dAtA[i] = 0x20
		}
		if x.TransactionSize != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.TransactionSize))
			i--
			dAtA[i] = 0x18
		}
		if len(x.Chain) > 0 {
			i -= len(x.Chain)
			copy(dAtA[i:], x.Chain)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Chain)))
			i--
			dAtA[i] = 0x12
		}
		if x.BlockHeight != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.BlockHeight))
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*MsgNetworkFee)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgNetworkFee: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgNetworkFee: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
				}
				x.BlockHeight = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.BlockHeight |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Chain = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field TransactionSize", wireType)
				}
				x.TransactionSize = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.TransactionSize |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 4:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field TransactionFeeRate", wireType)
				}
				x.TransactionFeeRate = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.TransactionFeeRate |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Signer = append(x.Signer[:0], dAtA[iNdEx:postIndex]...)
				if x.Signer == nil {
					x.Signer = []byte{}
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: types/msg_network_fee.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MsgNetworkFee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockHeight        int64  `protobuf:"varint,1,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	Chain              string `protobuf:"bytes,2,opt,name=chain,proto3" json:"chain,omitempty"`
	TransactionSize    uint64 `protobuf:"varint,3,opt,name=transaction_size,json=transactionSize,proto3" json:"transaction_size,omitempty"`
	TransactionFeeRate uint64 `protobuf:"varint,4,opt,name=transaction_fee_rate,json=transactionFeeRate,proto3" json:"transaction_fee_rate,omitempty"`
	Signer             []byte `protobuf:"bytes,5,opt,name=signer,proto3" json:"signer,omitempty"`
}

func (x *MsgNetworkFee) Reset() {
	*x = MsgNetworkFee{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_msg_network_fee_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgNetworkFee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgNetworkFee) ProtoMessage() {}

// Deprecated: Use MsgNetworkFee.ProtoReflect.Descriptor instead.
func (*MsgNetworkFee) Descriptor() ([]byte, []int) {
	return file_types_msg_network_fee_proto_rawDescGZIP(), []int{0}
}

func (x *MsgNetworkFee) GetBlockHeight() int64 {
	if x != nil {
		return x.BlockHeight
	}
	return 0
}

func (x *MsgNetworkFee) GetChain() string {
	if x != nil {
		return x.Chain
	}
	return ""
}

func (x *MsgNetworkFee) GetTransactionSize() uint64 {
	if x != nil {
		return x.TransactionSize
	}
	return 0
}

func (x *MsgNetworkFee) GetTransactionFeeRate() uint64 {
	if x != nil {
		return x.TransactionFeeRate
	}
	return 0
}

func (x *MsgNetworkFee) GetSigner() []byte {
	if x != nil {
		return x.Signer
	}
	return nil
}

var File_types_msg_network_fee_proto protoreflect.FileDescriptor

var file_types_msg_network_fee_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d, 0x73, 0x67, 0x5f, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x5f, 0x66, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa0, 0x02, 0x0a, 0x0d, 0x4d,
	0x73, 0x67, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x46, 0x65, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12,
	0x44, 0x0a, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2e,
	0xfa, 0xde, 0x1f, 0x2a, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74,
	0x68, 0x6f, 0x72, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x74, 0x68, 0x6f, 0x72, 0x6e, 0x6f, 0x64,
	0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x52, 0x05,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x30, 0x0a, 0x14, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x66, 0x65, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x12,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x65, 0x65, 0x52, 0x61,
	0x74, 0x65, 0x12, 0x49, 0x0a, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0c, 0x42, 0x31, 0xfa, 0xde, 0x1f, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x42, 0x7c, 0x0a,
	0x09, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x42, 0x12, 0x4d, 0x73, 0x67, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x46, 0x65, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x27, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x68, 0x6f,
	0x72, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x74, 0x68, 0x6f, 0x72, 0x6e, 0x6f, 0x64, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x54, 0x58, 0x58, 0xaa,
	0x02, 0x05, 0x54, 0x79, 0x70, 0x65, 0x73, 0xca, 0x02, 0x05, 0x54, 0x79, 0x70, 0x65, 0x73, 0xe2,
	0x02, 0x11, 0x54, 0x79, 0x70, 0x65, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x05, 0x54, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_types_msg_network_fee_proto_rawDescOnce sync.Once
	file_types_msg_network_fee_proto_rawDescData = file_types_msg_network_fee_proto_rawDesc
)

func file_types_msg_network_fee_proto_rawDescGZIP() []byte {
	file_types_msg_network_fee_proto_rawDescOnce.Do(func() {
		file_types_msg_network_fee_proto_rawDescData = protoimpl.X.CompressGZIP(file_types_msg_network_fee_proto_rawDescData)
	})
	return file_types_msg_network_fee_proto_rawDescData
}

var file_types_msg_network_fee_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_types_msg_network_fee_proto_goTypes = []interface{}{
	(*MsgNetworkFee)(nil), // 0: types.MsgNetworkFee
}
var file_types_msg_network_fee_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_types_msg_network_fee_proto_init() }
func file_types_msg_network_fee_proto_init() {
	if File_types_msg_network_fee_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_types_msg_network_fee_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgNetworkFee); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_types_msg_network_fee_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_types_msg_network_fee_proto_goTypes,
		DependencyIndexes: file_types_msg_network_fee_proto_depIdxs,
		MessageInfos:      file_types_msg_network_fee_proto_msgTypes,
	}.Build()
	File_types_msg_network_fee_proto = out.File
	file_types_msg_network_fee_proto_rawDesc = nil
	file_types_msg_network_fee_proto_goTypes = nil
	file_types_msg_network_fee_proto_depIdxs = nil
}
