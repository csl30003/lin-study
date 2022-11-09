package model

//
// Notice
//  @Description: 通知
//
type Notice struct {
	Model
	StudentId     uint   `gorm:"column:student_id;not null;comment:学生id" json:"student_id"`
	NotifierId    uint   `gorm:"column:notifier_id;not null;comment:通知人id" json:"notifier_id"`
	NoticeContent string `gorm:"column:notice_content;type:varchar(255);default:'';comment:通知内容" json:"notice_content"`
}

//   可以用桥接模式，用不同身份（管理员、好友、路人）来通知不同的消息（活动、加油、加好友）
