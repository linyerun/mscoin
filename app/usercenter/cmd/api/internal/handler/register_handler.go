package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mscoin/app/usercenter/cmd/api/internal/logic"
	"mscoin/app/usercenter/cmd/api/internal/svc"
	"mscoin/app/usercenter/cmd/api/internal/types"
)

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
