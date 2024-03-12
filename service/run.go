package service

import (
	"errors"
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	"log"
	"net/http"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func Run() {
	e := echo.New()

	//renderer := &TemplateRenderer{
	//	templates: template.Must(template.ParseGlob("*.html")),
	//}
	//e.Renderer = renderer

	e.POST("/send/storage", GetStorage)
	e.POST("/price", GetPrice)
	e.GET("/category/:category", GetCategory)
	e.GET("/location/:location", GetLocation)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	if err := e.Start(":2323"); !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
