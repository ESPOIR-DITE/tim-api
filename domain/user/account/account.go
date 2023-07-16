package account

import (
	"errors"
	"net/http"
	"time"
)

type Account struct {
	Id        string `json:"id" sql:"id" gorm:"primaryKey"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Date      string `json:"date"`
	Status    bool   `json:"status"`
	Token     string `json:"token"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u Account) Bind(r *http.Request) error {
	if u.Email == "" && u.Password == "" {
		return errors.New("missing required fields")
	}
	return nil
}

func (Account) TableName() string {
	return "account"
}
