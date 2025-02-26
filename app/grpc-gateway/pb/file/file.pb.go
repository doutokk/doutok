// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.2
// source: file.proto

package file

import (
	_ "github.com/doutokk/doutok/rpc_gen/kitex_gen/genproto/googleapis/api/annotations"
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

type UploadFileResp struct {
	state                protoimpl.MessageState `protogen:"open.v1"`
	Key                  string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Host                 string                 `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	Policy               string                 `protobuf:"bytes,3,opt,name=policy,proto3" json:"policy,omitempty"`
	SecurityToken        string                 `protobuf:"bytes,4,opt,name=security_token,json=securityToken,proto3" json:"security_token,omitempty"`
	Signature            string                 `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
	XOssCredential       string                 `protobuf:"bytes,6,opt,name=x_oss_credential,json=xOssCredential,proto3" json:"x_oss_credential,omitempty"`
	XOssDate             string                 `protobuf:"bytes,7,opt,name=x_oss_date,json=xOssDate,proto3" json:"x_oss_date,omitempty"`
	XOssSignatureVersion string                 `protobuf:"bytes,8,opt,name=x_oss_signature_version,json=xOssSignatureVersion,proto3" json:"x_oss_signature_version,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *UploadFileResp) Reset() {
	*x = UploadFileResp{}
	mi := &file_file_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadFileResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadFileResp) ProtoMessage() {}

func (x *UploadFileResp) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadFileResp.ProtoReflect.Descriptor instead.
func (*UploadFileResp) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{0}
}

func (x *UploadFileResp) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *UploadFileResp) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *UploadFileResp) GetPolicy() string {
	if x != nil {
		return x.Policy
	}
	return ""
}

func (x *UploadFileResp) GetSecurityToken() string {
	if x != nil {
		return x.SecurityToken
	}
	return ""
}

func (x *UploadFileResp) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *UploadFileResp) GetXOssCredential() string {
	if x != nil {
		return x.XOssCredential
	}
	return ""
}

func (x *UploadFileResp) GetXOssDate() string {
	if x != nil {
		return x.XOssDate
	}
	return ""
}

func (x *UploadFileResp) GetXOssSignatureVersion() string {
	if x != nil {
		return x.XOssSignatureVersion
	}
	return ""
}

type UploadFileReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FileName      string                 `protobuf:"bytes,2,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UploadFileReq) Reset() {
	*x = UploadFileReq{}
	mi := &file_file_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadFileReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadFileReq) ProtoMessage() {}

func (x *UploadFileReq) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadFileReq.ProtoReflect.Descriptor instead.
func (*UploadFileReq) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{1}
}

func (x *UploadFileReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UploadFileReq) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

type FrontendUploadFileReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FileName      string                 `protobuf:"bytes,2,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FrontendUploadFileReq) Reset() {
	*x = FrontendUploadFileReq{}
	mi := &file_file_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FrontendUploadFileReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FrontendUploadFileReq) ProtoMessage() {}

func (x *FrontendUploadFileReq) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FrontendUploadFileReq.ProtoReflect.Descriptor instead.
func (*FrontendUploadFileReq) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{2}
}

func (x *FrontendUploadFileReq) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

type FrontendUploadFileResp struct {
	state                protoimpl.MessageState `protogen:"open.v1"`
	Key                  string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Host                 string                 `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	Policy               string                 `protobuf:"bytes,3,opt,name=policy,proto3" json:"policy,omitempty"`
	SecurityToken        string                 `protobuf:"bytes,4,opt,name=security_token,json=securityToken,proto3" json:"security_token,omitempty"`
	Signature            string                 `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
	XOssCredential       string                 `protobuf:"bytes,6,opt,name=x_oss_credential,json=xOssCredential,proto3" json:"x_oss_credential,omitempty"`
	XOssDate             string                 `protobuf:"bytes,7,opt,name=x_oss_date,json=xOssDate,proto3" json:"x_oss_date,omitempty"`
	XOssSignatureVersion string                 `protobuf:"bytes,8,opt,name=x_oss_signature_version,json=xOssSignatureVersion,proto3" json:"x_oss_signature_version,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *FrontendUploadFileResp) Reset() {
	*x = FrontendUploadFileResp{}
	mi := &file_file_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FrontendUploadFileResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FrontendUploadFileResp) ProtoMessage() {}

func (x *FrontendUploadFileResp) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FrontendUploadFileResp.ProtoReflect.Descriptor instead.
func (*FrontendUploadFileResp) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{3}
}

func (x *FrontendUploadFileResp) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *FrontendUploadFileResp) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *FrontendUploadFileResp) GetPolicy() string {
	if x != nil {
		return x.Policy
	}
	return ""
}

func (x *FrontendUploadFileResp) GetSecurityToken() string {
	if x != nil {
		return x.SecurityToken
	}
	return ""
}

func (x *FrontendUploadFileResp) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *FrontendUploadFileResp) GetXOssCredential() string {
	if x != nil {
		return x.XOssCredential
	}
	return ""
}

func (x *FrontendUploadFileResp) GetXOssDate() string {
	if x != nil {
		return x.XOssDate
	}
	return ""
}

func (x *FrontendUploadFileResp) GetXOssSignatureVersion() string {
	if x != nil {
		return x.XOssSignatureVersion
	}
	return ""
}

var File_file_proto protoreflect.FileDescriptor

var file_file_proto_rawDesc = string([]byte{
	0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x66, 0x69,
	0x6c, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x92, 0x02, 0x0a, 0x0e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x65, 0x63, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x78, 0x5f, 0x6f, 0x73, 0x73, 0x5f,
	0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x78, 0x4f, 0x73, 0x73, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x12, 0x1c, 0x0a, 0x0a, 0x78, 0x5f, 0x6f, 0x73, 0x73, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x78, 0x4f, 0x73, 0x73, 0x44, 0x61, 0x74, 0x65, 0x12, 0x35,
	0x0a, 0x17, 0x78, 0x5f, 0x6f, 0x73, 0x73, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x14, 0x78, 0x4f, 0x73, 0x73, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x45, 0x0a, 0x0d, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x34, 0x0a, 0x15,
	0x46, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x22, 0x9a, 0x02, 0x0a, 0x16, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68,
	0x6f, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x12, 0x28, 0x0a, 0x10, 0x78, 0x5f, 0x6f, 0x73, 0x73, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x78, 0x4f, 0x73, 0x73,
	0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x1c, 0x0a, 0x0a, 0x78, 0x5f,
	0x6f, 0x73, 0x73, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x78, 0x4f, 0x73, 0x73, 0x44, 0x61, 0x74, 0x65, 0x12, 0x35, 0x0a, 0x17, 0x78, 0x5f, 0x6f, 0x73,
	0x73, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x78, 0x4f, 0x73, 0x73, 0x53,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x32,
	0xb2, 0x01, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x39, 0x0a, 0x0a, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x13, 0x2e,
	0x66, 0x69, 0x6c, 0x65, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x1a, 0x14, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x68, 0x0a, 0x12, 0x46, 0x72,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65,
	0x12, 0x1b, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e,
	0x66, 0x69, 0x6c, 0x65, 0x2e, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x17, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a, 0x22, 0x0c, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x75, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x42, 0x07, 0x5a, 0x05, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_file_proto_rawDescOnce sync.Once
	file_file_proto_rawDescData []byte
)

func file_file_proto_rawDescGZIP() []byte {
	file_file_proto_rawDescOnce.Do(func() {
		file_file_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_file_proto_rawDesc), len(file_file_proto_rawDesc)))
	})
	return file_file_proto_rawDescData
}

var file_file_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_file_proto_goTypes = []any{
	(*UploadFileResp)(nil),         // 0: file.UploadFileResp
	(*UploadFileReq)(nil),          // 1: file.UploadFileReq
	(*FrontendUploadFileReq)(nil),  // 2: file.FrontendUploadFileReq
	(*FrontendUploadFileResp)(nil), // 3: file.FrontendUploadFileResp
}
var file_file_proto_depIdxs = []int32{
	1, // 0: file.FileService.UploadFile:input_type -> file.UploadFileReq
	2, // 1: file.FileService.FrontendUploadFile:input_type -> file.FrontendUploadFileReq
	0, // 2: file.FileService.UploadFile:output_type -> file.UploadFileResp
	3, // 3: file.FileService.FrontendUploadFile:output_type -> file.FrontendUploadFileResp
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_file_proto_init() }
func file_file_proto_init() {
	if File_file_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_file_proto_rawDesc), len(file_file_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_file_proto_goTypes,
		DependencyIndexes: file_file_proto_depIdxs,
		MessageInfos:      file_file_proto_msgTypes,
	}.Build()
	File_file_proto = out.File
	file_file_proto_goTypes = nil
	file_file_proto_depIdxs = nil
}
