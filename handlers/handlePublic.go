package handlers

import (
	"context"
	"github.com/a-h/templ"
	"github.com/dnbeze/bigNote/components"
	"github.com/labstack/echo"
)

// ServeIndex serves the index.html page
func ServeIndex(c echo.Context) error {
	templ.Handler(components.HomePage()).ServeHTTP(c.Response(), c.Request())
	return nil
}

func TestUser(c echo.Context) error {
	return components.HomePage().Render(context.Background(), c.Response().Writer)
}
