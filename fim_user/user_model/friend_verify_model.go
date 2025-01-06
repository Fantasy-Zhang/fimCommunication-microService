package user_model

import (
	"fimCommunication-microService/common/model"
	"fimCommunication-microService/common/model/ctype"
)

// FriendVerifyModel 好友验证表
type FriendVerifyModel struct {
	model.Model
	SendUserId           uint                        `json:"send_user_id"`
	RevUserId            uint                        `json:"rev_user_id"`
	Notice               string                      `json:"notice"` // 备注
	Status               int                         `json:"status"` // 状态  1同意 2拒绝 3忽略
	AdditionalMessage    string                      `gorm:"size:128" json:"additional_message"`
	VerificationQuestion *ctype.VerificationQuestion `json:"verification_question"`
}
