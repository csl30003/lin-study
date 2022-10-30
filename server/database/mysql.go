package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"server/config"
)

//
// Mysql
//  @Description: MySQL
//
type Mysql struct {
	host     string
	port     string
	user     string
	password string
	database string
	db       *gorm.DB
}

//
// GetDBInstance
//  @Description: 获取数据库实例
//  @receiver m MySQL
//  @return interface{}
//
func (m *Mysql) GetDBInstance() any {
	return m.db
}

//
// Conn
//  @Description: 连接
//  @receiver m MySQL
//
func (m *Mysql) Conn() {
	var err error

	m.host = config.Cfg.Section("MYSQL").Key("host").String()
	m.port = config.Cfg.Section("MYSQL").Key("port").String()
	m.user = config.Cfg.Section("MYSQL").Key("user").String()
	m.password = config.Cfg.Section("MYSQL").Key("password").String()
	m.database = config.Cfg.Section("MYSQL").Key("database").String()

	//  连接MySQL
	dsn := m.user + ":" + m.password + "@tcp(" + m.host + ":" + m.port + ")/" + m.database + "?charset=utf8mb4&parseTime=True&loc=Local"
	m.db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("MySQL初始化失败", err)
	}
}
