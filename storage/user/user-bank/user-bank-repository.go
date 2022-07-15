package user_repo

import (
	"fmt"
	"github.com/google/uuid"
	"tim-api/config"
	user_details "tim-api/domain/user/user-details"
)

var connection = config.GetDatabase()

func CreateUserBankTable() bool {
	var tableData = &user_details.UserBank{}
	err := connection.AutoMigrate(tableData)
	if err != nil {
		return false
	}
	return true
}

func SetDatabase() {
	err := connection.Set("gorm:table_options", "ENGINE=Distributed(cluster, default, hits)").AutoMigrate(&user_details.UserBank{})
	if err != nil {
		fmt.Println("UserBank database config not set")
	} else {
		fmt.Println("UserBank database config set successfully")
	}
}

func CreateUserDetails(entity user_details.UserBank) user_details.UserBank {
	var tableData = user_details.UserBank{}
	id := "UB-" + uuid.New().String()
	userDetails := user_details.UserBank{id, entity.UserEmail, entity.BankType, entity.BankName, entity.BranchCode, entity.BankNumber, entity.CvcCode}
	connection.Create(userDetails).Find(&tableData)
	return tableData
}
func UpdateUserBank(entity user_details.UserBank) user_details.UserBank {
	var tableData = user_details.UserBank{}
	userBank := user_details.UserBank{entity.Id, entity.UserEmail, entity.BankType, entity.BankName, entity.BranchCode, entity.BankNumber, entity.CvcCode}
	connection.Updates(userBank).Find(&tableData)
	return tableData
}
func GetUserBank(id string) user_details.UserBank {
	entity := user_details.UserBank{}
	connection.Where("id = ?", id).Find(&entity)
	return entity
}
func GetUserBankByEmail(userEmail string) user_details.UserBank {
	entity := user_details.UserBank{}
	connection.Where("user_email = ?", userEmail).Find(&entity)
	return entity
}

func GetUserBanks() []user_details.UserBank {
	entity := []user_details.UserBank{}
	connection.Find(&entity)
	return entity
}
func DeleteUserBankByUserEmail(email string) bool {
	entity := user_details.UserBank{}
	connection.Where("user_email = ?", email).Delete(&entity)
	if entity.Id == "" {
		return true
	}
	return false
}
func DeleteUserBankById(id string) bool {
	entity := user_details.UserBank{}
	connection.Where("id = ?", id).Delete(&entity)
	if entity.Id == "" {
		return true
	}
	return false
}
func IsExistUserBankById(id string) bool {
	entity := user_details.UserBank{}
	connection.Where("id = ?", id).Find(&entity)
	if entity.Id != "" {
		return true
	}
	return false
}
func IsExistUserBankByEmail(email string) bool {
	entity := user_details.UserBank{}
	connection.Where("user_email = ?", email).Find(&entity)
	if entity.Id != "" {
		return true
	}
	return false
}
