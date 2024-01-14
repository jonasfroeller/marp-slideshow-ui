package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/a-h/templ"
    "marp-slideshow-ui/views"
)

func render(ctx echo.Context, cmp templ.Component) error {
	return cmp.Render(ctx.Request().Context(), ctx.Response())
}

func homeHandler(c echo.Context) error {
	return render(c, views.Home("Home"))
}

func main() {
	e := echo.New()
	e.GET("/", homeHandler)

	fmt.Println("Listening on localhost:1323")
	e.Logger.Fatal(e.Start(":1323"))
}