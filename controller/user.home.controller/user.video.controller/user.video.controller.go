package userVideoController

import (
	"encoding/json"
	"errors"
	"fmt"
	server_config "github.com/ESPOIR-DITE/tim-api/config/server.config"
	"github.com/ESPOIR-DITE/tim-api/controller/util"
	userVideoDomain "github.com/ESPOIR-DITE/tim-api/domain/user/user.video.domain"
	userVideoRepository "github.com/ESPOIR-DITE/tim-api/storage/user/user.video.repository"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Home(app *server_config.Env) http.Handler {
	r := chi.NewRouter()
	repo := userVideoRepository.NewUserVideoRepository(app.GormDB)
	r.Get("/get/{id}", get(app, repo))
	r.Get("/get-with-videoId/{videoId}", getWithVideoId(app, repo))
	r.Get("/get-all/{customerId}", getAllOf(app, repo))
	r.Get("/delete/{id}", delete(app, repo))
	r.Post("/create", create(app, repo))
	r.Post("/update", update(app, repo))
	r.Get("/getAll", getAll(app, repo))
	return r
}

func getWithVideoId(app *server_config.Env, repo *userVideoRepository.UserVideoRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		videoId := chi.URLParam(r, "videoId")
		if videoId != "" {
			object, err := repo.GetUserVideoWithVideoId(videoId)
			if err != nil {
				log.Errorf("fail to get user.home.controller video with video id: %s, err: %d", videoId, err)
				render.Render(w, r, util.InternalServeErr(err))
				return
			}
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

func getAllOf(app *server_config.Env, repo *userVideoRepository.UserVideoRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		customerId := chi.URLParam(r, "customerId")
		userVideos, err := repo.GetAllUserVideo(customerId)
		if err != nil {
			log.Errorf("fail to get all user.home.controller video with user.home.controller id: %s err: %d", customerId, err)
			render.Render(w, r, util.InternalServeErr(err))
			return
		}
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

func delete(app *server_config.Env, repo *userVideoRepository.UserVideoRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id != "" {
			role, err := repo.DeleteUserVideo(id)
			if err != nil {
				log.Errorf("fail to delete user.home.controller video with id: %s err: %d", id, err)
				render.Render(w, r, util.InternalServeErr(err))
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

func getAll(app *server_config.Env, repo *userVideoRepository.UserVideoRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := repo.GetUserVideos()
		if err != nil {
			log.Errorf("fail to get all user.home.controller video err: %d", err)
			render.Render(w, r, util.InternalServeErr(err))
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

func create(app *server_config.Env, repo *userVideoRepository.UserVideoRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &userVideoDomain.UserVideo{}
		err := render.Bind(r, data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}

		response, err := repo.CreateUserVideo(*data)
		if err != nil {
			log.Errorf("error creating User video err: %d", err)
			render.Render(w, r, util.InternalServeErr(errors.New("error creating User video.controller")))
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

func update(app *server_config.Env, repo *userVideoRepository.UserVideoRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &userVideoDomain.UserVideo{}
		err := render.Bind(r, data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		response, err := repo.UpdateUserVideo(*data)
		if err != nil {
			log.Errorf("fail to update User video error: %d", err)
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error creating User video.controller")))
			return
		}
		result, err := json.Marshal(&response)
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

func get(app *server_config.Env, repo *userVideoRepository.UserVideoRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id != "" {
			object, err := repo.GetUserVideo(id)
			if err != nil {
				log.Errorf("fail to get User video with id: %s error: %d", id, err)
				render.Render(w, r, util.ErrInvalidRequest(errors.New("error creating User video.controller")))
				return
			}
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
