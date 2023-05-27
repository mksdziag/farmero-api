package categories

import "github.com/google/uuid"

type Category struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Name        string    `json:"name" db:"name" validate:"required,min=3,max=30"`
	Description string    `json:"description" db:"description" validate:"required,min=3,max=80"`
	Key         string    `json:"key" db:"key" validate:"required,min=3,max=30"`
}
