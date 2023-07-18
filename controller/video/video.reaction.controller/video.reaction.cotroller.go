package videoReactionController

import (
	"encoding/json"
	"errors"
	"fmt"
	server_config "github.com/ESPOIR-DITE/tim-api/config/server.config"
	"github.com/ESPOIR-DITE/tim-api/controller/util"
	video_reaction "github.com/ESPOIR-DITE/tim-api/domain/video/video.reaction"
	videoReactionRepository "github.com/ESPOIR-DITE/tim-api/storage/video/video-reaction-repo"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Home(app *server_config.Env) http.Handler {
	r := chi.NewRouter()
	repo := videoReactionRepository.NewVideoReactionRepository(app.GormDB)
	r.Get("/get/{id}", get(app, repo))
	r.Post("/reaction", reactToVideo(app, repo))
	return r
}

func reactToVideo(app *server_config.Env, repo *videoReactionRepository.VideoReactionRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := video_reaction.VideoReaction{}
		err := render.Bind(r, &data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		var response *video_reaction.VideoReaction

		if data.Like > 0 {
			response, err = repo.LikeReact(data)
			if err != nil {
				app.ErrorLog.Println("error creating video reaction.")
				log.Error("error creating video reaction.")
				render.Render(w, r, util.ErrInvalidRequest(errors.New("error creating video.controller reaction")))
				return
			}
		} else {
			response, err = repo.UnLikeReact(data)
			if err != nil {
				app.ErrorLog.Println("error creating video reaction.")
				log.Error("error creating video reaction.")
				render.Render(w, r, util.ErrInvalidRequest(errors.New("error creating video.controller reaction")))
				return
			}
		}

		if response.VideoId == "" {
			app.ErrorLog.Println("error creating video.controller reaction.")
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error creating video.controller reaction")))
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

func get(app *server_config.Env, repo *videoReactionRepository.VideoReactionRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id != "" {
			object, err := repo.GetVideo(id)
			if err != nil {
				fmt.Println("couldn't marshal")
				log.Errorf("couldn't find video reaction with id: %s, err: %d", id, err)
				render.Render(w, r, util.InternalServeErr(err))
				return
			}
			result, err := json.Marshal(object)
			if err != nil {
				fmt.Println("couldn't marshal")
				render.Render(w, r, util.InternalServeErr(errors.New("error marshalling")))
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
