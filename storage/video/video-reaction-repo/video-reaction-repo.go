package videoReactionRepository

import (
	videoReactionDomain "github.com/ESPOIR-DITE/tim-api/domain/video/video.reaction"
	"gorm.io/gorm"
)

type VideoReactionRepository struct {
	GormDB *gorm.DB
}

func NewVideoReactionRepository(gormDB *gorm.DB) *VideoReactionRepository {
	return &VideoReactionRepository{
		GormDB: gormDB,
	}
}

func (vrr VideoReactionRepository) CreateVideoReaction(entity videoReactionDomain.VideoReaction) (*videoReactionDomain.VideoReaction, error) {
	var tableData = &videoReactionDomain.VideoReaction{}
	if err := vrr.GormDB.Create(entity).Find(&tableData).Error; err != nil {
		return nil, err
	}
	return tableData, nil
}

func (vrr VideoReactionRepository) LikeReact(entity videoReactionDomain.VideoReaction) (*videoReactionDomain.VideoReaction, error) {
	isExist, _ := vrr.IsVideoReactionExist(entity.VideoId)
	if isExist {
		return vrr.liked(entity.VideoId)
	}
	result, err := vrr.CreateVideoReaction(entity)
	if err != nil {
		return nil, err
	}
	if result.VideoId != "" {
		return vrr.liked(result.VideoId)
	}
	return result, nil
}
func (vrr VideoReactionRepository) UnLikeReact(entity videoReactionDomain.VideoReaction) (*videoReactionDomain.VideoReaction, error) {
	isExist, _ := vrr.IsVideoReactionExist(entity.VideoId)
	if isExist {
		return vrr.unLiked(entity.VideoId)
	}
	result, err := vrr.CreateVideoReaction(entity)
	if err != nil {
		return nil, err
	}
	return vrr.unLiked(result.VideoId)
}
func (vrr VideoReactionRepository) liked(videoId string) (*videoReactionDomain.VideoReaction, error) {
	videoReaction, err := vrr.GetVideo(videoId)
	if err != nil {
		return nil, err
	}
	videoReaction.Like++
	return vrr.UpdateVideoReaction(*videoReaction)
}
func (vrr VideoReactionRepository) unLiked(videoId string) (*videoReactionDomain.VideoReaction, error) {
	videoReaction, err := vrr.GetVideo(videoId)
	if err != nil {
		return nil, err
	}
	videoReaction.Unlike++
	return vrr.UpdateVideoReaction(*videoReaction)
}

func (vrr VideoReactionRepository) GetVideo(id string) (*videoReactionDomain.VideoReaction, error) {
	entity := &videoReactionDomain.VideoReaction{}
	if err := vrr.GormDB.Where("video_id = ?", id).Find(&entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func (vrr VideoReactionRepository) UpdateVideoReaction(entity videoReactionDomain.VideoReaction) (*videoReactionDomain.VideoReaction, error) {
	var tableData = &videoReactionDomain.VideoReaction{}
	if err := vrr.GormDB.Updates(entity).Find(tableData).Error; err != nil {
		return nil, err
	}
	return tableData, nil
}

func (vrr VideoReactionRepository) IsVideoReactionExist(videoId string) (bool, error) {
	var exists bool
	if err := vrr.GormDB.Model(&videoReactionDomain.VideoReaction{}).Select("count(*) > 0").Where("video_id = ?", videoId).Find(&exists).Error; err != nil {
		return false, err
	}
	return exists, nil
}

func (vrr VideoReactionRepository) DeleteVideoReaction(id string) (bool, error) {
	if err := vrr.GormDB.Where("video_id = ?", id).Delete(&videoReactionDomain.VideoReaction{}).Error; err != nil {
		return false, err
	}
	return true, nil
}
