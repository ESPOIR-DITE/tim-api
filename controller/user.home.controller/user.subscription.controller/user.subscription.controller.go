package userSubscriptionController

import (
	"encoding/json"
	"errors"
	"fmt"
	server_config "github.com/ESPOIR-DITE/tim-api/config/server.config"
	"github.com/ESPOIR-DITE/tim-api/controller/util"
	userSubscriptionDomain "github.com/ESPOIR-DITE/tim-api/domain/user/user.subscription.domain"
	userSubscriptionRepository "github.com/ESPOIR-DITE/tim-api/storage/user/user-sub-repo"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Home(app *server_config.Env) http.Handler {
	r := chi.NewRouter()
	repo := userSubscriptionRepository.NewUserSubscriptionRepository(app.GormDB)
	r.Get("/get/{id}", get(app, repo))
	r.Get("/delete/{id}", delete(app, repo))
	r.Post("/create", create(app, repo))
	r.Post("/create", update(app, repo))
	r.Get("/getAll", getAll(app, repo))
	return r
}

func delete(app *server_config.Env, repo *userSubscriptionRepository.UserSubscriptionRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id != "" {
			role, err := repo.DeleteUserSubscription(id)
			if err != nil {
				log.Error(err)
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
		render.Render(w, r, util.ErrInvalidRequest(errors.New("missing required param")))
		return
	}
}

func getAll(app *server_config.Env, repo *userSubscriptionRepository.UserSubscriptionRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := repo.GetUserSubscriptions()
		if err != nil {
			log.Error(err)
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

func create(app *server_config.Env, repo *userSubscriptionRepository.UserSubscriptionRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &userSubscriptionDomain.UserSubscription{}
		err := render.Bind(r, data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		response, err := repo.CreateUserSubscription(*data)
		if err != nil {
			log.Errorf("error creating User account error: %d", err)
			render.Render(w, r, util.InternalServeErr(errors.New("error creating User account")))
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

func update(app *server_config.Env, repo *userSubscriptionRepository.UserSubscriptionRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &userSubscriptionDomain.UserSubscription{}
		err := render.Bind(r, data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		response, err := repo.UpdateUserSubscription(*data)
		if err != nil {
			log.Error("error creating UserSubscription")
			render.Render(w, r, util.InternalServeErr(errors.New("error creating UserSubscription")))
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

func get(app *server_config.Env, repo *userSubscriptionRepository.UserSubscriptionRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id != "" {
			object, err := repo.GetUserSubscription(id)
			if err != nil {
				log.Error(err)
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
