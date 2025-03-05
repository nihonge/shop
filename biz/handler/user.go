package handler

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	"github.com/nihonge/tiktok/rpc/user/kitex_gen/user"
	"github.com/nihonge/tiktok/rpc/user/kitex_gen/user/userservice"
)

func Register(ctx context.Context, c *app.RequestContext) {
	//初始化一个req
	req := &user.RegisterReq{}
	//先解析请求数据，看看是否符合标准
	if err := c.Bind(req); err != nil {
		hlog.Errorf("bind failed: %v", err)
		c.String(400, "bind failed")
		return
	}
	//字段是必须的
	if req.Email == "" || req.Password == "" || req.ConfirmPassword == "" {
		hlog.Errorf("参数不全")
		c.String(400, "参数不全")
		return
	}
	hlog.Infof("user/register微服务被调用,email=", req.Email)
	cli, err := userservice.NewClient("user.register", client.WithHostPorts("localhost:8891"))
	if err != nil {
		hlog.Errorf("NewClient failed: %v", err)
		c.String(500, "NewClient failed")
		return
	}
	fmt.Println("开始发送请求……")
	resp, err := cli.Register(context.Background(), req)
	if err != nil {
		hlog.Errorf("Register failed: %v", err)
		c.String(500, "Register failed")
		return
	}
	c.JSON(200, resp)
}
