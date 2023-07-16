package roleRepository

import (
	roleDomain "github.com/ESPOIR-DITE/tim-api/domain/user/role.domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoleRepository struct {
	GormDB *gorm.DB
}

func NewRoleRepository(gormDB *gorm.DB) *RoleRepository {
	return &RoleRepository{
		GormDB: gormDB,
	}
}

func (rp *RoleRepository) CreateRole(entity roleDomain.Role) (*roleDomain.Role, error) {
	id := "R-" + uuid.New().String()
	entity.Id = id
	err := rp.GormDB.Create(&entity).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (rp *RoleRepository) UpdateRole(entity roleDomain.Role) (*roleDomain.Role, error) {
	var tableData = &roleDomain.Role{}
	if err := rp.GormDB.Create(entity).First(&tableData).Error; err != nil {
		return nil, err
	}
	return tableData, nil
}

func (rp *RoleRepository) GetRole(roleId string) (*roleDomain.Role, error) {
	entity := &roleDomain.Role{}
	if err := rp.GormDB.Where("id = ?", roleId).First(&entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}
func (rp *RoleRepository) GetRoles() ([]roleDomain.Role, error) {
	var entity []roleDomain.Role
	if err := rp.GormDB.Find(&entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}
func (rp *RoleRepository) DeleteRole(id string) (bool, error) {
	entity := &roleDomain.Role{}
	if err := rp.GormDB.Where("id = ?", id).Delete(&entity).Error; err != nil {
		return false, err
	}
	if entity.Id == "" {
		return true, nil
	}
	return false, nil
}
