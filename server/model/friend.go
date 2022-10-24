package model

type Friend struct {
	Model
	StudentId uint `gorm:"column:student_id;not null;comment:学生id" json:"student_id"`
	ObjectId  uint `gorm:"column:object_id;not null;comment:对方id" json:"object_id"`
	IsMutual  int8 `gorm:"column:is_mutual;type:tinyint(1);not null;default:0;comment:是否双向好友" json:"is_mutual"`
}
