package service

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"reflect"
	"server/model"
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
	claims, _ := c.Get("claims")
	claimsValueElem := reflect.ValueOf(claims).Elem()
	studentID := uint(claimsValueElem.FieldByName("ID").Uint())

	var chat model.Chat

	if err := c.ShouldBindJSON(&chat); err != nil {
		response.Failed(c, "参数错误")
		return
	}
	cipherText, err := base64.StdEncoding.DecodeString(chat.ChatContent)
	if err != nil {
		response.Failed(c, "参数错误")
		return
	}

	//  RSA解密 为了实现网络安全大作业加的
	plainText, err := tool.RSADecrypt(cipherText, tool.PrivateKeyStr)
	if err != nil {
		response.Failed(c, "解密失败")
		return
	}

	chat.ChatContent = string(plainText)
	chat.SendId = studentID

	model.AddChat(&chat)
	response.Success(c, "发送成功", nil)
}

//
// GetChat
//  @Description: 获取聊天记录
//  @param c 上下文
//
func GetChat(c *gin.Context) {
	claims, _ := c.Get("claims")
	claimsValueElem := reflect.ValueOf(claims).Elem()
	studentID := uint(claimsValueElem.FieldByName("ID").Uint())

	jsonObj := struct {
		RecipientId uint   `json:"recipient_id"`
		PublicKey   string `json:"public_key"`
	}{}
	if err := c.ShouldBindJSON(&jsonObj); err != nil {
		response.Failed(c, "参数错误")
		return
	}

	chatSlice, ok := model.GetChat(studentID, jsonObj.RecipientId)
	if !ok {
		response.Failed(c, "无法获取聊天记录")
		return
	}

	//  加密 为了实现网络安全大作业加的
	for i := range chatSlice {
		chatContentTemp, err := tool.RSAEncrypt([]byte(chatSlice[i].ChatContent), jsonObj.PublicKey)
		if err != nil {
			response.Failed(c, "无法加密")
			return
		}
		chatSlice[i].ChatContent = base64.StdEncoding.EncodeToString(chatContentTemp)
	}

	response.Success(c, "获取成功", chatSlice)
}
