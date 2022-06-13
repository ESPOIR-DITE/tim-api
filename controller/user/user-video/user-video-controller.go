package user_subscription

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
	"tim-api/config"
	"tim-api/controller/util"
	"tim-api/domain"
	repository "tim-api/storage/user/user-video-repo"
)

func Home(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/get/{id}", get(app))
	r.Get("/get-with-videoId/{videoId}", getWithVideoId(app))
	r.Get("/get-all/{customerId}", getAllOf(app))
	r.Get("/delete/{id}", delete(app))
	r.Post("/create", create(app))
	r.Post("/update", update(app))
	r.Get("/getAll", getAll(app))
	return r
}

func getWithVideoId(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		videoId := chi.URLParam(r, "videoId")
		if videoId != "" {
			object := repository.GetUserVideoWithVideoId(videoId)
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

func getAllOf(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		customerId := chi.URLParam(r, "customerId")
		userVideos := repository.GetAllUserVideo(customerId)
		result, err := json.Marshal(userVideos)
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

func delete(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id != "" {
			role := repository.DeleteUserVideo(id)
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
		user := repository.GetUserVideos()
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
		data := &domain.UserVideo{}
		err := render.Bind(r, data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}

		object := repository.GetUserVideoObject(data)
		response := repository.CreateUserVideo(object)
		if response.Id == "" {
			fmt.Println("error creating User video")
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error creating User video")))
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

func update(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &domain.UserVideo{}
		err := render.Bind(r, data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		object := repository.GetUserVideoObject(data)
		response := repository.UpdateUserVideo(object)
		if response.Id == "" {
			fmt.Println("error creating User video")
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error creating User video")))
			return
		}
		result, err := json.Marshal(repository.GetUserVideoObject(response))
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

func get(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id != "" {
			object := repository.GetUserVideo(id)
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
