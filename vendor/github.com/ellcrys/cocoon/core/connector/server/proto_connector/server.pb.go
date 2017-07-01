// Code generated by protoc-gen-gogo.
// source: server.proto
// DO NOT EDIT!

/*
Package proto_connector is a generated protocol buffer package.

It is generated from these files:
	server.proto

It has these top-level messages:
	Request
	LockOperation
	LedgerOperation
	CocoonCodeOperation
	Response
*/
package proto_connector

import proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// OpType represents operation types
type OpType int32

const (
	OpType_LedgerOp     OpType = 0
	OpType_CocoonCodeOp OpType = 1
	OpType_LockOp       OpType = 2
)

var OpType_name = map[int32]string{
	0: "LedgerOp",
	1: "CocoonCodeOp",
	2: "LockOp",
}
var OpType_value = map[string]int32{
	"LedgerOp":     0,
	"CocoonCodeOp": 1,
	"LockOp":       2,
}

func (x OpType) String() string {
	return proto.EnumName(OpType_name, int32(x))
}
func (OpType) EnumDescriptor() ([]byte, []int) { return fileDescriptorServer, []int{0} }

// Request represents a transaction request which can be
// either a ledger or cocoon code operation
type Request struct {
	OpType       OpType               `protobuf:"varint,1,opt,name=opType,proto3,enum=proto_connector.OpType" json:"opType,omitempty"`
	LedgerOp     *LedgerOperation     `protobuf:"bytes,2,opt,name=ledgerOp" json:"ledgerOp,omitempty"`
	CocoonCodeOp *CocoonCodeOperation `protobuf:"bytes,3,opt,name=cocoonCodeOp" json:"cocoonCodeOp,omitempty"`
	LockOp       *LockOperation       `protobuf:"bytes,4,opt,name=lockOp" json:"lockOp,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptorServer, []int{0} }

func (m *Request) GetOpType() OpType {
	if m != nil {
		return m.OpType
	}
	return OpType_LedgerOp
}

func (m *Request) GetLedgerOp() *LedgerOperation {
	if m != nil {
		return m.LedgerOp
	}
	return nil
}

func (m *Request) GetCocoonCodeOp() *CocoonCodeOperation {
	if m != nil {
		return m.CocoonCodeOp
	}
	return nil
}

func (m *Request) GetLockOp() *LockOperation {
	if m != nil {
		return m.LockOp
	}
	return nil
}

// LockOperation represents a key locking operation within a cocoon's scope.
type LockOperation struct {
	Name   string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Params []string `protobuf:"bytes,2,rep,name=params" json:"params,omitempty"`
	LinkTo string   `protobuf:"bytes,3,opt,name=linkTo,proto3" json:"linkTo,omitempty"`
}

func (m *LockOperation) Reset()                    { *m = LockOperation{} }
func (m *LockOperation) String() string            { return proto.CompactTextString(m) }
func (*LockOperation) ProtoMessage()               {}
func (*LockOperation) Descriptor() ([]byte, []int) { return fileDescriptorServer, []int{1} }

func (m *LockOperation) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LockOperation) GetParams() []string {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *LockOperation) GetLinkTo() string {
	if m != nil {
		return m.LinkTo
	}
	return ""
}

// LedgerOperation represents an operation against the ledger
type LedgerOperation struct {
	ID     string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name   string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Params []string `protobuf:"bytes,3,rep,name=params" json:"params,omitempty"`
	LinkTo string   `protobuf:"bytes,4,opt,name=linkTo,proto3" json:"linkTo,omitempty"`
	Body   []byte   `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *LedgerOperation) Reset()                    { *m = LedgerOperation{} }
func (m *LedgerOperation) String() string            { return proto.CompactTextString(m) }
func (*LedgerOperation) ProtoMessage()               {}
func (*LedgerOperation) Descriptor() ([]byte, []int) { return fileDescriptorServer, []int{2} }

func (m *LedgerOperation) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *LedgerOperation) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LedgerOperation) GetParams() []string {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *LedgerOperation) GetLinkTo() string {
	if m != nil {
		return m.LinkTo
	}
	return ""
}

func (m *LedgerOperation) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

// CocoonCodeOperation represents a cocoon code invoke operation
type CocoonCodeOperation struct {
	ID       string            `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Function string            `protobuf:"bytes,2,opt,name=function,proto3" json:"function,omitempty"`
	Params   []string          `protobuf:"bytes,3,rep,name=params" json:"params,omitempty"`
	Header   map[string]string `protobuf:"bytes,4,rep,name=header" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *CocoonCodeOperation) Reset()                    { *m = CocoonCodeOperation{} }
func (m *CocoonCodeOperation) String() string            { return proto.CompactTextString(m) }
func (*CocoonCodeOperation) ProtoMessage()               {}
func (*CocoonCodeOperation) Descriptor() ([]byte, []int) { return fileDescriptorServer, []int{3} }

func (m *CocoonCodeOperation) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *CocoonCodeOperation) GetFunction() string {
	if m != nil {
		return m.Function
	}
	return ""
}

func (m *CocoonCodeOperation) GetParams() []string {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *CocoonCodeOperation) GetHeader() map[string]string {
	if m != nil {
		return m.Header
	}
	return nil
}

// Response represents the response
type Response struct {
	ID     string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Status int32  `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Body   []byte `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptorServer, []int{4} }

func (m *Response) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Response) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Response) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "proto_connector.Request")
	proto.RegisterType((*LockOperation)(nil), "proto_connector.LockOperation")
	proto.RegisterType((*LedgerOperation)(nil), "proto_connector.LedgerOperation")
	proto.RegisterType((*CocoonCodeOperation)(nil), "proto_connector.CocoonCodeOperation")
	proto.RegisterType((*Response)(nil), "proto_connector.Response")
	proto.RegisterEnum("proto_connector.OpType", OpType_name, OpType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Connector service

type ConnectorClient interface {
	Transact(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type connectorClient struct {
	cc *grpc.ClientConn
}

func NewConnectorClient(cc *grpc.ClientConn) ConnectorClient {
	return &connectorClient{cc}
}

func (c *connectorClient) Transact(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/proto_connector.Connector/Transact", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Connector service

type ConnectorServer interface {
	Transact(context.Context, *Request) (*Response, error)
}

func RegisterConnectorServer(s *grpc.Server, srv ConnectorServer) {
	s.RegisterService(&_Connector_serviceDesc, srv)
}

func _Connector_Transact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorServer).Transact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_connector.Connector/Transact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorServer).Transact(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Connector_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto_connector.Connector",
	HandlerType: (*ConnectorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Transact",
			Handler:    _Connector_Transact_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server.proto",
}

func init() { proto.RegisterFile("server.proto", fileDescriptorServer) }

var fileDescriptorServer = []byte{
	// 443 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x51, 0x6f, 0xd3, 0x30,
	0x10, 0x26, 0x49, 0x1b, 0xd2, 0x6b, 0xd8, 0xa2, 0x03, 0x0d, 0xd3, 0x07, 0x14, 0x45, 0x3c, 0x44,
	0x3c, 0x14, 0x54, 0x10, 0x02, 0x84, 0xc4, 0x43, 0x07, 0xea, 0xa4, 0x49, 0x95, 0x4c, 0xdf, 0x91,
	0x97, 0x1a, 0x98, 0xda, 0xd9, 0xc6, 0x76, 0x27, 0xe5, 0xcf, 0xf2, 0x43, 0x78, 0x42, 0x75, 0x9c,
	0xd2, 0x2d, 0x9d, 0xb4, 0xa7, 0xdc, 0xf9, 0xbb, 0xef, 0xbb, 0xef, 0xce, 0x0e, 0xa4, 0x86, 0xeb,
	0x6b, 0xae, 0xc7, 0x4a, 0x4b, 0x2b, 0xf1, 0xd8, 0x7d, 0xbe, 0x57, 0x52, 0x08, 0x5e, 0x59, 0xa9,
	0x8b, 0xbf, 0x01, 0x3c, 0xa4, 0xfc, 0xf7, 0x86, 0x1b, 0x8b, 0xaf, 0x20, 0x96, 0x6a, 0x51, 0x2b,
	0x4e, 0x82, 0x3c, 0x28, 0x8f, 0x26, 0x4f, 0xc7, 0xb7, 0xaa, 0xc7, 0x73, 0x07, 0x53, 0x5f, 0x86,
	0x9f, 0x20, 0x59, 0xf3, 0xe5, 0x4f, 0xae, 0xe7, 0x8a, 0x84, 0x79, 0x50, 0x0e, 0x27, 0x79, 0x87,
	0x72, 0xee, 0x0b, 0xb8, 0x66, 0xf6, 0x52, 0x0a, 0xba, 0x63, 0xe0, 0x0c, 0xd2, 0x4a, 0x56, 0x52,
	0x8a, 0xa9, 0x5c, 0xf2, 0xb9, 0x22, 0x91, 0x53, 0x78, 0xd1, 0x51, 0x98, 0xee, 0x15, 0xb5, 0x2a,
	0x37, 0x98, 0xf8, 0x0e, 0xe2, 0xb5, 0xac, 0x56, 0x73, 0x45, 0x7a, 0x4e, 0xe3, 0x79, 0xd7, 0x85,
	0x83, 0x5b, 0xb6, 0xaf, 0x2e, 0xbe, 0xc1, 0xa3, 0x1b, 0x00, 0x22, 0xf4, 0x04, 0xbb, 0x6a, 0xe6,
	0x1f, 0x50, 0x17, 0xe3, 0x09, 0xc4, 0x8a, 0x69, 0x76, 0x65, 0x48, 0x98, 0x47, 0xe5, 0x80, 0xfa,
	0x6c, 0x7b, 0xbe, 0xbe, 0x14, 0xab, 0x85, 0x74, 0xc6, 0x07, 0xd4, 0x67, 0x45, 0x0d, 0xc7, 0xb7,
	0x66, 0xc6, 0x23, 0x08, 0xcf, 0x4e, 0xbd, 0x68, 0x78, 0x76, 0xba, 0x6b, 0x13, 0x1e, 0x6c, 0x13,
	0xdd, 0xd1, 0xa6, 0xb7, 0xdf, 0x66, 0xab, 0x71, 0x21, 0x97, 0x35, 0xe9, 0xe7, 0x41, 0x99, 0x52,
	0x17, 0x17, 0x7f, 0x02, 0x78, 0x7c, 0x60, 0x5b, 0x9d, 0xfe, 0x23, 0x48, 0x7e, 0x6c, 0x44, 0xb5,
	0xc5, 0xbc, 0x87, 0x5d, 0x7e, 0xa7, 0x8f, 0x19, 0xc4, 0xbf, 0x38, 0x5b, 0x72, 0x4d, 0x7a, 0x79,
	0x54, 0x0e, 0x27, 0xaf, 0xef, 0x73, 0x4f, 0xe3, 0x99, 0xa3, 0x7c, 0x11, 0x56, 0xd7, 0xd4, 0xf3,
	0x47, 0x1f, 0x60, 0xb8, 0x77, 0x8c, 0x19, 0x44, 0x2b, 0x5e, 0x7b, 0x77, 0xdb, 0x10, 0x9f, 0x40,
	0xff, 0x9a, 0xad, 0x37, 0xed, 0x7e, 0x9a, 0xe4, 0x63, 0xf8, 0x3e, 0x28, 0xbe, 0x42, 0x42, 0xb9,
	0x51, 0x52, 0x18, 0xde, 0x19, 0xea, 0x04, 0x62, 0x63, 0x99, 0xdd, 0x18, 0x47, 0xeb, 0x53, 0x9f,
	0xed, 0x16, 0x15, 0xfd, 0x5f, 0xd4, 0xcb, 0xb7, 0x10, 0x37, 0x4f, 0x19, 0x53, 0x48, 0xda, 0xdb,
	0xca, 0x1e, 0x60, 0x06, 0xe9, 0xfe, 0x14, 0x59, 0x80, 0x00, 0x71, 0xf3, 0x44, 0xb2, 0x70, 0x72,
	0x0e, 0x83, 0x69, 0x3b, 0x2d, 0x7e, 0x86, 0x64, 0xa1, 0x99, 0x30, 0xac, 0xb2, 0x48, 0x3a, 0xbb,
	0xf0, 0xbf, 0xd4, 0xe8, 0xd9, 0x01, 0xa4, 0xf1, 0x7f, 0x11, 0x3b, 0xe4, 0xcd, 0xbf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xe4, 0x1d, 0x19, 0x15, 0xa1, 0x03, 0x00, 0x00,
}