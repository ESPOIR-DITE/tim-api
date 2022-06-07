package video_category

import (
	"fmt"
	"github.com/google/uuid"
	"tim-api/config"
	"tim-api/domain"
)

var connection = config.GetDatabase()

func CreateVideoCommentTable() bool {
	var tableData = &domain.VideoComment{}
	err := connection.AutoMigrate(tableData)
	if err != nil {
		return false
	}
	return true
}
func SetDatabase() {
	erro := connection.Set("gorm:table_options", "ENGINE=Distributed(cluster, default, hits)").AutoMigrate(&domain.VideoComment{})
	if erro != nil {
		fmt.Println("VideoComment database config not set")
	} else {
		fmt.Println("VideoComment database config set successfully")
	}
}
func CreateVideoComment(entity domain.VideoComment) *domain.VideoComment {
	var tableData = &domain.VideoComment{}
	id := "VC-" + uuid.New().String()
	user := domain.VideoComment{id, entity.VideoId, entity.CommentId}
	connection.Create(user).Find(&tableData)
	return tableData
}
func UpdateVideoComment(entity domain.VideoComment) *domain.VideoComment {
	var tableData = &domain.VideoComment{}
	connection.Create(entity).Find(&tableData)
	return tableData
}
func GetVideoComment(id string) domain.VideoComment {
	entity := domain.VideoComment{}
	connection.Where("id = ?", id).Find(&entity)
	return entity
}
func GetVideoComments() []domain.VideoComment {
	entity := []domain.VideoComment{}
	connection.Find(&entity)
	return entity
}
func DeleteVideoComment(email string) bool {
	entity := domain.VideoComment{}
	connection.Where("id = ?", email).Delete(&entity)
	if entity.Id == "" {
		return true
	}
	return false
}
func GetVideoCommentObject(entity *domain.VideoComment) domain.VideoComment {
	return domain.VideoComment{entity.Id, entity.VideoId, entity.CommentId}
}
