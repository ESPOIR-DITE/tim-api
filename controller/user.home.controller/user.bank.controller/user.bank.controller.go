package userBankController

import (
	"encoding/json"
	"errors"
	"fmt"
	server_config "github.com/ESPOIR-DITE/tim-api/config/server.config"
	"github.com/ESPOIR-DITE/tim-api/controller/util"
	user_details "github.com/ESPOIR-DITE/tim-api/domain/user/user.details.domain"
	userBankRepository "github.com/ESPOIR-DITE/tim-api/storage/user/user.bank.repository"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Home(app *server_config.Env) http.Handler {
	r := chi.NewRouter()
	repo := userBankRepository.NewUserBankRepository(app.GormDB)

	r.Get("/get/{id}", get(app, repo))
	//r.Get("/getWithEmail/{email}", getWithEmail(app,repo))
	r.Get("/delete/{id}", delete(app, repo))
	//r.Get("/get-by-email/{id}", deleteByEmail(app,repo))
	r.Post("/create", create(app, repo))
	r.Post("/update", update(app, repo))
	r.Get("/getAll", getAll(app, repo))
	//r.Get("/isExist-by-email/{email}", isExist(app,repo))
	r.Get("/isExist-by-id/{id}", isExistById(app, repo))

	return r
}

func isExistById(app *server_config.Env, repo *userBankRepository.UserBankRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id != "" {
			role, err := repo.IsExistUserBankById(id)
			if err != nil {
				log.Errorf("fail to get all err %d", err)
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
		render.Render(w, r, util.ErrInvalidRequest(errors.New("missing param")))
		return
	}
}

//func isExist(app *config.Env, repo *userBankRepository.UserBankRepository) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		email := chi.URLParam(r, "email")
//		if email != "" {
//			role := repo.IsExistUserBankByEmail(email)
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

func delete(app *server_config.Env, repo *userBankRepository.UserBankRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id != "" {
			role, err := repo.DeleteUserBankById(id)
			if err != nil {
				log.Errorf("fail to delete with id: %s err: %d", id, err)
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
	}
}

//func deleteByEmail(app *config.Env,repo *userBankRepository.UserBankRepository) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		id := chi.URLParam(r, "id")
//		if id != "" {
//			role, := repo.DeleteUserBankByUserEmail(id)
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

func getAll(app *server_config.Env, repo *userBankRepository.UserBankRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := repo.GetUserBanks()
		if err != nil {
			log.Errorf("fail to get all  err: %d", err)
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

func create(app *server_config.Env, repo *userBankRepository.UserBankRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := user_details.UserBank{}
		err := render.Bind(r, &data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		response, err := repo.CreateUserDetails(data)
		if err != nil {
			log.Errorf("error creating User account err: %d", err)
			render.Render(w, r, util.InternalServeErr(errors.New("error creating User account")))
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

func update(app *server_config.Env, repo *userBankRepository.UserBankRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := user_details.UserBank{}
		err := render.Bind(r, &data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		response, err := repo.UpdateUserBank(data)
		if err != nil {
			log.Errorf("fail to update err: %d", err)
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

func get(app *server_config.Env, repo *userBankRepository.UserBankRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id != "" {
			object, err := repo.GetUserBank(id)
			if err != nil {
				log.Errorf("fail to get with id: %s err: %d", id, err)
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
				return
			}
		}
		render.Render(w, r, util.InternalServeErr(errors.New("missing required param")))
		return
	}
}

//func getWithEmail(app *config.Env,repo *userBankRepository.UserBankRepository) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		email := chi.URLParam(r, "email")
//		if email != "" {
//			object := repo.GetUserBankByEmail(email)
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
