package util

import (
	"errors"
	"fmt"
	server_config "github.com/ESPOIR-DITE/tim-api/config/server.config"
	videoDataDomain "github.com/ESPOIR-DITE/tim-api/domain/video/video.data.domain"
	videodataRepository "github.com/ESPOIR-DITE/tim-api/storage/video/video.data.repository"
	"github.com/mowshon/moviego"
	log "github.com/sirupsen/logrus"
	"image"
	"image/png"
	"io/ioutil"
	"os"
	"time"
)

const pictureURL = "files/pictures/"

//const pictureURL = "files/pictures/"

const videoURL = "files/videos/"

//const videoURL = "videos/"

func ReadVideoFile(id, extension string) ([]byte, error) {
	fileBytes, err := ioutil.ReadFile(videoURL + id + "." + extension)
	if err != nil {
		fmt.Println(err, "error reading video file.")
		return []byte{}, err
	}
	return fileBytes, nil
}

func VideoWriter(app *server_config.Env, date videoDataDomain.VideoData, isUpdate bool, repo *videodataRepository.VideoDataRepository) {
	tempFile, err := os.Create(videoURL + date.Id + "." + date.FileType)
	if err != nil {
		fmt.Println(err, " Creating video temp file")
	}
	_, err = tempFile.Write(date.Video)
	if err != nil {
		fmt.Println(err, "fail to write File")
	} else {
		fmt.Println("Successfully Wrote video.controller")
		time.Sleep(1 * time.Minute)
		err := ScreenShort(date.Id, date.FileType)
		if err == nil {
			time.Sleep(1 * time.Minute)
			err := CropImage(date.Id + ".png")
			if err != nil {
				fmt.Println(err, " error cropping")
			}
		}

		err = savePicture(date, isUpdate, repo)
		if err != nil {
			fmt.Println(err, " error saving image")
		}
	}
}

func VideoWriterWithS3(app *server_config.Env, date videoDataDomain.VideoData, isUpdate bool, repo *videodataRepository.VideoDataRepository) {
	tempFile, err := os.Create(videoURL + date.Id + "." + date.FileType)
	if err != nil {
		fmt.Println(err, " Creating video temp file")
	}
	_, err = tempFile.Write(date.Video)
	if err != nil {
		fmt.Println(err, "fail to write video File to S3")
		return
	}
	_, err = app.PhotoS3Bucket.AddJsonPayload("tim-api-videos", date.Id, date.Video)
	if err != nil {
		fmt.Println(err, "fail to write video File to S3")
	} else {
		fmt.Println("Successfully Wrote video.controller")
		time.Sleep(1 * time.Minute)
		err := ScreenShort(date.Id, date.FileType)
		if err == nil {
			time.Sleep(1 * time.Minute)
			err := CropImage(date.Id + ".png")
			if err != nil {
				fmt.Println(err, " error cropping")
			}
		}
		err = savePictureToS3(app, date, isUpdate, repo)
		if err != nil {
			fmt.Println(err, " error saving image")
		}
		err = os.Remove(videoURL + date.Id + "." + date.FileType)
		if err != nil {
			fmt.Println(err, " error saving image")
		}
	}
}

// Tested working
func ScreenShort(id, extension string) error {
	movie, err := moviego.Load(videoURL + id + "." + extension)
	if err != nil {
		fmt.Println(err, "error ScreenShorting")
		return err
	}
	_, err = movie.Screenshot(5, pictureURL+id+".png")
	if err != nil {
		fmt.Println(err, "error")
		return err
	}
	fmt.Println("screenshot successfully")
	return nil
}

// CropImage This crops an image removing extra size in the edge of
// the picture keeping the center portion of 200px X 200px.
func CropImage(imageURL string) error {
	img, err := ReadPngImage(pictureURL + imageURL)
	if err != nil {
		fmt.Println(err, " error reading image for cropping.")
		return err
	}
	size := img.Bounds().Size()
	xcenter := size.X/2 + 200
	ycenter := size.Y/2 + 200
	img, err = cropIt(img, image.Rect(ycenter, xcenter, 200, 200))
	if err != nil {
		fmt.Println(err, " error cutting image.")
		return err
	}

	return writePngImage(img, imageURL)
}

// savePicture this submits videoData data to the api.
func savePicture(data videoDataDomain.VideoData, isUpdate bool, repo *videodataRepository.VideoDataRepository) error {
	picture, err := os.ReadFile(pictureURL + data.Id + ".png")
	if err != nil {
		fmt.Println(err, " could not read picture!")
		return err
	} else {
		errRemove := os.Remove(pictureURL + data.Id + ".png")
		if errRemove != nil {
			fmt.Println(errRemove, " could not Remove picture!")
		}
		videoDataObject := videoDataDomain.VideoData{
			Id:       data.Id,
			Picture:  picture,
			Video:    []byte{},
			FileType: data.FileType,
			FileSize: data.FileSize,
		}
		if isUpdate {
			_, err := repo.UpdateVideoDate(videoDataObject)
			if err != nil {
				fmt.Println("error update videoData")
				return errors.New("Could not update videoData!")
			}
		}
		_, err := repo.CreateVideoData(videoDataObject)
		if err != nil {
			fmt.Println("error creating videoData")
			return errors.New("Could not create videoData!")
		}
	}
	return nil
}

func savePictureToS3(app *server_config.Env, data videoDataDomain.VideoData, isUpdate bool, repo *videodataRepository.VideoDataRepository) error {
	picture, err := os.ReadFile(pictureURL + data.Id + ".png")
	if err != nil {
		fmt.Println(err, " could not read picture!")
		return err
	} else {
		errRemove := os.Remove(pictureURL + data.Id + ".png")
		if errRemove != nil {
			fmt.Println(errRemove, " could not Remove picture!")
		}
		videoDataObject := videoDataDomain.VideoData{
			Id:       data.Id,
			Picture:  []byte{},
			Video:    []byte{},
			FileType: data.FileType,
			FileSize: data.FileSize,
		}
		if isUpdate {
			_, err := repo.UpdateVideoDate(videoDataObject)
			if err != nil {
				fmt.Println("error update videoData")
				return errors.New("Could not update videoData!")
			}
		}
		_, err := repo.CreateVideoData(videoDataObject)
		if err != nil {
			fmt.Println("error creating videoData")
			return errors.New("Could not create videoData!")
		}
		go func() {
			_, err := app.PhotoS3Bucket.AddJsonPayload("tim-api-photos", data.Id, picture)
			if err != nil {
				log.Errorf("picture upload fail in s3, error: %d", err)
			}
		}()
	}
	return nil
}

func GetVideoPictures(id string) ([]byte, error) {
	//dat, err := os.ReadFile("files/pictures/" + id + ".png")
	dat, err := os.ReadFile(pictureURL + id + ".png")
	if err != nil {
		fmt.Println(err, "error reading file")
	}
	return dat, err
}

func GetVideoPictureFromS3(app *server_config.Env, id string) ([]byte, error) {
	picture, err := app.PhotoS3Bucket.DownloadFile("tim-api-videos", id)
	if err != nil {
		fmt.Println(err, "error reading file")
	}
	return picture, err
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
