package videoDataDomain

import (
	"errors"
	"net/http"
	"time"
)

// VideoData represents the main entity of this application.
//
// swagger:model
type VideoData struct {
	Id        string `json:"id" gorm:"primaryKey"`
	Picture   []byte `json:"picture" gorm:"-:all"`
	Video     []byte `json:"video" gorm:"-:all"`
	FileType  string `json:"fileType" `
	FileSize  string `json:"fileSize"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (v VideoData) Bind(r *http.Request) error {
	if v.Id == "" && len(v.Picture) != 0 && len(v.Video) != 0 {
		return errors.New("missing required fields")
	}
	return nil
}

func (VideoData) TableName() string {
	return "video_data"
}
