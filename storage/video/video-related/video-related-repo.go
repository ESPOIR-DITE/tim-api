package videoRelatedRepository

import (
	"fmt"
	videoRelatedDomain "github.com/ESPOIR-DITE/tim-api/domain/video/video.related.domain"
	"gorm.io/gorm"
)

type VideoRelatedRepository struct {
	GormDB *gorm.DB
}

func NewVideoRelatedRepository(gormDB *gorm.DB) *VideoRelatedRepository {
	return &VideoRelatedRepository{
		GormDB: gormDB,
	}
}

func (vrr VideoRelatedRepository) CreateVideoRelated(entity videoRelatedDomain.VideoRelated) (*videoRelatedDomain.VideoRelated, error) {
	var tableData = &videoRelatedDomain.VideoRelated{}
	err := vrr.GormDB.Create(entity).Find(tableData).Error
	if err != nil {
		fmt.Println("error creating video.controller related")
		return tableData, err
	}
	return tableData, nil
}
func (vrr VideoRelatedRepository) GetVideosRelatedTo(videoId string) ([]videoRelatedDomain.VideoRelated, error) {
	entity := []videoRelatedDomain.VideoRelated{}
	err := vrr.GormDB.Where("current_video_id = ?", videoId).Or("related_video_id = ?", videoId).Find(&entity).Error
	if err != nil {
		fmt.Println("error reading video.controller related")
		return entity, err
	}
	return entity, nil
}
func (vrr VideoRelatedRepository) DeleteVideoRelated(videoRelated videoRelatedDomain.VideoRelated) (bool, error) {
	err := vrr.GormDB.Where("current_video_id = ? AND related_video_id = ?", videoRelated.CurrentVideoId, videoRelated.RelatedVideoId).Delete(&videoRelatedDomain.VideoRelated{}).Error
	if err != nil {
		fmt.Println("error deleting video.controller related")
		return false, err
	}
	return true, nil
}
