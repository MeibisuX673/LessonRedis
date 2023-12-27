package main

import (
	"github.com/MeibisuX673/LessonRedis/app/router"
	"github.com/MeibisuX673/LessonRedis/config/database"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	err := database.AllDatabases.Init()
	if err != nil {
		log.Fatal(err)
	}

	ge := router.AppRouter()

	ge.Run(":8080")

}
