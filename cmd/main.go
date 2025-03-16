package main

import (
	"strconv"
	"variasco/go-fiber-hw/config"
	"variasco/go-fiber-hw/internal/home"
	"variasco/go-fiber-hw/internal/pages"
	log "variasco/go-fiber-hw/pkg/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
	slogFiber "github.com/samber/slog-fiber"
)

func main() {
	conf := config.LoadConfig()
	logger := log.CreateLogger(conf.Log)

	engine := html.New("html", ".html")
	app := fiber.New(fiber.Config{Views: engine})

	app.Use(slogFiber.New(logger))
	app.Use(recover.New())

	home.NewHomeHandler(app)
	pages.NewPagesHandler(app)

	app.Listen(":" + strconv.Itoa(conf.Server.Port))
}
