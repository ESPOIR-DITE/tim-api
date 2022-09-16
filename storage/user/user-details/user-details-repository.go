package user_repo

import (
	"fmt"
	"github.com/google/uuid"
	"log"
	"tim-api/config"
	user_details "tim-api/domain/user/user-details"
)

var connection = config.GetDatabase()

func CreateUserDetailsTable() bool {
	var tableData = &user_details.UserDetails{}
	err := connection.AutoMigrate(tableData)
	if err != nil {
		return false
	}
	return true
}

func SetDatabase() {
	err := connection.Set("gorm:table_options", "ENGINE=Distributed(cluster, default, hits)").AutoMigrate(&user_details.UserDetails{})
	if err != nil {
		fmt.Println("User database config not set")
	} else {
		fmt.Println("User database config set successfully")
	}
}

func CreateUserDetails(entity user_details.UserDetails) user_details.UserDetails {
	var tableData = user_details.UserDetails{}
	id := "UD-" + uuid.New().String()
	userDetails := user_details.UserDetails{id, entity.UserEmail, entity.BankId, entity.CompanyRegisteredNumber, entity.TaxNumber}
	connection.Create(userDetails).Find(&tableData)
	return tableData
}
func UpdateUserDetails(entity user_details.UserDetails) user_details.UserDetails {
	var tableData = user_details.UserDetails{}
	user := user_details.UserDetails{entity.Id, entity.UserEmail, entity.BankId, entity.CompanyRegisteredNumber, entity.TaxNumber}
	connection.Updates(user).Find(&tableData)
	return tableData
}
func GetUserDetail(email string) user_details.UserDetails {
	entity := user_details.UserDetails{}
	connection.Where("id = ?", email).Find(&entity)
	return entity
}
func GetUserDetailsByEmail(email string) user_details.UserDetails {
	entity := user_details.UserDetails{}
	connection.Where("user_email = ?", email).Find(&entity)
	return entity
}

func GetUserDetails() []user_details.UserDetails {
	entity := []user_details.UserDetails{}
	connection.Find(&entity)
	return entity
}
func DeleteUserDetailsByUserEmail(email string) bool {
	entity := user_details.UserDetails{}
	connection.Where("user_email = ?", email).Delete(&entity)
	if entity.Id == "" {
		return true
	}
	return false
}
func DeleteUserDetailsById(id string) bool {
	entity := user_details.UserDetails{}
	connection.Where("id = ?", id).Delete(&entity)
	if entity.Id == "" {
		return true
	}
	return false
}
func IsExistUserDetailById(id string) bool {
	entity := user_details.UserDetails{}
	connection.Where("id = ?", id).Find(&entity)
	if entity.Id != "" {
		return true
	}
	return false
}
func IsExistUserDetailsByEmail(email string) bool {
	entity := user_details.UserDetails{}
	err := connection.Where("user_email = ?", email).Find(&entity).Error
	if err != nil {
		log.Fatal(err)
	}
	if entity.Id != "" {
		return true
	}
	return false
}
