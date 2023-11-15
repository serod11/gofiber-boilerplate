/*
 * @author serod
 */

package service

import (
	"github.com/serod11/gofiber-boilerplate/pkg/model"
	"github.com/serod11/gofiber-boilerplate/pkg/service"
	"github.com/serod11/gofiber-boilerplate/test/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestGetAllBooks(t *testing.T) {
	ctl := gomock.NewController(t)
	mockRepo := mock.NewMockBookRepo(ctl)
	mockRepo.EXPECT().FindAll().Return([]model.Book{{ID: 1, Name: "Clean Code", Page: 464}})
	bookService := service.BookService{
		BookRepo: mockRepo,
	}
	books, _ := bookService.GetAllBooks()
	t.Logf("%+v\n", books)
	assert.Equal(t, "Clean Code", books[0].Name, "Name doesn't match")
}

func TestGetBook(t *testing.T) {
	ctl := gomock.NewController(t)
	mockRepo := mock.NewMockBookRepo(ctl)
	mockRepo.EXPECT().FindById(uint(1)).Return(model.Book{ID: 1, Name: "Harry Potter", Page: 389})
	bookService := service.BookService{
		BookRepo: mockRepo,
	}
	book, err := bookService.GetBook(1)
	if err != nil {
		t.Error(err)
	}
	t.Log(book)
	assert.Equal(t, uint(389), book.Page, "Book page should be 389")
}
