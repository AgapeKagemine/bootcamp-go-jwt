package routes

import (
	"gowt/internal/handler"
	userHandler "gowt/internal/handler/user"
	userRoutes "gowt/internal/provider/routes/user"

	"github.com/gin-gonic/gin"
)

type Routes struct {
	Route *gin.Engine
}

func NewRoute(h userHandler.UserHandler) *Routes {
	routes := &Routes{
		Route: gin.New(),
	}

	api := routes.Route.Group("/api")
	api.GET("/welcome", handler.Welcome)
	userRoutes.RegisterUserRoutes(api, h)

	return routes
}
