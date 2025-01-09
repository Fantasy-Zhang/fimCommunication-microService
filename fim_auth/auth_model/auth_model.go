package auth_model

import "fimCommunication-microService/common/model"

type UserModel struct {
	model.Model
	Pwd      string `gorm:"size:64"json:"pwd"`
	Nickname string `gorm:"size:32"json:"nickname"`
	Abstract string `gorm:"size:128"json:"abstract"`
	Avatar   string `gorm:"size:256"json:"avatar"`
	Ip       string `gorm:"size:32"json:"ip"`
	Addr     string `gorm:"size:64"json:"addr"`
	Role     int    `json:"role"` //角色 1：管理员 2：普通用户
}
