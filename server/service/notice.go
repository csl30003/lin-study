package service

import (
	"github.com/gin-gonic/gin"
	"server/model"
	"server/response"
	"strconv"
)

//
// GetNotice
//  @Description: 获取通知
//  @param c 上下文
//
func GetNotice(c *gin.Context) {

}

//
// SendNotice
//  @Description: 发送通知
//  @param c
//
func SendNotice(c *gin.Context) {
	noticeTemp := struct {
		NotifierId     uint   `json:"notifier_id"`
		StudentIdSlice []uint `json:"student_id_slice"`
		NoticeContent  string `json:"notice_content"`
	}{}

	if err := c.ShouldBindJSON(&noticeTemp); err != nil {
		response.Failed(c, "参数错误")
		return
	}

	student, ok := model.GetStudentByStudentID(strconv.Itoa(int(noticeTemp.NotifierId)))
	if !ok {
		response.Failed(c, "参数错误")
		return
	}

	//  先理清通知人和对方的关系

	//  管理员(我)-活动 加油 加好友
	//  好友-加油
	//  路人-加油 加好友
	if len(noticeTemp.StudentIdSlice) > 1 && student.Name == "csl30003" {
		//  管理员发通知
		notifier := NewManagerNotifier(noticeTemp.NotifierId, student.Name, noticeTemp.StudentIdSlice)
		event := NewActivityNotify(notifier)
		err := event.EventNotify(noticeTemp.NoticeContent)
		if err != nil {
			response.Failed(c, "通知失败")
			return
		}
		response.Success(c, "通知成功", nil)
	} else if noticeTemp.NoticeContent == "加油" {

	}

}

//   桥接模式，用不同身份（管理员、好友、路人）来通知不同的消息（活动、加油、加好友）

//
// Notifier
//  @Description: 通知者
//
type Notifier interface {
	Notify(noticeContent string) error
}

//
// ManagerNotifier
//  @Description: 管理员通知者
//
type ManagerNotifier struct {
	notifierID     uint
	notifierName   string
	studentIDSlice []uint
}

func (mn *ManagerNotifier) Notify(noticeContent string) error {
	noticeContent = "管理员-" + mn.notifierName + "-" + noticeContent
	//  存数据库
	for i := 0; i < len(mn.studentIDSlice); i++ {
		notice := model.Notice{
			StudentId:     mn.studentIDSlice[i],
			NotifierId:    mn.notifierID,
			NoticeContent: noticeContent,
		}
		model.AddNotice(&notice)
	}
	return nil
}

func NewManagerNotifier(notifierID uint, notifierName string, studentIDSlice []uint) *ManagerNotifier {
	return &ManagerNotifier{
		notifierID:     notifierID,
		notifierName:   notifierName,
		studentIDSlice: studentIDSlice,
	}
}

//
// FriendNotifier
//  @Description: 好友通知者
//
type FriendNotifier struct {
	notifierID     uint
	notifierName   string
	studentIDSlice []uint
}

func (fn *FriendNotifier) Notify(noticeContent string) error {
	noticeContent = "好友-" + fn.notifierName + "-" + noticeContent
	//  存数据库
	return nil
}

func NewFriendNotifier(notifierID uint, notifierName string, studentIDSlice []uint) *FriendNotifier {
	return &FriendNotifier{
		notifierID:     notifierID,
		notifierName:   notifierName,
		studentIDSlice: studentIDSlice,
	}
}

//
// StrangerNotifier
//  @Description: 路人通知者
//
type StrangerNotifier struct {
	notifierID     uint
	notifierName   string
	studentIDSlice []uint
}

func (sn *StrangerNotifier) Notify(noticeContent string) error {
	noticeContent = "路人-" + sn.notifierName + "-" + noticeContent
	//  存数据库
	return nil
}

func NewStrangerNotifier(notifierID uint, notifierName string, studentIDSlice []uint) *StrangerNotifier {
	return &StrangerNotifier{
		notifierID:     notifierID,
		notifierName:   notifierName,
		studentIDSlice: studentIDSlice,
	}
}

//
// Event
//  @Description: 事件
//
type Event interface {
	EventNotify(noticeContent string) error
}

//
// ActivityNotify
//  @Description: 活动通知
//
type ActivityNotify struct {
	notifier Notifier
}

func (an *ActivityNotify) EventNotify(noticeContent string) error {
	noticeContent = "活动-" + noticeContent
	return an.notifier.Notify(noticeContent)
}

func NewActivityNotify(notifier Notifier) *ActivityNotify {
	return &ActivityNotify{
		notifier: notifier,
	}
}

//
// InspiredNotify
//  @Description: 加油通知
//
type InspiredNotify struct {
	notifier Notifier
}

func (in *InspiredNotify) EventNotify(noticeContent string) error {
	noticeContent = "加油-" + noticeContent
	return in.notifier.Notify(noticeContent)
}

func NewInspiredNotify(notifier Notifier) *InspiredNotify {
	return &InspiredNotify{
		notifier: notifier,
	}
}

//
// MakeFriendNotify
//  @Description: 加好友通知
//
type MakeFriendNotify struct {
	notifier Notifier
}

func (mn *MakeFriendNotify) EventNotify(noticeContent string) error {
	noticeContent = "申请加好友-" + noticeContent
	return mn.notifier.Notify(noticeContent)
}

func NewMakeFriendNotify(notifier Notifier) *MakeFriendNotify {
	return &MakeFriendNotify{
		notifier: notifier,
	}
}
