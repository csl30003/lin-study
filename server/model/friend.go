package model

import (
	"errors"
	"server/database"
)

//
// Friend
//  @Description: 朋友
//
type Friend struct {
	Model
	StudentId uint `gorm:"column:student_id;not null;comment:学生id" json:"student_id"`
	ObjectId  uint `gorm:"column:object_id;not null;comment:对方id" json:"object_id"`
	IsMutual  int8 `gorm:"column:is_mutual;type:tinyint(1);not null;default:0;comment:是否双向好友" json:"is_mutual"`
}

//
// FriendTemp
//  @Description: 朋友的一种展示结构体
//
type FriendTemp struct {
	ID                         uint   `json:"id"`
	Name                       string `json:"name"`
	Label                      string `json:"label"`
	Sex                        int8   `json:"sex"`
	Goal                       string `json:"goal"`
	AccumulatedConcentrateTime int32  `json:"accumulated_concentrate_time"`
}

//
// GetFriendByStudentID
//  @Description: 通过学生ID获取朋友
//  @param studentID 学生ID
//  @param mutual 是否相互跟随
//  @return []FriendTemp
//  @return bool
//
func GetFriendByStudentID(studentID, mutual uint) ([]FriendTemp, bool) {
	db := database.GetMysqlDBInstance()
	var friendTempSlice []FriendTemp
	rows, err := db.Model(&Friend{}).Select("friends.object_id as id, students.name, students.label, students.sex, students.goal, students.accumulated_concentrate_time").Joins("left join students on friends.object_id = students.id").Where("friends.student_id = ? and friends.is_mutual = ?", studentID, mutual).Rows()
	if err != nil {
		return nil, false
	}
	for rows.Next() {
		var friendTemp FriendTemp
		err := db.ScanRows(rows, &friendTemp)
		if err != nil {
			return nil, false
		}
		friendTempSlice = append(friendTempSlice, friendTemp)
	}
	if len(friendTempSlice) == 0 {
		return nil, false
	}
	return friendTempSlice, true
}

//
// GetFriendByObjectID
//  @Description: 通过对方ID获取朋友
//  @param objectID 对方ID
//  @return []FriendTemp
//  @return bool
//
func GetFriendByObjectID(objectID uint) ([]FriendTemp, bool) {
	db := database.GetMysqlDBInstance()
	var friendTempSlice []FriendTemp
	rows, err := db.Model(&Friend{}).Select("friends.student_id as id, students.name, students.label, students.sex, students.goal, students.accumulated_concentrate_time").Joins("left join students on friends.student_id = students.id").Where("friends.object_id = ? and friends.is_mutual = 0", objectID).Rows()
	if err != nil {
		return nil, false
	}
	for rows.Next() {
		var friendTemp FriendTemp
		err := db.ScanRows(rows, &friendTemp)
		if err != nil {
			return nil, false
		}
		friendTempSlice = append(friendTempSlice, friendTemp)
	}
	if len(friendTempSlice) == 0 {
		return nil, false
	}
	return friendTempSlice, true
}

//
// ExistFriend
//  @Description: 是否存在好友关系
//  @param studentID 学生ID
//  @param objectID 对方ID
//  @return bool
//
func ExistFriend(studentID, objectID uint) bool {
	db := database.GetMysqlDBInstance()
	var friend Friend
	if err := db.Where("student_id = ? and object_id = ?", studentID, objectID).First(&friend).Error; err != nil {
		return false
	}
	return true
}

//
// AddFriend
//  @Description: 添加朋友
//  @param friend 朋友
//
func AddFriend(friend *Friend) {
	db := database.GetMysqlDBInstance()
	db.Create(friend)
}

//
// UpdateIsMutual
//  @Description: 更新is_mutual值
//  @param studentID 学生ID
//  @param objectID 对方ID
//  @param value 值
//  @return err
//
func UpdateIsMutual(studentID, objectID, value uint) (err error) {
	db := database.GetMysqlDBInstance()
	result := db.Model(Friend{}).Where("student_id = ? and object_id = ?", studentID, objectID).Update("is_mutual", value)
	if result.RowsAffected == 0 {
		err = errors.New("update: no updated row")
	}
	if result.Error != nil {
		err = result.Error
	}
	return
}

//
// GetFriendID
//  @Description: 获取好友ID
//  @param friend 好友
//  @return uint
//  @return bool
//
func GetFriendID(friend *Friend) (uint, bool) {
	db := database.GetMysqlDBInstance()
	if err := db.Where("student_id = ? and object_id = ?", friend.StudentId, friend.ObjectId).First(&friend).Error; err != nil {
		return friend.ID, false
	}
	return friend.ID, true
}

//
// DeleteFriend
//  @Description: 删除朋友
//  @param friend 朋友
//
func DeleteFriend(friend *Friend) {
	db := database.GetMysqlDBInstance()
	db.Delete(&friend)
}
