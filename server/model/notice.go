package model

import "server/database"

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

func AddNotice(notice *Notice) {
	db := database.GetMysqlDBInstance()
	db.Create(notice)
}
