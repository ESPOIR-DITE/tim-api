package channel

import (
	"errors"
	"net/http"
	"time"
)

type Channel struct {
	Id            string `json:"id" gorm:"primaryKey"`
	Name          string `json:"name" gorm:"not null"`
	ChannelTypeId string `json:"channel_type_id" sql:"channel_type_id"`
	UserId        string `json:"account_id" sql:"account_id"`
	Region        string `json:"region"`
	Date          string `json:"date"`
	Image         []byte `json:"image" gorm:"-:all"`
	ImageBase64   string `json:"image_base_64" gorm:"-:all"`
	Description   string `json:"description"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (c Channel) Bind(r *http.Request) error {
	if c.UserId == "" || c.Name == "" || c.ChannelTypeId == "" {
		return errors.New("missing required fields")
	}
	return nil
}

func (Channel) TableName() string {
	return "channel.controller"
}
