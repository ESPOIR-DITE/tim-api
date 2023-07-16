package channel_type

import (
	"errors"
	"net/http"
)

type ChannelType struct {
	Id          string `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (c ChannelType) Bind(r *http.Request) error {
	if c.Name == "" {
		return errors.New("missing required fields")
	}
	return nil
}

func (ChannelType) TableName() string {
	return "channel_type"
}
