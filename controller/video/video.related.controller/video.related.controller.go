package videoRelatedController

import (
	"encoding/json"
	"errors"
	"fmt"
	server_config "github.com/ESPOIR-DITE/tim-api/config/server.config"
	"github.com/ESPOIR-DITE/tim-api/controller/util"
	video_related "github.com/ESPOIR-DITE/tim-api/domain/video/video.related.domain"
	videoRelatedDomain "github.com/ESPOIR-DITE/tim-api/storage/video/video-related"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Home(app *server_config.Env) http.Handler {
	r := chi.NewRouter()

	repo := videoRelatedDomain.NewVideoRelatedRepository(app.GormDB)
	r.Get("/get/{videoI}", get(app, repo))
	r.Post("/delete/{videoI}", deleteVideoRelated(app, repo))
	r.Post("/create", create(app, repo))
	return r
}

func create(app *server_config.Env, repo *videoRelatedDomain.VideoRelatedRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := video_related.VideoRelated{}
		err := render.Bind(r, &data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		response, err := repo.CreateVideoRelated(data)
		if err != nil {
			log.Error(err)
			render.Render(w, r, util.InternalServeErr(err))
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

func deleteVideoRelated(app *server_config.Env, repo *videoRelatedDomain.VideoRelatedRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := video_related.VideoRelated{}
		err := render.Bind(r, &data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		object, err := repo.DeleteVideoRelated(data)
		if err != nil {
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
}

func get(app *server_config.Env, repo *videoRelatedDomain.VideoRelatedRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id != "" {
			object, err := repo.GetVideosRelatedTo(id)
			if err != nil {
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
	}
}
