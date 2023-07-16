package videoReactionDomain

import (
	"errors"
	"net/http"
	"time"
)

// VideoReaction represent reactions of user.home.controller.domain.controller toward a video.controller
//
// this is the main attractive model that involve reactions of a user.home.controller.domain.controller and get them involved in the application.
//
// swagger:model
type VideoReaction struct {
	// Takes the video.controller's id reacted to as it Id.
	// Should exist in Video table
	// required: true
	// min: 1
	VideoId string `json:"id" gorm:"primaryKey" sql:"video_id"`
	UserId  string `json:"userId" gorm:"primaryKey" sql:"user_id"`
	// required: false
	Like int `json:"like"`
	// required: false
	Unlike    int `json:"unlike"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (v VideoReaction) Bind(r *http.Request) error {
	if v.VideoId == "" {
		return errors.New("missing required fields")
	}
	return nil
}

func (VideoReaction) TableName() string {
	return "video_reaction"
}
