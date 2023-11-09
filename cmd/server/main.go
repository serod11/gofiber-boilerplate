/*
 * @author serod
 */

package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/serod11/gofiber-boilerplate/pkg/config"
	"github.com/serod11/gofiber-boilerplate/pkg/repo"
	"github.com/serod11/gofiber-boilerplate/pkg/service"
	"github.com/serod11/gofiber-boilerplate/pkg/utils/db"
	"log"
	"strconv"
)

func main() {
	// load config from env
	c, e := config.LoadConfig()
	if e != nil {
		log.Fatalln("Failed at config", e)
	}

	// initialize db connection
	conn := db.Init(c)

	// dependency injection
	bookService := service.BookService{
		BookRepo: &repo.BookRepository{DB: conn},
	}

	// REST API server [gofiber]
	app := fiber.New()
	app.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{"status": "OK"})
	})

	bookRoute := app.Group("/book")
	bookRoute.Get("/", func(ctx *fiber.Ctx) error {
		books, _ := bookService.GetAllBooks()
		return ctx.JSON(books)
	})
	bookRoute.Get("/:id", func(ctx *fiber.Ctx) error {
		id, err := strconv.ParseUint(ctx.Params("id"), 10, 0)
		if err != nil {
			return ctx.SendStatus(fiber.StatusBadRequest)
		}
		book, _ := bookService.GetBook(uint(id))
		return ctx.JSON(book)
	})

	app.Listen(":8080") // use c.Port instead
}
