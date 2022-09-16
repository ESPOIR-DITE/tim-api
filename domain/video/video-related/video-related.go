package video_related

import (
	"errors"
	"net/http"
)

// VideoRelated represent a relationship between a video to it close related one
//
// when one video is being played the rest related will be loaded to the relate video panel
//
// swagger:model
type VideoRelated struct {
	// Takes video id
	// Should exist in Video table
	// required: true
	// min: 1
	CurrentVideoId string `json:"currentVideo" gorm:"primary_key"`
	// Takes video id
	// Should exist in Video table
	// required: true
	// min: 1
	RelatedVideoId string `json:"relatedVideoId" gorm:"primary_key"`
}

func (v VideoRelated) Bind(r *http.Request) error {
	if v.RelatedVideoId == "" && v.CurrentVideoId == "" {
		return errors.New("missing required fields")
	}
	return nil
}
