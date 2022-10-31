package common

import "server/model"

// ClassroomNum 教室的数量
var ClassroomNum uint

//
// init
//  @Description: 初始化常量
//
func init() {
	ClassroomNum = model.GetClassroomNum()
}
