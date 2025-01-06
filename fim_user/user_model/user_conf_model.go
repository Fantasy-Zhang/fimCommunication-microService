package user_model

import (
	"fimCommunication-microService/common/model"
	"fimCommunication-microService/common/model/ctype"
)

// UserConfModel 用户配置表
type UserConfModel struct {
	model.Model
	UserId               uint                        `json:"user_id"`
	RecallMessage        *string                     `gorm:"size:32"json:"recall_message"` //消息撤回
	FriendOnline         bool                        `json:"friend_online"`                //好友上线
	Sound                bool                        `json:"sound"`                        //声音
	SecureLink           bool                        `json:"secure_link"`                  //安全链接
	SavePwd              bool                        `json:"save_pwd"`                     //记住密码
	SearchUser           int                         `json:"search_user"`                  //用户查找方式  0不允许别人查找 1只能通过用户号查找 2允许别人通过昵称查找
	Verification         int                         `json:"verification"`                 //好友验证方式  0不允许任何人添加 1允许任何人添加 2需要验证消息 3需要正确回答问题  4需要正确回答问题
	VerificationQuestion *ctype.VerificationQuestion `json:"verification_question"`        //验证问题      3和4的手动验证问题
	Online               bool                        `json:"online"`                       //在线状态   true 在线  false离线
}

//type VerificationQuestion struct {
//	Problem1 string `json:"problem1"`
//	Problem2 string `json:"problem2"`
//	Problem3 string `json:"problem3"`
//	Answer1  string `json:"answer1"`
//	Answer2  string `json:"answer2"`
//	Answer3  string `json:"answer3"`
//}
