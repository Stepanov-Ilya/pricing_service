package service

import (
	"errors"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func Run() {
	e := echo.New()

	e.POST("/price", GetPrice)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	if err := e.Start(":2323"); !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
