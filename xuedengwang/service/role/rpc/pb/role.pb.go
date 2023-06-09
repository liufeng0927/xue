// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: role.proto

package pb

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

type RoleInterface struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RoleInterface) Reset() {
	*x = RoleInterface{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleInterface) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleInterface) ProtoMessage() {}

func (x *RoleInterface) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleInterface.ProtoReflect.Descriptor instead.
func (*RoleInterface) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{0}
}

type RoleDao struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreateTime string `protobuf:"bytes,1,opt,name=createTime,proto3" json:"createTime,omitempty"`
	CreateBy   int64  `protobuf:"varint,2,opt,name=createBy,proto3" json:"createBy,omitempty"`
	UpdateTime string `protobuf:"bytes,3,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
	UpdateBy   int64  `protobuf:"varint,4,opt,name=updateBy,proto3" json:"updateBy,omitempty"`
	Remarks    string `protobuf:"bytes,5,opt,name=remarks,proto3" json:"remarks,omitempty"`
	ID         int64  `protobuf:"varint,6,opt,name=ID,proto3" json:"ID,omitempty"`
	Name       string `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	Code       string `protobuf:"bytes,8,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *RoleDao) Reset() {
	*x = RoleDao{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleDao) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleDao) ProtoMessage() {}

func (x *RoleDao) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleDao.ProtoReflect.Descriptor instead.
func (*RoleDao) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{1}
}

func (x *RoleDao) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *RoleDao) GetCreateBy() int64 {
	if x != nil {
		return x.CreateBy
	}
	return 0
}

func (x *RoleDao) GetUpdateTime() string {
	if x != nil {
		return x.UpdateTime
	}
	return ""
}

func (x *RoleDao) GetUpdateBy() int64 {
	if x != nil {
		return x.UpdateBy
	}
	return 0
}

func (x *RoleDao) GetRemarks() string {
	if x != nil {
		return x.Remarks
	}
	return ""
}

func (x *RoleDao) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *RoleDao) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RoleDao) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type RoleAll struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID   int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *RoleAll) Reset() {
	*x = RoleAll{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleAll) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleAll) ProtoMessage() {}

func (x *RoleAll) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleAll.ProtoReflect.Descriptor instead.
func (*RoleAll) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{2}
}

func (x *RoleAll) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *RoleAll) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type RoleAllResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleAll []*RoleAll `protobuf:"bytes,1,rep,name=roleAll,proto3" json:"roleAll,omitempty"`
}

func (x *RoleAllResp) Reset() {
	*x = RoleAllResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleAllResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleAllResp) ProtoMessage() {}

func (x *RoleAllResp) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleAllResp.ProtoReflect.Descriptor instead.
func (*RoleAllResp) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{3}
}

func (x *RoleAllResp) GetRoleAll() []*RoleAll {
	if x != nil {
		return x.RoleAll
	}
	return nil
}

type GetRoleReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageIndex   int32  `protobuf:"varint,1,opt,name=pageIndex,proto3" json:"pageIndex,omitempty"`
	PageSize    int32  `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	SearchValue string `protobuf:"bytes,3,opt,name=searchValue,proto3" json:"searchValue,omitempty"`
}

func (x *GetRoleReq) Reset() {
	*x = GetRoleReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRoleReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoleReq) ProtoMessage() {}

func (x *GetRoleReq) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoleReq.ProtoReflect.Descriptor instead.
func (*GetRoleReq) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{4}
}

func (x *GetRoleReq) GetPageIndex() int32 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *GetRoleReq) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *GetRoleReq) GetSearchValue() string {
	if x != nil {
		return x.SearchValue
	}
	return ""
}

type GetRoleResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content       []*RoleDao `protobuf:"bytes,1,rep,name=content,proto3" json:"content,omitempty"`
	TotalElements int64      `protobuf:"varint,2,opt,name=totalElements,proto3" json:"totalElements,omitempty"`
}

func (x *GetRoleResp) Reset() {
	*x = GetRoleResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRoleResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoleResp) ProtoMessage() {}

func (x *GetRoleResp) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoleResp.ProtoReflect.Descriptor instead.
func (*GetRoleResp) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{5}
}

func (x *GetRoleResp) GetContent() []*RoleDao {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *GetRoleResp) GetTotalElements() int64 {
	if x != nil {
		return x.TotalElements
	}
	return 0
}

type GetRoleByIDReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *GetRoleByIDReq) Reset() {
	*x = GetRoleByIDReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRoleByIDReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoleByIDReq) ProtoMessage() {}

func (x *GetRoleByIDReq) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoleByIDReq.ProtoReflect.Descriptor instead.
func (*GetRoleByIDReq) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{6}
}

func (x *GetRoleByIDReq) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type GetRoleByIDResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreateTime string `protobuf:"bytes,1,opt,name=createTime,proto3" json:"createTime,omitempty"`
	CreateBy   int64  `protobuf:"varint,2,opt,name=createBy,proto3" json:"createBy,omitempty"`
	UpdateTime string `protobuf:"bytes,3,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
	UpdateBy   int64  `protobuf:"varint,4,opt,name=updateBy,proto3" json:"updateBy,omitempty"`
	Remarks    string `protobuf:"bytes,5,opt,name=remarks,proto3" json:"remarks,omitempty"`
	ID         int64  `protobuf:"varint,6,opt,name=ID,proto3" json:"ID,omitempty"`
	Name       string `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	Code       string `protobuf:"bytes,8,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *GetRoleByIDResp) Reset() {
	*x = GetRoleByIDResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRoleByIDResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoleByIDResp) ProtoMessage() {}

func (x *GetRoleByIDResp) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoleByIDResp.ProtoReflect.Descriptor instead.
func (*GetRoleByIDResp) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{7}
}

func (x *GetRoleByIDResp) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *GetRoleByIDResp) GetCreateBy() int64 {
	if x != nil {
		return x.CreateBy
	}
	return 0
}

func (x *GetRoleByIDResp) GetUpdateTime() string {
	if x != nil {
		return x.UpdateTime
	}
	return ""
}

func (x *GetRoleByIDResp) GetUpdateBy() int64 {
	if x != nil {
		return x.UpdateBy
	}
	return 0
}

func (x *GetRoleByIDResp) GetRemarks() string {
	if x != nil {
		return x.Remarks
	}
	return ""
}

func (x *GetRoleByIDResp) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *GetRoleByIDResp) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetRoleByIDResp) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type UpdateRoleReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code       string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	RoleID     int64  `protobuf:"varint,2,opt,name=roleID,proto3" json:"roleID,omitempty"`
	Name       string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Remarks    string `protobuf:"bytes,4,opt,name=remarks,proto3" json:"remarks,omitempty"`
	UpdateByID int64  `protobuf:"varint,5,opt,name=updateByID,proto3" json:"updateByID,omitempty"`
}

func (x *UpdateRoleReq) Reset() {
	*x = UpdateRoleReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRoleReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRoleReq) ProtoMessage() {}

func (x *UpdateRoleReq) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRoleReq.ProtoReflect.Descriptor instead.
func (*UpdateRoleReq) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateRoleReq) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *UpdateRoleReq) GetRoleID() int64 {
	if x != nil {
		return x.RoleID
	}
	return 0
}

func (x *UpdateRoleReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateRoleReq) GetRemarks() string {
	if x != nil {
		return x.Remarks
	}
	return ""
}

func (x *UpdateRoleReq) GetUpdateByID() int64 {
	if x != nil {
		return x.UpdateByID
	}
	return 0
}

type AddRoleReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Remarks string `protobuf:"bytes,4,opt,name=remarks,proto3" json:"remarks,omitempty"`
	ByID    int64  `protobuf:"varint,5,opt,name=byID,proto3" json:"byID,omitempty"`
}

func (x *AddRoleReq) Reset() {
	*x = AddRoleReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRoleReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRoleReq) ProtoMessage() {}

func (x *AddRoleReq) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRoleReq.ProtoReflect.Descriptor instead.
func (*AddRoleReq) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{9}
}

func (x *AddRoleReq) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *AddRoleReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddRoleReq) GetRemarks() string {
	if x != nil {
		return x.Remarks
	}
	return ""
}

func (x *AddRoleReq) GetByID() int64 {
	if x != nil {
		return x.ByID
	}
	return 0
}

type DeleteRoleByIDReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleID int64 `protobuf:"varint,1,opt,name=roleID,proto3" json:"roleID,omitempty"`
}

func (x *DeleteRoleByIDReq) Reset() {
	*x = DeleteRoleByIDReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRoleByIDReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRoleByIDReq) ProtoMessage() {}

func (x *DeleteRoleByIDReq) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRoleByIDReq.ProtoReflect.Descriptor instead.
func (*DeleteRoleByIDReq) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteRoleByIDReq) GetRoleID() int64 {
	if x != nil {
		return x.RoleID
	}
	return 0
}

var File_role_proto protoreflect.FileDescriptor

var file_role_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x22, 0x0f, 0x0a, 0x0d, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x22, 0xd3, 0x01, 0x0a, 0x07, 0x52, 0x6f, 0x6c, 0x65, 0x44, 0x61, 0x6f, 0x12, 0x1e, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x42, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x2d, 0x0a, 0x07, 0x52, 0x6f, 0x6c, 0x65, 0x41,
	0x6c, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x34, 0x0a, 0x0b, 0x52, 0x6f, 0x6c, 0x65, 0x41, 0x6c,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x12, 0x25, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x41, 0x6c, 0x6c,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x6f, 0x6c, 0x65,
	0x41, 0x6c, 0x6c, 0x52, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x41, 0x6c, 0x6c, 0x22, 0x68, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x5a, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x25, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x6f, 0x6c, 0x65,
	0x44, 0x61, 0x6f, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0d,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x79, 0x49,
	0x44, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x49, 0x44, 0x22, 0xdb, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65,
	0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x42, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x42, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x22, 0x89, 0x01, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6c, 0x65,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x44,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x12, 0x1e,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x49, 0x44, 0x22, 0x62,
	0x0a, 0x0a, 0x41, 0x64, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x62, 0x79, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x62, 0x79,
	0x49, 0x44, 0x22, 0x2b, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65,
	0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x44, 0x32,
	0xb7, 0x02, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65,
	0x41, 0x6c, 0x6c, 0x12, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x1a, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x6f, 0x6c, 0x65,
	0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2a, 0x0a, 0x07, 0x67, 0x65, 0x74, 0x52, 0x6f,
	0x6c, 0x65, 0x12, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x36, 0x0a, 0x0b, 0x67, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x79,
	0x49, 0x44, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x42,
	0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x52,
	0x6f, 0x6c, 0x65, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x12, 0x32, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x70,
	0x62, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12,
	0x2c, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0e, 0x2e, 0x70, 0x62, 0x2e,
	0x41, 0x64, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x70, 0x62, 0x2e,
	0x52, 0x6f, 0x6c, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x3a, 0x0a,
	0x0e, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x79, 0x49, 0x44, 0x12,
	0x15, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x42,
	0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x6f, 0x6c, 0x65,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_role_proto_rawDescOnce sync.Once
	file_role_proto_rawDescData = file_role_proto_rawDesc
)

func file_role_proto_rawDescGZIP() []byte {
	file_role_proto_rawDescOnce.Do(func() {
		file_role_proto_rawDescData = protoimpl.X.CompressGZIP(file_role_proto_rawDescData)
	})
	return file_role_proto_rawDescData
}

var file_role_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_role_proto_goTypes = []interface{}{
	(*RoleInterface)(nil),     // 0: pb.RoleInterface
	(*RoleDao)(nil),           // 1: pb.RoleDao
	(*RoleAll)(nil),           // 2: pb.RoleAll
	(*RoleAllResp)(nil),       // 3: pb.RoleAllResp
	(*GetRoleReq)(nil),        // 4: pb.GetRoleReq
	(*GetRoleResp)(nil),       // 5: pb.GetRoleResp
	(*GetRoleByIDReq)(nil),    // 6: pb.GetRoleByIDReq
	(*GetRoleByIDResp)(nil),   // 7: pb.GetRoleByIDResp
	(*UpdateRoleReq)(nil),     // 8: pb.UpdateRoleReq
	(*AddRoleReq)(nil),        // 9: pb.AddRoleReq
	(*DeleteRoleByIDReq)(nil), // 10: pb.DeleteRoleByIDReq
}
var file_role_proto_depIdxs = []int32{
	2,  // 0: pb.RoleAllResp.roleAll:type_name -> pb.RoleAll
	1,  // 1: pb.GetRoleResp.content:type_name -> pb.RoleDao
	0,  // 2: pb.role.roleAll:input_type -> pb.RoleInterface
	4,  // 3: pb.role.getRole:input_type -> pb.GetRoleReq
	6,  // 4: pb.role.getRoleByID:input_type -> pb.GetRoleByIDReq
	8,  // 5: pb.role.updateRole:input_type -> pb.UpdateRoleReq
	9,  // 6: pb.role.addRole:input_type -> pb.AddRoleReq
	10, // 7: pb.role.deleteRoleByID:input_type -> pb.DeleteRoleByIDReq
	3,  // 8: pb.role.roleAll:output_type -> pb.RoleAllResp
	5,  // 9: pb.role.getRole:output_type -> pb.GetRoleResp
	7,  // 10: pb.role.getRoleByID:output_type -> pb.GetRoleByIDResp
	0,  // 11: pb.role.updateRole:output_type -> pb.RoleInterface
	0,  // 12: pb.role.addRole:output_type -> pb.RoleInterface
	0,  // 13: pb.role.deleteRoleByID:output_type -> pb.RoleInterface
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_role_proto_init() }
func file_role_proto_init() {
	if File_role_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_role_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleInterface); i {
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
		file_role_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleDao); i {
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
		file_role_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleAll); i {
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
		file_role_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleAllResp); i {
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
		file_role_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRoleReq); i {
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
		file_role_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRoleResp); i {
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
		file_role_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRoleByIDReq); i {
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
		file_role_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRoleByIDResp); i {
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
		file_role_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRoleReq); i {
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
		file_role_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRoleReq); i {
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
		file_role_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRoleByIDReq); i {
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
			RawDescriptor: file_role_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_role_proto_goTypes,
		DependencyIndexes: file_role_proto_depIdxs,
		MessageInfos:      file_role_proto_msgTypes,
	}.Build()
	File_role_proto = out.File
	file_role_proto_rawDesc = nil
	file_role_proto_goTypes = nil
	file_role_proto_depIdxs = nil
}
