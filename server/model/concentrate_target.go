package model

type ConcentrateTarget struct {
	Model
	StudentId uint   `gorm:"column:student_id;not null;comment:学生id;index" json:"student_id"`
	Target    string `gorm:"column:target;type:varchar(30);not null;default:'';comment:目标" json:"target"`
}
