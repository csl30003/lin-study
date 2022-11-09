package model

import (
	"database/sql"
	"log"
	"server/database"
)

//
// CollectClassroom
//  @Description: 收藏教室
//
type CollectClassroom struct {
	Model
	StudentId   uint `gorm:"column:student_id;not null;comment:学生id;index" json:"student_id"`
	ClassroomId uint `gorm:"column:classroom_id;not null;comment:教室id" json:"classroom_id"`
}

//
// CollectClassroomTemp
//  @Description: 收藏教室的一种展示结构体
//
type CollectClassroomTemp struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

//
// GetCollectClassroomByStudentID
//  @Description: 通过学生ID获取收藏教室
//  @param studentID 学生ID
//  @return []CollectClassroomTemp
//  @return bool
//
func GetCollectClassroomByStudentID(studentID uint) ([]CollectClassroomTemp, bool) {
	db := database.GetMysqlDBInstance()
	var collectClassroomTempSlice []CollectClassroomTemp

	s := `
		select c.id ,concat(c.floor, "-", c.layer, "-", c.class) as name
		from classrooms c 
		where c.deleted_at is null and c.id in 
		(
			select cc.classroom_id  
			from collect_classrooms cc  
			where cc.deleted_at is null and cc.student_id = ?
		)
	`
	rows, err := db.Raw(s, studentID).Rows()
	if err != nil {
		return nil, false
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println("Rows 关闭失败")
		}
	}(rows)

	for rows.Next() {
		var collectClassroomTemp CollectClassroomTemp
		err := db.ScanRows(rows, &collectClassroomTemp)
		if err != nil {
			return nil, false
		}
		collectClassroomTempSlice = append(collectClassroomTempSlice, collectClassroomTemp)
	}
	return collectClassroomTempSlice, true
}

//
// ExistCollectClassroom
//  @Description: 是否存在收藏教室的记录
//  @param studentID 学生ID
//  @param classroomID 教室ID
//  @return bool
//
func ExistCollectClassroom(collectClassroom *CollectClassroom) bool {
	db := database.GetMysqlDBInstance()
	if err := db.Where("student_id = ? and classroom_id = ?", collectClassroom.StudentId, collectClassroom.ClassroomId).First(&collectClassroom).Error; err != nil {
		return false
	}
	return true
}

//
// GetCollectClassroomID
//  @Description: 获取收藏教室ID
//  @param classroom 收藏教室
//  @return uint
//  @return bool
//
func GetCollectClassroomID(collectClassroom *CollectClassroom) (uint, bool) {
	db := database.GetMysqlDBInstance()
	if err := db.Where("student_id = ? and classroom_id = ?", collectClassroom.StudentId, collectClassroom.ClassroomId).First(&collectClassroom).Error; err != nil {
		return collectClassroom.ID, false
	}
	return collectClassroom.ID, true
}

//
// AddCollectClassroom
//  @Description: 添加收藏教室
//  @param classroom 收藏教室
//
func AddCollectClassroom(collectClassroom *CollectClassroom) {
	db := database.GetMysqlDBInstance()
	db.Create(collectClassroom)
}

//
// DeleteCollectClassroom
//  @Description: 删除收藏教室
//  @param classroom 收藏教室
//
func DeleteCollectClassroom(collectClassroom *CollectClassroom) {
	db := database.GetMysqlDBInstance()
	db.Delete(&collectClassroom)
}
