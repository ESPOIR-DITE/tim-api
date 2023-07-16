package videoCategoryHomeController

import (
	"encoding/json"
	"errors"
	"fmt"
	server_config "github.com/ESPOIR-DITE/tim-api/config/server.config"
	"github.com/ESPOIR-DITE/tim-api/controller/util"
	videoCategoryDomain "github.com/ESPOIR-DITE/tim-api/domain/video/video.category.domain"
	video_category "github.com/ESPOIR-DITE/tim-api/storage/video/video-category"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Home(app *server_config.Env) http.Handler {
	r := chi.NewRouter()
	repo := video_category.NewVideoCategoryRepository(app.GormDB)
	r.Get("/get/{id}", getVideoCategory(app, repo))
	r.Get("/delete/{id}", deleteVideoCategory(app, repo))
	r.Post("/create", createVideoCategory(app, repo))
	r.Post("/update", updateVideoCategory(app, repo))
	r.Get("/getAll", getVideoCategories(app, repo))
	return r
}

func deleteVideoCategory(app *server_config.Env, repo *video_category.VideoCategoryRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id != "" {
			role, err := repo.DeleteVideoCategory(id)
			if err != nil {
				log.Errorf("Fail to delete %s, err: %d", id, err)
				app.WarningLog.Printf("Fail to delete %s", id)
				render.Render(w, r, util.ErrInvalidRequest(err))
				return
			}
			result, err := json.Marshal(role)
			if err != nil || role == false {
				app.WarningLog.Printf("Fail to Marshal %s", role)
				render.Render(w, r, util.ErrInvalidRequest(err))
				return
			}
			app.InfoLog.Printf("Delete success full %s", id)
			_, err = w.Write([]byte(result))
			if err != nil {
				return
			}
		}
		render.Render(w, r, util.ErrInvalidRequest(errors.New("missing required param")))
		return
	}
}

func getVideoCategories(app *server_config.Env, repo *video_category.VideoCategoryRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := repo.GetVideoCategories()
		if err != nil {
			log.Errorf("fail to retrieve video category, err: %d", err)
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error marshalling")))
			return
		}
		result, err := json.Marshal(user)
		if err != nil {
			app.ErrorLog.Printf("Fail to read all video.controller category")
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error marshalling")))
			return
		}

		_, err = w.Write([]byte(result))
		if err != nil {
			app.InfoLog.Println("Fail to create video.controller category ")
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error writing bytes")))
			return
		}
	}
}

func createVideoCategory(app *server_config.Env, repo *video_category.VideoCategoryRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &videoCategoryDomain.VideoCategory{}
		err := render.Bind(r, data)
		if err != nil {
			app.ErrorLog.Printf("Fail to render video.controller category %s", data)
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		response, err := repo.CreateVideoCategory(*data)
		if err != nil {
			app.ErrorLog.Printf("Fail to create video.controller category %s", data)
			log.Errorf("Fail to create video category %s, err: %d", data, err)
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error creating Category")))
			return
		}
		result, err := json.Marshal(response)
		if err != nil {
			log.Errorf("Fail to marshal video.controller category err: %d", err)
			app.ErrorLog.Printf("Fail to marshal video.controller category %s", data)
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error marshalling")))
			return
		}
		app.InfoLog.Printf("created video.controller category %s", data)
		_, err = w.Write([]byte(result))
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(errors.New("writing bytes")))
			return
		}
	}
}

func updateVideoCategory(app *server_config.Env, repo *video_category.VideoCategoryRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &videoCategoryDomain.VideoCategory{}
		err := render.Bind(r, data)
		if err != nil {
			app.ErrorLog.Printf("Fail to update video.controller category %s", data)
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		response, err := repo.UpdateVideoCategory(*data)
		if err != nil {
			log.Errorf("Fail to updating video category %d", err)
			app.ErrorLog.Printf("Fail to updating video.controller category %s", data)
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error creating VideoCategory")))
			return
		}
		result, err := json.Marshal(&response)
		if err != nil {
			app.ErrorLog.Printf("Fail to marshal video.controller category %s", data)
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

func getVideoCategory(app *server_config.Env, repo *video_category.VideoCategoryRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id != "" {
			videoCategory, err := repo.GetVideoCategory(id)
			if err != nil {
				log.Errorf("fail to create video category err: %d", err)
				render.Render(w, r, util.InternalServeErr(errors.New("fail to create video category")))
				return
			}
			result, err := json.Marshal(&videoCategory)
			if err != nil {
				fmt.Println("couldn't marshal")
				render.Render(w, r, util.ErrInvalidRequest(errors.New("error marshalling")))
				return
			}
			//app.InfoLog.Printf("reading video Category by %s result is: %s", id, role.domain.controller)
			_, err = w.Write([]byte(result))
			if err != nil {
				return
			}
		}
		log.Error("missing param in the link.")
		render.Render(w, r, util.ErrInvalidRequest(errors.New("missing param")))
		return
	}
}
