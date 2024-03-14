package service

import (
	"errors"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func Run() {
	e := echo.New()

	//renderer := &TemplateRenderer{
	//	templates: template.Must(template.ParseGlob("*.html")),
	//}
	//e.Renderer = renderer

	e.POST("/price", GetPrice)
	e.GET("/data", GetData)
	e.POST("/send/baseline", AddBaseline)
	e.POST("/send/discounts", AddDiscounts)
	e.POST("/update/storage", UpdateStorage)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	if err := e.Start(":2323"); !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
