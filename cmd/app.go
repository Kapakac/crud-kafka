package main

import (
	"github.com/Kapakac/crud-kafka/internal/pkg/database"
	"github.com/Kapakac/crud-kafka/internal/pkg/handler"
	"github.com/Kapakac/crud-kafka/internal/pkg/middleware"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) Initialize() *App {
	a.Router = mux.NewRouter()
	a.DB = database.ConnectionDatabase()
	a.initializeRoutes()

	return a
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initializeRoutes() {
	middleware.Get(a.Router, "/countries", handler.GetCountries)
	middleware.Get(a.Router, "/countries/{id}", handler.GetCountry)
}
