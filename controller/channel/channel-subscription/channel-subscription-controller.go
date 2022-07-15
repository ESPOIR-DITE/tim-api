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
	repository "tim-api/storage/chanel/channelSubscription"
)

func Home(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/get/{id}", getChannelSubscription(app))
	r.Get("/get-by-channel/{id}", countChannelSubscriptionByChannel(app))
	r.Get("/delete/{id}", deleteChannelSubscription(app))
	r.Get("/get-by-user/{userId}", getChannelSubscriptionsByUser(app))
	r.Get("/get-by-channel/{channelId}", getChannelsByChannel(app))
	r.Post("/create", createChannelSubscription(app))
	r.Post("/update", updateChannel(app))
	r.Get("/getAll", getChannelSubscriptions(app))

	//r.Use(middleware.LoginSession{SessionManager: app.Session}.RequireAuthenticatedUser)
	//r.Get("/home", indexHanler(app))
	//r.Get("/homeError", indexErrorHanler(app))
	return r
}

func countChannelSubscriptionByChannel(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		user := repository.GetChannelSubscription(id)
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

func deleteChannelSubscription(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id != "" {
			role := repository.CountSubscriptionByChannelId(id)
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

func getChannelsByChannel(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "channelId")
		if id != "" {
			role := repository.GetChannelSubscriptionsByChannelId(id)
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
func getChannelSubscriptionsByUser(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "userId")
		if id != "" {
			role := repository.GetChannelSubscriptionsByUser(id)
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
		data := domain.ChannelSubscription{}
		err := render.Bind(r, &data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		var response = repository.UpdateChannelSubscription(data)
		if response.Id == "" {
			fmt.Println("error creating role")
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error creating channel-subscription")))
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
func getChannelSubscription(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		user := repository.GetChannelSubscription(id)
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
func createChannelSubscription(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := domain.ChannelSubscription{}
		err := render.Bind(r, &data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		var response = repository.CreateChannelSubscription(data)
		if response.Id == "" {
			fmt.Println("error creating role")
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error creating channel-subscription")))
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
func getChannelSubscriptions(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		channels := repository.GetChannelSubscriptions()
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
