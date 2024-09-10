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
	md_RUNEPool                protoreflect.MessageDescriptor
	fd_RUNEPool_reserve_units  protoreflect.FieldDescriptor
	fd_RUNEPool_pool_units     protoreflect.FieldDescriptor
	fd_RUNEPool_rune_deposited protoreflect.FieldDescriptor
	fd_RUNEPool_rune_withdrawn protoreflect.FieldDescriptor
)

func init() {
	file_types_type_rune_pool_proto_init()
	md_RUNEPool = File_types_type_rune_pool_proto.Messages().ByName("RUNEPool")
	fd_RUNEPool_reserve_units = md_RUNEPool.Fields().ByName("reserve_units")
	fd_RUNEPool_pool_units = md_RUNEPool.Fields().ByName("pool_units")
	fd_RUNEPool_rune_deposited = md_RUNEPool.Fields().ByName("rune_deposited")
	fd_RUNEPool_rune_withdrawn = md_RUNEPool.Fields().ByName("rune_withdrawn")
}

var _ protoreflect.Message = (*fastReflection_RUNEPool)(nil)

type fastReflection_RUNEPool RUNEPool

func (x *RUNEPool) ProtoReflect() protoreflect.Message {
	return (*fastReflection_RUNEPool)(x)
}

func (x *RUNEPool) slowProtoReflect() protoreflect.Message {
	mi := &file_types_type_rune_pool_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_RUNEPool_messageType fastReflection_RUNEPool_messageType
var _ protoreflect.MessageType = fastReflection_RUNEPool_messageType{}

type fastReflection_RUNEPool_messageType struct{}

func (x fastReflection_RUNEPool_messageType) Zero() protoreflect.Message {
	return (*fastReflection_RUNEPool)(nil)
}
func (x fastReflection_RUNEPool_messageType) New() protoreflect.Message {
	return new(fastReflection_RUNEPool)
}
func (x fastReflection_RUNEPool_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_RUNEPool
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_RUNEPool) Descriptor() protoreflect.MessageDescriptor {
	return md_RUNEPool
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_RUNEPool) Type() protoreflect.MessageType {
	return _fastReflection_RUNEPool_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_RUNEPool) New() protoreflect.Message {
	return new(fastReflection_RUNEPool)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_RUNEPool) Interface() protoreflect.ProtoMessage {
	return (*RUNEPool)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_RUNEPool) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.ReserveUnits != "" {
		value := protoreflect.ValueOfString(x.ReserveUnits)
		if !f(fd_RUNEPool_reserve_units, value) {
			return
		}
	}
	if x.PoolUnits != "" {
		value := protoreflect.ValueOfString(x.PoolUnits)
		if !f(fd_RUNEPool_pool_units, value) {
			return
		}
	}
	if x.RuneDeposited != "" {
		value := protoreflect.ValueOfString(x.RuneDeposited)
		if !f(fd_RUNEPool_rune_deposited, value) {
			return
		}
	}
	if x.RuneWithdrawn != "" {
		value := protoreflect.ValueOfString(x.RuneWithdrawn)
		if !f(fd_RUNEPool_rune_withdrawn, value) {
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
func (x *fastReflection_RUNEPool) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "types.RUNEPool.reserve_units":
		return x.ReserveUnits != ""
	case "types.RUNEPool.pool_units":
		return x.PoolUnits != ""
	case "types.RUNEPool.rune_deposited":
		return x.RuneDeposited != ""
	case "types.RUNEPool.rune_withdrawn":
		return x.RuneWithdrawn != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: types.RUNEPool"))
		}
		panic(fmt.Errorf("message types.RUNEPool does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_RUNEPool) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "types.RUNEPool.reserve_units":
		x.ReserveUnits = ""
	case "types.RUNEPool.pool_units":
		x.PoolUnits = ""
	case "types.RUNEPool.rune_deposited":
		x.RuneDeposited = ""
	case "types.RUNEPool.rune_withdrawn":
		x.RuneWithdrawn = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: types.RUNEPool"))
		}
		panic(fmt.Errorf("message types.RUNEPool does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_RUNEPool) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "types.RUNEPool.reserve_units":
		value := x.ReserveUnits
		return protoreflect.ValueOfString(value)
	case "types.RUNEPool.pool_units":
		value := x.PoolUnits
		return protoreflect.ValueOfString(value)
	case "types.RUNEPool.rune_deposited":
		value := x.RuneDeposited
		return protoreflect.ValueOfString(value)
	case "types.RUNEPool.rune_withdrawn":
		value := x.RuneWithdrawn
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: types.RUNEPool"))
		}
		panic(fmt.Errorf("message types.RUNEPool does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_RUNEPool) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "types.RUNEPool.reserve_units":
		x.ReserveUnits = value.Interface().(string)
	case "types.RUNEPool.pool_units":
		x.PoolUnits = value.Interface().(string)
	case "types.RUNEPool.rune_deposited":
		x.RuneDeposited = value.Interface().(string)
	case "types.RUNEPool.rune_withdrawn":
		x.RuneWithdrawn = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: types.RUNEPool"))
		}
		panic(fmt.Errorf("message types.RUNEPool does not contain field %s", fd.FullName()))
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
func (x *fastReflection_RUNEPool) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "types.RUNEPool.reserve_units":
		panic(fmt.Errorf("field reserve_units of message types.RUNEPool is not mutable"))
	case "types.RUNEPool.pool_units":
		panic(fmt.Errorf("field pool_units of message types.RUNEPool is not mutable"))
	case "types.RUNEPool.rune_deposited":
		panic(fmt.Errorf("field rune_deposited of message types.RUNEPool is not mutable"))
	case "types.RUNEPool.rune_withdrawn":
		panic(fmt.Errorf("field rune_withdrawn of message types.RUNEPool is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: types.RUNEPool"))
		}
		panic(fmt.Errorf("message types.RUNEPool does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_RUNEPool) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "types.RUNEPool.reserve_units":
		return protoreflect.ValueOfString("")
	case "types.RUNEPool.pool_units":
		return protoreflect.ValueOfString("")
	case "types.RUNEPool.rune_deposited":
		return protoreflect.ValueOfString("")
	case "types.RUNEPool.rune_withdrawn":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: types.RUNEPool"))
		}
		panic(fmt.Errorf("message types.RUNEPool does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_RUNEPool) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in types.RUNEPool", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_RUNEPool) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_RUNEPool) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_RUNEPool) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_RUNEPool) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*RUNEPool)
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
		l = len(x.ReserveUnits)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.PoolUnits)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.RuneDeposited)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.RuneWithdrawn)
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
		x := input.Message.Interface().(*RUNEPool)
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
		if len(x.RuneWithdrawn) > 0 {
			i -= len(x.RuneWithdrawn)
			copy(dAtA[i:], x.RuneWithdrawn)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.RuneWithdrawn)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.RuneDeposited) > 0 {
			i -= len(x.RuneDeposited)
			copy(dAtA[i:], x.RuneDeposited)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.RuneDeposited)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.PoolUnits) > 0 {
			i -= len(x.PoolUnits)
			copy(dAtA[i:], x.PoolUnits)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.PoolUnits)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.ReserveUnits) > 0 {
			i -= len(x.ReserveUnits)
			copy(dAtA[i:], x.ReserveUnits)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ReserveUnits)))
			i--
			dAtA[i] = 0xa
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
		x := input.Message.Interface().(*RUNEPool)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: RUNEPool: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: RUNEPool: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ReserveUnits", wireType)
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
				x.ReserveUnits = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field PoolUnits", wireType)
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
				x.PoolUnits = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field RuneDeposited", wireType)
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
				x.RuneDeposited = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field RuneWithdrawn", wireType)
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
				x.RuneWithdrawn = string(dAtA[iNdEx:postIndex])
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
// source: types/type_rune_pool.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// RUNEPool represents ownership of currently active POL.
type RUNEPool struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReserveUnits  string `protobuf:"bytes,1,opt,name=reserve_units,json=reserveUnits,proto3" json:"reserve_units,omitempty"`
	PoolUnits     string `protobuf:"bytes,2,opt,name=pool_units,json=poolUnits,proto3" json:"pool_units,omitempty"`
	RuneDeposited string `protobuf:"bytes,3,opt,name=rune_deposited,json=runeDeposited,proto3" json:"rune_deposited,omitempty"`
	RuneWithdrawn string `protobuf:"bytes,4,opt,name=rune_withdrawn,json=runeWithdrawn,proto3" json:"rune_withdrawn,omitempty"`
}

func (x *RUNEPool) Reset() {
	*x = RUNEPool{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_type_rune_pool_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RUNEPool) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RUNEPool) ProtoMessage() {}

// Deprecated: Use RUNEPool.ProtoReflect.Descriptor instead.
func (*RUNEPool) Descriptor() ([]byte, []int) {
	return file_types_type_rune_pool_proto_rawDescGZIP(), []int{0}
}

func (x *RUNEPool) GetReserveUnits() string {
	if x != nil {
		return x.ReserveUnits
	}
	return ""
}

func (x *RUNEPool) GetPoolUnits() string {
	if x != nil {
		return x.PoolUnits
	}
	return ""
}

func (x *RUNEPool) GetRuneDeposited() string {
	if x != nil {
		return x.RuneDeposited
	}
	return ""
}

func (x *RUNEPool) GetRuneWithdrawn() string {
	if x != nil {
		return x.RuneWithdrawn
	}
	return ""
}

var File_types_type_rune_pool_proto protoreflect.FileDescriptor

var file_types_type_rune_pool_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x72, 0x75, 0x6e,
	0x65, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x02, 0x0a, 0x08, 0x52, 0x55,
	0x4e, 0x45, 0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x43, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1e, 0xc8,
	0xde, 0x1f, 0x00, 0xda, 0xde, 0x1f, 0x16, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x73, 0x64, 0x6b,
	0x2e, 0x69, 0x6f, 0x2f, 0x6d, 0x61, 0x74, 0x68, 0x2e, 0x55, 0x69, 0x6e, 0x74, 0x52, 0x0c, 0x72,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x12, 0x3d, 0x0a, 0x0a, 0x70,
	0x6f, 0x6f, 0x6c, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x1e, 0xc8, 0xde, 0x1f, 0x00, 0xda, 0xde, 0x1f, 0x16, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x73,
	0x64, 0x6b, 0x2e, 0x69, 0x6f, 0x2f, 0x6d, 0x61, 0x74, 0x68, 0x2e, 0x55, 0x69, 0x6e, 0x74, 0x52,
	0x09, 0x70, 0x6f, 0x6f, 0x6c, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x12, 0x45, 0x0a, 0x0e, 0x72, 0x75,
	0x6e, 0x65, 0x5f, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x1e, 0xc8, 0xde, 0x1f, 0x00, 0xda, 0xde, 0x1f, 0x16, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x73, 0x64, 0x6b, 0x2e, 0x69, 0x6f, 0x2f, 0x6d, 0x61, 0x74, 0x68, 0x2e, 0x55, 0x69,
	0x6e, 0x74, 0x52, 0x0d, 0x72, 0x75, 0x6e, 0x65, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x65,
	0x64, 0x12, 0x45, 0x0a, 0x0e, 0x72, 0x75, 0x6e, 0x65, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72,
	0x61, 0x77, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1e, 0xc8, 0xde, 0x1f, 0x00, 0xda,
	0xde, 0x1f, 0x16, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x73, 0x64, 0x6b, 0x2e, 0x69, 0x6f, 0x2f,
	0x6d, 0x61, 0x74, 0x68, 0x2e, 0x55, 0x69, 0x6e, 0x74, 0x52, 0x0d, 0x72, 0x75, 0x6e, 0x65, 0x57,
	0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x6e, 0x42, 0x7b, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x42, 0x11, 0x54, 0x79, 0x70, 0x65, 0x52, 0x75, 0x6e, 0x65, 0x50,
	0x6f, 0x6f, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x6c,
	0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x68, 0x6f, 0x72, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x2f, 0x74, 0x68, 0x6f, 0x72, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x54, 0x58, 0x58, 0xaa, 0x02, 0x05, 0x54, 0x79, 0x70, 0x65,
	0x73, 0xca, 0x02, 0x05, 0x54, 0x79, 0x70, 0x65, 0x73, 0xe2, 0x02, 0x11, 0x54, 0x79, 0x70, 0x65,
	0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x05,
	0x54, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_types_type_rune_pool_proto_rawDescOnce sync.Once
	file_types_type_rune_pool_proto_rawDescData = file_types_type_rune_pool_proto_rawDesc
)

func file_types_type_rune_pool_proto_rawDescGZIP() []byte {
	file_types_type_rune_pool_proto_rawDescOnce.Do(func() {
		file_types_type_rune_pool_proto_rawDescData = protoimpl.X.CompressGZIP(file_types_type_rune_pool_proto_rawDescData)
	})
	return file_types_type_rune_pool_proto_rawDescData
}

var file_types_type_rune_pool_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_types_type_rune_pool_proto_goTypes = []interface{}{
	(*RUNEPool)(nil), // 0: types.RUNEPool
}
var file_types_type_rune_pool_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_types_type_rune_pool_proto_init() }
func file_types_type_rune_pool_proto_init() {
	if File_types_type_rune_pool_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_types_type_rune_pool_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RUNEPool); i {
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
			RawDescriptor: file_types_type_rune_pool_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_types_type_rune_pool_proto_goTypes,
		DependencyIndexes: file_types_type_rune_pool_proto_depIdxs,
		MessageInfos:      file_types_type_rune_pool_proto_msgTypes,
	}.Build()
	File_types_type_rune_pool_proto = out.File
	file_types_type_rune_pool_proto_rawDesc = nil
	file_types_type_rune_pool_proto_goTypes = nil
	file_types_type_rune_pool_proto_depIdxs = nil
}
