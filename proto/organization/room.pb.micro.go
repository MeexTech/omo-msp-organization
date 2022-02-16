// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/organization/room.proto

package organization

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

// Client API for RoomService service

type RoomService interface {
	AddOne(ctx context.Context, in *ReqRoomAdd, opts ...client.CallOption) (*ReplyRoomInfo, error)
	UpdateBase(ctx context.Context, in *ReqRoomUpdate, opts ...client.CallOption) (*ReplyInfo, error)
	RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error)
	GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyRoomInfo, error)
	GetList(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyRoomList, error)
	GetStatistic(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyStatistic, error)
	AppendDevice(ctx context.Context, in *ReqRoomDevice, opts ...client.CallOption) (*ReplyRoomDevices, error)
	SubtractDevice(ctx context.Context, in *ReqRoomDevice, opts ...client.CallOption) (*ReplyRoomDevices, error)
}

type roomService struct {
	c    client.Client
	name string
}

func NewRoomService(name string, c client.Client) RoomService {
	return &roomService{
		c:    c,
		name: name,
	}
}

func (c *roomService) AddOne(ctx context.Context, in *ReqRoomAdd, opts ...client.CallOption) (*ReplyRoomInfo, error) {
	req := c.c.NewRequest(c.name, "RoomService.AddOne", in)
	out := new(ReplyRoomInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomService) UpdateBase(ctx context.Context, in *ReqRoomUpdate, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "RoomService.UpdateBase", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomService) RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "RoomService.RemoveOne", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomService) GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyRoomInfo, error) {
	req := c.c.NewRequest(c.name, "RoomService.GetOne", in)
	out := new(ReplyRoomInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomService) GetList(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyRoomList, error) {
	req := c.c.NewRequest(c.name, "RoomService.GetList", in)
	out := new(ReplyRoomList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomService) GetStatistic(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyStatistic, error) {
	req := c.c.NewRequest(c.name, "RoomService.GetStatistic", in)
	out := new(ReplyStatistic)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomService) AppendDevice(ctx context.Context, in *ReqRoomDevice, opts ...client.CallOption) (*ReplyRoomDevices, error) {
	req := c.c.NewRequest(c.name, "RoomService.AppendDevice", in)
	out := new(ReplyRoomDevices)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomService) SubtractDevice(ctx context.Context, in *ReqRoomDevice, opts ...client.CallOption) (*ReplyRoomDevices, error) {
	req := c.c.NewRequest(c.name, "RoomService.SubtractDevice", in)
	out := new(ReplyRoomDevices)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RoomService service

type RoomServiceHandler interface {
	AddOne(context.Context, *ReqRoomAdd, *ReplyRoomInfo) error
	UpdateBase(context.Context, *ReqRoomUpdate, *ReplyInfo) error
	RemoveOne(context.Context, *RequestInfo, *ReplyInfo) error
	GetOne(context.Context, *RequestInfo, *ReplyRoomInfo) error
	GetList(context.Context, *RequestFilter, *ReplyRoomList) error
	GetStatistic(context.Context, *RequestFilter, *ReplyStatistic) error
	AppendDevice(context.Context, *ReqRoomDevice, *ReplyRoomDevices) error
	SubtractDevice(context.Context, *ReqRoomDevice, *ReplyRoomDevices) error
}

func RegisterRoomServiceHandler(s server.Server, hdlr RoomServiceHandler, opts ...server.HandlerOption) error {
	type roomService interface {
		AddOne(ctx context.Context, in *ReqRoomAdd, out *ReplyRoomInfo) error
		UpdateBase(ctx context.Context, in *ReqRoomUpdate, out *ReplyInfo) error
		RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error
		GetOne(ctx context.Context, in *RequestInfo, out *ReplyRoomInfo) error
		GetList(ctx context.Context, in *RequestFilter, out *ReplyRoomList) error
		GetStatistic(ctx context.Context, in *RequestFilter, out *ReplyStatistic) error
		AppendDevice(ctx context.Context, in *ReqRoomDevice, out *ReplyRoomDevices) error
		SubtractDevice(ctx context.Context, in *ReqRoomDevice, out *ReplyRoomDevices) error
	}
	type RoomService struct {
		roomService
	}
	h := &roomServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&RoomService{h}, opts...))
}

type roomServiceHandler struct {
	RoomServiceHandler
}

func (h *roomServiceHandler) AddOne(ctx context.Context, in *ReqRoomAdd, out *ReplyRoomInfo) error {
	return h.RoomServiceHandler.AddOne(ctx, in, out)
}

func (h *roomServiceHandler) UpdateBase(ctx context.Context, in *ReqRoomUpdate, out *ReplyInfo) error {
	return h.RoomServiceHandler.UpdateBase(ctx, in, out)
}

func (h *roomServiceHandler) RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error {
	return h.RoomServiceHandler.RemoveOne(ctx, in, out)
}

func (h *roomServiceHandler) GetOne(ctx context.Context, in *RequestInfo, out *ReplyRoomInfo) error {
	return h.RoomServiceHandler.GetOne(ctx, in, out)
}

func (h *roomServiceHandler) GetList(ctx context.Context, in *RequestFilter, out *ReplyRoomList) error {
	return h.RoomServiceHandler.GetList(ctx, in, out)
}

func (h *roomServiceHandler) GetStatistic(ctx context.Context, in *RequestFilter, out *ReplyStatistic) error {
	return h.RoomServiceHandler.GetStatistic(ctx, in, out)
}

func (h *roomServiceHandler) AppendDevice(ctx context.Context, in *ReqRoomDevice, out *ReplyRoomDevices) error {
	return h.RoomServiceHandler.AppendDevice(ctx, in, out)
}

func (h *roomServiceHandler) SubtractDevice(ctx context.Context, in *ReqRoomDevice, out *ReplyRoomDevices) error {
	return h.RoomServiceHandler.SubtractDevice(ctx, in, out)
}
