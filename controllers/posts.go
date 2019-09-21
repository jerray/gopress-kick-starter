package controllers

import (
	"net/http"
	"strconv"

	"github.com/fpay/gopress"
	starter "github.com/jerray/gopress-kick-starter"
	"github.com/jerray/gopress-kick-starter/middlewares"
)

// PostsController 文章控制器
type PostsController struct {
	postService starter.PostService
}

// NewPostsController returns posts controller instance
func NewPostsController(postService starter.PostService) *PostsController {
	return &PostsController{postService: postService}
}

// ViewPost 查看指定的一篇文章
func (c *PostsController) ViewPost(ctx gopress.Context) error {
	id := ctx.Param("id")
	postID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return err
	}

	post, err := c.postService.GetPostByID(postID)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "failed to get post")
	}

	data := map[string]interface{}{
		"post": post,
	}
	return ctx.Render(http.StatusOK, "posts/detail", data)
}

// Profile 查看文章列表
func (c *PostsController) ListPosts(ctx gopress.Context) error {
	user := middlewares.ExtractUser(ctx)

	posts, err := c.postService.ListPostsByUser(user)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "failed to get posts")
	}
	data := map[string]interface{}{
		"user":  user,
		"posts": posts,
	}
	return ctx.Render(http.StatusOK, "posts/list", data)
}
