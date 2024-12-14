// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2
// Source: im.proto

package imclient

import (
	"context"

	"wetalk/apps/im/rpc/im"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ChatLog                     = im.ChatLog
	Conversation                = im.Conversation
	CreateGroupConversationReq  = im.CreateGroupConversationReq
	CreateGroupConversationResp = im.CreateGroupConversationResp
	GetChatLogReq               = im.GetChatLogReq
	GetChatLogResp              = im.GetChatLogResp
	GetConversationsReq         = im.GetConversationsReq
	GetConversationsResp        = im.GetConversationsResp
	PutConversationsReq         = im.PutConversationsReq
	PutConversationsResp        = im.PutConversationsResp
	SetUpUserConversationReq    = im.SetUpUserConversationReq
	SetUpUserConversationResp   = im.SetUpUserConversationResp

	Im interface {
		// 获取会话记录
		GetChatLog(ctx context.Context, in *GetChatLogReq, opts ...grpc.CallOption) (*GetChatLogResp, error)
		// 建立会话: 群聊, 私聊
		SetUpUserConversation(ctx context.Context, in *SetUpUserConversationReq, opts ...grpc.CallOption) (*SetUpUserConversationResp, error)
		// 获取会话
		GetConversations(ctx context.Context, in *GetConversationsReq, opts ...grpc.CallOption) (*GetConversationsResp, error)
		// 更新会话
		PutConversations(ctx context.Context, in *PutConversationsReq, opts ...grpc.CallOption) (*PutConversationsResp, error)
		CreateGroupConversation(ctx context.Context, in *CreateGroupConversationReq, opts ...grpc.CallOption) (*CreateGroupConversationResp, error)
	}

	defaultIm struct {
		cli zrpc.Client
	}
)

func NewIm(cli zrpc.Client) Im {
	return &defaultIm{
		cli: cli,
	}
}

// 获取会话记录
func (m *defaultIm) GetChatLog(ctx context.Context, in *GetChatLogReq, opts ...grpc.CallOption) (*GetChatLogResp, error) {
	client := im.NewImClient(m.cli.Conn())
	return client.GetChatLog(ctx, in, opts...)
}

// 建立会话: 群聊, 私聊
func (m *defaultIm) SetUpUserConversation(ctx context.Context, in *SetUpUserConversationReq, opts ...grpc.CallOption) (*SetUpUserConversationResp, error) {
	client := im.NewImClient(m.cli.Conn())
	return client.SetUpUserConversation(ctx, in, opts...)
}

// 获取会话
func (m *defaultIm) GetConversations(ctx context.Context, in *GetConversationsReq, opts ...grpc.CallOption) (*GetConversationsResp, error) {
	client := im.NewImClient(m.cli.Conn())
	return client.GetConversations(ctx, in, opts...)
}

// 更新会话
func (m *defaultIm) PutConversations(ctx context.Context, in *PutConversationsReq, opts ...grpc.CallOption) (*PutConversationsResp, error) {
	client := im.NewImClient(m.cli.Conn())
	return client.PutConversations(ctx, in, opts...)
}

func (m *defaultIm) CreateGroupConversation(ctx context.Context, in *CreateGroupConversationReq, opts ...grpc.CallOption) (*CreateGroupConversationResp, error) {
	client := im.NewImClient(m.cli.Conn())
	return client.CreateGroupConversation(ctx, in, opts...)
}
