package main

import (
	"github.com/fpay/gopress"
	"github.com/jerray/gopress-kick-starter/controllers"
	"github.com/jerray/gopress-kick-starter/services"
)

func main() {
	s := gopress.NewServer(gopress.ServerOptions{
		Port: 3000,
	})

	s.RegisterGlobalMiddlewares(
		gopress.NewLoggingMiddleware("global"),
	)

	s.RegisterServices(
		services.NewDatabaseService(),
	)

	s.RegisterControllers(
		new(controllers.Users),
		new(controllers.Posts),
	)

	s.Start()
}
