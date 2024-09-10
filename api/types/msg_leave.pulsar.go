// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package types

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/cosmos/gogoproto/gogoproto"
	common "gitlab.com/thorchain/thornode/api/common"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_MsgLeave              protoreflect.MessageDescriptor
	fd_MsgLeave_tx           protoreflect.FieldDescriptor
	fd_MsgLeave_node_address protoreflect.FieldDescriptor
	fd_MsgLeave_signer       protoreflect.FieldDescriptor
)

func init() {
	file_types_msg_leave_proto_init()
	md_MsgLeave = File_types_msg_leave_proto.Messages().ByName("MsgLeave")
	fd_MsgLeave_tx = md_MsgLeave.Fields().ByName("tx")
	fd_MsgLeave_node_address = md_MsgLeave.Fields().ByName("node_address")
	fd_MsgLeave_signer = md_MsgLeave.Fields().ByName("signer")
}

var _ protoreflect.Message = (*fastReflection_MsgLeave)(nil)

type fastReflection_MsgLeave MsgLeave

func (x *MsgLeave) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgLeave)(x)
}

func (x *MsgLeave) slowProtoReflect() protoreflect.Message {
	mi := &file_types_msg_leave_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgLeave_messageType fastReflection_MsgLeave_messageType
var _ protoreflect.MessageType = fastReflection_MsgLeave_messageType{}

type fastReflection_MsgLeave_messageType struct{}

func (x fastReflection_MsgLeave_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgLeave)(nil)
}
func (x fastReflection_MsgLeave_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgLeave)
}
func (x fastReflection_MsgLeave_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgLeave
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgLeave) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgLeave
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgLeave) Type() protoreflect.MessageType {
	return _fastReflection_MsgLeave_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgLeave) New() protoreflect.Message {
	return new(fastReflection_MsgLeave)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgLeave) Interface() protoreflect.ProtoMessage {
	return (*MsgLeave)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgLeave) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Tx != nil {
		value := protoreflect.ValueOfMessage(x.Tx.ProtoReflect())
		if !f(fd_MsgLeave_tx, value) {
			return
		}
	}
	if len(x.NodeAddress) != 0 {
		value := protoreflect.ValueOfBytes(x.NodeAddress)
		if !f(fd_MsgLeave_node_address, value) {
			return
		}
	}
	if len(x.Signer) != 0 {
		value := protoreflect.ValueOfBytes(x.Signer)
		if !f(fd_MsgLeave_signer, value) {
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
func (x *fastReflection_MsgLeave) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "types.MsgLeave.tx":
		return x.Tx != nil
	case "types.MsgLeave.node_address":
		return len(x.NodeAddress) != 0
	case "types.MsgLeave.signer":
		return len(x.Signer) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: types.MsgLeave"))
		}
		panic(fmt.Errorf("message types.MsgLeave does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgLeave) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "types.MsgLeave.tx":
		x.Tx = nil
	case "types.MsgLeave.node_address":
		x.NodeAddress = nil
	case "types.MsgLeave.signer":
		x.Signer = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: types.MsgLeave"))
		}
		panic(fmt.Errorf("message types.MsgLeave does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgLeave) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "types.MsgLeave.tx":
		value := x.Tx
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "types.MsgLeave.node_address":
		value := x.NodeAddress
		return protoreflect.ValueOfBytes(value)
	case "types.MsgLeave.signer":
		value := x.Signer
		return protoreflect.ValueOfBytes(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: types.MsgLeave"))
		}
		panic(fmt.Errorf("message types.MsgLeave does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgLeave) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "types.MsgLeave.tx":
		x.Tx = value.Message().Interface().(*common.Tx)
	case "types.MsgLeave.node_address":
		x.NodeAddress = value.Bytes()
	case "types.MsgLeave.signer":
		x.Signer = value.Bytes()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: types.MsgLeave"))
		}
		panic(fmt.Errorf("message types.MsgLeave does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgLeave) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "types.MsgLeave.tx":
		if x.Tx == nil {
			x.Tx = new(common.Tx)
		}
		return protoreflect.ValueOfMessage(x.Tx.ProtoReflect())
	case "types.MsgLeave.node_address":
		panic(fmt.Errorf("field node_address of message types.MsgLeave is not mutable"))
	case "types.MsgLeave.signer":
		panic(fmt.Errorf("field signer of message types.MsgLeave is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: types.MsgLeave"))
		}
		panic(fmt.Errorf("message types.MsgLeave does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgLeave) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "types.MsgLeave.tx":
		m := new(common.Tx)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "types.MsgLeave.node_address":
		return protoreflect.ValueOfBytes(nil)
	case "types.MsgLeave.signer":
		return protoreflect.ValueOfBytes(nil)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: types.MsgLeave"))
		}
		panic(fmt.Errorf("message types.MsgLeave does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgLeave) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in types.MsgLeave", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgLeave) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgLeave) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgLeave) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgLeave) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgLeave)
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
		if x.Tx != nil {
			l = options.Size(x.Tx)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.NodeAddress)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
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
		x := input.Message.Interface().(*MsgLeave)
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
			dAtA[i] = 0x1a
		}
		if len(x.NodeAddress) > 0 {
			i -= len(x.NodeAddress)
			copy(dAtA[i:], x.NodeAddress)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.NodeAddress)))
			i--
			dAtA[i] = 0x12
		}
		if x.Tx != nil {
			encoded, err := options.Marshal(x.Tx)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
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
		x := input.Message.Interface().(*MsgLeave)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgLeave: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgLeave: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Tx", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Tx == nil {
					x.Tx = &common.Tx{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Tx); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field NodeAddress", wireType)
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
				x.NodeAddress = append(x.NodeAddress[:0], dAtA[iNdEx:postIndex]...)
				if x.NodeAddress == nil {
					x.NodeAddress = []byte{}
				}
				iNdEx = postIndex
			case 3:
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
// source: types/msg_leave.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MsgLeave struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tx          *common.Tx `protobuf:"bytes,1,opt,name=tx,proto3" json:"tx,omitempty"`
	NodeAddress []byte     `protobuf:"bytes,2,opt,name=node_address,json=nodeAddress,proto3" json:"node_address,omitempty"`
	Signer      []byte     `protobuf:"bytes,3,opt,name=signer,proto3" json:"signer,omitempty"`
}

func (x *MsgLeave) Reset() {
	*x = MsgLeave{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_msg_leave_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgLeave) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgLeave) ProtoMessage() {}

// Deprecated: Use MsgLeave.ProtoReflect.Descriptor instead.
func (*MsgLeave) Descriptor() ([]byte, []int) {
	return file_types_msg_leave_proto_rawDescGZIP(), []int{0}
}

func (x *MsgLeave) GetTx() *common.Tx {
	if x != nil {
		return x.Tx
	}
	return nil
}

func (x *MsgLeave) GetNodeAddress() []byte {
	if x != nil {
		return x.NodeAddress
	}
	return nil
}

func (x *MsgLeave) GetSigner() []byte {
	if x != nil {
		return x.Signer
	}
	return nil
}

var File_types_msg_leave_proto protoreflect.FileDescriptor

var file_types_msg_leave_proto_rawDesc = []byte{
	0x0a, 0x15, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d, 0x73, 0x67, 0x5f, 0x6c, 0x65, 0x61, 0x76,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x1a, 0x13,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcd, 0x01, 0x0a, 0x08, 0x4d, 0x73,
	0x67, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x12, 0x20, 0x0a, 0x02, 0x74, 0x78, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x78, 0x42, 0x04,
	0xc8, 0xde, 0x1f, 0x00, 0x52, 0x02, 0x74, 0x78, 0x12, 0x54, 0x0a, 0x0c, 0x6e, 0x6f, 0x64, 0x65,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x31,
	0xfa, 0xde, 0x1f, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x52, 0x0b, 0x6e, 0x6f, 0x64, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x49,
	0x0a, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x31,
	0xfa, 0xde, 0x1f, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x52, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x42, 0x77, 0x0a, 0x09, 0x63, 0x6f, 0x6d,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x42, 0x0d, 0x4d, 0x73, 0x67, 0x4c, 0x65, 0x61, 0x76, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x68, 0x6f, 0x72, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x74, 0x68,
	0x6f, 0x72, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0xa2, 0x02, 0x03, 0x54, 0x58, 0x58, 0xaa, 0x02, 0x05, 0x54, 0x79, 0x70, 0x65, 0x73, 0xca, 0x02,
	0x05, 0x54, 0x79, 0x70, 0x65, 0x73, 0xe2, 0x02, 0x11, 0x54, 0x79, 0x70, 0x65, 0x73, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x05, 0x54, 0x79, 0x70,
	0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_types_msg_leave_proto_rawDescOnce sync.Once
	file_types_msg_leave_proto_rawDescData = file_types_msg_leave_proto_rawDesc
)

func file_types_msg_leave_proto_rawDescGZIP() []byte {
	file_types_msg_leave_proto_rawDescOnce.Do(func() {
		file_types_msg_leave_proto_rawDescData = protoimpl.X.CompressGZIP(file_types_msg_leave_proto_rawDescData)
	})
	return file_types_msg_leave_proto_rawDescData
}

var file_types_msg_leave_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_types_msg_leave_proto_goTypes = []interface{}{
	(*MsgLeave)(nil),  // 0: types.MsgLeave
	(*common.Tx)(nil), // 1: common.Tx
}
var file_types_msg_leave_proto_depIdxs = []int32{
	1, // 0: types.MsgLeave.tx:type_name -> common.Tx
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_types_msg_leave_proto_init() }
func file_types_msg_leave_proto_init() {
	if File_types_msg_leave_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_types_msg_leave_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgLeave); i {
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
			RawDescriptor: file_types_msg_leave_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_types_msg_leave_proto_goTypes,
		DependencyIndexes: file_types_msg_leave_proto_depIdxs,
		MessageInfos:      file_types_msg_leave_proto_msgTypes,
	}.Build()
	File_types_msg_leave_proto = out.File
	file_types_msg_leave_proto_rawDesc = nil
	file_types_msg_leave_proto_goTypes = nil
	file_types_msg_leave_proto_depIdxs = nil
}
