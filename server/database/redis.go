package database

import (
	"github.com/go-redis/redis"
	"log"
	"server/conf"
)

//
// Redis
//  @Description: Redis
//
type Redis struct {
	host     string
	port     string
	password string
	database int
	db       *redis.Client
}

//
// GetDBInstance
//  @Description: 获取数据库实例
//  @receiver r Redis
//  @return interface{}
//
func (r *Redis) GetDBInstance() any {
	return r.db
}

//
// Conn
//  @Description: 连接
//  @receiver r Redis
//
func (r *Redis) Conn() {
	var err error

	r.host = conf.Cfg.Section("REDIS").Key("host").String()
	r.port = conf.Cfg.Section("REDIS").Key("port").String()
	r.password = conf.Cfg.Section("REDIS").Key("password").String()
	r.database, err = conf.Cfg.Section("REDIS").Key("database").Int()
	if err != nil {
		log.Fatal("Redis配置文件类型转换失败失败", err)
	}

	r.db = redis.NewClient(&redis.Options{
		Addr:     r.host + ":" + r.port,
		Password: r.password,
		DB:       r.database,
	})

	_, err = r.db.Ping().Result()
	if err != nil {
		log.Fatal("Redis初始化失败", err)
	}
}
