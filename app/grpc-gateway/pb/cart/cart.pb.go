// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.21.12
// source: cart.proto

package cart

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

type FrontendGetCartReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FrontendGetCartReq) Reset() {
	*x = FrontendGetCartReq{}
	mi := &file_cart_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FrontendGetCartReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FrontendGetCartReq) ProtoMessage() {}

func (x *FrontendGetCartReq) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FrontendGetCartReq.ProtoReflect.Descriptor instead.
func (*FrontendGetCartReq) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{0}
}

type FrontendItem struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProductId     uint32                 `protobuf:"varint,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	ProductName   string                 `protobuf:"bytes,2,opt,name=product_name,json=productName,proto3" json:"product_name,omitempty"`
	Price         float32                `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
	Description   string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Img           string                 `protobuf:"bytes,5,opt,name=img,proto3" json:"img,omitempty"`
	Quantity      int32                  `protobuf:"varint,6,opt,name=quantity,proto3" json:"quantity,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FrontendItem) Reset() {
	*x = FrontendItem{}
	mi := &file_cart_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FrontendItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FrontendItem) ProtoMessage() {}

func (x *FrontendItem) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FrontendItem.ProtoReflect.Descriptor instead.
func (*FrontendItem) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{1}
}

func (x *FrontendItem) GetProductId() uint32 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *FrontendItem) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *FrontendItem) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *FrontendItem) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *FrontendItem) GetImg() string {
	if x != nil {
		return x.Img
	}
	return ""
}

func (x *FrontendItem) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type FrontendGetCartResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Items         []*FrontendItem        `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FrontendGetCartResp) Reset() {
	*x = FrontendGetCartResp{}
	mi := &file_cart_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FrontendGetCartResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FrontendGetCartResp) ProtoMessage() {}

func (x *FrontendGetCartResp) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FrontendGetCartResp.ProtoReflect.Descriptor instead.
func (*FrontendGetCartResp) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{2}
}

func (x *FrontendGetCartResp) GetItems() []*FrontendItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type EditCartReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        uint32                 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Items         []*CartItem            `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EditCartReq) Reset() {
	*x = EditCartReq{}
	mi := &file_cart_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EditCartReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditCartReq) ProtoMessage() {}

func (x *EditCartReq) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditCartReq.ProtoReflect.Descriptor instead.
func (*EditCartReq) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{3}
}

func (x *EditCartReq) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *EditCartReq) GetItems() []*CartItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type EditCartResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EditCartResp) Reset() {
	*x = EditCartResp{}
	mi := &file_cart_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EditCartResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditCartResp) ProtoMessage() {}

func (x *EditCartResp) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditCartResp.ProtoReflect.Descriptor instead.
func (*EditCartResp) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{4}
}

type CartItem struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProductId     uint32                 `protobuf:"varint,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Quantity      int32                  `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CartItem) Reset() {
	*x = CartItem{}
	mi := &file_cart_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CartItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItem) ProtoMessage() {}

func (x *CartItem) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItem.ProtoReflect.Descriptor instead.
func (*CartItem) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{5}
}

func (x *CartItem) GetProductId() uint32 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *CartItem) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type AddItemReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Item          *CartItem              `protobuf:"bytes,2,opt,name=item,proto3" json:"item,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddItemReq) Reset() {
	*x = AddItemReq{}
	mi := &file_cart_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddItemReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddItemReq) ProtoMessage() {}

func (x *AddItemReq) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddItemReq.ProtoReflect.Descriptor instead.
func (*AddItemReq) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{6}
}

func (x *AddItemReq) GetItem() *CartItem {
	if x != nil {
		return x.Item
	}
	return nil
}

type AddItemResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddItemResp) Reset() {
	*x = AddItemResp{}
	mi := &file_cart_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddItemResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddItemResp) ProtoMessage() {}

func (x *AddItemResp) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddItemResp.ProtoReflect.Descriptor instead.
func (*AddItemResp) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{7}
}

type EmptyCartReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        uint32                 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EmptyCartReq) Reset() {
	*x = EmptyCartReq{}
	mi := &file_cart_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EmptyCartReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyCartReq) ProtoMessage() {}

func (x *EmptyCartReq) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyCartReq.ProtoReflect.Descriptor instead.
func (*EmptyCartReq) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{8}
}

func (x *EmptyCartReq) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetCartReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        uint32                 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCartReq) Reset() {
	*x = GetCartReq{}
	mi := &file_cart_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCartReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCartReq) ProtoMessage() {}

func (x *GetCartReq) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCartReq.ProtoReflect.Descriptor instead.
func (*GetCartReq) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{9}
}

func (x *GetCartReq) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetCartResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Cart          *Cart                  `protobuf:"bytes,1,opt,name=cart,proto3" json:"cart,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCartResp) Reset() {
	*x = GetCartResp{}
	mi := &file_cart_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCartResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCartResp) ProtoMessage() {}

func (x *GetCartResp) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCartResp.ProtoReflect.Descriptor instead.
func (*GetCartResp) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{10}
}

func (x *GetCartResp) GetCart() *Cart {
	if x != nil {
		return x.Cart
	}
	return nil
}

type Cart struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        uint32                 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Items         []*CartItem            `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Cart) Reset() {
	*x = Cart{}
	mi := &file_cart_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Cart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cart) ProtoMessage() {}

func (x *Cart) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cart.ProtoReflect.Descriptor instead.
func (*Cart) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{11}
}

func (x *Cart) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Cart) GetItems() []*CartItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type EmptyCartResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EmptyCartResp) Reset() {
	*x = EmptyCartResp{}
	mi := &file_cart_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EmptyCartResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyCartResp) ProtoMessage() {}

func (x *EmptyCartResp) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyCartResp.ProtoReflect.Descriptor instead.
func (*EmptyCartResp) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{12}
}

var File_cart_proto protoreflect.FileDescriptor

var file_cart_proto_rawDesc = string([]byte{
	0x0a, 0x0a, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x61,
	0x72, 0x74, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x14, 0x0a, 0x12, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x47, 0x65, 0x74, 0x43,
	0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x22, 0xb6, 0x01, 0x0a, 0x0c, 0x46, 0x72, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x6d, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x69, 0x6d, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22,
	0x3f, 0x0a, 0x13, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x47, 0x65, 0x74, 0x43, 0x61,
	0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x28, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x46, 0x72, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x22, 0x4c, 0x0a, 0x0b, 0x65, 0x64, 0x69, 0x74, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x43,
	0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x0e,
	0x0a, 0x0c, 0x65, 0x64, 0x69, 0x74, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x45,
	0x0a, 0x08, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x30, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x65, 0x71, 0x12, 0x22, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x22, 0x0d, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x22, 0x27, 0x0a, 0x0c, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x43,
	0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x25, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2d, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1e, 0x0a, 0x04, 0x63, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52,
	0x04, 0x63, 0x61, 0x72, 0x74, 0x22, 0x45, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x43, 0x61, 0x72,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x0f, 0x0a, 0x0d,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x32, 0xe0, 0x02,
	0x0a, 0x0b, 0x43, 0x61, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a,
	0x07, 0x41, 0x64, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x10, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e,
	0x41, 0x64, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x63, 0x61, 0x72,
	0x74, 0x2e, 0x41, 0x64, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x22, 0x16, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x10, 0x3a, 0x01, 0x2a, 0x22, 0x0b, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2f,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x30, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74,
	0x12, 0x10, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x52,
	0x65, 0x71, 0x1a, 0x11, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x09, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x43, 0x61, 0x72, 0x74, 0x12, 0x12, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12,
	0x48, 0x0a, 0x08, 0x65, 0x64, 0x69, 0x74, 0x43, 0x61, 0x72, 0x74, 0x12, 0x11, 0x2e, 0x63, 0x61,
	0x72, 0x74, 0x2e, 0x65, 0x64, 0x69, 0x74, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x12,
	0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x65, 0x64, 0x69, 0x74, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x3a, 0x01, 0x2a, 0x22, 0x0a, 0x2f,
	0x63, 0x61, 0x72, 0x74, 0x2f, 0x65, 0x64, 0x69, 0x74, 0x12, 0x55, 0x0a, 0x0f, 0x46, 0x72, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x64, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x12, 0x18, 0x2e, 0x63,
	0x61, 0x72, 0x74, 0x2e, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x47, 0x65, 0x74, 0x43,
	0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x46, 0x72,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x0d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x07, 0x12, 0x05, 0x2f, 0x63, 0x61, 0x72, 0x74,
	0x42, 0x07, 0x5a, 0x05, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
})

var (
	file_cart_proto_rawDescOnce sync.Once
	file_cart_proto_rawDescData []byte
)

func file_cart_proto_rawDescGZIP() []byte {
	file_cart_proto_rawDescOnce.Do(func() {
		file_cart_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_cart_proto_rawDesc), len(file_cart_proto_rawDesc)))
	})
	return file_cart_proto_rawDescData
}

var file_cart_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_cart_proto_goTypes = []any{
	(*FrontendGetCartReq)(nil),  // 0: cart.FrontendGetCartReq
	(*FrontendItem)(nil),        // 1: cart.FrontendItem
	(*FrontendGetCartResp)(nil), // 2: cart.FrontendGetCartResp
	(*EditCartReq)(nil),         // 3: cart.editCartReq
	(*EditCartResp)(nil),        // 4: cart.editCartResp
	(*CartItem)(nil),            // 5: cart.CartItem
	(*AddItemReq)(nil),          // 6: cart.AddItemReq
	(*AddItemResp)(nil),         // 7: cart.AddItemResp
	(*EmptyCartReq)(nil),        // 8: cart.EmptyCartReq
	(*GetCartReq)(nil),          // 9: cart.GetCartReq
	(*GetCartResp)(nil),         // 10: cart.GetCartResp
	(*Cart)(nil),                // 11: cart.Cart
	(*EmptyCartResp)(nil),       // 12: cart.EmptyCartResp
}
var file_cart_proto_depIdxs = []int32{
	1,  // 0: cart.FrontendGetCartResp.items:type_name -> cart.FrontendItem
	5,  // 1: cart.editCartReq.items:type_name -> cart.CartItem
	5,  // 2: cart.AddItemReq.item:type_name -> cart.CartItem
	11, // 3: cart.GetCartResp.cart:type_name -> cart.Cart
	5,  // 4: cart.Cart.items:type_name -> cart.CartItem
	6,  // 5: cart.CartService.AddItem:input_type -> cart.AddItemReq
	9,  // 6: cart.CartService.GetCart:input_type -> cart.GetCartReq
	8,  // 7: cart.CartService.EmptyCart:input_type -> cart.EmptyCartReq
	3,  // 8: cart.CartService.editCart:input_type -> cart.editCartReq
	0,  // 9: cart.CartService.FrontendGetCart:input_type -> cart.FrontendGetCartReq
	7,  // 10: cart.CartService.AddItem:output_type -> cart.AddItemResp
	10, // 11: cart.CartService.GetCart:output_type -> cart.GetCartResp
	12, // 12: cart.CartService.EmptyCart:output_type -> cart.EmptyCartResp
	4,  // 13: cart.CartService.editCart:output_type -> cart.editCartResp
	2,  // 14: cart.CartService.FrontendGetCart:output_type -> cart.FrontendGetCartResp
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_cart_proto_init() }
func file_cart_proto_init() {
	if File_cart_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_cart_proto_rawDesc), len(file_cart_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cart_proto_goTypes,
		DependencyIndexes: file_cart_proto_depIdxs,
		MessageInfos:      file_cart_proto_msgTypes,
	}.Build()
	File_cart_proto = out.File
	file_cart_proto_goTypes = nil
	file_cart_proto_depIdxs = nil
}
