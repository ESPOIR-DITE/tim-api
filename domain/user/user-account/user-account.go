package user_account

import (
	"errors"
	"net/http"
)

type UserAccount struct {
	CustomerId string `json:"customerId" gorm:"primaryKey"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Date       string `json:"date"`
	Status     bool   `json:"status"`
}

func (u UserAccount) Bind(r *http.Request) error {
	if u.Email == "" && u.Password == "" {
		return errors.New("missing required fields")
	}
	return nil
}
