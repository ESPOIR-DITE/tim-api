package role_repo

import (
	"fmt"
	"github.com/google/uuid"
	"tim-api/config"
	"tim-api/domain"
)

func CreateRoleTable() bool {
	var tableData = &domain.Role{}
	err := config.GetDatabase().AutoMigrate(tableData)
	if err != nil {
		return false
	}
	return true
}
func SetRoleDatabase() {
	err := config.GetDatabase().Set("gorm:table_options", "ENGINE=Distributed(cluster, default, hits)").AutoMigrate(&domain.User{})
	if err != nil {
		fmt.Println("Role database config not set")
	} else {
		fmt.Println("Role database config set successfully")
	}
}
func CreateRole(entity domain.Role) *domain.Role {
	var tableData = &domain.Role{}
	id := "R-" + uuid.New().String()
	user := domain.Role{id, entity.Name, entity.Description}
	config.GetDatabase().Create(user).Find(&tableData)
	return tableData
}
func UpdateRole(entity domain.Role) *domain.Role {
	var tableData = &domain.Role{}
	//user := domain.Role{entity.Id, entity.Name, entity.Description}
	config.GetDatabase().Create(entity).Find(&tableData)
	return tableData
}
func GetRole(customerId string) domain.Role {
	entity := domain.Role{}
	config.GetDatabase().Where("id = ?", customerId).Find(&entity)
	return entity
}
func GetRoles() []domain.Role {
	var entity []domain.Role
	config.GetDatabase().Find(&entity)
	return entity
}
func DeleteRole(id string) bool {
	entity := domain.Role{}
	config.GetDatabase().Where("id = ?", id).Delete(&entity)
	if entity.Id == "" {
		return true
	}
	return false
}
func GetRoleObject(role *domain.Role) domain.Role {
	return domain.Role{role.Id, role.Name, role.Description}
}
