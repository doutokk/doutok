// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.2
// source: product.proto

package product

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

type GetProductBatchReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Ids           []uint32               `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProductBatchReq) Reset() {
	*x = GetProductBatchReq{}
	mi := &file_product_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProductBatchReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductBatchReq) ProtoMessage() {}

func (x *GetProductBatchReq) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductBatchReq.ProtoReflect.Descriptor instead.
func (*GetProductBatchReq) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{0}
}

func (x *GetProductBatchReq) GetIds() []uint32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type GetProductBatchResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Item          []*Product             `protobuf:"bytes,1,rep,name=item,proto3" json:"item,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProductBatchResp) Reset() {
	*x = GetProductBatchResp{}
	mi := &file_product_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProductBatchResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductBatchResp) ProtoMessage() {}

func (x *GetProductBatchResp) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductBatchResp.ProtoReflect.Descriptor instead.
func (*GetProductBatchResp) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{1}
}

func (x *GetProductBatchResp) GetItem() []*Product {
	if x != nil {
		return x.Item
	}
	return nil
}

type EditProductReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Product       *Product               `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EditProductReq) Reset() {
	*x = EditProductReq{}
	mi := &file_product_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EditProductReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditProductReq) ProtoMessage() {}

func (x *EditProductReq) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditProductReq.ProtoReflect.Descriptor instead.
func (*EditProductReq) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{2}
}

func (x *EditProductReq) GetProduct() *Product {
	if x != nil {
		return x.Product
	}
	return nil
}

type EditProductResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EditProductResp) Reset() {
	*x = EditProductResp{}
	mi := &file_product_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EditProductResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditProductResp) ProtoMessage() {}

func (x *EditProductResp) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditProductResp.ProtoReflect.Descriptor instead.
func (*EditProductResp) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{3}
}

type ListProductsReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Page          int32                  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize      int64                  `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	CategoryName  string                 `protobuf:"bytes,3,opt,name=categoryName,proto3" json:"categoryName,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListProductsReq) Reset() {
	*x = ListProductsReq{}
	mi := &file_product_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListProductsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProductsReq) ProtoMessage() {}

func (x *ListProductsReq) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProductsReq.ProtoReflect.Descriptor instead.
func (*ListProductsReq) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{4}
}

func (x *ListProductsReq) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListProductsReq) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListProductsReq) GetCategoryName() string {
	if x != nil {
		return x.CategoryName
	}
	return ""
}

type Product struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Picture       string                 `protobuf:"bytes,4,opt,name=picture,proto3" json:"picture,omitempty"`
	Price         float32                `protobuf:"fixed32,5,opt,name=price,proto3" json:"price,omitempty"`
	Categories    []string               `protobuf:"bytes,6,rep,name=categories,proto3" json:"categories,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Product) Reset() {
	*x = Product{}
	mi := &file_product_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{5}
}

func (x *Product) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Product) GetPicture() string {
	if x != nil {
		return x.Picture
	}
	return ""
}

func (x *Product) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Product) GetCategories() []string {
	if x != nil {
		return x.Categories
	}
	return nil
}

type GetProductResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Product       *Product               `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProductResp) Reset() {
	*x = GetProductResp{}
	mi := &file_product_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProductResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductResp) ProtoMessage() {}

func (x *GetProductResp) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductResp.ProtoReflect.Descriptor instead.
func (*GetProductResp) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{6}
}

func (x *GetProductResp) GetProduct() *Product {
	if x != nil {
		return x.Product
	}
	return nil
}

type ListProductsResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Item          []*Product             `protobuf:"bytes,1,rep,name=item,proto3" json:"item,omitempty"`
	Total         int32                  `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListProductsResp) Reset() {
	*x = ListProductsResp{}
	mi := &file_product_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListProductsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProductsResp) ProtoMessage() {}

func (x *ListProductsResp) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProductsResp.ProtoReflect.Descriptor instead.
func (*ListProductsResp) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{7}
}

func (x *ListProductsResp) GetItem() []*Product {
	if x != nil {
		return x.Item
	}
	return nil
}

func (x *ListProductsResp) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetProductReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProductReq) Reset() {
	*x = GetProductReq{}
	mi := &file_product_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProductReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductReq) ProtoMessage() {}

func (x *GetProductReq) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductReq.ProtoReflect.Descriptor instead.
func (*GetProductReq) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{8}
}

func (x *GetProductReq) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type SearchProductsReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Query         string                 `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Page          int32                  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	PageSize      int64                  `protobuf:"varint,3,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SearchProductsReq) Reset() {
	*x = SearchProductsReq{}
	mi := &file_product_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchProductsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchProductsReq) ProtoMessage() {}

func (x *SearchProductsReq) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchProductsReq.ProtoReflect.Descriptor instead.
func (*SearchProductsReq) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{9}
}

func (x *SearchProductsReq) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *SearchProductsReq) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *SearchProductsReq) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type SearchProductsResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Item          []*Product             `protobuf:"bytes,1,rep,name=item,proto3" json:"item,omitempty"`
	Total         int32                  `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SearchProductsResp) Reset() {
	*x = SearchProductsResp{}
	mi := &file_product_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchProductsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchProductsResp) ProtoMessage() {}

func (x *SearchProductsResp) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchProductsResp.ProtoReflect.Descriptor instead.
func (*SearchProductsResp) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{10}
}

func (x *SearchProductsResp) GetItem() []*Product {
	if x != nil {
		return x.Item
	}
	return nil
}

func (x *SearchProductsResp) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type CreateProductReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Picture       string                 `protobuf:"bytes,3,opt,name=picture,proto3" json:"picture,omitempty"`
	Price         float32                `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
	Categories    []string               `protobuf:"bytes,5,rep,name=categories,proto3" json:"categories,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateProductReq) Reset() {
	*x = CreateProductReq{}
	mi := &file_product_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateProductReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductReq) ProtoMessage() {}

func (x *CreateProductReq) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductReq.ProtoReflect.Descriptor instead.
func (*CreateProductReq) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{11}
}

func (x *CreateProductReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateProductReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateProductReq) GetPicture() string {
	if x != nil {
		return x.Picture
	}
	return ""
}

func (x *CreateProductReq) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateProductReq) GetCategories() []string {
	if x != nil {
		return x.Categories
	}
	return nil
}

type CreateProductResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateProductResp) Reset() {
	*x = CreateProductResp{}
	mi := &file_product_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateProductResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductResp) ProtoMessage() {}

func (x *CreateProductResp) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductResp.ProtoReflect.Descriptor instead.
func (*CreateProductResp) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{12}
}

func (x *CreateProductResp) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteProductReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteProductReq) Reset() {
	*x = DeleteProductReq{}
	mi := &file_product_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteProductReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProductReq) ProtoMessage() {}

func (x *DeleteProductReq) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProductReq.ProtoReflect.Descriptor instead.
func (*DeleteProductReq) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{13}
}

func (x *DeleteProductReq) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteProductResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteProductResp) Reset() {
	*x = DeleteProductResp{}
	mi := &file_product_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteProductResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProductResp) ProtoMessage() {}

func (x *DeleteProductResp) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProductResp.ProtoReflect.Descriptor instead.
func (*DeleteProductResp) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{14}
}

func (x *DeleteProductResp) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_product_proto protoreflect.FileDescriptor

var file_product_proto_rawDesc = string([]byte{
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03,
	0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x3b,
	0x0a, 0x13, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x12, 0x24, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x22, 0x3c, 0x0a, 0x0e, 0x45,
	0x64, 0x69, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x12, 0x2a, 0x0a,
	0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x11, 0x0a, 0x0f, 0x45, 0x64, 0x69,
	0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x65, 0x0a, 0x0f,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x22, 0x0a, 0x0c, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x9f, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x69, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x22, 0x3c, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2a, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x22, 0x4e, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x24, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x22, 0x1f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x59, 0x0a, 0x11, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22,
	0x50, 0x0a, 0x12, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x24, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x22, 0x98, 0x01, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70,
	0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x22, 0x23, 0x0a, 0x11,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x22, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2d, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x32, 0x97, 0x05, 0x0a, 0x15, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x58,
	0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x18,
	0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x3a, 0x01, 0x2a, 0x22, 0x08,
	0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x56, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x17,
	0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12,
	0x0f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x7b, 0x69, 0x64, 0x3d, 0x2a, 0x7d,
	0x12, 0x5b, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x73, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x1b,
	0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x10, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0a, 0x12, 0x08, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x5a, 0x0a,
	0x0b, 0x65, 0x64, 0x69, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x17, 0x2e, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e,
	0x45, 0x64, 0x69, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x3a, 0x01, 0x2a, 0x1a, 0x0d, 0x2f, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x2f, 0x65, 0x64, 0x69, 0x74, 0x12, 0x4e, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x1b, 0x2e, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x62, 0x0a, 0x0d, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22, 0x0f, 0x2f, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x5f, 0x0a,
	0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x19,
	0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x2a, 0x0f, 0x2f,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x7b, 0x69, 0x64, 0x3d, 0x2a, 0x7d, 0x42, 0x0a,
	0x5a, 0x08, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
})

var (
	file_product_proto_rawDescOnce sync.Once
	file_product_proto_rawDescData []byte
)

func file_product_proto_rawDescGZIP() []byte {
	file_product_proto_rawDescOnce.Do(func() {
		file_product_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_product_proto_rawDesc), len(file_product_proto_rawDesc)))
	})
	return file_product_proto_rawDescData
}

var file_product_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_product_proto_goTypes = []any{
	(*GetProductBatchReq)(nil),  // 0: product.GetProductBatchReq
	(*GetProductBatchResp)(nil), // 1: product.GetProductBatchResp
	(*EditProductReq)(nil),      // 2: product.EditProductReq
	(*EditProductResp)(nil),     // 3: product.EditProductResp
	(*ListProductsReq)(nil),     // 4: product.ListProductsReq
	(*Product)(nil),             // 5: product.Product
	(*GetProductResp)(nil),      // 6: product.GetProductResp
	(*ListProductsResp)(nil),    // 7: product.ListProductsResp
	(*GetProductReq)(nil),       // 8: product.GetProductReq
	(*SearchProductsReq)(nil),   // 9: product.SearchProductsReq
	(*SearchProductsResp)(nil),  // 10: product.SearchProductsResp
	(*CreateProductReq)(nil),    // 11: product.CreateProductReq
	(*CreateProductResp)(nil),   // 12: product.CreateProductResp
	(*DeleteProductReq)(nil),    // 13: product.DeleteProductReq
	(*DeleteProductResp)(nil),   // 14: product.DeleteProductResp
}
var file_product_proto_depIdxs = []int32{
	5,  // 0: product.GetProductBatchResp.item:type_name -> product.Product
	5,  // 1: product.EditProductReq.product:type_name -> product.Product
	5,  // 2: product.GetProductResp.product:type_name -> product.Product
	5,  // 3: product.ListProductsResp.item:type_name -> product.Product
	5,  // 4: product.SearchProductsResp.item:type_name -> product.Product
	4,  // 5: product.ProductCatalogService.ListProducts:input_type -> product.ListProductsReq
	8,  // 6: product.ProductCatalogService.GetProduct:input_type -> product.GetProductReq
	9,  // 7: product.ProductCatalogService.SearchProducts:input_type -> product.SearchProductsReq
	2,  // 8: product.ProductCatalogService.editProduct:input_type -> product.EditProductReq
	0,  // 9: product.ProductCatalogService.GetProductBatch:input_type -> product.GetProductBatchReq
	11, // 10: product.ProductCatalogService.CreateProduct:input_type -> product.CreateProductReq
	13, // 11: product.ProductCatalogService.DeleteProduct:input_type -> product.DeleteProductReq
	7,  // 12: product.ProductCatalogService.ListProducts:output_type -> product.ListProductsResp
	6,  // 13: product.ProductCatalogService.GetProduct:output_type -> product.GetProductResp
	10, // 14: product.ProductCatalogService.SearchProducts:output_type -> product.SearchProductsResp
	3,  // 15: product.ProductCatalogService.editProduct:output_type -> product.EditProductResp
	1,  // 16: product.ProductCatalogService.GetProductBatch:output_type -> product.GetProductBatchResp
	12, // 17: product.ProductCatalogService.CreateProduct:output_type -> product.CreateProductResp
	14, // 18: product.ProductCatalogService.DeleteProduct:output_type -> product.DeleteProductResp
	12, // [12:19] is the sub-list for method output_type
	5,  // [5:12] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_product_proto_init() }
func file_product_proto_init() {
	if File_product_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_product_proto_rawDesc), len(file_product_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_product_proto_goTypes,
		DependencyIndexes: file_product_proto_depIdxs,
		MessageInfos:      file_product_proto_msgTypes,
	}.Build()
	File_product_proto = out.File
	file_product_proto_goTypes = nil
	file_product_proto_depIdxs = nil
}
