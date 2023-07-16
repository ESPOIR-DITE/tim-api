package videoDataTest

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	server_config "github.com/ESPOIR-DITE/tim-api/config/server.config"
	"github.com/ESPOIR-DITE/tim-api/controller/util"
	controller_auth "github.com/ESPOIR-DITE/tim-api/controller/util/controller-auth"
	userVideoDomain "github.com/ESPOIR-DITE/tim-api/domain/user/user.video.domain"
	"github.com/ESPOIR-DITE/tim-api/domain/video/video"
	videoDataDomain "github.com/ESPOIR-DITE/tim-api/domain/video/video.data.domain"
	"github.com/ESPOIR-DITE/tim-api/logger"
	userVideoRepository "github.com/ESPOIR-DITE/tim-api/storage/user/user.video.repository"
	videodataRepository "github.com/ESPOIR-DITE/tim-api/storage/video/video.data.repository"
	videoRepository "github.com/ESPOIR-DITE/tim-api/storage/video/video.repository"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	"net/http"
	"os"
	"time"
)

func Home(app *server_config.Env) http.Handler {
	r := chi.NewRouter()
	repo := videodataRepository.NewVideoDataRepository(app.GormDB)
	videoRepo := videoRepository.NewVideoRepository(app.GormDB)
	userVideo := userVideoRepository.NewUserVideoRepository(app.GormDB)
	//r.Get("/get/{id}", get(app))
	r.Get("/delete/{id}", delete(app, repo))
	r.Post("/create", create(app, repo, userVideo))
	//r.Post("/update-video.controller", updateVideo(app))
	//r.Post("/create-picture", createPicture(app))
	//r.Post("/update-picture", updatePicture(app))
	r.Post("/update", update(app, repo))
	r.Get("/getAll", getAll(app, repo))
	r.Get("/raw-video/{id}", getRaw(app, repo))
	r.Get("/raw-all-video/{videoId}/{accountId}", getAllRaw(app, repo))
	////r.Get("/video-picture/{id}", getVideoPicture(app))
	r.Get("/video-public-picture", getPublicVideoPicture(app, repo, videoRepo))
	r.Get("/get-my-videos-picture", getMyVideoPicture(app, repo, videoRepo, userVideo))
	//r.Get("/video.controller-picture", getPublicVideoPicture(app))
	r.Get("/download/{videoId}", downloadVideo(app))
	return r
}

//func updateVideo(app *config.Env) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		//Authenticating user.home.controller.domain.controller by checking the token.
//		token := jwtauth.TokenFromHeader(r)
//		controller_auth.IsAllowed(token, w, r)
//		data := &domain.VideoData{}
//		err := render.Bind(r, data)
//		if err != nil {
//			render.Render(w, r, util.ErrInvalidRequest(err))
//			return
//		}
//		object := repository.GetCategoryObject(data)
//		if len(object.Picture) == 0 {
//			fmt.Println(object.Id)
//		} else {
//			fmt.Println()
//		}
//		go util.VideoWriter(object, true)
//
//		_, err = w.Write([]byte("Done"))
//		if err != nil {
//			render.Render(w, r, util.ErrInvalidRequest(errors.New("writing bytes")))
//			return
//		}
//		return
//	}
//}

func downloadVideo(app *server_config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		videoId := chi.URLParam(r, "videoId")

		if videoId == "" {
			render.Render(w, r, util.ErrInvalidRequest(errors.New("missing required field")))
			return
		}
		//videoDate,err := repo.GetVideoDate(videoId)
		//if videoDate.Id == "" {
		//	fmt.Println("error reading Video")
		//	return
		//}
		//fileBytes, err := util.ReadVideoFile(videoId, videoDate.FileType)
		//if err != nil {
		//	fmt.Println(err)
		//	return
		//}

		picture, err := app.PhotoS3Bucket.DownloadFile("tim-api-videos", videoId)
		if err != nil {
			logger.Log.Errorf("fail to get a picture of video with id: %s", videoId)
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "video.controller/mp4")
		w.Write(picture)
		return
	}
}

//func updatePicture(app *config.Env) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		//Authenticating user.home.controller.domain.controller by checking the token.
//		token := jwtauth.TokenFromHeader(r)
//		controller_auth.IsAllowed(token, w, r)
//		data := &domain.VideoData{}
//		err := render.Bind(r, data)
//		if err != nil {
//			render.Render(w, r, util.ErrInvalidRequest(err))
//			return
//		}
//		object := repository.GetCategoryObject(data)
//		videoDataObject := repository.UpdateVideoDate(object)
//		result, err := json.Marshal(videoDataObject)
//		if err != nil {
//			fmt.Println(err, " error creating videoData")
//		}
//		_, err = w.Write([]byte(result))
//		if err != nil {
//			render.Render(w, r, util.ErrInvalidRequest(errors.New("writing bytes")))
//			return
//		}
//		return
//	}
//}

type VideoVideoData struct {
	Video     video.Video
	VideoData videoDataDomain.VideoData
}

//func createPicture(app *config.Env) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		//Authenticating user.home.controller.domain.controller by checking the token.
//		token := jwtauth.TokenFromHeader(r)
//		controller_auth.IsAllowed(token, w, r)
//		data := &domain.VideoData{}
//		err := render.Bind(r, data)
//		if err != nil {
//			render.Render(w, r, util.ErrInvalidRequest(err))
//			return
//		}
//		object := repository.GetCategoryObject(data)
//		videoDataObject := repository.CreateVideoData(object)
//		result, err := json.Marshal(videoDataObject)
//		if err != nil {
//			fmt.Println(err, " error creating videoData")
//		}
//		_, err = w.Write([]byte(result))
//		if err != nil {
//			render.Render(w, r, util.ErrInvalidRequest(errors.New("writing bytes")))
//			return
//		}
//		return
//	}
//}

func getMyVideoPicture(app *server_config.Env, repo *videodataRepository.VideoDataRepository, videoRepo *videoRepository.VideoRepository, userVideo *userVideoRepository.UserVideoRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := jwtauth.TokenFromHeader(r)
		id, err := controller_auth.IsAllowed(app, token)
		if err != nil || id == nil {
			render.Render(w, r, util.ErrRecourseNotAllowed(err))
			return
		}

		userVideos, err := userVideo.GetUserVideosWithUserId(*id)
		if err != nil {
			logger.Log.Errorf("fail to read all user videos error: %d", err)
		}

		var videoVideoData = []VideoVideoData{}
		//todo this aggregation should happen in database level.
		for _, userVideo := range userVideos {
			video, err := videoRepo.GetVideo(userVideo.VideoId)
			if err != nil {
				logger.Log.Errorf("fail to read videos with id: %s, error: %d", userVideo.Id, err)
				continue
			}
			picture, err := app.PhotoS3Bucket.DownloadFile("tim-api-photos", video.Id)
			if err != nil {
				logger.Log.Errorf("fail to get a picture of video with id: %s", video.Id)
				continue
			}
			videoData, err := repo.GetVideoDate(video.Id)
			if err == nil {
				videoData.Picture = picture
				videoVideoData = append(videoVideoData, VideoVideoData{*video, *videoData})
			}
		}

		result, err := json.Marshal(videoVideoData)
		if err != nil {
			fmt.Println("couldn't marshal")
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error marshalling")))
			return
		}
		_, err = w.Write([]byte(result))
		if err != nil {
			return
		}
	}
}

func getPublicVideoPicture(app *server_config.Env, repo *videodataRepository.VideoDataRepository, videoRepo *videoRepository.VideoRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := jwtauth.TokenFromHeader(r)
		_, err := controller_auth.IsAllowed(app, token)
		if err != nil {
			render.Render(w, r, util.ErrRecourseNotAllowed(err))
			return
		}
		videos, err := videoRepo.GetAllPublicVideo()
		if err != nil {
			logger.Log.Errorf("fail to read all public videos error: %d", err)
			render.Render(w, r, util.InternalServeErr(err))
			return
		}
		var videoVideoData = []VideoVideoData{}

		//todo this aggregation should happen in database level.
		for _, video := range videos {
			picture, err := app.PhotoS3Bucket.DownloadFile("tim-api-photos", video.Id)
			if err != nil {
				logger.Log.Errorf("fail to get a picture of video with id: %s", video.Id)
				continue
			}
			videoData, err := repo.GetVideoDate(video.Id)
			if err == nil {
				videoData.Picture = picture
				videoVideoData = append(videoVideoData, VideoVideoData{video, *videoData})
			}
		}

		result, err := json.Marshal(videoVideoData)
		if err != nil {
			fmt.Println("couldn't marshal")
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error marshalling")))
			return
		}
		_, err = w.Write([]byte(result))
		if err != nil {
			return
		}
	}
}

func getAllRaw(app *server_config.Env, repo *videodataRepository.VideoDataRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		videoId := chi.URLParam(r, "videoId")
		accountId := chi.URLParam(r, "accountId")
		var extension = ".mp4"
		videData, err := repo.GetVideoDate(videoId)
		if err == nil {
			extension = "." + videData.FileType
		}
		fmt.Println(accountId)

		http.ServeFile(w, r, "videos/"+videoId+extension)
		//if isAdmin(email) {
		//	//http.ServeFile(w, r, "files/test/"+id+extension)
		//	//bytes,err :=readVideoFile(id,extension)
		//	//if err != nil {
		//	//	return
		//	//}
		//	http.ServeFile(w, r, "videos/"+id+extension)
		//	return
		//}
		return
	}
}

func getVideoPicture(app *server_config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := jwtauth.TokenFromHeader(r)
		_, err := controller_auth.IsAllowed(app, token)
		if err != nil {
			render.Render(w, r, util.ErrRecourseNotAllowed(err))
			return
		}
		id := chi.URLParam(r, "id")
		picture, err := app.PhotoS3Bucket.DownloadFile("tim-api-photos", id)
		if err != nil {
			logger.Log.Errorf("fail to get a picture of video with id: %s", id)
		}

		result, err := json.Marshal(picture)
		if err != nil {
			fmt.Println("couldn't marshal")
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error marshalling")))
			return
		}
		_, err = w.Write([]byte(result))
		if err != nil {
			return
		}
	}
}

//func isAdmin(email string) bool {
//	user.home.controller.domain.controller := user_repo.GetUser(email)
//	if user.home.controller.domain.controller.RoleId == "" {
//		return false
//	}
//	role.domain.controller := role_repo.GetRole(user.home.controller.domain.controller.RoleId)
//	if role.domain.controller.Name == "" {
//		return false
//	}
//	if role.domain.controller.Name == "agent" || role.domain.controller.Name == "admin" || role.domain.controller.Name == "superAdmin" {
//		return true
//	}
//	return false
//}

// @Summary getRaw  Returns a file of a video.controller.
// @ID getRaw-videoData
// @Produce json
// @Parameters:
//
//		-name: tags
//		 in: query
//	  required: true
//		 type: string
//
// @Success 200 {object} VideoData
// @Failure 404 {object} string
// @Router /video.controller/video.controller-data/getRwa/{id} [get]

func getRaw(app *server_config.Env, repo *videodataRepository.VideoDataRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Authenticating user.home.controller.domain.controller by checking the token.
		token := jwtauth.TokenFromHeader(r)
		_, err := controller_auth.IsAllowed(app, token)
		if err != nil {
			render.Render(w, r, util.ErrRecourseNotAllowed(err))
			return
		}
		id := chi.URLParam(r, "id")
		var videoObject videoDataDomain.VideoData
		videoDate, err := repo.GetVideoDate(id)
		if err != nil {
			logger.Log.Errorf("fail to retrieve video data with id: %s error: %d", id, err)
			render.Render(w, r, util.InternalServeErr(err))
			return
		}
		picture, err := app.PhotoS3Bucket.DownloadFile("tim-api-videos", id)
		if err != nil {
			logger.Log.Errorf("fail to get a picture of video with id: %s", id)
		}

		sEnc := base64.StdEncoding.EncodeToString(picture)
		videoObject = videoDataDomain.VideoData{
			Id:       videoDate.Id,
			Picture:  []byte{},
			Video:    []byte{},
			FileType: videoDate.FileType,
			FileSize: sEnc,
		}

		//http.ServeFile(w, r, "files/test/"+id+".mp4")
		//return
		result, err := json.Marshal(videoObject)
		if err != nil {
			fmt.Println("couldn't marshal")
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error marshalling")))
			return
		}
		_, err = w.Write([]byte(result))
		if err != nil {
			return
		}
	}
}

// @Summary delete  remove a specified videoData from DB
// @ID delete-videoData
// @Produce json
// @Parameters:
//
//		-name: tags
//		 in: query
//	  required: true
//		 type: string
//
// @Success 200 {object} VideoData
// @Failure 404 {object} string
// @Router /video.controller/video.controller-data/delete [get]
func delete(app *server_config.Env, repo *videodataRepository.VideoDataRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Authenticating user.home.controller.domain.controller by checking the token.
		token := jwtauth.TokenFromHeader(r)
		_, err := controller_auth.IsAllowed(app, token)
		if err != nil {
			render.Render(w, r, util.ErrRecourseNotAllowed(err))
			return
		}
		id := chi.URLParam(r, "id")
		if id != "" {
			result, err := repo.DeleteVideoData(id)
			if err != nil {
				logger.Log.Errorf("error deleting video with id: %s", id)
				render.Render(w, r, util.ErrInvalidRequest(err))
				return
			}
			bytes, err := json.Marshal(result)
			if err != nil {
				fmt.Println("couldn't marshal")
				render.Render(w, r, util.ErrInvalidRequest(errors.New("error marshalling")))
				return
			}
			_, err = w.Write([]byte(bytes))
			if err != nil {
				return
			}
		}
	}
}

// @Summary getAll returns a list of VideoData object
// @ID getAll-videoData
// @Produce json
// @Success 200 {object} VideoData
// @Failure 404 {object} string
// @Router /video.controller/video.controller-data/getAll [get]
func getAll(app *server_config.Env, repo *videodataRepository.VideoDataRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		token := jwtauth.TokenFromHeader(r)
		_, err := controller_auth.IsAllowed(app, token)
		if err != nil {
			render.Render(w, r, util.ErrRecourseNotAllowed(err))
			return
		}

		video, err := repo.GetVideoDatas()
		if err != nil {
			logger.Log.Errorf("error reading videos")
			render.Render(w, r, util.InternalServeErr(err))
			return
		}
		result, err := json.Marshal(video)
		if err != nil {
			fmt.Println("couldn't marshal")
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error marshalling")))
			return
		}
		_, err = w.Write([]byte(result))
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error writing bytes")))
			return
		}
	}
}

// @Summary create returns a VideoData object
// @ID create-videoData
// @Produce json
// @Parameters:
//
//		-name: tags
//		 in: query
//	  required: true
//		 type: VideoData
//
// @Success 200 {object} VideoData
// @Failure 404 {object} string
// @Router /video.controller/video.controller-data/create [post]
func create(app *server_config.Env, repo *videodataRepository.VideoDataRepository, userVideo *userVideoRepository.UserVideoRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Authenticating user.home.controller.domain.controller by checking the token.
		token := jwtauth.TokenFromHeader(r)
		id, err := controller_auth.IsAllowed(app, token)
		if err != nil {
			render.Render(w, r, util.ErrRecourseNotAllowed(err))
			return
		}
		data := &videoDataDomain.VideoData{}
		err = render.Bind(r, data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		logger.Log.Infof("%d start creating a video data with id: %s", id, data.Id)

		//go util.VideoWriter(app, *data, false, repo)
		go util.VideoWriterWithS3(app, *data, false, repo)

		_, err = userVideo.CreateUserVideo(userVideoDomain.UserVideo{
			Id:        "",
			AccountId: *id,
			VideoId:   data.Id,
			Date:      time.Now(),
		})
		if err != nil {
			logger.Log.Errorf("%d fail to creating a user video data with id: %s", id, data.Id)
		}

		_, err = w.Write([]byte("Done"))
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(errors.New("writing bytes")))
			return
		}
		return
	}
}

// @Summary update Updates an existing videoData
// @ID update-videoData
// @Produce json
// @Parameters:
//
//		-name: tags
//		 in: query
//	  required: true
//		 type: VideoData
//
// @Success 200 {object} VideoData
// @Failure 404 {object} string
// @Router /video.controller/video.controller-data/update [post]
func update(app *server_config.Env, repo *videodataRepository.VideoDataRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Authenticating user.home.controller.domain.controller by checking the token.
		token := jwtauth.TokenFromHeader(r)
		_, err := controller_auth.IsAllowed(app, token)
		if err != nil {
			render.Render(w, r, util.ErrRecourseNotAllowed(err))
			return
		}
		data := &videoDataDomain.VideoData{}
		err = render.Bind(r, data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		response, err := repo.UpdateVideoDate(*data)
		if response.Id == "" {
			fmt.Println("error creating VideoData")
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error creating Video")))
			return
		}
		result, err := json.Marshal(response)
		if err != nil {
			fmt.Println("couldn't marshal")
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error marshalling")))
			return
		}
		_, err = w.Write([]byte(result))
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(errors.New("writing bytes")))
			return
		}
	}
}

func readVideoFile(id string, extension string) ([]byte, error) {
	//file, err := os.ReadFile("files/test/"+id+""+extension)
	file, err := os.ReadFile("files/test/" + id + ".mp4")
	if err != nil {
		fmt.Println(err, " error reading file!")
		return nil, err
	}
	return file, nil
}

//func get(app *config.Env) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		id := chi.URLParam(r, "id")
//		if id != "" {
//			object := repository.GetVideoDate(id)
//			//file, err := readVideoFile(id, "")
//			result, err := json.Marshal(object)
//			if err != nil {
//				fmt.Println("couldn't marshal")
//				render.Render(w, r, util.ErrInvalidRequest(errors.New("error marshalling")))
//				return
//			}
//			_, err = w.Write([]byte(result))
//			if err != nil {
//				return
//			}
//		}
//	}
//}
