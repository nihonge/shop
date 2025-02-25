package main

import (
	"log"
	"net"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	auth "github.com/nihonge/tiktok/rpc/auth/kitex_gen/myauth/authservice"
)

func main() {
	//接入服务注册中心
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Println(err)
		return
	}
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8890") //设置rpc服务的IP
	svr := auth.NewServer(new(AuthServiceImpl),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithServerBasicInfo(
			&rpcinfo.EndpointBasicInfo{
				ServiceName: "myauth",
			},
		),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
