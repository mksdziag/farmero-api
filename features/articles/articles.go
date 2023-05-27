package articles

import (
	UUID "github.com/google/uuid"
)

type Article struct {
	ID          UUID.UUID `json:"id" db:"id"`
	Title       string    `json:"title" db:"title" validate:"required,min=3,max=80"`
	Description string    `json:"description" db:"description" validate:"required,min=3,max=80"`
	Content     string    `json:"content" db:"content" validate:"required,min=3"`
	Cover       string    `json:"cover" db:"cover" validate:"required"`
	Tags        []string  `json:"tags"`
	Categories  []string  `json:"categories"`
}
