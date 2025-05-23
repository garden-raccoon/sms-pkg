// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: sms-pkg-service.proto

package sms_pkg

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SmsByRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserUuid      []byte                 `protobuf:"bytes,1,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SmsByRequest) Reset() {
	*x = SmsByRequest{}
	mi := &file_sms_pkg_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SmsByRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SmsByRequest) ProtoMessage() {}

func (x *SmsByRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sms_pkg_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SmsByRequest.ProtoReflect.Descriptor instead.
func (*SmsByRequest) Descriptor() ([]byte, []int) {
	return file_sms_pkg_service_proto_rawDescGZIP(), []int{0}
}

func (x *SmsByRequest) GetUserUuid() []byte {
	if x != nil {
		return x.UserUuid
	}
	return nil
}

type SmsEmpty struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SmsEmpty) Reset() {
	*x = SmsEmpty{}
	mi := &file_sms_pkg_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SmsEmpty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SmsEmpty) ProtoMessage() {}

func (x *SmsEmpty) ProtoReflect() protoreflect.Message {
	mi := &file_sms_pkg_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SmsEmpty.ProtoReflect.Descriptor instead.
func (*SmsEmpty) Descriptor() ([]byte, []int) {
	return file_sms_pkg_service_proto_rawDescGZIP(), []int{1}
}

var File_sms_pkg_service_proto protoreflect.FileDescriptor

const file_sms_pkg_service_proto_rawDesc = "" +
	"\n" +
	"\x15sms-pkg-service.proto\x12\x06models\x1a\x14sms-pkg-models.proto\"+\n" +
	"\fSmsByRequest\x12\x1b\n" +
	"\tuser_uuid\x18\x01 \x01(\fR\buserUuid\"\n" +
	"\n" +
	"\bSmsEmpty2\xa1\x01\n" +
	"\n" +
	"SmsService\x122\n" +
	"\x11CreateOrUpdateSms\x12\v.models.Sms\x1a\x10.models.SmsEmpty\x12*\n" +
	"\x05SmsBy\x12\x14.models.SmsByRequest\x1a\v.models.Sms\x123\n" +
	"\tDeleteSms\x12\x14.models.SmsByRequest\x1a\x10.models.SmsEmptyB\x13Z\x11protocols/sms-pkgb\x06proto3"

var (
	file_sms_pkg_service_proto_rawDescOnce sync.Once
	file_sms_pkg_service_proto_rawDescData []byte
)

func file_sms_pkg_service_proto_rawDescGZIP() []byte {
	file_sms_pkg_service_proto_rawDescOnce.Do(func() {
		file_sms_pkg_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_sms_pkg_service_proto_rawDesc), len(file_sms_pkg_service_proto_rawDesc)))
	})
	return file_sms_pkg_service_proto_rawDescData
}

var file_sms_pkg_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_sms_pkg_service_proto_goTypes = []any{
	(*SmsByRequest)(nil), // 0: models.SmsByRequest
	(*SmsEmpty)(nil),     // 1: models.SmsEmpty
	(*Sms)(nil),          // 2: models.Sms
}
var file_sms_pkg_service_proto_depIdxs = []int32{
	2, // 0: models.SmsService.CreateOrUpdateSms:input_type -> models.Sms
	0, // 1: models.SmsService.SmsBy:input_type -> models.SmsByRequest
	0, // 2: models.SmsService.DeleteSms:input_type -> models.SmsByRequest
	1, // 3: models.SmsService.CreateOrUpdateSms:output_type -> models.SmsEmpty
	2, // 4: models.SmsService.SmsBy:output_type -> models.Sms
	1, // 5: models.SmsService.DeleteSms:output_type -> models.SmsEmpty
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sms_pkg_service_proto_init() }
func file_sms_pkg_service_proto_init() {
	if File_sms_pkg_service_proto != nil {
		return
	}
	file_sms_pkg_models_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_sms_pkg_service_proto_rawDesc), len(file_sms_pkg_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sms_pkg_service_proto_goTypes,
		DependencyIndexes: file_sms_pkg_service_proto_depIdxs,
		MessageInfos:      file_sms_pkg_service_proto_msgTypes,
	}.Build()
	File_sms_pkg_service_proto = out.File
	file_sms_pkg_service_proto_goTypes = nil
	file_sms_pkg_service_proto_depIdxs = nil
}
