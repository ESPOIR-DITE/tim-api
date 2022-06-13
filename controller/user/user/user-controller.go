package video_comment

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
	repository "tim-api/storage/user/user-repo"
)

func Home(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/get/{id}", get(app))
	r.Get("/delete/{id}", delete(app))
	r.Post("/create", create(app))
	r.Post("/update", update(app))
	r.Get("/getAll", getAll(app))
	r.Get("/getAllAgents", getAllAgents(app))
	r.Get("/getAllUsers", getAllUsers(app))
	r.Get("/getAllAdmins", getAllAdmins(app))
	r.Get("/getAllSuperAdmins", getAllSuperAdmins(app))
	return r
}

func getAllUsers(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := repository.GetAllUsers()
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

func getAllSuperAdmins(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := repository.GetAllSuperAdmins()
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

func getAllAdmins(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := repository.GetAllAdmins()
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

func getAllAgents(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := repository.GetAllAgents()
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

func delete(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id != "" {
			role := repository.Delete(id)
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

func getAll(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := repository.GetUsers()
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

func create(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &domain.User{}
		err := render.Bind(r, data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		object := repository.GetUserObject(data)
		response := repository.CreateUser(object)
		if response.Email == "" {
			fmt.Println("error creating User account")
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error creating User account")))
			return
		}
		result, err := json.Marshal(repository.GetUserObject(response))
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

func update(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &domain.User{}
		err := render.Bind(r, data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		object := repository.GetUserObject(data)
		response := repository.UpdateUser(object)
		if response.Email == "" {
			fmt.Println("error creating User")
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error creating User")))
			return
		}
		result, err := json.Marshal(repository.GetUserObject(response))
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

func get(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id != "" {
			object := repository.GetUser(id)
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
