// Code generated by hertz generator.

package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/nihonge/tiktok/biz/handler"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	r.GET("/ping", handler.Ping)
	r.GET("/ping1", func(c context.Context, ctx *app.RequestContext) {
		ctx.String(200, "lalala~")
	})
	v1 := r.Group("/1")
	{
		v1.GET("/1", func(c context.Context, ctx *app.RequestContext) {
			ctx.String(200, "wdnmd1")
		})
		v1.GET("/2", func(c context.Context, ctx *app.RequestContext) {
			ctx.String(200, "fuckyou2")
		})
	}
	// your code ...
	r.Static("/", "./pages/index.html")
}
