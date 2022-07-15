package channel_subscription

import (
	"errors"
	"net/http"
)

type ChannelSubscription struct {
	Id        string `json:"id" gorm:"primaryKey"`
	ChannelId string `json:"channel_id"`
	UserId    string `json:"user_id"`
	Date      string `json:"date"`
}

func (c ChannelSubscription) Bind(r *http.Request) error {
	if c.ChannelId == "" || c.UserId == "" {
		return errors.New("missing required fields")
	}
	return nil
}
