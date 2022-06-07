package domain

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

type UserSubscription struct {
	Id             string `json:"id" gorm:"primaryKey"`
	UserId         string `json:"userId"`
	Stat           string `json:"stat"`
	SubscriptionId string `json:"subscriptionId"`
	Date           string `json:"date"`
}

func (u UserSubscription) Bind(r *http.Request) error {
	if u.UserId == "" && u.SubscriptionId == "" {
		return errors.New("missing required fields")
	}
	return nil
}

type UserVideo struct {
	Id         string `json:"id" gorm:"primaryKey"`
	CustomerId string `json:"customerId"`
	VideoId    string `json:"videoId"`
	Date       string `json:"date"`
}

func (u UserVideo) Bind(r *http.Request) error {
	if u.CustomerId == "" && u.VideoId == "" {
		return errors.New("missing required fields")
	}
	return nil
}
