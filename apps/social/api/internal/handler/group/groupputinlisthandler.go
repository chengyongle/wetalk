package group

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"wetalk/apps/social/api/internal/logic/group"
	"wetalk/apps/social/api/internal/svc"
	"wetalk/apps/social/api/internal/types"
)

// 申请进群列表
func GroupPutinListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GroupPutinListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := group.NewGroupPutinListLogic(r.Context(), svcCtx)
		resp, err := l.GroupPutinList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
