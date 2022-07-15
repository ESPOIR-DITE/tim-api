package video_category

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
	repository "tim-api/storage/video/video-category"
)

func Home(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/get/{id}", getVideoCategory(app))
	r.Get("/delete/{id}", deleteVideoCategory(app))
	r.Post("/create", createVideoCategory(app))
	r.Post("/update", updateVideoCategory(app))
	r.Get("/getAll", getVideoCategories(app))
	return r
}

func deleteVideoCategory(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id != "" {
			role := repository.DeleteVideoCategory(id)
			if role == false {
				app.WarningLog.Printf("Fail to delete %s", id)
				render.Render(w, r, util.ErrInvalidRequest(errors.New("error marshalling")))
				return
			}
			result, err := json.Marshal(role)
			if err != nil || role == false {
				app.WarningLog.Printf("Fail to Marshal %s", role)
				render.Render(w, r, util.ErrInvalidRequest(errors.New("error marshalling")))
				return
			}
			app.InfoLog.Printf("Delete success fully %s", id)
			_, err = w.Write([]byte(result))
			if err != nil {
				return
			}
		}
	}
}

func getVideoCategories(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := repository.GetVideoCategories()
		result, err := json.Marshal(user)
		if err != nil {
			app.ErrorLog.Printf("Fail to read all video category")
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error marshalling")))
			return
		}

		_, err = w.Write([]byte(result))
		if err != nil {
			app.InfoLog.Println("Fail to create video category ")
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error writing bytes")))
			return
		}
		app.InfoLog.Println("Fail to create video category ")
		return
	}
}

func createVideoCategory(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &domain.VideoCategory{}
		err := render.Bind(r, data)
		if err != nil {
			app.ErrorLog.Printf("Fail to render video category %s", data)
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		categoryObject := repository.GetVideoObject(data)
		response := repository.CreateVideoCategory(categoryObject)
		if response.Id == "" {
			app.ErrorLog.Printf("Fail to create video category %s", data)
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error creating Category")))
			return
		}
		result, err := json.Marshal(repository.GetVideoObject(response))
		if err != nil {
			app.ErrorLog.Printf("Fail to marshal video category %s", data)
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error marshalling")))
			return
		}
		app.InfoLog.Printf("created video category %s", data)
		_, err = w.Write([]byte(result))
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(errors.New("writing bytes")))
			return
		}
	}
}

func updateVideoCategory(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &domain.VideoCategory{}
		err := render.Bind(r, data)
		if err != nil {
			app.ErrorLog.Printf("Fail to update video category %s", data)
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		categoryObject := repository.GetVideoObject(data)
		response := repository.UpdateVideoCategory(categoryObject)
		if response.Id == "" {
			app.ErrorLog.Printf("Fail to updating video category %s", data)
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error creating VideoCategory")))
			return
		}
		result, err := json.Marshal(repository.GetVideoObject(response))
		if err != nil {
			app.ErrorLog.Printf("Fail to marshal video category %s", data)
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

func getVideoCategory(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id != "" {
			role := repository.GetVideo(id)
			result, err := json.Marshal(role)
			if err != nil {
				fmt.Println("couldn't marshal")
				render.Render(w, r, util.ErrInvalidRequest(errors.New("error marshalling")))
				return
			}
			app.InfoLog.Printf("reading video Category by %s result is: %s", id, role)
			_, err = w.Write([]byte(result))
			if err != nil {
				return
			}
		}
		app.WarningLog.Println("missing URLParam in the link.")
		render.Render(w, r, util.ErrInvalidRequest(errors.New("error marshalling")))
		return
	}
}
