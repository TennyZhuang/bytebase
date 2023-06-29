// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: v1/iam_policy.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	v1alpha1 "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
	expr "google.golang.org/genproto/googleapis/type/expr"
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

type IamPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Collection of binding.
	// A binding binds one or more project members to a single project role.
	Bindings []*Binding `protobuf:"bytes,1,rep,name=bindings,proto3" json:"bindings,omitempty"`
}

func (x *IamPolicy) Reset() {
	*x = IamPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_iam_policy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IamPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IamPolicy) ProtoMessage() {}

func (x *IamPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_v1_iam_policy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IamPolicy.ProtoReflect.Descriptor instead.
func (*IamPolicy) Descriptor() ([]byte, []int) {
	return file_v1_iam_policy_proto_rawDescGZIP(), []int{0}
}

func (x *IamPolicy) GetBindings() []*Binding {
	if x != nil {
		return x.Bindings
	}
	return nil
}

type Binding struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The role that is assigned to the members.
	// Format: roles/{role}
	Role string `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	// Specifies the principals requesting access for a Bytebase resource.
	Members []string `protobuf:"bytes,2,rep,name=members,proto3" json:"members,omitempty"`
	// The condition that is associated with this binding.
	// If the condition evaluates to true, then this binding applies to the current request.
	// If the condition evaluates to false, then this binding does not apply to the current request. However, a different role binding might grant the same role to one or more of the principals in this binding.
	Condition *expr.Expr `protobuf:"bytes,3,opt,name=condition,proto3" json:"condition,omitempty"`
	// The parsed expression of the condition.
	ParsedExpr *v1alpha1.ParsedExpr `protobuf:"bytes,4,opt,name=parsed_expr,json=parsedExpr,proto3" json:"parsed_expr,omitempty"`
}

func (x *Binding) Reset() {
	*x = Binding{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_iam_policy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Binding) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Binding) ProtoMessage() {}

func (x *Binding) ProtoReflect() protoreflect.Message {
	mi := &file_v1_iam_policy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Binding.ProtoReflect.Descriptor instead.
func (*Binding) Descriptor() ([]byte, []int) {
	return file_v1_iam_policy_proto_rawDescGZIP(), []int{1}
}

func (x *Binding) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *Binding) GetMembers() []string {
	if x != nil {
		return x.Members
	}
	return nil
}

func (x *Binding) GetCondition() *expr.Expr {
	if x != nil {
		return x.Condition
	}
	return nil
}

func (x *Binding) GetParsedExpr() *v1alpha1.ParsedExpr {
	if x != nil {
		return x.ParsedExpr
	}
	return nil
}

var File_v1_iam_policy_proto protoreflect.FileDescriptor

var file_v1_iam_policy_proto_rawDesc = []byte{
	0x0a, 0x13, 0x76, 0x31, 0x2f, 0x69, 0x61, 0x6d, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x76, 0x31, 0x1a, 0x25, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65,
	0x78, 0x70, 0x72, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73, 0x79, 0x6e,
	0x74, 0x61, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61,
	0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x65, 0x78, 0x70, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x3d, 0x0a, 0x09, 0x49, 0x61, 0x6d, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12,
	0x30, 0x0a, 0x08, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x73, 0x22, 0xb4, 0x01, 0x0a, 0x07, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a,
	0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x2f, 0x0a, 0x09, 0x63,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x45, 0x78, 0x70,
	0x72, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4a, 0x0a, 0x0b,
	0x70, 0x61, 0x72, 0x73, 0x65, 0x64, 0x5f, 0x65, 0x78, 0x70, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65,
	0x78, 0x70, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x61, 0x72,
	0x73, 0x65, 0x64, 0x45, 0x78, 0x70, 0x72, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x70, 0x61,
	0x72, 0x73, 0x65, 0x64, 0x45, 0x78, 0x70, 0x72, 0x42, 0x11, 0x5a, 0x0f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_v1_iam_policy_proto_rawDescOnce sync.Once
	file_v1_iam_policy_proto_rawDescData = file_v1_iam_policy_proto_rawDesc
)

func file_v1_iam_policy_proto_rawDescGZIP() []byte {
	file_v1_iam_policy_proto_rawDescOnce.Do(func() {
		file_v1_iam_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_iam_policy_proto_rawDescData)
	})
	return file_v1_iam_policy_proto_rawDescData
}

var file_v1_iam_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v1_iam_policy_proto_goTypes = []interface{}{
	(*IamPolicy)(nil),           // 0: bytebase.v1.IamPolicy
	(*Binding)(nil),             // 1: bytebase.v1.Binding
	(*expr.Expr)(nil),           // 2: google.type.Expr
	(*v1alpha1.ParsedExpr)(nil), // 3: google.api.expr.v1alpha1.ParsedExpr
}
var file_v1_iam_policy_proto_depIdxs = []int32{
	1, // 0: bytebase.v1.IamPolicy.bindings:type_name -> bytebase.v1.Binding
	2, // 1: bytebase.v1.Binding.condition:type_name -> google.type.Expr
	3, // 2: bytebase.v1.Binding.parsed_expr:type_name -> google.api.expr.v1alpha1.ParsedExpr
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_v1_iam_policy_proto_init() }
func file_v1_iam_policy_proto_init() {
	if File_v1_iam_policy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_iam_policy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IamPolicy); i {
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
		file_v1_iam_policy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Binding); i {
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
			RawDescriptor: file_v1_iam_policy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_iam_policy_proto_goTypes,
		DependencyIndexes: file_v1_iam_policy_proto_depIdxs,
		MessageInfos:      file_v1_iam_policy_proto_msgTypes,
	}.Build()
	File_v1_iam_policy_proto = out.File
	file_v1_iam_policy_proto_rawDesc = nil
	file_v1_iam_policy_proto_goTypes = nil
	file_v1_iam_policy_proto_depIdxs = nil
}