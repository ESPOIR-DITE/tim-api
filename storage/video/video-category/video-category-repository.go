package video_category

import (
	videoCategoryDomain "github.com/ESPOIR-DITE/tim-api/domain/video/video.category.domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VideoCategoryRepository struct {
	GormDB *gorm.DB
}

func NewVideoCategoryRepository(gormDB *gorm.DB) *VideoCategoryRepository {
	return &VideoCategoryRepository{
		GormDB: gormDB,
	}
}

func (vcr VideoCategoryRepository) CreateVideoCategory(entity videoCategoryDomain.VideoCategory) (*videoCategoryDomain.VideoCategory, error) {
	var tableData = &videoCategoryDomain.VideoCategory{}
	id := "VC-" + uuid.New().String()
	entity.Id = id
	if err := vcr.GormDB.Create(entity).First(tableData).Error; err != nil {
		return nil, err
	}
	return tableData, nil
}

func (vcr VideoCategoryRepository) UpdateVideoCategory(entity videoCategoryDomain.VideoCategory) (*videoCategoryDomain.VideoCategory, error) {
	var tableData = &videoCategoryDomain.VideoCategory{}
	if err := vcr.GormDB.Create(entity).First(&tableData).Error; err != nil {
		return nil, err
	}
	return tableData, nil
}

func (vcr VideoCategoryRepository) GetVideoCategory(id string) (*videoCategoryDomain.VideoCategory, error) {
	entity := &videoCategoryDomain.VideoCategory{}
	if err := vcr.GormDB.Where("id = ?", id).First(entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func (vcr VideoCategoryRepository) GetVideoCategories() ([]videoCategoryDomain.VideoCategory, error) {
	entity := []videoCategoryDomain.VideoCategory{}
	if err := vcr.GormDB.Find(&entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}
func (vcr VideoCategoryRepository) DeleteVideoCategory(email string) (bool, error) {
	entity := &videoCategoryDomain.VideoCategory{}
	if err := vcr.GormDB.Where("id = ?", email).Delete(&entity).Error; err != nil {
		return false, err
	}
	if entity.Id == "" {
		return true, nil
	}
	return false, nil
}
