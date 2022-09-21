package video_data

import (
	"errors"
	"net/http"
)

// VideoData represents the main entity of this application.
//
// swagger:model
type VideoData struct {
	Id       string `json:"id" gorm:"primaryKey"`
	Picture  []byte `json:"picture"`
	Video    []byte `json:"video" gorm:"-:all"`
	FileType string `json:"fileType"`
	FileSize string `json:"fileSize"`
}

func (v VideoData) Bind(r *http.Request) error {
	if v.Id == "" && len(v.Picture) != 0 && len(v.Video) != 0 {
		return errors.New("missing required fields")
	}
	return nil
}
