package role_repo

import (
	"fmt"
	"github.com/google/uuid"
	"tim-api/config"
	"tim-api/domain"
)

var connection = config.GetDatabase()

func CreateChannelTable() bool {
	var tableData = &domain.Channel{}
	err := connection.AutoMigrate(tableData)
	if err != nil {
		return false
	}
	return true
}
func SetChannelDatabase() {
	err := connection.Set("gorm:table_options", "ENGINE=Distributed(cluster, default, hits)").AutoMigrate(&domain.Channel{})
	if err != nil {
		fmt.Println("Role database config not set")
	} else {
		fmt.Println("Role database config set successfully")
	}
}
func CreateChannel(channel domain.Channel) domain.Channel {
	var tableData = domain.Channel{}
	id := "C-" + uuid.New().String()
	user := domain.Channel{id, channel.Name, channel.ChannelTypeId, channel.UserId, channel.Region, channel.Date, channel.Image, "", channel.Description}
	connection.Create(user).Find(&tableData)
	return tableData
}
func UpdateChannel(entity domain.Channel) domain.Channel {
	var tableData = domain.Channel{}
	//user := domain.Role{entity.Id, entity.Name, entity.Description}
	connection.Where("id = ", entity.Id).Updates(entity).Find(&tableData)
	return tableData
}
func GetChannel(roleId string) domain.Channel {
	entity := domain.Channel{}
	connection.Where("id = ?", roleId).Find(&entity)
	return entity
}
func GetChannelsByUser(userId string) []domain.Channel {
	entity := []domain.Channel{}
	connection.Where("user_id = ?", userId).Find(&entity)
	return entity
}
func GetChannelsByRegion(region string) []domain.Channel {
	entity := []domain.Channel{}
	config.GetDatabase().Where("region = ?", region).Find(&entity)
	return entity
}
func GetChannels() []domain.Channel {
	var entity []domain.Channel
	connection.Find(&entity)
	return entity
}
func DeleteChannel(id string) bool {
	entity := domain.ChannelType{}
	connection.Where("id = ?", id).Delete(&entity)
	if entity.Id == "" {
		return true
	}
	return false
}
func CountChannel() int64 {
	var value int64
	connection.Table("channels").Count(&value)
	return value
}
func GetChannelObject(channel *domain.Channel) domain.Channel {
	return domain.Channel{channel.Id, channel.Name, channel.ChannelTypeId, channel.UserId, channel.Region, channel.Date, channel.Image, channel.ImageBase64, channel.Description}
}
