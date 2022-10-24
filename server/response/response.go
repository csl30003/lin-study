package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{200, message, data})
}

func Failed(c *gin.Context, message string) {
	c.JSON(http.StatusOK, Response{400, message, nil})
}
