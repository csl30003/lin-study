package api

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"server/model"
	"server/response"
	"strconv"
)

//
// GetLayer
//  @Description: 获取层
//  @param c 上下文
//
func GetLayer(c *gin.Context) {
	floor := c.Param("floor")
	layer := model.GetLayerByFloor(floor)
	if len(layer) == 0 {
		response.Failed(c, "获取失败")
		return
	}
	response.Success(c, "获取成功", layer)
}

//
// GetClass
//  @Description: 获取班
//  @param c 上下文
//
func GetClass(c *gin.Context) {
	floor, layer := c.Param("floor"), c.Param("layer")
	class := model.GetClassByFloorAndLayer(floor, layer)
	if len(class) == 0 {
		response.Failed(c, "获取失败")
		return
	}
	response.Success(c, "获取成功", class)
}

//
// GetSeat
//  @Description: 获取座位
//  @param c 上下文
//
func GetSeat(c *gin.Context) {
	floor, layer, class := c.Param("floor"), c.Param("layer"), c.Param("class")
	seat := model.GetSeatByFloorAndLayerAndClass(floor, layer, class)
	if len(seat) == 0 {
		response.Failed(c, "获取失败")
		return
	}
	response.Success(c, "获取成功", seat)
}

//
// GetClassroomID
//  @Description: 获取教室ID
//  @param c 上下文
//
func GetClassroomID(c *gin.Context) {
	floor, layer, class := c.Param("floor"), c.Param("layer"), c.Param("class")
	classroomID, ok := model.GetClassroomID(floor, layer, class)
	if !ok {
		response.Failed(c, "获取失败")
		return
	}
	response.Success(c, "获取成功", classroomID)
}

//
// Seat
//  @Description: 入座
//  @param c 上下文
//
func Seat(c *gin.Context) {
	// 反射获取student_id
	claims, _ := c.Get("claims")
	claimsValueElem := reflect.ValueOf(claims).Elem()
	studentID := uint(claimsValueElem.FieldByName("ID").Uint())

	// 查看学生是否在专注
	if model.GetStudentStatus(studentID) == 1 {
		response.Failed(c, "学生正在专注")
		return
	}

	classroomID, ok := c.GetPostForm("classroom_id")
	if !ok {
		response.Failed(c, "获取教室失败")
		return
	}
	intClassroomID, err := strconv.Atoi(classroomID)
	if err != nil {
		response.Failed(c, "获取教室失败")
		return
	}

	seat, ok := c.GetPostForm("seat")
	if !ok {
		response.Failed(c, "获取座位失败")
		return
	}

	if ok := model.UpdateSeat(studentID, uint(intClassroomID), seat); !ok {
		response.Failed(c, "更新失败")
		return
	}
	response.Success(c, "更新成功", nil)
}

//
// Unseat
//  @Description: 离座
//  @param c 上下文
//
func Unseat(c *gin.Context) {
	claims, _ := c.Get("claims")
	claimsValueElem := reflect.ValueOf(claims).Elem()
	studentID := uint(claimsValueElem.FieldByName("ID").Uint())

	// 查看学生是否在专注
	if model.GetStudentStatus(studentID) == 0 {
		response.Failed(c, "学生不在座位上")
		return
	}

	classroomID, ok := c.GetPostForm("classroom_id")
	if !ok {
		response.Failed(c, "获取教室失败")
		return
	}
	intClassroomID, err := strconv.Atoi(classroomID)
	if err != nil {
		response.Failed(c, "获取教室失败")
		return
	}

	seat, ok := c.GetPostForm("seat")
	if !ok {
		response.Failed(c, "获取座位失败")
		return
	}

	if ok := model.UpdateUnseat(studentID, uint(intClassroomID), seat); !ok {
		response.Failed(c, "更新失败")
		return
	}
	response.Success(c, "更新成功", nil)
}
