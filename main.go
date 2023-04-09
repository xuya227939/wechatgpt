package main

import (
	"wechatbot/bootstrap"
	"wechatbot/config"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)
	//log.SetLevel(log.InfoLevel)

	log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})

	log.Info("程序启动")
	err := config.LoadConfig()
	if err != nil {
		log.Warn("没有找到配置文件，尝试读取环境变量")
	}

	wechatEnv := config.GetWechat()
	log.Warn(*wechatEnv)
	telegramEnv := config.GetTelegram()
	if wechatEnv != nil {
		bootstrap.StartWebChat()
	} else if telegramEnv != nil && *wechatEnv == "true" {
		bootstrap.StartTelegramBot()
	}

	log.Info("程序退出")
}
