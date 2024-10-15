// Copyright The HTNN Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: api/tests/integration/config.proto

package integration

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NeedBuffer bool   `protobuf:"varint,1,opt,name=need_buffer,json=needBuffer,proto3" json:"need_buffer,omitempty"`
	Decode     bool   `protobuf:"varint,2,opt,name=decode,proto3" json:"decode,omitempty"`
	Encode     bool   `protobuf:"varint,3,opt,name=encode,proto3" json:"encode,omitempty"`
	Headers    bool   `protobuf:"varint,4,opt,name=headers,proto3" json:"headers,omitempty"`
	Data       bool   `protobuf:"varint,5,opt,name=data,proto3" json:"data,omitempty"`
	Trailers   bool   `protobuf:"varint,6,opt,name=trailers,proto3" json:"trailers,omitempty"`
	ReplyMsg   string `protobuf:"bytes,7,opt,name=reply_msg,json=replyMsg,proto3" json:"reply_msg,omitempty"`
	InGrpcMode bool   `protobuf:"varint,8,opt,name=in_grpc_mode,json=inGrpcMode,proto3" json:"in_grpc_mode,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_tests_integration_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_api_tests_integration_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_api_tests_integration_config_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetNeedBuffer() bool {
	if x != nil {
		return x.NeedBuffer
	}
	return false
}

func (x *Config) GetDecode() bool {
	if x != nil {
		return x.Decode
	}
	return false
}

func (x *Config) GetEncode() bool {
	if x != nil {
		return x.Encode
	}
	return false
}

func (x *Config) GetHeaders() bool {
	if x != nil {
		return x.Headers
	}
	return false
}

func (x *Config) GetData() bool {
	if x != nil {
		return x.Data
	}
	return false
}

func (x *Config) GetTrailers() bool {
	if x != nil {
		return x.Trailers
	}
	return false
}

func (x *Config) GetReplyMsg() string {
	if x != nil {
		return x.ReplyMsg
	}
	return ""
}

func (x *Config) GetInGrpcMode() bool {
	if x != nil {
		return x.InGrpcMode
	}
	return false
}

type BadPluginConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PanicInFactory bool `protobuf:"varint,1,opt,name=panic_in_factory,json=panicInFactory,proto3" json:"panic_in_factory,omitempty"`
	PanicInParse   bool `protobuf:"varint,2,opt,name=panic_in_parse,json=panicInParse,proto3" json:"panic_in_parse,omitempty"`
	ErrorInInit    bool `protobuf:"varint,3,opt,name=error_in_init,json=errorInInit,proto3" json:"error_in_init,omitempty"`
	PanicInInit    bool `protobuf:"varint,4,opt,name=panic_in_init,json=panicInInit,proto3" json:"panic_in_init,omitempty"`
}

func (x *BadPluginConfig) Reset() {
	*x = BadPluginConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_tests_integration_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BadPluginConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BadPluginConfig) ProtoMessage() {}

func (x *BadPluginConfig) ProtoReflect() protoreflect.Message {
	mi := &file_api_tests_integration_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BadPluginConfig.ProtoReflect.Descriptor instead.
func (*BadPluginConfig) Descriptor() ([]byte, []int) {
	return file_api_tests_integration_config_proto_rawDescGZIP(), []int{1}
}

func (x *BadPluginConfig) GetPanicInFactory() bool {
	if x != nil {
		return x.PanicInFactory
	}
	return false
}

func (x *BadPluginConfig) GetPanicInParse() bool {
	if x != nil {
		return x.PanicInParse
	}
	return false
}

func (x *BadPluginConfig) GetErrorInInit() bool {
	if x != nil {
		return x.ErrorInInit
	}
	return false
}

func (x *BadPluginConfig) GetPanicInInit() bool {
	if x != nil {
		return x.PanicInInit
	}
	return false
}

type ConsumerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ConsumerConfig) Reset() {
	*x = ConsumerConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_tests_integration_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsumerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsumerConfig) ProtoMessage() {}

func (x *ConsumerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_api_tests_integration_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsumerConfig.ProtoReflect.Descriptor instead.
func (*ConsumerConfig) Descriptor() ([]byte, []int) {
	return file_api_tests_integration_config_proto_rawDescGZIP(), []int{2}
}

func (x *ConsumerConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_api_tests_integration_config_proto protoreflect.FileDescriptor

var file_api_tests_integration_config_proto_rawDesc = []byte{
	0x0a, 0x22, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xe2, 0x01, 0x0a, 0x06,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x65, 0x65, 0x64, 0x5f, 0x62,
	0x75, 0x66, 0x66, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x6e, 0x65, 0x65,
	0x64, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x64, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72,
	0x73, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x4d, 0x73, 0x67, 0x12, 0x20,
	0x0a, 0x0c, 0x69, 0x6e, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x6e, 0x47, 0x72, 0x70, 0x63, 0x4d, 0x6f, 0x64, 0x65,
	0x22, 0xa9, 0x01, 0x0a, 0x0f, 0x42, 0x61, 0x64, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x28, 0x0a, 0x10, 0x70, 0x61, 0x6e, 0x69, 0x63, 0x5f, 0x69, 0x6e,
	0x5f, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e,
	0x70, 0x61, 0x6e, 0x69, 0x63, 0x49, 0x6e, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x24,
	0x0a, 0x0e, 0x70, 0x61, 0x6e, 0x69, 0x63, 0x5f, 0x69, 0x6e, 0x5f, 0x70, 0x61, 0x72, 0x73, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x70, 0x61, 0x6e, 0x69, 0x63, 0x49, 0x6e, 0x50,
	0x61, 0x72, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x69, 0x6e,
	0x5f, 0x69, 0x6e, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x49, 0x6e, 0x49, 0x6e, 0x69, 0x74, 0x12, 0x22, 0x0a, 0x0d, 0x70, 0x61, 0x6e, 0x69,
	0x63, 0x5f, 0x69, 0x6e, 0x5f, 0x69, 0x6e, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0b, 0x70, 0x61, 0x6e, 0x69, 0x63, 0x49, 0x6e, 0x49, 0x6e, 0x69, 0x74, 0x22, 0x24, 0x0a, 0x0e,
	0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x42, 0x24, 0x5a, 0x22, 0x6d, 0x6f, 0x73, 0x6e, 0x2e, 0x69, 0x6f, 0x2f, 0x68, 0x74,
	0x6e, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_tests_integration_config_proto_rawDescOnce sync.Once
	file_api_tests_integration_config_proto_rawDescData = file_api_tests_integration_config_proto_rawDesc
)

func file_api_tests_integration_config_proto_rawDescGZIP() []byte {
	file_api_tests_integration_config_proto_rawDescOnce.Do(func() {
		file_api_tests_integration_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_tests_integration_config_proto_rawDescData)
	})
	return file_api_tests_integration_config_proto_rawDescData
}

var file_api_tests_integration_config_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_tests_integration_config_proto_goTypes = []interface{}{
	(*Config)(nil),          // 0: api.tests.integration.Config
	(*BadPluginConfig)(nil), // 1: api.tests.integration.BadPluginConfig
	(*ConsumerConfig)(nil),  // 2: api.tests.integration.ConsumerConfig
}
var file_api_tests_integration_config_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_tests_integration_config_proto_init() }
func file_api_tests_integration_config_proto_init() {
	if File_api_tests_integration_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_tests_integration_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
		file_api_tests_integration_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BadPluginConfig); i {
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
		file_api_tests_integration_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsumerConfig); i {
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
			RawDescriptor: file_api_tests_integration_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_tests_integration_config_proto_goTypes,
		DependencyIndexes: file_api_tests_integration_config_proto_depIdxs,
		MessageInfos:      file_api_tests_integration_config_proto_msgTypes,
	}.Build()
	File_api_tests_integration_config_proto = out.File
	file_api_tests_integration_config_proto_rawDesc = nil
	file_api_tests_integration_config_proto_goTypes = nil
	file_api_tests_integration_config_proto_depIdxs = nil
}
