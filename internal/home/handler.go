package home

import (
	"github.com/gofiber/fiber/v2"
)

type HomeHandler struct {
	router fiber.Router
}

type Data struct {
	Tags      []Tag
	Names     []string
	AdminName string
}

type Tag struct {
	ID   int
	Name string
}

var data = Data{Tags: []Tag{
	{ID: 1, Name: "#Еда"},
	{ID: 2, Name: "#Животные"},
	{ID: 3, Name: "#Машины"},
	{ID: 4, Name: "#Спорт"},
	{ID: 5, Name: "#Музыка"},
	{ID: 6, Name: "#Технологии"},
	{ID: 7, Name: "#Прочее"},
}}

func NewHomeHandler(router fiber.Router) {
	handler := &HomeHandler{
		router: router,
	}

	handler.router.Get("/", handler.main)
}

func (handler *HomeHandler) main(c *fiber.Ctx) error {
	return c.Render("base", data)
}
