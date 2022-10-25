package model

//
// Diary
//  @Description: 日记
//
type Diary struct {
	Model
	StudentId    uint   `gorm:"column:student_id;not null;comment:学生id" json:"student_id"`
	ClassroomId  uint   `gorm:"column:classroom_id;not null;comment:教师id" json:"classroom_id"`
	DiaryContent string `gorm:"column:diary_content;type:varchar(255);default:'';comment:日记内容" json:"diary_content"`
	Like         int32  `gorm:"column:like;type:int(4);not null;default:0;comment:点赞数量" json:"like"`
}
