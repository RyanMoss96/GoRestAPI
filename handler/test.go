package handler

import (
	"RyanMoss96/GoRestAPI/database"
	"fmt"
)

func Test() {
	client := database.GetRedisClient()

	pong, err := client.Ping().Result()

	if err != nil {

	}

	fmt.Println("Test" + pong)
}
