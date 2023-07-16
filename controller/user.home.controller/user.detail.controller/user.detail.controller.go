package userDetailController

import (
	"encoding/json"
	"errors"
	"fmt"
	server_config "github.com/ESPOIR-DITE/tim-api/config/server.config"
	"github.com/ESPOIR-DITE/tim-api/controller/util"
	user_details "github.com/ESPOIR-DITE/tim-api/domain/user/user.details.domain"
	userDetailsRepository "github.com/ESPOIR-DITE/tim-api/storage/user/user.details.repository"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Home(app *server_config.Env) http.Handler {
	r := chi.NewRouter()
	repo := userDetailsRepository.NewUserDetailsRepository(app.GormDB)

	r.Get("/get/{id}", get(app, repo))
	//r.Get("/getWithEmail/{email}", getWithEmail(app, repo))
	r.Get("/delete/{id}", delete(app, repo))
	//r.Get("/isExist-by-email/{email}", isExist(app, repo))
	r.Get("/isExist-by-id/{email}", isExistById(app, repo))
	//r.Get("/get-by-email/{id}", deleteByEmail(app, repo))
	r.Post("/create", create(app, repo))
	r.Post("/update", update(app, repo))
	r.Get("/getAll", getAll(app, repo))

	return r
}

func isExistById(app *server_config.Env, repo *userDetailsRepository.UserDetailsRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id != "" {
			role, err := repo.IsExistUserDetailById(id)
			if err != nil {
				log.Errorf("fail to get user.home.controller deatild, err: %d", err)
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

//func isExist(app *config.Env, repo userDetailsRepository.UserDetailsRepository) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		email := chi.URLParam(r, "email")
//		if email != "" {
//			role := repo.IsExistUserDetailsByEmail(email)
//			result, err := json.Marshal(role)
//			if err != nil {
//				fmt.Println("couldn't marshal")
//				render.Render(w, r, util.ErrInvalidRequest(errors.New("error marshalling")))
//				return
//			}
//			_, err = w.Write([]byte(result))
//			if err != nil {
//				return
//			}
//		}
//	}
//}

func delete(app *server_config.Env, repo *userDetailsRepository.UserDetailsRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id != "" {
			role, err := repo.DeleteUserDetailsById(id)
			if err != nil {
				log.Errorf("fail to delete user.home.controller deatild with id: %s, err: %d", id, err)
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
		render.Render(w, r, util.ErrInvalidRequest(errors.New("error marshalling")))
		return
	}
}

//func deleteByEmail(app *config.Env, repo userDetailsRepository.UserDetailsRepository) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		id := chi.URLParam(r, "id")
//		if id != "" {
//			role := repository.DeleteUserDetailsByUserEmail(id)
//			result, err := json.Marshal(role)
//			if err != nil {
//				fmt.Println("couldn't marshal")
//				render.Render(w, r, util.ErrInvalidRequest(errors.New("error marshalling")))
//				return
//			}
//			_, err = w.Write([]byte(result))
//			if err != nil {
//				return
//			}
//		}
//	}
//}

func getAll(app *server_config.Env, repo *userDetailsRepository.UserDetailsRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := repo.GetUserDetails()
		if err != nil {
			log.Errorf("fail to get all user.home.controller deatials err: %d", err)
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

func create(app *server_config.Env, repo *userDetailsRepository.UserDetailsRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := user_details.AccountDetails{}
		err := render.Bind(r, &data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		response, err := repo.CreateUserDetails(data)
		if err != nil {
			log.Errorf("fail to create user.home.controller deatials err: %d", err)
			render.Render(w, r, util.InternalServeErr(errors.New("error marshalling")))
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

func update(app *server_config.Env, repo *userDetailsRepository.UserDetailsRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := user_details.AccountDetails{}
		err := render.Bind(r, &data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		response, err := repo.UpdateUserDetails(data)
		if err != nil {
			log.Errorf("fail updating user.home.controller details account err: %d", err)
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error creating Video")))
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

func get(app *server_config.Env, repo *userDetailsRepository.UserDetailsRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id != "" {
			object, err := repo.GetUserDetail(id)
			if err != nil {
				log.Errorf("fail to get user.home.controller details account with id: %s err: %d", id, err)
				render.Render(w, r, util.ErrInvalidRequest(errors.New("error creating Video")))
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

//func getWithEmail(app *config.Env, repo userDetailsRepository.UserDetailsRepository) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		email := chi.URLParam(r, "email")
//		if email != "" {
//			object := repo.GetUserDetailsByEmail(email)
//			result, err := json.Marshal(object)
//			if err != nil {
//				fmt.Println("couldn't marshal")
//				render.Render(w, r, util.ErrInvalidRequest(errors.New("error marshalling")))
//				return
//			}
//			_, err = w.Write([]byte(result))
//			if err != nil {
//				return
//			}
//		}
//	}
//}
