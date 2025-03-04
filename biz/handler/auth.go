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

func Auth(ctx context.Context, c *app.RequestContext) {
	//先解析请求数据，看看是否符合标准
	//初始化一个req
	req := &myauth.DeliverTokenReq{}
	if err := c.Bind(req); err != nil {
		hlog.Errorf("bind failed: %v", err)
		c.String(400, "bind failed")
		return
	}
	hlog.Infof("auth微服务被调用,userid=", req.UserId)
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
