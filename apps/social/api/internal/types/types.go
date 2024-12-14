// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2

package types

type FriendDeleteReq struct {
	UserId    string `json:"user_id,omitempty"`    // 用户ID
	FriendUid string `json:"friend_uid,omitempty"` // 需要删除的好友UID
}

type FriendDeleteResp struct {
	Success bool `json:"success,omitempty"` // 删除是否成功
}

type FriendListReq struct {
	UserId string `json:"user_id,omitempty"`
}

type FriendListResp struct {
	List []*Friends `json:"list,omitempty"`
}

type FriendPutInHandleReq struct {
	FriendReqId  int32 `json:"friend_req_id,omitempty"`
	HandleResult int32 `json:"handle_result,omitempty"` // 处理结果
}

type FriendPutInHandleResp struct {
	Success bool `json:"success,omitempty"` // 处理是否成功
}

type FriendPutInListReq struct {
	UserId string `json:"user_id,omitempty"`
}

type FriendPutInListResp struct {
	List []*FriendRequests `json:"list,omitempty"`
}

type FriendPutInReq struct {
	UserId string `json:"user_id,omitempty"`
	ReqMsg string `json:"req_msg,omitempty"`
}

type FriendPutInResp struct {
	Success bool `json:"success,omitempty"`
}

type FriendRequests struct {
	Id           int64  `json:"id,omitempty"`
	UserId       string `json:"user_id,omitempty"`
	ReqUid       string `json:"req_uid,omitempty"`
	ReqMsg       string `json:"req_msg,omitempty"`
	ReqTime      int64  `json:"req_time,omitempty"`
	HandleResult int    `json:"handle_result,omitempty"`
	HandleMsg    string `json:"handle_msg,omitempty"`
}

type Friends struct {
	Id        int32  `json:"id,omitempty"`
	FriendUid string `json:"friend_uid,omitempty"`
	Nickname  string `json:"nickname,omitempty"`
	Avatar    string `json:"avatar,omitempty"`
	Remark    string `json:"remark,omitempty"`
}

type GroupCreateReq struct {
	Name string `json:"name,omitempty"`
	Icon string `json:"icon,omitempty"`
}

type GroupCreateResp struct {
	Id string `json:"id,omitempty"`
}

type GroupExitReq struct {
	UserId  string `json:"user_id,omitempty"`  // 用户ID
	GroupId string `json:"group_id,omitempty"` // 群ID
}

type GroupExitResp struct {
	Success bool `json:"success,omitempty"` // 退出是否成功
}

type GroupListReq struct {
}

type GroupListResp struct {
	List []*Groups `json:"list,omitempty"`
}

type GroupMembers struct {
	Id            int32  `json:"id,omitempty"`
	GroupId       string `json:"group_id,omitempty"`
	UserId        string `json:"user_id,omitempty"`
	Nickname      string `json:"nickname,omitempty"`
	UserAvatarUrl string `json:"user_avatar_url,omitempty"`
	RoleLevel     int32  `json:"role_level,omitempty"`
	JoinTime      int64  `json:"join_time,omitempty"`
	JoinSource    int32  `json:"join_source,omitempty"`
	InviterUid    string `json:"inviter_uid,omitempty"`
	OperatorUid   string `json:"operator_uid,omitempty"`
}

type GroupPutInHandleReq struct {
	GroupReqId   int32  `json:"group_req_id,omitempty"`
	GroupId      string `json:"group_id,omitempty"`
	HandleResult int32  `json:"handle_result,omitempty"` // 处理结果
}

type GroupPutInHandleResp struct {
	GroupId string `json:"group_id,omitempty"`
}

type GroupPutinListReq struct {
	GroupId string `json:"group_id,omitempty"`
}

type GroupPutinListResp struct {
	List []*GroupRequests `json:"list,omitempty"`
}

type GroupPutinReq struct {
	GroupId    string `json:"group_id,omitempty"`
	ReqMsg     string `json:"req_msg,omitempty"`
	JoinSource int32  `json:"join_source,omitempty"`
	InviterUid string `json:"inviter_uid,omitempty"`
}

type GroupPutinResp struct {
	GroupId string `json:"group_id,omitempty"`
}

type GroupRequests struct {
	Id           int32  `json:"id,omitempty"`
	UserId       string `json:"user_id,omitempty"`
	GroupId      string `json:"group_id,omitempty"`
	ReqId        string `json:"req_id,omitempty"`
	ReqMsg       string `json:"req_msg,omitempty"`
	ReqTime      int64  `json:"req_time,omitempty"`
	JoinSource   int32  `json:"join_source,omitempty"`
	InviterUid   string `json:"inviter_uid,omitempty"`
	HandleUid    string `json:"handle_uid,omitempty"`
	HandleResult int32  `json:"handle_result,omitempty"` // 处理结果
}

type GroupUsersReq struct {
	GroupId string `json:"group_id,omitempty"`
}

type GroupUsersResp struct {
	List []*GroupMembers `json:"list,omitempty"`
}

type Groups struct {
	Id              string `json:"id,omitempty"`
	Name            string `json:"name,omitempty"`
	Icon            string `json:"icon,omitempty"`
	Status          int32  `json:"status,omitempty"`
	CreatorUid      string `json:"creator_uid,omitempty"`
	GroupType       int32  `json:"group_type,omitempty"`
	MemberNum       int64  `json:"member_num,omitempty"`
	IsVerify        bool   `json:"is_verify,omitempty"`
	Notification    string `json:"notification,omitempty"` // 公告通知
	NotificationUid string `json:"notification_uid,omitempty"`
}
