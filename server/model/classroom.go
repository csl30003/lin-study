package model

import (
	"errors"
	"gorm.io/gorm"
	"server/database"
)

//
// Classroom
//  @Description: 教室
//
type Classroom struct {
	Model
	Floor  string `gorm:"column:floor;type:varchar(20);default:'';comment:楼;index:seek" json:"floor"`
	Layer  string `gorm:"column:layer;type:varchar(20);default:'';comment:层;index:seek" json:"layer"`
	Class  string `gorm:"column:class;type:varchar(20);default:'';comment:班;index:seek" json:"class"`
	Seat1  string `gorm:"column:seat1;type:varchar(30);not null;default:'_';comment:座位1" json:"seat1"`
	Seat2  string `gorm:"column:seat2;type:varchar(30);not null;default:'_';comment:座位2" json:"seat2"`
	Seat3  string `gorm:"column:seat3;type:varchar(30);not null;default:'_';comment:座位3" json:"seat3"`
	Seat4  string `gorm:"column:seat4;type:varchar(30);not null;default:'_';comment:座位4" json:"seat4"`
	Seat5  string `gorm:"column:seat5;type:varchar(30);not null;default:'_';comment:座位5" json:"seat5"`
	Seat6  string `gorm:"column:seat6;type:varchar(30);not null;default:'_';comment:座位6" json:"seat6"`
	Seat7  string `gorm:"column:seat7;type:varchar(30);not null;default:'_';comment:座位7" json:"seat7"`
	Seat8  string `gorm:"column:seat8;type:varchar(30);not null;default:'_';comment:座位8" json:"seat8"`
	Seat9  string `gorm:"column:seat9;type:varchar(30);not null;default:'_';comment:座位9" json:"seat9"`
	Seat10 string `gorm:"column:seat10;type:varchar(30);not null;default:'_';comment:座位10" json:"seat10"`
}

//
// GetClassroomNum
//  @Description: 获取教室的数量
//  @return int
//
func GetClassroomNum() uint {
	db := database.GetMysqlDBInstance()
	var (
		classroom []Classroom
		count     int64
	)
	db.Model(&classroom).Count(&count)
	return uint(count)
}

//
// GetLayerByFloor
//  @Description: 通过楼获取层
//  @param floor 楼
//  @return layer 层
//
func GetLayerByFloor(floor string) (layer []string) {
	db := database.GetMysqlDBInstance()
	var classroom []Classroom
	db.Distinct("layer").Where("floor = ?", floor).Find(&classroom)

	for i := range classroom {
		layer = append(layer, classroom[i].Layer)
	}
	return
}

//
// GetClassByFloorAndLayer
//  @Description: 通过楼和层获取班
//  @param floor 楼
//  @param layer 层
//  @return class 班
//
func GetClassByFloorAndLayer(floor, layer string) (class []string) {
	db := database.GetMysqlDBInstance()
	var classroom []Classroom
	db.Select("class").Where("floor = ? and layer = ?", floor, layer).Find(&classroom)

	for i := range classroom {
		class = append(class, classroom[i].Class)
	}
	return
}

//
// GetSeatByFloorAndLayerAndClass
//  @Description: 通过楼和层和班获取座位
//  @param floor 楼
//  @param layer 层
//  @param class 班
//  @return map[string]string
//
func GetSeatByFloorAndLayerAndClass(floor, layer, class string) map[string]string {
	db := database.GetMysqlDBInstance()
	var classroom []Classroom
	db.Where("floor = ? and layer = ? and class = ?", floor, layer, class).First(&classroom)

	seat := make(map[string]string, 10)
	seat["seat1"] = classroom[0].Seat1
	seat["seat2"] = classroom[0].Seat2
	seat["seat3"] = classroom[0].Seat3
	seat["seat4"] = classroom[0].Seat4
	seat["seat5"] = classroom[0].Seat5
	seat["seat6"] = classroom[0].Seat6
	seat["seat7"] = classroom[0].Seat7
	seat["seat8"] = classroom[0].Seat8
	seat["seat9"] = classroom[0].Seat9
	seat["seat10"] = classroom[0].Seat10

	return seat
}

//
// GetClassroomID
//  @Description: 通过楼和层和班获取教室ID
//  @param floor 楼
//  @param layer 层
//  @param class 班
//  @return uint
//
func GetClassroomID(floor, layer, class string) (uint, bool) {
	db := database.GetMysqlDBInstance()
	var classroom Classroom
	if err := db.Where("floor = ? and layer = ? and class = ?", floor, layer, class).First(&classroom).Error; err != nil {
		return classroom.ID, false
	}
	return classroom.ID, true
}

//
// UpdateSeat
//  @Description: 入座
//  @param studentName 学生昵称
//  @param classroomID 教室ID
//  @param seat 座位
//  @return bool
//
func UpdateSeat(classroomID uint, studentName, seat string) bool {
	db := database.GetMysqlDBInstance()
	var classroom Classroom

	// 事务
	err := db.Transaction(func(tx *gorm.DB) error {
		seatStr := "seat" + seat
		whereStr := "id = ? and " + seatStr + " = '_'"
		result := db.Model(&classroom).Where(whereStr, classroomID).Update(seatStr, studentName)
		if result.RowsAffected == 0 {
			return errors.New("update: no updated row")
		}
		if result.Error != nil {
			return result.Error
		}

		if err := UpdateStudentStatus(studentName, 1); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return false
	}
	return true
}

//
// UpdateUnseat
//  @Description: 离座
//  @param studentName 学生昵称
//  @param classroomID 教室ID
//  @param seat 座位
//  @return bool
//
func UpdateUnseat(classroomID uint, studentName, seat string) bool {
	db := database.GetMysqlDBInstance()
	var classroom Classroom

	// 事务
	err := db.Transaction(func(tx *gorm.DB) error {
		seatStr := "seat" + seat
		whereStr := "id = ? and " + seatStr + " = ?"
		result := db.Model(&classroom).Where(whereStr, classroomID, studentName).Update(seatStr, "_")
		if result.RowsAffected == 0 {
			return errors.New("update: no updated row")
		}
		if result.Error != nil {
			return result.Error
		}

		if err := UpdateStudentStatus(studentName, 0); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return false
	}
	return true
}
