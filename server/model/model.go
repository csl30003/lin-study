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

//
// Model
//  @Description: 固定属性
//
type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

//  数据库实例
var db *gorm.DB

//
// init
//  @Description: MySQL初始化
//
func init() {
	var err error

	//  读取配置文件
	cfg, err := ini.Load("./config/dev/config.ini")
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

	//  连接MySQL
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to init db", err)
	}

	err = db.AutoMigrate(
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
	if err != nil {
		log.Println("数据库表初始化失败")
		return
	}
}
