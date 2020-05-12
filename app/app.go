package app

import (
	"RyanMoss96/GoRestAPI/database"
	"RyanMoss96/GoRestAPI/handler"
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

type App struct {
	router *mux.Router
}

func (app *App) InitializeAndStart() {
	app.router = mux.NewRouter()
	app.setupRoutes()

	database.NewRedisClient(redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	}))

	client := database.GetRedisClient()

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	handler.Test()

	log.Fatal(http.ListenAndServe(":8080", app.router))
}

func (app *App) setupRoutes() {
	app.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("router")
	})
}
