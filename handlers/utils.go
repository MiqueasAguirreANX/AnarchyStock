package handlers

import (
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetTokenExpiration() (exp int) {
	expString := os.Getenv("TOKEN_EXP")
	exp64, err := strconv.ParseInt(expString, 10, 32)
	if err != nil {
		return 3600
	}
	return int(exp64)
}

func ErrorMessage(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status":  "error",
		"message": message,
	})
}

func SuccessMessage(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": message,
	})
}

func DataMessage(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": data,
	})
}
