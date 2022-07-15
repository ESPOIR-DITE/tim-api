package category_controller

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
	"tim-api/config"
	"tim-api/controller/util"
	"tim-api/domain"
	repository "tim-api/storage/video/category"
)

func Home(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/get/{id}", GetCategory(app))
	r.Get("/delete/{id}", DeleteCategory(app))
	r.Post("/create", CreateCategory(app))
	r.Post("/update", UpdateCategory(app))
	r.Get("/getAll", GetCategories(app))
	return r
}

func DeleteCategory(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id != "" {
			category := repository.DeleteCategory(id)
			if category == false {
				app.ErrorLog.Printf("Error deleting category: %s", id)
				render.Render(w, r, util.ErrInvalidRequest(errors.New("error marshalling")))
				return
			}
			result, err := json.Marshal(category)
			if err != nil {
				app.ErrorLog.Printf("couldn't marshal category response: %s", id)
				render.Render(w, r, util.ErrRender(errors.New("error marshalling")))
				return
			}
			_, err = w.Write([]byte(result))
			if err != nil {
				app.ErrorLog.Printf("couldn't render bytes")
				render.Render(w, r, util.ErrRender(errors.New("error marshalling")))
				return
			} else {
				app.InfoLog.Printf("Successfully deleted category: %v", category)
				return
			}
		}
	}
}

func GetCategories(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := repository.GetCategories()
		if len(user) == 0 {
			app.ErrorLog.Println("couldn't read categories")
			render.Render(w, r, util.ErrInvalidRequest(errors.New("reading")))
			return
		}
		result, err := json.Marshal(user)
		if err != nil {
			app.ErrorLog.Printf("couldn't render bytes")
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error marshalling")))
			return
		}
		_, err = w.Write([]byte(result))
		if err != nil {
			app.ErrorLog.Printf("couldn't render bytes")
			render.Render(w, r, util.ErrRender(errors.New("error marshalling")))
			return
		} else {
			app.InfoLog.Println("Successfully read categories")
			return
		}
	}
}

func CreateCategory(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := domain.Category{}
		err := render.Bind(r, &data)
		if err != nil {
			app.ErrorLog.Println("couldn't render category data: ", data)
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		response := repository.CreateCategory(data)
		if response.Id == "" {
			app.ErrorLog.Println("couldn't create category: ", data)
			render.Render(w, r, util.InternalServeErr(errors.New("error creating Category")))
			return
		}
		result, err := json.Marshal(repository.GetCategoryObject(response))
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

func UpdateCategory(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &domain.Category{}
		err := render.Bind(r, data)
		if err != nil {
			render.Render(w, r, util.ErrRender(err))
			return
		}
		categoryObject := repository.GetCategoryObject(data)
		response := repository.UpdateCategory(categoryObject)
		if response.Id == "" {
			app.ErrorLog.Println("couldn't updating category: ", data)
			render.Render(w, r, util.InternalServeErr(errors.New("error updating Category")))
			return
		}
		result, err := json.Marshal(repository.GetCategoryObject(response))
		if err != nil {
			app.ErrorLog.Printf("couldn't render bytes")
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

func GetCategory(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id != "" {
			role := repository.GetCategory(id)
			if role.Id == "" {
				app.ErrorLog.Printf("could not find Category with id: %s", id)
				render.Render(w, r, util.ErrRecourseNotFind(errors.New("error marshalling")))
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
		render.Render(w, r, util.ErrRecourseNotFind(errors.New("missing params")))
		return
	}
}
