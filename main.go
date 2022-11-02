package main

import (
	"fmt"
	"github.com/yrzs/wechat-work/kernel"
	"github.com/yrzs/wechat-work/work"
)

func GetWorkConfig() *work.UserConfig {
	return &work.UserConfig{
		CorpID:  "CorpID",
		AgentID: 1,
		Secret:  "Secret",
		Log: work.Log{
			Level: "debug",
			File:  "./wechat.log",
		},
		OAuth: work.OAuth{
			Callback: "Callback",
			Scopes:   []string{},
		},
		Cache: kernel.NewRedisClient(&kernel.RedisOptions{
			Addr:     "127.0.0.1:6379",
			Password: "",
			DB:       1,
		}),
		HttpDebug: true,
		Debug: true,
		// server config
		AESKey: "AESKey",
	}
}

func main() {
	workConfig := GetWorkConfig()
	w, err := work.NewWork(workConfig)
	if err != nil {
		panic(err.Error())
	}
	var userList []string
	userList = append(userList, "1")
	res, err := w.OA.GetCheckinData(3, 1667232000, 1667318399, userList)
	fmt.Println(err)
	fmt.Println(res.CheckInData)
}