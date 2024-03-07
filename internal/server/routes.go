package server

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"

	"htmx-fiber/cmd/web"
	"htmx-fiber/internal/database"
	"htmx-fiber/internal/handlers"
)

func (s *FiberServer) RegisterFiberRoutes() {
	s.App.Get("/health", s.healthHandler)

	s.App.Static("/js", "./cmd/web/js")

	s.App.Get("/", adaptor.HTTPHandler(templ.Handler(web.Home())))

	// departments endpoints
	department := s.App.Group("/department")

	department.Post("/create", func(c *fiber.Ctx) error {
		return handlers.HandleViewCreatePage(c, database.DB)
	})
	department.Get("/create", func(c *fiber.Ctx) error {
		return handlers.HandleViewCreatePage(c, database.DB)
	})
	department.Get("/list", func(c *fiber.Ctx) error {
		return handlers.HandleViewDepartmentList(c, database.DB)
	})
	department.Delete("/delete/:id", func(c *fiber.Ctx) error {
		return handlers.HandleViewDeleteDepartment(c, database.DB)
	})
	department.Get("/detail/:id", func(c *fiber.Ctx) error {
		return handlers.HandleViewDepartmentDetail(c, database.DB)
	})

	// reviews endpoints
	review := s.App.Group("/review")

	review.Get("/create", func(c *fiber.Ctx) error {
		return handlers.HandleViewCreateReview(c, database.DB)
	})
	review.Post("/create", func(c *fiber.Ctx) error {
		return handlers.HandleViewCreateReview(c, database.DB)
	})
	review.Get("/success", adaptor.HTTPHandler(templ.Handler(web.CreateReviewSuccessPage())))

	// position endpoints
	position := s.App.Group("/position")
	position.Get("/detail/:id", func(c *fiber.Ctx) error {
		return handlers.HandleViewPositionDetail(c, database.DB)
	})
}

func (s *FiberServer) healthHandler(c *fiber.Ctx) error {
	return c.JSON(s.db.Health())
}
