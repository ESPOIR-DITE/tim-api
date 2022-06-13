package video_data

import (
	"fmt"
	"tim-api/config"
	"tim-api/domain"
)

var connection = config.GetDatabase()

func CreateVideoDataTable() bool {
	var tableData = &domain.VideoData{}
	err := connection.AutoMigrate(tableData)
	if err != nil {
		return false
	}
	return true
}
func SetDatabase() {
	erro := connection.Set("gorm:table_options", "ENGINE=Distributed(cluster, default, hits)").AutoMigrate(&domain.VideoData{})
	if erro != nil {
		fmt.Println("Category database config not set")
	} else {
		fmt.Println("Category database config set successfully")
	}
}
func CreateVideoData(entity domain.VideoData) *domain.VideoData {
	var tableData = &domain.VideoData{}
	//id := "C-" + uuid.New().String()
	user := domain.VideoData{entity.Id, entity.Picture, entity.Video, entity.FileType, entity.FileSize}
	connection.Create(user).Find(&tableData)
	return tableData
}
func UpdateVideoDate(entity domain.VideoData) *domain.VideoData {
	var tableData = &domain.VideoData{}
	connection.Updates(entity).Find(&tableData)
	return tableData
}
func GetVideoDate(id string) domain.VideoData {
	entity := domain.VideoData{}
	connection.Where("id = ?", id).Find(&entity)
	return entity
}
func GetVideoDatas() []domain.VideoData {
	entity := []domain.VideoData{}
	connection.Find(&entity)
	return entity
}
func DeleteVideoData(email string) bool {
	entity := domain.VideoData{}
	connection.Where("id = ?", email).Delete(&entity)
	if entity.Id == "" {
		return true
	}
	return false
}
func GetCategoryObject(entity *domain.VideoData) domain.VideoData {
	return domain.VideoData{entity.Id, entity.Picture, entity.Video, entity.FileType, entity.FileSize}
}
