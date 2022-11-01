package common

import "server/model"

var (
	// ClassroomNum 教室的数量
	ClassroomNum uint
	// RedisKeyStudentID Redis的key命名前缀student_id
	RedisKeyStudentID = "lin_study:student_id:"
)

//
// init
//  @Description: 初始化常量
//
func init() {
	ClassroomNum = model.GetClassroomNum()
}
