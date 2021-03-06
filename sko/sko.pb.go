// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sko.proto

package sko

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ObjectMeta struct {
	Key                  []byte   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Version              uint64   `protobuf:"varint,2,opt,name=version" json:"version,omitempty"`
	IncrId               uint64   `protobuf:"varint,3,opt,name=incr_id,json=incrId" json:"incr_id,omitempty"`
	Created              uint64   `protobuf:"varint,4,opt,name=created" json:"created,omitempty"`
	Updated              uint64   `protobuf:"varint,5,opt,name=updated" json:"updated,omitempty"`
	Attrs                uint64   `protobuf:"varint,6,opt,name=attrs" json:"attrs,omitempty"`
	Expired              uint64   `protobuf:"varint,11,opt,name=expired" json:"expired,omitempty"`
	DataAttrs            uint64   `protobuf:"varint,12,opt,name=data_attrs,json=dataAttrs" json:"data_attrs,omitempty"`
	DataCheck            uint64   `protobuf:"varint,13,opt,name=data_check,json=dataCheck" json:"data_check,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ObjectMeta) Reset()         { *m = ObjectMeta{} }
func (m *ObjectMeta) String() string { return proto.CompactTextString(m) }
func (*ObjectMeta) ProtoMessage()    {}
func (*ObjectMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_sko_6d21db2e82c33d93, []int{0}
}
func (m *ObjectMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectMeta.Unmarshal(m, b)
}
func (m *ObjectMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectMeta.Marshal(b, m, deterministic)
}
func (dst *ObjectMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectMeta.Merge(dst, src)
}
func (m *ObjectMeta) XXX_Size() int {
	return xxx_messageInfo_ObjectMeta.Size(m)
}
func (m *ObjectMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectMeta.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectMeta proto.InternalMessageInfo

func (m *ObjectMeta) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *ObjectMeta) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *ObjectMeta) GetIncrId() uint64 {
	if m != nil {
		return m.IncrId
	}
	return 0
}

func (m *ObjectMeta) GetCreated() uint64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *ObjectMeta) GetUpdated() uint64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *ObjectMeta) GetAttrs() uint64 {
	if m != nil {
		return m.Attrs
	}
	return 0
}

func (m *ObjectMeta) GetExpired() uint64 {
	if m != nil {
		return m.Expired
	}
	return 0
}

func (m *ObjectMeta) GetDataAttrs() uint64 {
	if m != nil {
		return m.DataAttrs
	}
	return 0
}

func (m *ObjectMeta) GetDataCheck() uint64 {
	if m != nil {
		return m.DataCheck
	}
	return 0
}

type ObjectData struct {
	Attrs                uint64   `protobuf:"varint,8,opt,name=attrs" json:"attrs,omitempty"`
	Check                uint64   `protobuf:"varint,9,opt,name=check" json:"check,omitempty"`
	Value                []byte   `protobuf:"bytes,10,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ObjectData) Reset()         { *m = ObjectData{} }
func (m *ObjectData) String() string { return proto.CompactTextString(m) }
func (*ObjectData) ProtoMessage()    {}
func (*ObjectData) Descriptor() ([]byte, []int) {
	return fileDescriptor_sko_6d21db2e82c33d93, []int{1}
}
func (m *ObjectData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectData.Unmarshal(m, b)
}
func (m *ObjectData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectData.Marshal(b, m, deterministic)
}
func (dst *ObjectData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectData.Merge(dst, src)
}
func (m *ObjectData) XXX_Size() int {
	return xxx_messageInfo_ObjectData.Size(m)
}
func (m *ObjectData) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectData.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectData proto.InternalMessageInfo

func (m *ObjectData) GetAttrs() uint64 {
	if m != nil {
		return m.Attrs
	}
	return 0
}

func (m *ObjectData) GetCheck() uint64 {
	if m != nil {
		return m.Check
	}
	return 0
}

func (m *ObjectData) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type ObjectItem struct {
	Meta                 *ObjectMeta `protobuf:"bytes,2,opt,name=meta" json:"meta,omitempty"`
	Data                 *ObjectData `protobuf:"bytes,3,opt,name=data" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ObjectItem) Reset()         { *m = ObjectItem{} }
func (m *ObjectItem) String() string { return proto.CompactTextString(m) }
func (*ObjectItem) ProtoMessage()    {}
func (*ObjectItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_sko_6d21db2e82c33d93, []int{2}
}
func (m *ObjectItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectItem.Unmarshal(m, b)
}
func (m *ObjectItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectItem.Marshal(b, m, deterministic)
}
func (dst *ObjectItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectItem.Merge(dst, src)
}
func (m *ObjectItem) XXX_Size() int {
	return xxx_messageInfo_ObjectItem.Size(m)
}
func (m *ObjectItem) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectItem.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectItem proto.InternalMessageInfo

func (m *ObjectItem) GetMeta() *ObjectMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *ObjectItem) GetData() *ObjectData {
	if m != nil {
		return m.Data
	}
	return nil
}

type ObjectReader struct {
	Mode                 uint64   `protobuf:"varint,1,opt,name=mode" json:"mode,omitempty"`
	Keys                 [][]byte `protobuf:"bytes,2,rep,name=keys,proto3" json:"keys,omitempty"`
	LimitNum             int64    `protobuf:"varint,8,opt,name=limit_num,json=limitNum" json:"limit_num,omitempty"`
	LimitSize            int64    `protobuf:"varint,9,opt,name=limit_size,json=limitSize" json:"limit_size,omitempty"`
	KeyOffset            []byte   `protobuf:"bytes,12,opt,name=key_offset,json=keyOffset,proto3" json:"key_offset,omitempty"`
	KeyCutset            []byte   `protobuf:"bytes,13,opt,name=key_cutset,json=keyCutset,proto3" json:"key_cutset,omitempty"`
	LogOffset            uint64   `protobuf:"varint,14,opt,name=log_offset,json=logOffset" json:"log_offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ObjectReader) Reset()         { *m = ObjectReader{} }
func (m *ObjectReader) String() string { return proto.CompactTextString(m) }
func (*ObjectReader) ProtoMessage()    {}
func (*ObjectReader) Descriptor() ([]byte, []int) {
	return fileDescriptor_sko_6d21db2e82c33d93, []int{3}
}
func (m *ObjectReader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectReader.Unmarshal(m, b)
}
func (m *ObjectReader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectReader.Marshal(b, m, deterministic)
}
func (dst *ObjectReader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectReader.Merge(dst, src)
}
func (m *ObjectReader) XXX_Size() int {
	return xxx_messageInfo_ObjectReader.Size(m)
}
func (m *ObjectReader) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectReader.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectReader proto.InternalMessageInfo

func (m *ObjectReader) GetMode() uint64 {
	if m != nil {
		return m.Mode
	}
	return 0
}

func (m *ObjectReader) GetKeys() [][]byte {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *ObjectReader) GetLimitNum() int64 {
	if m != nil {
		return m.LimitNum
	}
	return 0
}

func (m *ObjectReader) GetLimitSize() int64 {
	if m != nil {
		return m.LimitSize
	}
	return 0
}

func (m *ObjectReader) GetKeyOffset() []byte {
	if m != nil {
		return m.KeyOffset
	}
	return nil
}

func (m *ObjectReader) GetKeyCutset() []byte {
	if m != nil {
		return m.KeyCutset
	}
	return nil
}

func (m *ObjectReader) GetLogOffset() uint64 {
	if m != nil {
		return m.LogOffset
	}
	return 0
}

type ObjectWriter struct {
	Mode                 uint64      `protobuf:"varint,1,opt,name=mode" json:"mode,omitempty"`
	Meta                 *ObjectMeta `protobuf:"bytes,2,opt,name=meta" json:"meta,omitempty"`
	Data                 *ObjectData `protobuf:"bytes,3,opt,name=data" json:"data,omitempty"`
	PrevVersion          uint64      `protobuf:"varint,5,opt,name=prev_version,json=prevVersion" json:"prev_version,omitempty"`
	PrevDataCheck        uint64      `protobuf:"varint,6,opt,name=prev_data_check,json=prevDataCheck" json:"prev_data_check,omitempty"`
	IncrNamespace        string      `protobuf:"bytes,7,opt,name=incr_namespace,json=incrNamespace" json:"incr_namespace,omitempty"`
	ProposalExpired      uint64      `protobuf:"varint,16,opt,name=proposal_expired,json=proposalExpired" json:"proposal_expired,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ObjectWriter) Reset()         { *m = ObjectWriter{} }
func (m *ObjectWriter) String() string { return proto.CompactTextString(m) }
func (*ObjectWriter) ProtoMessage()    {}
func (*ObjectWriter) Descriptor() ([]byte, []int) {
	return fileDescriptor_sko_6d21db2e82c33d93, []int{4}
}
func (m *ObjectWriter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectWriter.Unmarshal(m, b)
}
func (m *ObjectWriter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectWriter.Marshal(b, m, deterministic)
}
func (dst *ObjectWriter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectWriter.Merge(dst, src)
}
func (m *ObjectWriter) XXX_Size() int {
	return xxx_messageInfo_ObjectWriter.Size(m)
}
func (m *ObjectWriter) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectWriter.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectWriter proto.InternalMessageInfo

func (m *ObjectWriter) GetMode() uint64 {
	if m != nil {
		return m.Mode
	}
	return 0
}

func (m *ObjectWriter) GetMeta() *ObjectMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *ObjectWriter) GetData() *ObjectData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *ObjectWriter) GetPrevVersion() uint64 {
	if m != nil {
		return m.PrevVersion
	}
	return 0
}

func (m *ObjectWriter) GetPrevDataCheck() uint64 {
	if m != nil {
		return m.PrevDataCheck
	}
	return 0
}

func (m *ObjectWriter) GetIncrNamespace() string {
	if m != nil {
		return m.IncrNamespace
	}
	return ""
}

func (m *ObjectWriter) GetProposalExpired() uint64 {
	if m != nil {
		return m.ProposalExpired
	}
	return 0
}

type ObjectResult struct {
	Status               uint64        `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
	Message              string        `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	Meta                 *ObjectMeta   `protobuf:"bytes,8,opt,name=meta" json:"meta,omitempty"`
	Items                []*ObjectItem `protobuf:"bytes,9,rep,name=items" json:"items,omitempty"`
	Next                 bool          `protobuf:"varint,12,opt,name=next" json:"next,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ObjectResult) Reset()         { *m = ObjectResult{} }
func (m *ObjectResult) String() string { return proto.CompactTextString(m) }
func (*ObjectResult) ProtoMessage()    {}
func (*ObjectResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_sko_6d21db2e82c33d93, []int{5}
}
func (m *ObjectResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectResult.Unmarshal(m, b)
}
func (m *ObjectResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectResult.Marshal(b, m, deterministic)
}
func (dst *ObjectResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectResult.Merge(dst, src)
}
func (m *ObjectResult) XXX_Size() int {
	return xxx_messageInfo_ObjectResult.Size(m)
}
func (m *ObjectResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectResult.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectResult proto.InternalMessageInfo

func (m *ObjectResult) GetStatus() uint64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *ObjectResult) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ObjectResult) GetMeta() *ObjectMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *ObjectResult) GetItems() []*ObjectItem {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *ObjectResult) GetNext() bool {
	if m != nil {
		return m.Next
	}
	return false
}

func init() {
	proto.RegisterType((*ObjectMeta)(nil), "sko.ObjectMeta")
	proto.RegisterType((*ObjectData)(nil), "sko.ObjectData")
	proto.RegisterType((*ObjectItem)(nil), "sko.ObjectItem")
	proto.RegisterType((*ObjectReader)(nil), "sko.ObjectReader")
	proto.RegisterType((*ObjectWriter)(nil), "sko.ObjectWriter")
	proto.RegisterType((*ObjectResult)(nil), "sko.ObjectResult")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Object service

type ObjectClient interface {
	Query(ctx context.Context, in *ObjectReader, opts ...grpc.CallOption) (*ObjectResult, error)
	Commit(ctx context.Context, in *ObjectWriter, opts ...grpc.CallOption) (*ObjectResult, error)
	Prepare(ctx context.Context, in *ObjectWriter, opts ...grpc.CallOption) (*ObjectResult, error)
	Accept(ctx context.Context, in *ObjectWriter, opts ...grpc.CallOption) (*ObjectResult, error)
}

type objectClient struct {
	cc *grpc.ClientConn
}

func NewObjectClient(cc *grpc.ClientConn) ObjectClient {
	return &objectClient{cc}
}

func (c *objectClient) Query(ctx context.Context, in *ObjectReader, opts ...grpc.CallOption) (*ObjectResult, error) {
	out := new(ObjectResult)
	err := grpc.Invoke(ctx, "/sko.Object/Query", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectClient) Commit(ctx context.Context, in *ObjectWriter, opts ...grpc.CallOption) (*ObjectResult, error) {
	out := new(ObjectResult)
	err := grpc.Invoke(ctx, "/sko.Object/Commit", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectClient) Prepare(ctx context.Context, in *ObjectWriter, opts ...grpc.CallOption) (*ObjectResult, error) {
	out := new(ObjectResult)
	err := grpc.Invoke(ctx, "/sko.Object/Prepare", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectClient) Accept(ctx context.Context, in *ObjectWriter, opts ...grpc.CallOption) (*ObjectResult, error) {
	out := new(ObjectResult)
	err := grpc.Invoke(ctx, "/sko.Object/Accept", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Object service

type ObjectServer interface {
	Query(context.Context, *ObjectReader) (*ObjectResult, error)
	Commit(context.Context, *ObjectWriter) (*ObjectResult, error)
	Prepare(context.Context, *ObjectWriter) (*ObjectResult, error)
	Accept(context.Context, *ObjectWriter) (*ObjectResult, error)
}

func RegisterObjectServer(s *grpc.Server, srv ObjectServer) {
	s.RegisterService(&_Object_serviceDesc, srv)
}

func _Object_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjectReader)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sko.Object/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectServer).Query(ctx, req.(*ObjectReader))
	}
	return interceptor(ctx, in, info, handler)
}

func _Object_Commit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjectWriter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectServer).Commit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sko.Object/Commit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectServer).Commit(ctx, req.(*ObjectWriter))
	}
	return interceptor(ctx, in, info, handler)
}

func _Object_Prepare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjectWriter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectServer).Prepare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sko.Object/Prepare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectServer).Prepare(ctx, req.(*ObjectWriter))
	}
	return interceptor(ctx, in, info, handler)
}

func _Object_Accept_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjectWriter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectServer).Accept(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sko.Object/Accept",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectServer).Accept(ctx, req.(*ObjectWriter))
	}
	return interceptor(ctx, in, info, handler)
}

var _Object_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sko.Object",
	HandlerType: (*ObjectServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Query",
			Handler:    _Object_Query_Handler,
		},
		{
			MethodName: "Commit",
			Handler:    _Object_Commit_Handler,
		},
		{
			MethodName: "Prepare",
			Handler:    _Object_Prepare_Handler,
		},
		{
			MethodName: "Accept",
			Handler:    _Object_Accept_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sko.proto",
}

func init() { proto.RegisterFile("sko.proto", fileDescriptor_sko_6d21db2e82c33d93) }

var fileDescriptor_sko_6d21db2e82c33d93 = []byte{
	// 600 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xd1, 0x6e, 0xd3, 0x4a,
	0x10, 0xbd, 0x8e, 0x13, 0x27, 0x9e, 0x26, 0x6d, 0xef, 0xaa, 0x02, 0x0b, 0x84, 0x14, 0x8c, 0x8a,
	0xc2, 0x4b, 0x29, 0xe5, 0x0b, 0x4a, 0x8b, 0x44, 0x1f, 0x68, 0x61, 0x91, 0xca, 0x63, 0xb4, 0xb5,
	0xa7, 0xc5, 0x38, 0xce, 0x5a, 0xbb, 0xeb, 0xaa, 0xe9, 0x1f, 0xf0, 0x17, 0xfc, 0x11, 0xfc, 0x0b,
	0x3f, 0x80, 0x76, 0xd6, 0xdb, 0x26, 0x08, 0x50, 0x1f, 0x78, 0xdb, 0x39, 0x67, 0xce, 0xae, 0xe7,
	0xcc, 0x49, 0x20, 0xd6, 0xa5, 0xdc, 0xa9, 0x95, 0x34, 0x92, 0x85, 0xba, 0x94, 0xe9, 0x8f, 0x00,
	0xe0, 0xe4, 0xec, 0x33, 0x66, 0xe6, 0x2d, 0x1a, 0xc1, 0x36, 0x21, 0x2c, 0x71, 0x91, 0x04, 0xe3,
	0x60, 0x32, 0xe4, 0xf6, 0xc8, 0x12, 0xe8, 0x5f, 0xa2, 0xd2, 0x85, 0x9c, 0x27, 0x9d, 0x71, 0x30,
	0xe9, 0x72, 0x5f, 0xb2, 0xfb, 0xd0, 0x2f, 0xe6, 0x99, 0x9a, 0x16, 0x79, 0x12, 0x12, 0x13, 0xd9,
	0xf2, 0x28, 0xb7, 0x92, 0x4c, 0xa1, 0x30, 0x98, 0x27, 0x5d, 0x27, 0x69, 0x4b, 0xcb, 0x34, 0x75,
	0x4e, 0x4c, 0xcf, 0x31, 0x6d, 0xc9, 0xb6, 0xa0, 0x27, 0x8c, 0x51, 0x3a, 0x89, 0x08, 0x77, 0x85,
	0xed, 0xc7, 0xab, 0xba, 0x50, 0x98, 0x27, 0x6b, 0xae, 0xbf, 0x2d, 0xd9, 0x23, 0x80, 0x5c, 0x18,
	0x31, 0x75, 0xa2, 0x21, 0x91, 0xb1, 0x45, 0xf6, 0x49, 0xe8, 0xe9, 0xec, 0x13, 0x66, 0x65, 0x32,
	0xba, 0xa5, 0x0f, 0x2c, 0x90, 0x1e, 0xfb, 0xa1, 0x0f, 0x85, 0x11, 0xb7, 0x6f, 0x0f, 0x96, 0xdf,
	0xde, 0x82, 0x9e, 0x53, 0xc7, 0x0e, 0xa5, 0xc2, 0xa2, 0x97, 0x62, 0xd6, 0x60, 0x02, 0x64, 0x91,
	0x2b, 0xd2, 0x53, 0x7f, 0xdf, 0x91, 0xc1, 0x8a, 0x3d, 0x81, 0x6e, 0x85, 0x46, 0x90, 0x5f, 0x6b,
	0x7b, 0x1b, 0x3b, 0xd6, 0xf2, 0x5b, 0x8f, 0x39, 0x91, 0xb6, 0xc9, 0x7e, 0x0f, 0x59, 0xb7, 0xda,
	0x64, 0xbf, 0x89, 0x13, 0x99, 0x7e, 0x0b, 0x60, 0xe8, 0x40, 0x8e, 0x22, 0x47, 0xc5, 0x18, 0x74,
	0x2b, 0x99, 0x23, 0x2d, 0xa8, 0xcb, 0xe9, 0x6c, 0xb1, 0x12, 0x17, 0x3a, 0xe9, 0x8c, 0xc3, 0xc9,
	0x90, 0xd3, 0x99, 0x3d, 0x84, 0x78, 0x56, 0x54, 0x85, 0x99, 0xce, 0x9b, 0x8a, 0xc6, 0x0a, 0xf9,
	0x80, 0x80, 0xe3, 0xa6, 0xb2, 0xe6, 0x38, 0x52, 0x17, 0xd7, 0x48, 0xe3, 0x85, 0xdc, 0xb5, 0x7f,
	0x28, 0xae, 0xd1, 0xd2, 0x25, 0x2e, 0xa6, 0xf2, 0xfc, 0x5c, 0xa3, 0x21, 0x6b, 0x87, 0x3c, 0x2e,
	0x71, 0x71, 0x42, 0x80, 0xa7, 0xb3, 0xc6, 0x58, 0x7a, 0x74, 0x43, 0x1f, 0x10, 0x40, 0x97, 0xcb,
	0x0b, 0xaf, 0x5e, 0x77, 0xce, 0xcf, 0xe4, 0x85, 0x53, 0xa7, 0x5f, 0x3a, 0x7e, 0xa2, 0x8f, 0xaa,
	0x30, 0x7f, 0x98, 0xe8, 0x9f, 0x19, 0xc8, 0x1e, 0xc3, 0xb0, 0x56, 0x78, 0x39, 0xf5, 0x11, 0x76,
	0xa9, 0x5b, 0xb3, 0xd8, 0x69, 0x1b, 0xe3, 0xa7, 0xb0, 0x41, 0x2d, 0x4b, 0x79, 0x71, 0x19, 0x1c,
	0x59, 0xf8, 0xd0, 0x67, 0x86, 0x6d, 0xc3, 0x3a, 0xc5, 0x7d, 0x2e, 0x2a, 0xd4, 0xb5, 0xc8, 0x30,
	0xe9, 0x8f, 0x83, 0x49, 0xcc, 0x47, 0x16, 0x3d, 0xf6, 0x20, 0x7b, 0x06, 0x9b, 0xb5, 0x92, 0xb5,
	0xd4, 0x62, 0x36, 0xf5, 0xd9, 0xdd, 0xa4, 0xfb, 0x36, 0x3c, 0xfe, 0xda, 0xc1, 0xe9, 0xd7, 0xa5,
	0xed, 0xea, 0x66, 0x66, 0xd8, 0x3d, 0x88, 0xb4, 0x11, 0xa6, 0xd1, 0xad, 0x1b, 0x6d, 0x65, 0x7f,
	0x06, 0x15, 0x6a, 0x2d, 0x2e, 0x90, 0x2c, 0x89, 0xb9, 0x2f, 0x6f, 0x9c, 0x1a, 0xfc, 0xcd, 0xa9,
	0x6d, 0xe8, 0x15, 0x06, 0x2b, 0x9d, 0xc4, 0xe3, 0xf0, 0x97, 0x2e, 0x9b, 0x57, 0xee, 0x58, 0xbb,
	0x89, 0x39, 0x5e, 0xb9, 0x8d, 0x0f, 0x38, 0x9d, 0xf7, 0xbe, 0x07, 0x10, 0xb9, 0x4e, 0xf6, 0x1c,
	0x7a, 0xef, 0x1b, 0x54, 0x0b, 0xf6, 0xff, 0x92, 0xde, 0xc5, 0xf2, 0xc1, 0x2a, 0x64, 0x67, 0x49,
	0xff, 0x63, 0xbb, 0x10, 0x1d, 0xc8, 0xaa, 0x2a, 0xcc, 0x8a, 0xc2, 0xad, 0xfd, 0xf7, 0x8a, 0x17,
	0xd0, 0x7f, 0xa7, 0xb0, 0x16, 0x0a, 0xef, 0x2c, 0xd9, 0x85, 0x68, 0x3f, 0xcb, 0xb0, 0xbe, 0xf3,
	0x23, 0xaf, 0x3a, 0x6f, 0xc2, 0xb3, 0x88, 0xfe, 0x01, 0x5f, 0xfe, 0x0c, 0x00, 0x00, 0xff, 0xff,
	0x04, 0x2f, 0xa8, 0x2f, 0x0e, 0x05, 0x00, 0x00,
}
