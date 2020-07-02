// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/acm/user.proto

package omo_msp_acm

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for UserService service

func NewUserServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for UserService service

type UserService interface {
	AddOne(ctx context.Context, in *ReqUserAdd, opts ...client.CallOption) (*ReplyUserLink, error)
	GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyUserLink, error)
	RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error)
	GetList(ctx context.Context, in *ReqUserList, opts ...client.CallOption) (*ReplyUserList, error)
	IsPermission(ctx context.Context, in *ReqUserPermission, opts ...client.CallOption) (*ReplyUserPermission, error)
	AppendRole(ctx context.Context, in *ReqLinkRole, opts ...client.CallOption) (*ReplyLinkRole, error)
	SubtractRole(ctx context.Context, in *ReqLinkRole, opts ...client.CallOption) (*ReplyLinkRole, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) AddOne(ctx context.Context, in *ReqUserAdd, opts ...client.CallOption) (*ReplyUserLink, error) {
	req := c.c.NewRequest(c.name, "UserService.AddOne", in)
	out := new(ReplyUserLink)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyUserLink, error) {
	req := c.c.NewRequest(c.name, "UserService.GetOne", in)
	out := new(ReplyUserLink)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "UserService.RemoveOne", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetList(ctx context.Context, in *ReqUserList, opts ...client.CallOption) (*ReplyUserList, error) {
	req := c.c.NewRequest(c.name, "UserService.GetList", in)
	out := new(ReplyUserList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) IsPermission(ctx context.Context, in *ReqUserPermission, opts ...client.CallOption) (*ReplyUserPermission, error) {
	req := c.c.NewRequest(c.name, "UserService.IsPermission", in)
	out := new(ReplyUserPermission)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) AppendRole(ctx context.Context, in *ReqLinkRole, opts ...client.CallOption) (*ReplyLinkRole, error) {
	req := c.c.NewRequest(c.name, "UserService.AppendRole", in)
	out := new(ReplyLinkRole)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) SubtractRole(ctx context.Context, in *ReqLinkRole, opts ...client.CallOption) (*ReplyLinkRole, error) {
	req := c.c.NewRequest(c.name, "UserService.SubtractRole", in)
	out := new(ReplyLinkRole)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	AddOne(context.Context, *ReqUserAdd, *ReplyUserLink) error
	GetOne(context.Context, *RequestInfo, *ReplyUserLink) error
	RemoveOne(context.Context, *RequestInfo, *ReplyInfo) error
	GetList(context.Context, *ReqUserList, *ReplyUserList) error
	IsPermission(context.Context, *ReqUserPermission, *ReplyUserPermission) error
	AppendRole(context.Context, *ReqLinkRole, *ReplyLinkRole) error
	SubtractRole(context.Context, *ReqLinkRole, *ReplyLinkRole) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) error {
	type userService interface {
		AddOne(ctx context.Context, in *ReqUserAdd, out *ReplyUserLink) error
		GetOne(ctx context.Context, in *RequestInfo, out *ReplyUserLink) error
		RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error
		GetList(ctx context.Context, in *ReqUserList, out *ReplyUserList) error
		IsPermission(ctx context.Context, in *ReqUserPermission, out *ReplyUserPermission) error
		AppendRole(ctx context.Context, in *ReqLinkRole, out *ReplyLinkRole) error
		SubtractRole(ctx context.Context, in *ReqLinkRole, out *ReplyLinkRole) error
	}
	type UserService struct {
		userService
	}
	h := &userServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&UserService{h}, opts...))
}

type userServiceHandler struct {
	UserServiceHandler
}

func (h *userServiceHandler) AddOne(ctx context.Context, in *ReqUserAdd, out *ReplyUserLink) error {
	return h.UserServiceHandler.AddOne(ctx, in, out)
}

func (h *userServiceHandler) GetOne(ctx context.Context, in *RequestInfo, out *ReplyUserLink) error {
	return h.UserServiceHandler.GetOne(ctx, in, out)
}

func (h *userServiceHandler) RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error {
	return h.UserServiceHandler.RemoveOne(ctx, in, out)
}

func (h *userServiceHandler) GetList(ctx context.Context, in *ReqUserList, out *ReplyUserList) error {
	return h.UserServiceHandler.GetList(ctx, in, out)
}

func (h *userServiceHandler) IsPermission(ctx context.Context, in *ReqUserPermission, out *ReplyUserPermission) error {
	return h.UserServiceHandler.IsPermission(ctx, in, out)
}

func (h *userServiceHandler) AppendRole(ctx context.Context, in *ReqLinkRole, out *ReplyLinkRole) error {
	return h.UserServiceHandler.AppendRole(ctx, in, out)
}

func (h *userServiceHandler) SubtractRole(ctx context.Context, in *ReqLinkRole, out *ReplyLinkRole) error {
	return h.UserServiceHandler.SubtractRole(ctx, in, out)
}
