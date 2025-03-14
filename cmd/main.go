package main

import (
	"strconv"
	"variasco/go-fiber-hw/config"
	"variasco/go-fiber-hw/internal/pages"

	"github.com/gofiber/fiber/v2"
)

func main() {
	conf := config.LoadConfig()

	app := fiber.New()

	pages.NewPagesHandler(app)

	app.Listen(":" + strconv.Itoa(conf.Server.Port))
}
