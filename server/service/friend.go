package service

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"server/model"
	"server/response"
)

//
// GetMutualFriend
//  @Description: 获取相互的朋友
//  @param c 上下文
//
func GetMutualFriend(c *gin.Context) {
	claims, _ := c.Get("claims")
	claimsValueElem := reflect.ValueOf(claims).Elem()
	studentID := uint(claimsValueElem.FieldByName("ID").Uint())

	friend, ok := model.GetFriendByStudentID(studentID, 1)
	if !ok {
		response.Success(c, "没朋友", friend)
		return
	}
	response.Success(c, "获取成功", friend)
}

//
// GetMyFriend
//  @Description: 获取我想交的朋友
//  @param c 上下文
//
func GetMyFriend(c *gin.Context) {
	claims, _ := c.Get("claims")
	claimsValueElem := reflect.ValueOf(claims).Elem()
	studentID := uint(claimsValueElem.FieldByName("ID").Uint())

	friend, ok := model.GetFriendByStudentID(studentID, 0)
	if !ok {
		response.Success(c, "没朋友", friend)
		return
	}
	response.Success(c, "获取成功", friend)
}

//
// GetNoMyFriend
//  @Description: 获取想交我的朋友
//  @param c 上下文
//
func GetNoMyFriend(c *gin.Context) {
	claims, _ := c.Get("claims")
	claimsValueElem := reflect.ValueOf(claims).Elem()
	studentID := uint(claimsValueElem.FieldByName("ID").Uint())

	friend, ok := model.GetFriendByObjectID(studentID)
	if !ok {
		response.Success(c, "没朋友", friend)
		return
	}
	response.Success(c, "获取成功", friend)
}

//
// Follow
//  @Description: 交朋友
//  @param c 上下文
//
func Follow(c *gin.Context) {
	claims, _ := c.Get("claims")
	claimsValueElem := reflect.ValueOf(claims).Elem()
	studentID := uint(claimsValueElem.FieldByName("ID").Uint())

	var friend model.Friend

	if err := c.ShouldBindJSON(&friend); err != nil {
		response.Failed(c, "参数错误")
		return
	}
	if friend.ObjectId == studentID {
		response.Failed(c, "参数错误")
		return
	}

	friend.StudentId = studentID

	//  判断是否有好友关系 是则返回
	if ok1 := model.ExistFriend(friend.StudentId, friend.ObjectId); ok1 {
		response.Failed(c, "已经申请过添加好友")
		return
	}

	//  否则创建记录 is_mutual=0
	model.AddFriend(&friend)

	//  查看对方是否也follow自己 否则返回
	if ok2 := model.ExistFriend(friend.ObjectId, friend.StudentId); !ok2 {
		response.Success(c, "申请添加好友成功", nil)
		return
	}
	//  是则将这两条记录的is_mutual设为1
	if err := model.UpdateIsMutual(friend.StudentId, friend.ObjectId, 1); err != nil {
		response.Failed(c, "更新好友关系失败")
		return
	}
	if err := model.UpdateIsMutual(friend.ObjectId, friend.StudentId, 1); err != nil {
		response.Failed(c, "更新好友关系失败")
		return
	}
	response.Success(c, "已相互跟随", nil)
}

//
// Unfollow
//  @Description: 绝交
//  @param c 上下文
//
func Unfollow(c *gin.Context) {
	claims, _ := c.Get("claims")
	claimsValueElem := reflect.ValueOf(claims).Elem()
	studentID := uint(claimsValueElem.FieldByName("ID").Uint())

	var friend model.Friend

	if err := c.ShouldBindJSON(&friend); err != nil {
		response.Failed(c, "参数错误")
		return
	}
	if friend.ObjectId == studentID {
		response.Failed(c, "参数错误")
		return
	}

	friend.StudentId = studentID

	//  查看是否存在这条记录 否则返回
	id, ok := model.GetFriendID(&friend)
	if !ok {
		response.Failed(c, "好友记录不存在")
		return
	}

	friend.ID = id

	//  软删除记录
	model.DeleteFriend(&friend)

	//  把对方的记录的is_mutual设为0
	if err := model.UpdateIsMutual(friend.ObjectId, friend.StudentId, 0); err != nil {
		//  更新失败代表对方没有跟随自己
		response.Success(c, "删除好友成功", nil)
		return
	}
	response.Success(c, "删除好友成功", nil)
}
