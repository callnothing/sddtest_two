// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: XXXBBB.proto

package xxxx

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type HelloRequestt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *HelloRequestt) Reset() {
	*x = HelloRequestt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_XXXBBB_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequestt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequestt) ProtoMessage() {}

func (x *HelloRequestt) ProtoReflect() protoreflect.Message {
	mi := &file_XXXBBB_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequestt.ProtoReflect.Descriptor instead.
func (*HelloRequestt) Descriptor() ([]byte, []int) {
	return file_XXXBBB_proto_rawDescGZIP(), []int{0}
}

func (x *HelloRequestt) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type HelloReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg  string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Code string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *HelloReply) Reset() {
	*x = HelloReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_XXXBBB_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReply) ProtoMessage() {}

func (x *HelloReply) ProtoReflect() protoreflect.Message {
	mi := &file_XXXBBB_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReply.ProtoReflect.Descriptor instead.
func (*HelloReply) Descriptor() ([]byte, []int) {
	return file_XXXBBB_proto_rawDescGZIP(), []int{1}
}

func (x *HelloReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *HelloReply) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

var File_XXXBBB_proto protoreflect.FileDescriptor

var file_XXXBBB_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x58, 0x58, 0x58, 0x42, 0x42, 0x42, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x78, 0x78, 0x78, 0x78, 0x78, 0x2e, 0x78, 0x78, 0x78, 0x78, 0x22, 0x21, 0x0a, 0x0d, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x32, 0x0a,
	0x0a, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x32, 0x47, 0x0a, 0x06, 0x58, 0x58, 0x58, 0x42, 0x42, 0x42, 0x12, 0x3d, 0x0a, 0x08, 0x53,
	0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x19, 0x2e, 0x78, 0x78, 0x78, 0x78, 0x78, 0x2e,
	0x78, 0x78, 0x78, 0x78, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x74, 0x1a, 0x16, 0x2e, 0x78, 0x78, 0x78, 0x78, 0x78, 0x2e, 0x78, 0x78, 0x78, 0x78, 0x2e,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x6c, 0x6c, 0x6e, 0x6f, 0x74,
	0x68, 0x69, 0x6e, 0x67, 0x2f, 0x73, 0x64, 0x64, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x77, 0x6f,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x6f, 0x2f, 0x78, 0x78, 0x78, 0x78, 0x78, 0x2f, 0x78,
	0x78, 0x78, 0x78, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_XXXBBB_proto_rawDescOnce sync.Once
	file_XXXBBB_proto_rawDescData = file_XXXBBB_proto_rawDesc
)

func file_XXXBBB_proto_rawDescGZIP() []byte {
	file_XXXBBB_proto_rawDescOnce.Do(func() {
		file_XXXBBB_proto_rawDescData = protoimpl.X.CompressGZIP(file_XXXBBB_proto_rawDescData)
	})
	return file_XXXBBB_proto_rawDescData
}

var file_XXXBBB_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_XXXBBB_proto_goTypes = []interface{}{
	(*HelloRequestt)(nil), // 0: xxxxx.xxxx.HelloRequestt
	(*HelloReply)(nil),    // 1: xxxxx.xxxx.HelloReply
}
var file_XXXBBB_proto_depIdxs = []int32{
	0, // 0: xxxxx.xxxx.XXXBBB.SayHello:input_type -> xxxxx.xxxx.HelloRequestt
	1, // 1: xxxxx.xxxx.XXXBBB.SayHello:output_type -> xxxxx.xxxx.HelloReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_XXXBBB_proto_init() }
func file_XXXBBB_proto_init() {
	if File_XXXBBB_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_XXXBBB_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequestt); i {
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
		file_XXXBBB_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloReply); i {
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
			RawDescriptor: file_XXXBBB_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_XXXBBB_proto_goTypes,
		DependencyIndexes: file_XXXBBB_proto_depIdxs,
		MessageInfos:      file_XXXBBB_proto_msgTypes,
	}.Build()
	File_XXXBBB_proto = out.File
	file_XXXBBB_proto_rawDesc = nil
	file_XXXBBB_proto_goTypes = nil
	file_XXXBBB_proto_depIdxs = nil
}
