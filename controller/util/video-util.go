package util

import (
	"fmt"
	"github.com/ItsJimi/gif/pkg/convert"
	"github.com/mowshon/moviego"
	"os"
	"os/exec"
)

func GetGif() {
	options := convert.Options{
		FPS:   30,
		Scale: -1,
		Crop:  "",
	}

	err := convert.FromFolder("../../files/test/upload-196722650.mp4", "../../files/gifs", options)
	if err != nil {
		fmt.Println(err)
	}
}

func GetGif2() {
	cmd := "ffmpeg -i files/test/upload-196722650.mp4 -vf \"fps=5,scale=320:-1:flags=lanczos\" -c:v pam -f image2pipe - | convert -delay 5 - -loop 0 -layers optimize test.gif"
	_, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		fmt.Println(fmt.Sprintf("Failed to execute command: %s", cmd))
	}
}
func VideoWriter(id string, file []byte, extension string) {
	tempFile, err := os.Create("files/test/" + id + "." + extension)
	if err != nil {
		fmt.Println(err)
	}
	_, err = tempFile.Write(file)
	if err != nil {
		fmt.Println(err, "fail to Upload File")
	}
	// return that we have successfully uploaded our file!
	fmt.Println("Successfully Wrote video")
	ScreenShort(id, extension)
}

//Tested working
func ScreenShort(id, extension string) {
	first, err := moviego.Load("files/test/" + id + "." + extension)
	if err != nil {
		fmt.Println(err, "error ScreenShorting")
	}
	result, err := first.Screenshot(5, "files/pictures/"+id+".png")
	if err != nil {
		fmt.Println(err, "error")
	} else {
		fmt.Println(result)
	}
}
func GetVideoPictures(id string) ([]byte, error) {
	dat, err := os.ReadFile("files/pictures/" + id + ".png")
	if err != nil {
		fmt.Println(err, "error reading file")
	}
	return dat, err
}
