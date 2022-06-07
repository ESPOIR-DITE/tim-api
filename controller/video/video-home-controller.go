package video

import (
	"github.com/go-chi/chi"
	"net/http"
	"tim-api/config"
	videoController "tim-api/controller/video/video"
	categoryController "tim-api/controller/video/video-category"
	videoCategoryController "tim-api/controller/video/video-category"
	videoCommentController "tim-api/controller/video/video-comment"
	video_data "tim-api/controller/video/video-data"
)

func Home(app *config.Env) http.Handler {
	mux := chi.NewMux()
	mux.Handle("/", homeHandler(app))
	mux.Mount("/category", categoryController.Home(app))
	mux.Mount("/video-category", videoCategoryController.Home(app))
	mux.Mount("/video-data", video_data.Home(app))
	mux.Mount("/video", videoController.Home(app))
	mux.Mount("/comment", videoCommentController.Home(app))
	return mux
}

func homeHandler(app *config.Env) http.Handler {
	return nil
}
