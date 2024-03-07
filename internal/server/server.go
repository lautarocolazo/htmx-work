package server

import (
	"github.com/gofiber/fiber/v2"

	"htmx-fiber/internal/database"
)

type FiberServer struct {
	*fiber.App
	db database.Service
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(),
		db:  database.New(),
	}
	return server
}

func (s *FiberServer) SetupStatic() {
	s.Static("/assets", "./assets")
}
