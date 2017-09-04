package controllers

import (
	"net/http"
	"time"

	"github.com/fpay/gopress"
	"github.com/jerray/gopress-kick-starter/models"
)

// UsersController 用户控制器
type UsersController struct {
	// Uncomment this line if you want to use services in the app
	// app *gopress.App
}

// NewUsersController returns users controller instance
func NewUsersController() *UsersController {
	return new(UsersController)
}

// RegisterRoutes 注册路由
func (u *UsersController) RegisterRoutes(app *gopress.App) {
	// Uncomment this line if you want to use services in the app
	// c.app = app
	app.GET("/login", u.Login)
	app.GET("/user", u.Profile)
}

// Login 登录action
func (u *UsersController) Login(c gopress.Context) error {
	return c.Render(http.StatusOK, "users/login", nil)
}

// Profile 查看用户详情
func (u *UsersController) Profile(c gopress.Context) error {
	// Or you can get app from request context
	// app := gopress.AppFromContext(ctx)
	user := &models.User{
		ID:        uint64(1),
		Name:      "Admin",
		CreatedAt: time.Now(),
	}
	data := map[string]interface{}{
		"user": user,
	}
	return c.Render(http.StatusOK, "users/profile", data)
}
