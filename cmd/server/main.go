/*
 * @author serod
 */

package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/serod11/gofiber-boilerplate/pkg/adapters/rest"
	"github.com/serod11/gofiber-boilerplate/pkg/domain/repository"
	"github.com/serod11/gofiber-boilerplate/pkg/domain/service"
	"github.com/serod11/gofiber-boilerplate/pkg/infra/postgres"
	"github.com/serod11/gofiber-boilerplate/pkg/utils/config"
	"log"
)

func main() {
	// load config from env
	c, e := config.LoadConfig()
	if e != nil {
		log.Fatalln("Failed at config", e)
	}

	// initialize postgres connection
	conn := postgres.Init(c)

	// REST API server [gofiber]
	app := fiber.New()
	app.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{"status": "OK"})
	})

	bookRepo := repository.NewBookRepo(conn)
	bookService := service.NewBookService(bookRepo)

	rest.RegisterRoutes(app, bookService)

	app.Listen(":8080") // use c.Port instead
}
