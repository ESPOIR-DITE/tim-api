package channel_video

import (
	"errors"
	"net/http"
)

type ChannelVideos struct {
	VideoId     string `json:"video_id" gorm:"primaryKey"`
	ChannelId   string `json:"channel_id" gorm:"primaryKey"`
	Description string `json:"description"`
}

func (c ChannelVideos) Bind(r *http.Request) error {
	if c.ChannelId == "" && c.VideoId == "" {
		return errors.New("missing required fields")
	}
	return nil
}

func (ChannelVideos) TableName() string {
	return "channel_video"
}
