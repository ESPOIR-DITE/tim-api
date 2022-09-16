package video_related

import (
	"fmt"
	"testing"
	video_related "tim-api/domain/video/video-related"
)

func TestCreateVideoRelatedTable(t *testing.T) {
	result := CreateVideoRelatedTable()
	fmt.Println(result)
}
func TestCreateVideoRelated(t *testing.T) {
	object := video_related.VideoRelated{"jjendnnd", "348rwere"}
	result, err := CreateVideoRelated(object)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
func TestGetVideosRelatedTo(t *testing.T) {
	result, err := GetVideosRelatedTo("348rwere")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
func TestDeleteVideoRelated(t *testing.T) {
	object := video_related.VideoRelated{"jjendnnd", "348rwere"}
	result, err := DeleteVideoRelated(object)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
