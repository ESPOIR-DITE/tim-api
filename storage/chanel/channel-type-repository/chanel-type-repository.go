package role_repo

import (
	"fmt"
	"github.com/google/uuid"
	"tim-api/config"
	"tim-api/domain"
)

func CreateChannelTypeTable() bool {
	var tableData = &domain.ChannelType{}
	err := config.GetDatabase().AutoMigrate(tableData)
	if err != nil {
		return false
	}
	return true
}
func SetChannelTypeDatabase() {
	err := config.GetDatabase().Set("gorm:table_options", "ENGINE=Distributed(cluster, default, hits)").AutoMigrate(&domain.ChannelType{})
	if err != nil {
		fmt.Println("Role database config not set")
	} else {
		fmt.Println("Role database config set successfully")
	}
}
func CreateChannelType(channelType domain.ChannelType) (domain.ChannelType, error) {
	var tableData = domain.ChannelType{}
	id := "CT-" + uuid.New().String()
	user := domain.ChannelType{id, channelType.Name, channelType.Description}
	err := config.GetDatabase().Create(user).Find(&tableData).Error
	if err != nil {
		return tableData, err
	}
	return tableData, nil
}
func UpdateChannelType(entity domain.ChannelType) domain.ChannelType {
	var tableData = domain.ChannelType{}
	config.GetDatabase().Where("id = ", entity.Id).Updates(entity).Find(&tableData)
	return tableData
}
func GetChannelType(roleId string) domain.ChannelType {
	entity := domain.ChannelType{}
	config.GetDatabase().Where("id = ?", roleId).Find(&entity)
	return entity
}
func GetChannelTypes() []domain.ChannelType {
	var entity []domain.ChannelType
	config.GetDatabase().Find(&entity)
	return entity
}
func DeleteChannelType(id string) bool {
	entity := domain.ChannelType{}
	config.GetDatabase().Where("id = ?", id).Delete(&entity)
	if entity.Id == "" {
		return true
	}
	return false
}
func GetChannelObject(channel *domain.ChannelType) domain.ChannelType {
	return domain.ChannelType{channel.Id, channel.Name, channel.Description}
}
