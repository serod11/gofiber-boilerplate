/*
 * @author serod
 */

package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/serod11/gofiber-boilerplate/pkg/domain/service"
)

type BookHandler struct {
	Service service.BookService
}

func RegisterRoutes(app *fiber.App, bookService service.BookService) {
	h := &BookHandler{
		Service: bookService,
	}

	group := app.Group("/book")
	group.Post("/", h.AddBook)
}
