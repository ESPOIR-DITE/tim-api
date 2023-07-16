package category

import (
	"errors"
	"net/http"
	"time"
)

type Category struct {
	Id          string `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (c Category) Bind(r *http.Request) error {
	if c.Name == "" {
		return errors.New("missing required fields")
	}
	return nil
}
func (Category) TableName() string {
	return "category"
}
