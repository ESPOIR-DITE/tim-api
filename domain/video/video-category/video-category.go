package video_category

import (
	"errors"
	"net/http"
)

type VideoCategory struct {
	Id         string `json:"id" gorm:"primaryKey"`
	VideoId    string `json:"videoId"`
	CategoryId string `json:"categoryId"`
}

func (v VideoCategory) Bind(r *http.Request) error {
	if v.VideoId == "" && v.CategoryId == "" {
		return errors.New("missing required fields")
	}
	return nil
}
