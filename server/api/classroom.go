package api

import (
	"github.com/gin-gonic/gin"
	"reflect"
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
	floor, layer, class := c.Param("floor"), c.Param("layer"), c.Param("class")
	seat := model.GetSeatByFloorAndLayerAndClass(floor, layer, class)
	if len(seat) == 0 {
		response.Failed(c, "获取失败")
		return
	}
	response.Success(c, "获取成功", seat)
}

//
// Seat
//  @Description: 入座
//  @param c 上下文
//
func Seat(c *gin.Context) {
	claims, _ := c.Get("claims")
	// 反射获取id
	claimsValueElem := reflect.ValueOf(claims).Elem()
	id := uint(claimsValueElem.FieldByName("ID").Uint())
	floor, layer, class := c.Param("floor"), c.Param("layer"), c.Param("class")
	seat, ok := c.GetPostForm("seat")
	if !ok {
		response.Failed(c, "获取座位失败")
		return
	}

	if ok := model.UpdateSeat(floor, layer, class, seat, id); !ok {
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
	id := uint(claimsValueElem.FieldByName("ID").Uint())
	floor, layer, class := c.Param("floor"), c.Param("layer"), c.Param("class")
	seat, ok := c.GetPostForm("seat")
	if !ok {
		response.Failed(c, "获取座位失败")
		return
	}

	if ok := model.UpdateUnseat(floor, layer, class, seat, id); !ok {
		response.Failed(c, "更新失败")
		return
	}
	response.Success(c, "更新成功", nil)
}
