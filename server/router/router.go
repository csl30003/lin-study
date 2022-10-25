package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"server/api"
	"time"
)

//
// Start
//  @Description: 路由初始化
//
func Start() {
	e := gin.Default()

	//  解决跨域请求
	mwCORS := cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Type"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 24 * time.Hour,
	})

	e.Use(mwCORS)

	e.POST("/login", api.Login)
	e.POST("/register", api.Register)
	e.GET("/logout", api.Logout)
	e.GET("/getStudentInfo", api.GetStudentInfo)

	e.GET("/student", api.ListStudent)
	e.POST("/student", api.AddStudent)

	err := e.Run()
	if err != nil {
		log.Println("服务器启动失败")
		return
	}
}
