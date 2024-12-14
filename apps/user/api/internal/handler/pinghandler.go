package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"wetalk/apps/user/api/internal/logic"
	"wetalk/apps/user/api/internal/svc"
	"wetalk/apps/user/api/internal/types"
)

func PingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewPingLogic(r.Context(), svcCtx)
		resp, err := l.Ping(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
