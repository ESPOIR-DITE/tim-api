package video

import (
	server_config "github.com/ESPOIR-DITE/tim-api/config/server.config"
	categoryController "github.com/ESPOIR-DITE/tim-api/controller/video/category.controller"
	videoDataTest "github.com/ESPOIR-DITE/tim-api/controller/video/video-data-test"
	videoCategoryHomeController "github.com/ESPOIR-DITE/tim-api/controller/video/video.category.controller"
	videoController "github.com/ESPOIR-DITE/tim-api/controller/video/video.controller"
	videoDataHomeController "github.com/ESPOIR-DITE/tim-api/controller/video/video.data.controller"
	"github.com/go-chi/chi"
	"net/http"
)

func Home(app *server_config.Env) http.Handler {
	mux := chi.NewMux()
	mux.Handle("/", homeHandler(app))
	mux.Mount("/category", categoryController.Home(app))
	mux.Mount("/video-category", videoCategoryHomeController.Home(app))
	mux.Mount("/video-data", videoDataHomeController.Home(app))
	mux.Mount("/video", videoController.Home(app))
	mux.Mount("/video-data-test", videoDataTest.Home(app))
	//mux.Mount("/comment", videoCommentController.Home(app))
	return mux
}

func homeHandler(app *server_config.Env) http.Handler {
	return nil
}
