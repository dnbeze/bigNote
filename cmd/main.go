package main

import (
	"fmt"
	"github.com/dnbeze/bigNote/db"
	"github.com/dnbeze/bigNote/handlers"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"log"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("failed to load env vars %v", err)
	}

	e := echo.New()
	client, err := db.InitDB()
	if err != nil {
		log.Fatalf("failed connecting to Mongo %v", err)
	}

	fmt.Print(client)
	fmt.Println("Connected to MONGO!~~~~!!!~~~")
	e.GET("/", handlers.ServeIndex)
	e.Logger.Fatal(e.Start(":8080"))

}
