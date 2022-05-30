// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/acm/scene.proto

package omo_msp_acm

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
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
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for SceneService service

type SceneService interface {
	AddOne(ctx context.Context, in *ReqSceneAdd, opts ...client.CallOption) (*ReplySceneLink, error)
	GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplySceneLink, error)
	RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error)
	GetList(ctx context.Context, in *RequestPage, opts ...client.CallOption) (*ReplySceneList, error)
	UpdateStatus(ctx context.Context, in *ReqSceneStatus, opts ...client.CallOption) (*ReplyInfo, error)
	UpdateModules(ctx context.Context, in *ReqSceneLinks, opts ...client.CallOption) (*ReplySceneLink, error)
}

type sceneService struct {
	c    client.Client
	name string
}

func NewSceneService(name string, c client.Client) SceneService {
	return &sceneService{
		c:    c,
		name: name,
	}
}

func (c *sceneService) AddOne(ctx context.Context, in *ReqSceneAdd, opts ...client.CallOption) (*ReplySceneLink, error) {
	req := c.c.NewRequest(c.name, "SceneService.AddOne", in)
	out := new(ReplySceneLink)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sceneService) GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplySceneLink, error) {
	req := c.c.NewRequest(c.name, "SceneService.GetOne", in)
	out := new(ReplySceneLink)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sceneService) RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "SceneService.RemoveOne", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sceneService) GetList(ctx context.Context, in *RequestPage, opts ...client.CallOption) (*ReplySceneList, error) {
	req := c.c.NewRequest(c.name, "SceneService.GetList", in)
	out := new(ReplySceneList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sceneService) UpdateStatus(ctx context.Context, in *ReqSceneStatus, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "SceneService.UpdateStatus", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sceneService) UpdateModules(ctx context.Context, in *ReqSceneLinks, opts ...client.CallOption) (*ReplySceneLink, error) {
	req := c.c.NewRequest(c.name, "SceneService.UpdateModules", in)
	out := new(ReplySceneLink)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SceneService service

type SceneServiceHandler interface {
	AddOne(context.Context, *ReqSceneAdd, *ReplySceneLink) error
	GetOne(context.Context, *RequestInfo, *ReplySceneLink) error
	RemoveOne(context.Context, *RequestInfo, *ReplyInfo) error
	GetList(context.Context, *RequestPage, *ReplySceneList) error
	UpdateStatus(context.Context, *ReqSceneStatus, *ReplyInfo) error
	UpdateModules(context.Context, *ReqSceneLinks, *ReplySceneLink) error
}

func RegisterSceneServiceHandler(s server.Server, hdlr SceneServiceHandler, opts ...server.HandlerOption) error {
	type sceneService interface {
		AddOne(ctx context.Context, in *ReqSceneAdd, out *ReplySceneLink) error
		GetOne(ctx context.Context, in *RequestInfo, out *ReplySceneLink) error
		RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error
		GetList(ctx context.Context, in *RequestPage, out *ReplySceneList) error
		UpdateStatus(ctx context.Context, in *ReqSceneStatus, out *ReplyInfo) error
		UpdateModules(ctx context.Context, in *ReqSceneLinks, out *ReplySceneLink) error
	}
	type SceneService struct {
		sceneService
	}
	h := &sceneServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&SceneService{h}, opts...))
}

type sceneServiceHandler struct {
	SceneServiceHandler
}

func (h *sceneServiceHandler) AddOne(ctx context.Context, in *ReqSceneAdd, out *ReplySceneLink) error {
	return h.SceneServiceHandler.AddOne(ctx, in, out)
}

func (h *sceneServiceHandler) GetOne(ctx context.Context, in *RequestInfo, out *ReplySceneLink) error {
	return h.SceneServiceHandler.GetOne(ctx, in, out)
}

func (h *sceneServiceHandler) RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error {
	return h.SceneServiceHandler.RemoveOne(ctx, in, out)
}

func (h *sceneServiceHandler) GetList(ctx context.Context, in *RequestPage, out *ReplySceneList) error {
	return h.SceneServiceHandler.GetList(ctx, in, out)
}

func (h *sceneServiceHandler) UpdateStatus(ctx context.Context, in *ReqSceneStatus, out *ReplyInfo) error {
	return h.SceneServiceHandler.UpdateStatus(ctx, in, out)
}

func (h *sceneServiceHandler) UpdateModules(ctx context.Context, in *ReqSceneLinks, out *ReplySceneLink) error {
	return h.SceneServiceHandler.UpdateModules(ctx, in, out)
}
