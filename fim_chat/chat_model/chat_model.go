package chat_model

import (
	"fimCommunication-microService/common/model"
	"fimCommunication-microService/common/model/ctype"
)

type ChatModel struct {
	model.Model
	SendUserId uint             `json:"send_user_id"`
	RevUserId  uint             `json:"rev_user_id"`
	MsgType    int              `json:"msg_type"`    //消息类型    1:文本消息 2:图片消息 3:视频消息 4:文件消息 5:语音消息 6:语言通话 7:视频通话 8:撤回消息 9:回复消息 10：引用消息
	MsgPreview string           `json:"msg_preview"` //消息预览
	Msg        ctype.Msg        `json:"msg"`         //消息内容
	SystemMsg  *ctype.SystemMsg `json:"system_msg"`
}

//type SystemMsg struct {
//	Type int `json:"type"` //违规类型 1:涉黄 2:涉恐 3:政治敏感 4:不当言论
//}

//type Msg struct {
//	Content      *string       `json:"content"` //消息类型MsgType=1时使用   和MsgType一样
//	ImageMsg     *ImageMsg     `json:"image_msg"`
//	VideoMsg     *VideoMsg     `json:"video_msg"`
//	FileMsg      *FileMsg      `json:"file_msg"`
//	VoiceMsg     *VoiceMsg     `json:"voice_msg"`
//	VoiceCallMsg *VoiceCallMsg `json:"voice_call_msg"`
//	WithdrawMsg  *WithdrawMsg  `json:"withdraw_msg"`
//	ReplyMsg     *ReplyMsg     `json:"reply_msg"`
//	ReferenceMsg *ReferenceMsg `json:"reference_msg"`
//}
//
//type ImageMsg struct {
//	Title string `json:"title"`
//	Url   string `json:"url"`
//}
//type VideoMsg struct {
//	Title    string `json:"title"`
//	Url      string `json:"url"`
//	Duration int    `json:"duration"`
//}
//type FileMsg struct {
//	Title string `json:"title"`
//	Url   string `json:"url"`
//	Size  int    `json:"size"`
//	Type  string `json:"type"`
//}
//
//// VoiceMsg 语音消息
//type VoiceMsg struct {
//	Title string `json:"title"`
//	Url   string `json:"url"`
//}
//type VoiceCallMsg struct {
//	StartTime time.Time `json:"start_time"` // 通话开始时间
//	EndTime   time.Time `json:"end_time"`   // 通话结束时间
//	EndReason int       `json:"end_reason"` // 通话结束原因 0:发起方挂断 1:接收方挂断 2:网络原因 3:呼叫失败
//}
//type WithdrawMsg struct {
//	Content     string `json:"content"`
//	OriginalMsg Msg    `json:"_"`
//}
//type ReplyMsg struct {
//	MsgId   uint   `json:"msg_id"`
//	Content string `json:"content"` //回复文本
//	Msg     Msg    `json:"msg"`
//}
//type ReferenceMsg struct {
//	MsgId   uint   `json:"msg_id"`
//	Content string `json:"content"` //回复文本
//	Msg     Msg    `json:"msg"`
//}
