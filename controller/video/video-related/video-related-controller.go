package video_related

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
	"tim-api/config"
	"tim-api/controller/util"
	video_related "tim-api/domain/video/video-related"
	repository "tim-api/storage/video/video-related"
)

func Home(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/get/{videoI}", get(app))
	r.Post("/delete/{videoI}", deleteVideoRelated(app))
	r.Post("/create", create(app))
	return r
}

func create(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := video_related.VideoRelated{}
		err := render.Bind(r, &data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		response, err := repository.CreateVideoRelated(data)
		if err != nil {
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

func deleteVideoRelated(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := video_related.VideoRelated{}
		err := render.Bind(r, &data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		object, err := repository.DeleteVideoRelated(data)
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

func get(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id != "" {
			object, err := repository.GetVideosRelatedTo(id)
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
