package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"html/template"
	"net/http"
	"tim-api/config"
	"tim-api/controller/user"
	"tim-api/controller/util"
	"tim-api/controller/video"
)

func Controllers(env *config.Env) http.Handler {
	mux := chi.NewMux()
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Logger)
	mux.Use(env.Session.LoadAndSave)

	mux.Handle("/", homeHandler(env))
	mux.Mount("/video", video.Home(env))
	mux.Mount("/user", user.Home(env))
	mux.Handle("/system-set-up", setSystemSetUp(env))
	fileServer := http.FileServer(http.Dir("./view/assets/"))
	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/assets/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	mux.Mount("/assets/", http.StripPrefix("/assets", fileServer))
	return mux
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
		_, err = w.Write([]byte(response))
		if err != nil {
			return
		}
	}
}

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
