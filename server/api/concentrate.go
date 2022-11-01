package api

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"server/common"
	"server/database"
	"server/model"
	"server/response"
	"strconv"
	"time"
)

//
// BeginConcentrate
//  @Description: 开始专注
//  @param c 上下文
//
func BeginConcentrate(c *gin.Context) {
	//  get data
	claims, _ := c.Get("claims")
	claimsValueElem := reflect.ValueOf(claims).Elem()
	uintStudentID := uint(claimsValueElem.FieldByName("ID").Uint())
	strStudentID := strconv.Itoa(int(uintStudentID))

	db := database.GetRedisDBInstance()
	val, _ := db.Get(common.RedisKeyStudentID + strStudentID).Result()
	if val != "" {
		//  Redis里有该学生的记录 不能再次进行专注
		response.Failed(c, "无法再次进入专注")
		return
	}

	//  通过学生ID查看是否正在专注 如果存在并且status为0就返回错误
	if ok := model.ExistConcentrateAndStatusIsZero(uintStudentID); ok {
		response.Failed(c, "无法再次进入专注")
		return
	}

	var concentrate model.Concentrate

	if err := c.ShouldBindJSON(&concentrate); err != nil {
		response.Failed(c, "参数错误")
		return
	}
	if concentrate.ConcentrateTime <= 0 || concentrate.ConcentrateTime > 180 {
		//  能设置的专注时间范围为0-3h
		response.Failed(c, "参数错误")
		return
	}

	concentrate.StudentId = uintStudentID

	//  MySQL create concentrate
	model.AddConcentrate(&concentrate)

	//  Redis setex student_id concentrate_time+3h ing
	db.Set(common.RedisKeyStudentID+strStudentID, "ing", time.Duration(concentrate.ConcentrateTime)*time.Minute+3*time.Hour)

	response.Success(c, "成功进入专注状态", nil)
}

//
// QuitConcentrate
//  @Description: 退出专注
//  @param c 上下文
//
func QuitConcentrate(c *gin.Context) {
	//  实现Redis的过期回调（待

	claims, _ := c.Get("claims")
	claimsValueElem := reflect.ValueOf(claims).Elem()
	uintStudentID := uint(claimsValueElem.FieldByName("ID").Uint())
	strStudentID := strconv.Itoa(int(uintStudentID))

	var concentrate model.Concentrate
	concentrate.StudentId = uintStudentID
	concentrate.Status = 0
	id, ok := model.GetConcentrateID(&concentrate)
	if !ok {
		response.Failed(c, "专注记录不存在")
		return
	}
	concentrate.ID = id

	//  MySQL delete concentrate
	model.DeleteConcentrate(&concentrate)

	//  Redis del student_id
	db := database.GetRedisDBInstance()
	db.Del(common.RedisKeyStudentID + strStudentID)

	response.Success(c, "成功退出专注状态", nil)
}

//
// EndConcentrate
//  @Description: 结束专注
//  @param c 上下文
//
func EndConcentrate(c *gin.Context) {
	claims, _ := c.Get("claims")
	claimsValueElem := reflect.ValueOf(claims).Elem()
	uintStudentID := uint(claimsValueElem.FieldByName("ID").Uint())
	strStudentID := strconv.Itoa(int(uintStudentID))

	//  if can get student_id and value == 0 and Redis ttl student_id < 3h then MySQL update and Redis del
	//  else then MySQL delete and Redis del
	db := database.GetRedisDBInstance()
	val, _ := db.Get(common.RedisKeyStudentID + strStudentID).Result()
	ttl := db.TTL(common.RedisKeyStudentID + strStudentID).Val().Minutes()
	if val == "0" && ttl > 0 && ttl < 180 {
		//  顺利结束专注
		if err := model.UpdateConcentrateStatus(uintStudentID); err != nil {
			response.Failed(c, "学生当前没有正在专注")
			return
		}
		db.Del(common.RedisKeyStudentID + strStudentID)
		response.Success(c, "顺利结束专注状态", nil)

		//  累加学生总专注时长（待
	} else {
		//  提前结束 即QuitConcentrate()
		var concentrate model.Concentrate
		concentrate.StudentId = uintStudentID
		concentrate.Status = 0
		id, ok := model.GetConcentrateID(&concentrate)
		if !ok {
			response.Failed(c, "专注记录不存在")
			return
		}
		concentrate.ID = id

		//  MySQL delete concentrate
		model.DeleteConcentrate(&concentrate)

		//  Redis del student_id
		db := database.GetRedisDBInstance()
		db.Del(common.RedisKeyStudentID + strStudentID)

		response.Success(c, "提前结束专注状态", nil)
	}
}

//
// OutConcentrate
//  @Description: 切出专注
//  @param c 上下文
//
func OutConcentrate(c *gin.Context) {
	claims, _ := c.Get("claims")
	claimsValueElem := reflect.ValueOf(claims).Elem()
	uintStudentID := uint(claimsValueElem.FieldByName("ID").Uint())
	strStudentID := strconv.Itoa(int(uintStudentID))

	//  Redis setex student_id 30s old_ttl
	db := database.GetRedisDBInstance()
	ttl := db.TTL(common.RedisKeyStudentID + strStudentID).Val().Seconds()
	if ttl <= 0 {
		response.Failed(c, "已结束专注")
		return
	}
	db.Set(common.RedisKeyStudentID+strStudentID, ttl, 30*time.Second)

	response.Success(c, "成功切出专注状态", nil)
}

//
// BackConcentrate
//  @Description: 返回专注
//  @param c 上下文
//
func BackConcentrate(c *gin.Context) {
	claims, _ := c.Get("claims")
	claimsValueElem := reflect.ValueOf(claims).Elem()
	uintStudentID := uint(claimsValueElem.FieldByName("ID").Uint())
	strStudentID := strconv.Itoa(int(uintStudentID))

	//  Redis setex student_id old_value ing
	db := database.GetRedisDBInstance()
	ttl := db.TTL(common.RedisKeyStudentID + strStudentID).Val().Seconds()
	if ttl <= 0 {
		response.Failed(c, "已结束专注")
		return
	}
	val, _ := db.Get(common.RedisKeyStudentID + strStudentID).Result()
	intVal, err := strconv.Atoi(val)
	if err != nil {
		response.Failed(c, "无需返回专注")
		return
	}
	db.Set(common.RedisKeyStudentID+strStudentID, "ing", time.Duration(intVal)*time.Second)

	response.Success(c, "成功返回专注状态", nil)
}
