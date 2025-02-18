package main

import (
	"log"
	"net"

	"github.com/cloudwego/kitex/server"
	auth "github.com/nihonge/tiktok/rpc/auth/kitex_gen/auth/authservice"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8890") //设置rpc服务的IP
	svr := auth.NewServer(new(AuthServiceImpl), server.WithServiceAddr(addr))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
