package user_repo

import (
	"fmt"
	"tim-api/config"
	"tim-api/domain"
)

var connection = config.GetDatabase()

func CreateUserTable() bool {
	var tableData = &domain.User{}
	err := connection.AutoMigrate(tableData)
	if err != nil {
		return false
	}
	return true
}
func SetDatabase() {
	erro := connection.Set("gorm:table_options", "ENGINE=Distributed(cluster, default, hits)").AutoMigrate(&domain.User{})
	if erro != nil {
		fmt.Println("User database config not set")
	} else {
		fmt.Println("User database config set successfully")
	}
}
func CreateUser(entity domain.User) *domain.User {
	var tableData = &domain.User{}
	//id := "U-"+uuid.New().String()
	user := domain.User{entity.Email, entity.Name, entity.Surname, entity.BirthDate, entity.RoleId}
	connection.Create(user).Find(&tableData)
	return tableData
}
func UpdateUser(entity domain.User) *domain.User {
	var tableData = &domain.User{}
	user := domain.User{entity.Email, entity.Name, entity.Surname, entity.BirthDate, entity.RoleId}
	connection.Updates(user).Find(&tableData)
	return tableData
}
func GetUser(email string) domain.User {
	entity := domain.User{}
	connection.Where("email = ?", email).Find(&entity)
	return entity
}
func GetUsers() []domain.User {
	entity := []domain.User{}
	connection.Find(&entity)
	return entity
}
func Delete(email string) bool {
	entity := domain.User{}
	connection.Where("email = ?", email).Delete(&entity)
	if entity.Email == "" {
		return true
	}
	return false
}
func GetUserObject(user *domain.User) domain.User {
	return domain.User{user.Email, user.Name, user.Surname, user.BirthDate, user.RoleId}
}
