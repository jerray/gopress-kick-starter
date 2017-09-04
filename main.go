package main

import (
	"context"
	"os"
	"os/signal"
	"time"

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
		gopress.NewLoggingMiddleware("global", nil),
		middlewares.NewAuthMiddleware(),
	)

	s.RegisterControllers(
		controllers.NewUsersController(),
		controllers.NewPostsController(),
	)

	go func() {
		if err := s.Start(); err != nil {
			s.Logger.Info("server shutdown")
		}
	}()

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		s.Logger.Fatalf("server shutdown error: %s", err)
	}
	time.Sleep(2 * time.Second)
}
