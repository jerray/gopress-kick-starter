package controllers

import (
	"net/http"
	"time"

	"github.com/fpay/gopress"
	"github.com/jerray/gopress-kick-starter/models"
)

// Users 用户控制器
type Users struct {
}

// RegisterRoutes 注册路由
func (u *Users) RegisterRoutes(app *gopress.App) {
	app.GET("/login", u.Login)
	app.GET("/user", u.Profile)
}

// Login 登录action
func (u *Users) Login(c gopress.Context) error {
	return c.Render(http.StatusOK, "users/login", nil)
}

// Profile 查看用户详情
func (u *Users) Profile(c gopress.Context) error {
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
