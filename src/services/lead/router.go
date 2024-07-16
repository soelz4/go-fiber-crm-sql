package lead

import (
	"github.com/gofiber/fiber/v3"
)

// Lead Handler
type Handler struct {
	store *Store
}

// Return New Lead Handler
func NewHandler(store *Store) *Handler {
	return &Handler{store: store}
}

// Routes
func (h *Handler) RegisterRoutes(app *fiber.App) {
	app.Get("/api/v1/leads", h.GetLeads)
	app.Get("/api/v1/lead/:id", h.GetLeadByID)
	app.Post("/api/v1/newlead", h.NewLead)
	app.Delete("/api/v1/lead/:id", h.DeleteLeadByID)
}
