// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: fizzy_ocr.proto

package fizzyocrapi

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

//文本的坐标，以四个顶点坐标表示 注意：此字段可能返回 null，表示取不到有效值
type Polygon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LeftTop     *Coord `protobuf:"bytes,1,opt,name=left_top,json=leftTop,proto3" json:"left_top,omitempty"`             // 左上顶点坐标
	RightTop    *Coord `protobuf:"bytes,2,opt,name=right_top,json=rightTop,proto3" json:"right_top,omitempty"`          // 右上顶点坐标
	RightBottom *Coord `protobuf:"bytes,3,opt,name=right_bottom,json=rightBottom,proto3" json:"right_bottom,omitempty"` // 右下顶点坐标
	LeftBottom  *Coord `protobuf:"bytes,4,opt,name=left_bottom,json=leftBottom,proto3" json:"left_bottom,omitempty"`    // 左下顶点坐标
}

func (x *Polygon) Reset() {
	*x = Polygon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fizzy_ocr_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Polygon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Polygon) ProtoMessage() {}

func (x *Polygon) ProtoReflect() protoreflect.Message {
	mi := &file_fizzy_ocr_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Polygon.ProtoReflect.Descriptor instead.
func (*Polygon) Descriptor() ([]byte, []int) {
	return file_fizzy_ocr_proto_rawDescGZIP(), []int{0}
}

func (x *Polygon) GetLeftTop() *Coord {
	if x != nil {
		return x.LeftTop
	}
	return nil
}

func (x *Polygon) GetRightTop() *Coord {
	if x != nil {
		return x.RightTop
	}
	return nil
}

func (x *Polygon) GetRightBottom() *Coord {
	if x != nil {
		return x.RightBottom
	}
	return nil
}

func (x *Polygon) GetLeftBottom() *Coord {
	if x != nil {
		return x.LeftBottom
	}
	return nil
}

//坐标点
type Coord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X float32 `protobuf:"fixed32,1,opt,name=x,proto3" json:"x,omitempty"` //横坐标
	Y float32 `protobuf:"fixed32,2,opt,name=y,proto3" json:"y,omitempty"` //纵坐标
}

func (x *Coord) Reset() {
	*x = Coord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fizzy_ocr_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Coord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Coord) ProtoMessage() {}

func (x *Coord) ProtoReflect() protoreflect.Message {
	mi := &file_fizzy_ocr_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Coord.ProtoReflect.Descriptor instead.
func (*Coord) Descriptor() ([]byte, []int) {
	return file_fizzy_ocr_proto_rawDescGZIP(), []int{1}
}

func (x *Coord) GetX() float32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Coord) GetY() float32 {
	if x != nil {
		return x.Y
	}
	return 0
}

//OCR识别请求体（支持的图片格式：PNG、JPG、JPEG、GIF 格式）
type GeneralOCRReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImageUrl    string `protobuf:"bytes,1,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`          //图片的 Url 地址。
	ImageBase64 string `protobuf:"bytes,2,opt,name=image_base64,json=imageBase64,proto3" json:"image_base64,omitempty"` //图片的 Base64 值
}

func (x *GeneralOCRReq) Reset() {
	*x = GeneralOCRReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fizzy_ocr_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeneralOCRReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeneralOCRReq) ProtoMessage() {}

func (x *GeneralOCRReq) ProtoReflect() protoreflect.Message {
	mi := &file_fizzy_ocr_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeneralOCRReq.ProtoReflect.Descriptor instead.
func (*GeneralOCRReq) Descriptor() ([]byte, []int) {
	return file_fizzy_ocr_proto_rawDescGZIP(), []int{2}
}

func (x *GeneralOCRReq) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *GeneralOCRReq) GetImageBase64() string {
	if x != nil {
		return x.ImageBase64
	}
	return ""
}

// OCR响应体
type GeneralPrintOCRRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code       int32                   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`                               //状态码
	Message    string                  `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`                          //响应信息
	ConsumTime int64                   `protobuf:"varint,3,opt,name=consum_time,json=consumTime,proto3" json:"consum_time,omitempty"` //接口耗时
	Data       *GeneralPrintOCRRspData `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`                                //响应数据
}

func (x *GeneralPrintOCRRsp) Reset() {
	*x = GeneralPrintOCRRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fizzy_ocr_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeneralPrintOCRRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeneralPrintOCRRsp) ProtoMessage() {}

func (x *GeneralPrintOCRRsp) ProtoReflect() protoreflect.Message {
	mi := &file_fizzy_ocr_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeneralPrintOCRRsp.ProtoReflect.Descriptor instead.
func (*GeneralPrintOCRRsp) Descriptor() ([]byte, []int) {
	return file_fizzy_ocr_proto_rawDescGZIP(), []int{3}
}

func (x *GeneralPrintOCRRsp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GeneralPrintOCRRsp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GeneralPrintOCRRsp) GetConsumTime() int64 {
	if x != nil {
		return x.ConsumTime
	}
	return 0
}

func (x *GeneralPrintOCRRsp) GetData() *GeneralPrintOCRRspData {
	if x != nil {
		return x.Data
	}
	return nil
}

// OCR识别数据集
type GeneralPrintOCRRspData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item []*GeneralPrintOCRRspDataItem `protobuf:"bytes,1,rep,name=item,proto3" json:"item,omitempty"` //识别数据集
}

func (x *GeneralPrintOCRRspData) Reset() {
	*x = GeneralPrintOCRRspData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fizzy_ocr_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeneralPrintOCRRspData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeneralPrintOCRRspData) ProtoMessage() {}

func (x *GeneralPrintOCRRspData) ProtoReflect() protoreflect.Message {
	mi := &file_fizzy_ocr_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeneralPrintOCRRspData.ProtoReflect.Descriptor instead.
func (*GeneralPrintOCRRspData) Descriptor() ([]byte, []int) {
	return file_fizzy_ocr_proto_rawDescGZIP(), []int{4}
}

func (x *GeneralPrintOCRRspData) GetItem() []*GeneralPrintOCRRspDataItem {
	if x != nil {
		return x.Item
	}
	return nil
}

// OCR识别内容
type GeneralPrintOCRRspDataItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Polygon *Polygon `protobuf:"bytes,1,opt,name=polygon,proto3" json:"polygon,omitempty"` //坐标点位
	Content string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"` //识别内容
	Score   float32  `protobuf:"fixed32,3,opt,name=Score,proto3" json:"Score,omitempty"`   //置信度
}

func (x *GeneralPrintOCRRspDataItem) Reset() {
	*x = GeneralPrintOCRRspDataItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fizzy_ocr_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeneralPrintOCRRspDataItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeneralPrintOCRRspDataItem) ProtoMessage() {}

func (x *GeneralPrintOCRRspDataItem) ProtoReflect() protoreflect.Message {
	mi := &file_fizzy_ocr_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeneralPrintOCRRspDataItem.ProtoReflect.Descriptor instead.
func (*GeneralPrintOCRRspDataItem) Descriptor() ([]byte, []int) {
	return file_fizzy_ocr_proto_rawDescGZIP(), []int{5}
}

func (x *GeneralPrintOCRRspDataItem) GetPolygon() *Polygon {
	if x != nil {
		return x.Polygon
	}
	return nil
}

func (x *GeneralPrintOCRRspDataItem) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *GeneralPrintOCRRspDataItem) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

var File_fizzy_ocr_proto protoreflect.FileDescriptor

var file_fizzy_ocr_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x66, 0x69, 0x7a, 0x7a, 0x79, 0x5f, 0x6f, 0x63, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0b, 0x66, 0x69, 0x7a, 0x7a, 0x79, 0x6f, 0x63, 0x72, 0x61, 0x70, 0x69, 0x22, 0xd5,
	0x01, 0x0a, 0x07, 0x50, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x08, 0x6c, 0x65,
	0x66, 0x74, 0x5f, 0x74, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x66,
	0x69, 0x7a, 0x7a, 0x79, 0x6f, 0x63, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6f, 0x72, 0x64,
	0x52, 0x07, 0x6c, 0x65, 0x66, 0x74, 0x54, 0x6f, 0x70, 0x12, 0x2f, 0x0a, 0x09, 0x72, 0x69, 0x67,
	0x68, 0x74, 0x5f, 0x74, 0x6f, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x66,
	0x69, 0x7a, 0x7a, 0x79, 0x6f, 0x63, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6f, 0x72, 0x64,
	0x52, 0x08, 0x72, 0x69, 0x67, 0x68, 0x74, 0x54, 0x6f, 0x70, 0x12, 0x35, 0x0a, 0x0c, 0x72, 0x69,
	0x67, 0x68, 0x74, 0x5f, 0x62, 0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x66, 0x69, 0x7a, 0x7a, 0x79, 0x6f, 0x63, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x43,
	0x6f, 0x6f, 0x72, 0x64, 0x52, 0x0b, 0x72, 0x69, 0x67, 0x68, 0x74, 0x42, 0x6f, 0x74, 0x74, 0x6f,
	0x6d, 0x12, 0x33, 0x0a, 0x0b, 0x6c, 0x65, 0x66, 0x74, 0x5f, 0x62, 0x6f, 0x74, 0x74, 0x6f, 0x6d,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x66, 0x69, 0x7a, 0x7a, 0x79, 0x6f, 0x63,
	0x72, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x52, 0x0a, 0x6c, 0x65, 0x66, 0x74,
	0x42, 0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x22, 0x23, 0x0a, 0x05, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x12,
	0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a,
	0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x79, 0x22, 0x4f, 0x0a, 0x0d, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x4f, 0x43, 0x52, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a, 0x09,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x36, 0x34, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x42, 0x61, 0x73, 0x65, 0x36, 0x34, 0x22, 0x9c, 0x01, 0x0a,
	0x12, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x4f, 0x43, 0x52,
	0x52, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x37, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x66, 0x69, 0x7a, 0x7a, 0x79, 0x6f, 0x63, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x4f, 0x43, 0x52, 0x52, 0x73,
	0x70, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x55, 0x0a, 0x16, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x4f, 0x43, 0x52, 0x52, 0x73,
	0x70, 0x44, 0x61, 0x74, 0x61, 0x12, 0x3b, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x66, 0x69, 0x7a, 0x7a, 0x79, 0x6f, 0x63, 0x72, 0x61, 0x70,
	0x69, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x4f, 0x43,
	0x52, 0x52, 0x73, 0x70, 0x44, 0x61, 0x74, 0x61, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x69, 0x74,
	0x65, 0x6d, 0x22, 0x7c, 0x0a, 0x1a, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x50, 0x72, 0x69,
	0x6e, 0x74, 0x4f, 0x43, 0x52, 0x52, 0x73, 0x70, 0x44, 0x61, 0x74, 0x61, 0x49, 0x74, 0x65, 0x6d,
	0x12, 0x2e, 0x0a, 0x07, 0x70, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x66, 0x69, 0x7a, 0x7a, 0x79, 0x6f, 0x63, 0x72, 0x61, 0x70, 0x69, 0x2e,
	0x50, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x52, 0x07, 0x70, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x3b, 0x66, 0x69, 0x7a, 0x7a, 0x79, 0x6f, 0x63, 0x72, 0x61,
	0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fizzy_ocr_proto_rawDescOnce sync.Once
	file_fizzy_ocr_proto_rawDescData = file_fizzy_ocr_proto_rawDesc
)

func file_fizzy_ocr_proto_rawDescGZIP() []byte {
	file_fizzy_ocr_proto_rawDescOnce.Do(func() {
		file_fizzy_ocr_proto_rawDescData = protoimpl.X.CompressGZIP(file_fizzy_ocr_proto_rawDescData)
	})
	return file_fizzy_ocr_proto_rawDescData
}

var file_fizzy_ocr_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_fizzy_ocr_proto_goTypes = []interface{}{
	(*Polygon)(nil),                    // 0: fizzyocrapi.Polygon
	(*Coord)(nil),                      // 1: fizzyocrapi.Coord
	(*GeneralOCRReq)(nil),              // 2: fizzyocrapi.GeneralOCRReq
	(*GeneralPrintOCRRsp)(nil),         // 3: fizzyocrapi.GeneralPrintOCRRsp
	(*GeneralPrintOCRRspData)(nil),     // 4: fizzyocrapi.GeneralPrintOCRRspData
	(*GeneralPrintOCRRspDataItem)(nil), // 5: fizzyocrapi.GeneralPrintOCRRspDataItem
}
var file_fizzy_ocr_proto_depIdxs = []int32{
	1, // 0: fizzyocrapi.Polygon.left_top:type_name -> fizzyocrapi.Coord
	1, // 1: fizzyocrapi.Polygon.right_top:type_name -> fizzyocrapi.Coord
	1, // 2: fizzyocrapi.Polygon.right_bottom:type_name -> fizzyocrapi.Coord
	1, // 3: fizzyocrapi.Polygon.left_bottom:type_name -> fizzyocrapi.Coord
	4, // 4: fizzyocrapi.GeneralPrintOCRRsp.data:type_name -> fizzyocrapi.GeneralPrintOCRRspData
	5, // 5: fizzyocrapi.GeneralPrintOCRRspData.item:type_name -> fizzyocrapi.GeneralPrintOCRRspDataItem
	0, // 6: fizzyocrapi.GeneralPrintOCRRspDataItem.polygon:type_name -> fizzyocrapi.Polygon
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_fizzy_ocr_proto_init() }
func file_fizzy_ocr_proto_init() {
	if File_fizzy_ocr_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fizzy_ocr_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Polygon); i {
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
		file_fizzy_ocr_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Coord); i {
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
		file_fizzy_ocr_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeneralOCRReq); i {
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
		file_fizzy_ocr_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeneralPrintOCRRsp); i {
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
		file_fizzy_ocr_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeneralPrintOCRRspData); i {
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
		file_fizzy_ocr_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeneralPrintOCRRspDataItem); i {
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
			RawDescriptor: file_fizzy_ocr_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_fizzy_ocr_proto_goTypes,
		DependencyIndexes: file_fizzy_ocr_proto_depIdxs,
		MessageInfos:      file_fizzy_ocr_proto_msgTypes,
	}.Build()
	File_fizzy_ocr_proto = out.File
	file_fizzy_ocr_proto_rawDesc = nil
	file_fizzy_ocr_proto_goTypes = nil
	file_fizzy_ocr_proto_depIdxs = nil
}
