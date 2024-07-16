package server

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"

	"go-fiber-crm-sql/src/services/lead"
)

// Server
type Server struct {
	address string
	db      *gorm.DB
}

// Return New Server
func NewServer(address string, db *gorm.DB) *Server {
	return &Server{address: address, db: db}
}

// Server RUN
func (s *Server) Run() error {
	// Fiber New App
	app := fiber.New()

	// Lead Handler
	leadStore := lead.NewStore(s.db)
	leadHandler := lead.NewHandler(leadStore)

	// Routes
	leadHandler.RegisterRoutes(app)

	// Logs
	log.Println("Listening On", s.address)

	// Create Server
	return app.Listen(s.address)
}
