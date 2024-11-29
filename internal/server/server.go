package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.mongodb.org/mongo-driver/mongo"
	"graphql/config"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	cfg      *config.Config
	fiberApp *fiber.App
	psqlDB   *pgxpool.Pool
	mngDB    *mongo.Database
}

func NewServer(
	cfg *config.Config,
	fiberApp *fiber.App,
	psqlDB *pgxpool.Pool,
	mngDB *mongo.Database,
) *Server {
	return &Server{
		cfg:      cfg,
		fiberApp: fiberApp,
		psqlDB:   psqlDB,
		mngDB:    mngDB,
	}
}

func (s *Server) Run() error {
	log.Println("Trying to run server...")

	s.MapHandlers()

	go func() {
		fiberAddress := fmt.Sprintf("%s:%s", s.cfg.Fiber.Host, s.cfg.Fiber.Port)

		s.fiberApp.Get("/health_check", func(c *fiber.Ctx) error {
			return c.SendStatus(fiber.StatusOK)
		})
		log.Println("Fiber server is started on " + fiberAddress)

		if err := s.fiberApp.Listen(fiberAddress); err != nil {
			log.Fatal(err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	return nil
}
