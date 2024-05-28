package util

import (
	"github.com/zeromicro/go-zero/core/logx"
	"mscoin/common/tool"
)

type CaptchaReq struct {
	Id        string `json:"id"`
	SecretKey string `json:"secretkey"`
	Scene     int    `json:"scene"`
	Token     string `json:"token"`
	Ip        string `json:"ip"`
}

type CaptchaResp struct {
	Success int    `json:"success"`
	Score   int    `json:"score"`
	Msg     string `json:"msg"`
}

func CaptchaVerify(server string, vid string, key string, token string, scene int, ip string) bool {
	req := &CaptchaReq{
		Id:        vid,
		SecretKey: key,
		Token:     token,
		Scene:     scene,
		Ip:        ip,
	}
	resp := new(CaptchaResp)

	err := tool.HttpPost(server, req, resp)
	if err != nil {
		logx.Error(err)
		return false
	}

	return resp.Success == 1
}
