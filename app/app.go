package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	router *mux.Router
}

func (app *App) InitializeAndStart() {
	app.router = mux.NewRouter()
	app.setupRoutes()

	log.Fatal(http.ListenAndServe(":8080", app.router))
}

func (app *App) setupRoutes() {
	app.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("router")
	})
}
