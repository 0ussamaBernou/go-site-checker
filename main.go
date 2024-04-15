package main

import (
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", root)

	log.Fatal(e.Start(":1323"))
}

func root(c echo.Context) error {
	return hello("You").Render(c.Request().Context(), c.Response())
}
