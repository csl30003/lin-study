package service

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"server/model"
	"server/response"
)

//
// KeepDiary
//  @Description: 写日记
//  @param c 上下文
//
func KeepDiary(c *gin.Context) {
	claims, _ := c.Get("claims")
	claimsValueElem := reflect.ValueOf(claims).Elem()
	studentID := uint(claimsValueElem.FieldByName("ID").Uint())

	var diary model.Diary

	if err := c.ShouldBindJSON(&diary); err != nil {
		response.Failed(c, "参数错误")
		return
	}

	diary.StudentId = studentID
	diary.Likes = 0

	model.AddDiary(&diary)
	response.Success(c, "成功提交日记", diary)
}

//
// EraseTheDiary
//  @Description: 擦掉日记
//  @param c 上下文
//
func EraseTheDiary(c *gin.Context) {
	claims, _ := c.Get("claims")
	claimsValueElem := reflect.ValueOf(claims).Elem()
	studentID := uint(claimsValueElem.FieldByName("ID").Uint())

	var diary model.Diary

	if err := c.ShouldBindJSON(&diary); err != nil {
		response.Failed(c, "参数错误")
		return
	}

	//  通过ID查学生ID 如果不符合则返回
	studentIDTemp := model.GetStudentIDByDiaryID(&diary)
	if studentIDTemp != studentID {
		response.Failed(c, "删除日记失败")
		return
	}

	model.DeleteDiary(&diary)
	response.Success(c, "成功删除日记", nil)
}

//
// ReadDiary
//  @Description: 看日记
//  @param c 上下文
//
func ReadDiary(c *gin.Context) {
	claims, _ := c.Get("claims")
	claimsValueElem := reflect.ValueOf(claims).Elem()
	studentID := uint(claimsValueElem.FieldByName("ID").Uint())

	diarySlice := model.GetDiaryByStudentID(studentID)
	if len(diarySlice) == 0 {
		response.Success(c, "还没写过日记", diarySlice)
		return
	}
	response.Success(c, "获取成功", diarySlice)
}

//
// ChangeDiary
//  @Description: 改日记
//  @param c 上下文
//
func ChangeDiary(c *gin.Context) {
	claims, _ := c.Get("claims")
	claimsValueElem := reflect.ValueOf(claims).Elem()
	studentID := uint(claimsValueElem.FieldByName("ID").Uint())

	var diary model.Diary

	if err := c.ShouldBindJSON(&diary); err != nil {
		response.Failed(c, "参数错误")
		return
	}

	diaryContentTemp := diary.DiaryContent

	//  通过ID查学生ID 如果不符合则返回
	studentIDTemp := model.GetStudentIDByDiaryID(&diary)
	if studentIDTemp != studentID {
		response.Failed(c, "修改日记失败")
		return
	}

	model.UpdateDiaryContent(&diary, diaryContentTemp)
	response.Success(c, "修改日记成功", nil)
}

//
// LikeDiary
//  @Description: 点赞日记
//  @param c 上下文
//
func LikeDiary(c *gin.Context) {
	var diary model.Diary

	if err := c.ShouldBindJSON(&diary); err != nil {
		response.Failed(c, "参数错误")
		return
	}

	if ok := model.LikeDiary(&diary); !ok {
		response.Failed(c, "点赞失败")
		return
	}
	response.Success(c, "成功点赞日记", nil)
}

//
// UnlikeDiary
//  @Description: 取消点赞日记
//  @param c 上下文
//
func UnlikeDiary(c *gin.Context) {
	var diary model.Diary

	if err := c.ShouldBindJSON(&diary); err != nil {
		response.Failed(c, "参数错误")
		return
	}

	//  查看点赞数量是否小于等于0 是则返回
	if likes, ok := model.GetDiaryLikes(&diary); !ok || likes <= 0 {
		response.Failed(c, "取消点赞失败")
		return
	}

	if ok := model.UnlikeDiary(&diary); !ok {
		response.Failed(c, "取消点赞失败")
		return
	}
	response.Success(c, "成功取消点赞日记", nil)
}
