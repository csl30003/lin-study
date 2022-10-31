package api

import (
	"github.com/gin-gonic/gin"
	"server/response"
)

//
// BeginConcentrate
//  @Description: 开始专注
//  @param c 上下文
//
func BeginConcentrate(c *gin.Context) {
	response.Success(c, "成功进入专注状态", nil)
}

//
// EndConcentrate
//  @Description: 结束专注
//  @param c 上下文
//
func EndConcentrate(c *gin.Context) {
	response.Success(c, "成功离开专注状态", nil)
}
