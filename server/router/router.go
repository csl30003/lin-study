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
		i.GET("/:floor/:layer/:class/getClassroomID", api.GetClassroomID)

		c := i.Group("/classroom")
		{
			c.PATCH("/seat", api.Seat)
			c.PATCH("/unseat", api.Unseat)

			c.POST("/concentrate", api.BeginConcentrate)
			c.DELETE("/concentrate", api.QuitConcentrate)
			c.PATCH("/concentrate", api.EndConcentrate)
			c.PATCH("/concentrateOut", api.OutConcentrate)
			c.PATCH("/concentrateBack", api.BackConcentrate)
		}

		i.GET("/collectClassroom", api.GetCollectClassroom)
		i.POST("/collectClassroom", api.AddCollectClassroom)
		i.DELETE("/collectClassroom", api.DeleteCollectClassroom)

		i.GET("/concentrateTarget", api.GetConcentrateTarget)
		i.POST("/concentrateTarget", api.AddConcentrateTarget)
		i.DELETE("/concentrateTarget", api.DeleteConcentrateTarget)
	}

	err := e.Run(":8080")
	if err != nil {
		log.Println("服务器启动失败")
		return
	}
}
