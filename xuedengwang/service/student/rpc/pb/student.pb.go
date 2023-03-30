// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: student.proto

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

type StudentInterface struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StudentInterface) Reset() {
	*x = StudentInterface{}
	if protoimpl.UnsafeEnabled {
		mi := &file_student_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StudentInterface) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StudentInterface) ProtoMessage() {}

func (x *StudentInterface) ProtoReflect() protoreflect.Message {
	mi := &file_student_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StudentInterface.ProtoReflect.Descriptor instead.
func (*StudentInterface) Descriptor() ([]byte, []int) {
	return file_student_proto_rawDescGZIP(), []int{0}
}

type StudentDao struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreateTime     string `protobuf:"bytes,1,opt,name=createTime,proto3" json:"createTime,omitempty"`
	CreateBy       int64  `protobuf:"varint,2,opt,name=createBy,proto3" json:"createBy,omitempty"`
	UpdateTime     string `protobuf:"bytes,3,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
	UpdateBy       int64  `protobuf:"varint,4,opt,name=updateBy,proto3" json:"updateBy,omitempty"`
	Remarks        string `protobuf:"bytes,5,opt,name=remarks,proto3" json:"remarks,omitempty"`
	ID             int64  `protobuf:"varint,6,opt,name=ID,proto3" json:"ID,omitempty"`
	Stuno          string `protobuf:"bytes,7,opt,name=stuno,proto3" json:"stuno,omitempty"`
	Sex            string `protobuf:"bytes,8,opt,name=sex,proto3" json:"sex,omitempty"`
	Phone          string `protobuf:"bytes,9,opt,name=phone,proto3" json:"phone,omitempty"`
	QQ             string `protobuf:"bytes,10,opt,name=QQ,proto3" json:"QQ,omitempty"`
	GradeClassName string `protobuf:"bytes,11,opt,name=GradeClassName,proto3" json:"GradeClassName,omitempty"`
	Name           string `protobuf:"bytes,12,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *StudentDao) Reset() {
	*x = StudentDao{}
	if protoimpl.UnsafeEnabled {
		mi := &file_student_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StudentDao) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StudentDao) ProtoMessage() {}

func (x *StudentDao) ProtoReflect() protoreflect.Message {
	mi := &file_student_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StudentDao.ProtoReflect.Descriptor instead.
func (*StudentDao) Descriptor() ([]byte, []int) {
	return file_student_proto_rawDescGZIP(), []int{1}
}

func (x *StudentDao) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *StudentDao) GetCreateBy() int64 {
	if x != nil {
		return x.CreateBy
	}
	return 0
}

func (x *StudentDao) GetUpdateTime() string {
	if x != nil {
		return x.UpdateTime
	}
	return ""
}

func (x *StudentDao) GetUpdateBy() int64 {
	if x != nil {
		return x.UpdateBy
	}
	return 0
}

func (x *StudentDao) GetRemarks() string {
	if x != nil {
		return x.Remarks
	}
	return ""
}

func (x *StudentDao) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *StudentDao) GetStuno() string {
	if x != nil {
		return x.Stuno
	}
	return ""
}

func (x *StudentDao) GetSex() string {
	if x != nil {
		return x.Sex
	}
	return ""
}

func (x *StudentDao) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *StudentDao) GetQQ() string {
	if x != nil {
		return x.QQ
	}
	return ""
}

func (x *StudentDao) GetGradeClassName() string {
	if x != nil {
		return x.GradeClassName
	}
	return ""
}

func (x *StudentDao) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetStudentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageIndex   int32  `protobuf:"varint,1,opt,name=pageIndex,proto3" json:"pageIndex,omitempty"`
	PageSize    int32  `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	SearchValue string `protobuf:"bytes,3,opt,name=searchValue,proto3" json:"searchValue,omitempty"`
}

func (x *GetStudentReq) Reset() {
	*x = GetStudentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_student_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStudentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStudentReq) ProtoMessage() {}

func (x *GetStudentReq) ProtoReflect() protoreflect.Message {
	mi := &file_student_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStudentReq.ProtoReflect.Descriptor instead.
func (*GetStudentReq) Descriptor() ([]byte, []int) {
	return file_student_proto_rawDescGZIP(), []int{2}
}

func (x *GetStudentReq) GetPageIndex() int32 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *GetStudentReq) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *GetStudentReq) GetSearchValue() string {
	if x != nil {
		return x.SearchValue
	}
	return ""
}

type GetStudentResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content       []*StudentDao `protobuf:"bytes,1,rep,name=content,proto3" json:"content,omitempty"`
	TotalElements int64         `protobuf:"varint,2,opt,name=totalElements,proto3" json:"totalElements,omitempty"`
}

func (x *GetStudentResp) Reset() {
	*x = GetStudentResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_student_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStudentResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStudentResp) ProtoMessage() {}

func (x *GetStudentResp) ProtoReflect() protoreflect.Message {
	mi := &file_student_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStudentResp.ProtoReflect.Descriptor instead.
func (*GetStudentResp) Descriptor() ([]byte, []int) {
	return file_student_proto_rawDescGZIP(), []int{3}
}

func (x *GetStudentResp) GetContent() []*StudentDao {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *GetStudentResp) GetTotalElements() int64 {
	if x != nil {
		return x.TotalElements
	}
	return 0
}

type GetStudentByIDReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *GetStudentByIDReq) Reset() {
	*x = GetStudentByIDReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_student_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStudentByIDReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStudentByIDReq) ProtoMessage() {}

func (x *GetStudentByIDReq) ProtoReflect() protoreflect.Message {
	mi := &file_student_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStudentByIDReq.ProtoReflect.Descriptor instead.
func (*GetStudentByIDReq) Descriptor() ([]byte, []int) {
	return file_student_proto_rawDescGZIP(), []int{4}
}

func (x *GetStudentByIDReq) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type GetStudentByIDResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreateTime     string `protobuf:"bytes,1,opt,name=createTime,proto3" json:"createTime,omitempty"`
	CreateBy       int64  `protobuf:"varint,2,opt,name=createBy,proto3" json:"createBy,omitempty"`
	UpdateTime     string `protobuf:"bytes,3,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
	UpdateBy       int64  `protobuf:"varint,4,opt,name=updateBy,proto3" json:"updateBy,omitempty"`
	Remarks        string `protobuf:"bytes,5,opt,name=remarks,proto3" json:"remarks,omitempty"`
	ID             int64  `protobuf:"varint,6,opt,name=ID,proto3" json:"ID,omitempty"`
	Stuno          string `protobuf:"bytes,7,opt,name=stuno,proto3" json:"stuno,omitempty"`
	Sex            string `protobuf:"bytes,8,opt,name=sex,proto3" json:"sex,omitempty"`
	Phone          string `protobuf:"bytes,9,opt,name=phone,proto3" json:"phone,omitempty"`
	QQ             string `protobuf:"bytes,10,opt,name=QQ,proto3" json:"QQ,omitempty"`
	GradeClassID   int64  `protobuf:"varint,11,opt,name=GradeClassID,proto3" json:"GradeClassID,omitempty"`
	Name           string `protobuf:"bytes,12,opt,name=name,proto3" json:"name,omitempty"`
	GradeClassName string `protobuf:"bytes,13,opt,name=GradeClassName,proto3" json:"GradeClassName,omitempty"`
}

func (x *GetStudentByIDResp) Reset() {
	*x = GetStudentByIDResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_student_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStudentByIDResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStudentByIDResp) ProtoMessage() {}

func (x *GetStudentByIDResp) ProtoReflect() protoreflect.Message {
	mi := &file_student_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStudentByIDResp.ProtoReflect.Descriptor instead.
func (*GetStudentByIDResp) Descriptor() ([]byte, []int) {
	return file_student_proto_rawDescGZIP(), []int{5}
}

func (x *GetStudentByIDResp) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *GetStudentByIDResp) GetCreateBy() int64 {
	if x != nil {
		return x.CreateBy
	}
	return 0
}

func (x *GetStudentByIDResp) GetUpdateTime() string {
	if x != nil {
		return x.UpdateTime
	}
	return ""
}

func (x *GetStudentByIDResp) GetUpdateBy() int64 {
	if x != nil {
		return x.UpdateBy
	}
	return 0
}

func (x *GetStudentByIDResp) GetRemarks() string {
	if x != nil {
		return x.Remarks
	}
	return ""
}

func (x *GetStudentByIDResp) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *GetStudentByIDResp) GetStuno() string {
	if x != nil {
		return x.Stuno
	}
	return ""
}

func (x *GetStudentByIDResp) GetSex() string {
	if x != nil {
		return x.Sex
	}
	return ""
}

func (x *GetStudentByIDResp) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *GetStudentByIDResp) GetQQ() string {
	if x != nil {
		return x.QQ
	}
	return ""
}

func (x *GetStudentByIDResp) GetGradeClassID() int64 {
	if x != nil {
		return x.GradeClassID
	}
	return 0
}

func (x *GetStudentByIDResp) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetStudentByIDResp) GetGradeClassName() string {
	if x != nil {
		return x.GradeClassName
	}
	return ""
}

type UpdateStudentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Remarks      string `protobuf:"bytes,1,opt,name=remarks,proto3" json:"remarks,omitempty"`
	ID           int64  `protobuf:"varint,2,opt,name=ID,proto3" json:"ID,omitempty"`
	Stuno        string `protobuf:"bytes,3,opt,name=stuno,proto3" json:"stuno,omitempty"`
	Sex          string `protobuf:"bytes,4,opt,name=sex,proto3" json:"sex,omitempty"`
	Phone        string `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
	QQ           string `protobuf:"bytes,6,opt,name=QQ,proto3" json:"QQ,omitempty"`
	GradeClassID int64  `protobuf:"varint,7,opt,name=GradeClassID,proto3" json:"GradeClassID,omitempty"`
	UpdateByID   int64  `protobuf:"varint,8,opt,name=updateByID,proto3" json:"updateByID,omitempty"`
	Name         string `protobuf:"bytes,9,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *UpdateStudentReq) Reset() {
	*x = UpdateStudentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_student_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStudentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStudentReq) ProtoMessage() {}

func (x *UpdateStudentReq) ProtoReflect() protoreflect.Message {
	mi := &file_student_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStudentReq.ProtoReflect.Descriptor instead.
func (*UpdateStudentReq) Descriptor() ([]byte, []int) {
	return file_student_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateStudentReq) GetRemarks() string {
	if x != nil {
		return x.Remarks
	}
	return ""
}

func (x *UpdateStudentReq) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *UpdateStudentReq) GetStuno() string {
	if x != nil {
		return x.Stuno
	}
	return ""
}

func (x *UpdateStudentReq) GetSex() string {
	if x != nil {
		return x.Sex
	}
	return ""
}

func (x *UpdateStudentReq) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UpdateStudentReq) GetQQ() string {
	if x != nil {
		return x.QQ
	}
	return ""
}

func (x *UpdateStudentReq) GetGradeClassID() int64 {
	if x != nil {
		return x.GradeClassID
	}
	return 0
}

func (x *UpdateStudentReq) GetUpdateByID() int64 {
	if x != nil {
		return x.UpdateByID
	}
	return 0
}

func (x *UpdateStudentReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type AddStudentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Remarks      string `protobuf:"bytes,1,opt,name=remarks,proto3" json:"remarks,omitempty"`
	Stuno        string `protobuf:"bytes,3,opt,name=stuno,proto3" json:"stuno,omitempty"`
	Sex          string `protobuf:"bytes,4,opt,name=sex,proto3" json:"sex,omitempty"`
	Phone        string `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
	QQ           string `protobuf:"bytes,6,opt,name=QQ,proto3" json:"QQ,omitempty"`
	GradeClassID int64  `protobuf:"varint,7,opt,name=GradeClassID,proto3" json:"GradeClassID,omitempty"`
	ByID         int64  `protobuf:"varint,8,opt,name=byID,proto3" json:"byID,omitempty"`
	Name         string `protobuf:"bytes,9,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *AddStudentReq) Reset() {
	*x = AddStudentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_student_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddStudentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddStudentReq) ProtoMessage() {}

func (x *AddStudentReq) ProtoReflect() protoreflect.Message {
	mi := &file_student_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddStudentReq.ProtoReflect.Descriptor instead.
func (*AddStudentReq) Descriptor() ([]byte, []int) {
	return file_student_proto_rawDescGZIP(), []int{7}
}

func (x *AddStudentReq) GetRemarks() string {
	if x != nil {
		return x.Remarks
	}
	return ""
}

func (x *AddStudentReq) GetStuno() string {
	if x != nil {
		return x.Stuno
	}
	return ""
}

func (x *AddStudentReq) GetSex() string {
	if x != nil {
		return x.Sex
	}
	return ""
}

func (x *AddStudentReq) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *AddStudentReq) GetQQ() string {
	if x != nil {
		return x.QQ
	}
	return ""
}

func (x *AddStudentReq) GetGradeClassID() int64 {
	if x != nil {
		return x.GradeClassID
	}
	return 0
}

func (x *AddStudentReq) GetByID() int64 {
	if x != nil {
		return x.ByID
	}
	return 0
}

func (x *AddStudentReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DeleteStudentByIDReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StudentID int64 `protobuf:"varint,1,opt,name=StudentID,proto3" json:"StudentID,omitempty"`
}

func (x *DeleteStudentByIDReq) Reset() {
	*x = DeleteStudentByIDReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_student_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteStudentByIDReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteStudentByIDReq) ProtoMessage() {}

func (x *DeleteStudentByIDReq) ProtoReflect() protoreflect.Message {
	mi := &file_student_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteStudentByIDReq.ProtoReflect.Descriptor instead.
func (*DeleteStudentByIDReq) Descriptor() ([]byte, []int) {
	return file_student_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteStudentByIDReq) GetStudentID() int64 {
	if x != nil {
		return x.StudentID
	}
	return 0
}

var File_student_proto protoreflect.FileDescriptor

var file_student_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x22, 0x12, 0x0a, 0x10, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x22, 0xb8, 0x02, 0x0a, 0x0a, 0x53, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x44, 0x61, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x42, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x42, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x18,
	0x0a, 0x07, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x75, 0x6e,
	0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x75, 0x6e, 0x6f, 0x12, 0x10,
	0x0a, 0x03, 0x73, 0x65, 0x78, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x65, 0x78,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x51, 0x51, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x51, 0x51, 0x12, 0x26, 0x0a, 0x0e, 0x47, 0x72, 0x61, 0x64, 0x65, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x47, 0x72, 0x61, 0x64, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x6b, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x60, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x28, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x44,
	0x61, 0x6f, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x22, 0x23, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x42,
	0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x22, 0xe4, 0x02, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1e, 0x0a,
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
	0x14, 0x0a, 0x05, 0x73, 0x74, 0x75, 0x6e, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x73, 0x74, 0x75, 0x6e, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x78, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x73, 0x65, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x51, 0x51, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x51, 0x51, 0x12, 0x22, 0x0a,
	0x0c, 0x47, 0x72, 0x61, 0x64, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x44, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0c, 0x47, 0x72, 0x61, 0x64, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49,
	0x44, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x47, 0x72, 0x61, 0x64, 0x65, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x47,
	0x72, 0x61, 0x64, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xe2, 0x01,
	0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x74, 0x75, 0x6e, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x75,
	0x6e, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x73, 0x65, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x51, 0x51,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x51, 0x51, 0x12, 0x22, 0x0a, 0x0c, 0x47, 0x72,
	0x61, 0x64, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x44, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0c, 0x47, 0x72, 0x61, 0x64, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x44, 0x12, 0x1e,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x49, 0x44, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x49, 0x44, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0xc3, 0x01, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x74, 0x75, 0x6e, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73,
	0x74, 0x75, 0x6e, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x73, 0x65, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x51, 0x51, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x51, 0x51, 0x12, 0x22, 0x0a, 0x0c,
	0x47, 0x72, 0x61, 0x64, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x44, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0c, 0x47, 0x72, 0x61, 0x64, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x44,
	0x12, 0x12, 0x0a, 0x04, 0x62, 0x79, 0x49, 0x44, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x62, 0x79, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x34, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71,
	0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x32, 0xb8,
	0x02, 0x0a, 0x07, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x33, 0x0a, 0x0a, 0x67, 0x65,
	0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65,
	0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x70, 0x62,
	0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x3f, 0x0a, 0x0e, 0x67, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49,
	0x44, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e,
	0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65,
	0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x3b, 0x0a, 0x0d, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e,
	0x74, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x35, 0x0a,
	0x0a, 0x61, 0x64, 0x64, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x11, 0x2e, 0x70, 0x62,
	0x2e, 0x41, 0x64, 0x64, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x14,
	0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x11, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x44,
	0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_student_proto_rawDescOnce sync.Once
	file_student_proto_rawDescData = file_student_proto_rawDesc
)

func file_student_proto_rawDescGZIP() []byte {
	file_student_proto_rawDescOnce.Do(func() {
		file_student_proto_rawDescData = protoimpl.X.CompressGZIP(file_student_proto_rawDescData)
	})
	return file_student_proto_rawDescData
}

var file_student_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_student_proto_goTypes = []interface{}{
	(*StudentInterface)(nil),     // 0: pb.StudentInterface
	(*StudentDao)(nil),           // 1: pb.StudentDao
	(*GetStudentReq)(nil),        // 2: pb.GetStudentReq
	(*GetStudentResp)(nil),       // 3: pb.GetStudentResp
	(*GetStudentByIDReq)(nil),    // 4: pb.GetStudentByIDReq
	(*GetStudentByIDResp)(nil),   // 5: pb.GetStudentByIDResp
	(*UpdateStudentReq)(nil),     // 6: pb.UpdateStudentReq
	(*AddStudentReq)(nil),        // 7: pb.AddStudentReq
	(*DeleteStudentByIDReq)(nil), // 8: pb.DeleteStudentByIDReq
}
var file_student_proto_depIdxs = []int32{
	1, // 0: pb.GetStudentResp.content:type_name -> pb.StudentDao
	2, // 1: pb.student.getStudent:input_type -> pb.GetStudentReq
	4, // 2: pb.student.getStudentByID:input_type -> pb.GetStudentByIDReq
	6, // 3: pb.student.updateStudent:input_type -> pb.UpdateStudentReq
	7, // 4: pb.student.addStudent:input_type -> pb.AddStudentReq
	8, // 5: pb.student.deleteStudentByID:input_type -> pb.DeleteStudentByIDReq
	3, // 6: pb.student.getStudent:output_type -> pb.GetStudentResp
	5, // 7: pb.student.getStudentByID:output_type -> pb.GetStudentByIDResp
	0, // 8: pb.student.updateStudent:output_type -> pb.StudentInterface
	0, // 9: pb.student.addStudent:output_type -> pb.StudentInterface
	0, // 10: pb.student.deleteStudentByID:output_type -> pb.StudentInterface
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_student_proto_init() }
func file_student_proto_init() {
	if File_student_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_student_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StudentInterface); i {
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
		file_student_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StudentDao); i {
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
		file_student_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStudentReq); i {
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
		file_student_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStudentResp); i {
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
		file_student_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStudentByIDReq); i {
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
		file_student_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStudentByIDResp); i {
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
		file_student_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStudentReq); i {
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
		file_student_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddStudentReq); i {
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
		file_student_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteStudentByIDReq); i {
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
			RawDescriptor: file_student_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_student_proto_goTypes,
		DependencyIndexes: file_student_proto_depIdxs,
		MessageInfos:      file_student_proto_msgTypes,
	}.Build()
	File_student_proto = out.File
	file_student_proto_rawDesc = nil
	file_student_proto_goTypes = nil
	file_student_proto_depIdxs = nil
}
