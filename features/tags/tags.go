package tags

import "github.com/google/uuid"

type Tag struct {
	ID   uuid.UUID `json:"id" db:"id"`
	Name string    `json:"name" db:"name" validate:"required,min=3,max=30"`
	Key  string    `json:"key" db:"key" validate:"required,min=3,max=30"`
}
