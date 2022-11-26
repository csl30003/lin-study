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

func GetStudentIDByDiaryID(diary *Diary) uint {
	db := database.GetMysqlDBInstance()
	if err := db.Where("id = ?", diary.ID).First(&diary).Error; err != nil {
		return 0
	}
	return diary.StudentId
}

func DeleteDiary(diary *Diary) {
	db := database.GetMysqlDBInstance()
	db.Delete(&diary)
}

func GetDiaryByStudentID(studentID uint) (diary []Diary) {
	db := database.GetMysqlDBInstance()
	db.Where("student_id = ?", studentID).Find(&diary)
	return
}

func UpdateDiaryContent(diary *Diary, diaryContent string) {
	db := database.GetMysqlDBInstance()
	db.Model(&diary).Update("diary_content", diaryContent)
}

func LikeDiary(diary *Diary) bool {
	db := database.GetMysqlDBInstance()
	if err := db.Model(&diary).Update("likes", gorm.Expr("likes+?", 1)).Error; err != nil {
		return false
	}
	return true
}

func GetDiaryLikes(diary *Diary) (int32, bool) {
	db := database.GetMysqlDBInstance()
	if err := db.Where("id = ?", diary.ID).First(&diary).Error; err != nil {
		return 0, false
	}
	return diary.Likes, true
}

func UnlikeDiary(diary *Diary) bool {
	db := database.GetMysqlDBInstance()
	if err := db.Model(&diary).Update("likes", gorm.Expr("likes-?", 1)).Error; err != nil {
		return false
	}
	return true
}
