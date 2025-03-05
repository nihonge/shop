// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/nihonge/tiktok/biz/handler"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	//认证服务url设置
	r.POST("/auth/token", handler.AuthToken)
	r.POST("/auth/verify", handler.AuthVerify)

	r.POST("/user/register", handler.AuthVerify)
	r.POST("/user/login", handler.AuthVerify)

	r.GET("/ping", handler.Ping)

}
