package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"wetalk/apps/user/api/internal/logic"
	"wetalk/apps/user/api/internal/svc"
	"wetalk/apps/user/api/internal/types"
)

func FindUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FindUserRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewFindUserLogic(r.Context(), svcCtx)
		resp, err := l.FindUser(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
