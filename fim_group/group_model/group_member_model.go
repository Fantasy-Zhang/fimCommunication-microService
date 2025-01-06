package group_model

import "fimCommunication-microService/common/model"

type GroupMemberModel struct {
	model.Model
	GroupId         uint   `json:"group_id"`
	UserId          uint   `json:"user_id"`
	MemberNickname  string `json:"member_nickname"`
	Role            int    `json:"role"` // 0普通成员 1管理员 2群主
	ProhibitionTime *int   `json:"prohibition_time"`
}
