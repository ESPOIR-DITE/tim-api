package videoCommentDomain

import (
	"errors"
	"net/http"
	"time"
)

type VideoComment struct {
	Id        string `json:"id" gorm:"primaryKey"`
	VideoId   string `json:"videoId" sql:"video_id"`
	UserId    string `json:"userId" sql:"user_id"`
	Comment   []byte `json:"picture"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (v VideoComment) Bind(r *http.Request) error {
	if v.VideoId == "" && len(v.Comment) < 0 {
		return errors.New("missing required fields")
	}
	return nil
}

func (VideoComment) TableName() string {
	return "video_comment"
}
