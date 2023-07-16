package userSubscriptionDomain

import (
	"errors"
	"net/http"
	"time"
)

type UserSubscription struct {
	Id             string `json:"id" gorm:"primaryKey"`
	AccountId      string `json:"accountId" sql:"account_id"`
	Stat           string `json:"stat"`
	SubscriptionId string `json:"subscriptionId" sql:"subscription_id"`
	Date           string `json:"date"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (u UserSubscription) Bind(r *http.Request) error {
	if u.AccountId == "" && u.SubscriptionId == "" {
		return errors.New("missing required fields")
	}
	return nil
}
func (UserSubscription) TableName() string {
	return "user_subscription"
}
