// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.25.1
// source: checkout.proto

package temporal

import (
	_ "github.com/cludden/protoc-gen-go-temporal/gen/temporal/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type CheckoutFlowRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Cart          *Cart                  `protobuf:"bytes,1,opt,name=cart,proto3" json:"cart,omitempty"`
	Customer      *Profile               `protobuf:"bytes,2,opt,name=customer,proto3" json:"customer,omitempty"`
	PaymentType   PaymentType            `protobuf:"varint,3,opt,name=paymentType,proto3,enum=temporal.PaymentType" json:"paymentType,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckoutFlowRequest) Reset() {
	*x = CheckoutFlowRequest{}
	mi := &file_checkout_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckoutFlowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckoutFlowRequest) ProtoMessage() {}

func (x *CheckoutFlowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_checkout_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckoutFlowRequest.ProtoReflect.Descriptor instead.
func (*CheckoutFlowRequest) Descriptor() ([]byte, []int) {
	return file_checkout_proto_rawDescGZIP(), []int{0}
}

func (x *CheckoutFlowRequest) GetCart() *Cart {
	if x != nil {
		return x.Cart
	}
	return nil
}

func (x *CheckoutFlowRequest) GetCustomer() *Profile {
	if x != nil {
		return x.Customer
	}
	return nil
}

func (x *CheckoutFlowRequest) GetPaymentType() PaymentType {
	if x != nil {
		return x.PaymentType
	}
	return PaymentType_CASH
}

type CreatePaymentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Price         int32                  `protobuf:"varint,1,opt,name=price,proto3" json:"price,omitempty"`
	CustomerId    string                 `protobuf:"bytes,2,opt,name=customerId,proto3" json:"customerId,omitempty"`
	OrderId       string                 `protobuf:"bytes,3,opt,name=orderId,proto3" json:"orderId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreatePaymentRequest) Reset() {
	*x = CreatePaymentRequest{}
	mi := &file_checkout_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreatePaymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePaymentRequest) ProtoMessage() {}

func (x *CreatePaymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_checkout_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePaymentRequest.ProtoReflect.Descriptor instead.
func (*CreatePaymentRequest) Descriptor() ([]byte, []int) {
	return file_checkout_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePaymentRequest) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreatePaymentRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *CreatePaymentRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type CreatePaymentResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status        PaymentStatus          `protobuf:"varint,2,opt,name=status,proto3,enum=temporal.PaymentStatus" json:"status,omitempty"`
	Url           string                 `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreatePaymentResponse) Reset() {
	*x = CreatePaymentResponse{}
	mi := &file_checkout_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreatePaymentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePaymentResponse) ProtoMessage() {}

func (x *CreatePaymentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_checkout_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePaymentResponse.ProtoReflect.Descriptor instead.
func (*CreatePaymentResponse) Descriptor() ([]byte, []int) {
	return file_checkout_proto_rawDescGZIP(), []int{2}
}

func (x *CreatePaymentResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreatePaymentResponse) GetStatus() PaymentStatus {
	if x != nil {
		return x.Status
	}
	return PaymentStatus_PaymentStatusNew
}

func (x *CreatePaymentResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_checkout_proto protoreflect.FileDescriptor

const file_checkout_proto_rawDesc = "" +
	"\n" +
	"\x0echeckout.proto\x12\btemporal\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1atemporal/v1/temporal.proto\x1a\x0ecustomer.proto\x1a\fcommon.proto\"\xa1\x01\n" +
	"\x13CheckoutFlowRequest\x12\"\n" +
	"\x04cart\x18\x01 \x01(\v2\x0e.temporal.CartR\x04cart\x12-\n" +
	"\bcustomer\x18\x02 \x01(\v2\x11.temporal.ProfileR\bcustomer\x127\n" +
	"\vpaymentType\x18\x03 \x01(\x0e2\x15.temporal.PaymentTypeR\vpaymentType\"f\n" +
	"\x14CreatePaymentRequest\x12\x14\n" +
	"\x05price\x18\x01 \x01(\x05R\x05price\x12\x1e\n" +
	"\n" +
	"customerId\x18\x02 \x01(\tR\n" +
	"customerId\x12\x18\n" +
	"\aorderId\x18\x03 \x01(\tR\aorderId\"j\n" +
	"\x15CreatePaymentResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12/\n" +
	"\x06status\x18\x02 \x01(\x0e2\x17.temporal.PaymentStatusR\x06status\x12\x10\n" +
	"\x03url\x18\x03 \x01(\tR\x03url2\xc1\x02\n" +
	"\bCheckout\x12g\n" +
	"\fCheckoutFlow\x12\x1d.temporal.CheckoutFlowRequest\x1a\x0f.temporal.Order\"'\x8a\xc4\x03#*\x1fcheckout/${! id.or(uuid_v4()) }0\x01\x12]\n" +
	"\x11AssortmentReserve\x12\".temporal.AssortmentReserveRequest\x1a\x16.google.protobuf.Empty\"\f\x92\xc4\x03\b\"\x02\b\x042\x02 \x03\x12^\n" +
	"\rCreatePayment\x12\x1e.temporal.CreatePaymentRequest\x1a\x1f.temporal.CreatePaymentResponse\"\f\x92\xc4\x03\b\"\x02\b\x042\x02 \x03\x1a\r\x8a\xc4\x03\t\n" +
	"\ageneralB)Z'temporalapp/internal/generated/temporalb\x06proto3"

var (
	file_checkout_proto_rawDescOnce sync.Once
	file_checkout_proto_rawDescData []byte
)

func file_checkout_proto_rawDescGZIP() []byte {
	file_checkout_proto_rawDescOnce.Do(func() {
		file_checkout_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_checkout_proto_rawDesc), len(file_checkout_proto_rawDesc)))
	})
	return file_checkout_proto_rawDescData
}

var file_checkout_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_checkout_proto_goTypes = []any{
	(*CheckoutFlowRequest)(nil),      // 0: temporal.CheckoutFlowRequest
	(*CreatePaymentRequest)(nil),     // 1: temporal.CreatePaymentRequest
	(*CreatePaymentResponse)(nil),    // 2: temporal.CreatePaymentResponse
	(*Cart)(nil),                     // 3: temporal.Cart
	(*Profile)(nil),                  // 4: temporal.Profile
	(PaymentType)(0),                 // 5: temporal.PaymentType
	(PaymentStatus)(0),               // 6: temporal.PaymentStatus
	(*AssortmentReserveRequest)(nil), // 7: temporal.AssortmentReserveRequest
	(*Order)(nil),                    // 8: temporal.Order
	(*emptypb.Empty)(nil),            // 9: google.protobuf.Empty
}
var file_checkout_proto_depIdxs = []int32{
	3, // 0: temporal.CheckoutFlowRequest.cart:type_name -> temporal.Cart
	4, // 1: temporal.CheckoutFlowRequest.customer:type_name -> temporal.Profile
	5, // 2: temporal.CheckoutFlowRequest.paymentType:type_name -> temporal.PaymentType
	6, // 3: temporal.CreatePaymentResponse.status:type_name -> temporal.PaymentStatus
	0, // 4: temporal.Checkout.CheckoutFlow:input_type -> temporal.CheckoutFlowRequest
	7, // 5: temporal.Checkout.AssortmentReserve:input_type -> temporal.AssortmentReserveRequest
	1, // 6: temporal.Checkout.CreatePayment:input_type -> temporal.CreatePaymentRequest
	8, // 7: temporal.Checkout.CheckoutFlow:output_type -> temporal.Order
	9, // 8: temporal.Checkout.AssortmentReserve:output_type -> google.protobuf.Empty
	2, // 9: temporal.Checkout.CreatePayment:output_type -> temporal.CreatePaymentResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_checkout_proto_init() }
func file_checkout_proto_init() {
	if File_checkout_proto != nil {
		return
	}
	file_customer_proto_init()
	file_common_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_checkout_proto_rawDesc), len(file_checkout_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_checkout_proto_goTypes,
		DependencyIndexes: file_checkout_proto_depIdxs,
		MessageInfos:      file_checkout_proto_msgTypes,
	}.Build()
	File_checkout_proto = out.File
	file_checkout_proto_goTypes = nil
	file_checkout_proto_depIdxs = nil
}
