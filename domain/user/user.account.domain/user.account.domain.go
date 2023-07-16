package userAccountDomain

import (
	"errors"
	"net/http"
	"time"
)

type UserAccount struct {
	Id           string `json:"id" sql:"customer_id" gorm:"primaryKey"`
	AccountId    string `json:"accountId" sql:"account_id"`
	UserDetailId string `json:"userDetailId" sql:"user_detail_id"`
	UserId       string `json:"userId"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (u UserAccount) Bind(r *http.Request) error {
	if u.AccountId == "" && u.UserId == "" && u.UserDetailId == "" {
		return errors.New("missing required fields")
	}
	return nil
}

func (UserAccount) TableName() string {
	return "user_account"
}
