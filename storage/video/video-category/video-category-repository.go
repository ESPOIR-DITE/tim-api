package video_category

import (
	"fmt"
	"github.com/google/uuid"
	"tim-api/config"
	"tim-api/domain"
)

var connection = config.GetDatabase()

func CreateVideoCategoryTable() bool {
	var tableData = &domain.VideoCategory{}
	err := connection.AutoMigrate(tableData)
	if err != nil {
		return false
	}
	return true
}
func SetDatabase() {
	erro := connection.Set("gorm:table_options", "ENGINE=Distributed(cluster, default, hits)").AutoMigrate(&domain.VideoCategory{})
	if erro != nil {
		fmt.Println("VideoCategory database config not set")
	} else {
		fmt.Println("VideoCategory database config set successfully")
	}
}
func CreateVideoCategory(entity domain.VideoCategory) *domain.VideoCategory {
	var tableData = &domain.VideoCategory{}
	id := "VC-" + uuid.New().String()
	user := domain.VideoCategory{id, entity.VideoId, entity.CategoryId}
	connection.Create(user).Find(&tableData)
	return tableData
}
func UpdateVideoCategory(entity domain.VideoCategory) *domain.VideoCategory {
	var tableData = &domain.VideoCategory{}
	connection.Create(entity).Find(&tableData)
	return tableData
}
func GetVideo(id string) domain.VideoCategory {
	entity := domain.VideoCategory{}
	connection.Where("id = ?", id).Find(&entity)
	return entity
}
func GetVideoCategories() []domain.VideoCategory {
	entity := []domain.VideoCategory{}
	connection.Find(&entity)
	return entity
}
func DeleteVideoCategory(email string) bool {
	entity := domain.VideoCategory{}
	connection.Where("id = ?", email).Delete(&entity)
	if entity.Id == "" {
		return true
	}
	return false
}
func GetVideoObject(entity *domain.VideoCategory) domain.VideoCategory {
	return domain.VideoCategory{entity.Id, entity.VideoId, entity.CategoryId}
}
