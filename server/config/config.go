package config

import (
	"gopkg.in/ini.v1"
	"log"
	"os"
)

var Cfg *ini.File

func init() {
	var err error
	Cfg, err = ini.Load("./config/dev/config.ini")
	if err != nil {
		log.Println("文件读取失败", err)
		os.Exit(1)
	}
}
