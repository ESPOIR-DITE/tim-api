package video_comment

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	"net/http"
	"tim-api/config"
	"tim-api/controller/util"
	controller_auth "tim-api/controller/util/controller-auth"
	"tim-api/domain"
	repository "tim-api/storage/video/video-repo"
)

func Home(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/get/{id}", get(app))
	r.Get("/get-pictures/{email}", getUserVideo(app))
	r.Get("/delete/{id}", delete(app))
	r.Post("/create", create(app))
	r.Post("/update", update(app))
	r.Get("/getAll", getAll(app))
	return r
}

// @Summary getUserVideo returns a list of Video of a user [agent]
// @ID getUserVideo-video
// @Produce json
// @Success 200 {object} Video
// @Failure 404 {object} message
// @Router /video/get-pictures/{email} [get]
func getUserVideo(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := repository.GetVideos()
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

// @Summary delete returns a boolean
// @ID delete-video
// @Produce json
// @Success 200 {object} Video
// @Failure 404 {object} message
// @Router /video/video/delete [get]
func delete(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id != "" {
			role := repository.DeleteVideo(id)
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

// @Summary homeHandler returns a list of Video
// @ID getAll-video
// @Produce json
// @Success 200 {object} Video
// @Failure 404 {object} message
// @Router /video/video/getAll [get]
func getAll(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := repository.GetVideos()
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

// @Summary homeHandler returns a Video object
// @ID getAll-video
// @Produce json
// @Param data body Model:Video true
// @Success 200 {object} Video
// @Failure 404 {object} message
// @Router /video/video/create [post]
func create(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := jwtauth.TokenFromHeader(r)
		controller_auth.IsAllowed(token, w, r)
		data := &domain.Video{}
		err := render.Bind(r, data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		object := repository.GetVideoObject(data)
		response := repository.CreateVideo(object)
		if response.Id == "" {
			fmt.Println("error creating video")
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error creating video")))
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

// @Summary update returns a Video object
// @ID update-video
// @Produce json
// @Param data body Model:Video true
// @Success 200 {object} Video
// @Failure 404 {object} message
// @Router /video/video/update [post]
func update(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := jwtauth.TokenFromHeader(r)
		controller_auth.IsAllowed(token, w, r)
		data := &domain.Video{}
		err := render.Bind(r, data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		object := repository.GetVideoObject(data)
		response := repository.UpdateVideo(object)
		if response.Id == "" {
			fmt.Println("error creating Video")
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

// @Summary get returns a Video object
// @ID get-video
// @Produce json
// @Param id path string true
// @Success 200 {object} Video
// @Failure 404 {object} message
// @Router /user/role/get/{id} [get]
func get(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id != "" {
			object := repository.GetVideo(id)
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
