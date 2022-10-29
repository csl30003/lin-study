package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"server/config"
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
var DB *gorm.DB

//
// GetDBInstance
//  @Description: 单例模式 获取数据库实例
//  @return *gorm.DB
//
func GetDBInstance() *gorm.DB {
	return DB
}

//
// init
//  @Description: MySQL初始化
//
func init() {
	var err error

	host := config.Cfg.Section("MYSQL").Key("host").String()
	port := config.Cfg.Section("MYSQL").Key("port").String()
	user := config.Cfg.Section("MYSQL").Key("user").String()
	password := config.Cfg.Section("MYSQL").Key("password").String()
	database := config.Cfg.Section("MYSQL").Key("database").String()

	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"

	//  连接MySQL
	db := GetDBInstance()
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
