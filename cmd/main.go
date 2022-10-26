package main

import (
	"flag"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tetsurowada/test_server_2022/internal/app"
)

var port = flag.String("port", "13233", "port to listen")

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", app.Hello)
	e.GET("/api/items", app.GetOwnedItems)

	e.Logger.Fatal(e.Start(":" + *port))
}

// ハンドラーを定義
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
