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
	CreateTxn() *gorm.DB
	FindAll() ([]model.Book, error)
	FindById(id uint) (model.Book, error)
	Create(db *gorm.DB, book model.Book) error
}

// BookRepository implementation of BookRepo
type BookRepository struct {
	DB *gorm.DB
}

func (b *BookRepository) CreateTxn() *gorm.DB {
	return b.DB.Begin()
}

func (b *BookRepository) FindAll() ([]model.Book, error) {
	var books []model.Book
	if err := b.DB.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (b *BookRepository) FindById(id uint) (model.Book, error) {
	var book model.Book
	b.DB.First(&book, id) // SELECT * FROM book where id = 1
	return book, nil
}

func (b *BookRepository) Create(txn *gorm.DB, book model.Book) error {
	if txn == nil {
		txn = b.DB
	}
	if err := txn.Create(&book).Error; err != nil {
		txn.Rollback()
		return err
	}
	return nil
}
