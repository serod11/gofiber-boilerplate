/*
 * @author serod
 */

package model

type Book struct {
	ID   uint   `gorm:"primarykey" json:"id"`
	Name string `gorm:"size:128" json:"name"`
	Page uint   `gorm:"" json:"page"`
}

type BookRequest struct {
	Name string `json:"name"`
	Page uint   `json:"page"`
}
