package model

import (
	"gorm.io/gorm"
	"server/database"
)

//
// Diary
//  @Description: 日记
//
type Diary struct {
	Model
	StudentId    uint   `gorm:"column:student_id;not null;comment:学生id" json:"student_id"`
	DiaryContent string `gorm:"column:diary_content;type:varchar(255);default:'';comment:日记内容" json:"diary_content"`
	Likes        int32  `gorm:"column:likes;type:int(4);not null;default:0;comment:点赞数量" json:"likes"`
}

//
// AddDiary
//  @Description: 添加日记
//  @param diary 日记
//
func AddDiary(diary *Diary) {
	db := database.GetMysqlDBInstance()
	db.Create(diary)
}

//
// GetStudentIDByDiaryID
//  @Description: 通过日记ID获取学生ID
//  @param diary 日记
//  @return uint
//
func GetStudentIDByDiaryID(diary *Diary) uint {
	db := database.GetMysqlDBInstance()
	if err := db.Where("id = ?", diary.ID).First(&diary).Error; err != nil {
		return 0
	}
	return diary.StudentId
}

//
// DeleteDiary
//  @Description: 删除日记
//  @param diary 日记
//
func DeleteDiary(diary *Diary) {
	db := database.GetMysqlDBInstance()
	db.Delete(&diary)
}

//
// GetDiaryByStudentID
//  @Description: 通过学生ID获取日记切片
//  @param studentID 学生ID
//  @return diary 日记切片
//
func GetDiaryByStudentID(studentID uint) (diary []Diary) {
	db := database.GetMysqlDBInstance()
	db.Where("student_id = ?", studentID).Find(&diary)
	return
}

//
// UpdateDiaryContent
//  @Description: 更新日记内容
//  @param diary 日记
//  @param diaryContent 日记内容
//
func UpdateDiaryContent(diary *Diary, diaryContent string) {
	db := database.GetMysqlDBInstance()
	db.Model(&diary).Update("diary_content", diaryContent)
}

//
// LikeDiary
//  @Description: 点赞日记
//  @param diary 日记
//  @return bool
//
func LikeDiary(diary *Diary) bool {
	db := database.GetMysqlDBInstance()
	if err := db.Model(&diary).Update("likes", gorm.Expr("likes+?", 1)).Error; err != nil {
		return false
	}
	return true
}

//
// GetDiaryLikes
//  @Description: 获取日记点赞数量
//  @param diary 日记
//  @return int32
//  @return bool
//
func GetDiaryLikes(diary *Diary) (int32, bool) {
	db := database.GetMysqlDBInstance()
	if err := db.Where("id = ?", diary.ID).First(&diary).Error; err != nil {
		return 0, false
	}
	return diary.Likes, true
}

//
// UnlikeDiary
//  @Description: 取消点赞日记
//  @param diary 日记
//  @return bool
//
func UnlikeDiary(diary *Diary) bool {
	db := database.GetMysqlDBInstance()
	if err := db.Model(&diary).Update("likes", gorm.Expr("likes-?", 1)).Error; err != nil {
		return false
	}
	return true
}
