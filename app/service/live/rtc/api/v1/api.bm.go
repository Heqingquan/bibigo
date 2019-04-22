// Code generated by protoc-gen-bm v0.1, DO NOT EDIT.
// source: api.proto

/*
Package v1 is a generated blademaster stub package.
This code was generated with go-common/app/tool/bmgen/protoc-gen-bm v0.1.

package 命名使用 {discovery_id}.{version} 的方式, version 形如 v1, v2, v1beta ..
NOTE: 不知道的 discovery_id 请询问大佬, 新项目找大佬申请 discovery_id，先到先得抢注
e.g. account.service.v1

It is generated from these files:
	api.proto
*/
package v1

import (
	"context"

	bm "go-common/library/net/http/blademaster"
	"go-common/library/net/http/blademaster/binding"
)

// to suppressed 'imported but not used warning'
var _ *bm.Context
var _ context.Context
var _ binding.StructValidator

var PathRtcJoinChannel = "/live.rtc.v1.Rtc/JoinChannel"
var PathRtcLeaveChannel = "/live.rtc.v1.Rtc/LeaveChannel"
var PathRtcPublishStream = "/live.rtc.v1.Rtc/PublishStream"
var PathRtcTerminateStream = "/live.rtc.v1.Rtc/TerminateStream"
var PathRtcChannel = "/live.rtc.v1.Rtc/Channel"
var PathRtcStream = "/live.rtc.v1.Rtc/Stream"
var PathRtcSetRtcConfig = "/live.rtc.v1.Rtc/SetRtcConfig"
var PathRtcVerifyToken = "/live.rtc.v1.Rtc/VerifyToken"

// =============
// Rtc Interface
// =============

type RtcBMServer interface {
	// `method:"POST"`
	JoinChannel(ctx context.Context, req *JoinChannelRequest) (resp *JoinChannelResponse, err error)

	// `method:"POST"`
	LeaveChannel(ctx context.Context, req *LeaveChannelRequest) (resp *LeaveChannelResponse, err error)

	// `method:"POST"`
	PublishStream(ctx context.Context, req *PublishStreamRequest) (resp *PublishStreamResponse, err error)

	// `method:"POST"`
	TerminateStream(ctx context.Context, req *TerminateStreamRequest) (resp *TerminateStreamResponse, err error)

	// `method:"GET"`
	Channel(ctx context.Context, req *ChannelRequest) (resp *ChannelResponse, err error)

	// `method:"GET"`
	Stream(ctx context.Context, req *StreamRequest) (resp *StreamResponse, err error)

	// `method:"POST"`
	SetRtcConfig(ctx context.Context, req *SetRtcConfigRequest) (resp *SetRtcConfigResponse, err error)

	// `method:"GET"`
	VerifyToken(ctx context.Context, req *VerifyTokenRequest) (resp *VerifyTokenResponse, err error)
}

var v1RtcSvc RtcBMServer

func rtcJoinChannel(c *bm.Context) {
	p := new(JoinChannelRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1RtcSvc.JoinChannel(c, p)
	c.JSON(resp, err)
}

func rtcLeaveChannel(c *bm.Context) {
	p := new(LeaveChannelRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1RtcSvc.LeaveChannel(c, p)
	c.JSON(resp, err)
}

func rtcPublishStream(c *bm.Context) {
	p := new(PublishStreamRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1RtcSvc.PublishStream(c, p)
	c.JSON(resp, err)
}

func rtcTerminateStream(c *bm.Context) {
	p := new(TerminateStreamRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1RtcSvc.TerminateStream(c, p)
	c.JSON(resp, err)
}

func rtcChannel(c *bm.Context) {
	p := new(ChannelRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1RtcSvc.Channel(c, p)
	c.JSON(resp, err)
}

func rtcStream(c *bm.Context) {
	p := new(StreamRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1RtcSvc.Stream(c, p)
	c.JSON(resp, err)
}

func rtcSetRtcConfig(c *bm.Context) {
	p := new(SetRtcConfigRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1RtcSvc.SetRtcConfig(c, p)
	c.JSON(resp, err)
}

func rtcVerifyToken(c *bm.Context) {
	p := new(VerifyTokenRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1RtcSvc.VerifyToken(c, p)
	c.JSON(resp, err)
}

// RegisterRtcBMServer Register the blademaster route
func RegisterRtcBMServer(e *bm.Engine, server RtcBMServer) {
	v1RtcSvc = server
	e.POST("/live.rtc.v1.Rtc/JoinChannel", rtcJoinChannel)
	e.POST("/live.rtc.v1.Rtc/LeaveChannel", rtcLeaveChannel)
	e.POST("/live.rtc.v1.Rtc/PublishStream", rtcPublishStream)
	e.POST("/live.rtc.v1.Rtc/TerminateStream", rtcTerminateStream)
	e.GET("/live.rtc.v1.Rtc/Channel", rtcChannel)
	e.GET("/live.rtc.v1.Rtc/Stream", rtcStream)
	e.POST("/live.rtc.v1.Rtc/SetRtcConfig", rtcSetRtcConfig)
	e.GET("/live.rtc.v1.Rtc/VerifyToken", rtcVerifyToken)
}