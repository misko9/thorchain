// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package types

import (
	_ "github.com/cosmos/gogoproto/gogoproto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: types/query.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_types_query_proto protoreflect.FileDescriptor

var file_types_query_proto_rawDesc = []byte{
	0x0a, 0x11, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x70, 0x6f, 0x6f, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x5f, 0x64, 0x65, 0x72, 0x69, 0x76, 0x65, 0x64, 0x5f, 0x70, 0x6f, 0x6f, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x5f, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x5f, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x73, 0x61, 0x76, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x5f, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x74,
	0x72, 0x61, 0x64, 0x65, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0xfd, 0x0a, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x5a, 0x0a, 0x04, 0x50, 0x6f, 0x6f,
	0x6c, 0x12, 0x17, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50,
	0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x74,
	0x68, 0x6f, 0x72, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x7b, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x7d, 0x12, 0x56, 0x0a, 0x05, 0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x12, 0x18,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x6f, 0x6f, 0x6c,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x74, 0x68,
	0x6f, 0x72, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x70, 0x6f, 0x6f, 0x6c, 0x73, 0x12, 0x70, 0x0a,
	0x0b, 0x44, 0x65, 0x72, 0x69, 0x76, 0x65, 0x64, 0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x1e, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x65, 0x72, 0x69, 0x76, 0x65,
	0x64, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x65, 0x72, 0x69, 0x76, 0x65,
	0x64, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f, 0x74, 0x68, 0x6f, 0x72, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x2f, 0x64, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x7b, 0x61, 0x73, 0x73, 0x65, 0x74, 0x7d, 0x12,
	0x6c, 0x0a, 0x0c, 0x44, 0x65, 0x72, 0x69, 0x76, 0x65, 0x64, 0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x12,
	0x1f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x65, 0x72,
	0x69, 0x76, 0x65, 0x64, 0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x20, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x65,
	0x72, 0x69, 0x76, 0x65, 0x64, 0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x74, 0x68, 0x6f,
	0x72, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x64, 0x70, 0x6f, 0x6f, 0x6c, 0x73, 0x12, 0x9e, 0x01,
	0x0a, 0x11, 0x4c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x12, 0x24, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x4c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x3c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x36, 0x12, 0x34, 0x2f, 0x74, 0x68, 0x6f, 0x72, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x2f, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x7b, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x7d, 0x2f, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x2f, 0x7b, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x7d, 0x12, 0x98,
	0x01, 0x0a, 0x12, 0x4c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x73, 0x12, 0x25, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x4c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4c, 0x69, 0x71, 0x75, 0x69, 0x64,
	0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x33, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2d, 0x12, 0x2b, 0x2f, 0x74,
	0x68, 0x6f, 0x72, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x7b, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x7d, 0x2f, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x5f,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x12, 0x6d, 0x0a, 0x05, 0x53, 0x61, 0x76,
	0x65, 0x72, 0x12, 0x18, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x53, 0x61, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x61, 0x76, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x12,
	0x27, 0x2f, 0x74, 0x68, 0x6f, 0x72, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x70, 0x6f, 0x6f, 0x6c,
	0x2f, 0x7b, 0x61, 0x73, 0x73, 0x65, 0x74, 0x7d, 0x2f, 0x73, 0x61, 0x76, 0x65, 0x72, 0x2f, 0x7b,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x7d, 0x12, 0x67, 0x0a, 0x06, 0x53, 0x61, 0x76, 0x65,
	0x72, 0x73, 0x12, 0x19, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x53, 0x61, 0x76, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x61, 0x76, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x20, 0x12, 0x1e, 0x2f, 0x74, 0x68, 0x6f, 0x72, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x70, 0x6f,
	0x6f, 0x6c, 0x2f, 0x7b, 0x61, 0x73, 0x73, 0x65, 0x74, 0x7d, 0x2f, 0x73, 0x61, 0x76, 0x65, 0x72,
	0x73, 0x12, 0x79, 0x0a, 0x08, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x1b, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42, 0x6f, 0x72, 0x72, 0x6f,
	0x77, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x32, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2c,
	0x12, 0x2a, 0x2f, 0x74, 0x68, 0x6f, 0x72, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x70, 0x6f, 0x6f,
	0x6c, 0x2f, 0x7b, 0x61, 0x73, 0x73, 0x65, 0x74, 0x7d, 0x2f, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77,
	0x65, 0x72, 0x2f, 0x7b, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x7d, 0x12, 0x73, 0x0a, 0x09,
	0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x12, 0x1c, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x12, 0x21,
	0x2f, 0x74, 0x68, 0x6f, 0x72, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x70, 0x6f, 0x6f, 0x6c, 0x2f,
	0x7b, 0x61, 0x73, 0x73, 0x65, 0x74, 0x7d, 0x2f, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x72,
	0x73, 0x12, 0x6f, 0x0a, 0x09, 0x54, 0x72, 0x61, 0x64, 0x65, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x1c,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x72, 0x61, 0x64,
	0x65, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x72, 0x61, 0x64, 0x65, 0x55,
	0x6e, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1f, 0x12, 0x1d, 0x2f, 0x74, 0x68, 0x6f, 0x72, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f,
	0x74, 0x72, 0x61, 0x64, 0x65, 0x2f, 0x75, 0x6e, 0x69, 0x74, 0x2f, 0x7b, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x7d, 0x12, 0x6b, 0x0a, 0x0a, 0x54, 0x72, 0x61, 0x64, 0x65, 0x55, 0x6e, 0x69, 0x74, 0x73,
	0x12, 0x1d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x72,
	0x61, 0x64, 0x65, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x72, 0x61,
	0x64, 0x65, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x74, 0x68, 0x6f, 0x72, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x2f, 0x74, 0x72, 0x61, 0x64, 0x65, 0x2f, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x42,
	0x78, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x42, 0x0a, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x6c,
	0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x68, 0x6f, 0x72, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x2f, 0x74, 0x68, 0x6f, 0x72, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x54, 0x58, 0x58, 0xaa, 0x02, 0x05, 0x54, 0x79, 0x70, 0x65,
	0x73, 0xca, 0x02, 0x05, 0x54, 0x79, 0x70, 0x65, 0x73, 0xe2, 0x02, 0x11, 0x54, 0x79, 0x70, 0x65,
	0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x05,
	0x54, 0x79, 0x70, 0x65, 0x73, 0xc8, 0xe2, 0x1e, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var file_types_query_proto_goTypes = []interface{}{
	(*QueryPoolRequest)(nil),                // 0: types.QueryPoolRequest
	(*QueryPoolsRequest)(nil),               // 1: types.QueryPoolsRequest
	(*QueryDerivedPoolRequest)(nil),         // 2: types.QueryDerivedPoolRequest
	(*QueryDerivedPoolsRequest)(nil),        // 3: types.QueryDerivedPoolsRequest
	(*QueryLiquidityProviderRequest)(nil),   // 4: types.QueryLiquidityProviderRequest
	(*QueryLiquidityProvidersRequest)(nil),  // 5: types.QueryLiquidityProvidersRequest
	(*QuerySaverRequest)(nil),               // 6: types.QuerySaverRequest
	(*QuerySaversRequest)(nil),              // 7: types.QuerySaversRequest
	(*QueryBorrowerRequest)(nil),            // 8: types.QueryBorrowerRequest
	(*QueryBorrowersRequest)(nil),           // 9: types.QueryBorrowersRequest
	(*QueryTradeUnitRequest)(nil),           // 10: types.QueryTradeUnitRequest
	(*QueryTradeUnitsRequest)(nil),          // 11: types.QueryTradeUnitsRequest
	(*QueryPoolResponse)(nil),               // 12: types.QueryPoolResponse
	(*QueryPoolsResponse)(nil),              // 13: types.QueryPoolsResponse
	(*QueryDerivedPoolResponse)(nil),        // 14: types.QueryDerivedPoolResponse
	(*QueryDerivedPoolsResponse)(nil),       // 15: types.QueryDerivedPoolsResponse
	(*QueryLiquidityProviderResponse)(nil),  // 16: types.QueryLiquidityProviderResponse
	(*QueryLiquidityProvidersResponse)(nil), // 17: types.QueryLiquidityProvidersResponse
	(*QuerySaverResponse)(nil),              // 18: types.QuerySaverResponse
	(*QuerySaversResponse)(nil),             // 19: types.QuerySaversResponse
	(*QueryBorrowerResponse)(nil),           // 20: types.QueryBorrowerResponse
	(*QueryBorrowersResponse)(nil),          // 21: types.QueryBorrowersResponse
	(*QueryTradeUnitResponse)(nil),          // 22: types.QueryTradeUnitResponse
	(*QueryTradeUnitsResponse)(nil),         // 23: types.QueryTradeUnitsResponse
}
var file_types_query_proto_depIdxs = []int32{
	0,  // 0: types.Query.Pool:input_type -> types.QueryPoolRequest
	1,  // 1: types.Query.Pools:input_type -> types.QueryPoolsRequest
	2,  // 2: types.Query.DerivedPool:input_type -> types.QueryDerivedPoolRequest
	3,  // 3: types.Query.DerivedPools:input_type -> types.QueryDerivedPoolsRequest
	4,  // 4: types.Query.LiquidityProvider:input_type -> types.QueryLiquidityProviderRequest
	5,  // 5: types.Query.LiquidityProviders:input_type -> types.QueryLiquidityProvidersRequest
	6,  // 6: types.Query.Saver:input_type -> types.QuerySaverRequest
	7,  // 7: types.Query.Savers:input_type -> types.QuerySaversRequest
	8,  // 8: types.Query.Borrower:input_type -> types.QueryBorrowerRequest
	9,  // 9: types.Query.Borrowers:input_type -> types.QueryBorrowersRequest
	10, // 10: types.Query.TradeUnit:input_type -> types.QueryTradeUnitRequest
	11, // 11: types.Query.TradeUnits:input_type -> types.QueryTradeUnitsRequest
	12, // 12: types.Query.Pool:output_type -> types.QueryPoolResponse
	13, // 13: types.Query.Pools:output_type -> types.QueryPoolsResponse
	14, // 14: types.Query.DerivedPool:output_type -> types.QueryDerivedPoolResponse
	15, // 15: types.Query.DerivedPools:output_type -> types.QueryDerivedPoolsResponse
	16, // 16: types.Query.LiquidityProvider:output_type -> types.QueryLiquidityProviderResponse
	17, // 17: types.Query.LiquidityProviders:output_type -> types.QueryLiquidityProvidersResponse
	18, // 18: types.Query.Saver:output_type -> types.QuerySaverResponse
	19, // 19: types.Query.Savers:output_type -> types.QuerySaversResponse
	20, // 20: types.Query.Borrower:output_type -> types.QueryBorrowerResponse
	21, // 21: types.Query.Borrowers:output_type -> types.QueryBorrowersResponse
	22, // 22: types.Query.TradeUnit:output_type -> types.QueryTradeUnitResponse
	23, // 23: types.Query.TradeUnits:output_type -> types.QueryTradeUnitsResponse
	12, // [12:24] is the sub-list for method output_type
	0,  // [0:12] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_types_query_proto_init() }
func file_types_query_proto_init() {
	if File_types_query_proto != nil {
		return
	}
	file_types_query_pool_proto_init()
	file_types_query_derived_pool_proto_init()
	file_types_query_liquidity_provider_proto_init()
	file_types_query_saver_proto_init()
	file_types_query_borrower_proto_init()
	file_types_query_trade_unit_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_types_query_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_types_query_proto_goTypes,
		DependencyIndexes: file_types_query_proto_depIdxs,
	}.Build()
	File_types_query_proto = out.File
	file_types_query_proto_rawDesc = nil
	file_types_query_proto_goTypes = nil
	file_types_query_proto_depIdxs = nil
}
