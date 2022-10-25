package model

//
// Concentrate
//  @Description: 专注
//
type Concentrate struct {
	Model
	StudentId           uint   `gorm:"column:student_id;not null;comment:学生id" json:"student_id"`
	ConcentrateTargetId uint   `gorm:"column:concentrate_target_id;not null;comment:目标id" json:"concentrate_target_id"`
	ConcentrateTime     string `gorm:"column:concentrate_time;type:varchar(30);not null;default:'';comment:专注时长" json:"concentrate_time"`
	RestTime            string `gorm:"column:rest_time;type:varchar(30);default:'';comment:小休时长" json:"rest_time"`
}
