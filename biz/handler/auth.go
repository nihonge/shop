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
	fmt.Println(req.UserId)
	hlog.Infof("auth微服务被调用,调用者IP: %s", c.ClientIP())
	cli, err := authservice.NewClient("auth", client.WithHostPorts("localhost:8888"))
	if err != nil {
		hlog.Errorf("NewClient failed: %v", err)
		c.String(500, "NewClient failed")
		return
	}
	cli.DeliverTokenByRPC(context.Background(), nil)
}
