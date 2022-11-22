package service

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"server/model"
	"server/response"
)

//
// GetConcentrateTarget
//  @Description: 获取专注目标
//  @param c 上下文
//
func GetConcentrateTarget(c *gin.Context) {
	claims, _ := c.Get("claims")
	claimsValueElem := reflect.ValueOf(claims).Elem()
	studentID := uint(claimsValueElem.FieldByName("ID").Uint())

	targetSlice := model.GetConcentrateTargetByStudentID(studentID)
	if len(targetSlice) == 0 {
		response.Success(c, "无专注目标", targetSlice)
		return
	}
	response.Success(c, "获取成功", targetSlice)
}

//
// AddConcentrateTarget
//  @Description: 添加专注目标
//  @param c 上下文
//
func AddConcentrateTarget(c *gin.Context) {
	claims, _ := c.Get("claims")
	claimsValueElem := reflect.ValueOf(claims).Elem()
	studentID := uint(claimsValueElem.FieldByName("ID").Uint())

	var concentrateTarget model.ConcentrateTarget

	if err := c.ShouldBindJSON(&concentrateTarget); err != nil {
		response.Failed(c, "参数错误")
		return
	}

	concentrateTarget.StudentId = studentID

	//   查看记录是否存在
	if ok := model.ExistConcentrateTarget(&concentrateTarget); ok {
		response.Failed(c, "专注目标已存在")
		return
	}

	concentrateTargetID := model.AddConcentrateTarget(&concentrateTarget)
	response.Success(c, "添加成功", concentrateTargetID)
}

//
// DeleteConcentrateTarget
//  @Description: 删除专注目标
//  @param c 上下文
//
func DeleteConcentrateTarget(c *gin.Context) {
	claims, _ := c.Get("claims")
	claimsValueElem := reflect.ValueOf(claims).Elem()
	studentID := uint(claimsValueElem.FieldByName("ID").Uint())

	var concentrateTarget model.ConcentrateTarget

	if err := c.ShouldBindJSON(&concentrateTarget); err != nil {
		response.Failed(c, "参数错误")
		return
	}

	concentrateTarget.StudentId = studentID

	id, ok := model.GetConcentrateTargetID(&concentrateTarget)
	if !ok {
		response.Failed(c, "收藏记录不存在")
		return
	}
	concentrateTarget.ID = id

	model.DeleteConcentrateTarget(&concentrateTarget)
	response.Success(c, "删除成功", nil)
}
