// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: api/api.proto

package miner_search_api

import (
	context "context"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type FindOfferRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Region    string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	CpuSingle int32  `protobuf:"varint,2,opt,name=cpuSingle,proto3" json:"cpuSingle,omitempty"`
	CpuMulti  int32  `protobuf:"varint,3,opt,name=cpuMulti,proto3" json:"cpuMulti,omitempty"`
	Ram       int32  `protobuf:"varint,4,opt,name=ram,proto3" json:"ram,omitempty"`
	DiskSpeed int32  `protobuf:"varint,5,opt,name=diskSpeed,proto3" json:"diskSpeed,omitempty"`
	DiskSize  int32  `protobuf:"varint,6,opt,name=diskSize,proto3" json:"diskSize,omitempty"`
	NetSpeed  int32  `protobuf:"varint,7,opt,name=netSpeed,proto3" json:"netSpeed,omitempty"`
}

func (x *FindOfferRequest) Reset() {
	*x = FindOfferRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindOfferRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindOfferRequest) ProtoMessage() {}

func (x *FindOfferRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindOfferRequest.ProtoReflect.Descriptor instead.
func (*FindOfferRequest) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{0}
}

func (x *FindOfferRequest) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *FindOfferRequest) GetCpuSingle() int32 {
	if x != nil {
		return x.CpuSingle
	}
	return 0
}

func (x *FindOfferRequest) GetCpuMulti() int32 {
	if x != nil {
		return x.CpuMulti
	}
	return 0
}

func (x *FindOfferRequest) GetRam() int32 {
	if x != nil {
		return x.Ram
	}
	return 0
}

func (x *FindOfferRequest) GetDiskSpeed() int32 {
	if x != nil {
		return x.DiskSpeed
	}
	return 0
}

func (x *FindOfferRequest) GetDiskSize() int32 {
	if x != nil {
		return x.DiskSize
	}
	return 0
}

func (x *FindOfferRequest) GetNetSpeed() int32 {
	if x != nil {
		return x.NetSpeed
	}
	return 0
}

type FindOfferResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offers []*Offer `protobuf:"bytes,1,rep,name=offers,proto3" json:"offers,omitempty"`
}

func (x *FindOfferResponse) Reset() {
	*x = FindOfferResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindOfferResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindOfferResponse) ProtoMessage() {}

func (x *FindOfferResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindOfferResponse.ProtoReflect.Descriptor instead.
func (*FindOfferResponse) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{1}
}

func (x *FindOfferResponse) GetOffers() []*Offer {
	if x != nil {
		return x.Offers
	}
	return nil
}

type VmType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Region    string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	CpuSingle int32  `protobuf:"varint,2,opt,name=cpuSingle,proto3" json:"cpuSingle,omitempty"`
	CpuMulti  int32  `protobuf:"varint,3,opt,name=cpuMulti,proto3" json:"cpuMulti,omitempty"`
	Ram       int32  `protobuf:"varint,4,opt,name=ram,proto3" json:"ram,omitempty"`
	DiskSpeed int32  `protobuf:"varint,5,opt,name=diskSpeed,proto3" json:"diskSpeed,omitempty"`
	DiskSize  int32  `protobuf:"varint,6,opt,name=diskSize,proto3" json:"diskSize,omitempty"`
	NetSpeed  int32  `protobuf:"varint,7,opt,name=netSpeed,proto3" json:"netSpeed,omitempty"`
}

func (x *VmType) Reset() {
	*x = VmType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VmType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VmType) ProtoMessage() {}

func (x *VmType) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VmType.ProtoReflect.Descriptor instead.
func (*VmType) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{2}
}

func (x *VmType) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *VmType) GetCpuSingle() int32 {
	if x != nil {
		return x.CpuSingle
	}
	return 0
}

func (x *VmType) GetCpuMulti() int32 {
	if x != nil {
		return x.CpuMulti
	}
	return 0
}

func (x *VmType) GetRam() int32 {
	if x != nil {
		return x.Ram
	}
	return 0
}

func (x *VmType) GetDiskSpeed() int32 {
	if x != nil {
		return x.DiskSpeed
	}
	return 0
}

func (x *VmType) GetDiskSize() int32 {
	if x != nil {
		return x.DiskSize
	}
	return 0
}

func (x *VmType) GetNetSpeed() int32 {
	if x != nil {
		return x.NetSpeed
	}
	return 0
}

type Offer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pps    int32   `protobuf:"varint,1,opt,name=pps,proto3" json:"pps,omitempty"`
	Miner  string  `protobuf:"bytes,2,opt,name=miner,proto3" json:"miner,omitempty"`
	VmType *VmType `protobuf:"bytes,3,opt,name=vmType,proto3" json:"vmType,omitempty"`
	Id     int32   `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Offer) Reset() {
	*x = Offer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Offer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Offer) ProtoMessage() {}

func (x *Offer) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Offer.ProtoReflect.Descriptor instead.
func (*Offer) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{3}
}

func (x *Offer) GetPps() int32 {
	if x != nil {
		return x.Pps
	}
	return 0
}

func (x *Offer) GetMiner() string {
	if x != nil {
		return x.Miner
	}
	return ""
}

func (x *Offer) GetVmType() *VmType {
	if x != nil {
		return x.VmType
	}
	return nil
}

func (x *Offer) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_api_api_proto protoreflect.FileDescriptor

var file_api_api_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x61, 0x70, 0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcc, 0x01, 0x0a, 0x10, 0x46, 0x69,
	0x6e, 0x64, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x70, 0x75, 0x53, 0x69, 0x6e,
	0x67, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x70, 0x75, 0x53, 0x69,
	0x6e, 0x67, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x70, 0x75, 0x4d, 0x75, 0x6c, 0x74, 0x69,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x70, 0x75, 0x4d, 0x75, 0x6c, 0x74, 0x69,
	0x12, 0x10, 0x0a, 0x03, 0x72, 0x61, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x72,
	0x61, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69, 0x73, 0x6b, 0x53, 0x70, 0x65, 0x65, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x64, 0x69, 0x73, 0x6b, 0x53, 0x70, 0x65, 0x65, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x64, 0x69, 0x73, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x6e, 0x65, 0x74, 0x53, 0x70, 0x65, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x6e, 0x65, 0x74, 0x53, 0x70, 0x65, 0x65, 0x64, 0x22, 0x3e, 0x0a, 0x11, 0x46, 0x69, 0x6e, 0x64,
	0x4f, 0x66, 0x66, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a,
	0x06, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x4f, 0x66, 0x66, 0x65, 0x72,
	0x52, 0x06, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x22, 0xc2, 0x01, 0x0a, 0x06, 0x56, 0x6d, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x63,
	0x70, 0x75, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x63, 0x70, 0x75, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x70, 0x75,
	0x4d, 0x75, 0x6c, 0x74, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x70, 0x75,
	0x4d, 0x75, 0x6c, 0x74, 0x69, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x61, 0x6d, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x72, 0x61, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69, 0x73, 0x6b, 0x53,
	0x70, 0x65, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x64, 0x69, 0x73, 0x6b,
	0x53, 0x70, 0x65, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x6b, 0x53, 0x69, 0x7a,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x69, 0x73, 0x6b, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x65, 0x74, 0x53, 0x70, 0x65, 0x65, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x6e, 0x65, 0x74, 0x53, 0x70, 0x65, 0x65, 0x64, 0x22, 0x6b, 0x0a,
	0x05, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x70, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x70, 0x70, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x69, 0x6e, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x2a,
	0x0a, 0x06, 0x76, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x56, 0x6d, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x06, 0x76, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x32, 0x78, 0x0a, 0x0d, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x67, 0x0a, 0x09, 0x46,
	0x69, 0x6e, 0x64, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12, 0x15, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2f, 0x66, 0x69, 0x6e, 0x64, 0x4f,
	0x66, 0x66, 0x65, 0x72, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x50, 0x32, 0x50, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d, 0x69, 0x6e, 0x65,
	0x72, 0x2d, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2d, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_api_proto_rawDescOnce sync.Once
	file_api_api_proto_rawDescData = file_api_api_proto_rawDesc
)

func file_api_api_proto_rawDescGZIP() []byte {
	file_api_api_proto_rawDescOnce.Do(func() {
		file_api_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_api_proto_rawDescData)
	})
	return file_api_api_proto_rawDescData
}

var file_api_api_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_api_proto_goTypes = []interface{}{
	(*FindOfferRequest)(nil),  // 0: search_api.FindOfferRequest
	(*FindOfferResponse)(nil), // 1: search_api.FindOfferResponse
	(*VmType)(nil),            // 2: search_api.VmType
	(*Offer)(nil),             // 3: search_api.Offer
}
var file_api_api_proto_depIdxs = []int32{
	3, // 0: search_api.FindOfferResponse.offers:type_name -> search_api.Offer
	2, // 1: search_api.Offer.vmType:type_name -> search_api.VmType
	0, // 2: search_api.SearchService.FindOffer:input_type -> search_api.FindOfferRequest
	1, // 3: search_api.SearchService.FindOffer:output_type -> search_api.FindOfferResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_api_proto_init() }
func file_api_api_proto_init() {
	if File_api_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindOfferRequest); i {
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
		file_api_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindOfferResponse); i {
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
		file_api_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VmType); i {
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
		file_api_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Offer); i {
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
			RawDescriptor: file_api_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_api_proto_goTypes,
		DependencyIndexes: file_api_api_proto_depIdxs,
		MessageInfos:      file_api_api_proto_msgTypes,
	}.Build()
	File_api_api_proto = out.File
	file_api_api_proto_rawDesc = nil
	file_api_api_proto_goTypes = nil
	file_api_api_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SearchServiceClient is the client API for SearchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SearchServiceClient interface {
	FindOffer(ctx context.Context, in *FindOfferRequest, opts ...grpc.CallOption) (*FindOfferResponse, error)
}

type searchServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSearchServiceClient(cc grpc.ClientConnInterface) SearchServiceClient {
	return &searchServiceClient{cc}
}

func (c *searchServiceClient) FindOffer(ctx context.Context, in *FindOfferRequest, opts ...grpc.CallOption) (*FindOfferResponse, error) {
	out := new(FindOfferResponse)
	err := c.cc.Invoke(ctx, "/search_api.SearchService/FindOffer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchServiceServer is the server API for SearchService service.
type SearchServiceServer interface {
	FindOffer(context.Context, *FindOfferRequest) (*FindOfferResponse, error)
}

// UnimplementedSearchServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSearchServiceServer struct {
}

func (*UnimplementedSearchServiceServer) FindOffer(context.Context, *FindOfferRequest) (*FindOfferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOffer not implemented")
}

func RegisterSearchServiceServer(s *grpc.Server, srv SearchServiceServer) {
	s.RegisterService(&_SearchService_serviceDesc, srv)
}

func _SearchService_FindOffer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindOfferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).FindOffer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/search_api.SearchService/FindOffer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).FindOffer(ctx, req.(*FindOfferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SearchService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "search_api.SearchService",
	HandlerType: (*SearchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindOffer",
			Handler:    _SearchService_FindOffer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/api.proto",
}
