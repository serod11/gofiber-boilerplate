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
	books, err := s.BookRepo.FindAll()
	return books, err
}

func (s *BookService) GetBook(id uint) (model.Book, error) {
	book, err := s.BookRepo.FindById(id)
	return book, err
}

func (s *BookService) AddBook(req model.BookRequest) error {
	txn := s.BookRepo.CreateTxn()
	err := s.BookRepo.Create(txn, model.Book{
		Name: req.Name,
		Page: req.Page,
	})
	if err != nil {
		return err
	}
	if txn != nil {
		txn.Commit()
	}
	return nil
}
