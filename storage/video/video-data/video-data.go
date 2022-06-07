package video_data

import (
	"fmt"
	"github.com/google/uuid"
	"tim-api/config"
	"tim-api/domain"
)

var connection = config.GetDatabase()

func CreateVideoDataTable() bool {
	var tableData = &domain.VideoDate{}
	err := connection.AutoMigrate(tableData)
	if err != nil {
		return false
	}
	return true
}
func SetDatabase() {
	erro := connection.Set("gorm:table_options", "ENGINE=Distributed(cluster, default, hits)").AutoMigrate(&domain.VideoDate{})
	if erro != nil {
		fmt.Println("Category database config not set")
	} else {
		fmt.Println("Category database config set successfully")
	}
}
func CreateVideoData(entity domain.VideoDate) *domain.VideoDate {
	var tableData = &domain.VideoDate{}
	id := "C-" + uuid.New().String()
	user := domain.VideoDate{id, entity.Video, entity.VideoGif, entity.FileType}
	connection.Create(user).Find(&tableData)
	return tableData
}
func UpdateVideoDate(entity domain.VideoDate) *domain.VideoDate {
	var tableData = &domain.VideoDate{}
	connection.Updates(entity).Find(&tableData)
	return tableData
}
func GetVideoDate(id string) domain.VideoDate {
	entity := domain.VideoDate{}
	connection.Where("id = ?", id).Find(&entity)
	return entity
}
func GetVideoDatas() []domain.VideoDate {
	entity := []domain.VideoDate{}
	connection.Find(&entity)
	return entity
}
func DeleteVideoData(email string) bool {
	entity := domain.VideoDate{}
	connection.Where("id = ?", email).Delete(&entity)
	if entity.Id == "" {
		return true
	}
	return false
}
func GetCategoryObject(entity *domain.VideoDate) domain.VideoDate {
	return domain.VideoDate{entity.Id, entity.Video, entity.VideoGif, entity.FileType}
}
