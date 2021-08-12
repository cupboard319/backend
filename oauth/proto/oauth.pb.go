// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: proto/oauth.proto

package oauth

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

// lifted from https://github.com/micro/go-micro/blob/master/auth/service/proto/auth.proto
type AuthToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken  string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken string `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	Created      int64  `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
	Expiry       int64  `protobuf:"varint,4,opt,name=expiry,proto3" json:"expiry,omitempty"`
}

func (x *AuthToken) Reset() {
	*x = AuthToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_oauth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthToken) ProtoMessage() {}

func (x *AuthToken) ProtoReflect() protoreflect.Message {
	mi := &file_proto_oauth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthToken.ProtoReflect.Descriptor instead.
func (*AuthToken) Descriptor() ([]byte, []int) {
	return file_proto_oauth_proto_rawDescGZIP(), []int{0}
}

func (x *AuthToken) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *AuthToken) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *AuthToken) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *AuthToken) GetExpiry() int64 {
	if x != nil {
		return x.Expiry
	}
	return 0
}

type GoogleURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GoogleURLRequest) Reset() {
	*x = GoogleURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_oauth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoogleURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoogleURLRequest) ProtoMessage() {}

func (x *GoogleURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_oauth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoogleURLRequest.ProtoReflect.Descriptor instead.
func (*GoogleURLRequest) Descriptor() ([]byte, []int) {
	return file_proto_oauth_proto_rawDescGZIP(), []int{1}
}

type GoogleURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *GoogleURLResponse) Reset() {
	*x = GoogleURLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_oauth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoogleURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoogleURLResponse) ProtoMessage() {}

func (x *GoogleURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_oauth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoogleURLResponse.ProtoReflect.Descriptor instead.
func (*GoogleURLResponse) Descriptor() ([]byte, []int) {
	return file_proto_oauth_proto_rawDescGZIP(), []int{2}
}

func (x *GoogleURLResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type GoogleLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State       string `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	Code        string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	ErrorReason string `protobuf:"bytes,3,opt,name=errorReason,proto3" json:"errorReason,omitempty"`
}

func (x *GoogleLoginRequest) Reset() {
	*x = GoogleLoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_oauth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoogleLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoogleLoginRequest) ProtoMessage() {}

func (x *GoogleLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_oauth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoogleLoginRequest.ProtoReflect.Descriptor instead.
func (*GoogleLoginRequest) Descriptor() ([]byte, []int) {
	return file_proto_oauth_proto_rawDescGZIP(), []int{3}
}

func (x *GoogleLoginRequest) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *GoogleLoginRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *GoogleLoginRequest) GetErrorReason() string {
	if x != nil {
		return x.ErrorReason
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthToken  *AuthToken `protobuf:"bytes,1,opt,name=authToken,proto3" json:"authToken,omitempty"`
	CustomerID string     `protobuf:"bytes,2,opt,name=customerID,proto3" json:"customerID,omitempty"`
	Namespace  string     `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_oauth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_oauth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_proto_oauth_proto_rawDescGZIP(), []int{4}
}

func (x *LoginResponse) GetAuthToken() *AuthToken {
	if x != nil {
		return x.AuthToken
	}
	return nil
}

func (x *LoginResponse) GetCustomerID() string {
	if x != nil {
		return x.CustomerID
	}
	return ""
}

func (x *LoginResponse) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type GithubURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GithubURLRequest) Reset() {
	*x = GithubURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_oauth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GithubURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GithubURLRequest) ProtoMessage() {}

func (x *GithubURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_oauth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GithubURLRequest.ProtoReflect.Descriptor instead.
func (*GithubURLRequest) Descriptor() ([]byte, []int) {
	return file_proto_oauth_proto_rawDescGZIP(), []int{5}
}

type GithubURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *GithubURLResponse) Reset() {
	*x = GithubURLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_oauth_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GithubURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GithubURLResponse) ProtoMessage() {}

func (x *GithubURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_oauth_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GithubURLResponse.ProtoReflect.Descriptor instead.
func (*GithubURLResponse) Descriptor() ([]byte, []int) {
	return file_proto_oauth_proto_rawDescGZIP(), []int{6}
}

func (x *GithubURLResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type GithubLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State       string `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	Code        string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	ErrorReason string `protobuf:"bytes,3,opt,name=errorReason,proto3" json:"errorReason,omitempty"`
}

func (x *GithubLoginRequest) Reset() {
	*x = GithubLoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_oauth_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GithubLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GithubLoginRequest) ProtoMessage() {}

func (x *GithubLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_oauth_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GithubLoginRequest.ProtoReflect.Descriptor instead.
func (*GithubLoginRequest) Descriptor() ([]byte, []int) {
	return file_proto_oauth_proto_rawDescGZIP(), []int{7}
}

func (x *GithubLoginRequest) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *GithubLoginRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *GithubLoginRequest) GetErrorReason() string {
	if x != nil {
		return x.ErrorReason
	}
	return ""
}

var File_proto_oauth_proto protoreflect.FileDescriptor

var file_proto_oauth_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x22, 0x85, 0x01, 0x0a, 0x09, 0x41,
	0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x72,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x79, 0x22, 0x12, 0x0a, 0x10, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x55, 0x52, 0x4c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x25, 0x0a, 0x11, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x60, 0x0a,
	0x12, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22,
	0x7d, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2e, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x09, 0x61, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x44,
	0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x12,
	0x0a, 0x10, 0x47, 0x69, 0x74, 0x68, 0x75, 0x62, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x25, 0x0a, 0x11, 0x47, 0x69, 0x74, 0x68, 0x75, 0x62, 0x55, 0x52, 0x4c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x60, 0x0a, 0x12, 0x47, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x32, 0x87, 0x02, 0x0a, 0x05,
	0x4f, 0x61, 0x75, 0x74, 0x68, 0x12, 0x3e, 0x0a, 0x09, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x55,
	0x52, 0x4c, 0x12, 0x17, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6f, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0b, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x12, 0x19, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x09, 0x47, 0x69, 0x74, 0x68, 0x75, 0x62, 0x55,
	0x52, 0x4c, 0x12, 0x17, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x47, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6f, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x47, 0x69, 0x74, 0x68, 0x75, 0x62, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0b, 0x47, 0x69, 0x74, 0x68, 0x75, 0x62, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x12, 0x19, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x47, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x3b, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_oauth_proto_rawDescOnce sync.Once
	file_proto_oauth_proto_rawDescData = file_proto_oauth_proto_rawDesc
)

func file_proto_oauth_proto_rawDescGZIP() []byte {
	file_proto_oauth_proto_rawDescOnce.Do(func() {
		file_proto_oauth_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_oauth_proto_rawDescData)
	})
	return file_proto_oauth_proto_rawDescData
}

var file_proto_oauth_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_oauth_proto_goTypes = []interface{}{
	(*AuthToken)(nil),          // 0: oauth.AuthToken
	(*GoogleURLRequest)(nil),   // 1: oauth.GoogleURLRequest
	(*GoogleURLResponse)(nil),  // 2: oauth.GoogleURLResponse
	(*GoogleLoginRequest)(nil), // 3: oauth.GoogleLoginRequest
	(*LoginResponse)(nil),      // 4: oauth.LoginResponse
	(*GithubURLRequest)(nil),   // 5: oauth.GithubURLRequest
	(*GithubURLResponse)(nil),  // 6: oauth.GithubURLResponse
	(*GithubLoginRequest)(nil), // 7: oauth.GithubLoginRequest
}
var file_proto_oauth_proto_depIdxs = []int32{
	0, // 0: oauth.LoginResponse.authToken:type_name -> oauth.AuthToken
	1, // 1: oauth.Oauth.GoogleURL:input_type -> oauth.GoogleURLRequest
	3, // 2: oauth.Oauth.GoogleLogin:input_type -> oauth.GoogleLoginRequest
	5, // 3: oauth.Oauth.GithubURL:input_type -> oauth.GithubURLRequest
	7, // 4: oauth.Oauth.GithubLogin:input_type -> oauth.GithubLoginRequest
	2, // 5: oauth.Oauth.GoogleURL:output_type -> oauth.GoogleURLResponse
	4, // 6: oauth.Oauth.GoogleLogin:output_type -> oauth.LoginResponse
	6, // 7: oauth.Oauth.GithubURL:output_type -> oauth.GithubURLResponse
	4, // 8: oauth.Oauth.GithubLogin:output_type -> oauth.LoginResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_oauth_proto_init() }
func file_proto_oauth_proto_init() {
	if File_proto_oauth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_oauth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthToken); i {
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
		file_proto_oauth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoogleURLRequest); i {
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
		file_proto_oauth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoogleURLResponse); i {
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
		file_proto_oauth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoogleLoginRequest); i {
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
		file_proto_oauth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
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
		file_proto_oauth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GithubURLRequest); i {
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
		file_proto_oauth_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GithubURLResponse); i {
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
		file_proto_oauth_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GithubLoginRequest); i {
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
			RawDescriptor: file_proto_oauth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_oauth_proto_goTypes,
		DependencyIndexes: file_proto_oauth_proto_depIdxs,
		MessageInfos:      file_proto_oauth_proto_msgTypes,
	}.Build()
	File_proto_oauth_proto = out.File
	file_proto_oauth_proto_rawDesc = nil
	file_proto_oauth_proto_goTypes = nil
	file_proto_oauth_proto_depIdxs = nil
}
