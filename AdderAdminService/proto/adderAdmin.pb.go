// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: AdderAdminService/proto/adderAdmin.proto

package adderAdmin

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

type Id struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Id) Reset() {
	*x = Id{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AdderAdminService_proto_adderAdmin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Id) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Id) ProtoMessage() {}

func (x *Id) ProtoReflect() protoreflect.Message {
	mi := &file_AdderAdminService_proto_adderAdmin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Id.ProtoReflect.Descriptor instead.
func (*Id) Descriptor() ([]byte, []int) {
	return file_AdderAdminService_proto_adderAdmin_proto_rawDescGZIP(), []int{0}
}

func (x *Id) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Ok struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *Ok) Reset() {
	*x = Ok{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AdderAdminService_proto_adderAdmin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ok) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ok) ProtoMessage() {}

func (x *Ok) ProtoReflect() protoreflect.Message {
	mi := &file_AdderAdminService_proto_adderAdmin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ok.ProtoReflect.Descriptor instead.
func (*Ok) Descriptor() ([]byte, []int) {
	return file_AdderAdminService_proto_adderAdmin_proto_rawDescGZIP(), []int{1}
}

func (x *Ok) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

var File_AdderAdminService_proto_adderAdmin_proto protoreflect.FileDescriptor

var file_AdderAdminService_proto_adderAdmin_proto_rawDesc = []byte{
	0x0a, 0x28, 0x41, 0x64, 0x64, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x64, 0x64, 0x65, 0x72, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x61, 0x64, 0x64, 0x65,
	0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x22, 0x14, 0x0a, 0x02, 0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x14, 0x0a, 0x02,
	0x6f, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02,
	0x6f, 0x6b, 0x32, 0x67, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x12, 0x2a, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x0e, 0x2e, 0x61,
	0x64, 0x64, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x69, 0x64, 0x1a, 0x0e, 0x2e, 0x61,
	0x64, 0x64, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6f, 0x6b, 0x12, 0x2d, 0x0a, 0x0b,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x0e, 0x2e, 0x61, 0x64,
	0x64, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x69, 0x64, 0x1a, 0x0e, 0x2e, 0x61, 0x64,
	0x64, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6f, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_AdderAdminService_proto_adderAdmin_proto_rawDescOnce sync.Once
	file_AdderAdminService_proto_adderAdmin_proto_rawDescData = file_AdderAdminService_proto_adderAdmin_proto_rawDesc
)

func file_AdderAdminService_proto_adderAdmin_proto_rawDescGZIP() []byte {
	file_AdderAdminService_proto_adderAdmin_proto_rawDescOnce.Do(func() {
		file_AdderAdminService_proto_adderAdmin_proto_rawDescData = protoimpl.X.CompressGZIP(file_AdderAdminService_proto_adderAdmin_proto_rawDescData)
	})
	return file_AdderAdminService_proto_adderAdmin_proto_rawDescData
}

var file_AdderAdminService_proto_adderAdmin_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_AdderAdminService_proto_adderAdmin_proto_goTypes = []interface{}{
	(*Id)(nil), // 0: adderAdmin.id
	(*Ok)(nil), // 1: adderAdmin.ok
}
var file_AdderAdminService_proto_adderAdmin_proto_depIdxs = []int32{
	0, // 0: adderAdmin.AdderAdmin.AddAdmin:input_type -> adderAdmin.id
	0, // 1: adderAdmin.AdderAdmin.DeleteAdmin:input_type -> adderAdmin.id
	1, // 2: adderAdmin.AdderAdmin.AddAdmin:output_type -> adderAdmin.ok
	1, // 3: adderAdmin.AdderAdmin.DeleteAdmin:output_type -> adderAdmin.ok
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_AdderAdminService_proto_adderAdmin_proto_init() }
func file_AdderAdminService_proto_adderAdmin_proto_init() {
	if File_AdderAdminService_proto_adderAdmin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_AdderAdminService_proto_adderAdmin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Id); i {
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
		file_AdderAdminService_proto_adderAdmin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ok); i {
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
			RawDescriptor: file_AdderAdminService_proto_adderAdmin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_AdderAdminService_proto_adderAdmin_proto_goTypes,
		DependencyIndexes: file_AdderAdminService_proto_adderAdmin_proto_depIdxs,
		MessageInfos:      file_AdderAdminService_proto_adderAdmin_proto_msgTypes,
	}.Build()
	File_AdderAdminService_proto_adderAdmin_proto = out.File
	file_AdderAdminService_proto_adderAdmin_proto_rawDesc = nil
	file_AdderAdminService_proto_adderAdmin_proto_goTypes = nil
	file_AdderAdminService_proto_adderAdmin_proto_depIdxs = nil
}