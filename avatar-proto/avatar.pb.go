// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.26.1
// source: avatar.proto

package avatar_proto

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

type SetAvatarIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserUuid string `protobuf:"bytes,1,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	Filename string `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
	Batch    []byte `protobuf:"bytes,3,opt,name=batch,proto3" json:"batch,omitempty"`
}

func (x *SetAvatarIn) Reset() {
	*x = SetAvatarIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_avatar_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetAvatarIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetAvatarIn) ProtoMessage() {}

func (x *SetAvatarIn) ProtoReflect() protoreflect.Message {
	mi := &file_avatar_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetAvatarIn.ProtoReflect.Descriptor instead.
func (*SetAvatarIn) Descriptor() ([]byte, []int) {
	return file_avatar_proto_rawDescGZIP(), []int{0}
}

func (x *SetAvatarIn) GetUserUuid() string {
	if x != nil {
		return x.UserUuid
	}
	return ""
}

func (x *SetAvatarIn) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *SetAvatarIn) GetBatch() []byte {
	if x != nil {
		return x.Batch
	}
	return nil
}

type SetAvatarOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Link string `protobuf:"bytes,1,opt,name=link,proto3" json:"link,omitempty"`
}

func (x *SetAvatarOut) Reset() {
	*x = SetAvatarOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_avatar_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetAvatarOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetAvatarOut) ProtoMessage() {}

func (x *SetAvatarOut) ProtoReflect() protoreflect.Message {
	mi := &file_avatar_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetAvatarOut.ProtoReflect.Descriptor instead.
func (*SetAvatarOut) Descriptor() ([]byte, []int) {
	return file_avatar_proto_rawDescGZIP(), []int{1}
}

func (x *SetAvatarOut) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

type GetAllAvatarsIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserUuid string `protobuf:"bytes,1,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
}

func (x *GetAllAvatarsIn) Reset() {
	*x = GetAllAvatarsIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_avatar_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllAvatarsIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllAvatarsIn) ProtoMessage() {}

func (x *GetAllAvatarsIn) ProtoReflect() protoreflect.Message {
	mi := &file_avatar_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllAvatarsIn.ProtoReflect.Descriptor instead.
func (*GetAllAvatarsIn) Descriptor() ([]byte, []int) {
	return file_avatar_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllAvatarsIn) GetUserUuid() string {
	if x != nil {
		return x.UserUuid
	}
	return ""
}

type Avatar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Link string `protobuf:"bytes,2,opt,name=link,proto3" json:"link,omitempty"`
}

func (x *Avatar) Reset() {
	*x = Avatar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_avatar_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Avatar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Avatar) ProtoMessage() {}

func (x *Avatar) ProtoReflect() protoreflect.Message {
	mi := &file_avatar_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Avatar.ProtoReflect.Descriptor instead.
func (*Avatar) Descriptor() ([]byte, []int) {
	return file_avatar_proto_rawDescGZIP(), []int{3}
}

func (x *Avatar) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Avatar) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

type GetAllAvatarsOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvatarList []*Avatar `protobuf:"bytes,1,rep,name=avatar_list,json=avatarList,proto3" json:"avatar_list,omitempty"`
}

func (x *GetAllAvatarsOut) Reset() {
	*x = GetAllAvatarsOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_avatar_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllAvatarsOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllAvatarsOut) ProtoMessage() {}

func (x *GetAllAvatarsOut) ProtoReflect() protoreflect.Message {
	mi := &file_avatar_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllAvatarsOut.ProtoReflect.Descriptor instead.
func (*GetAllAvatarsOut) Descriptor() ([]byte, []int) {
	return file_avatar_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllAvatarsOut) GetAvatarList() []*Avatar {
	if x != nil {
		return x.AvatarList
	}
	return nil
}

type DeleteAvatarIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvatarId string `protobuf:"bytes,1,opt,name=avatar_id,json=avatarId,proto3" json:"avatar_id,omitempty"`
}

func (x *DeleteAvatarIn) Reset() {
	*x = DeleteAvatarIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_avatar_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAvatarIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAvatarIn) ProtoMessage() {}

func (x *DeleteAvatarIn) ProtoReflect() protoreflect.Message {
	mi := &file_avatar_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAvatarIn.ProtoReflect.Descriptor instead.
func (*DeleteAvatarIn) Descriptor() ([]byte, []int) {
	return file_avatar_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteAvatarIn) GetAvatarId() string {
	if x != nil {
		return x.AvatarId
	}
	return ""
}

var File_avatar_proto protoreflect.FileDescriptor

var file_avatar_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x22, 0x5c, 0x0a, 0x0b, 0x53, 0x65, 0x74, 0x41, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x49, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x75, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x55, 0x75,
	0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x62, 0x61, 0x74, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x62,
	0x61, 0x74, 0x63, 0x68, 0x22, 0x22, 0x0a, 0x0c, 0x53, 0x65, 0x74, 0x41, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x4f, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x22, 0x2e, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x73, 0x49, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x22, 0x2c, 0x0a, 0x06, 0x41, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x22, 0x43, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x73, 0x4f, 0x75, 0x74, 0x12, 0x2f, 0x0a, 0x0b, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x2e, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52,
	0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x2d, 0x0a, 0x0e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x6e, 0x12, 0x1b, 0x0a,
	0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x64, 0x32, 0xcb, 0x01, 0x0a, 0x0d, 0x41,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x09,
	0x53, 0x65, 0x74, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x13, 0x2e, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x2e, 0x53, 0x65, 0x74, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x6e, 0x1a, 0x14,
	0x2e, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x2e, 0x53, 0x65, 0x74, 0x41, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x28, 0x01, 0x12, 0x44, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x73, 0x12, 0x17, 0x2e, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x73,
	0x49, 0x6e, 0x1a, 0x18, 0x2e, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x73, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x12, 0x38,
	0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x16,
	0x2e, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x49, 0x6e, 0x1a, 0x0e, 0x2e, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x2e,
	0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_avatar_proto_rawDescOnce sync.Once
	file_avatar_proto_rawDescData = file_avatar_proto_rawDesc
)

func file_avatar_proto_rawDescGZIP() []byte {
	file_avatar_proto_rawDescOnce.Do(func() {
		file_avatar_proto_rawDescData = protoimpl.X.CompressGZIP(file_avatar_proto_rawDescData)
	})
	return file_avatar_proto_rawDescData
}

var file_avatar_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_avatar_proto_goTypes = []any{
	(*SetAvatarIn)(nil),      // 0: avatar.SetAvatarIn
	(*SetAvatarOut)(nil),     // 1: avatar.SetAvatarOut
	(*GetAllAvatarsIn)(nil),  // 2: avatar.GetAllAvatarsIn
	(*Avatar)(nil),           // 3: avatar.Avatar
	(*GetAllAvatarsOut)(nil), // 4: avatar.GetAllAvatarsOut
	(*DeleteAvatarIn)(nil),   // 5: avatar.DeleteAvatarIn
}
var file_avatar_proto_depIdxs = []int32{
	3, // 0: avatar.GetAllAvatarsOut.avatar_list:type_name -> avatar.Avatar
	0, // 1: avatar.AvatarService.SetAvatar:input_type -> avatar.SetAvatarIn
	2, // 2: avatar.AvatarService.GetAllAvatars:input_type -> avatar.GetAllAvatarsIn
	5, // 3: avatar.AvatarService.DeleteAvatar:input_type -> avatar.DeleteAvatarIn
	1, // 4: avatar.AvatarService.SetAvatar:output_type -> avatar.SetAvatarOut
	4, // 5: avatar.AvatarService.GetAllAvatars:output_type -> avatar.GetAllAvatarsOut
	3, // 6: avatar.AvatarService.DeleteAvatar:output_type -> avatar.Avatar
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_avatar_proto_init() }
func file_avatar_proto_init() {
	if File_avatar_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_avatar_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*SetAvatarIn); i {
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
		file_avatar_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*SetAvatarOut); i {
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
		file_avatar_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllAvatarsIn); i {
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
		file_avatar_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Avatar); i {
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
		file_avatar_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllAvatarsOut); i {
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
		file_avatar_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteAvatarIn); i {
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
			RawDescriptor: file_avatar_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_avatar_proto_goTypes,
		DependencyIndexes: file_avatar_proto_depIdxs,
		MessageInfos:      file_avatar_proto_msgTypes,
	}.Build()
	File_avatar_proto = out.File
	file_avatar_proto_rawDesc = nil
	file_avatar_proto_goTypes = nil
	file_avatar_proto_depIdxs = nil
}
