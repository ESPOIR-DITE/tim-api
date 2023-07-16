package channelVideoController

import (
	"encoding/json"
	"errors"
	"fmt"
	server_config "github.com/ESPOIR-DITE/tim-api/config/server.config"
	"github.com/ESPOIR-DITE/tim-api/controller/util"
	channel_video "github.com/ESPOIR-DITE/tim-api/domain/channel/channel-video"
	repository "github.com/ESPOIR-DITE/tim-api/storage/chanel/channel.video.repository"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Home(app *server_config.Env) http.Handler {
	r := chi.NewRouter()
	repo := repository.NewChannelVideoRepository(app.GormDB)

	r.Get("/get/{id}", getChannel(app, repo))
	r.Get("/delete/{id}", deleteChannel(app, repo))
	//r.Get("/get-by-video.controller/{videoId}", getChannelsByVideoId(app, repo))
	r.Get("/get-by-channel.controller/{channelId}", getChannelsByChannel(app, repo))
	r.Post("/create", createChannel(app, repo))
	r.Post("/update", updateChannel(app, repo))
	r.Get("/", getChannels(app, repo))

	return r
}

func deleteChannel(app *server_config.Env, repo *repository.ChannelVideoRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id != "" {
			role, err := repo.DeleteChannelVideo(id)
			if err != nil {
				fmt.Println("couldn't marshal")
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

func getChannelsByChannel(app *server_config.Env, repo *repository.ChannelVideoRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "channelId")
		if id != "" {
			channels, err := repo.GetChannelVideosByChannelId(id)
			if err != nil {
				log.Errorf("error retrieving channels %d", err)
				render.Render(w, r, util.InternalServeErr(err))
				return
			}
			result, err := json.Marshal(channels)
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

//	func getChannelsByVideoId(app *config.Env, repo *repository.ChannelVideoRepository) http.HandlerFunc {
//		return func(w http.ResponseWriter, r *http.Request) {
//			id := chi.URLParam(r, "videoId")
//			if id != "" {
//				role.domain.controller := repo.GetChannelVideosByVideoId(id)
//				result, err := json.Marshal(role.domain.controller)
//				if err != nil {
//					fmt.Println("couldn't marshal")
//					render.Render(w, r, util.ErrInvalidRequest(errors.New("error marshalling")))
//					return
//				}
//				_, err = w.Write([]byte(result))
//				if err != nil {
//					return
//				}
//			}
//		}
//	}
func updateChannel(app *server_config.Env, repo *repository.ChannelVideoRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := channel_video.ChannelVideos{}
		err := render.Bind(r, &data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		response, err := repo.UpdateChannelVideo(data)
		if err != nil {
			fmt.Println("error creating role.domain.controller")
			render.Render(w, r, util.InternalServeErr(err))
			return
		}
		result, err := json.Marshal(response)
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
func getChannel(app *server_config.Env, repo *repository.ChannelVideoRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		channelVideo, err := repo.GetChannelVideo(id)
		if err != nil {
			render.Render(w, r, util.InternalServeErr(err))
			return
		}
		result, err := json.Marshal(channelVideo)
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
func createChannel(app *server_config.Env, repo *repository.ChannelVideoRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, claims, _ := jwtauth.FromContext(r.Context())
		fmt.Println(claims)
		data := channel_video.ChannelVideos{}
		err := render.Bind(r, &data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		response, err := repo.CreateChannelVideo(data)
		if err != nil {
			fmt.Println("error creating video.controller channel.controller")
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error creating video.controller channel.controller")))
			return
		}
		result, err := json.Marshal(response)
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
func getChannels(app *server_config.Env, repo *repository.ChannelVideoRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		channels, err := repo.GetChannelVideos()
		if err != nil {
			log.Errorf("error retrieving channels %d", err)
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error marshalling")))
			return
		}
		result, err := json.Marshal(channels)
		if err != nil {
			log.Errorf("couldn't marshal channels %d", err)
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error marshalling")))
			return
		}
		_, err = w.Write([]byte(result))
		if err != nil {
			return
		}
	}
}
