// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.0
// source: exchange.proto

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

// The request message containing the order details.
type OrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    int32  `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Type      string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`           // BUY or SELL
	OrderType string `protobuf:"bytes,3,opt,name=orderType,proto3" json:"orderType,omitempty"` // MARKET or LIMIT
	Amount    int32  `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Price     int64  `protobuf:"varint,5,opt,name=price,proto3" json:"price,omitempty"` // Ignored for MARKET orders
}

func (x *OrderRequest) Reset() {
	*x = OrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exchange_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderRequest) ProtoMessage() {}

func (x *OrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_exchange_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderRequest.ProtoReflect.Descriptor instead.
func (*OrderRequest) Descriptor() ([]byte, []int) {
	return file_exchange_proto_rawDescGZIP(), []int{0}
}

func (x *OrderRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *OrderRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *OrderRequest) GetOrderType() string {
	if x != nil {
		return x.OrderType
	}
	return ""
}

func (x *OrderRequest) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *OrderRequest) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

// The response message containing the result of the order.
type OrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Details string `protobuf:"bytes,2,opt,name=details,proto3" json:"details,omitempty"`
}

func (x *OrderResponse) Reset() {
	*x = OrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exchange_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderResponse) ProtoMessage() {}

func (x *OrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_exchange_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderResponse.ProtoReflect.Descriptor instead.
func (*OrderResponse) Descriptor() ([]byte, []int) {
	return file_exchange_proto_rawDescGZIP(), []int{1}
}

func (x *OrderResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *OrderResponse) GetDetails() string {
	if x != nil {
		return x.Details
	}
	return ""
}

var File_exchange_proto protoreflect.FileDescriptor

var file_exchange_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x22, 0x86, 0x01, 0x0a, 0x0c, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x22, 0x41, 0x0a, 0x0d, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x32, 0x4e, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x09, 0x53, 0x65, 0x6e, 0x64, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x65, 0x78,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x69, 0x63, 0x68, 0x61, 0x6c, 0x50, 0x69, 0x74, 0x72, 0x2f,
	0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_exchange_proto_rawDescOnce sync.Once
	file_exchange_proto_rawDescData = file_exchange_proto_rawDesc
)

func file_exchange_proto_rawDescGZIP() []byte {
	file_exchange_proto_rawDescOnce.Do(func() {
		file_exchange_proto_rawDescData = protoimpl.X.CompressGZIP(file_exchange_proto_rawDescData)
	})
	return file_exchange_proto_rawDescData
}

var file_exchange_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_exchange_proto_goTypes = []interface{}{
	(*OrderRequest)(nil),  // 0: exchange.OrderRequest
	(*OrderResponse)(nil), // 1: exchange.OrderResponse
}
var file_exchange_proto_depIdxs = []int32{
	0, // 0: exchange.OrderService.SendOrder:input_type -> exchange.OrderRequest
	1, // 1: exchange.OrderService.SendOrder:output_type -> exchange.OrderResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_exchange_proto_init() }
func file_exchange_proto_init() {
	if File_exchange_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_exchange_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderRequest); i {
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
		file_exchange_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderResponse); i {
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
			RawDescriptor: file_exchange_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_exchange_proto_goTypes,
		DependencyIndexes: file_exchange_proto_depIdxs,
		MessageInfos:      file_exchange_proto_msgTypes,
	}.Build()
	File_exchange_proto = out.File
	file_exchange_proto_rawDesc = nil
	file_exchange_proto_goTypes = nil
	file_exchange_proto_depIdxs = nil
}
