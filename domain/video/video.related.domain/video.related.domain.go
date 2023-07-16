package videoRelatedDomain

import (
	"errors"
	"net/http"
	"time"
)

// VideoRelated represent a relationship between a video.controller to it close related one
//
// when one video.controller is being played the rest related will be loaded to the relate video.controller panel
//
// swagger:model
type VideoRelated struct {
	// Takes video.controller id
	// Should exist in Video table
	// required: true
	// min: 1
	CurrentVideoId string `json:"currentVideo" gorm:"primary_key"`
	// Takes video.controller id
	// Should exist in Video table
	// required: true
	// min: 1
	RelatedVideoId string `json:"relatedVideoId" gorm:"primary_key"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (v VideoRelated) Bind(r *http.Request) error {
	if v.RelatedVideoId == "" && v.CurrentVideoId == "" {
		return errors.New("missing required fields")
	}
	return nil
}

func (VideoRelated) TableName() string {
	return "video_related"
}
