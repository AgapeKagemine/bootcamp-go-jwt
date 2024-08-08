package server

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gowt/internal/handler/user"
	"gowt/internal/provider/database"
	"gowt/internal/provider/routes"
	"gowt/internal/provider/server/domain"
	"gowt/internal/repository"
	"gowt/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type Server struct {
	Server *gin.Engine
}

func autowired(db *sql.DB) user.UserHandler {
	repo := repository.NewUserRepository(db)
	uc := usecase.NewUserUsecase(repo)
	return user.NewUserHandler(uc)
}

func Run() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	defer stop()

	serverCfg := domain.NewServerConfig()

	db, err := database.NewDB()
	if err != nil {
		log.Fatal().Err(err).Msg("Error connecting to database")
	}

	defer db.Close()

	h := autowired(db)

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", serverCfg.Address, serverCfg.Port),
		Handler: routes.NewRoute(h).Route,
	}

	go func() {
		log.Info().Msg(fmt.Sprintf("Starting server on port %d...", serverCfg.Port))

		err := server.ListenAndServe()

		if err != nil && err != http.ErrServerClosed {
			log.Error().Err(err).Msg("Error starting server")
		}
	}()

	<-ctx.Done()
	stop()
	log.Info().Msg("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = server.Shutdown(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("Error shutting down server")
	}

	log.Info().Msg("HTTP server stopped")
}
