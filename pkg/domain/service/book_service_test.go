package service

import (
	"github.com/ing-bank/gormtestutil"
	"github.com/serod11/gofiber-boilerplate/pkg/domain/command"
	"github.com/serod11/gofiber-boilerplate/pkg/domain/model"
	"github.com/serod11/gofiber-boilerplate/pkg/domain/repository"
	"testing"
)

func TestBook(t *testing.T) {
	db := gormtestutil.NewMemoryDatabase(t)
	db.AutoMigrate(&model.Book{})

	bookService := NewBookService(repository.NewBookRepo(db))
	bookService.AddBook(command.CreateBookCommand{
		Page: 12,
		Name: "GG",
	})
}
