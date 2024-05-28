package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mscoin/app/usercenter/cmd/api/internal/logic"
	"mscoin/app/usercenter/cmd/api/internal/svc"
	"mscoin/app/usercenter/cmd/api/internal/types"
)

func SendCodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SendCodeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewSendCodeLogic(r.Context(), svcCtx)
		resp, err := l.SendCode(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
