package util

import (
	"errors"
	"fmt"
	"github.com/mowshon/moviego"
	"image"
	"image/png"
	"io/ioutil"
	"os"
	"tim-api/domain"
	video_data "tim-api/storage/video/video-data"
	"time"
)

const videoURL = "files/videos/"
const pictureURL = "files/pictures/"
const pictureReserveURL = "files/test/"

func ReadVideoFile(id, extension string) ([]byte, error) {
	fileBytes, err := ioutil.ReadFile(videoURL + id + "." + extension)
	if err != nil {
		fmt.Println(err, "error reading video file.")
		return []byte{}, err
	}
	return fileBytes, nil
}

func VideoWriter(date domain.VideoData, isUpdate bool) {
	tempFile, err := os.Create(videoURL + date.Id + "." + date.FileType)
	if err != nil {
		fmt.Println(err, " Creating video temp file")
	}
	//file, err := os.ReadFile("1234.mp4")
	//_, err = tempFile.Write(file)
	_, err = tempFile.Write(date.Video)
	if err != nil {
		fmt.Println(err, "fail to write File")
	} else {
		fmt.Println("Successfully Wrote video")
		time.Sleep(1 * time.Minute)
		ScreenShort(date.Id, date.FileType)
		time.Sleep(1 * time.Minute)
		err := CropImage(date.Id + ".png")
		if err != nil {
			fmt.Println(err, " error cropping")
		}
		err = savePicture(date, isUpdate)
		if err != nil {
			fmt.Println(err, " error saving image")
		}
	}
}

// Tested working
func ScreenShort(id, extension string) {
	fmt.Println("opening video for screenshotting...")
	fmt.Println("id: ", id)
	fmt.Println("extension: ", extension)
	first, err := moviego.Load(videoURL + id + "." + extension)
	//first, err := moviego.Load(id + "." + extension)
	if err != nil {
		fmt.Println(err, "error ScreenShorting")
	}
	result, err := first.Screenshot(5, pictureURL+id+".png")
	if err != nil {
		fmt.Println(err, "error")
	} else {
		fmt.Println("Result: ", result)
	}
	fmt.Println("screenshot successfully")
}

func CropImage(imageURL string) error {
	img, err := ReadPngImage(pictureURL + imageURL)
	if err != nil {
		fmt.Println(err, " error reading image for cropping.")
		return err
	}
	//err = os.Remove(imageURL)
	//if err != nil {
	//	fmt.Println(err, " error deleting image")
	//}
	info := img.Bounds().Size()

	xcenter := info.X/2 + 200
	ycenter := info.Y/2 + 200
	//fmt.Println("picture center: x= ", xcenter, " y= ", ycenter)
	//I've hard-coded a crop rectangle, start (0,0), end (100, 100).
	//img, err = cropImage(img, image.Rect(0, 0, 100, 100))
	img, err = cropIt(img, image.Rect(ycenter, xcenter, 200, 200))
	if err != nil {
		fmt.Println(err, " error cutting image.")
		return err
	}

	return writePngImage(img, imageURL)
}

// readImage reads a
func savePicture(data domain.VideoData, isUpdate bool) error {
	picture, err := os.ReadFile(pictureURL + data.Id + ".png")
	if err != nil {
		fmt.Println(err, " could not read picture!")
		return err
	} else {
		errRemove := os.Remove(pictureURL + data.Id + ".png")
		if errRemove != nil {
			fmt.Println(errRemove, " could not Remove picture!")
		}
		videoDataObject := domain.VideoData{data.Id, picture, []byte{}, data.FileType, data.FileSize}
		if isUpdate {
			result := video_data.UpdateVideoDate(videoDataObject)
			if result.Id == "" {
				fmt.Println("error update videoData")
				return errors.New("Could not update videoData!")
			}
		}
		result := video_data.CreateVideoData(videoDataObject)
		if result.Id == "" {
			fmt.Println("error creating videoData")
			return errors.New("Could not create videoData!")
		}
	}
	return nil
}

func GetVideoPictures(id string) ([]byte, error) {
	dat, err := os.ReadFile("files/pictures/" + id + ".png")
	if err != nil {
		fmt.Println(err, "error reading file")
	}
	return dat, err
}

// readImage reads a image file from disk. We're assuming the file will be png
// format.
func ReadPngImage(url string) (image.Image, error) {
	fd, err := os.Open(url)
	if err != nil {
		return nil, err
	}
	defer fd.Close()

	// image.Decode requires that you import the right image package. We've
	// imported "image/png", so Decode will work for png files. If we needed to
	// decode jpeg files then we would need to import "image/jpeg".
	//
	// Ignored return value is image format name.
	img, _, err := image.Decode(fd)
	if err != nil {
		return nil, err
	}

	return img, nil
}

// cropImage takes an image and crops it to the specified rectangle.
func cropIt(img image.Image, crop image.Rectangle) (image.Image, error) {
	type subImager interface {
		SubImage(r image.Rectangle) image.Image
	}
	// img is an Image interface. This checks if the underlying value has a
	// method called SubImage. If it does, then we can use SubImage to crop the
	// image.
	simg, ok := img.(subImager)
	if !ok {
		return nil, fmt.Errorf("image does not support cropping")
	}

	return simg.SubImage(crop), nil
}

// writeImage writes an Image back to the disk.
func writePngImage(img image.Image, name string) error {
	fd, err := os.Create(pictureURL + name)
	if err != nil {
		fmt.Println(err, "error creating cropped temp file")
		return err
	}
	defer fd.Close()

	return png.Encode(fd, img)
}
