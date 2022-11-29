package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"server/conf"
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

	db *gorm.DB
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

	m.host = conf.Cfg.Section("MYSQL").Key("host").String()
	m.port = conf.Cfg.Section("MYSQL").Key("port").String()
	m.user = conf.Cfg.Section("MYSQL").Key("user").String()
	m.password = conf.Cfg.Section("MYSQL").Key("password").String()
	m.database = conf.Cfg.Section("MYSQL").Key("database").String()

	//  连接MySQL
	dsn := m.user + ":" + m.password + "@tcp(" + m.host + ":" + m.port + ")/" + m.database + "?charset=utf8mb4&parseTime=True&loc=UTC"
	m.db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("MySQL初始化失败", err)
	}

	sqlDB, err := m.db.DB()
	if err != nil {
		log.Fatal("MySQL配置初始化失败", err)
	}

	maxOpenConn, err := conf.Cfg.Section("DATASOURCE").Key("max_open_conn").Int()
	if err != nil {
		log.Fatal("DataSource配置文件类型转换失败失败", err)
	}
	maxIdleConn, err := conf.Cfg.Section("DATASOURCE").Key("max_idle_conn").Int()
	if err != nil {
		log.Fatal("DataSource配置文件类型转换失败失败", err)
	}

	//  设置最大连接数
	sqlDB.SetMaxOpenConns(maxOpenConn)
	//  设置最大空闲连接数
	sqlDB.SetMaxIdleConns(maxIdleConn)
}
