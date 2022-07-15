package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
	"tim-api/config"
	roleController "tim-api/controller/user/role"
	userController "tim-api/controller/user/user"
	userAccountController "tim-api/controller/user/user-account"
	userSubscriptionController "tim-api/controller/user/user-account"
	userBank "tim-api/controller/user/user-bank"
	userDetails "tim-api/controller/user/user-detail"
	userVideoController "tim-api/controller/user/user-video"
	"tim-api/controller/util"
)

func Home(app *config.Env) http.Handler {
	mux := chi.NewMux()
	mux.Handle("/", homeHandler(app))
	mux.Mount("/role", roleController.Home(app))
	mux.Mount("/user", userController.Home(app))
	mux.Mount("/user-account", userAccountController.Home(app))
	mux.Mount("/user-subscription", userSubscriptionController.Home(app))
	mux.Mount("/user-video", userVideoController.Home(app))
	mux.Mount("/user-detail", userDetails.Home(app))
	mux.Mount("/user-bank", userBank.Home(app))

	return mux
}

func homeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response, err := json.Marshal("data")
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
