package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/cloudwego/kitex/client"
	"github.com/nihonge/tiktok/rpc/auth/kitex_gen/myauth"
	"github.com/nihonge/tiktok/rpc/auth/kitex_gen/myauth/authservice"
)

func TestDeliverTokenByRPC(t *testing.T) {
	c, err := authservice.NewClient("myauth", client.WithHostPorts("localhost:8888"))
	if err != nil {
		fmt.Println("NewClient failed:", err)
		return
	}
	req := &myauth.DeliverTokenReq{
		UserId: 114514,
	}
	resp, err := c.DeliverTokenByRPC(context.Background(), req)
	if err != nil {
		fmt.Println("DeliverTokenByRPC failed:", err)
		return
	}
	fmt.Println(resp)
}
