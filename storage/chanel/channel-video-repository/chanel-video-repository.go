package role_repo

import (
	"fmt"
	"github.com/google/uuid"
	"tim-api/config"
	"tim-api/domain/channel/channel-video"
)

var connection = config.GetDatabase()

func CreateChannelVideoTable() bool {
	var tableData = &channel_video.ChannelVideos{}
	err := connection.AutoMigrate(tableData)
	if err != nil {
		return false
	}
	return true
}
func SetChannelDatabase() {
	err := connection.Set("gorm:table_options", "ENGINE=Distributed(cluster, default, hits)").AutoMigrate(&channel_video.ChannelVideos{})
	if err != nil {
		fmt.Println("Role database config not set")
	} else {
		fmt.Println("Role database config set successfully")
	}
}
func CreateChannelVideo(channel channel_video.ChannelVideos) channel_video.ChannelVideos {
	var tableData = channel_video.ChannelVideos{}
	id := "CV-" + uuid.New().String()
	user := channel_video.ChannelVideos{id, channel.VideoId, channel.ChannelId, channel.Description}
	connection.Create(user).Find(&tableData)
	return tableData
}
func UpdateChannelVideo(entity channel_video.ChannelVideos) channel_video.ChannelVideos {
	var tableData = channel_video.ChannelVideos{}
	//user := domain.Role{entity.Id, entity.Name, entity.Description}
	connection.Where("id = ", entity.Id).Updates(entity).Find(&tableData)
	return tableData
}
func GetChannelVideo(roleId string) channel_video.ChannelVideos {
	entity := channel_video.ChannelVideos{}
	connection.Where("id = ?", roleId).Find(&entity)
	return entity
}
func GetChannelVideosByChannelId(channelId string) []channel_video.ChannelVideos {
	entity := []channel_video.ChannelVideos{}
	connection.Where("channel_id = ?", channelId).Find(&entity)
	return entity
}
func GetChannelVideosByVideoId(videoId string) channel_video.ChannelVideos {
	entity := channel_video.ChannelVideos{}
	config.GetDatabase().Where("video_id = ?", videoId).Find(&entity)
	return entity
}
func GetChannelVideos() []channel_video.ChannelVideos {
	var entity []channel_video.ChannelVideos
	connection.Find(&entity)
	return entity
}
func DeleteChannelVideo(id string) bool {
	entity := channel_video.ChannelVideos{}
	connection.Where("id = ?", id).Delete(&entity)
	if entity.Id == "" {
		return true
	}
	return false
}
func CountChannelVideo() int64 {
	var value int64
	connection.Count(&value).Find(channel_video.ChannelVideos{})
	return value
}
