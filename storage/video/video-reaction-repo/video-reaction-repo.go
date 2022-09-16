package video_reaction_repo

import (
	"log"
	"tim-api/config"
	video_reaction "tim-api/domain/video/video-reaction"
)

var connection = config.GetDatabase()

func CreateVideoReactionTable() bool {
	var tableData = &video_reaction.VideoReaction{}
	err := connection.AutoMigrate(tableData)
	if err != nil {
		return false
	}
	return true
}

func CreateVideoReaction(entity video_reaction.VideoReaction) video_reaction.VideoReaction {
	var tableData = video_reaction.VideoReaction{}
	connection.Create(entity).Find(&tableData)
	return tableData
}

func LikeReact(entity video_reaction.VideoReaction) video_reaction.VideoReaction {
	isExist := IsVideoReactionExist(entity.VideoId)
	if isExist {
		return liked(entity.VideoId)
	}
	var result = CreateVideoReaction(entity)
	if result.VideoId != "" {
		return liked(result.VideoId)
	}
	return result
}
func UnLikeReact(entity video_reaction.VideoReaction) video_reaction.VideoReaction {
	isExist := IsVideoReactionExist(entity.VideoId)
	if isExist {
		return unLiked(entity.VideoId)
	}
	var result = CreateVideoReaction(entity)
	if result.VideoId != "" {
		return unLiked(result.VideoId)
	}
	return result
}
func liked(videoId string) video_reaction.VideoReaction {
	videoReaction := GetVideo(videoId)
	videoReaction.Like++
	return UpdateVideoReaction(videoReaction)
}
func unLiked(videoId string) video_reaction.VideoReaction {
	videoReaction := GetVideo(videoId)
	videoReaction.UnLike++
	return UpdateVideoReaction(videoReaction)
}

func GetVideo(id string) video_reaction.VideoReaction {
	entity := video_reaction.VideoReaction{}
	connection.Where("video_id = ?", id).Find(&entity)
	return entity
}

func UpdateVideoReaction(entity video_reaction.VideoReaction) video_reaction.VideoReaction {
	var tableData = video_reaction.VideoReaction{}
	err := connection.Updates(entity).Find(&tableData).Error
	if err != nil {
		log.Fatal(err)
		return tableData
	}
	return tableData
}

func IsVideoReactionExist(videoId string) bool {
	var exists bool
	err := connection.Model(&video_reaction.VideoReaction{}).Select("count(*) > 0").Where("video_id = ?", videoId).Find(&exists).Error
	if err != nil {
		log.Fatal(err)
		return false
	}
	return exists
}

func DeleteVideoReaction(id string) bool {
	err := connection.Where("video_id = ?", id).Delete(&video_reaction.VideoReaction{}).Error
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
