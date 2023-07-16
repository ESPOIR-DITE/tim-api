package userAccountRepository

import (
	userAccountDomain "github.com/ESPOIR-DITE/tim-api/domain/user/user.account.domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserAccountRepository struct {
	GormDB *gorm.DB
}

func NewUserAccountRepository(gormDB *gorm.DB) *UserAccountRepository {
	return &UserAccountRepository{
		GormDB: gormDB,
	}
}

func (uap *UserAccountRepository) CreateUserAccount(entity userAccountDomain.UserAccount) (*userAccountDomain.UserAccount, error) {
	var tableData = &userAccountDomain.UserAccount{}
	id := "UA-" + uuid.New().String()
	entity.Id = id
	if err := uap.GormDB.Create(entity).First(&tableData).Error; err != nil {
		return nil, err
	}
	return tableData, nil
}

func (uap *UserAccountRepository) UpdateUserAccount(entity userAccountDomain.UserAccount) (*userAccountDomain.UserAccount, error) {
	var tableData = &userAccountDomain.UserAccount{}
	if err := uap.GormDB.Updates(entity).First(&tableData).Error; err != nil {
		return nil, err
	}
	return tableData, nil
}

func (uap *UserAccountRepository) UpdateToken(entity userAccountDomain.UserAccount, token string) (userAccountDomain.UserAccount, error) {
	var tableData = userAccountDomain.UserAccount{}
	err := uap.GormDB.Model(&tableData).Where("user_id = ?", entity.UserId).Update("token", token).Error
	if err != nil {
		return entity, err
	}
	return entity, nil
}

func (uap *UserAccountRepository) GetUserAccountWithAccountIdAndUserId(userId, accountId string) (*userAccountDomain.UserAccount, error) {
	entity := &userAccountDomain.UserAccount{}
	if err := uap.GormDB.Where("user_id = ?", userId).Where("account_id = ? ", accountId).First(&entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func (uap *UserAccountRepository) GetUserAccountWithUserDetail(userDetailId string) (*userAccountDomain.UserAccount, error) {
	entity := &userAccountDomain.UserAccount{}
	if err := uap.GormDB.Where("user_detail_id = ?", userDetailId).First(&entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func (uap *UserAccountRepository) GetUserAccountWithAccountId(accountId string) (*userAccountDomain.UserAccount, error) {
	entity := &userAccountDomain.UserAccount{}
	if err := uap.GormDB.Where("account_id = ?", accountId).Find(&entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func (uap *UserAccountRepository) GetUserAccountWithUserId(userId string) ([]userAccountDomain.UserAccount, error) {
	entity := []userAccountDomain.UserAccount{}
	if err := uap.GormDB.Where("user_id = ?", userId).Find(&entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}
func (uap *UserAccountRepository) GetUserAccounts() ([]userAccountDomain.UserAccount, error) {
	var entity []userAccountDomain.UserAccount
	if err := uap.GormDB.Find(&entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}
func (uap *UserAccountRepository) DeleteUserAccount(id string) (bool, error) {
	entity := userAccountDomain.UserAccount{}
	if err := uap.GormDB.Where("id = ?", id).Delete(&entity).Error; err != nil {
		return false, err
	}
	if entity.Id == "" {
		return true, nil
	}
	return false, nil
}
func (uap *UserAccountRepository) countUserAccount() (*int64, error) {
	var value int64
	if err := uap.GormDB.Table("userAccountDomain").Count(&value).Error; err != nil {
		return nil, err
	}
	return &value, nil
}
