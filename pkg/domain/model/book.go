/*
 * @author serod
 */

package model

import "github.com/google/uuid"

type Book struct {
	ID   uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Name string    `gorm:"size:128" json:"name"`
	Page uint      `gorm:"" json:"page"`
}
