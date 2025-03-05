package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/nihonge/tiktok/database"
	user "github.com/nihonge/tiktok/rpc/user/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// TODO: Your code here...
	//调用user的GORM方法
	db, err := database.InitializeDB(&database.DBConfig{}) // 你的数据库初始化方法
	if err != nil {
		return nil, err
	}
	userService := database.NewUserService(db)

	_, err = userService.CreateUser(req.Email, req.Password, req.ConfirmPassword)
	if err != nil {
		if strings.Contains(err.Error(), "邮箱已存在") {
			return nil, fmt.Errorf("邮箱已存在")
		}
		return nil, err
	}
	resp = &user.RegisterResp{}
	resp.UserId = 2
	return resp, nil
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginReq) (resp *user.LoginResp, err error) {
	// TODO: Your code here...
	return
}
