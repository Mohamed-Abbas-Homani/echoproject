package main

import (
	"echoproject/internal/routes"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	routes.InitRoutes()
}
func main() {
	routes.Run()
}
