package videoController

import (
	"encoding/json"
	"errors"
	"fmt"
	server_config "github.com/ESPOIR-DITE/tim-api/config/server.config"
	"github.com/ESPOIR-DITE/tim-api/controller/util"
	controller_auth "github.com/ESPOIR-DITE/tim-api/controller/util/controller-auth"
	"github.com/ESPOIR-DITE/tim-api/domain/video/video"
	videoRepository "github.com/ESPOIR-DITE/tim-api/storage/video/video.repository"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Home(app *server_config.Env) http.Handler {
	r := chi.NewRouter()
	repo := videoRepository.NewVideoRepository(app.GormDB)
	r.Get("/get/{id}", get(app, repo))
	r.Get("/get-pictures/{email}", getUserVideo(app, repo))
	r.Get("/delete/{id}", delete(app, repo))
	r.Post("/create", create(app, repo))
	r.Post("/update", update(app, repo))
	r.Get("/getAll", getAll(app, repo))
	return r
}

// @Summary getUserVideo returns a list of Video of a user.home.controller.domain.controller [agent]
// @ID getUserVideo-video.controller
// @Produce json
// @responses:
//
//		200: Video
//	 404: string
//	 500: string
//
// @Router /video.controller/video.controller/get-pictures/{email} [get]
func getUserVideo(app *server_config.Env, repo *videoRepository.VideoRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		videos, err := repo.GetVideos()
		if err != nil {
			log.Error(err)
			render.Render(w, r, util.InternalServeErr(err))
			return
		}
		result, err := json.Marshal(videos)
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
// @ID deleteVideo
// @Produce json
// @responses:
//
//		200: Video
//	 404: string
//	 500: string
//
// @Router /video.controller/video.controller/delete [get]
func delete(app *server_config.Env, repo *videoRepository.VideoRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id != "" {
			role, err := repo.DeleteVideo(id)
			if err != nil {
				log.Errorf("failed to delete video: %s, error: %d", id, err)
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
		render.Render(w, r, util.ErrInvalidRequest(errors.New("missing required field")))
		return
	}
}

// @Summary homeHandler returns a list of video.controller
// @ID getAllVideo
// @Produce json
// @responses:
//
//		200: Video
//	 404: string
//	 500: string
//
// @Router /video.controller/video.controller/getAll [get]
func getAll(app *server_config.Env, repo *videoRepository.VideoRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := jwtauth.TokenFromHeader(r)
		_, err := controller_auth.IsAllowed(app, token)
		if err != nil {
			render.Render(w, r, util.ErrRecourseNotAllowed(err))
			return
		}
		user, err := repo.GetVideos()
		if err != nil {
			log.Errorf("failed to retrieve videos:  error: %d", err)
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

// @Summary create returns a video.controller object
// @ID createVideo
// @Produce json
// @Parameters:
//
//		-name: tags
//		 in: query
//	  required: true
//		 type: Video
//
// @responses:
//
//		200: Video
//	 404: string
//	 500: string
//
// @Router /video.controller/video.controller/create [post]
func create(app *server_config.Env, repo *videoRepository.VideoRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := jwtauth.TokenFromHeader(r)
		_, err := controller_auth.IsAllowed(app, token)
		if err != nil {
			render.Render(w, r, util.ErrRecourseNotAllowed(err))
			return
		}
		data := &video.Video{}
		err = render.Bind(r, data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		response, err := repo.CreateVideo(*data)
		if err != nil {
			log.Errorf("error creating video.controller, error: %d", err)
			render.Render(w, r, util.InternalServeErr(errors.New("error creating video.controller")))
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

// @Summary update returns a video.controller object
// @ID updateVideo
// @Produce json
// @Parameters:
//
//		-name: tags
//		 in: query
//	  required: true
//		 type: Video
//
// @responses:
//
//		200: Video
//	 404: string
//	 500: string
//
// @Router /video.controller/video.controller/update [post]
func update(app *server_config.Env, repo *videoRepository.VideoRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := jwtauth.TokenFromHeader(r)
		_, err := controller_auth.IsAllowed(app, token)
		if err != nil {
			render.Render(w, r, util.ErrRecourseNotAllowed(err))
			return
		}
		data := &video.Video{}
		err = render.Bind(r, data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		response, err := repo.UpdateVideo(*data)
		if err != nil {
			log.Errorf("error updating Video err: %d", err)
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

// @Summary get returns a video.controller object
// @ID get-video.controller
// @Produce json
// @Parameters:
//
//		-name: tags
//		 in: query
//	  required: true
//		 type: Video
//
// @responses:
//
//		200: Video
//	 404: string
//	 500: string
//
// @Router /video.controller/video.controller/get/{id} [get]
func get(app *server_config.Env, repo *videoRepository.VideoRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id != "" {
			object, err := repo.GetVideo(id)
			if err != nil {
				log.Errorf("couldn't retieve video with id: %s, err: %d", id, err)
				render.Render(w, r, util.ErrInvalidRequest(errors.New("error marshalling")))
				return
			}
			result, err := json.Marshal(object)
			if err != nil {
				log.Errorf("couldn't marshal video with id: %s, err: %d", id, err)
				render.Render(w, r, util.ErrInvalidRequest(errors.New("error marshalling")))
				return
			}
			_, err = w.Write([]byte(result))
			if err != nil {
				return
			}
		}
		render.Render(w, r, util.ErrInvalidRequest(errors.New("missing required field")))
		return
	}
}
