package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"server/middleware"
	"server/service"
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
		AllowOrigins:     []string{"http://localhost:8081/"},
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

	e.POST("/login", service.Login)
	e.POST("/register", service.Register)

	i := e.Group("/index", middleware.JWT())
	{
		i.GET("/logout", service.Logout)
		i.GET("/getStudentInfo", service.GetStudentInfo)

		i.GET("/:floor", service.GetLayer)
		i.GET("/:floor/:layer", service.GetClass)
		i.GET("/:floor/:layer/:class", service.GetSeat)
		i.GET("/:floor/:layer/:class/getClassroomID", service.GetClassroomID)

		c := i.Group("/classroom")
		{
			c.PATCH("/seat", service.Seat)
			c.PATCH("/unseat", service.Unseat)

			c.POST("/concentrate", service.BeginConcentrate)
			c.DELETE("/concentrate", service.QuitConcentrate)
			c.PATCH("/concentrate", service.EndConcentrate)
			c.PATCH("/concentrateOut", service.OutConcentrate)
			c.PATCH("/concentrateBack", service.BackConcentrate)
		}

		i.GET("/collectClassroom", service.GetCollectClassroom)
		i.POST("/collectClassroom", service.AddCollectClassroom)
		i.DELETE("/collectClassroom", service.DeleteCollectClassroom)

		i.GET("/concentrateTarget", service.GetConcentrateTarget)
		i.POST("/concentrateTarget", service.AddConcentrateTarget)
		i.DELETE("/concentrateTarget", service.DeleteConcentrateTarget)

		i.GET("/mutualFriend", service.GetMutualFriend)
		i.GET("/myFriend", service.GetMyFriend)
		i.GET("/noMyFriend", service.GetNoMyFriend)
		i.POST("/friend", service.Follow)
		i.DELETE("/friend", service.Unfollow)

		i.GET("/getPublicKey", service.GetPublicKey)
		i.POST("/chat", service.SendMessage)
		i.GET("/chat", service.GetChat)
	}

	err := e.Run(":8080")
	if err != nil {
		log.Println("服务器启动失败")
		return
	}
}
