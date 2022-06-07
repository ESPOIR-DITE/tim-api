package user_video_repo

import (
	"fmt"
	"github.com/google/uuid"
	"tim-api/config"
	"tim-api/domain"
)

func CreateUserVideoTable() bool {
	var tableData = &domain.UserVideo{}
	err := config.GetDatabase().AutoMigrate(tableData)
	if err != nil {
		return false
	}
	return true
}
func SetUserVideo() {
	err := config.GetDatabase().Set("gorm:table_options", "ENGINE=Distributed(cluster, default, hits)").AutoMigrate(&domain.UserVideo{})
	if err != nil {
		fmt.Println("UserVideo database config not set")
	} else {
		fmt.Println("UserVideo database config set successfully")
	}
}
func CreateUserVideo(entity domain.UserVideo) domain.UserVideo {
	var tableData = &domain.UserVideo{}
	id := "UV-" + uuid.New().String()
	user := domain.UserVideo{id, entity.CustomerId, entity.VideoId, entity.Date}
	config.GetDatabase().Create(user).Find(&tableData)
	return user
}
func UpdateUserVideo(entity domain.UserVideo) *domain.UserVideo {
	var tableData = &domain.UserVideo{}
	//userSubscription := domain.UserSubscription{entity.Id, entity.Name, entity.Description}
	config.GetDatabase().Create(entity).Find(&tableData)
	return tableData
}
func GetAllUserVideo(customerId string) []domain.UserVideo {
	entity := []domain.UserVideo{}
	config.GetDatabase().Where("customer_id = ?", customerId).Find(&entity)
	return entity
}
func GetUserVideo(id string) domain.UserVideo {
	entity := domain.UserVideo{}
	config.GetDatabase().Where("id = ?", id).Find(&entity)
	return entity
}
func GetUserVideos() []domain.UserVideo {
	var entity []domain.UserVideo
	config.GetDatabase().Find(&entity)
	return entity
}
func DeleteUserVideo(id string) bool {
	entity := domain.UserVideo{}
	config.GetDatabase().Where("id = ?", id).Delete(&entity)
	if entity.CustomerId == "" {
		return true
	}
	return false
}
func GetUserVideoObject(userVideo *domain.UserVideo) domain.UserVideo {
	return domain.UserVideo{userVideo.Id, userVideo.CustomerId, userVideo.VideoId, userVideo.Date}
}
