package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})
	e.GET("/user/getUser/:id", getUser)
	e.GET("/user/show", show)
	e.Logger.Fatal(e.Start(":1323"))
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func show(c echo.Context) error {
	name := c.QueryParam("name")
	age := c.QueryParam("age")
	return c.String(http.StatusOK, "name:"+name+",age:"+age)
}
