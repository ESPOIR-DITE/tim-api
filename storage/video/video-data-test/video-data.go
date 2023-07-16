package video_data_test

//var connection = config.GetDatabase()
//
//func CreateVideoDataTestTable() bool {
//	var tableData = &domain.VideoDataTest{}
//	err := connection.AutoMigrate(tableData)
//	if err != nil {
//		return false
//	}
//	return true
//}
//func SetDatabase() {
//	erro := connection.Set("gorm:table_options", "ENGINE=Distributed(cluster, default, hits)").AutoMigrate(&domain.VideoDataTest{})
//	if erro != nil {
//		fmt.Println("Category database config not set")
//	} else {
//		fmt.Println("Category database config set successfully")
//	}
//}
//func CreateVideoData(entity domain.VideoDataTest) *domain.VideoDataTest {
//	var tableData = &domain.VideoDataTest{}
//	//id := "C-" + uuid.New().String()
//	user.home.controller.domain.controller := domain.VideoDataTest{entity.Id, entity.Picture, entity.Video, entity.FileType, entity.FileSize}
//	connection.Create(user.home.controller.domain.controller).Find(&tableData)
//	return tableData
//}
//func UpdateVideoDate(entity domain.VideoDataTest) domain.VideoDataTest {
//	var tableData = &domain.VideoDataTest{}
//	connection.Updates(entity).Find(&tableData)
//	return entity
//}
//func GetVideoDate(id string) domain.VideoDataTest {
//	entity := domain.VideoDataTest{}
//	connection.Where("id = ?", id).Find(&entity)
//	return entity
//}
//func GetVideoDatas() []domain.VideoDataTest {
//	entity := []domain.VideoDataTest{}
//	connection.Find(&entity)
//	return entity
//}
//func DeleteVideoData(email string) bool {
//	entity := domain.VideoDataTest{}
//	connection.Where("id = ?", email).Delete(&entity)
//	if entity.Id == "" {
//		return true
//	}
//	return false
//}
//func GetVideoDataObject(entity *domain.VideoDataTest) domain.VideoDataTest {
//	return domain.VideoDataTest{entity.Id, entity.Picture, entity.Video, entity.FileType, entity.FileSize}
//}
