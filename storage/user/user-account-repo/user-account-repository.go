package user_account_repo

import (
	"fmt"
	"github.com/google/uuid"
	"tim-api/config"
	user_account "tim-api/domain/user/user-account"
	"tim-api/security"
)

var connection = config.GetDatabase()

func CreateUserAccountTable() bool {
	var tableData = &user_account.UserAccount{}
	err := connection.AutoMigrate(tableData)
	if err != nil {
		return false
	}
	return true
}
func SetUserAccountDatabase() {
	err := connection.Set("gorm:table_options", "ENGINE=Distributed(cluster, default, hits)").AutoMigrate(&user_account.UserAccount{})
	if err != nil {
		fmt.Println("User Account database config not set")
	} else {
		fmt.Println("User Account database config set successfully")
	}
}
func CreateUserAccount(entity user_account.UserAccount) user_account.UserAccount {
	var tableData = user_account.UserAccount{}
	id := "UA-" + uuid.New().String()
	user := user_account.UserAccount{id, entity.Email, entity.Password, entity.Date, entity.Status, ""}
	config.GetDatabase().Create(user).Find(&tableData)
	return tableData
}
func UpdateUserAccount(entity user_account.UserAccount) user_account.UserAccount {
	var tableData = user_account.UserAccount{}
	user := user_account.UserAccount{entity.CustomerId, entity.Email, entity.Password, entity.Date, entity.Status, entity.Token}
	connection.Updates(user).Find(&tableData)
	return tableData
}
func UpdateToken(entity user_account.UserAccount, token string) (user_account.UserAccount, error) {
	var tableData = user_account.UserAccount{}
	err := connection.Model(&tableData).Where("customer_id = ?", entity.CustomerId).Update("token", token).Error
	if err != nil {
		return entity, err
	}
	return entity, nil
}
func GetAllUserAccountByEmail(email string) []user_account.UserAccount {
	entity := []user_account.UserAccount{}
	connection.Where("email = ?", email).Find(&entity)
	return entity
}
func GetUserAccount(customerId string) user_account.UserAccount {
	entity := user_account.UserAccount{}
	connection.Where("customer_id = ?", customerId).Find(&entity)
	return entity
}
func GetUserAccountWithEmail(email string) user_account.UserAccount {
	entity := user_account.UserAccount{}
	connection.Where("Email = ?", email).Find(&entity)
	return entity
}
func GetUserAccounts() []user_account.UserAccount {
	var entity []user_account.UserAccount
	connection.Find(&entity)
	return entity
}
func Login(userAccount user_account.UserAccount) (user_account.UserAccount, error) {
	entity := user_account.UserAccount{}
	err := connection.Where("email = ? and password = ?", userAccount.Email, userAccount.Password).Find(&entity).Error
	if err != nil {
		return entity, err
	}
	return entity, nil
}
func DeleteUserAccount(customerId string) bool {
	entity := user_account.UserAccount{}
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
func GetUserAccountObject(account *user_account.UserAccount) user_account.UserAccount {
	password, err := security.EncryptPassword(account.Password)
	if err != nil {
		password = account.Password
	}
	return user_account.UserAccount{account.CustomerId, account.Email, password, account.Date, account.Status, account.Token}
}
func GetDecodedUserAccountObject(account *user_account.UserAccount) user_account.UserAccount {
	password, err := security.EncryptPassword(account.Password)
	if err != nil {
		password = account.Password
	}
	return user_account.UserAccount{account.CustomerId, account.Email, password, account.Date, account.Status, account.Token}
}
