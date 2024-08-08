package user

import (
	"context"
	"net/http"

	"gowt/internal/domain"
	"gowt/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (h *UserHandlerImpl) Login(c *gin.Context) {
	log.Info().Msg("Login Handle")

	user := &domain.User{
		Username: "",
		Password: "",
	}

	response := &domain.Response{
		Code:    0,
		Message: "",
	}

	defer c.JSON(response.Code, response)

	response.Code = http.StatusBadRequest

	name := c.Query("username")
	if !utils.IsValid(name) {
		response.Message = "error: username is empty or not valid"
		response.Payload = name
		return
	}
	user.Username = name

	password := c.Query("password")
	if !utils.IsValid(password) {
		response.Message = "error: password is empty or not valid"
		response.Payload = password
		return
	}
	user.Password = password

	ctx := context.WithValue(c.Request.Context(), domain.Str("user"), user)

	data, err := h.uc.Login(ctx)
	response.Payload = data

	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = "failed to login"
		log.Error().Err(err).Msg("failed to login")
		return
	}

	response.Code = http.StatusOK
	response.Message = http.StatusText(response.Code)
}
