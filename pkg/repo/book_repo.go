/*
 * @author serod
 */

package repo

import (
	"github.com/serod11/gofiber-boilerplate/pkg/model"
	"gorm.io/gorm"
)

// BookRepo Data access layer abstraction
type BookRepo interface {
	FindAll() []model.Book
	FindById(id uint) model.Book
}

// BookRepository implementation of BookRepo
type BookRepository struct {
	DB *gorm.DB
}

func (b *BookRepository) FindAll() []model.Book {
	var books []model.Book
	b.DB.Find(&books) // SELECT * FROM book
	return books
}

func (b *BookRepository) FindById(id uint) model.Book {
	var book model.Book
	b.DB.First(&book, id) // SELECT * FROM book where id = 1
	return book
}
