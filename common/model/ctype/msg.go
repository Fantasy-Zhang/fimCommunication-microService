package ctype

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type Msg struct {
	Content      *string       `json:"content"` //消息类型MsgType=1时使用   和MsgType一样
	ImageMsg     *ImageMsg     `json:"image_msg"`
	VideoMsg     *VideoMsg     `json:"video_msg"`
	FileMsg      *FileMsg      `json:"file_msg"`
	VoiceMsg     *VoiceMsg     `json:"voice_msg"`
	VoiceCallMsg *VoiceCallMsg `json:"voice_call_msg"`
	WithdrawMsg  *WithdrawMsg  `json:"withdraw_msg"`
	ReplyMsg     *ReplyMsg     `json:"reply_msg"`
	ReferenceMsg *ReferenceMsg `json:"reference_msg"`
	AtMsg        *AtMsg        `json:"at_msg"`
}

func (c *Msg) Scan(val interface{}) error {
	return json.Unmarshal(val.([]byte), c)
}
func (c *Msg) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

type ImageMsg struct {
	Title string `json:"title"`
	Url   string `json:"url"`
}
type VideoMsg struct {
	Title    string `json:"title"`
	Url      string `json:"url"`
	Duration int    `json:"duration"`
}
type FileMsg struct {
	Title string `json:"title"`
	Url   string `json:"url"`
	Size  int    `json:"size"`
	Type  string `json:"type"`
}

// VoiceMsg 语音消息
type VoiceMsg struct {
	Title string `json:"title"`
	Url   string `json:"url"`
}
type VoiceCallMsg struct {
	StartTime time.Time `json:"start_time"` // 通话开始时间
	EndTime   time.Time `json:"end_time"`   // 通话结束时间
	EndReason int       `json:"end_reason"` // 通话结束原因 0:发起方挂断 1:接收方挂断 2:网络原因 3:呼叫失败
}
type WithdrawMsg struct {
	Content     string `json:"content"`
	OriginalMsg Msg    `json:"_"`
}
type ReplyMsg struct {
	MsgId   uint   `json:"msg_id"`
	Content string `json:"content"` //回复文本
	Msg     Msg    `json:"msg"`
}
type ReferenceMsg struct {
	MsgId   uint   `json:"msg_id"`
	Content string `json:"content"` //回复文本
	Msg     Msg    `json:"msg"`
}
type AtMsg struct {
	UserId  uint   `json:"user_id"`
	Content string `json:"content"` //回复文本
	Msg     Msg    `json:"msg"`
}
