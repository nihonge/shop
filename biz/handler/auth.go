package handler

import (
	"context"
	"fmt"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	"github.com/nihonge/tiktok/rpc/auth/kitex_gen/myauth"
	"github.com/nihonge/tiktok/rpc/auth/kitex_gen/myauth/authservice"
)

func Auth(ctx context.Context, c *app.RequestContext) {
	userIDstr := c.Query("user_id") // 假设参数名是 user_id（小写下划线）
	if userIDstr == "" {
		c.String(400, "user_id is required")
		return
	}
	fmt.Println(userIDstr)
	// 将字符串转换为 int32
	userID, err := strconv.ParseInt(userIDstr, 10, 32)
	if err != nil {
		fmt.Println("转换失败:", err)
		return
	}
	//先解析请求数据，看看是否符合标准
	//初始化一个req
	req := &myauth.DeliverTokenReq{}
	if err := c.Bind(req); err != nil {
		hlog.Errorf("bind failed: %v", err)
		c.String(400, "bind failed")
		return
	}
	req.UserId = int32(userID)
	//如果userid是0，说明没有传userid，返回错误
	if req.UserId == 0 {
		c.String(400, "userid is required")
		return
	}
	hlog.Infof("auth微服务被调用,调用者IP: %s", c.ClientIP())
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
