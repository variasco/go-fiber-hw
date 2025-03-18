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
		{ID: 1, Name: "#Еда", Enum: "food"},
		{ID: 2, Name: "#Животные", Enum: "animals"},
		{ID: 3, Name: "#Машины", Enum: "cars"},
		{ID: 4, Name: "#Спорт", Enum: "sport"},
		{ID: 5, Name: "#Музыка", Enum: "music"},
		{ID: 6, Name: "#Технологии", Enum: "tech"},
		{ID: 7, Name: "#Прочее", Enum: "other"},
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
