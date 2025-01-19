// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.2
// source: merchant.proto

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

type ShowAllProductRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	MerchantId    string                 `protobuf:"bytes,1,opt,name=merchant_id,json=merchantId,proto3" json:"merchant_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ShowAllProductRequest) Reset() {
	*x = ShowAllProductRequest{}
	mi := &file_merchant_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ShowAllProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShowAllProductRequest) ProtoMessage() {}

func (x *ShowAllProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_merchant_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShowAllProductRequest.ProtoReflect.Descriptor instead.
func (*ShowAllProductRequest) Descriptor() ([]byte, []int) {
	return file_merchant_proto_rawDescGZIP(), []int{0}
}

func (x *ShowAllProductRequest) GetMerchantId() string {
	if x != nil {
		return x.MerchantId
	}
	return ""
}

type ShowAllProductResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Products      []*ProductResponse     `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ShowAllProductResponse) Reset() {
	*x = ShowAllProductResponse{}
	mi := &file_merchant_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ShowAllProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShowAllProductResponse) ProtoMessage() {}

func (x *ShowAllProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_merchant_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShowAllProductResponse.ProtoReflect.Descriptor instead.
func (*ShowAllProductResponse) Descriptor() ([]byte, []int) {
	return file_merchant_proto_rawDescGZIP(), []int{1}
}

func (x *ShowAllProductResponse) GetProducts() []*ProductResponse {
	if x != nil {
		return x.Products
	}
	return nil
}

type ProductResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	MerchantId    string                 `protobuf:"bytes,2,opt,name=merchant_id,json=merchantId,proto3" json:"merchant_id,omitempty"`
	Name          string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Price         string                 `protobuf:"bytes,4,opt,name=price,proto3" json:"price,omitempty"`
	Stock         int32                  `protobuf:"varint,5,opt,name=stock,proto3" json:"stock,omitempty"`
	Category      string                 `protobuf:"bytes,6,opt,name=category,proto3" json:"category,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProductResponse) Reset() {
	*x = ProductResponse{}
	mi := &file_merchant_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductResponse) ProtoMessage() {}

func (x *ProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_merchant_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductResponse.ProtoReflect.Descriptor instead.
func (*ProductResponse) Descriptor() ([]byte, []int) {
	return file_merchant_proto_rawDescGZIP(), []int{2}
}

func (x *ProductResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProductResponse) GetMerchantId() string {
	if x != nil {
		return x.MerchantId
	}
	return ""
}

func (x *ProductResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProductResponse) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *ProductResponse) GetStock() int32 {
	if x != nil {
		return x.Stock
	}
	return 0
}

func (x *ProductResponse) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

var File_merchant_proto protoreflect.FileDescriptor

var file_merchant_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x22, 0x38, 0x0a, 0x15, 0x53, 0x68,
	0x6f, 0x77, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61,
	0x6e, 0x74, 0x49, 0x64, 0x22, 0x4f, 0x0a, 0x16, 0x53, 0x68, 0x6f, 0x77, 0x41, 0x6c, 0x6c, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35,
	0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x22, 0x9e, 0x01, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x72,
	0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x32, 0x68, 0x0a, 0x0f, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61,
	0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x55, 0x0a, 0x0e, 0x53, 0x68, 0x6f,
	0x77, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1f, 0x2e, 0x6d, 0x65,
	0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x53, 0x68, 0x6f, 0x77, 0x41, 0x6c, 0x6c, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6d,
	0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x53, 0x68, 0x6f, 0x77, 0x41, 0x6c, 0x6c, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x08, 0x5a, 0x06, 0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_merchant_proto_rawDescOnce sync.Once
	file_merchant_proto_rawDescData = file_merchant_proto_rawDesc
)

func file_merchant_proto_rawDescGZIP() []byte {
	file_merchant_proto_rawDescOnce.Do(func() {
		file_merchant_proto_rawDescData = protoimpl.X.CompressGZIP(file_merchant_proto_rawDescData)
	})
	return file_merchant_proto_rawDescData
}

var file_merchant_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_merchant_proto_goTypes = []any{
	(*ShowAllProductRequest)(nil),  // 0: merchant.ShowAllProductRequest
	(*ShowAllProductResponse)(nil), // 1: merchant.ShowAllProductResponse
	(*ProductResponse)(nil),        // 2: merchant.ProductResponse
}
var file_merchant_proto_depIdxs = []int32{
	2, // 0: merchant.ShowAllProductResponse.products:type_name -> merchant.ProductResponse
	0, // 1: merchant.MerchantService.ShowAllProduct:input_type -> merchant.ShowAllProductRequest
	1, // 2: merchant.MerchantService.ShowAllProduct:output_type -> merchant.ShowAllProductResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_merchant_proto_init() }
func file_merchant_proto_init() {
	if File_merchant_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_merchant_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_merchant_proto_goTypes,
		DependencyIndexes: file_merchant_proto_depIdxs,
		MessageInfos:      file_merchant_proto_msgTypes,
	}.Build()
	File_merchant_proto = out.File
	file_merchant_proto_rawDesc = nil
	file_merchant_proto_goTypes = nil
	file_merchant_proto_depIdxs = nil
}
