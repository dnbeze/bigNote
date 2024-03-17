package main

import (
	"github.com/a-h/templ"
	"github.com/dnbeze/bigNote/components"
	"github.com/labstack/echo"
)

func main() {

	e := echo.New()
	e.GET("/", templ.Handler(components.HomePage()))
	e.Logger.Fatal(e.Start(":8080"))

}
