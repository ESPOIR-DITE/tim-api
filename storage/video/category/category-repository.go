package video_repo

import (
	"fmt"
	"github.com/google/uuid"
	"tim-api/config"
	"tim-api/domain"
)

var connection = config.GetDatabase()

func CreateCategoryTable() bool {
	var tableData = &domain.Category{}
	err := connection.AutoMigrate(tableData)
	if err != nil {
		return false
	}
	return true
}
func SetDatabase() {
	erro := connection.Set("gorm:table_options", "ENGINE=Distributed(cluster, default, hits)").AutoMigrate(&domain.Category{})
	if erro != nil {
		fmt.Println("Category database config not set")
	} else {
		fmt.Println("Category database config set successfully")
	}
}
func CreateCategory(entity domain.Category) *domain.Category {
	var tableData = &domain.Category{}
	id := "C-" + uuid.New().String()
	user := domain.Category{id, entity.Name, entity.Description}
	connection.Create(user).Find(&tableData)
	return tableData
}
func UpdateCategory(entity domain.Category) *domain.Category {
	var tableData = &domain.Category{}
	connection.Create(entity).Find(&tableData)
	return tableData
}
func GetCategory(id string) domain.Category {
	entity := domain.Category{}
	connection.Where("id = ?", id).Find(&entity)
	return entity
}
func GetCategories() []domain.Category {
	entity := []domain.Category{}
	connection.Find(&entity)
	return entity
}
func DeleteCategory(email string) bool {
	entity := domain.Category{}
	connection.Where("id = ?", email).Delete(&entity)
	if entity.Id == "" {
		return true
	}
	return false
}
func GetCategoryObject(entity *domain.Category) domain.Category {
	return domain.Category{entity.Id, entity.Name, entity.Description}
}
