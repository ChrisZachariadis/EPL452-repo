// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rate.proto

package rate

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Request message with hotel IDs and date range
type Request struct {
	HotelIds             []string `protobuf:"bytes,1,rep,name=hotelIds,proto3" json:"hotelIds,omitempty"`
	InDate               string   `protobuf:"bytes,2,opt,name=inDate,proto3" json:"inDate,omitempty"`
	OutDate              string   `protobuf:"bytes,3,opt,name=outDate,proto3" json:"outDate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_426335fde4cae2d1, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetHotelIds() []string {
	if m != nil {
		return m.HotelIds
	}
	return nil
}

func (m *Request) GetInDate() string {
	if m != nil {
		return m.InDate
	}
	return ""
}

func (m *Request) GetOutDate() string {
	if m != nil {
		return m.OutDate
	}
	return ""
}

// Result message containing a list of RatePlan
type Result struct {
	RatePlans            []*RatePlan `protobuf:"bytes,1,rep,name=ratePlans,proto3" json:"ratePlans,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_426335fde4cae2d1, []int{1}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetRatePlans() []*RatePlan {
	if m != nil {
		return m.RatePlans
	}
	return nil
}

// RatePlan message with rate information
type RatePlan struct {
	HotelId              string    `protobuf:"bytes,1,opt,name=hotelId,proto3" json:"hotelId,omitempty"`
	InDate               string    `protobuf:"bytes,2,opt,name=inDate,proto3" json:"inDate,omitempty"`
	OutDate              string    `protobuf:"bytes,3,opt,name=outDate,proto3" json:"outDate,omitempty"`
	RoomType             *RoomType `protobuf:"bytes,4,opt,name=roomType,proto3" json:"roomType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RatePlan) Reset()         { *m = RatePlan{} }
func (m *RatePlan) String() string { return proto.CompactTextString(m) }
func (*RatePlan) ProtoMessage()    {}
func (*RatePlan) Descriptor() ([]byte, []int) {
	return fileDescriptor_426335fde4cae2d1, []int{2}
}

func (m *RatePlan) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RatePlan.Unmarshal(m, b)
}
func (m *RatePlan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RatePlan.Marshal(b, m, deterministic)
}
func (m *RatePlan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RatePlan.Merge(m, src)
}
func (m *RatePlan) XXX_Size() int {
	return xxx_messageInfo_RatePlan.Size(m)
}
func (m *RatePlan) XXX_DiscardUnknown() {
	xxx_messageInfo_RatePlan.DiscardUnknown(m)
}

var xxx_messageInfo_RatePlan proto.InternalMessageInfo

func (m *RatePlan) GetHotelId() string {
	if m != nil {
		return m.HotelId
	}
	return ""
}

func (m *RatePlan) GetInDate() string {
	if m != nil {
		return m.InDate
	}
	return ""
}

func (m *RatePlan) GetOutDate() string {
	if m != nil {
		return m.OutDate
	}
	return ""
}

func (m *RatePlan) GetRoomType() *RoomType {
	if m != nil {
		return m.RoomType
	}
	return nil
}

// RoomType message with rate details
type RoomType struct {
	TotalRate            float64  `protobuf:"fixed64,1,opt,name=totalRate,proto3" json:"totalRate,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoomType) Reset()         { *m = RoomType{} }
func (m *RoomType) String() string { return proto.CompactTextString(m) }
func (*RoomType) ProtoMessage()    {}
func (*RoomType) Descriptor() ([]byte, []int) {
	return fileDescriptor_426335fde4cae2d1, []int{3}
}

func (m *RoomType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoomType.Unmarshal(m, b)
}
func (m *RoomType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoomType.Marshal(b, m, deterministic)
}
func (m *RoomType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomType.Merge(m, src)
}
func (m *RoomType) XXX_Size() int {
	return xxx_messageInfo_RoomType.Size(m)
}
func (m *RoomType) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomType.DiscardUnknown(m)
}

var xxx_messageInfo_RoomType proto.InternalMessageInfo

func (m *RoomType) GetTotalRate() float64 {
	if m != nil {
		return m.TotalRate
	}
	return 0
}

func (m *RoomType) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "rate.Request")
	proto.RegisterType((*Result)(nil), "rate.Result")
	proto.RegisterType((*RatePlan)(nil), "rate.RatePlan")
	proto.RegisterType((*RoomType)(nil), "rate.RoomType")
}

func init() {
	proto.RegisterFile("rate.proto", fileDescriptor_426335fde4cae2d1)
}

var fileDescriptor_426335fde4cae2d1 = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0x89, 0x5b, 0xba, 0xe9, 0xac, 0x7a, 0x98, 0x83, 0x84, 0xc5, 0x43, 0xe9, 0xc5, 0x45,
	0x64, 0x85, 0x15, 0x7c, 0x02, 0x41, 0xf4, 0x24, 0x41, 0xf0, 0x5c, 0xdd, 0x01, 0x0b, 0xb5, 0x53,
	0x9b, 0xe9, 0xc1, 0xab, 0x4f, 0x2e, 0x49, 0x93, 0xdd, 0x5e, 0xbd, 0xe5, 0xff, 0x66, 0x98, 0xf9,
	0x98, 0x00, 0x0c, 0xb5, 0xd0, 0xb6, 0x1f, 0x58, 0x18, 0x33, 0xff, 0xae, 0xde, 0x60, 0x69, 0xe9,
	0x7b, 0x24, 0x27, 0xb8, 0x06, 0xfd, 0xc9, 0x42, 0xed, 0xd3, 0xde, 0x19, 0x55, 0x2e, 0x36, 0x85,
	0x3d, 0x64, 0xbc, 0x80, 0xbc, 0xe9, 0x1e, 0x6a, 0x21, 0x73, 0x52, 0xaa, 0x4d, 0x61, 0x63, 0x42,
	0x03, 0x4b, 0x1e, 0x25, 0x14, 0x16, 0xa1, 0x90, 0x62, 0x75, 0x0f, 0xb9, 0x25, 0x37, 0xb6, 0x82,
	0x37, 0x50, 0xf8, 0x55, 0x2f, 0x6d, 0xdd, 0x4d, 0x83, 0x57, 0xbb, 0xf3, 0x6d, 0x10, 0xb1, 0x11,
	0xdb, 0x63, 0x43, 0xf5, 0xab, 0x40, 0x27, 0xee, 0xc7, 0x47, 0x05, 0xa3, 0xa6, 0xf1, 0x31, 0xfe,
	0x5f, 0x08, 0xaf, 0x41, 0x0f, 0xcc, 0x5f, 0xaf, 0x3f, 0x3d, 0x99, 0xac, 0x54, 0x33, 0x8b, 0x48,
	0xed, 0xa1, 0x5e, 0x3d, 0x83, 0x4e, 0x14, 0x2f, 0xa1, 0x10, 0x96, 0xba, 0xf5, 0x52, 0xc1, 0x42,
	0xd9, 0x23, 0xc0, 0x12, 0x56, 0x7b, 0x72, 0x1f, 0x43, 0xd3, 0x4b, 0xc3, 0x5d, 0x94, 0x99, 0xa3,
	0xdd, 0x2d, 0x64, 0xa1, 0xf3, 0x0a, 0xf4, 0x23, 0x89, 0x7f, 0x3a, 0x3c, 0x8b, 0x9b, 0xa7, 0xcb,
	0xaf, 0x4f, 0x53, 0xf4, 0xf7, 0x7a, 0xcf, 0xc3, 0xff, 0xdc, 0xfd, 0x05, 0x00, 0x00, 0xff, 0xff,
	0xc0, 0x18, 0xd0, 0xdf, 0xad, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RateClient is the client API for Rate service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RateClient interface {
	GetRates(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error)
}

type rateClient struct {
	cc grpc.ClientConnInterface
}

func NewRateClient(cc grpc.ClientConnInterface) RateClient {
	return &rateClient{cc}
}

func (c *rateClient) GetRates(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/rate.Rate/GetRates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RateServer is the server API for Rate service.
type RateServer interface {
	GetRates(context.Context, *Request) (*Result, error)
}

// UnimplementedRateServer can be embedded to have forward compatible implementations.
type UnimplementedRateServer struct {
}

func (*UnimplementedRateServer) GetRates(ctx context.Context, req *Request) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRates not implemented")
}

func RegisterRateServer(s *grpc.Server, srv RateServer) {
	s.RegisterService(&_Rate_serviceDesc, srv)
}

func _Rate_GetRates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RateServer).GetRates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rate.Rate/GetRates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RateServer).GetRates(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Rate_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rate.Rate",
	HandlerType: (*RateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRates",
			Handler:    _Rate_GetRates_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rate.proto",
}