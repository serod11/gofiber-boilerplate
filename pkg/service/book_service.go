/*
 * @author serod
 */

package service

import (
	"github.com/serod11/gofiber-boilerplate/pkg/model"
	"github.com/serod11/gofiber-boilerplate/pkg/repo"
)

type BookService struct {
	BookRepo repo.BookRepo
}

func (s *BookService) GetAllBooks() ([]model.Book, error) {
	books := s.BookRepo.FindAll()
	return books, nil
}

func (s *BookService) GetBook(id uint) (model.Book, error) {
	book := s.BookRepo.FindById(id)
	return book, nil
}
