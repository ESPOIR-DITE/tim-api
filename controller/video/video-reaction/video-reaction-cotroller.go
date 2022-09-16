package video_reaction

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
	"tim-api/config"
	"tim-api/controller/util"
	video_reaction "tim-api/domain/video/video-reaction"
	repository "tim-api/storage/video/video-reaction-repo"
)

func Home(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/get/{id}", get(app))
	r.Post("/reaction", reactToVideo(app))
	return r
}

func reactToVideo(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := video_reaction.VideoReaction{}
		err := render.Bind(r, &data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		var response video_reaction.VideoReaction

		if data.Like > 0 {
			response = repository.LikeReact(data)
		} else {
			response = repository.UnLikeReact(data)
		}

		if response.VideoId == "" {
			app.ErrorLog.Println("error creating video reaction.")
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error creating video reaction")))
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
