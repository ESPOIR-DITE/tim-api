package video

import (
	"errors"
	"net/http"
)

// Video represents the main entity of this application.
//
// swagger:model Video
type Video struct {
	// The id for this video
	//
	//required: true
	//min: 1
	Id           string  `json:"id" gorm:"primaryKey"`
	Title        string  `json:"title"`
	Date         string  `json:"date"`
	DateUploaded string  `json:"dateUploaded"`
	Description  string  `json:"description"`
	IsPrivate    bool    `json:"isPrivate"`
	Price        float64 `json:"price"`
	URL          string  `json:"url"`
}

func (v Video) Bind(r *http.Request) error {
	if v.Title == "" && v.Description == "" {
		return errors.New("missing required fields")
	}
	return nil
}
