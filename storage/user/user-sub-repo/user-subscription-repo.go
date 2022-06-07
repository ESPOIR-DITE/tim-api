package user_sub_repo

import (
	"fmt"
	"github.com/google/uuid"
	"tim-api/config"
	"tim-api/domain"
)

func CreateUserSubscriptionTable() bool {
	var tableData = &domain.UserSubscription{}
	err := config.GetDatabase().AutoMigrate(tableData)
	if err != nil {
		return false
	}
	return true
}
func SetUserSubscriptionDatabase() {
	err := config.GetDatabase().Set("gorm:table_options", "ENGINE=Distributed(cluster, default, hits)").AutoMigrate(&domain.UserSubscription{})
	if err != nil {
		fmt.Println("Role database config not set")
	} else {
		fmt.Println("Role database config set successfully")
	}
}
func CreateUserSubscription(entity domain.UserSubscription) *domain.UserSubscription {
	var tableData = &domain.UserSubscription{}
	id := "R-" + uuid.New().String()
	user := domain.UserSubscription{id, entity.UserId, entity.Stat, entity.SubscriptionId, entity.Date}
	config.GetDatabase().Create(user).Find(&tableData)
	return tableData
}
func UpdateUserSubscription(entity domain.UserSubscription) *domain.UserSubscription {
	var tableData = &domain.UserSubscription{}
	//userSubscription := domain.UserSubscription{entity.Id, entity.Name, entity.Description}
	config.GetDatabase().Create(entity).Find(&tableData)
	return tableData
}
func GetUserSubscription(customerId string) domain.UserSubscription {
	entity := domain.UserSubscription{}
	config.GetDatabase().Where("id = ?", customerId).Find(&entity)
	return entity
}
func GetUserSubscriptions() []domain.UserSubscription {
	var entity []domain.UserSubscription
	config.GetDatabase().Find(&entity)
	return entity
}
func DeleteUserSubscription(id string) bool {
	entity := domain.UserSubscription{}
	config.GetDatabase().Where("id = ?", id).Delete(&entity)
	if entity.Id == "" {
		return true
	}
	return false
}
func GetUserSubscriptionObject(subscription *domain.UserSubscription) domain.UserSubscription {
	return domain.UserSubscription{subscription.Id, subscription.UserId, subscription.Stat, subscription.SubscriptionId, subscription.Date}
}
