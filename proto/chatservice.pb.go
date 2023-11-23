// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: proto/chatservice.proto

package proto

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

type MessageType int32

const (
	MessageType_PUBLISH MessageType = 0
	MessageType_JOIN    MessageType = 1
	MessageType_LEAVE   MessageType = 2
)

// Enum value maps for MessageType.
var (
	MessageType_name = map[int32]string{
		0: "PUBLISH",
		1: "JOIN",
		2: "LEAVE",
	}
	MessageType_value = map[string]int32{
		"PUBLISH": 0,
		"JOIN":    1,
		"LEAVE":   2,
	}
)

func (x MessageType) Enum() *MessageType {
	p := new(MessageType)
	*p = x
	return p
}

func (x MessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_chatservice_proto_enumTypes[0].Descriptor()
}

func (MessageType) Type() protoreflect.EnumType {
	return &file_proto_chatservice_proto_enumTypes[0]
}

func (x MessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageType.Descriptor instead.
func (MessageType) EnumDescriptor() ([]byte, []int) {
	return file_proto_chatservice_proto_rawDescGZIP(), []int{0}
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientName int32       `protobuf:"varint,1,opt,name=clientName,proto3" json:"clientName,omitempty"`
	Message    string      `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Type       MessageType `protobuf:"varint,3,opt,name=type,proto3,enum=proto.MessageType" json:"type,omitempty"`
	Clock      []int32     `protobuf:"varint,4,rep,packed,name=clock,proto3" json:"clock,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_chatservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chatservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_proto_chatservice_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetClientName() int32 {
	if x != nil {
		return x.ClientName
	}
	return 0
}

func (x *Request) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Request) GetType() MessageType {
	if x != nil {
		return x.Type
	}
	return MessageType_PUBLISH
}

func (x *Request) GetClock() []int32 {
	if x != nil {
		return x.Clock
	}
	return nil
}

type Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientName int32       `protobuf:"varint,1,opt,name=clientName,proto3" json:"clientName,omitempty"`
	Message    string      `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Type       MessageType `protobuf:"varint,3,opt,name=type,proto3,enum=proto.MessageType" json:"type,omitempty"`
	Clock      []int32     `protobuf:"varint,4,rep,packed,name=clock,proto3" json:"clock,omitempty"`
}

func (x *Reply) Reset() {
	*x = Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_chatservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reply) ProtoMessage() {}

func (x *Reply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chatservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reply.ProtoReflect.Descriptor instead.
func (*Reply) Descriptor() ([]byte, []int) {
	return file_proto_chatservice_proto_rawDescGZIP(), []int{1}
}

func (x *Reply) GetClientName() int32 {
	if x != nil {
		return x.ClientName
	}
	return 0
}

func (x *Reply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Reply) GetType() MessageType {
	if x != nil {
		return x.Type
	}
	return MessageType_PUBLISH
}

func (x *Reply) GetClock() []int32 {
	if x != nil {
		return x.Clock
	}
	return nil
}

var File_proto_chatservice_proto protoreflect.FileDescriptor

var file_proto_chatservice_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x81, 0x01, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x04, 0x20, 0x03, 0x28, 0x05, 0x52, 0x05, 0x63,
	0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x7f, 0x0a, 0x05, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1e, 0x0a,
	0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x04, 0x20, 0x03, 0x28, 0x05, 0x52, 0x05,
	0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x2a, 0x2f, 0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x53, 0x48, 0x10,
	0x00, 0x12, 0x08, 0x0a, 0x04, 0x4a, 0x4f, 0x49, 0x4e, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x4c,
	0x45, 0x41, 0x56, 0x45, 0x10, 0x02, 0x32, 0x3a, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x28, 0x01,
	0x30, 0x01, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x65, 0x6d, 0x6a, 0x61, 0x6b, 0x6f, 0x62, 0x73, 0x65, 0x6e, 0x31, 0x2f, 0x64, 0x73, 0x79,
	0x73, 0x35, 0x32, 0x30, 0x32, 0x33, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_chatservice_proto_rawDescOnce sync.Once
	file_proto_chatservice_proto_rawDescData = file_proto_chatservice_proto_rawDesc
)

func file_proto_chatservice_proto_rawDescGZIP() []byte {
	file_proto_chatservice_proto_rawDescOnce.Do(func() {
		file_proto_chatservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_chatservice_proto_rawDescData)
	})
	return file_proto_chatservice_proto_rawDescData
}

var file_proto_chatservice_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_chatservice_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_chatservice_proto_goTypes = []interface{}{
	(MessageType)(0), // 0: proto.MessageType
	(*Request)(nil),  // 1: proto.Request
	(*Reply)(nil),    // 2: proto.Reply
}
var file_proto_chatservice_proto_depIdxs = []int32{
	0, // 0: proto.Request.type:type_name -> proto.MessageType
	0, // 1: proto.Reply.type:type_name -> proto.MessageType
	1, // 2: proto.ChatService.Message:input_type -> proto.Request
	2, // 3: proto.ChatService.Message:output_type -> proto.Reply
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_chatservice_proto_init() }
func file_proto_chatservice_proto_init() {
	if File_proto_chatservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_chatservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_proto_chatservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reply); i {
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
			RawDescriptor: file_proto_chatservice_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_chatservice_proto_goTypes,
		DependencyIndexes: file_proto_chatservice_proto_depIdxs,
		EnumInfos:         file_proto_chatservice_proto_enumTypes,
		MessageInfos:      file_proto_chatservice_proto_msgTypes,
	}.Build()
	File_proto_chatservice_proto = out.File
	file_proto_chatservice_proto_rawDesc = nil
	file_proto_chatservice_proto_goTypes = nil
	file_proto_chatservice_proto_depIdxs = nil
}
