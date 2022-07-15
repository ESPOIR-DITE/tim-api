package domain

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
