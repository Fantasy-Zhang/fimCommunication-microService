package group_model

import (
	"fimCommunication-microService/common/model"
	"fimCommunication-microService/common/model/ctype"
)

// GroupMsgModel 群聊消息
type GroupMsgModel struct {
	model.Model
	SendUserId uint             `json:"send_user_id"`
	RevUserId  uint             `json:"rev_user_id"`
	MsgType    int              `json:"msg_type"`    //消息类型    1:文本消息 2:图片消息 3:视频消息 4:文件消息 5:语音消息 6:语言通话 7:视频通话 8:撤回消息 9:回复消息 10：引用消息  11:@消息
	MsgPreview string           `json:"msg_preview"` //消息预览
	Msg        ctype.Msg        `json:"msg"`         //消息内容
	SystemMsg  *ctype.SystemMsg `json:"system_msg"`
}
