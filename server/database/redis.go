package database

import (
	"github.com/go-redis/redis"
	"log"
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
// Conn
//  @Description: 连接
//  @receiver r Redis
//
func (r Redis) Conn() interface{} {
	var err error

	r.db = redis.NewClient(&redis.Options{
		Addr:     r.host + ":" + r.port,
		Password: r.password,
		DB:       r.database,
	})

	_, err = r.db.Ping().Result()
	if err != nil {
		log.Fatal("Redis初始化失败", err)
	}
	return r.db
}
