package main

import (
	"context"

	auth "github.com/nihonge/tiktok/rpc/auth/kitex_gen/auth"
)

// AuthServiceImpl implements the last service interface defined in the IDL.
type AuthServiceImpl struct{}

// DeliverTokenByRPC implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) DeliverTokenByRPC(ctx context.Context, req *auth.DeliverTokenReq) (resp *auth.DeliveryResp, err error) {
	// TODO: Your code here...
	//获取request中的userid
	userid := req.UserId
	//生成token
	_ = userid
	//存储token
	//返回token
	resp = &auth.DeliveryResp{}
	resp.Token = "nihonge"
	return
}

// VerifyTokenByRPC implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) VerifyTokenByRPC(ctx context.Context, req *auth.VerifyTokenReq) (resp *auth.VerifyResp, err error) {
	// TODO: Your code here...
	return
}
