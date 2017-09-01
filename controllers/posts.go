package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/fpay/gopress"
	"github.com/jerray/gopress-kick-starter/models"
	"github.com/jerray/gopress-kick-starter/services"
)

// Posts 用户控制器
type Posts struct {
	app *gopress.App
	db  *services.DatabaseService
}

// RegisterRoutes 注册路由
func (u *Posts) RegisterRoutes(app *gopress.App) {
	u.app = app
	u.db = app.Services.Get(services.DatabaseServiceName).(*services.DatabaseService)

	app.GET("/posts", u.ListPosts)
	app.GET("/posts/:id", u.ViewPost)
}

// ViewPost 查看指定的一篇文章
func (u *Posts) ViewPost(c gopress.Context) error {
	id := c.Param("id")
	postID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return err
	}

	title := fmt.Sprintf("Article %d", postID)
	post := &models.Post{
		ID:      postID,
		UserID:  1,
		Title:   title,
		Content: "Generated content! Using service: " + u.db.ServiceName(),
	}

	data := map[string]interface{}{
		"post": post,
	}
	return c.Render(http.StatusOK, "posts/detail", data)
}

// Profile 查看文章列表
func (u *Posts) ListPosts(c gopress.Context) error {
	user := &models.User{
		ID:        uint64(1),
		Name:      "Admin",
		CreatedAt: time.Now(),
	}

	posts := []*models.Post{
		{ID: 1, UserID: 1, Title: "Echo framework", Content: "Awesome!"},
		{ID: 2, UserID: 1, Title: "Handlebars template engine", Content: "Powerful!"},
		{ID: 3, UserID: 2, Title: "Using GORM", Content: "Handsome!"},
		{ID: 4, UserID: 1, Title: "Static supported", Content: "Yo!"},
	}
	data := map[string]interface{}{
		"user":  user,
		"posts": posts,
	}
	return c.Render(http.StatusOK, "posts/list", data)
}
