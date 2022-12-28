package handlers

import (
	"AnarchyStock/helpers"
	"AnarchyStock/models"
	"os"

	"github.com/gin-gonic/gin"
)

func CurrentUser(c *gin.Context) {

	user_id, err := helpers.ExtractTokenID(c)

	if err != nil {
		ErrorMessage(c, err.Error())
		return
	}

	u, err := models.GetUserByID(user_id)

	if err != nil {
		ErrorMessage(c, err.Error())
		return
	}

	DataMessage(c, u)
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {

	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		ErrorMessage(c, err.Error())
		return
	}

	u := models.User{}

	u.Username = input.Username
	u.Password = input.Password

	token, err := models.LoginCheck(u.Username, u.Password)

	if err != nil {
		ErrorMessage(c, "username or password is incorrect.")
		return
	}

	c.SetCookie("Authorization", token, GetTokenExpiration(), "/", os.Getenv("DOMAIN"), false, true)
	SuccessMessage(c, "User logged in")
}

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

func Register(c *gin.Context) {

	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		ErrorMessage(c, err.Error())
		return
	}

	u := models.User{}

	u.Username = input.Username
	u.Password = input.Password

	_, err := u.SaveUser()

	if err != nil {
		ErrorMessage(c, err.Error())
		return
	}

	SuccessMessage(c, "Registration Success")
}
