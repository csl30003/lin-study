package model

type CollectClassroom struct {
	Model
	StudentId   uint `gorm:"column:student_id;not null;comment:学生id;index" json:"student_id"`
	ClassroomId uint `gorm:"column:classroom_id;not null;comment:教室id" json:"classroom_id"`
}
