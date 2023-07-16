package videodataRepository

import (
	videoDataDomain "github.com/ESPOIR-DITE/tim-api/domain/video/video.data.domain"
	"gorm.io/gorm"
)

type VideoDataRepository struct {
	GormDB *gorm.DB
}

func NewVideoDataRepository(gormDB *gorm.DB) *VideoDataRepository {
	return &VideoDataRepository{
		GormDB: gormDB,
	}
}

func (vdr VideoDataRepository) CreateVideoData(entity videoDataDomain.VideoData) (*videoDataDomain.VideoData, error) {
	if err := vdr.GormDB.Create(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (vdr VideoDataRepository) UpdateVideoDate(entity videoDataDomain.VideoData) (*videoDataDomain.VideoData, error) {
	var tableData = &videoDataDomain.VideoData{}
	if err := vdr.GormDB.Updates(entity).Find(&tableData).Error; err != nil {
		return nil, err
	}
	return tableData, nil
}
func (vdr VideoDataRepository) GetVideoDate(id string) (*videoDataDomain.VideoData, error) {
	entity := &videoDataDomain.VideoData{}
	if err := vdr.GormDB.Where("id = ?", id).Find(&entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}
func (vdr VideoDataRepository) GetVideoDatas() ([]videoDataDomain.VideoData, error) {
	entity := []videoDataDomain.VideoData{}
	if err := vdr.GormDB.Find(&entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}
func (vdr VideoDataRepository) DeleteVideoData(email string) (bool, error) {
	entity := videoDataDomain.VideoData{}
	if err := vdr.GormDB.Where("id = ?", email).Delete(&entity).Error; err != nil {
		return false, err
	}
	if entity.Id == "" {
		return true, nil
	}
	return false, nil
}
