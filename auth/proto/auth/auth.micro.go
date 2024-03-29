// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/auth/auth.proto

package sxx_micro_book_srv_auth

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

// Client API for Auth service

type AuthService interface {
	MakeAccessToken(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	DelUserAccessToken(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	GetCachedAccessToken(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type authService struct {
	c    client.Client
	name string
}

func NewAuthService(name string, c client.Client) AuthService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "sxx.micro.book.srv.auth"
	}
	return &authService{
		c:    c,
		name: name,
	}
}

func (c *authService) MakeAccessToken(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Auth.MakeAccessToken", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) DelUserAccessToken(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Auth.DelUserAccessToken", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) GetCachedAccessToken(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Auth.GetCachedAccessToken", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Auth service

type AuthHandler interface {
	MakeAccessToken(context.Context, *Request, *Response) error
	DelUserAccessToken(context.Context, *Request, *Response) error
	GetCachedAccessToken(context.Context, *Request, *Response) error
}

func RegisterAuthHandler(s server.Server, hdlr AuthHandler, opts ...server.HandlerOption) error {
	type auth interface {
		MakeAccessToken(ctx context.Context, in *Request, out *Response) error
		DelUserAccessToken(ctx context.Context, in *Request, out *Response) error
		GetCachedAccessToken(ctx context.Context, in *Request, out *Response) error
	}
	type Auth struct {
		auth
	}
	h := &authHandler{hdlr}
	return s.Handle(s.NewHandler(&Auth{h}, opts...))
}

type authHandler struct {
	AuthHandler
}

func (h *authHandler) MakeAccessToken(ctx context.Context, in *Request, out *Response) error {
	return h.AuthHandler.MakeAccessToken(ctx, in, out)
}

func (h *authHandler) DelUserAccessToken(ctx context.Context, in *Request, out *Response) error {
	return h.AuthHandler.DelUserAccessToken(ctx, in, out)
}

func (h *authHandler) GetCachedAccessToken(ctx context.Context, in *Request, out *Response) error {
	return h.AuthHandler.GetCachedAccessToken(ctx, in, out)
}
