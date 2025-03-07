// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v4.23.4
// source: protos/sharding_rules.proto

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

type ShardingRuleEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Column       string `protobuf:"bytes,2,opt,name=column,proto3" json:"column,omitempty"`
	HashFunction string `protobuf:"bytes,3,opt,name=hashFunction,proto3" json:"hashFunction,omitempty"`
}

func (x *ShardingRuleEntry) Reset() {
	*x = ShardingRuleEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_sharding_rules_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShardingRuleEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShardingRuleEntry) ProtoMessage() {}

func (x *ShardingRuleEntry) ProtoReflect() protoreflect.Message {
	mi := &file_protos_sharding_rules_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShardingRuleEntry.ProtoReflect.Descriptor instead.
func (*ShardingRuleEntry) Descriptor() ([]byte, []int) {
	return file_protos_sharding_rules_proto_rawDescGZIP(), []int{0}
}

func (x *ShardingRuleEntry) GetColumn() string {
	if x != nil {
		return x.Column
	}
	return ""
}

func (x *ShardingRuleEntry) GetHashFunction() string {
	if x != nil {
		return x.HashFunction
	}
	return ""
}

type ShardingRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	TableName         string               `protobuf:"bytes,2,opt,name=tableName,proto3" json:"tableName,omitempty"`
	ShardingRuleEntry []*ShardingRuleEntry `protobuf:"bytes,3,rep,name=ShardingRuleEntry,proto3" json:"ShardingRuleEntry,omitempty"`
}

func (x *ShardingRule) Reset() {
	*x = ShardingRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_sharding_rules_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShardingRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShardingRule) ProtoMessage() {}

func (x *ShardingRule) ProtoReflect() protoreflect.Message {
	mi := &file_protos_sharding_rules_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShardingRule.ProtoReflect.Descriptor instead.
func (*ShardingRule) Descriptor() ([]byte, []int) {
	return file_protos_sharding_rules_proto_rawDescGZIP(), []int{1}
}

func (x *ShardingRule) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ShardingRule) GetTableName() string {
	if x != nil {
		return x.TableName
	}
	return ""
}

func (x *ShardingRule) GetShardingRuleEntry() []*ShardingRuleEntry {
	if x != nil {
		return x.ShardingRuleEntry
	}
	return nil
}

type AddShardingRuleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rules []*ShardingRule `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
}

func (x *AddShardingRuleRequest) Reset() {
	*x = AddShardingRuleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_sharding_rules_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddShardingRuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddShardingRuleRequest) ProtoMessage() {}

func (x *AddShardingRuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_sharding_rules_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddShardingRuleRequest.ProtoReflect.Descriptor instead.
func (*AddShardingRuleRequest) Descriptor() ([]byte, []int) {
	return file_protos_sharding_rules_proto_rawDescGZIP(), []int{2}
}

func (x *AddShardingRuleRequest) GetRules() []*ShardingRule {
	if x != nil {
		return x.Rules
	}
	return nil
}

type AddShardingRuleReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddShardingRuleReply) Reset() {
	*x = AddShardingRuleReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_sharding_rules_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddShardingRuleReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddShardingRuleReply) ProtoMessage() {}

func (x *AddShardingRuleReply) ProtoReflect() protoreflect.Message {
	mi := &file_protos_sharding_rules_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddShardingRuleReply.ProtoReflect.Descriptor instead.
func (*AddShardingRuleReply) Descriptor() ([]byte, []int) {
	return file_protos_sharding_rules_proto_rawDescGZIP(), []int{3}
}

type ListShardingRuleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListShardingRuleRequest) Reset() {
	*x = ListShardingRuleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_sharding_rules_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListShardingRuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListShardingRuleRequest) ProtoMessage() {}

func (x *ListShardingRuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_sharding_rules_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListShardingRuleRequest.ProtoReflect.Descriptor instead.
func (*ListShardingRuleRequest) Descriptor() ([]byte, []int) {
	return file_protos_sharding_rules_proto_rawDescGZIP(), []int{4}
}

type ListShardingRuleReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rules []*ShardingRule `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
}

func (x *ListShardingRuleReply) Reset() {
	*x = ListShardingRuleReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_sharding_rules_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListShardingRuleReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListShardingRuleReply) ProtoMessage() {}

func (x *ListShardingRuleReply) ProtoReflect() protoreflect.Message {
	mi := &file_protos_sharding_rules_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListShardingRuleReply.ProtoReflect.Descriptor instead.
func (*ListShardingRuleReply) Descriptor() ([]byte, []int) {
	return file_protos_sharding_rules_proto_rawDescGZIP(), []int{5}
}

func (x *ListShardingRuleReply) GetRules() []*ShardingRule {
	if x != nil {
		return x.Rules
	}
	return nil
}

type DropShardingRuleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id []string `protobuf:"bytes,1,rep,name=id,proto3" json:"id,omitempty"`
}

func (x *DropShardingRuleRequest) Reset() {
	*x = DropShardingRuleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_sharding_rules_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DropShardingRuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DropShardingRuleRequest) ProtoMessage() {}

func (x *DropShardingRuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_sharding_rules_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DropShardingRuleRequest.ProtoReflect.Descriptor instead.
func (*DropShardingRuleRequest) Descriptor() ([]byte, []int) {
	return file_protos_sharding_rules_proto_rawDescGZIP(), []int{6}
}

func (x *DropShardingRuleRequest) GetId() []string {
	if x != nil {
		return x.Id
	}
	return nil
}

type DropShardingRuleReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DropShardingRuleReply) Reset() {
	*x = DropShardingRuleReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_sharding_rules_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DropShardingRuleReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DropShardingRuleReply) ProtoMessage() {}

func (x *DropShardingRuleReply) ProtoReflect() protoreflect.Message {
	mi := &file_protos_sharding_rules_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DropShardingRuleReply.ProtoReflect.Descriptor instead.
func (*DropShardingRuleReply) Descriptor() ([]byte, []int) {
	return file_protos_sharding_rules_proto_rawDescGZIP(), []int{7}
}

var File_protos_sharding_rules_proto protoreflect.FileDescriptor

var file_protos_sharding_rules_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e,
	0x67, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x73,
	0x70, 0x71, 0x72, 0x22, 0x4f, 0x0a, 0x11, 0x53, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x52,
	0x75, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6c, 0x75,
	0x6d, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e,
	0x12, 0x22, 0x0a, 0x0c, 0x68, 0x61, 0x73, 0x68, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x68, 0x61, 0x73, 0x68, 0x46, 0x75, 0x6e, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x83, 0x01, 0x0a, 0x0c, 0x53, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e,
	0x67, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x45, 0x0a, 0x11, 0x53, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x52,
	0x75, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x73, 0x70, 0x71, 0x72, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x75,
	0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x11, 0x53, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e,
	0x67, 0x52, 0x75, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x42, 0x0a, 0x16, 0x41, 0x64,
	0x64, 0x53, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x70, 0x71, 0x72, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x64,
	0x69, 0x6e, 0x67, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x22, 0x16,
	0x0a, 0x14, 0x41, 0x64, 0x64, 0x53, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x75, 0x6c,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x19, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x68,
	0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x41, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e,
	0x67, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x28, 0x0a, 0x05, 0x72, 0x75,
	0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x70, 0x71, 0x72,
	0x2e, 0x53, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x05, 0x72,
	0x75, 0x6c, 0x65, 0x73, 0x22, 0x29, 0x0a, 0x17, 0x44, 0x72, 0x6f, 0x70, 0x53, 0x68, 0x61, 0x72,
	0x64, 0x69, 0x6e, 0x67, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x17, 0x0a, 0x15, 0x44, 0x72, 0x6f, 0x70, 0x53, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x52,
	0x75, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0x8c, 0x02, 0x0a, 0x14, 0x53, 0x68, 0x61,
	0x72, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x4e, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x53, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67,
	0x52, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x1c, 0x2e, 0x73, 0x70, 0x71, 0x72, 0x2e, 0x41, 0x64, 0x64,
	0x53, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x70, 0x71, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x53, 0x68,
	0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x12, 0x51, 0x0a, 0x11, 0x44, 0x72, 0x6f, 0x70, 0x53, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e,
	0x67, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x1d, 0x2e, 0x73, 0x70, 0x71, 0x72, 0x2e, 0x44, 0x72,
	0x6f, 0x70, 0x53, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x70, 0x71, 0x72, 0x2e, 0x44, 0x72, 0x6f,
	0x70, 0x53, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x68, 0x61, 0x72,
	0x64, 0x69, 0x6e, 0x67, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x1d, 0x2e, 0x73, 0x70, 0x71, 0x72,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x75, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x70, 0x71, 0x72, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x53, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x75, 0x6c, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x73, 0x70, 0x71, 0x72, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_sharding_rules_proto_rawDescOnce sync.Once
	file_protos_sharding_rules_proto_rawDescData = file_protos_sharding_rules_proto_rawDesc
)

func file_protos_sharding_rules_proto_rawDescGZIP() []byte {
	file_protos_sharding_rules_proto_rawDescOnce.Do(func() {
		file_protos_sharding_rules_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_sharding_rules_proto_rawDescData)
	})
	return file_protos_sharding_rules_proto_rawDescData
}

var file_protos_sharding_rules_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_protos_sharding_rules_proto_goTypes = []interface{}{
	(*ShardingRuleEntry)(nil),       // 0: spqr.ShardingRuleEntry
	(*ShardingRule)(nil),            // 1: spqr.ShardingRule
	(*AddShardingRuleRequest)(nil),  // 2: spqr.AddShardingRuleRequest
	(*AddShardingRuleReply)(nil),    // 3: spqr.AddShardingRuleReply
	(*ListShardingRuleRequest)(nil), // 4: spqr.ListShardingRuleRequest
	(*ListShardingRuleReply)(nil),   // 5: spqr.ListShardingRuleReply
	(*DropShardingRuleRequest)(nil), // 6: spqr.DropShardingRuleRequest
	(*DropShardingRuleReply)(nil),   // 7: spqr.DropShardingRuleReply
}
var file_protos_sharding_rules_proto_depIdxs = []int32{
	0, // 0: spqr.ShardingRule.ShardingRuleEntry:type_name -> spqr.ShardingRuleEntry
	1, // 1: spqr.AddShardingRuleRequest.rules:type_name -> spqr.ShardingRule
	1, // 2: spqr.ListShardingRuleReply.rules:type_name -> spqr.ShardingRule
	2, // 3: spqr.ShardingRulesService.AddShardingRules:input_type -> spqr.AddShardingRuleRequest
	6, // 4: spqr.ShardingRulesService.DropShardingRules:input_type -> spqr.DropShardingRuleRequest
	4, // 5: spqr.ShardingRulesService.ListShardingRules:input_type -> spqr.ListShardingRuleRequest
	3, // 6: spqr.ShardingRulesService.AddShardingRules:output_type -> spqr.AddShardingRuleReply
	7, // 7: spqr.ShardingRulesService.DropShardingRules:output_type -> spqr.DropShardingRuleReply
	5, // 8: spqr.ShardingRulesService.ListShardingRules:output_type -> spqr.ListShardingRuleReply
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_protos_sharding_rules_proto_init() }
func file_protos_sharding_rules_proto_init() {
	if File_protos_sharding_rules_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_sharding_rules_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShardingRuleEntry); i {
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
		file_protos_sharding_rules_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShardingRule); i {
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
		file_protos_sharding_rules_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddShardingRuleRequest); i {
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
		file_protos_sharding_rules_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddShardingRuleReply); i {
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
		file_protos_sharding_rules_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListShardingRuleRequest); i {
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
		file_protos_sharding_rules_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListShardingRuleReply); i {
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
		file_protos_sharding_rules_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DropShardingRuleRequest); i {
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
		file_protos_sharding_rules_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DropShardingRuleReply); i {
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
			RawDescriptor: file_protos_sharding_rules_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_sharding_rules_proto_goTypes,
		DependencyIndexes: file_protos_sharding_rules_proto_depIdxs,
		MessageInfos:      file_protos_sharding_rules_proto_msgTypes,
	}.Build()
	File_protos_sharding_rules_proto = out.File
	file_protos_sharding_rules_proto_rawDesc = nil
	file_protos_sharding_rules_proto_goTypes = nil
	file_protos_sharding_rules_proto_depIdxs = nil
}
