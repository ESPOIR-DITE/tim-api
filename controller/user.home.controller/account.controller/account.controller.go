package accountController

import (
	"encoding/json"
	"errors"
	"fmt"
	server_config "github.com/ESPOIR-DITE/tim-api/config/server.config"
	"github.com/ESPOIR-DITE/tim-api/controller/util"
	controller_auth "github.com/ESPOIR-DITE/tim-api/controller/util/controller-auth"
	"github.com/ESPOIR-DITE/tim-api/domain/user/account"
	accountRepository "github.com/ESPOIR-DITE/tim-api/storage/user/account.repository"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Home(app *server_config.Env) http.Handler {
	r := chi.NewRouter()
	repo := accountRepository.NewAccountRepository(app.GormDB)
	//r.Get("/get/{id}", get(app, repo))
	//r.Get("/getWithEmail/{email}", getWithEmail(app, repo))
	r.Get("/delete/{id}", delete(app, repo))
	r.Post("/create", create(app, repo))
	r.Post("/login", login(app, repo))
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
// todo to move to account controller
func login(app *server_config.Env, repo *accountRepository.AccountRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &account.Account{}
		err := render.Bind(r, data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}

		accountObject, err := repo.LoginUser(data.Email)
		if err != nil {
			log.Fatalf("error login User account. invalid login credentials")
			app.InfoLog.Println("error login User account. invalid credentials")
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error login User account. invalid credentials")))
			return
		}
		isCorrect, err := app.SecurityService.ComparePasswords(accountObject.Password, data.Password)
		if err != nil {
			log.Fatalf("error login User account error: %d", err)
			render.Render(w, r, util.InternalServeErr(errors.New("error login User account. password decryption")))
			return
		}
		if !isCorrect {
			render.Render(w, r, util.ErrRecourseNotAllowed(errors.New("error login User account. invalid credentials")))
			return
		}
		token, err := app.SecurityService.GenerateJWT(accountObject.Email, accountObject.Id)
		if err != nil {
			fmt.Println("error generating token")
			render.Render(w, r, util.InternalServeErr(errors.New("error generating token")))
			return
		}
		userAccountUpdated, err := repo.UpdateToken(*accountObject, token)
		if err != nil {
			fmt.Println("error updating User account")
			render.Render(w, r, util.InternalServeErr(errors.New("error updating userAccount")))
			return
		}

		result, err := json.Marshal(userAccountUpdated)
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
func delete(app *server_config.Env, repo *accountRepository.AccountRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := jwtauth.TokenFromHeader(r)
		email, err := controller_auth.IsAllowed(app, token)
		if err != nil {
			render.Render(w, r, util.ErrRecourseNotAllowed(err))
			return
		}
		id := chi.URLParam(r, "id")
		if id != "" {
			userAccount, err := repo.DeleteAccount(id)
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
			log.Infof("%s requested deletetion of account with id: %s ", *email, id)
			_, err = w.Write([]byte(result))
			if err != nil {
				render.Render(w, r, util.ErrInvalidRequest(errors.New("error marshalling")))
				return
			}
		}
		render.Render(w, r, util.ErrInvalidRequest(errors.New("missing required param")))
		return
	}
}

func getAll(app *server_config.Env, repo *accountRepository.AccountRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := jwtauth.TokenFromHeader(r)
		email, err := controller_auth.IsAllowed(app, token)
		if err != nil {
			render.Render(w, r, util.ErrRecourseNotAllowed(err))
			return
		}
		log.Infof("%s rested get all account", *email)
		user, err := repo.GetAccounts()
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

func create(app *server_config.Env, repo *accountRepository.AccountRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &account.Account{}
		err := render.Bind(r, data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		encryptedData, err := app.SecurityService.EncryptPassword(data.Password)
		if err == nil {
			data.Password = encryptedData
		} else {
			log.Fatalf("could not encrypt account password for email: %s, error: %d", data.Email, err)
		}

		response, err := repo.CreateAccount(*data)
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

func update(app *server_config.Env, repo *accountRepository.AccountRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := jwtauth.TokenFromHeader(r)
		email, err := controller_auth.IsAllowed(app, token)
		if err != nil {
			render.Render(w, r, util.ErrRecourseNotAllowed(err))
			return
		}
		data := &account.Account{}
		err = render.Bind(r, data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		response, err := repo.UpdateAccount(*data)
		if err != nil {
			log.Errorf("fail to update user.home.controller account error: %d", err)
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error creating Video")))
			return
		}
		log.Infof("%s resquested an update of account with email %s", *email, data.Email)

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

func get(app *server_config.Env, repo *accountRepository.AccountRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := jwtauth.TokenFromHeader(r)
		_, err := controller_auth.IsAllowed(app, token)
		if err != nil {
			render.Render(w, r, util.ErrRecourseNotAllowed(err))
			return
		}
		id := chi.URLParam(r, "id")
		if id != "" {
			object, err := repo.GetAccount(id)
			if err != nil {
				log.Error(err)
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

func getWithEmail(app *server_config.Env, repo *accountRepository.AccountRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := chi.URLParam(r, "email")
		if email != "" {
			object, err := repo.GetAccountWithEmail(email)
			if err != nil {
				log.Error(err)
				render.Render(w, r, util.ErrInvalidRequest(err))
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
