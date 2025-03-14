package pages

import "github.com/gofiber/fiber/v2"

type PagesHandler struct {
	router fiber.Router
}

func NewPagesHandler(router fiber.Router) {
	handler := &PagesHandler{
		router: router,
	}

	handler.router.Get("/pages", handler.main)
}

func (handler *PagesHandler) main(c *fiber.Ctx) error {
	return c.SendString("Pages main")
}
