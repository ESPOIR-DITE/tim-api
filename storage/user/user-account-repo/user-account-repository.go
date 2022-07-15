package user_account_repo

import (
	"fmt"
	"github.com/google/uuid"
	"tim-api/config"
	"tim-api/domain"
)

var connection = config.GetDatabase()

func CreateUserAccountTable() bool {
	var tableData = &domain.UserAccount{}
	err := connection.AutoMigrate(tableData)
	if err != nil {
		return false
	}
	return true
}
func SetUserAccountDatabase() {
	err := connection.Set("gorm:table_options", "ENGINE=Distributed(cluster, default, hits)").AutoMigrate(&domain.UserAccount{})
	if err != nil {
		fmt.Println("User Account database config not set")
	} else {
		fmt.Println("User Account database config set successfully")
	}
}
func CreateUserAccount(entity domain.UserAccount) domain.UserAccount {
	var tableData = domain.UserAccount{}
	id := "UA-" + uuid.New().String()
	user := domain.UserAccount{id, entity.Email, entity.Password, entity.Date, entity.Status}
	config.GetDatabase().Create(user).Find(&tableData)
	return tableData
}
func UpdateUserAccount(entity domain.UserAccount) domain.UserAccount {
	var tableData = domain.UserAccount{}
	user := domain.UserAccount{entity.CustomerId, entity.Email, entity.Password, entity.Date, entity.Status}
	connection.Updates(user).Find(&tableData)
	return tableData
}
func GetUserAccount(customerId string) domain.UserAccount {
	entity := domain.UserAccount{}
	connection.Where("customer_id = ?", customerId).Find(&entity)
	return entity
}
func GetUserAccountWithEmail(email string) domain.UserAccount {
	entity := domain.UserAccount{}
	connection.Where("Email = ?", email).Find(&entity)
	return entity
}
func GetUserAccounts() []domain.UserAccount {
	var entity []domain.UserAccount
	connection.Find(&entity)
	return entity
}
func Login(userAccount domain.UserAccount) domain.UserAccount {
	entity := domain.UserAccount{}
	connection.Where("email = ? and password = ?", userAccount.Email, userAccount.Password).Find(&entity)
	return entity
}
func DeleteUserAccount(customerId string) bool {
	entity := domain.UserAccount{}
	connection.Where("CustomerId = ?", customerId).Delete(&entity)
	if entity.Email == "" {
		return true
	}
	return false
}
func countUserAccount() int64 {
	var value int64
	connection.Table("user_accounts").Count(&value)
	return value
}
func GetUserAccountObject(account *domain.UserAccount) domain.UserAccount {
	return domain.UserAccount{account.CustomerId, account.Email, account.Password, account.Date, account.Status}
}
