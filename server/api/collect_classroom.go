package api

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"server/model"
	"server/response"
)

var classroomNum uint

//
// init
//  @Description: 初始化计算教室的数量
//
func init() {
	classroomNum = model.GetClassroomNum()
}

//
// GetCollectClassroom
//  @Description: 获取收藏教室
//  @param c 上下文
//
func GetCollectClassroom(c *gin.Context) {
	claims, _ := c.Get("claims")
	claimsValueElem := reflect.ValueOf(claims).Elem()
	studentID := uint(claimsValueElem.FieldByName("ID").Uint())

	collectClassroom := model.GetCollectClassroomByStudentID(studentID)
	//   判断是否有数据 没有就返回 response.Success(c, "无收藏教室", collectClassroom)
	response.Success(c, "获取成功", collectClassroom)
}

//
// AddCollectClassroom
//  @Description: 添加收藏教室
//  @param c 上下文
//
func AddCollectClassroom(c *gin.Context) {
	claims, _ := c.Get("claims")
	claimsValueElem := reflect.ValueOf(claims).Elem()
	studentID := uint(claimsValueElem.FieldByName("ID").Uint())

	var collectClassroom model.CollectClassroom

	if err := c.ShouldBindJSON(&collectClassroom); err != nil {
		response.Failed(c, "参数错误")
		return
	}
	if collectClassroom.ClassroomId <= 0 || collectClassroom.ClassroomId > classroomNum {
		response.Failed(c, "参数错误")
		return
	}

	collectClassroom.StudentId = studentID

	//   查看记录是否存在
	if ok := model.ExistCollectClassroom(&collectClassroom); ok {
		response.Failed(c, "收藏记录已存在")
		return
	}

	model.AddCollectClassroom(&collectClassroom)
	response.Success(c, "添加成功", nil)
}

//
// CancelCollectClassroom
//  @Description: 取消收藏教室
//  @param c 上下文
//
func CancelCollectClassroom(c *gin.Context) {
	claims, _ := c.Get("claims")
	claimsValueElem := reflect.ValueOf(claims).Elem()
	studentID := uint(claimsValueElem.FieldByName("ID").Uint())

	var collectClassroom model.CollectClassroom

	if err := c.ShouldBindJSON(&collectClassroom); err != nil {
		response.Failed(c, "参数错误")
		return
	}
	if collectClassroom.ClassroomId <= 0 || collectClassroom.ClassroomId > classroomNum {
		response.Failed(c, "参数错误")
		return
	}

	collectClassroom.StudentId = studentID

	id, ok := model.GetCollectClassroomID(&collectClassroom)
	if !ok {
		response.Failed(c, "收藏记录不存在")
		return
	}
	collectClassroom.ID = id

	model.DeleteCollectClassroom(&collectClassroom)
	response.Success(c, "删除成功", nil)
}
