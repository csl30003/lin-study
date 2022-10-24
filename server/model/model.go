package model

import (
	"fmt"
	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

var db *gorm.DB

func init() {
	var err error

	cfg, err := ini.Load("../config/dev/config.ini")
	if err != nil {
		fmt.Println("文件读取失败", err)
		os.Exit(1)
	}
	host := cfg.Section("MYSQL").Key("host").String()
	port := cfg.Section("MYSQL").Key("port").String()
	user := cfg.Section("MYSQL").Key("user").String()
	password := cfg.Section("MYSQL").Key("password").String()
	database := cfg.Section("MYSQL").Key("database").String()

	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to init db", err)
	}
	db.AutoMigrate(
		&Student{},
		&Diary{},
		&Friend{},
		&Notice{},
		&Chat{},
		&ConcentrateTarget{},
		&Concentrate{},
		&Classroom{},
		&CollectClassroom{},
	)
}
