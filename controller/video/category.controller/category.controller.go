package categoryController

import (
	"encoding/json"
	"errors"
	server_config "github.com/ESPOIR-DITE/tim-api/config/server.config"
	"github.com/ESPOIR-DITE/tim-api/controller/util"
	"github.com/ESPOIR-DITE/tim-api/domain/video/category"
	categoryRepository "github.com/ESPOIR-DITE/tim-api/storage/video/category"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Home(app *server_config.Env) http.Handler {
	r := chi.NewRouter()
	repo := categoryRepository.NewCategoryRepository(app.GormDB)

	r.Get("/get/{id}", GetCategory(app, repo))
	r.Get("/delete/{id}", DeleteCategory(app, repo))
	r.Post("/create", CreateCategory(app, repo))
	r.Post("/update", UpdateCategory(app, repo))
	r.Get("/getAll", GetCategories(app, repo))
	return r
}

func DeleteCategory(app *server_config.Env, repo *categoryRepository.CategoryRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id != "" {
			category, err := repo.DeleteCategory(id)
			if err != nil {
				log.Errorf("Error deleting categoryId: %s err:%d", id, err)
				render.Render(w, r, util.InternalServeErr(errors.New("error deleting category")))
				return
			}
			result, err := json.Marshal(category)
			if err != nil {
				log.Errorf("couldn't marshal category response: %s", id)
				render.Render(w, r, util.ErrRender(errors.New("error marshalling")))
				return
			}
			_, err = w.Write([]byte(result))
			if err != nil {
				log.Errorf("couldn't render bytes")
				render.Render(w, r, util.ErrRender(errors.New("error marshalling")))
				return
			} else {
				app.InfoLog.Printf("Successfully deleted category: %v", category)
				return
			}
		}
	}
}

func GetCategories(app *server_config.Env, repo *categoryRepository.CategoryRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := repo.GetCategories()
		if err != nil {
			log.Errorf("couldn't read categories err: %d", err)
			render.Render(w, r, util.ErrInvalidRequest(errors.New("couldn't read categories")))
			return
		}
		result, err := json.Marshal(user)
		if err != nil {
			log.Errorf("couldn't render bytes")
			app.ErrorLog.Printf("couldn't render bytes")
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error marshalling")))
			return
		}
		_, err = w.Write([]byte(result))
		if err != nil {
			log.Errorf("couldn't render bytes")
			app.ErrorLog.Printf("couldn't render bytes")
			render.Render(w, r, util.ErrRender(errors.New("error marshalling")))
			return
		}
	}
}

func CreateCategory(app *server_config.Env, repo *categoryRepository.CategoryRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := category.Category{}
		err := render.Bind(r, &data)
		if err != nil {
			log.Errorf("couldn't render category data: %d", err)
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		response, err := repo.CreateCategory(data)
		if err != nil {
			log.Errorf("couldn't create category: %d", err)
			app.ErrorLog.Printf("couldn't create category: %d", data)
			render.Render(w, r, util.InternalServeErr(errors.New("error creating Category")))
			return
		}
		result, err := json.Marshal(response)
		if err != nil {
			app.ErrorLog.Printf("couldn't render bytes")
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error marshalling")))
			return
		}
		_, err = w.Write([]byte(result))
		if err != nil {
			app.ErrorLog.Printf("couldn't render bytes")
			render.Render(w, r, util.ErrRender(errors.New("writing bytes")))
			return
		}

		app.InfoLog.Printf("Successfully created category: %v", response)
		return

	}
}

func UpdateCategory(app *server_config.Env, repo *categoryRepository.CategoryRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &category.Category{}
		err := render.Bind(r, data)
		if err != nil {
			render.Render(w, r, util.ErrRender(err))
			return
		}
		response, err := repo.UpdateCategory(*data)
		if err != nil {
			app.ErrorLog.Println("couldn't updating category: ", data)
			log.Errorf("couldn't updating category: %d", err)
			render.Render(w, r, util.InternalServeErr(err))
			return
		}
		result, err := json.Marshal(response)
		if err != nil {
			app.ErrorLog.Printf("couldn't render bytes")
			log.Errorf("couldn't render bytes in updating")
			render.Render(w, r, util.ErrRender(errors.New("error marshalling")))
			return
		}
		_, err = w.Write([]byte(result))
		if err != nil {
			render.Render(w, r, util.ErrRender(errors.New("writing bytes")))
			return
		}
		app.InfoLog.Printf("Successfully updated category: %v", response)
		return
	}
}

func GetCategory(app *server_config.Env, repo *categoryRepository.CategoryRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id != "" {
			role, err := repo.GetCategory(id)
			if err != nil {
				app.ErrorLog.Printf("could not find Category with id: %s", id)
				log.Errorf("could not find Category with id: %d", err)
				render.Render(w, r, util.ErrResourceNotFind(err))
				return

			}
			result, err := json.Marshal(role)
			if err != nil {
				app.ErrorLog.Printf("err: %s could not render Category with id: %s", err, id)
				render.Render(w, r, util.ErrRender(errors.New("error marshalling")))
				return
			}
			app.InfoLog.Printf("read successful %s", id)
			_, err = w.Write([]byte(result))
			if err != nil {
				return
			}
		}
		app.ErrorLog.Printf("error could not read Category with id: %s", id)
		render.Render(w, r, util.ErrInvalidRequest(errors.New("missing params")))
		return
	}
}
