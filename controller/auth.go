package controller

import (
	"galon-app/models"
	token "galon-app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserReg struct {
	Fullname string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role_id  int    `json:"role_id" binding:"required,numeric"`
}
type UserLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var input UserLogin
	if err := c.ShouldBindJSON(&input); err != nil {
		ApiErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	u := models.User{}
	u.Username = input.Username
	u.Password = input.Password

	token, err := models.LoginCheck(u.Username, u.Password)
	if err != nil {
		ApiErrorResponse(c, http.StatusBadRequest, "username or password is incorrect")
		return
	}
	ApiResponse(c, http.StatusOK, "User validated.", token)
}

func Register(c *gin.Context) {
	var input UserReg
	if err := c.ShouldBindJSON(&input); err != nil {
		ApiErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	u := models.User{}
	u.Fullname=input.Fullname
	u.Username = input.Username
	u.Email=input.Email
	u.Password = input.Password
	u.Role_id = input.Role_id
	err := u.Create()
	if err != nil {
		ApiErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	ApiResponse(c, http.StatusOK, "Registrarion success")
}
func CurrentUser(c *gin.Context) {
	user_id, err := token.ExtractTokenID(c)
	if err != nil {
		ApiErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	u, err := models.GetUserByID(user_id)
	if err != nil {
		ApiErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	ApiResponse(c, http.StatusOK, "Success", u)
}
