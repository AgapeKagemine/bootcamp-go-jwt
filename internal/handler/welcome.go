package handler

import (
	"net/http"

	"gowt/internal/utils/jwt"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func Welcome(c *gin.Context) {
	log.Info().Msg("Welcome Handler")

	auth := c.GetHeader("Authorization")
	if auth == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "authorization header required",
		})
		return
	}

	err := jwt.VerifyToken(auth)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		log.Error().Err(err).Msg("failed to verify token")
		return
	}

	c.JSON(http.StatusOK, http.StatusText(http.StatusOK))
}
