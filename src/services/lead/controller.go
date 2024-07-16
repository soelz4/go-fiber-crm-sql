package lead

import (
	"encoding/json"

	"github.com/gofiber/fiber/v3"

	"go-fiber-crm-sql/src/types"
)

func (h *Handler) GetLeads(c fiber.Ctx) error {
	var leads []types.Lead
	// SQL Query
	h.store.db.Find(&leads)
	// Convert to JSON
	return c.JSON(leads)
}

func (h *Handler) GetLeadByID(c fiber.Ctx) error {
	var lead types.Lead
	id := c.Params("id")
	// SQL Query
	h.store.db.Find(&lead, id)
	// Convert to JSON
	return c.JSON(lead)
}

func (h *Handler) NewLead(c fiber.Ctx) error {
	lead := new(types.Lead)
	// Body
	body := c.Body()
	// Unmarshal Parses the JSON-Encoded Data and Stores the Result
	err := json.Unmarshal(body, lead)
	if err != nil {
		return c.Status(503).SendString("JSON Decoding Failed")
	} else {
		// SQL Query
		h.store.db.Create(&lead)
		// Convert to JSON
		return c.JSON(lead)
	}
}

func (h *Handler) DeleteLeadByID(c fiber.Ctx) error {
	id := c.Params("id")
	var lead types.Lead
	// SQL Query
	h.store.db.First(&lead, id)
	if lead.Name == "" {
		return c.Status(500).SendString("No Lead Found with ID")
	} else {
		h.store.db.Delete(&lead)
		return c.SendString("Lead Successfully Deleted")
	}
}
