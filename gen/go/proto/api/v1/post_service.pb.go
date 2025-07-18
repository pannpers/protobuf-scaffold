// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: proto/api/v1/post_service.proto

package v1

import (
	v1 "github.com/pannpers/protobuf-scaffold/gen/go/proto/entity/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// GetPostRequest is the request message for GetPost
type GetPostRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The unique identifier of the post to retrieve
	PostId        string `protobuf:"bytes,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPostRequest) Reset() {
	*x = GetPostRequest{}
	mi := &file_proto_api_v1_post_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostRequest) ProtoMessage() {}

func (x *GetPostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_v1_post_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostRequest.ProtoReflect.Descriptor instead.
func (*GetPostRequest) Descriptor() ([]byte, []int) {
	return file_proto_api_v1_post_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetPostRequest) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

// GetPostResponse is the response message for GetPost
type GetPostResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The post resource
	Post          *v1.Post `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPostResponse) Reset() {
	*x = GetPostResponse{}
	mi := &file_proto_api_v1_post_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostResponse) ProtoMessage() {}

func (x *GetPostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_v1_post_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostResponse.ProtoReflect.Descriptor instead.
func (*GetPostResponse) Descriptor() ([]byte, []int) {
	return file_proto_api_v1_post_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetPostResponse) GetPost() *v1.Post {
	if x != nil {
		return x.Post
	}
	return nil
}

// CreatePostRequest is the request message for CreatePost
type CreatePostRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The post resource to create
	Post          *v1.Post `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreatePostRequest) Reset() {
	*x = CreatePostRequest{}
	mi := &file_proto_api_v1_post_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreatePostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePostRequest) ProtoMessage() {}

func (x *CreatePostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_v1_post_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePostRequest.ProtoReflect.Descriptor instead.
func (*CreatePostRequest) Descriptor() ([]byte, []int) {
	return file_proto_api_v1_post_service_proto_rawDescGZIP(), []int{2}
}

func (x *CreatePostRequest) GetPost() *v1.Post {
	if x != nil {
		return x.Post
	}
	return nil
}

// CreatePostResponse is the response message for CreatePost
type CreatePostResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The created post resource
	Post          *v1.Post `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreatePostResponse) Reset() {
	*x = CreatePostResponse{}
	mi := &file_proto_api_v1_post_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreatePostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePostResponse) ProtoMessage() {}

func (x *CreatePostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_v1_post_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePostResponse.ProtoReflect.Descriptor instead.
func (*CreatePostResponse) Descriptor() ([]byte, []int) {
	return file_proto_api_v1_post_service_proto_rawDescGZIP(), []int{3}
}

func (x *CreatePostResponse) GetPost() *v1.Post {
	if x != nil {
		return x.Post
	}
	return nil
}

var File_proto_api_v1_post_service_proto protoreflect.FileDescriptor

const file_proto_api_v1_post_service_proto_rawDesc = "" +
	"\n" +
	"\x1fproto/api/v1/post_service.proto\x12\x03api\x1a\x1aproto/entity/v1/post.proto\x1a\x1aproto/entity/v1/user.proto\")\n" +
	"\x0eGetPostRequest\x12\x17\n" +
	"\apost_id\x18\x01 \x01(\tR\x06postId\"6\n" +
	"\x0fGetPostResponse\x12#\n" +
	"\x04post\x18\x01 \x01(\v2\x0f.entity.v1.PostR\x04post\"8\n" +
	"\x11CreatePostRequest\x12#\n" +
	"\x04post\x18\x01 \x01(\v2\x0f.entity.v1.PostR\x04post\"9\n" +
	"\x12CreatePostResponse\x12#\n" +
	"\x04post\x18\x01 \x01(\v2\x0f.entity.v1.PostR\x04post2\x82\x01\n" +
	"\vPostService\x124\n" +
	"\aGetPost\x12\x13.api.GetPostRequest\x1a\x14.api.GetPostResponse\x12=\n" +
	"\n" +
	"CreatePost\x12\x16.api.CreatePostRequest\x1a\x17.api.CreatePostResponseB;Z9github.com/pannpers/protobuf-scaffold/gen/go/proto/api/v1b\x06proto3"

var (
	file_proto_api_v1_post_service_proto_rawDescOnce sync.Once
	file_proto_api_v1_post_service_proto_rawDescData []byte
)

func file_proto_api_v1_post_service_proto_rawDescGZIP() []byte {
	file_proto_api_v1_post_service_proto_rawDescOnce.Do(func() {
		file_proto_api_v1_post_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_api_v1_post_service_proto_rawDesc), len(file_proto_api_v1_post_service_proto_rawDesc)))
	})
	return file_proto_api_v1_post_service_proto_rawDescData
}

var file_proto_api_v1_post_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_api_v1_post_service_proto_goTypes = []any{
	(*GetPostRequest)(nil),     // 0: api.GetPostRequest
	(*GetPostResponse)(nil),    // 1: api.GetPostResponse
	(*CreatePostRequest)(nil),  // 2: api.CreatePostRequest
	(*CreatePostResponse)(nil), // 3: api.CreatePostResponse
	(*v1.Post)(nil),            // 4: entity.v1.Post
}
var file_proto_api_v1_post_service_proto_depIdxs = []int32{
	4, // 0: api.GetPostResponse.post:type_name -> entity.v1.Post
	4, // 1: api.CreatePostRequest.post:type_name -> entity.v1.Post
	4, // 2: api.CreatePostResponse.post:type_name -> entity.v1.Post
	0, // 3: api.PostService.GetPost:input_type -> api.GetPostRequest
	2, // 4: api.PostService.CreatePost:input_type -> api.CreatePostRequest
	1, // 5: api.PostService.GetPost:output_type -> api.GetPostResponse
	3, // 6: api.PostService.CreatePost:output_type -> api.CreatePostResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_api_v1_post_service_proto_init() }
func file_proto_api_v1_post_service_proto_init() {
	if File_proto_api_v1_post_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_api_v1_post_service_proto_rawDesc), len(file_proto_api_v1_post_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_api_v1_post_service_proto_goTypes,
		DependencyIndexes: file_proto_api_v1_post_service_proto_depIdxs,
		MessageInfos:      file_proto_api_v1_post_service_proto_msgTypes,
	}.Build()
	File_proto_api_v1_post_service_proto = out.File
	file_proto_api_v1_post_service_proto_goTypes = nil
	file_proto_api_v1_post_service_proto_depIdxs = nil
}
