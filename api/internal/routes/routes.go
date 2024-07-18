package routes

import (
	"github.com/JoaoPauloFontana/data-import/api/internal/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{"Content-Type", "Authorization"},
	}))

	e.POST("/login", handlers.Login)

	r := e.Group("/api")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("/records", handlers.GetRecords)

	return e
}
