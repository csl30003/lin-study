package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//
// Response
//  @Description: 返回的JSON
//
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//
// Success
//  @Description: 成功返回
//  @param c 上下文
//  @param message 提示信息
//  @param data 数据
//
func Success(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{200, message, data})
}

//
// Failed
//  @Description: 失败返回
//  @param c 上下文
//  @param message 提示信息
//
func Failed(c *gin.Context, message string) {
	c.JSON(http.StatusOK, Response{400, message, nil})
}
