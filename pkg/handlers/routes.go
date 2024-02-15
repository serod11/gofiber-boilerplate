/*
 * @author serod
 */

package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/serod11/gofiber-boilerplate/pkg/service"
	"gorm.io/gorm"
)

type BookHandler struct {
	Service service.BookService
}

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	h := &BookHandler{
		Service: service.BookService{DB: db},
	}

	group := app.Group("/book")
	group.Get("/", h.GetAllBooks)
	group.Get("/:id", h.GetBook)
	group.Post("/", h.AddBook)
}
