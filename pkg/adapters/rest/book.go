/*
 * @author serod
 */

package rest

import "github.com/gofiber/fiber/v2"

// AddBook
// @Summary Add book
// @Description Add book
// @Tags BOOK
// @Success 200
// @Router /book/ [post]
func (h *BookHandler) AddBook(ctx *fiber.Ctx) error {
	body := new(BookRequest)
	if err := ctx.BodyParser(body); err != nil {
		return err
	}
	h.Service.AddBook(body.ToCommand())
	return ctx.SendStatus(fiber.StatusOK)
}
