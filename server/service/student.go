package service

import (
	"github.com/gin-gonic/gin"
	"server/model"
	"server/response"
)

//
// SearchStudent
//  @Description: 搜索学生
//  @param c 上下文
//
func SearchStudent(c *gin.Context) {
	searchContext, ok := c.GetQuery("searchContext")
	if !ok {
		response.Failed(c, "查询条件错误")
		return
	}
	student, ok := model.GetStudentByName(searchContext)
	if !ok {
		response.Failed(c, "查询失败")
		return
	}
	response.Success(c, "搜索成功", student)
}

//
// GetStudent
//  @Description: 获取学生
//  @param c
//
func GetStudent(c *gin.Context) {
	studentID, ok := c.GetQuery("id")
	if !ok {
		response.Failed(c, "查询条件错误")
		return
	}
	student, ok := model.GetStudentByStudentID(studentID)
	if !ok {
		response.Failed(c, "查询失败")
		return
	}
	response.Success(c, "搜索成功", student)
}
