package model

import (
	"errors"
	"gorm.io/gorm"
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
// GetLayerByFloor
//  @Description: 通过楼获取层
//  @param floor 楼
//  @return layer 层
//
func GetLayerByFloor(floor string) (layer []string) {
	var classroom []Classroom
	db.Distinct("layer").Where("floor = ? and deleted_at is null", floor).Find(&classroom)

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
	var classroom []Classroom
	db.Select("class").Where("floor = ? and layer = ? and deleted_at is null", floor, layer).Find(&classroom)

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
	var classroom []Classroom
	db.Where("floor = ? and layer = ? and class = ? and deleted_at is null", floor, layer, class).First(&classroom)

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
//  @Description: 通过楼和层和班获取教室id
//  @param floor 楼
//  @param layer 层
//  @param class 班
//  @return uint
//
func GetClassroomID(floor, layer, class string) uint {
	var classroom Classroom
	db.Where("floor = ? and layer = ? and class = ? and deleted_at is null", floor, layer, class).First(&classroom)

	return classroom.ID
}

//
// UpdateSeat
//  @Description: 更新座位
//  @param floor 楼
//  @param layer 层
//  @param class 班
//  @param seat 座位
//  @param id 学生id
//  @return bool
//
func UpdateSeat(floor, layer, class, seat string, id uint) bool {
	var classroom Classroom

	// 事务
	err := db.Transaction(func(tx *gorm.DB) error {
		// 待修改 先获取

		seatStr := "seat" + seat
		whereStr := "floor = ? and layer = ? and class = ? and " + seatStr + " = 0 and deleted_at is null"
		result := db.Model(&classroom).Where(whereStr, floor, layer, class).Update(seatStr, id)
		if result.RowsAffected == 0 {
			return errors.New("update: no updated row")
		}
		if result.Error != nil {
			return result.Error
		}

		// status为1代表正在座位上 为0代表无状态
		var status int8 = 1
		if id == 0 {
			status = 0
		}
		if err := UpdateStudentStatus(id, status); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return false
	}
	return true
}
