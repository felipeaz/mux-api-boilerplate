package app

import (
	"github.com/gorilla/mux"
	"log"
	"mux-api-boilerplate/internal/app/domain"
	"net/http"
)

type App struct {
	handler domain.Handler
	router  *mux.Router
}

func NewApp(handler domain.Handler) *App {
	return &App{
		handler: handler,
	}
}

func (app *App) InitializeRoutes() {
	app.router.HandleFunc("/sample", app.handler.Get).Methods(http.MethodGet)
	app.router.HandleFunc("/sample/{id}", app.handler.Find).Methods(http.MethodGet)
	app.router.HandleFunc("/sample", app.handler.Post).Methods(http.MethodPost)
	app.router.HandleFunc("/sample/{id}", app.handler.Put).Methods(http.MethodPut)
	app.router.HandleFunc("/sample/{id}", app.handler.Patch).Methods(http.MethodPatch)
	app.router.HandleFunc("/sample/{id}", app.handler.Delete).Methods(http.MethodDelete)
}

func (app *App) Run(addr string) {
	err := http.ListenAndServe(addr, app.router)
	if err != nil {
		log.Fatal(err)
	}
}
