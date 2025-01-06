package group_model

import (
	"fimCommunication-microService/common/model"
	"fimCommunication-microService/common/model/ctype"
)

// GroupVerifyModel 群验证消息
type GroupVerifyModel struct {
	model.Model
	GroupId              int64                       `json:"group_id"`
	UserId               int64                       `json:"user_id"`
	Status               int                         `json:"status"` // 状态  1同意 2拒绝 3忽略
	AdditionalMessage    string                      `json:"additional_message"`
	VerificationQuestion *ctype.VerificationQuestion `json:"verification_question"`
	Type                 int                         `json:"type"` //退群 1加群 2退群
}
