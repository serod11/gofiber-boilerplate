package repository

import (
	"fmt"
	"github.com/serod11/gofiber-boilerplate/pkg/domain/model"
	"gorm.io/gorm"
)

type AtomicCallback = func(r BookRepo) error

type bookRepository struct {
	DB *gorm.DB
}

type BookRepo interface {
	Create(book *model.Book) error
	Transactional(callback AtomicCallback) error
}

func NewBookRepo(db *gorm.DB) BookRepo {
	return &bookRepository{DB: db}
}

func (repo *bookRepository) Create(book *model.Book) error {
	err := repo.DB.Create(book).Error
	return err
}

// Transactional rolls back the operations if the callback returns an error
func (repo *bookRepository) Transactional(callback func(txn BookRepo) error) (err error) {
	tx := repo.DB.Begin()

	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				err = fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
			}
		} else {
			tx.Commit()
		}
	}()

	bookRepo := repo.withTx(tx)
	err = callback(bookRepo)
	return
}

func (repo *bookRepository) withTx(tx *gorm.DB) BookRepo {
	br := NewBookRepo(tx)
	return br
}
