package main

import (
	"context"
	"fmt"
	"github.com/dnbeze/bigNote/db"
	"github.com/dnbeze/bigNote/handlers"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"log"
	"os"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("failed to load env vars")
	}

	e := echo.New()
	client, err := db.InitDB(e.AcquireContext(), os.Getenv("MONGO_URI"))
	if err != nil {
		log.Fatalf("failed connecting to Mongo")
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("failed db ping")
	}
	fmt.Println("Connected to MONGO!~~~~!!!~~~")
	e.GET("/", handlers.ServeIndex)
	e.Logger.Fatal(e.Start(":8080"))

}
