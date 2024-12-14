package code

import "wetalk/pkg/xcode"

var (
	AlreadyFriend         = xcode.New(200001, "已经是好友关系")
	AlreadyApplied        = xcode.New(200002, "已经申请过")
	FriendReqBeforePass   = xcode.New(200003, "好友申请并已经通过")
	FriendReqBeforeRefuse = xcode.New(200004, "好友申请已经被拒绝")
	NotFriend             = xcode.New(200005, "不是好友")
	CannotJoinAgain       = xcode.New(200006, "已经退群，无法再次加入")
	CannotReapply         = xcode.New(200007, "无法再次申请")
	AlreadyJoin           = xcode.New(200008, "已在群中")
	GroupNotExist         = xcode.New(200009, "该群不存在")
	InvalidInvitation     = xcode.New(200010, "无效的邀请")
	UserNotExistGroup     = xcode.New(200011, "用户不在群中")
	GroupReqBeforePass    = xcode.New(200012, "群申请已经通过")
	GroupReqBeforeRefuse  = xcode.New(200013, "群申请已经被拒绝")
	InvalidApplyId        = xcode.New(200014, "无效的申请ID")
)
