package user_video

import (
	"errors"
	"net/http"
)

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
