package channelTypeRepository

import (
	channel_type "github.com/ESPOIR-DITE/tim-api/domain/channel/channel-type"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ChannelTypeRepository struct {
	GormDb *gorm.DB
}

func NewChannelTypeRepository(gormDb *gorm.DB) *ChannelTypeRepository {
	return &ChannelTypeRepository{
		GormDb: gormDb,
	}
}

func (c *ChannelTypeRepository) CreateChannelType(channelType channel_type.ChannelType) (*channel_type.ChannelType, error) {
	var channelTypeObject = &channel_type.ChannelType{}
	id := "CT-" + uuid.New().String()
	channelType.Id = id
	err := c.GormDb.Create(channelType).Find(channelTypeObject).Error
	if err != nil {
		return channelTypeObject, err
	}
	return channelTypeObject, nil
}

func (c *ChannelTypeRepository) UpdateChannelType(entity channel_type.ChannelType) (*channel_type.ChannelType, error) {
	var tableData = &channel_type.ChannelType{}
	if err := c.GormDb.Where("id = ", entity.Id).Updates(entity).Find(&tableData).Error; err != nil {
		return nil, err
	}
	return tableData, nil
}

func (c *ChannelTypeRepository) GetChannelType(roleId string) (*channel_type.ChannelType, error) {
	entity := &channel_type.ChannelType{}
	if err := c.GormDb.Where("id = ?", roleId).Find(&entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func (c *ChannelTypeRepository) GetChannelTypes() ([]channel_type.ChannelType, error) {
	var entity []channel_type.ChannelType
	if err := c.GormDb.Find(&entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func (c *ChannelTypeRepository) DeleteChannelType(id string) (bool, error) {
	entity := channel_type.ChannelType{}
	if err := c.GormDb.Where("id = ?", id).Delete(&entity).Error; err != nil {
		return false, err
	}
	if entity.Id == "" {
		return true, nil
	}
	return false, nil
}
