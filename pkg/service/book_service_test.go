/*
 * @author serod
 */

package service

import (
	"github.com/serod11/gofiber-boilerplate/pkg/model"
	"testing"
)

// MockBookRepository mock implementation of BookRepo for testing purpose
type MockBookRepository struct {
}

func (b *MockBookRepository) FindAll() []model.Book {
	return []model.Book{{ID: 1, Name: "Clean Code", Page: 464}}
}

func (b *MockBookRepository) FindById(id uint) model.Book {
	return model.Book{ID: id, Name: "Clean Code", Page: 464}
}

func TestFindBook(t *testing.T) {
	bookService := BookService{
		BookRepo: &MockBookRepository{},
	}
	books, _ := bookService.GetAllBooks()
	//book, err := bookService.GetBook(1)
	t.Logf("%+v\n", books)
}
