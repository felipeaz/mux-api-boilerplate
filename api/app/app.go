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
		router:  mux.NewRouter(),
	}
}

func (app *App) initializeRoutes() {
	app.router.HandleFunc("/sample", app.handler.Get).Methods(http.MethodGet)
	app.router.HandleFunc("/sample/{id:[0-9]+}", app.handler.Find).Methods(http.MethodGet)
	app.router.HandleFunc("/sample", app.handler.Post).Methods(http.MethodPost)
	app.router.HandleFunc("/sample/{id:[0-9]+}", app.handler.Put).Methods(http.MethodPut)
	app.router.HandleFunc("/sample/{id:[0-9]+}", app.handler.Patch).Methods(http.MethodPatch)
	app.router.HandleFunc("/sample/{id:[0-9]+}", app.handler.Delete).Methods(http.MethodDelete)
}

func (app *App) Run(addr string) {
	log.Println(addr)
	app.initializeRoutes()
	err := http.ListenAndServe(addr, app.router)
	if err != nil {
		log.Fatal(err)
	}
}
