package main

import (
	"fmt"
	"variasco/go-fiber-hw/config"
	"variasco/go-fiber-hw/internal/home"
	"variasco/go-fiber-hw/internal/user"
	log "variasco/go-fiber-hw/pkg/logger"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	conf := config.LoadConfig()
	logger := log.NewLogger(conf.Log)

	app := fiber.New()

	app.Use(fiberzerolog.New(fiberzerolog.Config{Logger: logger}))
	app.Use(recover.New())

	app.Static("/public", "public")

	home.NewHandler(home.HomeHandlerDeps{
		Router: app,
	})
	user.NewHandler(user.UserHandlerDeps{
		Router: app,
		Logger: logger,
	})

	app.Listen(fmt.Sprintf(": %d", conf.Server.Port))
}
