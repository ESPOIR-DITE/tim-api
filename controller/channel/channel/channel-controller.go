package role

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
	repository "tim-api/storage/chanel/channel-repository"
)

func Home(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/get/{id}", getChannel(app))
	r.Get("/delete/{id}", deleteChannel(app))
	r.Get("/get-by-user/{userId}", getChannelsByUser(app))
	r.Get("/get-by-region/{region}", getChannelsByRegion(app))
	r.Post("/create", createChannel(app))
	r.Post("/update", updateChannel(app))
	r.Get("/getAll", getChannels(app))

	//r.Use(middleware.LoginSession{SessionManager: app.Session}.RequireAuthenticatedUser)
	//r.Get("/home", indexHanler(app))
	//r.Get("/homeError", indexErrorHanler(app))
	return r
}

func deleteChannel(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id != "" {
			role := repository.DeleteChannel(id)
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

func getChannelsByRegion(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "region")
		if id != "" {
			role := repository.GetChannelsByRegion(id)
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
func getChannelsByUser(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "userId")
		if id != "" {
			role := repository.GetChannelsByUser(id)
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
func updateChannel(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := domain.Channel{}
		err := render.Bind(r, &data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		var response = repository.UpdateChannel(data)
		if response.Id == "" {
			fmt.Println("error creating role")
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error creating channel")))
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
func getChannel(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		user := repository.GetChannel(id)
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
func createChannel(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := domain.Channel{}
		err := render.Bind(r, &data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		var response = repository.CreateChannel(data)
		if response.Id == "" {
			fmt.Println("error creating role")
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error creating channel")))
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
func getChannels(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		channels := repository.GetChannels()
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
