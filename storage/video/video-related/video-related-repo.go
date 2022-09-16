package video_related

import (
	"fmt"
	"tim-api/config"
	video_related "tim-api/domain/video/video-related"
)

var connection = config.GetDatabase()

func CreateVideoRelatedTable() bool {
	var tableData = &video_related.VideoRelated{}
	err := connection.AutoMigrate(tableData)
	if err != nil {
		return false
	}
	return true
}
func CreateVideoRelated(entity video_related.VideoRelated) (video_related.VideoRelated, error) {
	var tableData = video_related.VideoRelated{}
	err := connection.Create(entity).Find(&tableData).Error
	if err != nil {
		fmt.Println("error creating video related")
		return tableData, err
	}
	return tableData, nil
}
func GetVideosRelatedTo(videoId string) ([]video_related.VideoRelated, error) {
	entity := []video_related.VideoRelated{}
	err := connection.Where("current_video_id = ?", videoId).Or("related_video_id = ?", videoId).Find(&entity).Error
	if err != nil {
		fmt.Println("error reading video related")
		return entity, err
	}
	return entity, nil
}
func DeleteVideoRelated(videoRelated video_related.VideoRelated) (bool, error) {
	err := connection.Where("current_video_id = ? AND related_video_id = ?", videoRelated.CurrentVideoId, videoRelated.RelatedVideoId).Delete(&video_related.VideoRelated{}).Error
	if err != nil {
		fmt.Println("error deleting video related")
		return false, err
	}
	return true, nil
}
