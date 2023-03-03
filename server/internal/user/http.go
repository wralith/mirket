package user

import "github.com/labstack/echo/v4"

func InitHTTPEndpoints(g *echo.Group, cnt *HTTPController) {
	g.POST("/register", cnt.Register)
	g.POST("/login", cnt.Login)
}
