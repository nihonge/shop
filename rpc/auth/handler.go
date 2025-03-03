package main

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	auth "github.com/nihonge/tiktok/rpc/auth/kitex_gen/myauth"
)

// AuthServiceImpl implements the last service interface defined in the IDL.
type AuthServiceImpl struct{}

// DeliverTokenByRPC implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) DeliverTokenByRPC(ctx context.Context, req *auth.DeliverTokenReq) (resp *auth.DeliveryResp, err error) {
	// TODO: Your code here...
	klog.Infof("auth微服务被调用")
	//获取request中的userid
	userid := req.UserId
	klog.Infof("userid: %d", userid)
	//生成token
	_ = userid
	//存储token
	//返回token
	resp = &auth.DeliveryResp{}
	resp.Token = "nihonge"
	return resp, nil
}

// VerifyTokenByRPC implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) VerifyTokenByRPC(ctx context.Context, req *auth.VerifyTokenReq) (resp *auth.VerifyResp, err error) {
	// TODO: Your code here...
	return
}
