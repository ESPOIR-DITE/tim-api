package domain

import (
	"errors"
	"net/http"
)

type UserStack struct {
	SuperAdmin int64
	Admin      int64
	Agent      int64
}

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

// Role represents the attribute that a user has in this application.
//
// swagger:model
type Role struct {
	// The id for this user
	//
	//required: true
	//min: 1
	Id string `json:"id" gorm:"primaryKey"`
	// The name of the role
	// required: true
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (role Role) Bind(r *http.Request) error {
	if role.Name == "" {
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
