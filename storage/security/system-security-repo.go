package security

import (
	"github.com/ESPOIR-DITE/tim-api/config"
	"github.com/ESPOIR-DITE/tim-api/domain/security"
	"log"
)

var connection = config.GetDatabase()

func CreateSystemSecurityTable() bool {
	var tableData = &security.SystemData{}
	err := connection.AutoMigrate(tableData)
	if err != nil {
		return false
	}
	return true
}
func CreateSystemData(entity security.SystemData) (security.SystemData, error) {
	var tableData = security.SystemData{}
	err := connection.Create(entity).Find(&tableData).Error
	if err != nil {
		return tableData, err
	}
	return tableData, nil
}

func GetSystemData(id string) (security.SystemData, error) {
	entity := security.SystemData{}
	err := connection.Where("identifier = ?", id).Find(&entity).Error
	if err != nil {
		return entity, err
	}
	return entity, nil
}

func UpdateSystemData(entity security.SystemData) (security.SystemData, error) {
	var tableData = security.SystemData{}
	err := connection.Updates(entity).Find(&tableData).Error
	if err != nil {
		log.Fatal(err)
		return tableData, err
	}
	return tableData, nil
}
