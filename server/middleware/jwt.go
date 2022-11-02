package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/conf"
	"server/response"
)

//
// Claims
//  @Description: JWT返回的JSON
//
type Claims struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	jwt.StandardClaims
}

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		//  获取Cookie
		ck, err := c.Request.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				// 没有设置Cookie
				response.Failed(c, "未授权")
				c.Abort()
				return
			}
			response.Failed(c, "未知错误")
			c.Abort()
			return
		}

		tokenString := ck.Value
		claims := &Claims{}

		jwtKey := []byte(conf.Cfg.Section("JWT").Key("secret_key").String())

		//  解析JWT字符串并吧结果存储在claims中
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				response.Failed(c, "未授权")
				c.Abort()
				return
			}
			response.Failed(c, "未知错误")
			c.Abort()
			return
		}
		if !token.Valid {
			response.Failed(c, "未授权")
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
