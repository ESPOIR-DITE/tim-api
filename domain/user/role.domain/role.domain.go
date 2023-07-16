package roleDomain

import (
	"errors"
	_ "gorm.io/datatypes"
	"net/http"
	"time"
)

type Role struct {
	Id          string `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (role Role) Bind(r *http.Request) error {
	if role.Name == "" {
		return errors.New("missing required fields")
	}
	return nil
}

func (Role) TableName() string {
	return "role"
}
