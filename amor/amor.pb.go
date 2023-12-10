// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: amor/amor.proto

package amor

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetHomeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetHomeRequest) Reset() {
	*x = GetHomeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_amor_amor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHomeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHomeRequest) ProtoMessage() {}

func (x *GetHomeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_amor_amor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHomeRequest.ProtoReflect.Descriptor instead.
func (*GetHomeRequest) Descriptor() ([]byte, []int) {
	return file_amor_amor_proto_rawDescGZIP(), []int{0}
}

func (x *GetHomeRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetHomeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Location string  `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	Type     string  `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Rating   float64 `protobuf:"fixed64,5,opt,name=rating,proto3" json:"rating,omitempty"`
	Owner    string  `protobuf:"bytes,6,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (x *GetHomeResponse) Reset() {
	*x = GetHomeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_amor_amor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHomeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHomeResponse) ProtoMessage() {}

func (x *GetHomeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_amor_amor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHomeResponse.ProtoReflect.Descriptor instead.
func (*GetHomeResponse) Descriptor() ([]byte, []int) {
	return file_amor_amor_proto_rawDescGZIP(), []int{1}
}

func (x *GetHomeResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetHomeResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetHomeResponse) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *GetHomeResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *GetHomeResponse) GetRating() float64 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *GetHomeResponse) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

type UpdateHomeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Location string `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	Type     string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Owner    string `protobuf:"bytes,6,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (x *UpdateHomeRequest) Reset() {
	*x = UpdateHomeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_amor_amor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateHomeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateHomeRequest) ProtoMessage() {}

func (x *UpdateHomeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_amor_amor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateHomeRequest.ProtoReflect.Descriptor instead.
func (*UpdateHomeRequest) Descriptor() ([]byte, []int) {
	return file_amor_amor_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateHomeRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateHomeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateHomeRequest) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *UpdateHomeRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *UpdateHomeRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

type DeleteHomeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteHomeRequest) Reset() {
	*x = DeleteHomeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_amor_amor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteHomeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteHomeRequest) ProtoMessage() {}

func (x *DeleteHomeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_amor_amor_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteHomeRequest.ProtoReflect.Descriptor instead.
func (*DeleteHomeRequest) Descriptor() ([]byte, []int) {
	return file_amor_amor_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteHomeRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type AddHomeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type     string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Location string `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	Owner    string `protobuf:"bytes,4,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (x *AddHomeRequest) Reset() {
	*x = AddHomeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_amor_amor_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddHomeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddHomeRequest) ProtoMessage() {}

func (x *AddHomeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_amor_amor_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddHomeRequest.ProtoReflect.Descriptor instead.
func (*AddHomeRequest) Descriptor() ([]byte, []int) {
	return file_amor_amor_proto_rawDescGZIP(), []int{4}
}

func (x *AddHomeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddHomeRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *AddHomeRequest) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *AddHomeRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

type AddHomeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *AddHomeResponse) Reset() {
	*x = AddHomeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_amor_amor_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddHomeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddHomeResponse) ProtoMessage() {}

func (x *AddHomeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_amor_amor_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddHomeResponse.ProtoReflect.Descriptor instead.
func (*AddHomeResponse) Descriptor() ([]byte, []int) {
	return file_amor_amor_proto_rawDescGZIP(), []int{5}
}

func (x *AddHomeResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_amor_amor_proto protoreflect.FileDescriptor

var file_amor_amor_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x61, 0x6d, 0x6f, 0x72, 0x2f, 0x61, 0x6d, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x20, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x48, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x93,
	0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x14,
	0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x22, 0x7d, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x6f,
	0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x6f, 0x6d,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x6a, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x48,
	0x6f, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x22, 0x25, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x48, 0x6f, 0x6d, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0xb8, 0x02, 0x0a, 0x0b,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x41, 0x6d, 0x6f, 0x72, 0x12, 0x42, 0x0a, 0x07, 0x41,
	0x64, 0x64, 0x48, 0x6f, 0x6d, 0x65, 0x12, 0x0f, 0x2e, 0x41, 0x64, 0x64, 0x48, 0x6f, 0x6d, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x41, 0x64, 0x64, 0x48, 0x6f, 0x6d,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0e, 0x12, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x6f, 0x6d, 0x65, 0x2f, 0x61, 0x64, 0x64, 0x12,
	0x4f, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x6f, 0x6d, 0x65, 0x12, 0x12, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0f, 0x2a, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x6f, 0x6d, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x12, 0x4f, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x6d, 0x65, 0x12, 0x12,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0f, 0x1a, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x6f, 0x6d, 0x65, 0x2f, 0x7b, 0x69, 0x64,
	0x7d, 0x12, 0x43, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x6d, 0x65, 0x12, 0x0f, 0x2e, 0x47,
	0x65, 0x74, 0x48, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e,
	0x47, 0x65, 0x74, 0x48, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x6f, 0x6d,
	0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x69, 0x6e, 0x6f, 0x6c, 0x69, 0x74, 0x6f, 0x2f, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x61, 0x6d, 0x6f, 0x72, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6d,
	0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_amor_amor_proto_rawDescOnce sync.Once
	file_amor_amor_proto_rawDescData = file_amor_amor_proto_rawDesc
)

func file_amor_amor_proto_rawDescGZIP() []byte {
	file_amor_amor_proto_rawDescOnce.Do(func() {
		file_amor_amor_proto_rawDescData = protoimpl.X.CompressGZIP(file_amor_amor_proto_rawDescData)
	})
	return file_amor_amor_proto_rawDescData
}

var file_amor_amor_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_amor_amor_proto_goTypes = []interface{}{
	(*GetHomeRequest)(nil),    // 0: GetHomeRequest
	(*GetHomeResponse)(nil),   // 1: GetHomeResponse
	(*UpdateHomeRequest)(nil), // 2: UpdateHomeRequest
	(*DeleteHomeRequest)(nil), // 3: DeleteHomeRequest
	(*AddHomeRequest)(nil),    // 4: AddHomeRequest
	(*AddHomeResponse)(nil),   // 5: AddHomeResponse
	(*empty.Empty)(nil),       // 6: google.protobuf.Empty
}
var file_amor_amor_proto_depIdxs = []int32{
	4, // 0: ProjectAmor.AddHome:input_type -> AddHomeRequest
	3, // 1: ProjectAmor.DeleteHome:input_type -> DeleteHomeRequest
	2, // 2: ProjectAmor.UpdateHome:input_type -> UpdateHomeRequest
	0, // 3: ProjectAmor.GetHome:input_type -> GetHomeRequest
	5, // 4: ProjectAmor.AddHome:output_type -> AddHomeResponse
	6, // 5: ProjectAmor.DeleteHome:output_type -> google.protobuf.Empty
	6, // 6: ProjectAmor.UpdateHome:output_type -> google.protobuf.Empty
	1, // 7: ProjectAmor.GetHome:output_type -> GetHomeResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_amor_amor_proto_init() }
func file_amor_amor_proto_init() {
	if File_amor_amor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_amor_amor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHomeRequest); i {
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
		file_amor_amor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHomeResponse); i {
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
		file_amor_amor_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateHomeRequest); i {
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
		file_amor_amor_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteHomeRequest); i {
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
		file_amor_amor_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddHomeRequest); i {
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
		file_amor_amor_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddHomeResponse); i {
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
			RawDescriptor: file_amor_amor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_amor_amor_proto_goTypes,
		DependencyIndexes: file_amor_amor_proto_depIdxs,
		MessageInfos:      file_amor_amor_proto_msgTypes,
	}.Build()
	File_amor_amor_proto = out.File
	file_amor_amor_proto_rawDesc = nil
	file_amor_amor_proto_goTypes = nil
	file_amor_amor_proto_depIdxs = nil
}
