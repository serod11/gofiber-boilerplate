/*
 * @author serod
 */

package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/serod11/gofiber-boilerplate/pkg/config"
	"github.com/serod11/gofiber-boilerplate/pkg/handlers"
	"github.com/serod11/gofiber-boilerplate/pkg/utils/db"
	"log"
)

func main() {
	// load config from env
	c, e := config.LoadConfig()
	if e != nil {
		log.Fatalln("Failed at config", e)
	}

	// initialize db connection
	conn := db.Init(c)

	// REST API server [gofiber]
	app := fiber.New()
	app.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{"status": "OK"})
	})

	handlers.RegisterRoutes(app, conn)

	app.Listen(":8080") // use c.Port instead
}
