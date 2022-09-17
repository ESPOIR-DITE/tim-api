package security

import (
	"errors"
	"net/http"
)

// SystemDate represent data that the system needs to function properly.
//
// Whenever a record is created the identifier should be noted down.
//
// swagger:model
type SystemData struct {
	// This value should be unique
	// required: true
	// min: 1
	Identifier string `json:"identifier" gorm:"primaryKey"`
	// required: true
	Value string `json:"value"`
}

func (s SystemData) Bind(r *http.Request) error {
	if s.Value == "" {
		return errors.New("Id field missing.")
	}
	return nil
}
