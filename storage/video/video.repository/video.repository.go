package videoRepository

import (
	"github.com/ESPOIR-DITE/tim-api/domain/video/video"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VideoRepository struct {
	GormDB *gorm.DB
}

func NewVideoRepository(gormDB *gorm.DB) *VideoRepository {
	return &VideoRepository{
		GormDB: gormDB,
	}
}

func (vr VideoRepository) CreateVideo(entity video.Video) (*video.Video, error) {
	id := "V-" + uuid.New().String()
	entity.Id = id
	if err := vr.GormDB.Create(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (vr VideoRepository) UpdateVideo(entity video.Video) (*video.Video, error) {
	var tableData = &video.Video{}
	if err := vr.GormDB.Updates(entity).Find(&tableData).Error; err != nil {
		return nil, err
	}
	return tableData, nil
}

func (vr VideoRepository) GetVideo(id string) (*video.Video, error) {
	entity := &video.Video{}
	if err := vr.GormDB.Where("id = ?", id).Find(entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func (vr VideoRepository) GetAllPublicVideo() ([]video.Video, error) {
	entity := []video.Video{}
	if err := vr.GormDB.Where("is_private = ?", false).Find(&entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}
func (vr VideoRepository) GetAllPrivateVideo() ([]video.Video, error) {
	entity := []video.Video{}
	if err := vr.GormDB.Where("is_private = ?", true).Find(&entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}
func (vr VideoRepository) GetVideos() ([]video.Video, error) {
	entity := []video.Video{}
	if err := vr.GormDB.Find(&entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}
func (vr VideoRepository) DeleteVideo(email string) (bool, error) {
	entity := video.Video{}
	if err := vr.GormDB.Where("id = ?", email).Delete(&entity).Error; err != nil {
		return false, err
	}
	return false, nil
}
func (vr VideoRepository) CountVideo() (*int64, error) {
	var value int64
	if err := vr.GormDB.Table("videos").Count(&value).Error; err != nil {
		return nil, err
	}
	return &value, nil
}
