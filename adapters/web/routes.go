package web

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func (w *WebServer) buildRoutes(echoServer *echo.Echo) {

	groupV1 := echoServer.Group("/api", middleware.BasicAuth(w.Authorization))

	w.pointsRoutes(groupV1)
}

func (w *WebServer) pointsRoutes(route *echo.Group) {

	route.GET("/points", w.HandlerGetPoints)
}
