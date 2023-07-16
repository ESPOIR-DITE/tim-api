package userVideoRepository

import (
	"github.com/ESPOIR-DITE/tim-api/config"
	user_video "github.com/ESPOIR-DITE/tim-api/domain/user/user.video.domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserVideoRepository struct {
	GormDb *gorm.DB
}

func NewUserVideoRepository(gormDb *gorm.DB) *UserVideoRepository {
	return &UserVideoRepository{
		GormDb: gormDb,
	}
}

func (uvr UserVideoRepository) CreateUserVideo(entity user_video.UserVideo) (*user_video.UserVideo, error) {
	var tableData = &user_video.UserVideo{}
	id := "UV-" + uuid.New().String()
	entity.Id = id
	if err := uvr.GormDb.Create(entity).First(&tableData).Error; err != nil {
		return nil, err
	}
	return tableData, nil
}
func (uvr UserVideoRepository) UpdateUserVideo(entity user_video.UserVideo) (*user_video.UserVideo, error) {
	var tableData = &user_video.UserVideo{}
	if err := config.GetDatabase().Create(entity).First(&tableData).Error; err != nil {
		return nil, err
	}
	return tableData, nil
}
func (uvr UserVideoRepository) GetAllUserVideo(accountId string) ([]user_video.UserVideo, error) {
	entity := []user_video.UserVideo{}
	if err := uvr.GormDb.Where("account_id = ?", accountId).Find(&entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}
func (uvr UserVideoRepository) GetUserVideo(id string) (*user_video.UserVideo, error) {
	entity := &user_video.UserVideo{}
	if err := uvr.GormDb.Where("id = ?", id).Find(entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}
func (uvr UserVideoRepository) GetUserVideosWithUserId(id string) ([]user_video.UserVideo, error) {
	entity := []user_video.UserVideo{}
	if err := uvr.GormDb.Where("account_id = ?", id).Find(entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}
func (uvr UserVideoRepository) GetUserVideoWithVideoId(videoId string) (*user_video.UserVideo, error) {
	entity := &user_video.UserVideo{}
	if err := uvr.GormDb.Where("video_id = ?", videoId).Find(entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}
func (uvr UserVideoRepository) GetUserVideos() ([]user_video.UserVideo, error) {
	var entity []user_video.UserVideo
	if err := uvr.GormDb.Find(&entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}
func (uvr UserVideoRepository) DeleteUserVideo(id string) (bool, error) {
	entity := user_video.UserVideo{}
	if err := uvr.GormDb.Where("id = ?", id).Delete(&entity).Error; err != nil {
		return false, err
	}
	if entity.AccountId == "" {
		return true, nil
	}
	return false, nil
}
