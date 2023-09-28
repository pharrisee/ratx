package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pharrisee/ratx/content"
	"github.com/pharrisee/ratx/internal/renderer"
)

func main() {
	e := echo.New()
	e.Renderer = renderer.NewViews(content.ViewsFS)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Filesystem: http.FS(content.StaticFS),
	}))

	e.GET("/", func(c echo.Context) error {
		c.Set("layout", "layout")
		return c.Render(200, "index", nil)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
