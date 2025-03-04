// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: gocica/v1/index_entry.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

// IndexEntry is a single entry in the index.
type IndexEntry struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OutputId      string                 `protobuf:"bytes,1,opt,name=output_id,json=outputId,proto3" json:"output_id,omitempty"`
	Size          int64                  `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Timenano      int64                  `protobuf:"varint,3,opt,name=timenano,proto3" json:"timenano,omitempty"`
	LastUsedAt    *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=last_used_at,json=lastUsedAt,proto3" json:"last_used_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IndexEntry) Reset() {
	*x = IndexEntry{}
	mi := &file_gocica_v1_index_entry_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IndexEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexEntry) ProtoMessage() {}

func (x *IndexEntry) ProtoReflect() protoreflect.Message {
	mi := &file_gocica_v1_index_entry_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexEntry.ProtoReflect.Descriptor instead.
func (*IndexEntry) Descriptor() ([]byte, []int) {
	return file_gocica_v1_index_entry_proto_rawDescGZIP(), []int{0}
}

func (x *IndexEntry) GetOutputId() string {
	if x != nil {
		return x.OutputId
	}
	return ""
}

func (x *IndexEntry) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *IndexEntry) GetTimenano() int64 {
	if x != nil {
		return x.Timenano
	}
	return 0
}

func (x *IndexEntry) GetLastUsedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUsedAt
	}
	return nil
}

// IndexEntryMap is a map of IndexEntry.
type IndexEntryMap struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Entries       map[string]*IndexEntry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IndexEntryMap) Reset() {
	*x = IndexEntryMap{}
	mi := &file_gocica_v1_index_entry_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IndexEntryMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexEntryMap) ProtoMessage() {}

func (x *IndexEntryMap) ProtoReflect() protoreflect.Message {
	mi := &file_gocica_v1_index_entry_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexEntryMap.ProtoReflect.Descriptor instead.
func (*IndexEntryMap) Descriptor() ([]byte, []int) {
	return file_gocica_v1_index_entry_proto_rawDescGZIP(), []int{1}
}

func (x *IndexEntryMap) GetEntries() map[string]*IndexEntry {
	if x != nil {
		return x.Entries
	}
	return nil
}

var File_gocica_v1_index_entry_proto protoreflect.FileDescriptor

var file_gocica_v1_index_entry_proto_rawDesc = string([]byte{
	0x0a, 0x1b, 0x67, 0x6f, 0x63, 0x69, 0x63, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x67,
	0x6f, 0x63, 0x69, 0x63, 0x61, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x01, 0x0a, 0x0a, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d,
	0x65, 0x6e, 0x61, 0x6e, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x69, 0x6d,
	0x65, 0x6e, 0x61, 0x6e, 0x6f, 0x12, 0x3c, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x73,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x73, 0x65,
	0x64, 0x41, 0x74, 0x22, 0xa3, 0x01, 0x0a, 0x0d, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x4d, 0x61, 0x70, 0x12, 0x3f, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x67, 0x6f, 0x63, 0x69, 0x63, 0x61, 0x2e,
	0x76, 0x31, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x4d, 0x61, 0x70,
	0x2e, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x65,
	0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x1a, 0x51, 0x0a, 0x0c, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2b, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x6f, 0x63, 0x69, 0x63, 0x61,
	0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x7a, 0x72, 0x65, 0x61, 0x6e, 0x2f,
	0x67, 0x6f, 0x63, 0x69, 0x63, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x63,
	0x69, 0x63, 0x61, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_gocica_v1_index_entry_proto_rawDescOnce sync.Once
	file_gocica_v1_index_entry_proto_rawDescData []byte
)

func file_gocica_v1_index_entry_proto_rawDescGZIP() []byte {
	file_gocica_v1_index_entry_proto_rawDescOnce.Do(func() {
		file_gocica_v1_index_entry_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_gocica_v1_index_entry_proto_rawDesc), len(file_gocica_v1_index_entry_proto_rawDesc)))
	})
	return file_gocica_v1_index_entry_proto_rawDescData
}

var file_gocica_v1_index_entry_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_gocica_v1_index_entry_proto_goTypes = []any{
	(*IndexEntry)(nil),            // 0: gocica.v1.IndexEntry
	(*IndexEntryMap)(nil),         // 1: gocica.v1.IndexEntryMap
	nil,                           // 2: gocica.v1.IndexEntryMap.EntriesEntry
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_gocica_v1_index_entry_proto_depIdxs = []int32{
	3, // 0: gocica.v1.IndexEntry.last_used_at:type_name -> google.protobuf.Timestamp
	2, // 1: gocica.v1.IndexEntryMap.entries:type_name -> gocica.v1.IndexEntryMap.EntriesEntry
	0, // 2: gocica.v1.IndexEntryMap.EntriesEntry.value:type_name -> gocica.v1.IndexEntry
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_gocica_v1_index_entry_proto_init() }
func file_gocica_v1_index_entry_proto_init() {
	if File_gocica_v1_index_entry_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_gocica_v1_index_entry_proto_rawDesc), len(file_gocica_v1_index_entry_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gocica_v1_index_entry_proto_goTypes,
		DependencyIndexes: file_gocica_v1_index_entry_proto_depIdxs,
		MessageInfos:      file_gocica_v1_index_entry_proto_msgTypes,
	}.Build()
	File_gocica_v1_index_entry_proto = out.File
	file_gocica_v1_index_entry_proto_goTypes = nil
	file_gocica_v1_index_entry_proto_depIdxs = nil
}
