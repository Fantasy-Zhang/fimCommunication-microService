package group_model

import (
	"fimCommunication-microService/common/model"
	"fimCommunication-microService/common/model/ctype"
)

type GroupModel struct {
	model.Model
	Title                string                      `gorm:"size:32"json:"title"`     // 群名称
	Abstract             string                      `gorm:"size:128"json:"abstract"` // 群简介
	Avatar               string                      `gorm:"size:256"json:"avatar"`   // 群头像
	Creator              uint                        `json:"creator"`                 // 群主
	IsSearch             bool                        `json:"is_search"`               // 是否允许搜索
	Verification         int                         `json:"verification"`            //群验证方式  0不允许任何人添加 1允许任何人添加 2需要验证消息 3需要正确回答问题  4需要正确回答问题
	VerificationQuestion *ctype.VerificationQuestion `json:"verification_question"`   //验证问题      3和4的手动验证问题
	ISInvite             bool                        `json:"is_invite"`               // 是否允许邀请
	IsTemporarySession   bool                        `json:"is_temporary_session"`    // 是否允许临时会话
	IsProhibition        bool                        `json:"is_prohibition"`          // 是否禁言
	Size                 int                         `json:"size"`                    // 群大小
}
