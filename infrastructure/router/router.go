package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"majezanu/capstone/internal/contracts/controller"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	pokemonRoutes := e.Group("/pokemon")
	pokemonRoutes.GET("", func(context echo.Context) error { return c.Pokemon.GetAll(context) })
	pokemonRoutes.GET("/multiple", func(context echo.Context) error { return c.Pokemon.GetMultiple(context) })
	pokemonRoutes.GET("/fetch/:id", func(context echo.Context) error { return c.Pokemon.FetchByIdAndSave(context) })
	pokemonRoutes.GET("/by-name/:name", func(context echo.Context) error { return c.Pokemon.GetByName(context) })
	pokemonRoutes.GET("/:id", func(context echo.Context) error { return c.Pokemon.GetById(context) })

	return e

}
