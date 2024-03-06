/*
 * @author serod
 */

package service

import (
	"github.com/google/uuid"
	"github.com/serod11/gofiber-boilerplate/pkg/domain/command"
	"github.com/serod11/gofiber-boilerplate/pkg/domain/model"
	"github.com/serod11/gofiber-boilerplate/pkg/domain/repository"
)

type BookService struct {
	BookRepo repository.BookRepo
}

func NewBookService(repo repository.BookRepo) BookService {
	return BookService{
		BookRepo: repo,
	}
}

func (s *BookService) AddBook(bookCommand command.CreateBookCommand) error {
	if txnErr := s.BookRepo.Transactional(func(repo repository.BookRepo) error {
		if err := repo.Create(&model.Book{
			ID:   uuid.New(),
			Name: bookCommand.Name,
			Page: bookCommand.Page,
		}); err != nil {
			return err
		}
		return nil
	}); txnErr != nil {
		return txnErr
	}
	return nil
}
