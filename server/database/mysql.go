package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
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
// Conn
//  @Description: 连接
//  @receiver m MySQL
//
func (m Mysql) Conn() interface{} {
	var err error

	//  连接MySQL
	dsn := m.user + ":" + m.password + "@tcp(" + m.host + ":" + m.port + ")/" + m.database + "?charset=utf8mb4&parseTime=True&loc=Local"
	m.db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("MySQL初始化失败", err)
	}
	return m.db
}
