package category

import (
	"errors"
	"net/http"
)

type Category struct {
	Id          string `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (c Category) Bind(r *http.Request) error {
	if c.Name == "" {
		return errors.New("missing required fields")
	}
	return nil
}
