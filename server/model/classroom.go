package model

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
//  @return seat 座位
//
func GetSeatByFloorAndLayerAndClass(floor, layer, class string) (seat map[int]int) {
	return
}
