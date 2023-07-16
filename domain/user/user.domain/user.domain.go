package userDomain

import (
	"errors"
	"net/http"
	"time"
)

type User struct {
	Email     string `json:"email" gorm:"primaryKey"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	BirthDate string `json:"birthDate" sql:"birth_date"`
	RoleId    string `json:"roleId" sql:"role_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u User) Bind(r *http.Request) error {
	if u.Name == "" && u.RoleId == "" && u.Email == "" {
		return errors.New("missing required fields")
	}
	return nil
}

func (User) TableName() string {
	return "user"
}
