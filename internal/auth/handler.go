package auth

import (
	"variasco/go-fiber-hw/pkg/tadapter"
	"variasco/go-fiber-hw/views/pages"

	"github.com/gofiber/fiber/v2"
)

type AuthHandlerDeps struct {
	Router fiber.Router
}

type AuthHandler struct {
	router fiber.Router
}

func NewHandler(deps AuthHandlerDeps) {
	handler := &AuthHandler{
		router: deps.Router,
	}

	handler.router.Get("/register", handler.register)
}

func (handler *AuthHandler) register(c *fiber.Ctx) error {
	return tadapter.Render(c, pages.Register())
}
