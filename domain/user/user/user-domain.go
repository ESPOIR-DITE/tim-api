package user

import (
	"errors"
	"net/http"
)

type User struct {
	Email     string `json:"email" gorm:"primaryKey"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	BirthDate string `json:"birthDate"`
	RoleId    string `json:"roleId"`
}

func (u User) Bind(r *http.Request) error {
	if u.Name == "" && u.RoleId == "" && u.Email == "" {
		return errors.New("missing required fields")
	}
	return nil
}
