package channelVideoRepository

import (
	"github.com/ESPOIR-DITE/tim-api/config"
	channel_video "github.com/ESPOIR-DITE/tim-api/domain/channel/channel-video"
	"gorm.io/gorm"
)

type ChannelVideoRepository struct {
	GormDB *gorm.DB
}

func NewChannelVideoRepository(gormDb *gorm.DB) *ChannelVideoRepository {
	return &ChannelVideoRepository{
		GormDB: gormDb,
	}
}

func (cv *ChannelVideoRepository) CreateChannelVideo(channel channel_video.ChannelVideos) (*channel_video.ChannelVideos, error) {
	tableData := &channel_video.ChannelVideos{}
	err := cv.GormDB.Create(channel).First(&tableData).Error
	if err != nil {
		return tableData, err
	}
	return tableData, nil
}
func (cv *ChannelVideoRepository) UpdateChannelVideo(entity channel_video.ChannelVideos) (*channel_video.ChannelVideos, error) {
	tableData := &channel_video.ChannelVideos{}
	err := cv.GormDB.Where("channel_id = ? And video_id = ?", entity.ChannelId, entity.VideoId).Updates(entity).First(&tableData).Error
	if err != nil {
		return tableData, err
	}
	return tableData, nil
}
func (cv *ChannelVideoRepository) GetChannelVideo(roleId string) (*channel_video.ChannelVideos, error) {
	entity := &channel_video.ChannelVideos{}
	if err := cv.GormDB.Where("id = ?", roleId).First(entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}
func (cv *ChannelVideoRepository) GetChannelVideosByChannelId(channelId string) ([]channel_video.ChannelVideos, error) {
	entity := []channel_video.ChannelVideos{}
	if err := cv.GormDB.Where("channel_id = ?", channelId).First(&entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func GetChannelVideosByVideoId(videoId string) (*channel_video.ChannelVideos, error) {
	entity := &channel_video.ChannelVideos{}
	if err := config.GetDatabase().Where("video_id = ?", videoId).Find(&entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func (cv *ChannelVideoRepository) GetChannelVideos() ([]channel_video.ChannelVideos, error) {
	var entity []channel_video.ChannelVideos
	if err := cv.GormDB.First(&entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func (cv *ChannelVideoRepository) DeleteChannelVideo(id string) (bool, error) {
	entity := channel_video.ChannelVideos{}
	err := cv.GormDB.Where("id = ?", id).Delete(&entity).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (cv *ChannelVideoRepository) CountChannelVideo() (*int64, error) {
	var value int64
	if err := cv.GormDB.Count(&value).Find(channel_video.ChannelVideos{}).Error; err != nil {
		return nil, err
	}
	return &value, nil
}
