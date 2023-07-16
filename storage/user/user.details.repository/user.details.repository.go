package userDetailsRepository

import (
	userDetailsDomain "github.com/ESPOIR-DITE/tim-api/domain/user/user.details.domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserDetailsRepository struct {
	GormDB *gorm.DB
}

func NewUserDetailsRepository(gormDB *gorm.DB) *UserDetailsRepository {
	return &UserDetailsRepository{
		GormDB: gormDB,
	}
}

func (udr *UserDetailsRepository) CreateUserDetails(entity userDetailsDomain.AccountDetails) (*userDetailsDomain.AccountDetails, error) {
	var tableData = &userDetailsDomain.AccountDetails{}
	id := "UD-" + uuid.New().String()
	entity.Id = id
	if err := udr.GormDB.Create(entity).First(&tableData).Error; err != nil {
		return nil, err
	}
	return tableData, nil
}

func (udr *UserDetailsRepository) UpdateUserDetails(entity userDetailsDomain.AccountDetails) (*userDetailsDomain.AccountDetails, error) {
	var tableData = &userDetailsDomain.AccountDetails{}
	if err := udr.GormDB.Updates(entity).First(&tableData).Error; err != nil {
		return nil, err
	}
	return tableData, nil
}

func (udr UserDetailsRepository) GetUserDetail(email string) (*userDetailsDomain.AccountDetails, error) {
	entity := &userDetailsDomain.AccountDetails{}
	if err := udr.GormDB.Where("id = ?", email).First(&entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func (udr UserDetailsRepository) GetUserDetails() ([]userDetailsDomain.AccountDetails, error) {
	entity := []userDetailsDomain.AccountDetails{}
	if err := udr.GormDB.Find(&entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func (udr UserDetailsRepository) DeleteUserDetailsById(id string) (bool, error) {
	entity := userDetailsDomain.AccountDetails{}
	if err := udr.GormDB.Where("id = ?", id).Delete(&entity).Error; err != nil {
		return false, err
	}
	if entity.Id == "" {
		return true, nil
	}
	return false, nil
}

func (udr UserDetailsRepository) IsExistUserDetailById(id string) (bool, error) {
	entity := userDetailsDomain.AccountDetails{}
	if err := udr.GormDB.Where("id = ?", id).Find(&entity).Error; err != nil {
		return false, err
	}
	if entity.Id != "" {
		return true, nil
	}
	return false, nil
}

func (udr UserDetailsRepository) IsUserExistWithBankId(bankId string) (bool, error) {
	entity := userDetailsDomain.AccountDetails{}
	err := udr.GormDB.Where("bank_id = ?", bankId).First(&entity).Error
	if err != nil {
		return false, err
	}
	if entity.Id != "" {
		return true, nil
	}
	return false, nil
}
