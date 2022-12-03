package model

import (
	"gorm.io/gorm"
	"time"
)

//
// Model
//  @Description: 固定属性
//
type Model struct {
	ID        uint           `gorm:"primary_key" json:"id"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `sql:"index" json:"deleted_at,omitempty"`
}

//
// init
//  @Description: 初始化MySQL数据库表
//
//func init() {
//	db := database.GetMysqlDBInstance()
//	err := db.AutoMigrate(
//		&Student{},
//		&Diary{},
//		&Friend{},
//		&Notice{},
//		&Chat{},
//		&ConcentrateTarget{},
//		&Concentrate{},
//		&Classroom{},
//		&CollectClassroom{},
//	)
//	if err != nil {
//		log.Fatal("MySQL表初始化失败", err)
//		return
//	}
//}
