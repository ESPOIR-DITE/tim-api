package video_comment

import (
	"errors"
	"net/http"
)

type VideoComment struct {
	Id      string `json:"id" gorm:"primaryKey"`
	VideoId string `json:"videoId"`
	UserId  string `json:"userId"`
	Comment []byte `json:"picture"`
}

func (v VideoComment) Bind(r *http.Request) error {
	if v.VideoId == "" && len(v.Comment) < 0 {
		return errors.New("missing required fields")
	}
	return nil
}
