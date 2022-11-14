package model

import "server/database"

//
// Chat
//  @Description: 聊天
//
type Chat struct {
	Model
	SendId      uint   `gorm:"column:send_id;not null;comment:发送者id" json:"send_id"`
	RecipientId uint   `gorm:"column:recipient_id;not null;comment:接收者id" json:"recipient_id"`
	ChatContent string `gorm:"column:chat_content;type:varchar(255);default:'';comment:聊天内容" json:"chat_content"`
}

//
// AddChat
//  @Description: 添加聊天
//  @param chat 聊天
//
func AddChat(chat *Chat) {
	db := database.GetMysqlDBInstance()
	db.Create(chat)
}

//
// GetChat
//  @Description: 获取聊天
//  @param sendID 发送者ID
//  @param recipientId 接收者ID
//  @return []Chat
//  @return bool
//
func GetChat(sendID, recipientId uint) ([]Chat, bool) {
	db := database.GetMysqlDBInstance()
	var chatSlice []Chat
	rows, err := db.Model(&Chat{}).Where("(send_id = ? and recipient_id = ?) or (send_id = ? and recipient_id = ?)", sendID, recipientId, recipientId, sendID).Rows()
	if err != nil {
		return nil, false
	}
	for rows.Next() {
		var chat Chat
		err := db.ScanRows(rows, &chat)
		if err != nil {
			return nil, false
		}
		chatSlice = append(chatSlice, chat)
	}
	if len(chatSlice) == 0 {
		return nil, false
	}
	return chatSlice, true
}
