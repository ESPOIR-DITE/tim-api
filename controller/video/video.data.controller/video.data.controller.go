package videoDataHomeController

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	server_config "github.com/ESPOIR-DITE/tim-api/config/server.config"
	"github.com/ESPOIR-DITE/tim-api/controller/util"
	controller_auth "github.com/ESPOIR-DITE/tim-api/controller/util/controller-auth"
	"github.com/ESPOIR-DITE/tim-api/domain/video/video"
	videoDataDomain "github.com/ESPOIR-DITE/tim-api/domain/video/video.data.domain"
	videodataRepository "github.com/ESPOIR-DITE/tim-api/storage/video/video.data.repository"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func Home(app *server_config.Env) http.Handler {
	r := chi.NewRouter()
	repo := videodataRepository.NewVideoDataRepository(app.GormDB)
	r.Get("/get/{id}", get(app, repo))
	r.Get("/delete/{id}", delete(app, repo))
	r.Post("/create", create(app, repo))
	r.Post("/update-video.controller", updateVideo(app, repo))
	r.Post("/create-picture", createPicture(app, repo))
	r.Post("/update-picture", updatePicture(app, repo))
	r.Post("/update", update(app, repo))
	r.Get("/getAll", getAll(app, repo))
	r.Get("/getRwa/{id}", getRaw(app, repo))
	r.Get("/getRwa/{id}/{email}", getAllRaw(app, repo))
	//r.Get("/video.controller-picture/{id}", getVideoPicture(app))
	//r.Get("/video.controller-public-picture", getPublicVideoPicture(app, repo))
	//r.Get("/video.controller-picture", getPublicVideoPicture(app, repo))
	r.Get("/stream/{videoId}", stream(app, repo))
	return r
}

func updateVideo(app *server_config.Env, repo *videodataRepository.VideoDataRepository) http.HandlerFunc {
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
		go util.VideoWriter(app, *data, true, repo)

		_, err = w.Write([]byte("Done"))
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(errors.New("writing bytes")))
			return
		}
		return
	}
}

func stream(app *server_config.Env, repo *videodataRepository.VideoDataRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := jwtauth.TokenFromHeader(r)
		_, err := controller_auth.IsAllowed(app, token)
		if err != nil {
			render.Render(w, r, util.ErrRecourseNotAllowed(err))
			return
		}
		videoId := chi.URLParam(r, "videoId")
		videoDate, err := repo.GetVideoDate(videoId)
		if err != nil {
			fmt.Println("error reading Video")
			log.Errorf("error reading Video err: %d", err)
			return
		}
		fileBytes, err := util.ReadVideoFile(videoId, videoDate.FileType)
		if err != nil {
			log.Errorf("error reading Video err: %d", err)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "video.controller/mp4")
		w.Write(fileBytes)
		return
	}
}

func updatePicture(app *server_config.Env, repo *videodataRepository.VideoDataRepository) http.HandlerFunc {
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
		videoDataObject, err := repo.UpdateVideoDate(*data)
		if err != nil {
			render.Render(w, r, util.InternalServeErr(err))
			return
		}
		result, err := json.Marshal(videoDataObject)
		if err != nil {
			fmt.Println(err, "Marshal error updating videoData")
		}
		_, err = w.Write([]byte(result))
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(errors.New("writing bytes")))
			return
		}
		return
	}
}

type VideoVideoData struct {
	Video     video.Video
	VideoData videoDataDomain.VideoData
}

func createPicture(app *server_config.Env, repo *videodataRepository.VideoDataRepository) http.HandlerFunc {
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
		videoDataObject, err := repo.CreateVideoData(*data)
		if err != nil {
			log.Error(err)
			render.Render(w, r, util.InternalServeErr(err))
			return
		}
		result, err := json.Marshal(videoDataObject)
		if err != nil {
			fmt.Println(err, " error creating videoData")
		}
		_, err = w.Write([]byte(result))
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(errors.New("writing bytes")))
			return
		}
		return
	}
}

//func getPublicVideoPicture(app *config.Env, repo *videodataRepository.VideoDataRepository) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		videos := video_repo.GetAllPublicVideo()
//		var videoVideoData = []VideoVideoData{}
//		for _, video := range videos {
//			videoData := repository.GetVideoDate(video.Id)
//			if videoData.Id != "" {
//				videoVideoData = append(videoVideoData, VideoVideoData{video, videoData})
//			}
//		}
//		result, err := json.Marshal(videoVideoData)
//		if err != nil {
//			fmt.Println("couldn't marshal")
//			render.Render(w, r, util.ErrInvalidRequest(errors.New("error marshalling")))
//			return
//		}
//		_, err = w.Write([]byte(result))
//		if err != nil {
//			return
//		}
//	}
//}

// Todo to be implemented.
func getAllRaw(app *server_config.Env, repo *videodataRepository.VideoDataRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := jwtauth.TokenFromHeader(r)
		_, err := controller_auth.IsAllowed(app, token)
		if err != nil {
			render.Render(w, r, util.ErrRecourseNotAllowed(err))
			return
		}
		id := chi.URLParam(r, "id")
		email := chi.URLParam(r, "email")
		var extension = ".mp4"
		videData, err := repo.GetVideoDate(id)
		if err != nil {
			//http.ServeFile(w, r, "files/test/"+id+extension)
			//bytes,err :=readVideoFile(id,extension)
			//if err != nil {
			//	return
			//}
			http.ServeFile(w, r, "videos/"+id+extension)
			return
		}
		if videData.FileType != "" {
			extension = "." + videData.FileType
		}
		if isAdmin(email) {
			//http.ServeFile(w, r, "files/test/"+id+extension)
			//bytes,err :=readVideoFile(id,extension)
			//if err != nil {
			//	return
			//}
			http.ServeFile(w, r, "videos/"+id+extension)
			return
		}
		return
	}
}

func getVideoPicture(app *server_config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		picture, err := util.GetVideoPictures(id)
		if err != nil {
			fmt.Println(err, " error reading picture!")
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

func isAdmin(email string) bool {
	//user.home.controller.domain.controller := user_repo.GetUser(email)
	//if user.home.controller.domain.controller.RoleId == "" {
	//	return false
	//}
	//role.domain.controller := role_repo.GetRole(user.home.controller.domain.controller.RoleId)
	//if role.domain.controller.Name == "" {
	//	return false
	//}
	//if role.domain.controller.Name == "agent" || role.domain.controller.Name == "admin" || role.domain.controller.Name == "superAdmin" {
	//	return true
	//}
	return false
}

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
			fmt.Println("error reading Video")
		} else {
			//videoObject = video_repo.GetVideo(id)
			//file, err := os.ReadFile("files/test/" + id + "." + videoDate.FileType)
			file, err := os.ReadFile("files/videos/" + id + "." + videoDate.FileType)
			if err != nil {
				fmt.Println(err, " error reading file")
			}
			sEnc := base64.StdEncoding.EncodeToString(file)
			videoDate.FileSize = sEnc
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
			role, err := repo.DeleteVideoData(id)
			if err != nil {
				log.Errorf("fail to retrieve data with %d", err)
				render.Render(w, r, util.InternalServeErr(errors.New("error marshalling")))
				return
			}
			result, err := json.Marshal(role)
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
		user, err := repo.GetVideoDatas()
		if err != nil {
			fmt.Println("couldn't retrieve data")
			render.Render(w, r, util.InternalServeErr(errors.New("error marshalling")))
			return
		}
		result, err := json.Marshal(user)
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
func create(app *server_config.Env, repo *videodataRepository.VideoDataRepository) http.HandlerFunc {
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
		go util.VideoWriter(app, *data, false, repo)

		_, err = w.Write([]byte("Done"))
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
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
		if err != nil {
			log.Errorf("fail tp create VideoData, err: %d", err)
			render.Render(w, r, util.InternalServeErr(errors.New("error creating Video")))
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

func get(app *server_config.Env, repo *videodataRepository.VideoDataRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := jwtauth.TokenFromHeader(r)
		_, err := controller_auth.IsAllowed(app, token)
		if err != nil {
			render.Render(w, r, util.ErrRecourseNotAllowed(err))
			return
		}
		id := chi.URLParam(r, "id")
		if id != "" {
			object, err := repo.GetVideoDate(id)
			if err != nil {
				log.Errorf("fail to retrieve video data with id: %s, err; %d", id, err)
				render.Render(w, r, util.ErrInvalidRequest(errors.New("error marshalling")))
				return
			}
			//file, err := readVideoFile(id, "")
			result, err := json.Marshal(object)
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
		render.Render(w, r, util.ErrInvalidRequest(errors.New("missing required param")))
		return
	}
}
