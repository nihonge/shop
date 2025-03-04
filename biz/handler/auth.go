package handler

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	"github.com/nihonge/tiktok/rpc/auth/kitex_gen/myauth"
	"github.com/nihonge/tiktok/rpc/auth/kitex_gen/myauth/authservice"
)

func AuthToken(ctx context.Context, c *app.RequestContext) {
	//初始化一个req
	req := &myauth.DeliverTokenReq{}
	//先解析请求数据，看看是否符合标准
	if err := c.Bind(req); err != nil {
		hlog.Errorf("bind failed: %v", err)
		c.String(400, "bind failed")
		return
	}
	//UserId是必须的
	if req.UserId == 0 {
		hlog.Errorf("UserId is required")
		c.String(400, "UserId is required")
		return
	}
	hlog.Infof("auth/token微服务被调用,userid=", req.UserId)
	cli, err := authservice.NewClient("auth", client.WithHostPorts("localhost:8890"))
	if err != nil {
		hlog.Errorf("NewClient failed: %v", err)
		c.String(500, "NewClient failed")
		return
	}
	fmt.Println("开始发送请求……")
	resp, err := cli.DeliverTokenByRPC(context.Background(), req)
	if err != nil {
		hlog.Errorf("DeliverTokenByRPC failed: %v", err)
		c.String(500, "DeliverTokenByRPC failed")
		return
	}
	c.JSON(200, resp)
}

func AuthVerify(ctx context.Context, c *app.RequestContext) {
	//先解析请求数据，看看是否符合标准
	//初始化一个req
	req := &myauth.VerifyTokenReq{}
	if err := c.Bind(req); err != nil {
		hlog.Errorf("bind failed: %v", err)
		c.String(400, "bind failed")
		return
	}
	//Token是必须的
	if req.Token == "" {
		hlog.Errorf("Token is required")
		c.String(400, "Token is required")
		return
	}
	hlog.Infof("auth/verify微服务被调用,token=", req.Token)
	cli, err := authservice.NewClient("auth", client.WithHostPorts("localhost:8890"))
	if err != nil {
		hlog.Errorf("NewClient failed: %v", err)
		c.String(500, "NewClient failed")
		return
	}
	fmt.Println("开始发送请求……")
	resp, err := cli.VerifyTokenByRPC(context.Background(), req)
	if err != nil {
		hlog.Errorf("VerifyTokenByRPC failed: %v", err)
		c.String(500, "VerifyTokenByRPC failed")
		return
	}
	c.JSON(200, resp)
}
