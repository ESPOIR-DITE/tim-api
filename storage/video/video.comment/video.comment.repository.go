package videoCommentRepository

import (
	videoComment "github.com/ESPOIR-DITE/tim-api/domain/video/video.comment.domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VideoCommentRepository struct {
	GormDB *gorm.DB
}

func NewVideoCommentRepository(gormDB *gorm.DB) *VideoCommentRepository {
	return &VideoCommentRepository{
		GormDB: gormDB,
	}
}

func (vr VideoCommentRepository) CreateVideoComment(entity videoComment.VideoComment) (*videoComment.VideoComment, error) {
	var tableData = &videoComment.VideoComment{}
	id := "VC-" + uuid.New().String()
	entity.Id = id
	if err := vr.GormDB.Create(entity).First(tableData).Error; err != nil {
		return nil, err
	}
	return tableData, nil
}

func (vr VideoCommentRepository) UpdateVideoComment(entity videoComment.VideoComment) (*videoComment.VideoComment, error) {
	var tableData = &videoComment.VideoComment{}
	if err := vr.GormDB.Create(entity).First(&tableData).Error; err != nil {
		return nil, err
	}
	return tableData, nil
}

func (vr VideoCommentRepository) GetVideoComment(id string) (*videoComment.VideoComment, error) {
	entity := &videoComment.VideoComment{}
	if err := vr.GormDB.Where("id = ?", id).First(entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func (vr VideoCommentRepository) GetVideoComments() ([]videoComment.VideoComment, error) {
	entity := []videoComment.VideoComment{}
	if err := vr.GormDB.Find(&entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}
func (vr VideoCommentRepository) DeleteVideoComment(email string) (bool, error) {
	entity := videoComment.VideoComment{}
	if err := vr.GormDB.Where("id = ?", email).Delete(&entity).Error; err != nil {
		return false, err
	}
	if entity.Id == "" {
		return true, nil
	}
	return false, nil
}
