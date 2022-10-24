package model

type Chat struct {
	Model
	SendId      uint   `gorm:"column:send_id;not null;comment:发送者id" json:"send_id"`
	RecipientId uint   `gorm:"column:recipient_id;not null;comment:接收者id" json:"recipient_id"`
	ChatContent string `gorm:"column:chat_content;type:varchar(255);default:'';comment:聊天内容" json:"chat_content"`
}
