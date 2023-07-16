package channelHomeController

import (
	server_config "github.com/ESPOIR-DITE/tim-api/config/server.config"
	channelController "github.com/ESPOIR-DITE/tim-api/controller/channel/channel.controller"
	channelSubscriptionController "github.com/ESPOIR-DITE/tim-api/controller/channel/channel.subscription.controller"
	channelTypeController "github.com/ESPOIR-DITE/tim-api/controller/channel/channel.type.controller"
	channelVideo "github.com/ESPOIR-DITE/tim-api/controller/channel/channel.video.controller"
	"github.com/go-chi/chi"
	"net/http"
)

func Home(app *server_config.Env) http.Handler {
	mux := chi.NewMux()
	mux.Handle("/", homeHandler(app))
	mux.Mount("/channel", channelController.Home(app))
	mux.Mount("/channel-video", channelVideo.Home(app))
	mux.Mount("/channel-type", channelTypeController.Home(app))
	mux.Mount("/channel-subscription", channelSubscriptionController.Home(app))
	return mux
}

func homeHandler(app *server_config.Env) http.Handler {
	return nil
}
