package categoryRepository

import (
	"github.com/ESPOIR-DITE/tim-api/domain/video/category"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	GormDB *gorm.DB
}

func NewCategoryRepository(gormDB *gorm.DB) *CategoryRepository {
	return &CategoryRepository{
		GormDB: gormDB,
	}
}

func (cr CategoryRepository) CreateCategory(entity category.Category) (*category.Category, error) {
	var tableData = &category.Category{}
	id := "C-" + uuid.New().String()
	entity.Id = id
	if err := cr.GormDB.Create(entity).First(&tableData).Error; err != nil {
		return nil, err
	}
	return tableData, nil
}
func (cr CategoryRepository) UpdateCategory(entity category.Category) (*category.Category, error) {
	var tableData = &category.Category{}
	if err := cr.GormDB.Create(entity).Find(&tableData).Error; err != nil {
		return nil, err
	}
	return tableData, nil
}
func (cr CategoryRepository) GetCategory(id string) (*category.Category, error) {
	entity := &category.Category{}
	if err := cr.GormDB.Where("id = ?", id).First(entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}
func (cr CategoryRepository) GetCategories() ([]category.Category, error) {
	entity := []category.Category{}
	if err := cr.GormDB.Find(&entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}
func (cr CategoryRepository) DeleteCategory(id string) (bool, error) {
	entity := &category.Category{}
	if err := cr.GormDB.Where("id = ?", id).Delete(&entity).Error; err != nil {
		return false, err
	}
	if entity.Id == "" {
		return true, nil
	}
	return false, nil
}
