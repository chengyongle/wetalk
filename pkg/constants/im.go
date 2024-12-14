package constants

type MType int

const (
	TextMType MType = iota
)

type ChatType int

const (
	GroupChatType  ChatType = iota + 1 //群聊
	SingleChatType                     //私聊
)

type ContentType int

const (
	ContentChatMsg ContentType = iota
	ContentMakeRead
)
