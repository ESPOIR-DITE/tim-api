package video

import (
	"github.com/go-chi/chi"
	"net/http"
	"tim-api/config"
	channel_controller "tim-api/controller/channel/channel"
	channel_subscription_controller "tim-api/controller/channel/channel-subscription"
	channel_type_controller "tim-api/controller/channel/channel-type"
	channel_video_controller "tim-api/controller/channel/channel-video"
)

func Home(app *config.Env) http.Handler {
	mux := chi.NewMux()
	mux.Handle("/", homeHandler(app))
	mux.Mount("/channel", channel_controller.Home(app))
	mux.Mount("/channel-video", channel_video_controller.Home(app))
	mux.Mount("/channel-type", channel_type_controller.Home(app))
	mux.Mount("/channel-subscription", channel_subscription_controller.Home(app))
	return mux
}

func homeHandler(app *config.Env) http.Handler {
	return nil
}
