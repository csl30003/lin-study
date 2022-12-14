package model

import (
	"server/database"
)

//
// ConcentrateTarget
//  @Description: 专注目标
//
type ConcentrateTarget struct {
	Model
	StudentId uint   `gorm:"column:student_id;not null;comment:学生id;index" json:"student_id"`
	Target    string `gorm:"column:target;type:varchar(30);not null;default:'';comment:目标" json:"target"`
}

//
// GetConcentrateTargetByStudentID
//  @Description: 通过学生ID获取专注目标
//  @param studentID 学生ID
//  @return []ConcentrateTarget
//
func GetConcentrateTargetByStudentID(studentID uint) (concentrateTarget []ConcentrateTarget) {
	db := database.GetMysqlDBInstance()
	db.Where("student_id = ?", studentID).Find(&concentrateTarget)
	return
}

//
// ExistConcentrateTarget
//  @Description:
//  @param concentrateTarget
//  @return bool
//
func ExistConcentrateTarget(concentrateTarget *ConcentrateTarget) bool {
	db := database.GetMysqlDBInstance()
	if err := db.Where("student_id = ? and target = ?", concentrateTarget.StudentId, concentrateTarget.Target).First(&concentrateTarget).Error; err != nil {
		return false
	}
	return true
}

//
// GetConcentrateTargetID
//  @Description: 获取专注目标ID
//  @param concentrateTarget 专注目标
//  @return uint
//  @return bool
//
func GetConcentrateTargetID(concentrateTarget *ConcentrateTarget) (uint, bool) {
	db := database.GetMysqlDBInstance()
	if err := db.Where("student_id = ? and target = ?", concentrateTarget.StudentId, concentrateTarget.Target).First(&concentrateTarget).Error; err != nil {
		return concentrateTarget.ID, false
	}
	return concentrateTarget.ID, true
}

//
// AddConcentrateTarget
//  @Description: 添加专注目标
//  @param concentrateTarget 专注目标
//  @return uint
//
func AddConcentrateTarget(concentrateTarget *ConcentrateTarget) uint {
	db := database.GetMysqlDBInstance()
	db.Create(concentrateTarget)
	return concentrateTarget.ID
}

//
// DeleteConcentrateTarget
//  @Description: 删除专注目标
//  @param concentrateTarget 专注目标
//
func DeleteConcentrateTarget(concentrateTarget *ConcentrateTarget) {
	db := database.GetMysqlDBInstance()
	db.Delete(&concentrateTarget)
}
