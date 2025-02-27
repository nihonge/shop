package test

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/nihonge/tiktok/rpc/auth/kitex_gen/myauth"
	"github.com/nihonge/tiktok/rpc/auth/kitex_gen/myauth/authservice"
)

// 以下测试需要先启动hertz和微服务以及etcd
func TestAuthAPI(t *testing.T) {
	c, err := authservice.NewClient("auth", client.WithHostPorts("localhost:8888"))
	if err != nil {
		log.Fatal(err)
	}
	req := myauth.DeliverTokenReq{
		UserId: 114514,
	}
	resp, err := c.DeliverTokenByRPC(context.Background(), &req, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}
