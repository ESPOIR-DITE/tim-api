package channel

import (
	"errors"
	"net/http"
)

type Channel struct {
	Id            string `json:"id" gorm:"primaryKey"`
	Name          string `json:"name" gorm:"not null"`
	ChannelTypeId string `json:"channel_type_id"`
	UserId        string `json:"user_id"`
	Region        string `json:"region"`
	Date          string `json:"date"`
	Image         []byte `json:"image"`
	ImageBase64   string `json:"image_base_64" gorm:"-:all"`
	Description   string `json:"description"`
}

func (c Channel) Bind(r *http.Request) error {
	if c.UserId == "" || c.Name == "" || c.ChannelTypeId == "" {
		return errors.New("missing required fields")
	}
	return nil
}
