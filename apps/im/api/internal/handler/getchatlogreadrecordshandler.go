package handler

import (
	"net/http"
	"wetalk/apps/im/api/internal/logic"
	"wetalk/apps/im/api/internal/svc"
	"wetalk/apps/im/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取已读记录
func getChatLogReadRecordsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetChatLogReadRecordsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetChatLogReadRecordsLogic(r.Context(), svcCtx)
		resp, err := l.GetChatLogReadRecords(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
