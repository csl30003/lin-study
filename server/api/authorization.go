package api

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/config"
	"server/middleware"
	"server/model"
	"server/response"
	"time"
)

//
// Login
//  @Description: 登录
//  @param c 上下文
//
func Login(c *gin.Context) {
	var student model.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		response.Failed(c, "参数错误")
		return
	}
	//  校验用户名和密码
	if ok := model.ExistStudentByNameAndPassword(student.Name, student.Password); !ok {
		response.Failed(c, "登录失败")
		return
	}

	//  过期时间
	expirationTime := time.Now().Add(12 * time.Hour)
	//  创建JWT声明
	claims := &middleware.Claims{
		ID:   student.ID,
		Name: student.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	//  使用用于签名的算法和令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtKey := []byte(config.Cfg.Section("JWT").Key("secret_key").String())
	//  创建JWT字符串
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		fmt.Println(err)
		response.Failed(c, "内部服务器错误")
		return
	}
	//  将客户端cookie token设置成JWT
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	response.Success(c, "登录成功", nil)
}

//
// Register
//  @Description: 注册
//  @param c 上下文
//
func Register(c *gin.Context) {
	var student model.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		response.Failed(c, "参数错误")
		return
	}

	if ok := model.ExistStudentByName(student.Name); ok {
		response.Failed(c, "用户名已存在")
		return
	}

	model.AddStudent(&student)
	response.Success(c, "注册成功", nil)
}

//
// Logout
//  @Description: 退出登录
//  @param c 上下文
//
func Logout(c *gin.Context) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:   "token",
		Value:  "",
		MaxAge: -1,
		Path:   "/",
	})

	response.Success(c, "退出登录成功", nil)
}

//
// GetStudentInfo
//  @Description: 获取用户信息
//  @param c 上下文
//
func GetStudentInfo(c *gin.Context) {
	claims, _ := c.Get("claims")

	response.Success(c, "获取学生信息成功", claims)
}
