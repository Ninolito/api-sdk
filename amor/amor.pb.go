// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: amor/amor.proto

package amor

import (
	amor_var "github.com/Ninolito/api-sdk/other/amor_var"
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

type ListHomeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListHomeRequest) Reset() {
	*x = ListHomeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_amor_amor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListHomeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListHomeRequest) ProtoMessage() {}

func (x *ListHomeRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ListHomeRequest.ProtoReflect.Descriptor instead.
func (*ListHomeRequest) Descriptor() ([]byte, []int) {
	return file_amor_amor_proto_rawDescGZIP(), []int{0}
}

type ListHomeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Latlong []float64 `protobuf:"fixed64,2,rep,packed,name=latlong,proto3" json:"latlong,omitempty"`
}

func (x *ListHomeResponse) Reset() {
	*x = ListHomeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_amor_amor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListHomeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListHomeResponse) ProtoMessage() {}

func (x *ListHomeResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ListHomeResponse.ProtoReflect.Descriptor instead.
func (*ListHomeResponse) Descriptor() ([]byte, []int) {
	return file_amor_amor_proto_rawDescGZIP(), []int{1}
}

func (x *ListHomeResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ListHomeResponse) GetLatlong() []float64 {
	if x != nil {
		return x.Latlong
	}
	return nil
}

type GetHomeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetHomeRequest) Reset() {
	*x = GetHomeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_amor_amor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHomeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHomeRequest) ProtoMessage() {}

func (x *GetHomeRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetHomeRequest.ProtoReflect.Descriptor instead.
func (*GetHomeRequest) Descriptor() ([]byte, []int) {
	return file_amor_amor_proto_rawDescGZIP(), []int{2}
}

func (x *GetHomeRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetHomeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string              `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Address   string              `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Type      string              `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Rating    float64             `protobuf:"fixed64,5,opt,name=rating,proto3" json:"rating,omitempty"`
	Owner     string              `protobuf:"bytes,6,opt,name=owner,proto3" json:"owner,omitempty"`
	Latlong   []float64           `protobuf:"fixed64,7,rep,packed,name=latlong,proto3" json:"latlong,omitempty"`
	Utilities *amor_var.Utilities `protobuf:"bytes,8,opt,name=utilities,proto3" json:"utilities,omitempty"`
}

func (x *GetHomeResponse) Reset() {
	*x = GetHomeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_amor_amor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHomeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHomeResponse) ProtoMessage() {}

func (x *GetHomeResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetHomeResponse.ProtoReflect.Descriptor instead.
func (*GetHomeResponse) Descriptor() ([]byte, []int) {
	return file_amor_amor_proto_rawDescGZIP(), []int{3}
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

func (x *GetHomeResponse) GetAddress() string {
	if x != nil {
		return x.Address
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

func (x *GetHomeResponse) GetLatlong() []float64 {
	if x != nil {
		return x.Latlong
	}
	return nil
}

func (x *GetHomeResponse) GetUtilities() *amor_var.Utilities {
	if x != nil {
		return x.Utilities
	}
	return nil
}

type UpdateHomeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Address string    `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Type    string    `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Owner   string    `protobuf:"bytes,5,opt,name=owner,proto3" json:"owner,omitempty"`
	Latlong []float64 `protobuf:"fixed64,6,rep,packed,name=latlong,proto3" json:"latlong,omitempty"`
}

func (x *UpdateHomeRequest) Reset() {
	*x = UpdateHomeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_amor_amor_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateHomeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateHomeRequest) ProtoMessage() {}

func (x *UpdateHomeRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use UpdateHomeRequest.ProtoReflect.Descriptor instead.
func (*UpdateHomeRequest) Descriptor() ([]byte, []int) {
	return file_amor_amor_proto_rawDescGZIP(), []int{4}
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

func (x *UpdateHomeRequest) GetAddress() string {
	if x != nil {
		return x.Address
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

func (x *UpdateHomeRequest) GetLatlong() []float64 {
	if x != nil {
		return x.Latlong
	}
	return nil
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
		mi := &file_amor_amor_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteHomeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteHomeRequest) ProtoMessage() {}

func (x *DeleteHomeRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DeleteHomeRequest.ProtoReflect.Descriptor instead.
func (*DeleteHomeRequest) Descriptor() ([]byte, []int) {
	return file_amor_amor_proto_rawDescGZIP(), []int{5}
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

	Name    string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type    string    `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Address string    `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Owner   string    `protobuf:"bytes,4,opt,name=owner,proto3" json:"owner,omitempty"`
	Latlong []float64 `protobuf:"fixed64,5,rep,packed,name=latlong,proto3" json:"latlong,omitempty"`
}

func (x *AddHomeRequest) Reset() {
	*x = AddHomeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_amor_amor_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddHomeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddHomeRequest) ProtoMessage() {}

func (x *AddHomeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_amor_amor_proto_msgTypes[6]
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
	return file_amor_amor_proto_rawDescGZIP(), []int{6}
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

func (x *AddHomeRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *AddHomeRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *AddHomeRequest) GetLatlong() []float64 {
	if x != nil {
		return x.Latlong
	}
	return nil
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
		mi := &file_amor_amor_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddHomeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddHomeResponse) ProtoMessage() {}

func (x *AddHomeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_amor_amor_proto_msgTypes[7]
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
	return file_amor_amor_proto_rawDescGZIP(), []int{7}
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
	0x6f, 0x12, 0x17, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x61, 0x6d, 0x6f, 0x72, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x6d, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x2f,
	0x61, 0x6d, 0x6f, 0x72, 0x5f, 0x76, 0x61, 0x72, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x11, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x6f, 0x6d,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3c, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74,
	0x48, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x6c, 0x61, 0x74, 0x6c, 0x6f, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x03, 0x28, 0x01, 0x52, 0x07, 0x6c,
	0x61, 0x74, 0x6c, 0x6f, 0x6e, 0x67, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x6d,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0xd5, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x48, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06,
	0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07,
	0x6c, 0x61, 0x74, 0x6c, 0x6f, 0x6e, 0x67, 0x18, 0x07, 0x20, 0x03, 0x28, 0x01, 0x52, 0x07, 0x6c,
	0x61, 0x74, 0x6c, 0x6f, 0x6e, 0x67, 0x12, 0x28, 0x0a, 0x09, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x55, 0x74, 0x69, 0x6c,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x09, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x22, 0x95, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x6d, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x18,
	0x0a, 0x07, 0x6c, 0x61, 0x74, 0x6c, 0x6f, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x03, 0x28, 0x01, 0x52,
	0x07, 0x6c, 0x61, 0x74, 0x6c, 0x6f, 0x6e, 0x67, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x48, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x82, 0x01,
	0x0a, 0x0e, 0x41, 0x64, 0x64, 0x48, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x61, 0x74, 0x6c,
	0x6f, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x03, 0x28, 0x01, 0x52, 0x07, 0x6c, 0x61, 0x74, 0x6c, 0x6f,
	0x6e, 0x67, 0x22, 0x25, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x48, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0xc6, 0x04, 0x0a, 0x0b, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x41, 0x6d, 0x6f, 0x72, 0x12, 0x72, 0x0a, 0x07, 0x41, 0x64, 0x64,
	0x48, 0x6f, 0x6d, 0x65, 0x12, 0x27, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x61, 0x6d,
	0x6f, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6d, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x64, 0x64, 0x48, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x61, 0x6d, 0x6f, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e,
	0x61, 0x6d, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x48, 0x6f, 0x6d, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12,
	0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x6f, 0x6d, 0x65, 0x2f, 0x61, 0x64, 0x64, 0x12, 0x67, 0x0a,
	0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x6f, 0x6d, 0x65, 0x12, 0x2a, 0x2e, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x61, 0x6d, 0x6f, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6d,
	0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x6f, 0x6d, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x2a, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x6f, 0x6d,
	0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x67, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x48, 0x6f, 0x6d, 0x65, 0x12, 0x2a, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x61, 0x6d,
	0x6f, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6d, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f,
	0x1a, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x6f, 0x6d, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12,
	0x73, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x6d, 0x65, 0x12, 0x27, 0x2e, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x61, 0x6d, 0x6f, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6d, 0x6f,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x61, 0x6d, 0x6f,
	0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6d, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x48, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x6f, 0x6d, 0x65, 0x2f,
	0x7b, 0x69, 0x64, 0x7d, 0x12, 0x7c, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x6f, 0x6d, 0x65,
	0x12, 0x28, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x61, 0x6d, 0x6f, 0x72, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x6d, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x48,
	0x6f, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x61, 0x6d, 0x6f, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6d, 0x6f,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x22, 0x11, 0x2f,
	0x76, 0x31, 0x2f, 0x68, 0x6f, 0x6d, 0x65, 0x2f, 0x61, 0x6c, 0x6c, 0x3a, 0x72, 0x65, 0x61, 0x64,
	0x30, 0x01, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x4e, 0x69, 0x6e, 0x6f, 0x6c, 0x69, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x61, 0x6d, 0x6f, 0x72, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6d, 0x6f, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_amor_amor_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_amor_amor_proto_goTypes = []interface{}{
	(*ListHomeRequest)(nil),    // 0: projectamor_api.amor.v1.ListHomeRequest
	(*ListHomeResponse)(nil),   // 1: projectamor_api.amor.v1.ListHomeResponse
	(*GetHomeRequest)(nil),     // 2: projectamor_api.amor.v1.GetHomeRequest
	(*GetHomeResponse)(nil),    // 3: projectamor_api.amor.v1.GetHomeResponse
	(*UpdateHomeRequest)(nil),  // 4: projectamor_api.amor.v1.UpdateHomeRequest
	(*DeleteHomeRequest)(nil),  // 5: projectamor_api.amor.v1.DeleteHomeRequest
	(*AddHomeRequest)(nil),     // 6: projectamor_api.amor.v1.AddHomeRequest
	(*AddHomeResponse)(nil),    // 7: projectamor_api.amor.v1.AddHomeResponse
	(*amor_var.Utilities)(nil), // 8: Utilities
	(*empty.Empty)(nil),        // 9: google.protobuf.Empty
}
var file_amor_amor_proto_depIdxs = []int32{
	8, // 0: projectamor_api.amor.v1.GetHomeResponse.utilities:type_name -> Utilities
	6, // 1: projectamor_api.amor.v1.ProjectAmor.AddHome:input_type -> projectamor_api.amor.v1.AddHomeRequest
	5, // 2: projectamor_api.amor.v1.ProjectAmor.DeleteHome:input_type -> projectamor_api.amor.v1.DeleteHomeRequest
	4, // 3: projectamor_api.amor.v1.ProjectAmor.UpdateHome:input_type -> projectamor_api.amor.v1.UpdateHomeRequest
	2, // 4: projectamor_api.amor.v1.ProjectAmor.GetHome:input_type -> projectamor_api.amor.v1.GetHomeRequest
	0, // 5: projectamor_api.amor.v1.ProjectAmor.ListHome:input_type -> projectamor_api.amor.v1.ListHomeRequest
	7, // 6: projectamor_api.amor.v1.ProjectAmor.AddHome:output_type -> projectamor_api.amor.v1.AddHomeResponse
	9, // 7: projectamor_api.amor.v1.ProjectAmor.DeleteHome:output_type -> google.protobuf.Empty
	9, // 8: projectamor_api.amor.v1.ProjectAmor.UpdateHome:output_type -> google.protobuf.Empty
	3, // 9: projectamor_api.amor.v1.ProjectAmor.GetHome:output_type -> projectamor_api.amor.v1.GetHomeResponse
	1, // 10: projectamor_api.amor.v1.ProjectAmor.ListHome:output_type -> projectamor_api.amor.v1.ListHomeResponse
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_amor_amor_proto_init() }
func file_amor_amor_proto_init() {
	if File_amor_amor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_amor_amor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListHomeRequest); i {
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
			switch v := v.(*ListHomeResponse); i {
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
		file_amor_amor_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_amor_amor_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_amor_amor_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_amor_amor_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_amor_amor_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
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
			NumMessages:   8,
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
