package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/swaggo/http-swagger"
	_ "github.com/swaggo/http-swagger/example/go-chi/docs"
	"html/template"
	"net/http"
	"tim-api/config"
	channel_controller "tim-api/controller/channel"
	"tim-api/controller/user"
	"tim-api/controller/util"
	"tim-api/controller/video"
	"tim-api/domain"
	role_repo "tim-api/storage/chanel/channel-repository"
	user_repo "tim-api/storage/user/user-repo"
	video_repo "tim-api/storage/video/video-repo"
)

// @title chi-swagger example APIs
// @version 1.0
// @description chi-swagger example APIs
// @BasePath /

func Controllers(env *config.Env) http.Handler {
	mux := chi.NewMux()
	mux.Use(config.CORS().Handler)
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Logger)
	mux.Use(env.Session.LoadAndSave)

	mux.Handle("/", homeHandler(env))
	mux.Mount("/video", video.Home(env))
	mux.Mount("/channel", channel_controller.Home(env))
	mux.Mount("/user", user.Home(env))
	mux.Handle("/system-set-up", setSystemSetUp(env))
	mux.Handle("/board", Dashboard(env))
	fileServer := http.FileServer(http.Dir("./view/assets/"))
	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/assets/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	mux.Mount("/assets/", http.StripPrefix("/assets", fileServer))
	mux.Mount("/swagger/", httpSwagger.Handler(httpSwagger.URL("http://localhost:8081/swagger/doc.json")))
	return mux
}

func Dashboard(env *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type Dashboard struct {
			PendingUsers int64 `json:"pending_users"`
			Users        int64 `json:"users"`
			Videos       int64 `json:"videos"`
			Channels     int64 `json:"channels"`
			UserStack    domain.UserStack
		}
		userStack := user_repo.GetUserStack()
		pendingUser := user_repo.GetAllPendingUsers()
		channels := role_repo.CountChannel()
		videos := video_repo.CountVideo()
		users := user_repo.CountUsers()

		data := Dashboard{pendingUser, users, videos, channels, userStack}
		response, err := json.Marshal(data)
		if err != nil {
			env.ErrorLog.Printf("couldn't marshal %s", err)
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error marshalling")))
			return
		}
		_, err = w.Write([]byte(response))
		env.InfoLog.Println("read Dashboard success")
		if err != nil {
			return
		}
	}
}

func setSystemSetUp(env *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type SetUp struct {
			TableSetUp []util.TableSetUpReport
		}
		data := SetUp{util.TableSetUp()}
		response, err := json.Marshal(data)
		if err != nil {
			fmt.Println("couldn't marshal")
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error marshalling")))
			return
		}
		env.InfoLog.Println("System setup")
		_, err = w.Write([]byte(response))
		if err != nil {
			return
		}
	}
}

// @Summary homeHandler all items in the todo list
// @ID get-all-todos
// @Produce json
// @Success 200 {object} todo
// @Router /todo [get]
func homeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		files := []string{
			app.Path + "index.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.Execute(w, nil)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}
}
