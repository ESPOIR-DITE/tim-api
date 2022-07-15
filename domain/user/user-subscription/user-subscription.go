package user_subscription

import (
	"errors"
	"net/http"
)

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
