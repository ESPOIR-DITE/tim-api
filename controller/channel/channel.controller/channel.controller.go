package channelController

import (
	"encoding/json"
	"errors"
	"fmt"
	server_config "github.com/ESPOIR-DITE/tim-api/config/server.config"
	"github.com/ESPOIR-DITE/tim-api/controller/util"
	"github.com/ESPOIR-DITE/tim-api/domain/channel/channel"
	repository "github.com/ESPOIR-DITE/tim-api/storage/chanel/channel-repository"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
)

func Home(app *server_config.Env) http.Handler {
	r := chi.NewRouter()
	repo := repository.NewChanelRepository(app.GormDB)

	r.Get("/get/{id}", getChannel(repo))
	r.Get("/delete/{id}", deleteChannel(repo))
	r.Get("/get-by-user.home.controller.domain.controller/{userId}", getChannelsByUser(repo))
	r.Get("/get-by-region/{region}", getChannelsByRegion(repo))
	r.Post("/create", createChannel(repo))
	r.Post("/update", updateChannel(repo))
	r.Get("/getAll", getChannels(repo))

	return r
}

func deleteChannel(repo *repository.ChanelRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id != "" {
			role := repo.DeleteChannel(id)
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
	}
}

func getChannelsByRegion(repo *repository.ChanelRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "region")
		if id != "" {
			role := repo.GetChannelsByRegion(id)
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
	}
}
func getChannelsByUser(repo *repository.ChanelRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "userId")
		if id != "" {
			role := repo.GetChannelsByUser(id)
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
	}
}
func updateChannel(repo *repository.ChanelRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := channel.Channel{}
		err := render.Bind(r, &data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		var response = repo.UpdateChannel(data)
		if response.Id == "" {
			fmt.Println("error creating role.domain.controller")
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error creating channel.controller")))
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
func getChannel(repo *repository.ChanelRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		user := repo.GetChannel(id)
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
func createChannel(repo *repository.ChanelRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := channel.Channel{}
		err := render.Bind(r, &data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		var response = repo.CreateChannel(data)
		if response.Id == "" {
			fmt.Println("error creating role.domain.controller")
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error creating channel.controller")))
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
func getChannels(repo *repository.ChanelRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		channels := repo.GetChannels()
		result, err := json.Marshal(channels)
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
