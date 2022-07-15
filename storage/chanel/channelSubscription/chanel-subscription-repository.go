package role_repo

import (
	"fmt"
	"github.com/google/uuid"
	"tim-api/config"
	"tim-api/domain"
)

var connection = config.GetDatabase()

func CreateChannelSubscriptionTable() bool {
	var tableData = &domain.ChannelSubscription{}
	err := connection.AutoMigrate(tableData)
	if err != nil {
		return false
	}
	return true
}
func SetChannelSubscriptionDatabase() {
	err := connection.Set("gorm:table_options", "ENGINE=Distributed(cluster, default, hits)").AutoMigrate(&domain.ChannelSubscription{})
	if err != nil {
		fmt.Println("Role database config not set")
	} else {
		fmt.Println("Role database config set successfully")
	}
}
func CreateChannelSubscription(channel domain.ChannelSubscription) domain.ChannelSubscription {
	var tableData = domain.ChannelSubscription{}
	id := "CS-" + uuid.New().String()
	user := domain.ChannelSubscription{id, channel.ChannelId, channel.UserId, channel.Date}
	connection.Create(user).Find(&tableData)
	return tableData
}
func UpdateChannelSubscription(entity domain.ChannelSubscription) domain.ChannelSubscription {
	var tableData = domain.ChannelSubscription{}
	//user := domain.Role{entity.Id, entity.Name, entity.Description}
	connection.Where("id = ", entity.Id).Updates(entity).Find(&tableData)
	return tableData
}
func GetChannelSubscription(roleId string) domain.ChannelSubscription {
	entity := domain.ChannelSubscription{}
	connection.Where("id = ?", roleId).Find(&entity)
	return entity
}
func GetChannelSubscriptionsByUser(userId string) []domain.ChannelSubscription {
	entity := []domain.ChannelSubscription{}
	config.GetDatabase().Where("user_id = ?", userId).Find(&entity)
	return entity
}
func GetChannelSubscriptionsByChannelId(channelId string) []domain.ChannelSubscription {
	entity := []domain.ChannelSubscription{}
	connection.Where("channel_id = ?", channelId).Find(&entity)
	return entity
}

func GetChannelSubscriptions() []domain.ChannelSubscription {
	var entity []domain.ChannelSubscription
	connection.Find(&entity)
	return entity
}
func DeleteChannelSubscription(id string) bool {
	entity := domain.ChannelSubscription{}
	connection.Where("id = ?", id).Delete(&entity)
	if entity.Id == "" {
		return true
	}
	return false
}

func CountSubscriptionByChannelId(channelId string) int64 {
	var value int64
	connection.Table("channel_subscription").Where("channel_id = ?", channelId).Count(&value)
	return value
}
func GetChannelSubscriptionObject(channel *domain.ChannelSubscription) domain.ChannelSubscription {
	return domain.ChannelSubscription{channel.Id, channel.ChannelId, channel.UserId, channel.Date}
}
