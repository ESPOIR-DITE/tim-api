package util

import (
	"testing"
)

func TestGetGif(t *testing.T) {
	ScreenShort("V-0218470b-c41b-4e3d-b0c5-7b019784c0c3", "mp4")
}
func TestScaleDown(t *testing.T) {
	//videoData := domain.VideoData{"1234", []byte{}, []byte{}, "mp4", "19000"}
	//VideoWriter(videoData)
}
func TestCropImage(t *testing.T) {
	CropImage("test.png")
}
