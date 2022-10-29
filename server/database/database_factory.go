package database

import (
	"github.com/go-redis/redis"
	"gorm.io/gorm"
	"log"
	"server/config"
)

//
// Database
//  @Description: 数据库接口
//
type Database interface {
	Conn() interface{}
}

//
// Factory
//  @Description: 工厂方法接口
//
type Factory interface {
	CreateDatabase() Database
}

//  MySQL工厂
type mysqlFactory struct{}

//
// CreateDatabase
//  @Description: MySQL工厂创建数据库
//  @receiver m MySQL
//  @return Database
//
func (m mysqlFactory) CreateDatabase() Database {
	host := config.Cfg.Section("MYSQL").Key("host").String()
	port := config.Cfg.Section("MYSQL").Key("port").String()
	user := config.Cfg.Section("MYSQL").Key("user").String()
	password := config.Cfg.Section("MYSQL").Key("password").String()
	database := config.Cfg.Section("MYSQL").Key("database").String()

	return Mysql{host, port, user, password, database, nil}
}

//  Redis工厂
type redisFactory struct{}

//
// CreateDatabase
//  @Description: Redis工厂创建数据库
//  @receiver r Redis
//  @return Database
//
func (r redisFactory) CreateDatabase() Database {
	host := config.Cfg.Section("REDIS").Key("host").String()
	port := config.Cfg.Section("REDIS").Key("port").String()
	password := config.Cfg.Section("REDIS").Key("password").String()
	database, err := config.Cfg.Section("REDIS").Key("database").Int()
	if err != nil {
		log.Fatal("Redis配置文件类型转换失败失败", err)
		return nil
	}

	return Redis{host, port, password, database, nil}
}

//
// NewFactory
//  @Description: 生成工厂
//  @param s 工厂名称
//  @return Factory
//
func NewFactory(s string) Factory {
	switch s {
	case "mysql":
		return mysqlFactory{}
	case "redis":
		return redisFactory{}
	}
	return nil
}

//  数据库实例
var (
	mysqlDB *gorm.DB
	redisDB *redis.Client
)

//
// GetMysqlDBInstance
//  @Description: 获取MySQL数据库实例
//  @return *gorm.DB
//
func GetMysqlDBInstance() *gorm.DB {
	return mysqlDB
}

//
// GetRedisDBInstance
//  @Description: 获取Redis数据库实例
//  @return *redis.Client
//
func GetRedisDBInstance() *redis.Client {
	return redisDB
}

func init() {
	//  初始化MySQL
	mf := NewFactory("mysql")
	md := mf.CreateDatabase()
	mysqlDB = md.Conn().(*gorm.DB)

	////  初始化Redis
	//rf := NewFactory("redis")
	//rd := rf.CreateDatabase()
	//redisDB = rd.Conn().(*redis.Client)
}
