package util

import (
	"testing"
	"tim-api/domain"
)

func TestGetGif(t *testing.T) {
	ScreenShort("V-162609bc-07a0-41fa-83f0-15705f816167", ".mp4")
}
func TestScaleDown(t *testing.T) {
	videoData := domain.VideoData{"1234", []byte{}, []byte{}, "mp4", "19000"}
	VideoWriter(videoData)
}
func TestCropImage(t *testing.T) {
	CropImage("test.png")
}
