/*
 * @author serod
 */

package service

import (
	"github.com/serod11/gofiber-boilerplate/pkg/model"
	"gorm.io/gorm"
)

type BookService struct {
	DB *gorm.DB
}

func (s *BookService) GetAllBooks() ([]model.Book, error) {
	var books []model.Book
	if err := s.DB.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (s *BookService) GetBook(id uint) (model.Book, error) {
	var book model.Book
	s.DB.First(&book, id) // SELECT * FROM book where id = 1
	return book, nil
}

func (s *BookService) AddBook(req model.BookRequest) error {
	s.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&model.Book{
			Name: req.Name,
			Page: req.Page,
		}).Error; err != nil {
			return err
		}
		return nil
	})
	return nil
}
