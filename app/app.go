package app

import (
	"RyanMoss96/GoRestAPI/database"
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
	app.router = setupRouterAndRoutes()
	redisClient := initializeRedisClient()

	pong, err := redisClient.Ping().Result()

	if err != nil {
		log.Fatal("Unable to create Redis Client. Exiting")
	}

	log.Println(pong + " - Redis Client Successfully Created")

	log.Fatal(http.ListenAndServe(":8080", app.router))
}

func initializeRedisClient() *redis.Client {
	return database.NewRedisClient(redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	}))
}

func setupRouterAndRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("router")
	})

	return router
}
