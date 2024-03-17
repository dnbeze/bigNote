package handlers

import (
	"context"
	"github.com/dnbeze/bigNote/components"
	"github.com/labstack/echo"
	"net/http"
	"os"
)

// ServeIndex serves the index.html page
func ServeIndex(c echo.Context) error {
	htmlContent, err := os.ReadFile("templates/index.html")
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to load index page")
	}
	return c.HTML(http.StatusOK, string(htmlContent))
}

func TestUser(c echo.Context) error {
	return components.HomePage().Render(context.Background(), echo.ResponseWriter())
}
