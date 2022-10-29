package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"server/api"
	"server/middleware"
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

	i := e.Group("/index", middleware.JWT())
	{
		i.GET("/logout", api.Logout)
		i.GET("/getStudentInfo", api.GetStudentInfo)

		i.GET("/:floor", api.GetLayer)
		i.GET("/:floor/:layer", api.GetClass)
		i.GET("/:floor/:layer/:class", api.GetSeat)
		i.PATCH("/:floor/:layer/:class/seat", api.Seat)
		i.PATCH("/:floor/:layer/:class/unseat", api.Unseat)

		i.POST("")
	}

	err := e.Run(":8080")
	if err != nil {
		log.Println("服务器启动失败")
		return
	}
}
