package video_reaction

import (
	"errors"
	"net/http"
)

// VideoReaction represent reactions of user toward a video
//
// this is the main attractive model that involve reactions of a user and get them involved in the application.
//
// swagger:model
type VideoReaction struct {
	// Takes the video's id reacted to as it Id.
	// Should exist in Video table
	// required: true
	// min: 1
	VideoId string `json:"id" gorm:"primaryKey"`
	// required: false
	Like int `json:"like"`
	// required: false
	UnLike int `json:"unLike"`
}

func (v VideoReaction) Bind(r *http.Request) error {
	if v.VideoId == "" {
		return errors.New("missing required fields")
	}
	return nil
}
