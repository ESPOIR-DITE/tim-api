package video_repo

import (
	"fmt"
	"github.com/google/uuid"
	"tim-api/config"
	"tim-api/domain"
)

var connection = config.GetDatabase()

func CreateVideoTable() bool {
	var tableData = &domain.Video{}
	err := connection.AutoMigrate(tableData)
	if err != nil {
		return false
	}
	return true
}
func SetDatabase() {
	erro := connection.Set("gorm:table_options", "ENGINE=Distributed(cluster, default, hits)").AutoMigrate(&domain.Video{})
	if erro != nil {
		fmt.Println("User database config not set")
	} else {
		fmt.Println("User database config set successfully")
	}
}
func CreateVideo(entity domain.Video) *domain.Video {
	var tableData = &domain.Video{}
	id := "V-" + uuid.New().String()
	user := domain.Video{id, entity.Title, entity.Date, entity.DateUploaded, entity.Description, entity.IsPrivate, entity.Price, entity.URL}
	connection.Create(user).Find(&tableData)
	return tableData
}
func UpdateVideo(entity domain.Video) *domain.Video {
	var tableData = &domain.Video{}
	connection.Create(entity).Find(&tableData)
	return tableData
}
func GetVideo(id string) domain.Video {
	entity := domain.Video{}
	connection.Where("id = ?", id).Find(&entity)
	return entity
}
func GetVideos() []domain.Video {
	entity := []domain.Video{}
	connection.Find(&entity)
	return entity
}
func DeleteVideo(email string) bool {
	entity := domain.User{}
	connection.Where("id = ?", email).Delete(&entity)
	if entity.Email == "" {
		return true
	}
	return false
}
func GetVideoObject(entity *domain.Video) domain.Video {
	return domain.Video{entity.Id, entity.Title, entity.Date, entity.DateUploaded, entity.Description, entity.IsPrivate, entity.Price, entity.URL}
}
