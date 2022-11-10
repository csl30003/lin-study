package service

import (
	"github.com/gin-gonic/gin"
	"server/response"
	"server/tool"
)

//
// GetPublicKey
//  @Description: 获取公钥
//  @param c 上下文
//
func GetPublicKey(c *gin.Context) {
	response.Success(c, "获取公钥成功", tool.PublicKeyStr)
}

//
// SendMessage
//  @Description: 发送聊天信息
//  @param c 上下文
//
func SendMessage(c *gin.Context) {

}

//
// GetChat
//  @Description: 获取聊天记录
//  @param c 上下文
//
func GetChat(c *gin.Context) {

}
