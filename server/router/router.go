package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"server/api"
	"time"
)

func Start() {
	e := gin.Default()
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
	e.GET("/getStudentInfo", api.GetStudentInfo)

	e.GET("/student", api.ListStudent)
	e.POST("/student", api.AddStudent)

	e.Run()
}
