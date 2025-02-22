package handler

import (
	"context"
	"log"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/client"
	"github.com/nihonge/tiktok/rpc/auth/kitex_gen/auth"
	"github.com/nihonge/tiktok/rpc/auth/kitex_gen/auth/authservice"
)

func Auth(ctx context.Context, c *app.RequestContext) {
	cli, err := authservice.NewClient("auth", client.WithHostPorts("0.0.0.0:8890"))
	if err != nil {
		log.Fatal(err)
	}
	req := auth.DeliverTokenReq{}
	req.UserId = 114514
	resp, err := cli.DeliverTokenByRPC(ctx, &req)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, resp)
}
