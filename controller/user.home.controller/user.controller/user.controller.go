package userController

import (
	"encoding/json"
	"errors"
	"fmt"
	server_config "github.com/ESPOIR-DITE/tim-api/config/server.config"
	"github.com/ESPOIR-DITE/tim-api/controller/util"
	userDomain "github.com/ESPOIR-DITE/tim-api/domain/user/user.domain"
	userRepository "github.com/ESPOIR-DITE/tim-api/storage/user/user.repository"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Home(app *server_config.Env) http.Handler {
	r := chi.NewRouter()
	repo := userRepository.NewUserRepository(app.GormDB)
	r.Get("/get/{id}", get(app, repo))
	r.Get("/get-with-token/{token}", getWithToken(app, repo))
	r.Get("/delete/{id}", delete(app, repo))
	r.Post("/create", create(app, repo))
	r.Post("/update", update(app, repo))
	r.Get("/getAll", getAll(app, repo))
	r.Get("/getAllAgents", getAllAgents(app, repo))
	r.Get("/getAllUsers", getAllUsers(app, repo))
	r.Get("/getAllAdmins", getAllAdmins(app, repo))
	r.Get("/getAllSuperAdmins", getAllSuperAdmins(app, repo))
	return r
}

func getAllUsers(app *server_config.Env, repo *userRepository.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := repo.GetUsers()
		if err != nil {
			log.Errorf("fail to retrieve users err: %d", err)
			render.Render(w, r, util.InternalServeErr(errors.New("error marshalling")))
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

func getAllSuperAdmins(app *server_config.Env, repo *userRepository.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := repo.GetAllSuperAdmins()
		if err != nil {
			log.Errorf("fail to retrieve all super admins err: %d", err)
			render.Render(w, r, util.InternalServeErr(errors.New("error marshalling")))
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

func getAllAdmins(app *server_config.Env, repo *userRepository.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := repo.GetAllAdmins()
		if err != nil {
			log.Errorf("fail retrieving all admins err: %d", err)
			render.Render(w, r, util.InternalServeErr(errors.New("error marshalling")))
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

func getAllAgents(app *server_config.Env, repo *userRepository.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := repo.GetAllAgents()
		if err != nil {
			log.Errorf("fail to retrieve all agents err: %d", err)
			render.Render(w, r, util.InternalServeErr(errors.New("error marshalling")))
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

func delete(app *server_config.Env, repo *userRepository.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id != "" {
			role, err := repo.Delete(id)
			if err != nil {
				log.Errorf("fail to delete with id: %s, err %d", id, err)
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

func getAll(app *server_config.Env, repo *userRepository.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := repo.GetUsers()
		if err != nil {
			log.Errorf("fail to get all err %d", err)
			render.Render(w, r, util.InternalServeErr(errors.New("error marshalling")))
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

func create(app *server_config.Env, repo *userRepository.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &userDomain.User{}
		err := render.Bind(r, data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		object, err := repo.CreateUser(*data)
		if err != nil {
			log.Errorf("fail to create user.home.controller err %d", err)
			render.Render(w, r, util.InternalServeErr(errors.New("error marshalling")))
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
			render.Render(w, r, util.ErrInvalidRequest(errors.New("writing bytes")))
			return
		}
	}
}

func update(app *server_config.Env, repo *userRepository.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &userDomain.User{}
		err := render.Bind(r, data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		response, err := repo.UpdateUser(*data)
		if err != nil {
			fmt.Println("error updating user.home.controller err: %d", err)
			log.Errorf("fail to updating user.home.controller err: %d", err)
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error creating User")))
			return
		}
		result, err := json.Marshal(*response)
		if err != nil {
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

func get(app *server_config.Env, repo *userRepository.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id != "" {
			object, err := repo.GetUser(id)
			if err != nil {
				log.Errorf("fail to get user.home.controller err %d", err)
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
	}
}
func getWithToken(app *server_config.Env, repo *userRepository.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := chi.URLParam(r, "token")
		if token != "" {
			claims, err := app.SecurityService.ValidateToken(token)
			if err != nil {
				log.Errorf("fail to get user err %d", err)
				render.Render(w, r, util.ErrRecourseNotAllowed(err))
				return
			}
			object, err := repo.GetUser(claims.Email)
			if err != nil {
				log.Errorf("fail to get user.home.controller err %d", err)
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
	}
}
