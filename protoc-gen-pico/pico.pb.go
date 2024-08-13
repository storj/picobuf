// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v5.27.3
// source: pico.proto

package main

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MessageOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AlwaysPresent             bool `protobuf:"varint,1,opt,name=always_present,json=alwaysPresent,proto3" json:"always_present,omitempty"`
	CaptureUnrecognizedFields bool `protobuf:"varint,2,opt,name=capture_unrecognized_fields,json=captureUnrecognizedFields,proto3" json:"capture_unrecognized_fields,omitempty"`
}

func (x *MessageOptions) Reset() {
	*x = MessageOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pico_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageOptions) ProtoMessage() {}

func (x *MessageOptions) ProtoReflect() protoreflect.Message {
	mi := &file_pico_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageOptions.ProtoReflect.Descriptor instead.
func (*MessageOptions) Descriptor() ([]byte, []int) {
	return file_pico_proto_rawDescGZIP(), []int{0}
}

func (x *MessageOptions) GetAlwaysPresent() bool {
	if x != nil {
		return x.AlwaysPresent
	}
	return false
}

func (x *MessageOptions) GetCaptureUnrecognizedFields() bool {
	if x != nil {
		return x.CaptureUnrecognizedFields
	}
	return false
}

type FieldOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AlwaysPresent   bool   `protobuf:"varint,1,opt,name=always_present,json=alwaysPresent,proto3" json:"always_present,omitempty"`
	CustomType      string `protobuf:"bytes,2,opt,name=custom_type,json=customType,proto3" json:"custom_type,omitempty"`
	CustomSerialize string `protobuf:"bytes,3,opt,name=custom_serialize,json=customSerialize,proto3" json:"custom_serialize,omitempty"`
}

func (x *FieldOptions) Reset() {
	*x = FieldOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pico_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldOptions) ProtoMessage() {}

func (x *FieldOptions) ProtoReflect() protoreflect.Message {
	mi := &file_pico_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldOptions.ProtoReflect.Descriptor instead.
func (*FieldOptions) Descriptor() ([]byte, []int) {
	return file_pico_proto_rawDescGZIP(), []int{1}
}

func (x *FieldOptions) GetAlwaysPresent() bool {
	if x != nil {
		return x.AlwaysPresent
	}
	return false
}

func (x *FieldOptions) GetCustomType() string {
	if x != nil {
		return x.CustomType
	}
	return ""
}

func (x *FieldOptions) GetCustomSerialize() string {
	if x != nil {
		return x.CustomSerialize
	}
	return ""
}

var file_pico_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*MessageOptions)(nil),
		Field:         28980,
		Name:          "pico.message",
		Tag:           "bytes,28980,opt,name=message",
		Filename:      "pico.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*FieldOptions)(nil),
		Field:         28980,
		Name:          "pico.field",
		Tag:           "bytes,28980,opt,name=field",
		Filename:      "pico.proto",
	},
}

// Extension fields to descriptorpb.MessageOptions.
var (
	// optional pico.MessageOptions message = 28980;
	E_Message = &file_pico_proto_extTypes[0]
)

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional pico.FieldOptions field = 28980;
	E_Field = &file_pico_proto_extTypes[1]
)

var File_pico_proto protoreflect.FileDescriptor

var file_pico_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x70, 0x69, 0x63, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x70, 0x69,
	0x63, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x77, 0x0a, 0x0e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x6c, 0x77, 0x61, 0x79, 0x73,
	0x5f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d,
	0x61, 0x6c, 0x77, 0x61, 0x79, 0x73, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x12, 0x3e, 0x0a,
	0x1b, 0x63, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x75, 0x6e, 0x72, 0x65, 0x63, 0x6f, 0x67,
	0x6e, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x19, 0x63, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x55, 0x6e, 0x72, 0x65, 0x63,
	0x6f, 0x67, 0x6e, 0x69, 0x7a, 0x65, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x22, 0x81, 0x01,
	0x0a, 0x0c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x25,
	0x0a, 0x0e, 0x61, 0x6c, 0x77, 0x61, 0x79, 0x73, 0x5f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x61, 0x6c, 0x77, 0x61, 0x79, 0x73, 0x50, 0x72,
	0x65, 0x73, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x5f, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a,
	0x65, 0x3a, 0x51, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xb4, 0xe2,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x69, 0x63, 0x6f, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x3a, 0x49, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xb4, 0xe2, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x69, 0x63, 0x6f, 0x2e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x42,
	0x17, 0x5a, 0x15, 0x73, 0x74, 0x6f, 0x72, 0x6a, 0x2e, 0x69, 0x6f, 0x2f, 0x70, 0x69, 0x63, 0x6f,
	0x62, 0x75, 0x66, 0x3b, 0x6d, 0x61, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pico_proto_rawDescOnce sync.Once
	file_pico_proto_rawDescData = file_pico_proto_rawDesc
)

func file_pico_proto_rawDescGZIP() []byte {
	file_pico_proto_rawDescOnce.Do(func() {
		file_pico_proto_rawDescData = protoimpl.X.CompressGZIP(file_pico_proto_rawDescData)
	})
	return file_pico_proto_rawDescData
}

var file_pico_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pico_proto_goTypes = []interface{}{
	(*MessageOptions)(nil),              // 0: pico.MessageOptions
	(*FieldOptions)(nil),                // 1: pico.FieldOptions
	(*descriptorpb.MessageOptions)(nil), // 2: google.protobuf.MessageOptions
	(*descriptorpb.FieldOptions)(nil),   // 3: google.protobuf.FieldOptions
}
var file_pico_proto_depIdxs = []int32{
	2, // 0: pico.message:extendee -> google.protobuf.MessageOptions
	3, // 1: pico.field:extendee -> google.protobuf.FieldOptions
	0, // 2: pico.message:type_name -> pico.MessageOptions
	1, // 3: pico.field:type_name -> pico.FieldOptions
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	2, // [2:4] is the sub-list for extension type_name
	0, // [0:2] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pico_proto_init() }
func file_pico_proto_init() {
	if File_pico_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pico_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageOptions); i {
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
		file_pico_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldOptions); i {
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
			RawDescriptor: file_pico_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_pico_proto_goTypes,
		DependencyIndexes: file_pico_proto_depIdxs,
		MessageInfos:      file_pico_proto_msgTypes,
		ExtensionInfos:    file_pico_proto_extTypes,
	}.Build()
	File_pico_proto = out.File
	file_pico_proto_rawDesc = nil
	file_pico_proto_goTypes = nil
	file_pico_proto_depIdxs = nil
}
