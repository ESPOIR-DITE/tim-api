package controller

import (
	"github.com/ESPOIR-DITE/tim-api/config"
	server_config "github.com/ESPOIR-DITE/tim-api/config/server.config"
	channelHomeController "github.com/ESPOIR-DITE/tim-api/controller/channel"
	userHomeController "github.com/ESPOIR-DITE/tim-api/controller/user.home.controller"
	"github.com/ESPOIR-DITE/tim-api/controller/video"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/swaggo/http-swagger"
	_ "github.com/swaggo/http-swagger/example/go-chi/docs"
	"html/template"
	"net/http"
)

// @title chi-swagger example APIs
// @version 1.0
// @description chi-swagger example APIs
// @BasePath: /tim-api/

func Controllers(env *server_config.Env) http.Handler {
	mux := chi.NewMux()
	mux.Use(config.CORS().Handler)
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Logger)
	mux.Use(env.Session.LoadAndSave)

	mux.Handle("/", homeHandler(env))
	mux.Mount("/video", video.Home(env))
	mux.Mount("/channel", channelHomeController.Home(env))
	mux.Mount("/user", userHomeController.Home(env))
	//mux.Handle("/board", Dashboard(env))
	fileServer := http.FileServer(http.Dir("./view/assets/"))
	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/assets/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	mux.Mount("/assets/", http.StripPrefix("/assets", fileServer))
	mux.Mount("/swagger/", httpSwagger.Handler(httpSwagger.URL("http://localhost:8081/swagger/doc.json")))
	return mux
}

//func Dashboard(env *config.Env) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		type Dashboard struct {
//			PendingUsers int64 `json:"pending_users"`
//			Users        int64 `json:"users"`
//			Videos       int64 `json:"videos"`
//			//Channels     int64 `json:"channels"`
//			UserStack domain.UserStack
//		}
//		userStack := user_repo.GetUserStack()
//		pendingUser := user_repo.GetAllPendingUsers()
//		//channels := role_repo.CountChannel()
//		videos := video_repo.CountVideo()
//		users := user_repo.CountUsers()
//
//		data := Dashboard{pendingUser, users, videos,
//			//channels,
//			userStack}
//		response, err := json.Marshal(data)
//		if err != nil {
//			env.ErrorLog.Printf("couldn't marshal %s", err)
//			render.Render(w, r, util.ErrInvalidRequest(errors.New("error marshalling")))
//			return
//		}
//		_, err = w.Write([]byte(response))
//		env.InfoLog.Println("read Dashboard success")
//		if err != nil {
//			return
//		}
//	}
//}

// @Summary homeHandler all items in the
// @ID get-all-todos
// @Produce json
// @Success 200
// @Router  / [get]
func homeHandler(app *server_config.Env) http.HandlerFunc {
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
