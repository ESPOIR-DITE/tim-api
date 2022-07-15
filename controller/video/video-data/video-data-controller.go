package video_data

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
	"os"
	"tim-api/config"
	"tim-api/controller/util"
	"tim-api/domain"
	role_repo "tim-api/storage/user/role-repo"
	user_repo "tim-api/storage/user/user-repo"
	repository "tim-api/storage/video/video-data"
	video_repo "tim-api/storage/video/video-repo"
)

func Home(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/get/{id}", get(app))
	r.Get("/delete/{id}", delete(app))
	r.Post("/create", create(app))
	r.Post("/update-video", updateVideo(app))
	r.Post("/create-picture", createPicture(app))
	r.Post("/update-picture", updatePicture(app))
	r.Post("/update", update(app))
	r.Get("/getAll", getAll(app))
	r.Get("/getRwa/{id}", getRaw(app))
	r.Get("/getRwa/{id}/{email}", getAllRaw(app))
	//r.Get("/video-picture/{id}", getVideoPicture(app))
	r.Get("/video-public-picture", getPublicVideoPicture(app))
	r.Get("/video-picture", getPublicVideoPicture(app))
	r.Get("/stream/{videoId}", stream(app))
	return r
}

func updateVideo(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &domain.VideoData{}
		err := render.Bind(r, data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		object := repository.GetCategoryObject(data)
		if len(object.Picture) == 0 {
			fmt.Println(object.Id)
		} else {
			fmt.Println()
		}
		go util.VideoWriter(object, true)

		_, err = w.Write([]byte("Done"))
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(errors.New("writing bytes")))
			return
		}
		return
	}
}

func stream(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		videoId := chi.URLParam(r, "videoId")
		videoDate := repository.GetVideoDate(videoId)
		if videoDate.Id == "" {
			fmt.Println("error reading Video")
			return
		}
		fileBytes, err := util.ReadVideoFile(videoId, videoDate.FileType)
		if err != nil {
			fmt.Println(err)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "video/mp4")
		w.Write(fileBytes)
		return
	}
}

func updatePicture(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &domain.VideoData{}
		err := render.Bind(r, data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		object := repository.GetCategoryObject(data)
		videoDataObject := repository.UpdateVideoDate(object)
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

type VideoVideoData struct {
	Video     domain.Video
	VideoData domain.VideoData
}

func createPicture(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &domain.VideoData{}
		err := render.Bind(r, data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		object := repository.GetCategoryObject(data)
		videoDataObject := repository.CreateVideoData(object)
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

func getPublicVideoPicture(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		videos := video_repo.GetAllPublicVideo()
		var videoVideoData = []VideoVideoData{}
		for _, video := range videos {
			videoData := repository.GetVideoDate(video.Id)
			if videoData.Id != "" {
				videoVideoData = append(videoVideoData, VideoVideoData{video, videoData})
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

func getAllRaw(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		email := chi.URLParam(r, "email")
		var extension = ".mp4"
		videData := repository.GetVideoDate(id)
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

func getVideoPicture(app *config.Env) http.HandlerFunc {
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
	user := user_repo.GetUser(email)
	if user.RoleId == "" {
		return false
	}
	role := role_repo.GetRole(user.RoleId)
	if role.Name == "" {
		return false
	}
	if role.Name == "agent" || role.Name == "admin" || role.Name == "superAdmin" {
		return true
	}
	return false
}
func getRaw(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		var videoObject domain.VideoData
		videoDate := repository.GetVideoDate(id)
		if videoDate.Id == "" {
			fmt.Println("error reading Video")
		} else {
			//videoObject = video_repo.GetVideo(id)
			//file, err := os.ReadFile("files/test/" + id + "." + videoDate.FileType)
			file, err := os.ReadFile("files/videos/" + id + "." + videoDate.FileType)
			if err != nil {
				fmt.Println(err, " error reading file")
			}
			sEnc := base64.StdEncoding.EncodeToString(file)
			videoObject = domain.VideoData{videoDate.Id, []byte{}, []byte{}, videoDate.FileType, sEnc}
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

func delete(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id != "" {
			role := repository.DeleteVideoData(id)
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

func getAll(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := repository.GetVideoDatas()
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

func create(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &domain.VideoData{}
		err := render.Bind(r, data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		object := repository.GetCategoryObject(data)
		if len(object.Picture) == 0 {
			fmt.Println(object.Id)
		} else {
			fmt.Println()
		}
		go util.VideoWriter(object, false)

		_, err = w.Write([]byte("Done"))
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(errors.New("writing bytes")))
			return
		}
		return
	}
}

func update(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &domain.VideoData{}
		err := render.Bind(r, data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		object := repository.GetCategoryObject(data)
		response := repository.UpdateVideoDate(object)
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

func get(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id != "" {
			object := repository.GetVideoDate(id)
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
	}
}
