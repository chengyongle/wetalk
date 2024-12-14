package logic

import (
	"context"
	"github.com/pkg/errors"
	"wetalk/apps/im/rpc/im"
	"wetalk/apps/im/rpc/internal/svc"
	"wetalk/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetChatLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetChatLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetChatLogLogic {
	return &GetChatLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取会话记录
func (l *GetChatLogLogic) GetChatLog(in *im.GetChatLogReq) (*im.GetChatLogResp, error) {
	// 根据id
	if in.MsgId != "" {
		chatlog, err := l.svcCtx.ChatLogModel.FindOne(l.ctx, in.MsgId)
		if err != nil {
			return nil, errors.Wrapf(xerr.NewDBErr(), "find chatLog by msgId err %v, req %v", err, in.MsgId)
		}

		return &im.GetChatLogResp{
			List: []*im.ChatLog{{
				Id:             chatlog.ID.Hex(),
				ConversationId: chatlog.ConversationId,
				SendId:         chatlog.SendId,
				RecvId:         chatlog.RecvId,
				MsgType:        int32(chatlog.MsgType),
				MsgContent:     chatlog.MsgContent,
				ChatType:       int32(chatlog.ChatType),
				SendTime:       chatlog.SendTime,
				ReadRecords:    chatlog.ReadRecords,
			}},
		}, nil
	}
	//时间段分段查询
	data, err := l.svcCtx.ChatLogModel.ListBySendTime(l.ctx, in.ConversationId, in.StartSendTime, in.EndSendTime, in.Count)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "find chatLog list by SendTime err %v, req %v", err, in)
	}

	res := make([]*im.ChatLog, 0, len(data))
	for _, datum := range data {
		res = append(res, &im.ChatLog{
			Id:             datum.ID.Hex(),
			ConversationId: datum.ConversationId,
			SendId:         datum.SendId,
			RecvId:         datum.RecvId,
			MsgType:        int32(datum.MsgType),
			MsgContent:     datum.MsgContent,
			ChatType:       int32(datum.ChatType),
			SendTime:       datum.SendTime,
			ReadRecords:    datum.ReadRecords,
		})
	}

	return &im.GetChatLogResp{
		List: res,
	}, nil
}