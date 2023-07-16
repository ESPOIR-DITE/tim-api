package userAccountController

import (
	"encoding/json"
	"errors"
	"fmt"
	server_config "github.com/ESPOIR-DITE/tim-api/config/server.config"
	"github.com/ESPOIR-DITE/tim-api/controller/util"
	user_account "github.com/ESPOIR-DITE/tim-api/domain/user/user.account.domain"
	userAccountRepository "github.com/ESPOIR-DITE/tim-api/storage/user/user-account-repo"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Home(app *server_config.Env) http.Handler {
	r := chi.NewRouter()
	repo := userAccountRepository.NewUserAccountRepository(app.GormDB)
	//r.Get("/get/{id}", get(app, repo))
	//r.Get("/getWithEmail/{email}", getWithEmail(app, repo))
	r.Get("/delete/{id}", delete(app, repo))
	r.Post("/create", create(app, repo))
	//r.Post("/login", login(app, repo))
	r.Post("/update", update(app, repo))
	r.Get("/getAll", getAll(app, repo))
	return r
}

// login godoc
// @Summary Authenticate User based on User struct containing email and password
// @Produce json
// @Parameters:
//
//		-name: tags
//		 in: query
//	  required: true
//		 type: UserAccount
//
// @Success 200 {object} string
// @Router /user.home.controller.domain.controller/user.home.controller.domain.controller-account/login [post]
//todo to move to account controller
//func login(app *config.Env) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		data := &user_account.UserAccount{}
//		err := render.Bind(r, data)
//		if err != nil {
//			render.Render(w, r, util.ErrInvalidRequest(err))
//			return
//		}
//		isSuccessful, response := authenticateUser(data.Password, data.Email)
//		if !isSuccessful {
//			app.InfoLog.Println("error login User account. Password mismatch.")
//			render.Render(w, r, util.ErrInvalidRequest(errors.New("error login User account. Password mismatch.")))
//			return
//		}
//		//response, err := repository.Login(object)
//		//if err != nil {
//		//	fmt.Println("error creating User account")
//		//	render.Render(w, r, util.ErrInvalidRequest(errors.New("error creating User account")))
//		//	return
//		//}
//		token, err := security.GenerateJWT(response.Email)
//		if err != nil {
//			fmt.Println("error generating token")
//			render.Render(w, r, util.InternalServeErr(errors.New("error generating token")))
//			return
//		}
//		userAccountUpdated, err := repository.UpdateToken(response, token)
//		if err != nil {
//			fmt.Println("error updating User account")
//			render.Render(w, r, util.InternalServeErr(errors.New("error updating userAccount")))
//			return
//		}
//
//		result, err := json.Marshal(userAccountUpdated)
//		if err != nil {
//			fmt.Println("couldn't marshal")
//			render.Render(w, r, util.ErrInvalidRequest(errors.New("error marshalling")))
//			return
//		}
//		_, err = w.Write([]byte(result))
//		if err != nil {
//			render.Render(w, r, util.ErrInvalidRequest(errors.New("writing bytes")))
//			return
//		}
//	}
//}

// delete godoc
// @Summary Deletes UserAccount based on id
// @Produce json
// @Parameters:
//
//		-name: tags
//		 in: query
//	  required: true
//		 type: UserAccount
//
// @Success 200 {object} bool
// @Router /user.home.controller.domain.controller/user.home.controller.domain.controller-account/delete [get]
func delete(app *server_config.Env, repo *userAccountRepository.UserAccountRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id != "" {
			userAccount, err := repo.DeleteUserAccount(id)
			if err != nil {
				log.Error(err)
				render.Render(w, r, util.InternalServeErr(errors.New("error marshalling")))
				return
			}
			result, err := json.Marshal(userAccount)
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

func getAll(app *server_config.Env, repo *userAccountRepository.UserAccountRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := repo.GetUserAccounts()
		if err != nil {
			log.Error(err)
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

func create(app *server_config.Env, repo *userAccountRepository.UserAccountRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &user_account.UserAccount{}
		err := render.Bind(r, data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		response, err := repo.CreateUserAccount(*data)
		if err != nil {
			log.Error(err)
			render.Render(w, r, util.InternalServeErr(err))
			return
		}
		result, err := json.Marshal(response)
		if err != nil {
			log.Error(err)
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

func update(app *server_config.Env, repo *userAccountRepository.UserAccountRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &user_account.UserAccount{}
		err := render.Bind(r, data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		response, err := repo.UpdateUserAccount(*data)
		if err != nil {
			log.Errorf("fail to update user.home.controller account error: %d", err)
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error creating Video")))
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

//func get(app *config.Env, repo * userAccountRepository.UserAccountRepository) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		id := chi.URLParam(r, "id")
//		if id != "" {
//			object := repo.GetUserAccount(id)
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
//func getWithEmail(app *config.Env, repo * userAccountRepository.UserAccountRepository) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		email := chi.URLParam(r, "email")
//		if email != "" {
//			object := repo.GetUserAccountWithEmail(email)
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

//func authenticateUser(password, email string, repo * userAccountRepository.UserAccountRepository) (bool, user_account.UserAccount) {
//	result := repo.GetAllUserAccountByEmail(email)
//	for _, userAccount := range result {
//		ok, _ := security.ComparePasswords(userAccount.Password, password)
//		if ok {
//			return true, userAccount
//		}
//	}
//	return false, user_account.UserAccount{}
//}
