/*
 * @author serod
 */

package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/serod11/gofiber-boilerplate/pkg/model"
	"strconv"
)

// GetAllBooks
// @Summary Get all books
// @Description Get all books
// @Tags BOOK
// @Success 200 {object} []model.Book
// @Router /book/ [get]
func (h *BookHandler) GetAllBooks(ctx *fiber.Ctx) error {
	books, _ := h.Service.GetAllBooks()
	return ctx.JSON(books)
}

// GetBook
// @Summary Get book
// @Description Get book
// @Tags BOOK
// @Success 200 {object} model.Book
// @Router /book/:id [get]
func (h *BookHandler) GetBook(ctx *fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 0)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	book, _ := h.Service.GetBook(uint(id))
	return ctx.JSON(book)
}

// AddBook
// @Summary Add book
// @Description Add book
// @Tags BOOK
// @Success 200
// @Router /book/ [post]
func (h *BookHandler) AddBook(ctx *fiber.Ctx) error {
	body := new(model.BookRequest)
	if err := ctx.BodyParser(body); err != nil {
		return err
	}
	h.Service.AddBook(*body)
	return ctx.SendStatus(fiber.StatusOK)
}
