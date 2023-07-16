package channelSubscriptionRepository

import (
	channelSubscriptionDomain "github.com/ESPOIR-DITE/tim-api/domain/channel/channel.subscription.domain"
	"gorm.io/gorm"
)

type ChannelSubscriptionRepository struct {
	GormDB *gorm.DB
}

func NewChannelSubscriptionRepository(gormDb *gorm.DB) *ChannelSubscriptionRepository {
	return &ChannelSubscriptionRepository{
		GormDB: gormDb,
	}
}
func (cs *ChannelSubscriptionRepository) CreateChannelSubscription(channel channelSubscriptionDomain.ChannelSubscription) (*channelSubscriptionDomain.ChannelSubscription, error) {
	var tableData = &channelSubscriptionDomain.ChannelSubscription{}
	if err := cs.GormDB.Create(channel).First(&tableData).Error; err != nil {
		return nil, err
	}
	return tableData, nil
}

func (cs *ChannelSubscriptionRepository) UpdateChannelSubscription(entity channelSubscriptionDomain.ChannelSubscription) (*channelSubscriptionDomain.ChannelSubscription, error) {
	var tableData = &channelSubscriptionDomain.ChannelSubscription{}
	if err := cs.GormDB.Where("channel_id = ?", entity.ChannelId).Where("user_id = ?", entity.UserId).Updates(entity).First(&tableData).Error; err != nil {
		return nil, err
	}
	return tableData, nil
}

func (cs *ChannelSubscriptionRepository) GetChannelSubscription(roleId string) (*channelSubscriptionDomain.ChannelSubscription, error) {
	entity := &channelSubscriptionDomain.ChannelSubscription{}
	if err := cs.GormDB.Where("id = ?", roleId).First(&entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func (cs *ChannelSubscriptionRepository) GetChannelSubscriptionsByUser(userId string) ([]channelSubscriptionDomain.ChannelSubscription, error) {
	entity := []channelSubscriptionDomain.ChannelSubscription{}
	if err := cs.GormDB.Where("user_id = ?", userId).First(&entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func (cs *ChannelSubscriptionRepository) GetChannelSubscriptionsByChannelId(channelId string) ([]channelSubscriptionDomain.ChannelSubscription, error) {
	entity := []channelSubscriptionDomain.ChannelSubscription{}
	if err := cs.GormDB.Where("channel_id = ?", channelId).First(&entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func (cs *ChannelSubscriptionRepository) GetChannelSubscriptions() ([]channelSubscriptionDomain.ChannelSubscription, error) {
	var entity []channelSubscriptionDomain.ChannelSubscription
	if err := cs.GormDB.First(&entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}
func (cs *ChannelSubscriptionRepository) DeleteChannelSubscription(id string) (bool, error) {
	entity := channelSubscriptionDomain.ChannelSubscription{}
	if err := cs.GormDB.Where("id = ?", id).Delete(&entity).Error; err != nil {
		return false, err
	}
	if entity.ChannelId == "" {
		return true, nil
	}
	return false, nil
}

func (cs *ChannelSubscriptionRepository) CountSubscriptionByChannelId(channelId string) (*int64, error) {
	var value int64
	if err := cs.GormDB.Table("channel_subscription").Where("channel_id = ?", channelId).Count(&value).Error; err != nil {
		return nil, err
	}
	return &value, nil
}
