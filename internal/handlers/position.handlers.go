package handlers

import (
	"database/sql"
	"strconv"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"

	"htmx-fiber/cmd/web"
	"htmx-fiber/internal/models"
)

// Handler
func HandleViewPositionDetail(c *fiber.Ctx, db *sql.DB) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	jobPosition, err := models.GetJobPositionByID(db, id)
	if err != nil {
		return err
	}

	reviews, err := models.GetReviewsByJobPositionID(db, id)
	if err != nil {
		return err
	}

	positionDetail := web.JobPositionDetail(*jobPosition, reviews)
	handler := adaptor.HTTPHandler(templ.Handler(positionDetail))
	return handler(c)
}
