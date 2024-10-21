// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: recommendations.proto

package protos

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

// The request message containing strengths and weaknesses.
type RecommendationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Strengths  string `protobuf:"bytes,1,opt,name=strengths,proto3" json:"strengths,omitempty"`
	Weaknesses string `protobuf:"bytes,2,opt,name=weaknesses,proto3" json:"weaknesses,omitempty"`
}

func (x *RecommendationRequest) Reset() {
	*x = RecommendationRequest{}
	mi := &file_recommendations_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RecommendationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecommendationRequest) ProtoMessage() {}

func (x *RecommendationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_recommendations_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecommendationRequest.ProtoReflect.Descriptor instead.
func (*RecommendationRequest) Descriptor() ([]byte, []int) {
	return file_recommendations_proto_rawDescGZIP(), []int{0}
}

func (x *RecommendationRequest) GetStrengths() string {
	if x != nil {
		return x.Strengths
	}
	return ""
}

func (x *RecommendationRequest) GetWeaknesses() string {
	if x != nil {
		return x.Weaknesses
	}
	return ""
}

// The response message containing recommendations.
type RecommendationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Recommendations string `protobuf:"bytes,1,opt,name=recommendations,proto3" json:"recommendations,omitempty"`
}

func (x *RecommendationResponse) Reset() {
	*x = RecommendationResponse{}
	mi := &file_recommendations_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RecommendationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecommendationResponse) ProtoMessage() {}

func (x *RecommendationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_recommendations_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecommendationResponse.ProtoReflect.Descriptor instead.
func (*RecommendationResponse) Descriptor() ([]byte, []int) {
	return file_recommendations_proto_rawDescGZIP(), []int{1}
}

func (x *RecommendationResponse) GetRecommendations() string {
	if x != nil {
		return x.Recommendations
	}
	return ""
}

var File_recommendations_proto protoreflect.FileDescriptor

var file_recommendations_proto_rawDesc = []byte{
	0x0a, 0x15, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x55, 0x0a, 0x15, 0x52, 0x65, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x72, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x72, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x73, 0x12,
	0x1e, 0x0a, 0x0a, 0x77, 0x65, 0x61, 0x6b, 0x6e, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x65, 0x61, 0x6b, 0x6e, 0x65, 0x73, 0x73, 0x65, 0x73, 0x22,
	0x42, 0x0a, 0x16, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x72, 0x65, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x32, 0x80, 0x01, 0x0a, 0x15, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x67, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x26, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x72, 0x65,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x52, 0x65,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67,
	0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_recommendations_proto_rawDescOnce sync.Once
	file_recommendations_proto_rawDescData = file_recommendations_proto_rawDesc
)

func file_recommendations_proto_rawDescGZIP() []byte {
	file_recommendations_proto_rawDescOnce.Do(func() {
		file_recommendations_proto_rawDescData = protoimpl.X.CompressGZIP(file_recommendations_proto_rawDescData)
	})
	return file_recommendations_proto_rawDescData
}

var file_recommendations_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_recommendations_proto_goTypes = []any{
	(*RecommendationRequest)(nil),  // 0: recommendations.RecommendationRequest
	(*RecommendationResponse)(nil), // 1: recommendations.RecommendationResponse
}
var file_recommendations_proto_depIdxs = []int32{
	0, // 0: recommendations.RecommendationService.GetRecommendations:input_type -> recommendations.RecommendationRequest
	1, // 1: recommendations.RecommendationService.GetRecommendations:output_type -> recommendations.RecommendationResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_recommendations_proto_init() }
func file_recommendations_proto_init() {
	if File_recommendations_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_recommendations_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_recommendations_proto_goTypes,
		DependencyIndexes: file_recommendations_proto_depIdxs,
		MessageInfos:      file_recommendations_proto_msgTypes,
	}.Build()
	File_recommendations_proto = out.File
	file_recommendations_proto_rawDesc = nil
	file_recommendations_proto_goTypes = nil
	file_recommendations_proto_depIdxs = nil
}
