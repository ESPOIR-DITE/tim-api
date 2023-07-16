package accountRepository

import (
	"github.com/ESPOIR-DITE/tim-api/domain/user/account"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AccountRepository struct {
	GormDB *gorm.DB
}

func NewAccountRepository(gormDB *gorm.DB) *AccountRepository {
	return &AccountRepository{
		GormDB: gormDB,
	}
}

func (uap *AccountRepository) CreateAccount(entity account.Account) (*account.Account, error) {
	id := "UA-" + uuid.New().String()
	entity.Id = id
	if err := uap.GormDB.Create(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (uap *AccountRepository) UpdateAccount(entity account.Account) (*account.Account, error) {
	var tableData = &account.Account{}
	if err := uap.GormDB.Updates(entity).First(&tableData).Error; err != nil {
		return nil, err
	}
	return tableData, nil
}

func (uap *AccountRepository) UpdateToken(entity account.Account, token string) (*account.Account, error) {
	var tableData = &account.Account{}
	err := uap.GormDB.Model(&tableData).Where("id = ?", entity.Id).Update("token", token).Error
	if err != nil {
		return nil, err
	}
	return tableData, nil
}

func (uap *AccountRepository) GetAllAccountByEmail(email string) ([]account.Account, error) {
	entity := []account.Account{}
	if err := uap.GormDB.Where("email = ?", email).Find(&entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func (uap *AccountRepository) GetAccount(id string) (*account.Account, error) {
	entity := &account.Account{}
	if err := uap.GormDB.Where("id = ?", id).First(&entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func (uap *AccountRepository) GetAccountWithEmail(email string) (*account.Account, error) {
	entity := &account.Account{}
	if err := uap.GormDB.Where("email = ?", email).First(&entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}
func (uap *AccountRepository) GetAccounts() ([]account.Account, error) {
	var entity []account.Account
	if err := uap.GormDB.Find(&entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func (uap *AccountRepository) Login(userAccount account.Account) (*account.Account, error) {
	entity := &account.Account{}
	err := uap.GormDB.Where("email = ? and password = ?", userAccount.Email, userAccount.Password).First(&entity).Error
	if err != nil {
		return entity, err
	}
	return entity, nil
}
func (uap *AccountRepository) DeleteAccount(id string) (bool, error) {
	entity := &account.Account{}
	if err := uap.GormDB.Where("id = ?", id).Delete(&entity).Error; err != nil {
		return false, err
	}
	if entity.Id == "" {
		return true, nil
	}
	return false, nil
}
func (uap *AccountRepository) countAccount() (*int64, error) {
	var value int64
	if err := uap.GormDB.Table("account").Count(&value).Error; err != nil {
		return nil, err
	}
	return &value, nil
}

//	func (uap *AccountRepository) GetUserAccountObject(account *account.Account) (*account.Account, error) {
//		password, err := security.EncryptPassword(account.Password)
//		if err != nil {
//			return account, err
//		}
//		account.Password = password
//		return account, err
//	}
//
//	func (uap *AccountRepository) GetDecodedUserAccountObject(account *account.Account) (*account.Account, error) {
//		password, err := security.EncryptPassword(account.Password)
//		if err != nil {
//			return account, err
//		}
//		account.Password = password
//		return account, nil
//	}
func (uap *AccountRepository) LoginUser(email string) (*account.Account, error) {
	entity := &account.Account{}
	if err := uap.GormDB.Where("email = ?", email).First(entity).Error; err != nil {
		return nil, err
	}
	if entity.Id == "" {
		return nil, nil
	}
	return entity, nil
}
