// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: corda/view_data.proto

package corda

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

type ViewData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Notarizations []*ViewData_Notarization `protobuf:"bytes,1,rep,name=notarizations,proto3" json:"notarizations,omitempty"`
	// Bytes of InteropPayload
	Payload []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *ViewData) Reset() {
	*x = ViewData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_corda_view_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewData) ProtoMessage() {}

func (x *ViewData) ProtoReflect() protoreflect.Message {
	mi := &file_corda_view_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewData.ProtoReflect.Descriptor instead.
func (*ViewData) Descriptor() ([]byte, []int) {
	return file_corda_view_data_proto_rawDescGZIP(), []int{0}
}

func (x *ViewData) GetNotarizations() []*ViewData_Notarization {
	if x != nil {
		return x.Notarizations
	}
	return nil
}

func (x *ViewData) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

type ViewData_Notarization struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signature   string `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	Certificate string `protobuf:"bytes,2,opt,name=certificate,proto3" json:"certificate,omitempty"`
	Id          string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ViewData_Notarization) Reset() {
	*x = ViewData_Notarization{}
	if protoimpl.UnsafeEnabled {
		mi := &file_corda_view_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewData_Notarization) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewData_Notarization) ProtoMessage() {}

func (x *ViewData_Notarization) ProtoReflect() protoreflect.Message {
	mi := &file_corda_view_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewData_Notarization.ProtoReflect.Descriptor instead.
func (*ViewData_Notarization) Descriptor() ([]byte, []int) {
	return file_corda_view_data_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ViewData_Notarization) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *ViewData_Notarization) GetCertificate() string {
	if x != nil {
		return x.Certificate
	}
	return ""
}

func (x *ViewData_Notarization) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_corda_view_data_proto protoreflect.FileDescriptor

var file_corda_view_data_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x6f, 0x72, 0x64, 0x61, 0x2f, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x63, 0x6f, 0x72, 0x64, 0x61, 0x22, 0xc8,
	0x01, 0x0a, 0x08, 0x56, 0x69, 0x65, 0x77, 0x44, 0x61, 0x74, 0x61, 0x12, 0x42, 0x0a, 0x0d, 0x6e,
	0x6f, 0x74, 0x61, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x72, 0x64, 0x61, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x44,
	0x61, 0x74, 0x61, 0x2e, 0x4e, 0x6f, 0x74, 0x61, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0d, 0x6e, 0x6f, 0x74, 0x61, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x5e, 0x0a, 0x0c, 0x4e, 0x6f, 0x74,
	0x61, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x42, 0x69, 0x0a, 0x17, 0x63, 0x6f, 0x6d,
	0x2e, 0x77, 0x65, 0x61, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x63,
	0x6f, 0x72, 0x64, 0x61, 0x5a, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x68, 0x79, 0x70, 0x65, 0x72, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2d, 0x6c, 0x61, 0x62,
	0x73, 0x2f, 0x77, 0x65, 0x61, 0x76, 0x65, 0x72, 0x2d, 0x64, 0x6c, 0x74, 0x2d, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2d, 0x67, 0x6f, 0x2f, 0x63,
	0x6f, 0x72, 0x64, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_corda_view_data_proto_rawDescOnce sync.Once
	file_corda_view_data_proto_rawDescData = file_corda_view_data_proto_rawDesc
)

func file_corda_view_data_proto_rawDescGZIP() []byte {
	file_corda_view_data_proto_rawDescOnce.Do(func() {
		file_corda_view_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_corda_view_data_proto_rawDescData)
	})
	return file_corda_view_data_proto_rawDescData
}

var file_corda_view_data_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_corda_view_data_proto_goTypes = []interface{}{
	(*ViewData)(nil),              // 0: corda.ViewData
	(*ViewData_Notarization)(nil), // 1: corda.ViewData.Notarization
}
var file_corda_view_data_proto_depIdxs = []int32{
	1, // 0: corda.ViewData.notarizations:type_name -> corda.ViewData.Notarization
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_corda_view_data_proto_init() }
func file_corda_view_data_proto_init() {
	if File_corda_view_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_corda_view_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewData); i {
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
		file_corda_view_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewData_Notarization); i {
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
			RawDescriptor: file_corda_view_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_corda_view_data_proto_goTypes,
		DependencyIndexes: file_corda_view_data_proto_depIdxs,
		MessageInfos:      file_corda_view_data_proto_msgTypes,
	}.Build()
	File_corda_view_data_proto = out.File
	file_corda_view_data_proto_rawDesc = nil
	file_corda_view_data_proto_goTypes = nil
	file_corda_view_data_proto_depIdxs = nil
}
