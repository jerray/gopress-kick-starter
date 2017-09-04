package main

import (
	"github.com/fpay/gopress"
	"github.com/jerray/gopress-kick-starter/controllers"
	"github.com/jerray/gopress-kick-starter/middlewares"
	"github.com/jerray/gopress-kick-starter/services"
)

func main() {
	s := gopress.NewServer(gopress.ServerOptions{
		Port: 3000,
	})

	s.RegisterServices(
		services.NewDatabaseService(),
	)

	s.RegisterGlobalMiddlewares(
		gopress.NewLoggingMiddleware("global"),
		middlewares.NewAuthMiddleware(),
	)

	s.RegisterControllers(
		controllers.NewUsersController(),
		controllers.NewPostsController(),
	)

	s.Start()
}
