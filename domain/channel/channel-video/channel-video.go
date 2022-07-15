package channel_video

import (
	"errors"
	"net/http"
)

type ChannelVideos struct {
	Id          string `json:"id" gorm:"primaryKey"`
	VideoId     string `json:"video_id"`
	ChannelId   string `json:"channel_id"`
	Description string `json:"description"`
}

func (c ChannelVideos) Bind(r *http.Request) error {
	if c.ChannelId == "" && c.VideoId == "" {
		return errors.New("missing required fields")
	}
	return nil
}
