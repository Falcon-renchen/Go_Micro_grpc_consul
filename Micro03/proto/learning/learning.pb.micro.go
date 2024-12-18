// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/learning/learning.proto

package learning

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Greeter service

type GreeterService interface {
	Hi(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type greeterService struct {
	c    client.Client
	name string
}

func NewGreeterService(name string, c client.Client) GreeterService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "greeter"
	}
	return &greeterService{
		c:    c,
		name: name,
	}
}

func (c *greeterService) Hi(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Greeter.Hi", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Greeter service

type GreeterHandler interface {
	Hi(context.Context, *Request, *Response) error
}

func RegisterGreeterHandler(s server.Server, hdlr GreeterHandler, opts ...server.HandlerOption) error {
	type greeter interface {
		Hi(ctx context.Context, in *Request, out *Response) error
	}
	type Greeter struct {
		greeter
	}
	h := &greeterHandler{hdlr}
	return s.Handle(s.NewHandler(&Greeter{h}, opts...))
}

type greeterHandler struct {
	GreeterHandler
}

func (h *greeterHandler) Hi(ctx context.Context, in *Request, out *Response) error {
	return h.GreeterHandler.Hi(ctx, in, out)
}

// Client API for Learning service

type LearningService interface {
	Hi(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type learningService struct {
	c    client.Client
	name string
}

func NewLearningService(name string, c client.Client) LearningService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "learning"
	}
	return &learningService{
		c:    c,
		name: name,
	}
}

func (c *learningService) Hi(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Learning.Hi", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Learning service

type LearningHandler interface {
	Hi(context.Context, *Request, *Response) error
}

func RegisterLearningHandler(s server.Server, hdlr LearningHandler, opts ...server.HandlerOption) error {
	type learning interface {
		Hi(ctx context.Context, in *Request, out *Response) error
	}
	type Learning struct {
		learning
	}
	h := &learningHandler{hdlr}
	return s.Handle(s.NewHandler(&Learning{h}, opts...))
}

type learningHandler struct {
	LearningHandler
}

func (h *learningHandler) Hi(ctx context.Context, in *Request, out *Response) error {
	return h.LearningHandler.Hi(ctx, in, out)
}
