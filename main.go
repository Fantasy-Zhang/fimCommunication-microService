package main

import (
	"fimCommunication-microService/core"
	"fimCommunication-microService/fim_chat/chat_model"
	"fimCommunication-microService/fim_group/group_model"
	"fimCommunication-microService/fim_user/user_model"
	"flag"
)

type options struct {
	DB bool
}

func main() {
	var opt options
	flag.BoolVar(&opt.DB, "db", false, "db")
	flag.Parse()

	if opt.DB {
		//db := core.InitGorm()
		db := core.InitGorm("root:123456@tcp(127.0.0.1:3306)/fim_server_db?charset=utf8mb4&parseTime=True&loc=Local")
		err := db.AutoMigrate(
			&user_model.UserModel{},
			&user_model.UserConfModel{},
			&user_model.FriendModel{},
			&user_model.FriendVerifyModel{},
			&chat_model.ChatModel{},

			&group_model.GroupModel{},
			&group_model.GroupMemberModel{},
			&group_model.GroupVerifyModel{},
			&group_model.GroupMsgModel{},
		)
		if err != nil {
			panic(err)
		}
	}
}
