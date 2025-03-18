package home

import (
	"variasco/go-fiber-hw/pkg/tadapter"
	"variasco/go-fiber-hw/views/components"
	"variasco/go-fiber-hw/views/pages"

	"github.com/gofiber/fiber/v2"
)

type HomeHandler struct {
	router fiber.Router
}

var mockData = pages.MainProps{
	Tags: []components.Tag{
		{ID: 1, Name: "#Еда"},
		{ID: 2, Name: "#Животные"},
		{ID: 3, Name: "#Машины"},
		{ID: 4, Name: "#Спорт"},
		{ID: 5, Name: "#Музыка"},
		{ID: 6, Name: "#Технологии"},
		{ID: 7, Name: "#Прочее"},
	},
}

func NewHomeHandler(router fiber.Router) {
	handler := &HomeHandler{
		router: router,
	}

	handler.router.Get("/", handler.main)
}

func (handler *HomeHandler) main(c *fiber.Ctx) error {
	return tadapter.Render(c, pages.Main(mockData))
}
