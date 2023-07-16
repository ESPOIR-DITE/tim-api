package videoCategoryDomain

import (
	"errors"
	"net/http"
	"time"
)

type VideoCategory struct {
	Id         string `json:"id" gorm:"primaryKey"`
	VideoId    string `json:"videoId" sql:"video_id"`
	CategoryId string `json:"categoryId" sql:"category_id"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (v VideoCategory) Bind(r *http.Request) error {
	if v.VideoId == "" && v.CategoryId == "" {
		return errors.New("missing required fields")
	}
	return nil
}

func (VideoCategory) TableName() string {
	return "video_category"
}
