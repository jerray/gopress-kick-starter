package controllers

import (
	"net/http"

	"github.com/fpay/gopress"
	starter "github.com/jerray/gopress-kick-starter"
	"github.com/jerray/gopress-kick-starter/middlewares"
)

// UsersController 用户控制器
type UsersController struct {
	userService starter.UserService
}

// NewUsersController returns users controller instance
func NewUsersController(userService starter.UserService) *UsersController {
	return &UsersController{userService: userService}
}

// Login 登录action
func (u *UsersController) Login(c gopress.Context) error {
	return c.Render(http.StatusOK, "users/login", nil)
}

// Profile 查看用户详情
func (u *UsersController) Profile(c gopress.Context) error {
	user := middlewares.ExtractUser(c)

	user, err := u.userService.GetUserByID(user.ID)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	data := map[string]interface{}{
		"user": user,
	}
	return c.Render(http.StatusOK, "users/profile", data)
}
