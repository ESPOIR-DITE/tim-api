package userSubscriptionRepository

import (
	userSubscriptionDomain "github.com/ESPOIR-DITE/tim-api/domain/user/user.subscription.domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserSubscriptionRepository struct {
	GormDB *gorm.DB
}

func NewUserSubscriptionRepository(gormDb *gorm.DB) *UserSubscriptionRepository {
	return &UserSubscriptionRepository{
		GormDB: gormDb,
	}
}

func (usr UserSubscriptionRepository) CreateUserSubscription(entity userSubscriptionDomain.UserSubscription) (*userSubscriptionDomain.UserSubscription, error) {
	var tableData = &userSubscriptionDomain.UserSubscription{}
	id := "R-" + uuid.New().String()
	entity.Id = id
	if err := usr.GormDB.Create(entity).First(&tableData).Error; err != nil {
		return nil, err
	}
	return tableData, nil
}
func (usr UserSubscriptionRepository) UpdateUserSubscription(entity userSubscriptionDomain.UserSubscription) (*userSubscriptionDomain.UserSubscription, error) {
	var tableData = &userSubscriptionDomain.UserSubscription{}
	if err := usr.GormDB.Create(entity).First(&tableData).Error; err != nil {
		return nil, err
	}
	return tableData, nil
}
func (usr UserSubscriptionRepository) GetUserSubscription(customerId string) (*userSubscriptionDomain.UserSubscription, error) {
	entity := &userSubscriptionDomain.UserSubscription{}
	if err := usr.GormDB.Where("id = ?", customerId).Find(entity).Error; err != nil {
		return entity, err
	}
	return entity, nil
}
func (usr UserSubscriptionRepository) GetUserSubscriptions() ([]userSubscriptionDomain.UserSubscription, error) {
	var entity []userSubscriptionDomain.UserSubscription
	if err := usr.GormDB.Find(&entity).Error; err != nil {
		return entity, err
	}
	return entity, nil
}
func (usr UserSubscriptionRepository) DeleteUserSubscription(id string) (bool, error) {
	entity := &userSubscriptionDomain.UserSubscription{}
	if err := usr.GormDB.Where("id = ?", id).Delete(entity).Error; err != nil {
		return false, err
	}
	if entity.Id == "" {
		return true, nil
	}
	return false, nil
}
