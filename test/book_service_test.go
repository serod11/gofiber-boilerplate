/*
 * @author serod
 */

package test

import (
	"github.com/ing-bank/gormtestutil"
	"github.com/serod11/gofiber-boilerplate/pkg/model"
	"github.com/serod11/gofiber-boilerplate/pkg/service"
	"testing"
)

func TestBook(t *testing.T) {
	db := gormtestutil.NewMemoryDatabase(t)
	db.AutoMigrate(&model.Book{})
	bookService := service.BookService{
		DB: db,
	}
	bookService.AddBook(model.BookRequest{
		Name: "Peter Pan",
		Page: 250,
	})
	book, err := bookService.GetBook(1)
	if err != nil {
		t.Error(err)
	}
	t.Log(book)
}
