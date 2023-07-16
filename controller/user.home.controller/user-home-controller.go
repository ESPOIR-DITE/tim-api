package userHomeController

import (
	"encoding/json"
	"errors"
	"fmt"
	server_config "github.com/ESPOIR-DITE/tim-api/config/server.config"
	accountController "github.com/ESPOIR-DITE/tim-api/controller/user.home.controller/account.controller"
	roleController "github.com/ESPOIR-DITE/tim-api/controller/user.home.controller/role.controller"
	userAccountController "github.com/ESPOIR-DITE/tim-api/controller/user.home.controller/user.account.controller"
	userBankController "github.com/ESPOIR-DITE/tim-api/controller/user.home.controller/user.bank.controller"
	userController "github.com/ESPOIR-DITE/tim-api/controller/user.home.controller/user.controller"
	userDetailController "github.com/ESPOIR-DITE/tim-api/controller/user.home.controller/user.detail.controller"
	user_subscription "github.com/ESPOIR-DITE/tim-api/controller/user.home.controller/user.subscription.controller"
	userVideoController "github.com/ESPOIR-DITE/tim-api/controller/user.home.controller/user.video.controller"
	"github.com/ESPOIR-DITE/tim-api/controller/util"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
)

func Home(app *server_config.Env) http.Handler {
	mux := chi.NewMux()
	mux.Handle("/", homeHandler(app))
	mux.Mount("/role", roleController.Home(app))
	mux.Mount("/user", userController.Home(app))
	mux.Mount("/account", accountController.Home(app))
	mux.Mount("/user-account", userAccountController.Home(app))
	mux.Mount("/user-subscription", user_subscription.Home(app))
	mux.Mount("/user-video", userVideoController.Home(app))
	mux.Mount("/user-detail", userDetailController.Home(app))
	mux.Mount("/user-bank", userBankController.Home(app))

	return mux
}

func homeHandler(app *server_config.Env) http.HandlerFunc {
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
