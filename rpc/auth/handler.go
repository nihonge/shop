package main

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	auth "github.com/nihonge/tiktok/rpc/auth/kitex_gen/myauth"
)

// AuthServiceImpl implements the last service interface defined in the IDL.
type AuthServiceImpl struct {
}

// DeliverTokenByRPC implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) DeliverTokenByRPC(ctx context.Context, req *auth.DeliverTokenReq) (resp *auth.DeliveryResp, err error) {
	// TODO: Your code here...
	klog.Infof("auth微服务DeliverTokenByRPC被调用")
	//获取request中的userid
	userid := req.UserId
	klog.Infof("userid: %d", userid)
	//生成token
	token, err := GenerateToken(int64(userid))
	if err != nil {
		klog.Errorf("生成token失败:%v", err)
		return nil, err
	}
	klog.Infof("生成token:%s", token)
	//返回token
	resp = &auth.DeliveryResp{}
	resp.Token = token
	return resp, nil
}

// VerifyTokenByRPC implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) VerifyTokenByRPC(ctx context.Context, req *auth.VerifyTokenReq) (resp *auth.VerifyResp, err error) {
	// TODO: Your code here...
	klog.Infof("auth微服务VerifyTokenByRPC被调用")
	//获取request中的token
	token := req.Token
	klog.Infof("token: %s", token)
	//定义返回的resp
	resp = &auth.VerifyResp{}
	//验证token
	claims, err := VerifyToken(token)
	if err != nil {
		klog.Errorf("验证token失败:%v", err)
		resp.Res = false
		return nil, err
	}
	klog.Infof("验证token成功")
	klog.Infof("claims: %v", claims)
	//返回验证结果
	resp.Res = true
	return resp, nil
}
