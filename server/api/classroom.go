package api

import (
	"github.com/gin-gonic/gin"
	"server/model"
	"server/response"
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

}
