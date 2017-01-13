// Code generated by protoc-gen-go.
// source: control.proto
// DO NOT EDIT!

/*
Package control is a generated protocol buffer package.

It is generated from these files:
	control.proto

It has these top-level messages:
	OpenRequest
	OpenResponse
*/
package control

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type OpenRequest struct {
	Url string `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
}

func (m *OpenRequest) Reset()                    { *m = OpenRequest{} }
func (m *OpenRequest) String() string            { return proto.CompactTextString(m) }
func (*OpenRequest) ProtoMessage()               {}
func (*OpenRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type OpenResponse struct {
	Error string `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
}

func (m *OpenResponse) Reset()                    { *m = OpenResponse{} }
func (m *OpenResponse) String() string            { return proto.CompactTextString(m) }
func (*OpenResponse) ProtoMessage()               {}
func (*OpenResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*OpenRequest)(nil), "OpenRequest")
	proto.RegisterType((*OpenResponse)(nil), "OpenResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Control service

type ControlClient interface {
	Open(ctx context.Context, in *OpenRequest, opts ...client.CallOption) (*OpenResponse, error)
}

type controlClient struct {
	c           client.Client
	serviceName string
}

func NewControlClient(serviceName string, c client.Client) ControlClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "control"
	}
	return &controlClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *controlClient) Open(ctx context.Context, in *OpenRequest, opts ...client.CallOption) (*OpenResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Control.Open", in)
	out := new(OpenResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Control service

type ControlHandler interface {
	Open(context.Context, *OpenRequest, *OpenResponse) error
}

func RegisterControlHandler(s server.Server, hdlr ControlHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Control{hdlr}, opts...))
}

type Control struct {
	ControlHandler
}

func (h *Control) Open(ctx context.Context, in *OpenRequest, out *OpenResponse) error {
	return h.ControlHandler.Open(ctx, in, out)
}

func init() { proto.RegisterFile("control.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 125 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xce, 0xcf, 0x2b,
	0x29, 0xca, 0xcf, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x92, 0xe7, 0xe2, 0xf6, 0x2f, 0x48,
	0xcd, 0x0b, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe0, 0x62, 0x2e, 0x2d, 0xca, 0x91,
	0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0x31, 0x95, 0x54, 0xb8, 0x78, 0x20, 0x0a, 0x8a, 0x0b,
	0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x44, 0xb8, 0x58, 0x53, 0x8b, 0x8a, 0xf2, 0x8b, 0xa0, 0x6a, 0x20,
	0x1c, 0x23, 0x03, 0x2e, 0x76, 0x67, 0x88, 0xb9, 0x42, 0xaa, 0x5c, 0x2c, 0x20, 0x0d, 0x42, 0x3c,
	0x7a, 0x48, 0x06, 0x4b, 0xf1, 0xea, 0x21, 0x9b, 0xa2, 0xc4, 0x90, 0xc4, 0x06, 0xb6, 0xdf, 0x18,
	0x10, 0x00, 0x00, 0xff, 0xff, 0xfb, 0x03, 0xcc, 0x0d, 0x90, 0x00, 0x00, 0x00,
}