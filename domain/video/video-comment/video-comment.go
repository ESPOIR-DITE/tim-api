package video_comment

import (
	"errors"
	"net/http"
)

type VideoComment struct {
	Id        string `json:"id" gorm:"primaryKey"`
	VideoId   string `json:"videoId"`
	CommentId string `json:"categoryId"`
}

func (v VideoComment) Bind(r *http.Request) error {
	if v.VideoId == "" && v.CommentId == "" {
		return errors.New("missing required fields")
	}
	return nil
}
