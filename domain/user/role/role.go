package role

import (
	"errors"
	"net/http"
)

type Role struct {
	Id          string `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (role Role) Bind(r *http.Request) error {
	if role.Name == "" {
		return errors.New("missing required fields")
	}
	return nil
}
