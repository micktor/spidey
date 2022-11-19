package httpd

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/micktor/spidey/internal/config"
	"github.com/micktor/spidey/internal/ent"
)

func Run(config *config.Config, db *ent.Client) {
	a := newApplicationConfig(config, db)

	app := fiber.New()
	//app.Use(cors.New())
	app.Use(logger.New())

	setupRoutes(app, a)

	fmt.Println(app.Listen(":3000"))
}

func setupRoutes(app *fiber.App, a *applicationConfig) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	spideyRoutes := v1.Group("/spidey")
	spideyRoutes.Post("/new", a.spideyHandler.CreateSpidey)
}
