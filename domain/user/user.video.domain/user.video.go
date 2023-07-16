package userVideoDomain

import (
	"errors"
	"net/http"
	"time"
)

type UserVideo struct {
	Id        string    `json:"id" gorm:"primaryKey"`
	AccountId string    `json:"accountId" sql:"account_id"`
	VideoId   string    `json:"videoId" sql:"video_id"`
	Date      time.Time `json:"date"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u UserVideo) Bind(r *http.Request) error {
	if u.AccountId == "" && u.VideoId == "" {
		return errors.New("missing required fields")
	}
	return nil
}
func (UserVideo) TableName() string {
	return "user_video"
}
