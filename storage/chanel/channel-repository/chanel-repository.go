package role_repo

import (
	"github.com/ESPOIR-DITE/tim-api/domain/channel/channel"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ChanelRepository struct {
	GormDB *gorm.DB
}

func NewChanelRepository(gormDB *gorm.DB) *ChanelRepository {
	return &ChanelRepository{
		GormDB: gormDB,
	}
}

func (db *ChanelRepository) CreateChannel(channelObject channel.Channel) *channel.Channel {
	var tableData = &channel.Channel{}
	id := "C-" + uuid.New().String()
	//user.home.controller.domain.controller := domain.Channel{id, channel.controller.Name, channel.controller.ChannelTypeId, channel.controller.UserId, channel.controller.Region, channel.controller.Date, channel.controller.Image, "", channel.controller.Description}
	channelObject.Id = id
	db.GormDB.Create(channelObject).Find(&tableData)
	return tableData
}
func (db *ChanelRepository) UpdateChannel(entity channel.Channel) *channel.Channel {
	var tableData = &channel.Channel{}
	db.GormDB.Where("id = ", entity.Id).Updates(entity).Find(&tableData)
	return tableData
}
func (db *ChanelRepository) GetChannel(roleId string) *channel.Channel {
	entity := &channel.Channel{}
	db.GormDB.Where("id = ?", roleId).Find(&entity)
	return entity
}
func (db *ChanelRepository) GetChannelsByUser(userId string) []channel.Channel {
	entity := []channel.Channel{}
	db.GormDB.Where("user_id = ?", userId).Find(&entity)
	return entity
}
func (db *ChanelRepository) GetChannelsByRegion(region string) []channel.Channel {
	entity := []channel.Channel{}
	db.GormDB.Where("region = ?", region).Find(&entity)
	return entity
}
func (db *ChanelRepository) GetChannels() []channel.Channel {
	var entity []channel.Channel
	db.GormDB.Find(&entity)
	return entity
}
func (db *ChanelRepository) DeleteChannel(id string) bool {
	entity := channel.Channel{}
	db.GormDB.Where("id = ?", id).Delete(&entity)
	if entity.Id == "" {
		return true
	}
	return false
}
func (db *ChanelRepository) CountChannel() *int64 {
	var value int64
	db.GormDB.Table("channels").Count(&value)
	return &value
}
