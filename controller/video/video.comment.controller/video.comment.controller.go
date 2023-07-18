package video_comment_controller

import (
	"encoding/json"
	"errors"
	"fmt"
	server_config "github.com/ESPOIR-DITE/tim-api/config/server.config"
	"github.com/ESPOIR-DITE/tim-api/controller/util"
	videoCommentDomain "github.com/ESPOIR-DITE/tim-api/domain/video/video.comment.domain"
	videoCommentRepository "github.com/ESPOIR-DITE/tim-api/storage/video/video.comment"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Home(app *server_config.Env) http.Handler {
	r := chi.NewRouter()
	repo := videoCommentRepository.NewVideoCommentRepository(app.GormDB)

	r.Get("/get/{id}", get(app, repo))
	r.Get("/delete/{id}", delete(app, repo))
	r.Post("/create", create(app, repo))
	r.Post("/create", update(app, repo))
	r.Get("/getAll", getAll(app, repo))
	return r
}

func delete(app *server_config.Env, repo *videoCommentRepository.VideoCommentRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id != "" {
			videoComment, err := repo.DeleteVideoComment(id)
			if err != nil {
				errorResponse := fmt.Sprintf("couldn't delete video comment with id: %s, err: %d", id, err)
				log.Errorf(errorResponse)
				render.Render(w, r, util.ErrInvalidRequest(errors.New(errorResponse)))
				return
			}
			result, err := json.Marshal(&videoComment)
			if err != nil {
				fmt.Println("couldn't marshal")
				log.Error(err)
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

func getAll(app *server_config.Env, repo *videoCommentRepository.VideoCommentRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := repo.GetVideoComments()
		if err != nil {
			log.Error(err)
			render.Render(w, r, util.InternalServeErr(errors.New(err.Error())))
			return
		}
		result, err := json.Marshal(user)
		if err != nil {
			log.Error(err)
			render.Render(w, r, util.InternalServeErr(errors.New("error marshalling")))
			return
		}
		_, err = w.Write([]byte(result))
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error writing bytes")))
			return
		}
	}
}

func create(app *server_config.Env, repo *videoCommentRepository.VideoCommentRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &videoCommentDomain.VideoComment{}
		err := render.Bind(r, data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		response, err := repo.CreateVideoComment(*data)
		if err != nil {
			log.Errorf("fail to create video comment err: %d", err)
			render.Render(w, r, util.InternalServeErr(errors.New("fail to creating video comment")))
			return
		}

		result, err := json.Marshal(*response)
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

func update(app *server_config.Env, repo *videoCommentRepository.VideoCommentRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &videoCommentDomain.VideoComment{}
		err := render.Bind(r, data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		response, err := repo.UpdateVideoComment(*data)
		if response.Id == "" {
			fmt.Println("fail to update Video comment")
			log.Errorf("fail to update Video comment err: %d", err)
			render.Render(w, r, util.ErrInvalidRequest(errors.New("fail to update Video comment")))
			return
		}
		result, err := json.Marshal(*response)
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

func get(app *server_config.Env, repo *videoCommentRepository.VideoCommentRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id != "" {
			object, err := repo.GetVideoComment(id)
			if err != nil {
				fmt.Println("couldn't marshal")
				render.Render(w, r, util.ErrInvalidRequest(errors.New("error marshalling")))
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
