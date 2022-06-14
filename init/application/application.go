package application

import (
	"mux-api-boilerplate/api/app"
	"mux-api-boilerplate/internal/app/handler"
	"mux-api-boilerplate/internal/app/service"
)

func Run() error {
	h := handler.NewHandler(
		service.NewService(),
	)
	a := app.NewApp(h)
	a.Run(":8080")
	return nil
}
