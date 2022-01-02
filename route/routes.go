package route

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"goprojects/urlshortener/api"
)

func InitServer() *echo.Echo {

	e := echo.New()

	e.Use(middleware.Logger())

	e.POST("api/url/short", api.APIShortURL)
	e.GET("/:path", api.APIRedirectToURL)

	return e

}
