package main

import (
	"strconv"
	"variasco/go-fiber-hw/config"
	"variasco/go-fiber-hw/internal/home"
	log "variasco/go-fiber-hw/pkg/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	slogFiber "github.com/samber/slog-fiber"
)

func main() {
	conf := config.LoadConfig()
	logger := log.CreateLogger(conf.Log)

	app := fiber.New()

	app.Use(slogFiber.New(logger))
	app.Use(recover.New())

	app.Static("/public", "public")

	home.NewHomeHandler(app)

	app.Listen(":" + strconv.Itoa(conf.Server.Port))
}
