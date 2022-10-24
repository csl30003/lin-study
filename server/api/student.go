package api

import (
	"github.com/gin-gonic/gin"
	"server/model"
	"server/response"
)

func AddStudent(c *gin.Context) {
	var student model.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		response.Failed(c, "参数错误")
		return
	}

	model.AddStudent(&student)
	response.Success(c, "添加成功", student)
}

func ListStudent(c *gin.Context) {
	students := model.ListStudent()
	response.Success(c, "查询成功", students)
}
