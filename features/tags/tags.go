package tags

type Tag struct {
	ID   string `json:"id" db:"id"`
	Name string `json:"name" db:"name" validate:"required,min=3,max=30"`
	Key  string `json:"key" db:"key" validate:"required,min=3,max=30"`
}
