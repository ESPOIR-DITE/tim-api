package userBankRepository

import (
	userDetailsDomain "github.com/ESPOIR-DITE/tim-api/domain/user/user.details.domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserBankRepository struct {
	GormDB *gorm.DB
}

func NewUserBankRepository(gormDB *gorm.DB) *UserBankRepository {
	return &UserBankRepository{
		GormDB: gormDB,
	}
}

func (ubr *UserBankRepository) CreateUserDetails(entity userDetailsDomain.UserBank) (*userDetailsDomain.UserBank, error) {
	var tableData = &userDetailsDomain.UserBank{}
	id := "UB-" + uuid.New().String()
	entity.Id = id
	if err := ubr.GormDB.Create(entity).First(&tableData).Error; err != nil {
		return nil, err
	}
	return tableData, nil
}

func (ubr *UserBankRepository) UpdateUserBank(entity userDetailsDomain.UserBank) (*userDetailsDomain.UserBank, error) {
	var tableData = &userDetailsDomain.UserBank{}
	if err := ubr.GormDB.Updates(entity).Where("id = ?", entity.Id).First(&tableData).Error; err != nil {
		return nil, err
	}
	return tableData, nil
}

func (ubr *UserBankRepository) GetUserBank(id string) (*userDetailsDomain.UserBank, error) {
	entity := &userDetailsDomain.UserBank{}
	if err := ubr.GormDB.Where("id = ?", id).Find(&entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func (ubr *UserBankRepository) GetUserBanks() ([]userDetailsDomain.UserBank, error) {
	entity := []userDetailsDomain.UserBank{}
	if err := ubr.GormDB.Find(&entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func (ubr *UserBankRepository) DeleteUserBankById(id string) (bool, error) {
	entity := userDetailsDomain.UserBank{}
	if err := ubr.GormDB.Where("id = ?", id).Delete(&entity).Error; err != nil {
		return false, err
	}
	if entity.Id == "" {
		return true, nil
	}
	return false, nil
}

func (ubr *UserBankRepository) IsExistUserBankById(id string) (bool, error) {
	entity := userDetailsDomain.UserBank{}
	if err := ubr.GormDB.Where("id = ?", id).Find(&entity).Error; err != nil {
		return false, err
	}
	if entity.Id != "" {
		return true, nil
	}
	return false, nil
}
