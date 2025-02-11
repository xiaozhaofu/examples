// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: result/result.proto

package result

import (
	score_server "github.com/gin-gonic/examples/protobuf/score_server"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 定义Result消息
type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url      string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Title    *string  `protobuf:"bytes,2,opt,name=title,proto3,oneof" json:"title,omitempty"`
	Snippets []string `protobuf:"bytes,3,rep,name=snippets,proto3" json:"snippets,omitempty"` // 字符串数组类型
	// Types that are assignable to NoticeWay:
	//
	//	*Result_Email
	//	*Result_Phone
	NoticeWay isResult_NoticeWay     `protobuf_oneof:"notice_way"`
	Price     *wrapperspb.Int64Value `protobuf:"bytes,6,opt,name=price,proto3" json:"price,omitempty"`
	Content   string                 `protobuf:"bytes,11,opt,name=content,proto3" json:"content,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Details   *anypb.Any             `protobuf:"bytes,8,opt,name=details,proto3" json:"details,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_result_result_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_result_result_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_result_result_proto_rawDescGZIP(), []int{0}
}

func (x *Result) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Result) GetTitle() string {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return ""
}

func (x *Result) GetSnippets() []string {
	if x != nil {
		return x.Snippets
	}
	return nil
}

func (m *Result) GetNoticeWay() isResult_NoticeWay {
	if m != nil {
		return m.NoticeWay
	}
	return nil
}

func (x *Result) GetEmail() string {
	if x, ok := x.GetNoticeWay().(*Result_Email); ok {
		return x.Email
	}
	return ""
}

func (x *Result) GetPhone() string {
	if x, ok := x.GetNoticeWay().(*Result_Phone); ok {
		return x.Phone
	}
	return ""
}

func (x *Result) GetPrice() *wrapperspb.Int64Value {
	if x != nil {
		return x.Price
	}
	return nil
}

func (x *Result) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Result) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *Result) GetDetails() *anypb.Any {
	if x != nil {
		return x.Details
	}
	return nil
}

type isResult_NoticeWay interface {
	isResult_NoticeWay()
}

type Result_Email struct {
	Email string `protobuf:"bytes,4,opt,name=email,proto3,oneof"`
}

type Result_Phone struct {
	Phone string `protobuf:"bytes,5,opt,name=phone,proto3,oneof"`
}

func (*Result_Email) isResult_NoticeWay() {}

func (*Result_Phone) isResult_NoticeWay() {}

// 定义SearchResponse消息
type SearchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 引用上面定义的Result消息类型，作为results字段的类型
	Results []*Result                 `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"` // repeated关键词标记，说明results字段是一个数组
	Rtaresp *score_server.RtaResponse `protobuf:"bytes,2,opt,name=rtaresp,proto3" json:"rtaresp,omitempty"` // 自定义消息类型RtaResponse
}

func (x *SearchResponse) Reset() {
	*x = SearchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_result_result_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResponse) ProtoMessage() {}

func (x *SearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_result_result_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResponse.ProtoReflect.Descriptor instead.
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return file_result_result_proto_rawDescGZIP(), []int{1}
}

func (x *SearchResponse) GetResults() []*Result {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *SearchResponse) GetRtaresp() *score_server.RtaResponse {
	if x != nil {
		return x.Rtaresp
	}
	return nil
}

// 定义SearchRequest消息
type SearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query         string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	PageNumber    int32  `protobuf:"varint,2,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	ResultPerPage int32  `protobuf:"varint,3,opt,name=result_per_page,json=resultPerPage,proto3" json:"result_per_page,omitempty"`
}

func (x *SearchRequest) Reset() {
	*x = SearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_result_result_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRequest) ProtoMessage() {}

func (x *SearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_result_result_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRequest.ProtoReflect.Descriptor instead.
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return file_result_result_proto_rawDescGZIP(), []int{2}
}

func (x *SearchRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *SearchRequest) GetPageNumber() int32 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *SearchRequest) GetResultPerPage() int32 {
	if x != nil {
		return x.ResultPerPage
	}
	return 0
}

type Price struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MarketPrice int64 `protobuf:"varint,1,opt,name=market_price,json=marketPrice,proto3" json:"market_price,omitempty"` // 建议使用下划线的命名方式
	SalePrice   int64 `protobuf:"varint,2,opt,name=sale_price,json=salePrice,proto3" json:"sale_price,omitempty"`
}

func (x *Price) Reset() {
	*x = Price{}
	if protoimpl.UnsafeEnabled {
		mi := &file_result_result_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Price) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Price) ProtoMessage() {}

func (x *Price) ProtoReflect() protoreflect.Message {
	mi := &file_result_result_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Price.ProtoReflect.Descriptor instead.
func (*Price) Descriptor() ([]byte, []int) {
	return file_result_result_proto_rawDescGZIP(), []int{3}
}

func (x *Price) GetMarketPrice() int64 {
	if x != nil {
		return x.MarketPrice
	}
	return 0
}

func (x *Price) GetSalePrice() int64 {
	if x != nil {
		return x.SalePrice
	}
	return 0
}

type Book struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Price *Price `protobuf:"bytes,2,opt,name=price,proto3" json:"price,omitempty"`
	// Timestamp是大写T!大写T!大写T!
	Date *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"` // 注意包名前缀
}

func (x *Book) Reset() {
	*x = Book{}
	if protoimpl.UnsafeEnabled {
		mi := &file_result_result_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_result_result_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_result_result_proto_rawDescGZIP(), []int{4}
}

func (x *Book) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Book) GetPrice() *Price {
	if x != nil {
		return x.Price
	}
	return nil
}

func (x *Book) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

var File_result_result_proto protoreflect.FileDescriptor

var file_result_result_proto_rawDesc = []byte{
	0x0a, 0x13, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x1a, 0x1d, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77,
	0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61,
	0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd0, 0x02, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x72, 0x6c, 0x12, 0x19, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x01, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1a,
	0x0a, 0x08, 0x73, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x08, 0x73, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x16, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x31, 0x0a, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36,
	0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x2e, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x42, 0x0c, 0x0a, 0x0a, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x5f, 0x77, 0x61, 0x79, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x6f, 0x0a, 0x0e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x07, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x33, 0x0a, 0x07, 0x72, 0x74, 0x61, 0x72, 0x65, 0x73, 0x70,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x52, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x52, 0x07, 0x72, 0x74, 0x61, 0x72, 0x65, 0x73, 0x70, 0x22, 0x6e, 0x0a, 0x0d, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x26, 0x0a, 0x0f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x70, 0x65, 0x72,
	0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x50, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x22, 0x49, 0x0a, 0x05, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x5f, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6d, 0x61, 0x72, 0x6b, 0x65,
	0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x61, 0x6c, 0x65, 0x5f, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x61, 0x6c, 0x65,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x71, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x32, 0x50, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x09,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x15, 0x2e, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0x48, 0x0a, 0x0b, 0x42, 0x6f,
	0x6f, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x06, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x0c, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x42, 0x6f, 0x6f,
	0x6b, 0x1a, 0x0c, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x22,
	0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x3a, 0x01, 0x2a, 0x22, 0x08, 0x2f, 0x76, 0x31, 0x2f,
	0x62, 0x6f, 0x6f, 0x6b, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x67, 0x69, 0x6e, 0x2d, 0x67, 0x6f, 0x6e, 0x69, 0x63, 0x2f, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_result_result_proto_rawDescOnce sync.Once
	file_result_result_proto_rawDescData = file_result_result_proto_rawDesc
)

func file_result_result_proto_rawDescGZIP() []byte {
	file_result_result_proto_rawDescOnce.Do(func() {
		file_result_result_proto_rawDescData = protoimpl.X.CompressGZIP(file_result_result_proto_rawDescData)
	})
	return file_result_result_proto_rawDescData
}

var file_result_result_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_result_result_proto_goTypes = []interface{}{
	(*Result)(nil),                   // 0: result.Result
	(*SearchResponse)(nil),           // 1: result.SearchResponse
	(*SearchRequest)(nil),            // 2: result.SearchRequest
	(*Price)(nil),                    // 3: result.Price
	(*Book)(nil),                     // 4: result.Book
	(*wrapperspb.Int64Value)(nil),    // 5: google.protobuf.Int64Value
	(*timestamppb.Timestamp)(nil),    // 6: google.protobuf.Timestamp
	(*anypb.Any)(nil),                // 7: google.protobuf.Any
	(*score_server.RtaResponse)(nil), // 8: score_server.RtaResponse
}
var file_result_result_proto_depIdxs = []int32{
	5, // 0: result.Result.price:type_name -> google.protobuf.Int64Value
	6, // 1: result.Result.timestamp:type_name -> google.protobuf.Timestamp
	7, // 2: result.Result.details:type_name -> google.protobuf.Any
	0, // 3: result.SearchResponse.results:type_name -> result.Result
	8, // 4: result.SearchResponse.rtaresp:type_name -> score_server.RtaResponse
	3, // 5: result.Book.price:type_name -> result.Price
	6, // 6: result.Book.date:type_name -> google.protobuf.Timestamp
	2, // 7: result.GetResultService.GetResult:input_type -> result.SearchRequest
	4, // 8: result.BookService.Create:input_type -> result.Book
	1, // 9: result.GetResultService.GetResult:output_type -> result.SearchResponse
	4, // 10: result.BookService.Create:output_type -> result.Book
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_result_result_proto_init() }
func file_result_result_proto_init() {
	if File_result_result_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_result_result_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
		file_result_result_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchResponse); i {
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
		file_result_result_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchRequest); i {
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
		file_result_result_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Price); i {
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
		file_result_result_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Book); i {
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
	file_result_result_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Result_Email)(nil),
		(*Result_Phone)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_result_result_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_result_result_proto_goTypes,
		DependencyIndexes: file_result_result_proto_depIdxs,
		MessageInfos:      file_result_result_proto_msgTypes,
	}.Build()
	File_result_result_proto = out.File
	file_result_result_proto_rawDesc = nil
	file_result_result_proto_goTypes = nil
	file_result_result_proto_depIdxs = nil
}
