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
	repository "tim-api/storage/user/role-repo"
)

func Home(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Handle("/", homeHandler(app))
	r.Get("/get/{id}", getRole(app))
	r.Post("/create", createRole(app))
	r.Get("/getAll", getRoles(app))

	//r.Use(middleware.LoginSession{SessionManager: app.Session}.RequireAuthenticatedUser)
	//r.Get("/home", indexHanler(app))
	//r.Get("/homeError", indexErrorHanler(app))
	return r
}

// @Summary homeHandler returns a string
// @ID homeHandler-roles
// @Produce json
// @Success 200 {object} Role
// @Failure 404 {object} string
// @Router /user/role [get]
func homeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response, err := json.Marshal("Role")
		if err != nil {
			fmt.Println("couldn't marshal")
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error marshalling")))
			return
		}
		_, err = w.Write([]byte(response))
		if err != nil {
			return
		}
	}
}

// @Summary getRoles returns a list of role object
// @ID get-roles
// @Produce json
// @Success 200 {object} Role
// @Failure 404 {object} string
// @Router /user/role/getAll [get]
func getRoles(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Authenticating user by checking the token.
		token := jwtauth.TokenFromHeader(r)
		controller_auth.IsAllowed(token, w, r)
		user := repository.GetRoles()
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

// @Summary create returns a role object
// @ID create-role
// @Produce json
// @Parameters:
//
//		-name: tags
//		 in: query
//	  required: true
//		 type: Role
//
// @Success 200 {object} Role
// @Failure 404 {object} string
// @Router /user/role/create [post]
func createRole(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Authenticating user by checking the token.
		token := jwtauth.TokenFromHeader(r)
		controller_auth.IsAllowed(token, w, r)
		data := &domain.Role{}
		err := render.Bind(r, data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		var role = repository.GetRoleObject(data)
		response := repository.CreateRole(role)
		if response.Id == "" {
			fmt.Println("error creating role")
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error creating role")))
			return
		}
		result, err := json.Marshal(repository.GetRoleObject(response))
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

// @Summary update returns a role object
// @ID update-role
// @Produce json
// @Parameters:
//
//		-name: tags
//		 in: query
//	  required: true
//		 type: Role
//
// @Success 200 {object} Role
// @Failure 404 {object} string
// @Router /user/role/update [post]
func update(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Authenticating user by checking the token.
		token := jwtauth.TokenFromHeader(r)
		controller_auth.IsAllowed(token, w, r)

		data := &domain.Role{}
		err := render.Bind(r, data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		var role = repository.GetRoleObject(data)
		response := repository.UpdateRole(role)
		if response.Id == "" {
			fmt.Println("error updating role")
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error updating role")))
			return
		}
		result, err := json.Marshal(repository.GetRoleObject(response))
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

// @Summary getRole returns a role object
// @ID get-role
// @Produce:  application/json
// @Parameters:
//
//		-name: tags
//		 in: query
//	  required: true
//		 type: Role
//
// @Success 200 {object} Role
// @Failure 404 {object} string
// @Router /user/role/get/{id} [get]
func getRole(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := jwtauth.TokenFromHeader(r)
		controller_auth.IsAllowed(token, w, r)
		id := chi.URLParam(r, "id")
		if id != "" {
			role := repository.GetRole(id)
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
