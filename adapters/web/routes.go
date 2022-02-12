package web

import (
	"github.com/labstack/echo"
)

func (w *WebServer) buildRoutes(echoServer *echo.Echo) {

	groupV1 := echoServer.Group("/api")

	w.pointsRoutes(groupV1)
}

func (w *WebServer) pointsRoutes(route *echo.Group) {

	route.GET("/points", w.HandlerGetPoints)
}
