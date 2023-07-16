package channelTypeController

import (
	"encoding/json"
	"errors"
	"fmt"
	server_config "github.com/ESPOIR-DITE/tim-api/config/server.config"
	"github.com/ESPOIR-DITE/tim-api/controller/util"
	controller_auth "github.com/ESPOIR-DITE/tim-api/controller/util/controller-auth"
	channel_type "github.com/ESPOIR-DITE/tim-api/domain/channel/channel-type"
	repository "github.com/ESPOIR-DITE/tim-api/storage/chanel/channel.type.repository"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Home(app *server_config.Env) http.Handler {
	r := chi.NewRouter()
	repo := repository.NewChannelTypeRepository(app.GormDB)

	r.Get("/get/{id}", getChannelType(app, repo))
	r.Get("/delete/{id}", deleteChannelType(app, repo))
	r.Post("/create", createChannelType(app, repo))
	r.Post("/update", updateChannelType(app, repo))
	r.Get("/getAll", getChannelTypes(app, repo))

	//r.Use(middleware.LoginSession{SessionManager: app.Session}.RequireAuthenticatedUser)
	//r.Get("/home", indexHanler(app))
	//r.Get("/homeError", indexErrorHanler(app))
	return r
}

func deleteChannelType(app *server_config.Env, repo *repository.ChannelTypeRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id != "" {
			role, err := repo.DeleteChannelType(id)
			if err != nil {
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
		render.Render(w, r, util.ErrInvalidRequest(errors.New("error marshalling")))
		return
	}
}
func updateChannelType(app *server_config.Env, repo *repository.ChannelTypeRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := channel_type.ChannelType{}
		if err := render.Bind(r, &data); err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		response, err := repo.UpdateChannelType(data)
		if err != nil {
			log.Error(err)
			render.Render(w, r, util.InternalServeErr(err))
			return
		}
		result, err := json.Marshal(response)
		if err != nil {
			log.Error(err)
			render.Render(w, r, util.InternalServeErr(errors.New("error marshalling")))
			return
		}
		_, err = w.Write([]byte(result))
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(errors.New("writing bytes")))
			return
		}
	}
}
func getChannelType(app *server_config.Env, repo *repository.ChannelTypeRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		user, err := repo.GetChannelType(id)
		if err != nil {
			log.Error(err)
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error marshalling")))
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
			return
		}
	}
}
func createChannelType(app *server_config.Env, repo *repository.ChannelTypeRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Authenticate Token
		token := jwtauth.TokenFromHeader(r)
		_, err := controller_auth.IsAllowed(app, token)
		if err != nil {
			render.Render(w, r, util.ErrRecourseNotAllowed(err))
			return
		}
		data := channel_type.ChannelType{}
		err = render.Bind(r, &data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		response, err := repo.CreateChannelType(data)
		if err != nil {
			fmt.Println("error creating role.domain.controller")
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error creating channel.controller-subscription")))
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
func getChannelTypes(app *server_config.Env, repo *repository.ChannelTypeRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		channels, err := repo.GetChannelTypes()
		if err != nil {
			log.Error(err)
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		result, err := json.Marshal(channels)
		if err != nil {
			log.Error(err)
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		_, err = w.Write([]byte(result))
		if err != nil {
			return
		}
	}
}
