package user_model

import "fimCommunication-microService/common/model"

// FriendModel 好友表
type FriendModel struct {
	model.Model
	SendUserId uint   `json:"send_user_id"`
	RevUserId  uint   `json:"rev_user_id"`
	Notice     string `gorm:"size:128" json:"notice"` // 备注
}
