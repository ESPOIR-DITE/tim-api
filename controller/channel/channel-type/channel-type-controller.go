package role

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	"net/http"
	"tim-api/config"
	"tim-api/controller/util"
	controller_auth "tim-api/controller/util/controller-auth"
	"tim-api/domain"
	repository "tim-api/storage/chanel/channel-type-repository"
)

func Home(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/get/{id}", getChannelType(app))
	r.Get("/delete/{id}", deleteChannelType(app))
	r.Post("/create", createChannelType(app))
	r.Post("/update", updateChannelType(app))
	r.Get("/getAll", getChannelTypes(app))

	//r.Use(middleware.LoginSession{SessionManager: app.Session}.RequireAuthenticatedUser)
	//r.Get("/home", indexHanler(app))
	//r.Get("/homeError", indexErrorHanler(app))
	return r
}

func deleteChannelType(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id != "" {
			role := repository.DeleteChannelType(id)
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
func updateChannelType(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := domain.ChannelType{}
		err := render.Bind(r, &data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		var response = repository.UpdateChannelType(data)
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
func getChannelType(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		user := repository.GetChannelType(id)
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
func createChannelType(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Authenticate Token
		token := jwtauth.TokenFromHeader(r)
		controller_auth.IsAllowed(token, w, r)

		data := domain.ChannelType{}
		err := render.Bind(r, &data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		response, err := repository.CreateChannelType(data)
		if err != nil {
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
func getChannelTypes(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		channels := repository.GetChannelTypes()
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
