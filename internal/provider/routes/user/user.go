package user

import (
	"gowt/internal/handler/user"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup, h user.UserHandler) {
	user := rg.Group("/user")
	user.GET("/register", h.Register)
	user.GET("/login", h.Login)
}
