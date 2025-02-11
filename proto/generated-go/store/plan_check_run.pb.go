// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: store/plan_check_run.proto

package store

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

type PlanCheckRunResult_Result_Status int32

const (
	PlanCheckRunResult_Result_STATUS_UNSPECIFIED PlanCheckRunResult_Result_Status = 0
	PlanCheckRunResult_Result_ERROR              PlanCheckRunResult_Result_Status = 1
	PlanCheckRunResult_Result_WARNING            PlanCheckRunResult_Result_Status = 2
	PlanCheckRunResult_Result_SUCCESS            PlanCheckRunResult_Result_Status = 3
)

// Enum value maps for PlanCheckRunResult_Result_Status.
var (
	PlanCheckRunResult_Result_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "ERROR",
		2: "WARNING",
		3: "SUCCESS",
	}
	PlanCheckRunResult_Result_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"ERROR":              1,
		"WARNING":            2,
		"SUCCESS":            3,
	}
)

func (x PlanCheckRunResult_Result_Status) Enum() *PlanCheckRunResult_Result_Status {
	p := new(PlanCheckRunResult_Result_Status)
	*p = x
	return p
}

func (x PlanCheckRunResult_Result_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlanCheckRunResult_Result_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_store_plan_check_run_proto_enumTypes[0].Descriptor()
}

func (PlanCheckRunResult_Result_Status) Type() protoreflect.EnumType {
	return &file_store_plan_check_run_proto_enumTypes[0]
}

func (x PlanCheckRunResult_Result_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlanCheckRunResult_Result_Status.Descriptor instead.
func (PlanCheckRunResult_Result_Status) EnumDescriptor() ([]byte, []int) {
	return file_store_plan_check_run_proto_rawDescGZIP(), []int{1, 0, 0}
}

type PlanCheckRunConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SheetId    int32 `protobuf:"varint,1,opt,name=sheet_id,json=sheetId,proto3" json:"sheet_id,omitempty"`
	DatabaseId int32 `protobuf:"varint,2,opt,name=database_id,json=databaseId,proto3" json:"database_id,omitempty"`
}

func (x *PlanCheckRunConfig) Reset() {
	*x = PlanCheckRunConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_plan_check_run_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlanCheckRunConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlanCheckRunConfig) ProtoMessage() {}

func (x *PlanCheckRunConfig) ProtoReflect() protoreflect.Message {
	mi := &file_store_plan_check_run_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlanCheckRunConfig.ProtoReflect.Descriptor instead.
func (*PlanCheckRunConfig) Descriptor() ([]byte, []int) {
	return file_store_plan_check_run_proto_rawDescGZIP(), []int{0}
}

func (x *PlanCheckRunConfig) GetSheetId() int32 {
	if x != nil {
		return x.SheetId
	}
	return 0
}

func (x *PlanCheckRunConfig) GetDatabaseId() int32 {
	if x != nil {
		return x.DatabaseId
	}
	return 0
}

type PlanCheckRunResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*PlanCheckRunResult_Result `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	Error   string                       `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *PlanCheckRunResult) Reset() {
	*x = PlanCheckRunResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_plan_check_run_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlanCheckRunResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlanCheckRunResult) ProtoMessage() {}

func (x *PlanCheckRunResult) ProtoReflect() protoreflect.Message {
	mi := &file_store_plan_check_run_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlanCheckRunResult.ProtoReflect.Descriptor instead.
func (*PlanCheckRunResult) Descriptor() ([]byte, []int) {
	return file_store_plan_check_run_proto_rawDescGZIP(), []int{1}
}

func (x *PlanCheckRunResult) GetResults() []*PlanCheckRunResult_Result {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *PlanCheckRunResult) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type PlanCheckRunResult_Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  PlanCheckRunResult_Result_Status `protobuf:"varint,1,opt,name=status,proto3,enum=bytebase.store.PlanCheckRunResult_Result_Status" json:"status,omitempty"`
	Title   string                           `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content string                           `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Code    int64                            `protobuf:"varint,4,opt,name=code,proto3" json:"code,omitempty"`
	// Types that are assignable to Report:
	//
	//	*PlanCheckRunResult_Result_SqlSummaryReport_
	//	*PlanCheckRunResult_Result_SqlReviewReport_
	Report isPlanCheckRunResult_Result_Report `protobuf_oneof:"report"`
}

func (x *PlanCheckRunResult_Result) Reset() {
	*x = PlanCheckRunResult_Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_plan_check_run_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlanCheckRunResult_Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlanCheckRunResult_Result) ProtoMessage() {}

func (x *PlanCheckRunResult_Result) ProtoReflect() protoreflect.Message {
	mi := &file_store_plan_check_run_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlanCheckRunResult_Result.ProtoReflect.Descriptor instead.
func (*PlanCheckRunResult_Result) Descriptor() ([]byte, []int) {
	return file_store_plan_check_run_proto_rawDescGZIP(), []int{1, 0}
}

func (x *PlanCheckRunResult_Result) GetStatus() PlanCheckRunResult_Result_Status {
	if x != nil {
		return x.Status
	}
	return PlanCheckRunResult_Result_STATUS_UNSPECIFIED
}

func (x *PlanCheckRunResult_Result) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PlanCheckRunResult_Result) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *PlanCheckRunResult_Result) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (m *PlanCheckRunResult_Result) GetReport() isPlanCheckRunResult_Result_Report {
	if m != nil {
		return m.Report
	}
	return nil
}

func (x *PlanCheckRunResult_Result) GetSqlSummaryReport() *PlanCheckRunResult_Result_SqlSummaryReport {
	if x, ok := x.GetReport().(*PlanCheckRunResult_Result_SqlSummaryReport_); ok {
		return x.SqlSummaryReport
	}
	return nil
}

func (x *PlanCheckRunResult_Result) GetSqlReviewReport() *PlanCheckRunResult_Result_SqlReviewReport {
	if x, ok := x.GetReport().(*PlanCheckRunResult_Result_SqlReviewReport_); ok {
		return x.SqlReviewReport
	}
	return nil
}

type isPlanCheckRunResult_Result_Report interface {
	isPlanCheckRunResult_Result_Report()
}

type PlanCheckRunResult_Result_SqlSummaryReport_ struct {
	SqlSummaryReport *PlanCheckRunResult_Result_SqlSummaryReport `protobuf:"bytes,5,opt,name=sql_summary_report,json=sqlSummaryReport,proto3,oneof"`
}

type PlanCheckRunResult_Result_SqlReviewReport_ struct {
	SqlReviewReport *PlanCheckRunResult_Result_SqlReviewReport `protobuf:"bytes,6,opt,name=sql_review_report,json=sqlReviewReport,proto3,oneof"`
}

func (*PlanCheckRunResult_Result_SqlSummaryReport_) isPlanCheckRunResult_Result_Report() {}

func (*PlanCheckRunResult_Result_SqlReviewReport_) isPlanCheckRunResult_Result_Report() {}

type PlanCheckRunResult_Result_SqlSummaryReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatementType string `protobuf:"bytes,1,opt,name=statement_type,json=statementType,proto3" json:"statement_type,omitempty"`
	AffectedRows  int64  `protobuf:"varint,2,opt,name=affected_rows,json=affectedRows,proto3" json:"affected_rows,omitempty"`
}

func (x *PlanCheckRunResult_Result_SqlSummaryReport) Reset() {
	*x = PlanCheckRunResult_Result_SqlSummaryReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_plan_check_run_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlanCheckRunResult_Result_SqlSummaryReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlanCheckRunResult_Result_SqlSummaryReport) ProtoMessage() {}

func (x *PlanCheckRunResult_Result_SqlSummaryReport) ProtoReflect() protoreflect.Message {
	mi := &file_store_plan_check_run_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlanCheckRunResult_Result_SqlSummaryReport.ProtoReflect.Descriptor instead.
func (*PlanCheckRunResult_Result_SqlSummaryReport) Descriptor() ([]byte, []int) {
	return file_store_plan_check_run_proto_rawDescGZIP(), []int{1, 0, 0}
}

func (x *PlanCheckRunResult_Result_SqlSummaryReport) GetStatementType() string {
	if x != nil {
		return x.StatementType
	}
	return ""
}

func (x *PlanCheckRunResult_Result_SqlSummaryReport) GetAffectedRows() int64 {
	if x != nil {
		return x.AffectedRows
	}
	return 0
}

type PlanCheckRunResult_Result_SqlReviewReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Line   int64  `protobuf:"varint,1,opt,name=line,proto3" json:"line,omitempty"`
	Detail string `protobuf:"bytes,2,opt,name=detail,proto3" json:"detail,omitempty"`
	// Code from sql review.
	Code int64 `protobuf:"varint,3,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *PlanCheckRunResult_Result_SqlReviewReport) Reset() {
	*x = PlanCheckRunResult_Result_SqlReviewReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_plan_check_run_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlanCheckRunResult_Result_SqlReviewReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlanCheckRunResult_Result_SqlReviewReport) ProtoMessage() {}

func (x *PlanCheckRunResult_Result_SqlReviewReport) ProtoReflect() protoreflect.Message {
	mi := &file_store_plan_check_run_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlanCheckRunResult_Result_SqlReviewReport.ProtoReflect.Descriptor instead.
func (*PlanCheckRunResult_Result_SqlReviewReport) Descriptor() ([]byte, []int) {
	return file_store_plan_check_run_proto_rawDescGZIP(), []int{1, 0, 1}
}

func (x *PlanCheckRunResult_Result_SqlReviewReport) GetLine() int64 {
	if x != nil {
		return x.Line
	}
	return 0
}

func (x *PlanCheckRunResult_Result_SqlReviewReport) GetDetail() string {
	if x != nil {
		return x.Detail
	}
	return ""
}

func (x *PlanCheckRunResult_Result_SqlReviewReport) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

var File_store_plan_check_run_proto protoreflect.FileDescriptor

var file_store_plan_check_run_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x5f, 0x72, 0x75, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x62, 0x79,
	0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x22, 0x50, 0x0a, 0x12,
	0x50, 0x6c, 0x61, 0x6e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x75, 0x6e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x68, 0x65, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x68, 0x65, 0x65, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x49, 0x64, 0x22, 0xe1,
	0x05, 0x0a, 0x12, 0x50, 0x6c, 0x61, 0x6e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x75, 0x6e, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x43, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x1a, 0xef, 0x04, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x48, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e, 0x62, 0x79,
	0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x6c, 0x61,
	0x6e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x6a, 0x0a, 0x12, 0x73, 0x71, 0x6c,
	0x5f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x52, 0x75, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x2e, 0x53, 0x71, 0x6c, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x48, 0x00, 0x52, 0x10, 0x73, 0x71, 0x6c, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x67, 0x0a, 0x11, 0x73, 0x71, 0x6c, 0x5f, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x39, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x75, 0x6e, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x53, 0x71, 0x6c, 0x52,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x48, 0x00, 0x52, 0x0f, 0x73,
	0x71, 0x6c, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x1a, 0x5e,
	0x0a, 0x10, 0x53, 0x71, 0x6c, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x66, 0x66,
	0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0c, 0x61, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x52, 0x6f, 0x77, 0x73, 0x1a, 0x51,
	0x0a, 0x0f, 0x53, 0x71, 0x6c, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x22, 0x45, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x0b,
	0x0a, 0x07, 0x57, 0x41, 0x52, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x53,
	0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x03, 0x42, 0x08, 0x0a, 0x06, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x42, 0x14, 0x5a, 0x12, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2d,
	0x67, 0x6f, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_store_plan_check_run_proto_rawDescOnce sync.Once
	file_store_plan_check_run_proto_rawDescData = file_store_plan_check_run_proto_rawDesc
)

func file_store_plan_check_run_proto_rawDescGZIP() []byte {
	file_store_plan_check_run_proto_rawDescOnce.Do(func() {
		file_store_plan_check_run_proto_rawDescData = protoimpl.X.CompressGZIP(file_store_plan_check_run_proto_rawDescData)
	})
	return file_store_plan_check_run_proto_rawDescData
}

var file_store_plan_check_run_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_store_plan_check_run_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_store_plan_check_run_proto_goTypes = []interface{}{
	(PlanCheckRunResult_Result_Status)(0),              // 0: bytebase.store.PlanCheckRunResult.Result.Status
	(*PlanCheckRunConfig)(nil),                         // 1: bytebase.store.PlanCheckRunConfig
	(*PlanCheckRunResult)(nil),                         // 2: bytebase.store.PlanCheckRunResult
	(*PlanCheckRunResult_Result)(nil),                  // 3: bytebase.store.PlanCheckRunResult.Result
	(*PlanCheckRunResult_Result_SqlSummaryReport)(nil), // 4: bytebase.store.PlanCheckRunResult.Result.SqlSummaryReport
	(*PlanCheckRunResult_Result_SqlReviewReport)(nil),  // 5: bytebase.store.PlanCheckRunResult.Result.SqlReviewReport
}
var file_store_plan_check_run_proto_depIdxs = []int32{
	3, // 0: bytebase.store.PlanCheckRunResult.results:type_name -> bytebase.store.PlanCheckRunResult.Result
	0, // 1: bytebase.store.PlanCheckRunResult.Result.status:type_name -> bytebase.store.PlanCheckRunResult.Result.Status
	4, // 2: bytebase.store.PlanCheckRunResult.Result.sql_summary_report:type_name -> bytebase.store.PlanCheckRunResult.Result.SqlSummaryReport
	5, // 3: bytebase.store.PlanCheckRunResult.Result.sql_review_report:type_name -> bytebase.store.PlanCheckRunResult.Result.SqlReviewReport
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_store_plan_check_run_proto_init() }
func file_store_plan_check_run_proto_init() {
	if File_store_plan_check_run_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_store_plan_check_run_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlanCheckRunConfig); i {
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
		file_store_plan_check_run_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlanCheckRunResult); i {
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
		file_store_plan_check_run_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlanCheckRunResult_Result); i {
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
		file_store_plan_check_run_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlanCheckRunResult_Result_SqlSummaryReport); i {
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
		file_store_plan_check_run_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlanCheckRunResult_Result_SqlReviewReport); i {
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
	file_store_plan_check_run_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*PlanCheckRunResult_Result_SqlSummaryReport_)(nil),
		(*PlanCheckRunResult_Result_SqlReviewReport_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_store_plan_check_run_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_store_plan_check_run_proto_goTypes,
		DependencyIndexes: file_store_plan_check_run_proto_depIdxs,
		EnumInfos:         file_store_plan_check_run_proto_enumTypes,
		MessageInfos:      file_store_plan_check_run_proto_msgTypes,
	}.Build()
	File_store_plan_check_run_proto = out.File
	file_store_plan_check_run_proto_rawDesc = nil
	file_store_plan_check_run_proto_goTypes = nil
	file_store_plan_check_run_proto_depIdxs = nil
}
