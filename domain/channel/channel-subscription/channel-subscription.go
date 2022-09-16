package channel_subscription

import (
	"errors"
	"net/http"
)

type ChannelSubscription struct {
	ChannelId string `json:"channel_id" gorm:"primaryKey"`
	UserId    string `json:"user_id" gorm:"primaryKey"`
	Date      string `json:"date"`
}

func (c ChannelSubscription) Bind(r *http.Request) error {
	if c.ChannelId == "" || c.UserId == "" {
		return errors.New("missing required fields")
	}
	return nil
}
