//
//Copyright 2020 Wim Henderickx.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.17.3
// source: config/configpb/config.proto

package configpb

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

type Status_ResourceStatus int32

const (
	Status_None          Status_ResourceStatus = 0
	Status_Success       Status_ResourceStatus = 1
	Status_Failed        Status_ResourceStatus = 2
	Status_CreatePending Status_ResourceStatus = 3
	Status_UpdatePending Status_ResourceStatus = 4
	Status_DeletePending Status_ResourceStatus = 5
)

// Enum value maps for Status_ResourceStatus.
var (
	Status_ResourceStatus_name = map[int32]string{
		0: "None",
		1: "Success",
		2: "Failed",
		3: "CreatePending",
		4: "UpdatePending",
		5: "DeletePending",
	}
	Status_ResourceStatus_value = map[string]int32{
		"None":          0,
		"Success":       1,
		"Failed":        2,
		"CreatePending": 3,
		"UpdatePending": 4,
		"DeletePending": 5,
	}
)

func (x Status_ResourceStatus) Enum() *Status_ResourceStatus {
	p := new(Status_ResourceStatus)
	*p = x
	return p
}

func (x Status_ResourceStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status_ResourceStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_config_configpb_config_proto_enumTypes[0].Descriptor()
}

func (Status_ResourceStatus) Type() protoreflect.EnumType {
	return &file_config_configpb_config_proto_enumTypes[0]
}

func (x Status_ResourceStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status_ResourceStatus.Descriptor instead.
func (Status_ResourceStatus) EnumDescriptor() ([]byte, []int) {
	return file_config_configpb_config_proto_rawDescGZIP(), []int{3, 0}
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Level int32  `protobuf:"varint,2,opt,name=level,proto3" json:"level,omitempty"`
	Path  *Path  `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	Data  []byte `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_configpb_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_config_configpb_config_proto_msgTypes[0]
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
	return file_config_configpb_config_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Request) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *Request) GetPath() *Path {
	if x != nil {
		return x.Path
	}
	return nil
}

func (x *Request) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type ResourceKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Level int32  `protobuf:"varint,2,opt,name=level,proto3" json:"level,omitempty"`
}

func (x *ResourceKey) Reset() {
	*x = ResourceKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_configpb_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceKey) ProtoMessage() {}

func (x *ResourceKey) ProtoReflect() protoreflect.Message {
	mi := &file_config_configpb_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceKey.ProtoReflect.Descriptor instead.
func (*ResourceKey) Descriptor() ([]byte, []int) {
	return file_config_configpb_config_proto_rawDescGZIP(), []int{1}
}

func (x *ResourceKey) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ResourceKey) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

type Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Reply) Reset() {
	*x = Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_configpb_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reply) ProtoMessage() {}

func (x *Reply) ProtoReflect() protoreflect.Message {
	mi := &file_config_configpb_config_proto_msgTypes[2]
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
	return file_config_configpb_config_proto_rawDescGZIP(), []int{2}
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Level     int32                 `protobuf:"varint,2,opt,name=level,proto3" json:"level,omitempty"`
	Path      *Path                 `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	Data      []byte                `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	Status    Status_ResourceStatus `protobuf:"varint,5,opt,name=status,proto3,enum=config.Status_ResourceStatus" json:"status,omitempty"`
	Deviation []*Deviation          `protobuf:"bytes,6,rep,name=deviation,proto3" json:"deviation,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_configpb_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_config_configpb_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_config_configpb_config_proto_rawDescGZIP(), []int{3}
}

func (x *Status) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Status) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *Status) GetPath() *Path {
	if x != nil {
		return x.Path
	}
	return nil
}

func (x *Status) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Status) GetStatus() Status_ResourceStatus {
	if x != nil {
		return x.Status
	}
	return Status_None
}

func (x *Status) GetDeviation() []*Deviation {
	if x != nil {
		return x.Deviation
	}
	return nil
}

type Deviation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path  *Path  `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Deviation) Reset() {
	*x = Deviation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_configpb_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Deviation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Deviation) ProtoMessage() {}

func (x *Deviation) ProtoReflect() protoreflect.Message {
	mi := &file_config_configpb_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Deviation.ProtoReflect.Descriptor instead.
func (*Deviation) Descriptor() ([]byte, []int) {
	return file_config_configpb_config_proto_rawDescGZIP(), []int{4}
}

func (x *Deviation) GetPath() *Path {
	if x != nil {
		return x.Path
	}
	return nil
}

func (x *Deviation) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// Path encodes a data tree path as a series of repeated strings, with
// each element of the path representing a data tree node name and the
// associated attributes.
// Reference: gNMI Specification Section 2.2.2.
type Path struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Elem []*PathElem `protobuf:"bytes,3,rep,name=elem,proto3" json:"elem,omitempty"` // Elements of the path.
}

func (x *Path) Reset() {
	*x = Path{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_configpb_config_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Path) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Path) ProtoMessage() {}

func (x *Path) ProtoReflect() protoreflect.Message {
	mi := &file_config_configpb_config_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Path.ProtoReflect.Descriptor instead.
func (*Path) Descriptor() ([]byte, []int) {
	return file_config_configpb_config_proto_rawDescGZIP(), []int{5}
}

func (x *Path) GetElem() []*PathElem {
	if x != nil {
		return x.Elem
	}
	return nil
}

// PathElem encodes an element of a path, along with any attributes (keys)
// that may be associated with it.
type PathElem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`                                                                                       // The name of the element in the path.
	Key       map[string]string `protobuf:"bytes,2,rep,name=key,proto3" json:"key,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // Map of key (attribute) name to value.
	Attribute *Attribute        `protobuf:"bytes,3,opt,name=attribute,proto3" json:"attribute,omitempty"`
}

func (x *PathElem) Reset() {
	*x = PathElem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_configpb_config_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PathElem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PathElem) ProtoMessage() {}

func (x *PathElem) ProtoReflect() protoreflect.Message {
	mi := &file_config_configpb_config_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PathElem.ProtoReflect.Descriptor instead.
func (*PathElem) Descriptor() ([]byte, []int) {
	return file_config_configpb_config_proto_rawDescGZIP(), []int{6}
}

func (x *PathElem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PathElem) GetKey() map[string]string {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *PathElem) GetAttribute() *Attribute {
	if x != nil {
		return x.Attribute
	}
	return nil
}

// Attaribute is used for code generation
type Attribute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ElemType  string   `protobuf:"bytes,1,opt,name=elemType,proto3" json:"elemType,omitempty"`
	Enum      []string `protobuf:"bytes,2,rep,name=enum,proto3" json:"enum,omitempty"`
	Range     []int32  `protobuf:"varint,3,rep,packed,name=range,proto3" json:"range,omitempty"`
	Length    []int32  `protobuf:"varint,4,rep,packed,name=length,proto3" json:"length,omitempty"`
	Pattern   []string `protobuf:"bytes,5,rep,name=pattern,proto3" json:"pattern,omitempty"`
	Union     bool     `protobuf:"varint,6,opt,name=union,proto3" json:"union,omitempty"`
	Mandatory bool     `protobuf:"varint,7,opt,name=mandatory,proto3" json:"mandatory,omitempty"`
	Default   string   `protobuf:"bytes,8,opt,name=default,proto3" json:"default,omitempty"`
}

func (x *Attribute) Reset() {
	*x = Attribute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_configpb_config_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Attribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attribute) ProtoMessage() {}

func (x *Attribute) ProtoReflect() protoreflect.Message {
	mi := &file_config_configpb_config_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attribute.ProtoReflect.Descriptor instead.
func (*Attribute) Descriptor() ([]byte, []int) {
	return file_config_configpb_config_proto_rawDescGZIP(), []int{7}
}

func (x *Attribute) GetElemType() string {
	if x != nil {
		return x.ElemType
	}
	return ""
}

func (x *Attribute) GetEnum() []string {
	if x != nil {
		return x.Enum
	}
	return nil
}

func (x *Attribute) GetRange() []int32 {
	if x != nil {
		return x.Range
	}
	return nil
}

func (x *Attribute) GetLength() []int32 {
	if x != nil {
		return x.Length
	}
	return nil
}

func (x *Attribute) GetPattern() []string {
	if x != nil {
		return x.Pattern
	}
	return nil
}

func (x *Attribute) GetUnion() bool {
	if x != nil {
		return x.Union
	}
	return false
}

func (x *Attribute) GetMandatory() bool {
	if x != nil {
		return x.Mandatory
	}
	return false
}

func (x *Attribute) GetDefault() string {
	if x != nil {
		return x.Default
	}
	return ""
}

var File_config_configpb_config_proto protoreflect.FileDescriptor

var file_config_configpb_config_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x70,
	0x62, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x69, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x20, 0x0a, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x37, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4b, 0x65, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x07, 0x0a, 0x05, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0xbe, 0x02, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x20, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x50, 0x61, 0x74, 0x68, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x35,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2f, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x64, 0x65, 0x76,
	0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x6c, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x6f, 0x6e, 0x65,
	0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x10, 0x01, 0x12,
	0x0a, 0x0a, 0x06, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x10, 0x03, 0x12, 0x11,
	0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x10,
	0x04, 0x12, 0x11, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x65, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x10, 0x05, 0x22, 0x43, 0x0a, 0x09, 0x44, 0x65, 0x76, 0x69, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x20, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x52, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x2c, 0x0a, 0x04, 0x50, 0x61, 0x74,
	0x68, 0x12, 0x24, 0x0a, 0x04, 0x65, 0x6c, 0x65, 0x6d, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x45, 0x6c, 0x65,
	0x6d, 0x52, 0x04, 0x65, 0x6c, 0x65, 0x6d, 0x22, 0xb4, 0x01, 0x0a, 0x08, 0x50, 0x61, 0x74, 0x68,
	0x45, 0x6c, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x50,
	0x61, 0x74, 0x68, 0x45, 0x6c, 0x65, 0x6d, 0x2e, 0x4b, 0x65, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2f, 0x0a, 0x09, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x09, 0x61, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x1a, 0x36, 0x0a, 0x08, 0x4b, 0x65, 0x79, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xd1,
	0x01, 0x0a, 0x09, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x65, 0x6c, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x65, 0x6c, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x6e, 0x75, 0x6d,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x65, 0x6e, 0x75, 0x6d, 0x12, 0x14, 0x0a, 0x05,
	0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52, 0x05, 0x72, 0x61, 0x6e,
	0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x05, 0x52, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61,
	0x74, 0x74, 0x65, 0x72, 0x6e, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x74,
	0x74, 0x65, 0x72, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x05, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x61,
	0x6e, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x6d,
	0x61, 0x6e, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x32, 0xc5, 0x01, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0f,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0d, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x12, 0x2c, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x13, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4b, 0x65, 0x79, 0x1a, 0x0e, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x2a,
	0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x06, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4b, 0x65, 0x79, 0x1a, 0x0d, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x2d, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x2d, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2f, 0x6e, 0x64, 0x64, 0x2d,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_configpb_config_proto_rawDescOnce sync.Once
	file_config_configpb_config_proto_rawDescData = file_config_configpb_config_proto_rawDesc
)

func file_config_configpb_config_proto_rawDescGZIP() []byte {
	file_config_configpb_config_proto_rawDescOnce.Do(func() {
		file_config_configpb_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_configpb_config_proto_rawDescData)
	})
	return file_config_configpb_config_proto_rawDescData
}

var file_config_configpb_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_config_configpb_config_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_config_configpb_config_proto_goTypes = []interface{}{
	(Status_ResourceStatus)(0), // 0: config.Status.ResourceStatus
	(*Request)(nil),            // 1: config.Request
	(*ResourceKey)(nil),        // 2: config.ResourceKey
	(*Reply)(nil),              // 3: config.Reply
	(*Status)(nil),             // 4: config.Status
	(*Deviation)(nil),          // 5: config.Deviation
	(*Path)(nil),               // 6: config.Path
	(*PathElem)(nil),           // 7: config.PathElem
	(*Attribute)(nil),          // 8: config.Attribute
	nil,                        // 9: config.PathElem.KeyEntry
}
var file_config_configpb_config_proto_depIdxs = []int32{
	6,  // 0: config.Request.path:type_name -> config.Path
	6,  // 1: config.Status.path:type_name -> config.Path
	0,  // 2: config.Status.status:type_name -> config.Status.ResourceStatus
	5,  // 3: config.Status.deviation:type_name -> config.Deviation
	6,  // 4: config.Deviation.path:type_name -> config.Path
	7,  // 5: config.Path.elem:type_name -> config.PathElem
	9,  // 6: config.PathElem.key:type_name -> config.PathElem.KeyEntry
	8,  // 7: config.PathElem.attribute:type_name -> config.Attribute
	1,  // 8: config.Configuration.Create:input_type -> config.Request
	2,  // 9: config.Configuration.Get:input_type -> config.ResourceKey
	1,  // 10: config.Configuration.Update:input_type -> config.Request
	2,  // 11: config.Configuration.Delete:input_type -> config.ResourceKey
	3,  // 12: config.Configuration.Create:output_type -> config.Reply
	4,  // 13: config.Configuration.Get:output_type -> config.Status
	3,  // 14: config.Configuration.Update:output_type -> config.Reply
	3,  // 15: config.Configuration.Delete:output_type -> config.Reply
	12, // [12:16] is the sub-list for method output_type
	8,  // [8:12] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_config_configpb_config_proto_init() }
func file_config_configpb_config_proto_init() {
	if File_config_configpb_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_configpb_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_config_configpb_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceKey); i {
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
		file_config_configpb_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_config_configpb_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_config_configpb_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Deviation); i {
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
		file_config_configpb_config_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Path); i {
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
		file_config_configpb_config_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PathElem); i {
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
		file_config_configpb_config_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Attribute); i {
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
			RawDescriptor: file_config_configpb_config_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_config_configpb_config_proto_goTypes,
		DependencyIndexes: file_config_configpb_config_proto_depIdxs,
		EnumInfos:         file_config_configpb_config_proto_enumTypes,
		MessageInfos:      file_config_configpb_config_proto_msgTypes,
	}.Build()
	File_config_configpb_config_proto = out.File
	file_config_configpb_config_proto_rawDesc = nil
	file_config_configpb_config_proto_goTypes = nil
	file_config_configpb_config_proto_depIdxs = nil
}
