package model

import (
	"errors"
	"server/database"
)

//
// Concentrate
//  @Description: 专注
//
type Concentrate struct {
	Model
	StudentId           uint  `gorm:"column:student_id;not null;comment:学生id" json:"student_id"`
	ConcentrateTargetId uint  `gorm:"column:concentrate_target_id;not null;comment:目标id" json:"concentrate_target_id"`
	ConcentrateTime     int32 `gorm:"column:concentrate_time;type:int(4);not null;default:0;comment:专注时长" json:"concentrate_time"`
	Status              int8  `gorm:"column:status;type:tinyint(1);not null;default:0;comment:状态" json:"status"`
}

//
// ExistConcentrateAndStatusIsZero
//  @Description: 通过学生ID获悉是否存在专注记录并且状态为0
//  @param studentID 学生ID
//  @return bool
//
func ExistConcentrateAndStatusIsZero(studentID uint) bool {
	db := database.GetMysqlDBInstance()
	var concentrate Concentrate
	if err := db.Where("student_id = ? and status = 0", studentID).First(&concentrate).Error; err == nil {
		return true
	}
	return false
}

//
// GetConcentrateID
//  @Description: 获取专注ID
//  @param concentrate 专注
//  @return uint
//  @return bool
//
func GetConcentrateID(concentrate *Concentrate) (uint, bool) {
	db := database.GetMysqlDBInstance()
	if err := db.Where("student_id = ? and status = ?", concentrate.StudentId, concentrate.Status).First(&concentrate).Error; err != nil {
		return concentrate.ID, false
	}
	return concentrate.ID, true
}

//
// AddConcentrate
//  @Description: 添加专注
//  @param concentrate 专注
//
func AddConcentrate(concentrate *Concentrate) {
	db := database.GetMysqlDBInstance()
	db.Create(concentrate)
}

//
// DeleteConcentrate
//  @Description: 删除专注
//  @param concentrate
//
func DeleteConcentrate(concentrate *Concentrate) {
	db := database.GetMysqlDBInstance()
	db.Delete(&concentrate)
}

func UpdateConcentrateStatus(studentID uint) (err error) {
	db := database.GetMysqlDBInstance()
	var concentrate Concentrate
	result := db.Model(&concentrate).Where("student_id = ? and status = 0", studentID).Update("status", 1)
	if result.RowsAffected == 0 {
		err = errors.New("update: no updated row")
	}
	if result.Error != nil {
		err = result.Error
	}
	return
}
