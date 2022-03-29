package main

import (
	"github.com/Kapakac/funny-sanctions/pkg/database"
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

	return a
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}
