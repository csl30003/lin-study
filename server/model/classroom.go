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
	Seat1  uint   `gorm:"column:seat1;not null;default:0;comment:座位1" json:"seat1"`
	Seat2  uint   `gorm:"column:seat2;not null;default:0;comment:座位2" json:"seat2"`
	Seat3  uint   `gorm:"column:seat3;not null;default:0;comment:座位3" json:"seat3"`
	Seat4  uint   `gorm:"column:seat4;not null;default:0;comment:座位4" json:"seat4"`
	Seat5  uint   `gorm:"column:seat5;not null;default:0;comment:座位5" json:"seat5"`
	Seat6  uint   `gorm:"column:seat6;not null;default:0;comment:座位6" json:"seat6"`
	Seat7  uint   `gorm:"column:seat7;not null;default:0;comment:座位7" json:"seat7"`
	Seat8  uint   `gorm:"column:seat8;not null;default:0;comment:座位8" json:"seat8"`
	Seat9  uint   `gorm:"column:seat9;not null;default:0;comment:座位9" json:"seat9"`
	Seat10 uint   `gorm:"column:seat10;not null;default:0;comment:座位10" json:"seat10"`
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
//  @return map[int]uint
//
func GetSeatByFloorAndLayerAndClass(floor, layer, class string) map[int]uint {
	db := database.GetMysqlDBInstance()
	var classroom []Classroom
	db.Where("floor = ? and layer = ? and class = ?", floor, layer, class).First(&classroom)

	seat := make(map[int]uint, 10)
	seat[1] = classroom[0].Seat1
	seat[2] = classroom[0].Seat2
	seat[3] = classroom[0].Seat3
	seat[4] = classroom[0].Seat4
	seat[5] = classroom[0].Seat5
	seat[6] = classroom[0].Seat6
	seat[7] = classroom[0].Seat7
	seat[8] = classroom[0].Seat8
	seat[9] = classroom[0].Seat9
	seat[10] = classroom[0].Seat10

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
//  @param studentID 学生ID
//  @param classroomID 教室ID
//  @param seat 座位
//  @return bool
//
func UpdateSeat(studentID, classroomID uint, seat string) bool {
	db := database.GetMysqlDBInstance()
	var classroom Classroom

	// 事务
	err := db.Transaction(func(tx *gorm.DB) error {
		seatStr := "seat" + seat
		whereStr := "id = ? and " + seatStr + " = 0"
		result := db.Model(&classroom).Where(whereStr, classroomID).Update(seatStr, studentID)
		if result.RowsAffected == 0 {
			return errors.New("update: no updated row")
		}
		if result.Error != nil {
			return result.Error
		}

		if err := UpdateStudentStatus(studentID, 1); err != nil {
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
//  @param studentID 学生ID
//  @param classroomID 教室ID
//  @param seat 座位
//  @return bool
//
func UpdateUnseat(studentID, classroomID uint, seat string) bool {
	db := database.GetMysqlDBInstance()
	var classroom Classroom

	// 事务
	err := db.Transaction(func(tx *gorm.DB) error {
		seatStr := "seat" + seat
		whereStr := "id = ? and " + seatStr + " = ?"
		result := db.Model(&classroom).Where(whereStr, classroomID, studentID).Update(seatStr, 0)
		if result.RowsAffected == 0 {
			return errors.New("update: no updated row")
		}
		if result.Error != nil {
			return result.Error
		}

		if err := UpdateStudentStatus(studentID, 0); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return false
	}
	return true
}
