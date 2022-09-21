package user_account

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
	"tim-api/config"
	"tim-api/controller/util"
	user_account "tim-api/domain/user/user-account"
	"tim-api/security"
	repository "tim-api/storage/user/user-account-repo"
)

func Home(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/get/{id}", get(app))
	r.Get("/getWithEmail/{email}", getWithEmail(app))
	r.Get("/delete/{id}", delete(app))
	r.Post("/create", create(app))
	r.Post("/login", login(app))
	r.Post("/update", update(app))
	r.Get("/getAll", getAll(app))
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
// @Router /user/user-account/login [post]
func login(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &user_account.UserAccount{}
		err := render.Bind(r, data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		isSuccessful, response := authenticateUser(data.Password, data.Email)
		if !isSuccessful {
			app.InfoLog.Println("error login User account. Password mismatch.")
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error login User account. Password mismatch.")))
			return
		}
		//response, err := repository.Login(object)
		//if err != nil {
		//	fmt.Println("error creating User account")
		//	render.Render(w, r, util.ErrInvalidRequest(errors.New("error creating User account")))
		//	return
		//}
		token, err := security.GenerateJWT(response.Email)
		if err != nil {
			fmt.Println("error generating token")
			render.Render(w, r, util.InternalServeErr(errors.New("error generating token")))
			return
		}
		userAccountUpdated, err := repository.UpdateToken(response, token)
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
// @Router /user/user-account/delete [get]
func delete(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id != "" {
			role := repository.DeleteUserAccount(id)
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
		user := repository.GetUserAccounts()
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
		data := &user_account.UserAccount{}
		err := render.Bind(r, data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		object := repository.GetUserAccountObject(data)
		response := repository.CreateUserAccount(object)
		if response.CustomerId == "" {
			fmt.Println("error creating User account")
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error creating User account")))
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

func update(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &user_account.UserAccount{}
		err := render.Bind(r, data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		object := repository.GetUserAccountObject(data)
		response := repository.UpdateUserAccount(object)
		if response.CustomerId == "" {
			fmt.Println("error updating user account")
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

func get(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id != "" {
			object := repository.GetUserAccount(id)
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
func getWithEmail(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := chi.URLParam(r, "email")
		if email != "" {
			object := repository.GetUserAccountWithEmail(email)
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

func authenticateUser(password, email string) (bool, user_account.UserAccount) {
	result := repository.GetAllUserAccountByEmail(email)
	for _, userAccount := range result {
		ok, _ := security.ComparePasswords(userAccount.Password, password)
		if ok {
			return true, userAccount
		}
	}
	return false, user_account.UserAccount{}
}
