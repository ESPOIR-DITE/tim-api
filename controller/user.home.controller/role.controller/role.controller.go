package roleController

import (
	"encoding/json"
	"errors"
	"fmt"
	server_config "github.com/ESPOIR-DITE/tim-api/config/server.config"
	"github.com/ESPOIR-DITE/tim-api/controller/util"
	controller_auth "github.com/ESPOIR-DITE/tim-api/controller/util/controller-auth"
	roleDomain "github.com/ESPOIR-DITE/tim-api/domain/user/role.domain"
	roleRepository "github.com/ESPOIR-DITE/tim-api/storage/user/role.repository"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Home(app *server_config.Env) http.Handler {
	r := chi.NewRouter()
	repo := roleRepository.NewRoleRepository(app.GormDB)
	r.Handle("/", homeHandler(app, repo))
	r.Get("/get/{id}", getRole(app, repo))
	r.Post("/create", createRole(app, repo))
	r.Get("/getAll", getRoles(app, repo))

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
// @Router /user.home.controller.domain.controller/role.domain.controller [get]
func homeHandler(app *server_config.Env, repo *roleRepository.RoleRepository) http.HandlerFunc {
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

// @Summary getRoles returns a list of role.domain.controller object
// @ID get-roles
// @Produce json
// @Success 200 {object} Role
// @Failure 404 {object} string
// @Router /user.home.controller.domain.controller/role.domain.controller/getAll [get]
func getRoles(app *server_config.Env, repo *roleRepository.RoleRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Authenticating user.home.controller.domain.controller by checking the token.
		token := jwtauth.TokenFromHeader(r)
		_, err := controller_auth.IsAllowed(app, token)
		if err != nil {
			render.Render(w, r, util.ErrRecourseNotAllowed(err))
			return
		}
		user, err := repo.GetRoles()
		if err != nil {
			log.Errorf("fail to retreive role.domain err: %d", err)
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
			return
		}
	}
}

// @Summary create returns a role.domain.controller object
// @ID create-role.domain.controller
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
// @Router /user.home.controller.domain.controller/role.domain.controller/create [post]
func createRole(app *server_config.Env, repo *roleRepository.RoleRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Authenticating user.home.controller.domain.controller by checking the token.
		token := jwtauth.TokenFromHeader(r)
		_, err := controller_auth.IsAllowed(app, token)
		if err != nil {
			render.Render(w, r, util.ErrRecourseNotAllowed(err))
			return
		}
		data := &roleDomain.Role{}
		err = render.Bind(r, data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		response, err := repo.CreateRole(*data)
		if err != nil {
			fmt.Println("error creating role err: %d", err)
			log.Errorf("error creating role err: %d", err)
			render.Render(w, r, util.InternalServeErr(errors.New("error creating role")))
			return
		}
		result, err := json.Marshal(&response)
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

// @Summary update returns a role.domain.controller object
// @ID update-role.domain.controller
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
// @Router /user.home.controller.domain.controller/role.domain.controller/update [post]
func update(app *server_config.Env, repo *roleRepository.RoleRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Authenticating user.home.controller.domain.controller by checking the token.
		token := jwtauth.TokenFromHeader(r)
		_, err := controller_auth.IsAllowed(app, token)
		if err != nil {
			render.Render(w, r, util.ErrRecourseNotAllowed(err))
			return
		}

		data := &roleDomain.Role{}
		err = render.Bind(r, data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		response, err := repo.UpdateRole(*data)
		if err != nil {
			log.Errorf("fail to update role err: %d", err)
			render.Render(w, r, util.InternalServeErr(errors.New("error updating role")))
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

// @Summary getRole returns a role.domain.controller object
// @ID get-role.domain.controller
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
// @Router /user.home.controller.domain.controller/role.domain.controller/get/{id} [get]
func getRole(app *server_config.Env, repo *roleRepository.RoleRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := jwtauth.TokenFromHeader(r)
		_, err := controller_auth.IsAllowed(app, token)
		if err != nil {
			render.Render(w, r, util.ErrRecourseNotAllowed(err))
			return
		}
		id := chi.URLParam(r, "id")
		if id != "" {
			role, err := repo.GetRole(id)
			if err != nil {
				log.Error(err)
				render.Render(w, r, util.InternalServeErr(err))
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
