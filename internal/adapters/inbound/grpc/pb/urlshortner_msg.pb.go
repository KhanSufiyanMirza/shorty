// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: urlshortner_msg.proto

package pb

import (
	duration "github.com/golang/protobuf/ptypes/duration"
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

type ShortnerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url         string             `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	CustomShort string             `protobuf:"bytes,2,opt,name=custom_short,json=customShort,proto3" json:"custom_short,omitempty"`
	Expiry      *duration.Duration `protobuf:"bytes,3,opt,name=expiry,proto3" json:"expiry,omitempty"`
}

func (x *ShortnerRequest) Reset() {
	*x = ShortnerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_urlshortner_msg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortnerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortnerRequest) ProtoMessage() {}

func (x *ShortnerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_urlshortner_msg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortnerRequest.ProtoReflect.Descriptor instead.
func (*ShortnerRequest) Descriptor() ([]byte, []int) {
	return file_urlshortner_msg_proto_rawDescGZIP(), []int{0}
}

func (x *ShortnerRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *ShortnerRequest) GetCustomShort() string {
	if x != nil {
		return x.CustomShort
	}
	return ""
}

func (x *ShortnerRequest) GetExpiry() *duration.Duration {
	if x != nil {
		return x.Expiry
	}
	return nil
}

type ShortnerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url             string             `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	CustomShort     string             `protobuf:"bytes,2,opt,name=custom_short,json=customShort,proto3" json:"custom_short,omitempty"`
	Expiry          *duration.Duration `protobuf:"bytes,3,opt,name=expiry,proto3" json:"expiry,omitempty"`
	XRateRemaining  int64              `protobuf:"varint,4,opt,name=x_rate_remaining,json=xRateRemaining,proto3" json:"x_rate_remaining,omitempty"`
	XRateLimitReset *duration.Duration `protobuf:"bytes,5,opt,name=x_rate_limit_reset,json=xRateLimitReset,proto3" json:"x_rate_limit_reset,omitempty"`
}

func (x *ShortnerResponse) Reset() {
	*x = ShortnerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_urlshortner_msg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortnerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortnerResponse) ProtoMessage() {}

func (x *ShortnerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_urlshortner_msg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortnerResponse.ProtoReflect.Descriptor instead.
func (*ShortnerResponse) Descriptor() ([]byte, []int) {
	return file_urlshortner_msg_proto_rawDescGZIP(), []int{1}
}

func (x *ShortnerResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *ShortnerResponse) GetCustomShort() string {
	if x != nil {
		return x.CustomShort
	}
	return ""
}

func (x *ShortnerResponse) GetExpiry() *duration.Duration {
	if x != nil {
		return x.Expiry
	}
	return nil
}

func (x *ShortnerResponse) GetXRateRemaining() int64 {
	if x != nil {
		return x.XRateRemaining
	}
	return 0
}

func (x *ShortnerResponse) GetXRateLimitReset() *duration.Duration {
	if x != nil {
		return x.XRateLimitReset
	}
	return nil
}

type UrlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortUrl string `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
}

func (x *UrlRequest) Reset() {
	*x = UrlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_urlshortner_msg_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UrlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UrlRequest) ProtoMessage() {}

func (x *UrlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_urlshortner_msg_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UrlRequest.ProtoReflect.Descriptor instead.
func (*UrlRequest) Descriptor() ([]byte, []int) {
	return file_urlshortner_msg_proto_rawDescGZIP(), []int{2}
}

func (x *UrlRequest) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

type UrlResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActualUrl string `protobuf:"bytes,1,opt,name=actual_url,json=actualUrl,proto3" json:"actual_url,omitempty"`
}

func (x *UrlResponse) Reset() {
	*x = UrlResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_urlshortner_msg_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UrlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UrlResponse) ProtoMessage() {}

func (x *UrlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_urlshortner_msg_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UrlResponse.ProtoReflect.Descriptor instead.
func (*UrlResponse) Descriptor() ([]byte, []int) {
	return file_urlshortner_msg_proto_rawDescGZIP(), []int{3}
}

func (x *UrlResponse) GetActualUrl() string {
	if x != nil {
		return x.ActualUrl
	}
	return ""
}

var File_urlshortner_msg_proto protoreflect.FileDescriptor

var file_urlshortner_msg_proto_rawDesc = []byte{
	0x0a, 0x15, 0x75, 0x72, 0x6c, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x5f, 0x6d, 0x73,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x79, 0x0a, 0x0f, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c,
	0x12, 0x21, 0x0a, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x53, 0x68,
	0x6f, 0x72, 0x74, 0x12, 0x31, 0x0a, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x22, 0xec, 0x01, 0x0a, 0x10, 0x53, 0x68, 0x6f, 0x72, 0x74,
	0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x21, 0x0a,
	0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x53, 0x68, 0x6f, 0x72, 0x74,
	0x12, 0x31, 0x0a, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x79, 0x12, 0x28, 0x0a, 0x10, 0x78, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x65,
	0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x78,
	0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x46, 0x0a,
	0x12, 0x78, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x72, 0x65,
	0x73, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x78, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x52, 0x65, 0x73, 0x65, 0x74, 0x22, 0x29, 0x0a, 0x0a, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c,
	0x22, 0x2c, 0x0a, 0x0b, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x55, 0x72, 0x6c, 0x42, 0x06,
	0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_urlshortner_msg_proto_rawDescOnce sync.Once
	file_urlshortner_msg_proto_rawDescData = file_urlshortner_msg_proto_rawDesc
)

func file_urlshortner_msg_proto_rawDescGZIP() []byte {
	file_urlshortner_msg_proto_rawDescOnce.Do(func() {
		file_urlshortner_msg_proto_rawDescData = protoimpl.X.CompressGZIP(file_urlshortner_msg_proto_rawDescData)
	})
	return file_urlshortner_msg_proto_rawDescData
}

var file_urlshortner_msg_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_urlshortner_msg_proto_goTypes = []interface{}{
	(*ShortnerRequest)(nil),   // 0: pb.ShortnerRequest
	(*ShortnerResponse)(nil),  // 1: pb.ShortnerResponse
	(*UrlRequest)(nil),        // 2: pb.UrlRequest
	(*UrlResponse)(nil),       // 3: pb.UrlResponse
	(*duration.Duration)(nil), // 4: google.protobuf.Duration
}
var file_urlshortner_msg_proto_depIdxs = []int32{
	4, // 0: pb.ShortnerRequest.expiry:type_name -> google.protobuf.Duration
	4, // 1: pb.ShortnerResponse.expiry:type_name -> google.protobuf.Duration
	4, // 2: pb.ShortnerResponse.x_rate_limit_reset:type_name -> google.protobuf.Duration
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_urlshortner_msg_proto_init() }
func file_urlshortner_msg_proto_init() {
	if File_urlshortner_msg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_urlshortner_msg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortnerRequest); i {
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
		file_urlshortner_msg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortnerResponse); i {
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
		file_urlshortner_msg_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UrlRequest); i {
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
		file_urlshortner_msg_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UrlResponse); i {
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
			RawDescriptor: file_urlshortner_msg_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_urlshortner_msg_proto_goTypes,
		DependencyIndexes: file_urlshortner_msg_proto_depIdxs,
		MessageInfos:      file_urlshortner_msg_proto_msgTypes,
	}.Build()
	File_urlshortner_msg_proto = out.File
	file_urlshortner_msg_proto_rawDesc = nil
	file_urlshortner_msg_proto_goTypes = nil
	file_urlshortner_msg_proto_depIdxs = nil
}
