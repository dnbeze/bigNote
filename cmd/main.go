package main

import (
	"github.com/dnbeze/bigNote/db"
	"github.com/dnbeze/bigNote/handlers"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func main() {

	godotenv.Load(".env")

	db.InitDB()

	e := echo.New()
	e.GET("/", handlers.ServeIndex)
	e.Logger.Fatal(e.Start(":8080"))

}
