package channelSubscriptionController

import (
	"encoding/json"
	"errors"
	"fmt"
	server_config "github.com/ESPOIR-DITE/tim-api/config/server.config"
	"github.com/ESPOIR-DITE/tim-api/controller/util"
	channelSubscriptionDomain "github.com/ESPOIR-DITE/tim-api/domain/channel/channel.subscription.domain"
	channelSubscriptionRepository "github.com/ESPOIR-DITE/tim-api/storage/chanel/channel.subscription.repository"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Home(app *server_config.Env) http.Handler {
	r := chi.NewRouter()
	repo := channelSubscriptionRepository.NewChannelSubscriptionRepository(app.GormDB)

	r.Get("/get/{id}", getChannelSubscription(app, repo))
	r.Get("/get-by-channel.controller/{id}", countChannelSubscriptionByChannel(app, repo))
	r.Get("/delete/{id}", deleteChannelSubscription(app, repo))
	r.Get("/get-by-user.home.controller.domain.controller/{userId}", getChannelSubscriptionsByUser(app, repo))
	r.Get("/get-by-channel.controller/{channelId}", getChannelsByChannel(app, repo))
	r.Post("/create", createChannelSubscription(app, repo))
	r.Post("/update", updateChannel(app, repo))
	r.Get("/getAll", getChannelSubscriptions(app, repo))
	return r
}

func countChannelSubscriptionByChannel(app *server_config.Env, repo *channelSubscriptionRepository.ChannelSubscriptionRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		user, err := repo.CountSubscriptionByChannelId(id)
		if err != nil {
			fmt.Println("couldn't marshal")
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error marshalling")))
			return
		}
		result, err := json.Marshal(user)
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
func deleteChannelSubscription(app *server_config.Env, repo *channelSubscriptionRepository.ChannelSubscriptionRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id != "" {
			role, err := repo.DeleteChannelSubscription(id)
			if err != nil {
				log.Error(err)
				render.Render(w, r, util.InternalServeErr(err))
				return
			}
			result, err := json.Marshal(role)
			if err != nil {
				log.Error(err)
				render.Render(w, r, util.ErrInvalidRequest(errors.New("error marshalling")))
				return
			}
			_, err = w.Write([]byte(result))
			if err != nil {
				return
			}
		}
		render.Render(w, r, util.ErrInvalidRequest(errors.New("missing required field")))
		return
	}
}
func getChannelsByChannel(app *server_config.Env, repo *channelSubscriptionRepository.ChannelSubscriptionRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "channelId")
		if id != "" {
			role, err := repo.GetChannelSubscriptionsByChannelId(id)
			if err != nil {
				log.Error(err)
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
		render.Render(w, r, util.ErrInvalidRequest(errors.New("missing required field")))
		return
	}
}
func getChannelSubscriptionsByUser(app *server_config.Env, repo *channelSubscriptionRepository.ChannelSubscriptionRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "userId")
		if id != "" {
			role, err := repo.GetChannelSubscriptionsByUser(id)
			if err != nil {
				log.Error(err)
				render.Render(w, r, util.InternalServeErr(err))
				return
			}
			result, err := json.Marshal(role)
			if err != nil {
				render.Render(w, r, util.ErrInvalidRequest(errors.New("error marshalling")))
				return
			}
			_, err = w.Write([]byte(result))
			if err != nil {
				return
			}
		}
		render.Render(w, r, util.ErrInvalidRequest(errors.New("missing required field")))
		return
	}
}
func updateChannel(app *server_config.Env, repo *channelSubscriptionRepository.ChannelSubscriptionRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := channelSubscriptionDomain.ChannelSubscription{}
		err := render.Bind(r, &data)
		if err != nil {
			log.Error(err)
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		response, err := repo.UpdateChannelSubscription(data)
		if err != nil {
			log.Errorf("error updating channel subscription: %d", err)
			render.Render(w, r, util.InternalServeErr(errors.New("error updating channel.controller-subscription")))
			return
		}
		result, err := json.Marshal(response)
		if err != nil {
			marshalError := fmt.Sprintf("couldn't marshal updating channel subscription struct, error: %d", err)
			log.Errorf(marshalError)
			render.Render(w, r, util.InternalServeErr(errors.New(marshalError)))
			return
		}
		_, err = w.Write([]byte(result))
		if err != nil {
			writingError := fmt.Sprintf("Error writing bytes, error: %d", err)
			log.Errorf(writingError)
			render.Render(w, r, util.ErrInvalidRequest(errors.New(writingError)))
			return
		}
	}
}
func getChannelSubscription(app *server_config.Env, repo *channelSubscriptionRepository.ChannelSubscriptionRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id != "" {
			user, err := repo.GetChannelSubscription(id)
			if err != nil {
				log.Error(err)
				render.Render(w, r, util.InternalServeErr(err))
				return
			}
			result, err := json.Marshal(user)
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
		render.Render(w, r, util.ErrInvalidRequest(errors.New("missing required field")))
		return
	}
}
func createChannelSubscription(app *server_config.Env, repo *channelSubscriptionRepository.ChannelSubscriptionRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := channelSubscriptionDomain.ChannelSubscription{}
		err := render.Bind(r, &data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		response, err := repo.CreateChannelSubscription(data)
		if err != nil {
			createError := fmt.Sprintf("error creating channel subscription err: %d", err)
			log.Error(createError)
			render.Render(w, r, util.InternalServeErr(errors.New(createError)))
			return
		}
		result, err := json.Marshal(response)
		if err != nil {
			log.Error(err)
			render.Render(w, r, util.InternalServeErr(errors.New("error marshalling")))
			return
		}
		_, err = w.Write([]byte(result))
		if err != nil {
			log.Error(err)
			render.Render(w, r, util.InternalServeErr(errors.New("writing bytes")))
			return
		}
	}
}
func getChannelSubscriptions(app *server_config.Env, repo *channelSubscriptionRepository.ChannelSubscriptionRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		channels, err := repo.GetChannelSubscriptions()
		if err != nil {
			log.Error(err)
			render.Render(w, r, util.InternalServeErr(errors.New("error marshalling")))
			return
		}
		result, err := json.Marshal(channels)
		if err != nil {
			log.Error(err)
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error marshalling")))
			return
		}
		_, err = w.Write([]byte(result))
		if err != nil {
			return
		}
	}
}
