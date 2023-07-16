package userRepository

import (
	"github.com/ESPOIR-DITE/tim-api/api"
	userDomain "github.com/ESPOIR-DITE/tim-api/domain/user/user.domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	GormDB *gorm.DB
}

func NewUserRepository(gormDB *gorm.DB) *UserRepository {
	return &UserRepository{
		GormDB: gormDB,
	}
}

func (ur *UserRepository) CreateUser(entity userDomain.User) (*userDomain.User, error) {
	if err := ur.GormDB.Create(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (ur *UserRepository) UpdateUser(entity userDomain.User) (*userDomain.User, error) {
	var tableData = &userDomain.User{}
	if err := ur.GormDB.Updates(entity).First(tableData).Error; err != nil {
		return nil, err
	}
	return tableData, nil
}

func (ur *UserRepository) GetUser(email string) (*userDomain.User, error) {
	entity := &userDomain.User{}
	if err := ur.GormDB.Where("email = ?", email).First(entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func (ur *UserRepository) GetAllAgents() ([]userDomain.User, error) {
	entity := []userDomain.User{}
	if err := ur.GormDB.Where("role_id = ?", api.AgentId).Find(&entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func (ur *UserRepository) GetAllSuperAdmins() ([]userDomain.User, error) {
	entity := []userDomain.User{}
	if err := ur.GormDB.Where("role_id = ?", api.SuperAdminId).Find(&entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

//	func (ur *UserRepository) GetUserStack() user.home.controller.domain.controller.UserStack {
//		var superAdmin int64
//		var admin int64
//		var agent int64
//		ur.GormDB.Table("users").Where("role_id = ?", api.SuperAdminId).Count(&superAdmin)
//		ur.GormDB.Table("users").Where("role_id = ?", api.AdminId).Count(&admin)
//		ur.GormDB.Table("users").Where("role_id = ?", api.AgentId).Count(&agent)
//		return domain.UserStack{superAdmin, admin, agent}
//	}
func (ur *UserRepository) GetAllUsersByRole(roleId string) ([]userDomain.User, error) {
	entity := []userDomain.User{}
	if err := ur.GormDB.Where("role_id = ?", roleId).Find(&entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func (ur *UserRepository) GetAllAdmins() ([]userDomain.User, error) {
	entity := []userDomain.User{}
	if err := ur.GormDB.Where("role_id = ?", api.AdminId).Find(&entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func (ur *UserRepository) CountUsers() (*int64, error) {
	var value int64
	if err := ur.GormDB.Table("users").Count(&value).Error; err != nil {
		return nil, err
	}
	return &value, nil
}

func (ur *UserRepository) GetUsers() ([]userDomain.User, error) {
	entity := []userDomain.User{}
	if err := ur.GormDB.Find(&entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func (ur *UserRepository) Delete(email string) (bool, error) {
	entity := userDomain.User{}
	if err := ur.GormDB.Where("email = ?", email).Delete(&entity).Error; err != nil {
		return false, err
	}
	if entity.Email == "" {
		return true, nil
	}
	return false, nil
}
