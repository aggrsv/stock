// Code generated by protoc-gen-go.
// source: analyze.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	analyze.proto

It has these top-level messages:
	CodeNameListReq
	CodeNameListResp
	CodeName
	GrowingCompareReq
	GrowingCompareResp
	Growth
	ProfitCompareReq
	ProfitCompareResp
	Profit
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type CodeNameListReq struct {
	Sybmol string `protobuf:"bytes,1,opt,name=sybmol" json:"sybmol,omitempty"`
}

func (m *CodeNameListReq) Reset()                    { *m = CodeNameListReq{} }
func (m *CodeNameListReq) String() string            { return proto1.CompactTextString(m) }
func (*CodeNameListReq) ProtoMessage()               {}
func (*CodeNameListReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CodeNameListReq) GetSybmol() string {
	if m != nil {
		return m.Sybmol
	}
	return ""
}

type CodeNameListResp struct {
	Codenames []*CodeName `protobuf:"bytes,1,rep,name=codenames" json:"codenames,omitempty"`
}

func (m *CodeNameListResp) Reset()                    { *m = CodeNameListResp{} }
func (m *CodeNameListResp) String() string            { return proto1.CompactTextString(m) }
func (*CodeNameListResp) ProtoMessage()               {}
func (*CodeNameListResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CodeNameListResp) GetCodenames() []*CodeName {
	if m != nil {
		return m.Codenames
	}
	return nil
}

type CodeName struct {
	Code string `protobuf:"bytes,1,opt,name=code" json:"code,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *CodeName) Reset()                    { *m = CodeName{} }
func (m *CodeName) String() string            { return proto1.CompactTextString(m) }
func (*CodeName) ProtoMessage()               {}
func (*CodeName) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CodeName) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *CodeName) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GrowingCompareReq struct {
	Code string `protobuf:"bytes,1,opt,name=code" json:"code,omitempty"`
}

func (m *GrowingCompareReq) Reset()                    { *m = GrowingCompareReq{} }
func (m *GrowingCompareReq) String() string            { return proto1.CompactTextString(m) }
func (*GrowingCompareReq) ProtoMessage()               {}
func (*GrowingCompareReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GrowingCompareReq) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type GrowingCompareResp struct {
	Growths []*Growth `protobuf:"bytes,1,rep,name=growths" json:"growths,omitempty"`
}

func (m *GrowingCompareResp) Reset()                    { *m = GrowingCompareResp{} }
func (m *GrowingCompareResp) String() string            { return proto1.CompactTextString(m) }
func (*GrowingCompareResp) ProtoMessage()               {}
func (*GrowingCompareResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GrowingCompareResp) GetGrowths() []*Growth {
	if m != nil {
		return m.Growths
	}
	return nil
}

// 成长能力
type Growth struct {
	Code string  `protobuf:"bytes,1,opt,name=code" json:"code,omitempty"`
	Name string  `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Year int32   `protobuf:"varint,9,opt,name=year" json:"year,omitempty"`
	Mbrg float32 `protobuf:"fixed32,3,opt,name=mbrg" json:"mbrg,omitempty"`
	Nprg float32 `protobuf:"fixed32,4,opt,name=nprg" json:"nprg,omitempty"`
	Nav  float32 `protobuf:"fixed32,5,opt,name=nav" json:"nav,omitempty"`
	Targ float32 `protobuf:"fixed32,6,opt,name=targ" json:"targ,omitempty"`
	Epsg float32 `protobuf:"fixed32,7,opt,name=epsg" json:"epsg,omitempty"`
	Seg  float32 `protobuf:"fixed32,8,opt,name=seg" json:"seg,omitempty"`
}

func (m *Growth) Reset()                    { *m = Growth{} }
func (m *Growth) String() string            { return proto1.CompactTextString(m) }
func (*Growth) ProtoMessage()               {}
func (*Growth) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Growth) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Growth) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Growth) GetYear() int32 {
	if m != nil {
		return m.Year
	}
	return 0
}

func (m *Growth) GetMbrg() float32 {
	if m != nil {
		return m.Mbrg
	}
	return 0
}

func (m *Growth) GetNprg() float32 {
	if m != nil {
		return m.Nprg
	}
	return 0
}

func (m *Growth) GetNav() float32 {
	if m != nil {
		return m.Nav
	}
	return 0
}

func (m *Growth) GetTarg() float32 {
	if m != nil {
		return m.Targ
	}
	return 0
}

func (m *Growth) GetEpsg() float32 {
	if m != nil {
		return m.Epsg
	}
	return 0
}

func (m *Growth) GetSeg() float32 {
	if m != nil {
		return m.Seg
	}
	return 0
}

type ProfitCompareReq struct {
	Code string `protobuf:"bytes,1,opt,name=code" json:"code,omitempty"`
}

func (m *ProfitCompareReq) Reset()                    { *m = ProfitCompareReq{} }
func (m *ProfitCompareReq) String() string            { return proto1.CompactTextString(m) }
func (*ProfitCompareReq) ProtoMessage()               {}
func (*ProfitCompareReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ProfitCompareReq) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type ProfitCompareResp struct {
	Profits []*Profit `protobuf:"bytes,1,rep,name=profits" json:"profits,omitempty"`
}

func (m *ProfitCompareResp) Reset()                    { *m = ProfitCompareResp{} }
func (m *ProfitCompareResp) String() string            { return proto1.CompactTextString(m) }
func (*ProfitCompareResp) ProtoMessage()               {}
func (*ProfitCompareResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ProfitCompareResp) GetProfits() []*Profit {
	if m != nil {
		return m.Profits
	}
	return nil
}

type Profit struct {
	Code            string  `protobuf:"bytes,1,opt,name=code" json:"code,omitempty"`
	Name            string  `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Year            int32   `protobuf:"varint,3,opt,name=year" json:"year,omitempty"`
	Roe             float32 `protobuf:"fixed32,4,opt,name=roe" json:"roe,omitempty"`
	NetProfitRatio  float32 `protobuf:"fixed32,5,opt,name=net_profit_ratio,json=netProfitRatio" json:"net_profit_ratio,omitempty"`
	GrossProfitRate float32 `protobuf:"fixed32,6,opt,name=gross_profit_rate,json=grossProfitRate" json:"gross_profit_rate,omitempty"`
	NetProfits      float32 `protobuf:"fixed32,7,opt,name=net_profits,json=netProfits" json:"net_profits,omitempty"`
	Esp             float32 `protobuf:"fixed32,8,opt,name=esp" json:"esp,omitempty"`
	BusinessIncome  float32 `protobuf:"fixed32,9,opt,name=business_income,json=businessIncome" json:"business_income,omitempty"`
	Bips            float32 `protobuf:"fixed32,10,opt,name=bips" json:"bips,omitempty"`
}

func (m *Profit) Reset()                    { *m = Profit{} }
func (m *Profit) String() string            { return proto1.CompactTextString(m) }
func (*Profit) ProtoMessage()               {}
func (*Profit) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Profit) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Profit) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Profit) GetYear() int32 {
	if m != nil {
		return m.Year
	}
	return 0
}

func (m *Profit) GetRoe() float32 {
	if m != nil {
		return m.Roe
	}
	return 0
}

func (m *Profit) GetNetProfitRatio() float32 {
	if m != nil {
		return m.NetProfitRatio
	}
	return 0
}

func (m *Profit) GetGrossProfitRate() float32 {
	if m != nil {
		return m.GrossProfitRate
	}
	return 0
}

func (m *Profit) GetNetProfits() float32 {
	if m != nil {
		return m.NetProfits
	}
	return 0
}

func (m *Profit) GetEsp() float32 {
	if m != nil {
		return m.Esp
	}
	return 0
}

func (m *Profit) GetBusinessIncome() float32 {
	if m != nil {
		return m.BusinessIncome
	}
	return 0
}

func (m *Profit) GetBips() float32 {
	if m != nil {
		return m.Bips
	}
	return 0
}

func init() {
	proto1.RegisterType((*CodeNameListReq)(nil), "proto.CodeNameListReq")
	proto1.RegisterType((*CodeNameListResp)(nil), "proto.CodeNameListResp")
	proto1.RegisterType((*CodeName)(nil), "proto.CodeName")
	proto1.RegisterType((*GrowingCompareReq)(nil), "proto.GrowingCompareReq")
	proto1.RegisterType((*GrowingCompareResp)(nil), "proto.GrowingCompareResp")
	proto1.RegisterType((*Growth)(nil), "proto.Growth")
	proto1.RegisterType((*ProfitCompareReq)(nil), "proto.ProfitCompareReq")
	proto1.RegisterType((*ProfitCompareResp)(nil), "proto.ProfitCompareResp")
	proto1.RegisterType((*Profit)(nil), "proto.Profit")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for AnalyzeService service

type AnalyzeServiceClient interface {
	CodeNameList(ctx context.Context, in *CodeNameListReq, opts ...grpc.CallOption) (*CodeNameListResp, error)
	GrowingCompare(ctx context.Context, in *GrowingCompareReq, opts ...grpc.CallOption) (*GrowingCompareResp, error)
	ProfitCompare(ctx context.Context, in *ProfitCompareReq, opts ...grpc.CallOption) (*ProfitCompareResp, error)
}

type analyzeServiceClient struct {
	cc *grpc.ClientConn
}

func NewAnalyzeServiceClient(cc *grpc.ClientConn) AnalyzeServiceClient {
	return &analyzeServiceClient{cc}
}

func (c *analyzeServiceClient) CodeNameList(ctx context.Context, in *CodeNameListReq, opts ...grpc.CallOption) (*CodeNameListResp, error) {
	out := new(CodeNameListResp)
	err := grpc.Invoke(ctx, "/proto.AnalyzeService/CodeNameList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyzeServiceClient) GrowingCompare(ctx context.Context, in *GrowingCompareReq, opts ...grpc.CallOption) (*GrowingCompareResp, error) {
	out := new(GrowingCompareResp)
	err := grpc.Invoke(ctx, "/proto.AnalyzeService/GrowingCompare", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyzeServiceClient) ProfitCompare(ctx context.Context, in *ProfitCompareReq, opts ...grpc.CallOption) (*ProfitCompareResp, error) {
	out := new(ProfitCompareResp)
	err := grpc.Invoke(ctx, "/proto.AnalyzeService/ProfitCompare", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AnalyzeService service

type AnalyzeServiceServer interface {
	CodeNameList(context.Context, *CodeNameListReq) (*CodeNameListResp, error)
	GrowingCompare(context.Context, *GrowingCompareReq) (*GrowingCompareResp, error)
	ProfitCompare(context.Context, *ProfitCompareReq) (*ProfitCompareResp, error)
}

func RegisterAnalyzeServiceServer(s *grpc.Server, srv AnalyzeServiceServer) {
	s.RegisterService(&_AnalyzeService_serviceDesc, srv)
}

func _AnalyzeService_CodeNameList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CodeNameListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyzeServiceServer).CodeNameList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AnalyzeService/CodeNameList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyzeServiceServer).CodeNameList(ctx, req.(*CodeNameListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnalyzeService_GrowingCompare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GrowingCompareReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyzeServiceServer).GrowingCompare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AnalyzeService/GrowingCompare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyzeServiceServer).GrowingCompare(ctx, req.(*GrowingCompareReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnalyzeService_ProfitCompare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProfitCompareReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyzeServiceServer).ProfitCompare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AnalyzeService/ProfitCompare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyzeServiceServer).ProfitCompare(ctx, req.(*ProfitCompareReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _AnalyzeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AnalyzeService",
	HandlerType: (*AnalyzeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CodeNameList",
			Handler:    _AnalyzeService_CodeNameList_Handler,
		},
		{
			MethodName: "GrowingCompare",
			Handler:    _AnalyzeService_GrowingCompare_Handler,
		},
		{
			MethodName: "ProfitCompare",
			Handler:    _AnalyzeService_ProfitCompare_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "analyze.proto",
}

func init() { proto1.RegisterFile("analyze.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 486 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x86, 0x95, 0x76, 0xcd, 0xd6, 0x33, 0xda, 0xa6, 0xbe, 0xd8, 0xcc, 0x6e, 0xa8, 0x72, 0xc1,
	0x02, 0x12, 0xbb, 0x28, 0xb7, 0x20, 0x34, 0x26, 0x34, 0x21, 0x21, 0x84, 0xc2, 0x03, 0x54, 0x49,
	0x77, 0x08, 0x91, 0x96, 0xd8, 0xf8, 0x84, 0x4d, 0xe5, 0x59, 0x78, 0x10, 0x1e, 0x86, 0x87, 0x41,
	0xc7, 0x76, 0xb6, 0xa4, 0x2a, 0x82, 0x5d, 0xf5, 0xef, 0x77, 0x7e, 0xdb, 0x27, 0xbf, 0x7d, 0x60,
	0x92, 0xd5, 0xd9, 0xf5, 0xe6, 0x07, 0x9e, 0x69, 0xa3, 0x1a, 0x25, 0x46, 0xf6, 0x27, 0x7e, 0x06,
	0xb3, 0x0b, 0x75, 0x85, 0x1f, 0xb3, 0x0a, 0x3f, 0x94, 0xd4, 0xa4, 0xf8, 0x4d, 0x1c, 0x41, 0x48,
	0x9b, 0xbc, 0x52, 0xd7, 0x32, 0x58, 0x04, 0xc9, 0x38, 0xf5, 0xff, 0xe2, 0x73, 0x88, 0xfa, 0x56,
	0xd2, 0xe2, 0x05, 0x8c, 0xd7, 0xea, 0x0a, 0xeb, 0xac, 0x42, 0x92, 0xc1, 0x62, 0x98, 0x1c, 0x2e,
	0x67, 0xee, 0x80, 0xb3, 0xd6, 0x9b, 0xde, 0x3b, 0xe2, 0x25, 0x1c, 0xb4, 0x58, 0x08, 0xd8, 0xe3,
	0x82, 0x3f, 0xc4, 0x6a, 0x66, 0x6c, 0x94, 0x03, 0xc7, 0x58, 0xc7, 0xa7, 0x30, 0xbf, 0x34, 0xea,
	0xb6, 0xac, 0x8b, 0x0b, 0x55, 0xe9, 0xcc, 0x20, 0xf7, 0xb8, 0x63, 0x71, 0xfc, 0x1a, 0xc4, 0xb6,
	0x91, 0xb4, 0x38, 0x85, 0xfd, 0xc2, 0xa8, 0xdb, 0xe6, 0x6b, 0xdb, 0xdf, 0xc4, 0xf7, 0x77, 0x69,
	0x69, 0xda, 0x56, 0xe3, 0x5f, 0x01, 0x84, 0x8e, 0xfd, 0x6f, 0x6b, 0xcc, 0x36, 0x98, 0x19, 0x39,
	0x5e, 0x04, 0xc9, 0x28, 0xb5, 0x9a, 0x59, 0x95, 0x9b, 0x42, 0x0e, 0x17, 0x41, 0x32, 0x48, 0xad,
	0xb6, 0x6b, 0xb5, 0x29, 0xe4, 0x9e, 0x63, 0xac, 0x45, 0x04, 0xc3, 0x3a, 0xbb, 0x91, 0x23, 0x8b,
	0x58, 0xb2, 0xab, 0xc9, 0x4c, 0x21, 0x43, 0xe7, 0x62, 0xcd, 0x0c, 0x35, 0x15, 0x72, 0xdf, 0x31,
	0xd6, 0xbc, 0x92, 0xb0, 0x90, 0x07, 0x6e, 0x25, 0x61, 0x11, 0x3f, 0x85, 0xe8, 0x93, 0x51, 0x5f,
	0xca, 0xe6, 0x1f, 0x09, 0xbd, 0x82, 0xf9, 0x96, 0xcf, 0x05, 0xa4, 0x2d, 0xdc, 0x0e, 0xc8, 0x59,
	0xd3, 0xb6, 0x1a, 0xff, 0x1c, 0x40, 0xe8, 0xd8, 0x83, 0x03, 0x1a, 0x76, 0x02, 0x8a, 0x60, 0x68,
	0x14, 0xfa, 0x2c, 0x58, 0x8a, 0x04, 0xa2, 0x1a, 0x9b, 0x95, 0x3b, 0x67, 0x65, 0xb2, 0xa6, 0x54,
	0x3e, 0x97, 0x69, 0x8d, 0x8d, 0x6f, 0x83, 0xa9, 0x78, 0x0e, 0xf3, 0xc2, 0x28, 0xa2, 0x8e, 0x17,
	0x7d, 0x5e, 0x33, 0x5b, 0xb8, 0x33, 0xa3, 0x78, 0x02, 0x87, 0xf7, 0xbb, 0x92, 0x4f, 0x10, 0xee,
	0x36, 0x24, 0x6e, 0x04, 0x49, 0xb7, 0x39, 0xba, 0x28, 0x66, 0xf9, 0x77, 0x2a, 0x6b, 0x24, 0x5a,
	0x95, 0xf5, 0x5a, 0x55, 0x68, 0xaf, 0x76, 0x90, 0x4e, 0x5b, 0xfc, 0xde, 0x52, 0xfe, 0xae, 0xbc,
	0xd4, 0x24, 0xc1, 0x5d, 0x0b, 0xeb, 0xe5, 0xef, 0x00, 0xa6, 0xe7, 0x6e, 0xc4, 0x3e, 0xa3, 0xb9,
	0x29, 0xd7, 0x28, 0xde, 0xc0, 0xa3, 0xee, 0xc4, 0x88, 0xa3, 0xad, 0xd1, 0xf0, 0x13, 0x77, 0x72,
	0xbc, 0x93, 0x93, 0x16, 0xef, 0x60, 0xda, 0x7f, 0xd2, 0x42, 0x76, 0x5e, 0x6f, 0x6f, 0x24, 0x4e,
	0x1e, 0xff, 0xa5, 0x42, 0x5a, 0xbc, 0x85, 0x49, 0xef, 0xde, 0xc5, 0x71, 0xef, 0x8a, 0x3b, 0x9b,
	0xc8, 0xdd, 0x05, 0xd2, 0x79, 0x68, 0x0b, 0x2f, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x7a, 0x08,
	0x63, 0xae, 0x47, 0x04, 0x00, 0x00,
}
