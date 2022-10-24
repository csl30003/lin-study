package model

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
	AccumulatedConcentrateTime string `gorm:"column:accumulated_concentrate_time;type:varchar(30);not null;default:'0min';comment:累计专注次数" json:"accumulated_concentrate_time"`
	Inspired                   int32  `gorm:"column:inspired;type:int(4);not null;default:0;comment:加油次数" json:"inspired"`
	Inspire                    int32  `gorm:"column:inspire;type:int(4);not null;default:0;comment:余香次数" json:"inspire"`
	Money                      int32  `gorm:"column:money;type:int(4);default:0;comment:金币" json:"money"`
	Status                     int8   `gorm:"column:status;type:tinyint(1);not null;default:0;comment:上线状态" json:"status"`
}

func ExistStudentByNameAndPassword(name string, password string) bool {
	var student Student
	if err := db.Where(&Student{Name: name, Password: password}).First(&student).Error; err != nil {
		return false
	}
	return true
}

func ExistStudentByName(name string) bool {
	var student Student
	if err := db.Where(&Student{Name: name}).First(&student).Error; err != nil {
		return false
	}
	return true
}

func AddStudent(student *Student) {
	db.Create(student)
}

func ListStudent() []Student {
	var students []Student
	db.Find(&students)
	return students
}
