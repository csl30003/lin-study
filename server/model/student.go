package model

import (
	"gorm.io/gorm"
	"server/database"
)

//
// Student
//  @Description: 学生
//
type Student struct {
	Model
	Name                       string `gorm:"column:name;type:varchar(30);not null;default:'未命名';comment:昵称;unique" json:"name"`
	Password                   string `gorm:"column:password;type:varchar(30);not null;comment:密码" json:"password"`
	Label                      string `gorm:"column:label;type:varchar(30);default:'';comment:标签" json:"label"`
	Image                      string `gorm:"column:image;type:mediumtext;comment:头像" json:"image"`
	Sex                        int8   `gorm:"column:sex;type:tinyint(1);not null;default:0;comment:性别" json:"sex"`
	Area                       string `gorm:"column:area;type:varchar(30);default:'';comment:地区" json:"area"`
	Stage                      string `gorm:"column:stage;type:varchar(10);default:'';comment:阶段" json:"stage"`
	School                     string `gorm:"column:school;type:varchar(20);default:'';comment:学校" json:"school"`
	Goal                       string `gorm:"column:goal;type:varchar(30);default:'';comment:备考目标" json:"goal"`
	TotalStudyDay              int32  `gorm:"column:total_study_day;type:int(4);not null;default:0;comment:总自习天数" json:"total_study_day"`
	ContinuousStudyDay         int32  `gorm:"column:continuous_study_day;type:int(4);not null;default:0;comment:当前连续自习天数" json:"continuous_study_day"`
	AccumulatedConcentrateTime int32  `gorm:"column:accumulated_concentrate_time;type:int(4);not null;default:0;comment:累计专注时长" json:"accumulated_concentrate_time"`
	Inspired                   int32  `gorm:"column:inspired;type:int(4);not null;default:0;comment:加油次数" json:"inspired"`
	Inspire                    int32  `gorm:"column:inspire;type:int(4);not null;default:0;comment:余香次数" json:"inspire"`
	Money                      int32  `gorm:"column:money;type:int(4);default:0;comment:金币" json:"money"`
	Status                     int8   `gorm:"column:status;type:tinyint(1);not null;default:0;comment:状态" json:"status"`
}

//
// StudentTemp
//  @Description: 学生的一种展示结构体
//
type StudentTemp struct {
	ID                         uint   `json:"id"`
	Name                       string `json:"name"`
	Label                      string `json:"label"`
	Sex                        int8   `json:"sex"`
	Goal                       string `json:"goal"`
	AccumulatedConcentrateTime int32  `json:"accumulated_concentrate_time"`
}

type StudentDetailsCanBeDisplayed struct {
	ID                         uint   `json:"id"`
	Name                       string `json:"name"`
	Label                      string `json:"label"`
	Sex                        int8   `json:"sex"`
	Area                       string `json:"area"`
	Stage                      string `json:"stage"`
	School                     string `json:"school"`
	Goal                       string `json:"goal"`
	TotalStudyDay              int32  `json:"total_study_day"`
	ContinuousStudyDay         int32  `json:"continuous_study_day"`
	AccumulatedConcentrateTime int32  `json:"accumulated_concentrate_time"`
	Inspired                   int32  `json:"inspired"`
	Inspire                    int32  `json:"inspire"`
	Status                     int8   `json:"status"`
}

// GetStudentByNameAndPassword
//  @Description: 通过昵称和密码获取学生
//  @param name 昵称
//  @param password 密码
//  @return bool
//
func GetStudentByNameAndPassword(name, password string) (Student, bool) {
	db := database.GetMysqlDBInstance()
	var student Student
	if err := db.Where("name = ? and password = ?", name, password).First(&student).Error; err != nil {
		return student, false
	}
	return student, true
}

//
// ExistStudentByName
//  @Description: 通过昵称判断学生是否存在
//  @param name 昵称
//  @return bool
//
func ExistStudentByName(name string) bool {
	db := database.GetMysqlDBInstance()
	var student Student
	if err := db.Where("name = ?", name).First(&student).Error; err != nil {
		return false
	}
	return true
}

//
// AddStudent
//  @Description: 添加学生
//  @param student 学生
//
func AddStudent(student *Student) {
	db := database.GetMysqlDBInstance()
	db.Create(student)
}

//
// GetStudentStatus
//  @Description: 获取学生状态
//  @param id 学生ID
//  @return int8
//
func GetStudentStatus(id uint) int8 {
	db := database.GetMysqlDBInstance()
	var student Student
	db.Where("id = ?", id).First(&student)

	return student.Status
}

//
// UpdateStudentStatus
//  @Description: 更新学生状态
//  @param name 学生昵称
//  @param status 状态
//  @return err
//
func UpdateStudentStatus(name string, status int8) (err error) {
	db := database.GetMysqlDBInstance()
	var student Student
	err = db.Model(&student).Where("name = ?", name).Update("status", status).Error
	return
}

//
// GetStudentByName
//  @Description: 通过昵称获取学生(模糊搜索)
//  @param name 昵称
//  @return []StudentTemp
//  @return bool
//
func GetStudentByName(name string) ([]StudentTemp, bool) {
	db := database.GetMysqlDBInstance()
	var studentTempSlice []StudentTemp
	rows, err := db.Model(&Student{}).Select("students.id, students.name, students.label, students.sex, students.goal, students.accumulated_concentrate_time").Where("students.name like ?", "%"+name+"%").Rows()
	if err != nil {
		return nil, false
	}
	for rows.Next() {
		var studentTemp StudentTemp
		err := db.ScanRows(rows, &studentTemp)
		if err != nil {
			return nil, false
		}
		studentTempSlice = append(studentTempSlice, studentTemp)
	}
	if len(studentTempSlice) == 0 {
		return nil, false
	}
	return studentTempSlice, true
}

//
// GetStudentByStudentID
//  @Description: 通过学生ID获取学生
//  @param id 学生ID
//  @return Student
//  @return bool
//
func GetStudentByStudentID(id string) (StudentDetailsCanBeDisplayed, bool) {
	db := database.GetMysqlDBInstance()
	var studentDetailsCanBeDisplayed StudentDetailsCanBeDisplayed
	if err := db.Model(&Student{}).Where("students.id = ?", id).First(&studentDetailsCanBeDisplayed).Error; err != nil {
		return studentDetailsCanBeDisplayed, false
	}
	return studentDetailsCanBeDisplayed, true
}

//
// UpdateStudentAccumulatedConcentrateTime
//  @Description: 更新学生的总专注时长
//  @param id 学生ID
//  @param concentrateTime 专注时长
//  @return err
//
func UpdateStudentAccumulatedConcentrateTime(id uint, concentrateTime int32) (err error) {
	db := database.GetMysqlDBInstance()
	var student Student
	err = db.Model(&student).Where("id = ?", id).Update("accumulated_concentrate_time", gorm.Expr("accumulated_concentrate_time+?", concentrateTime)).Error
	return
}
