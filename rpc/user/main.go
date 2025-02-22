package main

import (
	"log"
	"net"

	"github.com/cloudwego/kitex/server"
	user "github.com/nihonge/tiktok/rpc/user/kitex_gen/user/userservice"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8891") //设置rpc服务的IP
	svr := user.NewServer(new(UserServiceImpl), server.WithServiceAddr(addr))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
