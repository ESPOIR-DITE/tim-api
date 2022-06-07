package video_data

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
	"os"
	"tim-api/config"
	"tim-api/controller/util"
	"tim-api/domain"
	repository "tim-api/storage/video/video-data"
)

func Home(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/get/{id}", get(app))
	r.Get("/delete/{id}", delete(app))
	r.Post("/create", create(app))
	r.Post("/update", update(app))
	r.Get("/getAll", getAll(app))
	r.Get("/getRwa", getRaw(app))
	return r
}

func getRaw(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "files/test/upload-196722650.mp4")
	}
}

func delete(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id != "" {
			role := repository.DeleteVideoData(id)
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

func getAll(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := repository.GetVideoDatas()
		result, err := json.Marshal(user)
		if err != nil {
			fmt.Println("couldn't marshal")
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error marshalling")))
			return
		}
		_, err = w.Write([]byte(result))
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error writing bytes")))
			return
		}
	}
}

func create(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &domain.VideoDate{}
		err := render.Bind(r, data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		object := repository.GetCategoryObject(data)
		if len(object.Video) == 0 {
			fmt.Println(object.Id)
		} else {
			fmt.Println()
		}
		go util.VideoWriter(data.Id, data.Video, data.FileType)

		//response := repository.CreateVideoData(object)
		//if response.Id == "" {
		//	fmt.Println("error creating videoData")
		//	render.Render(w, r, util.ErrInvalidRequest(errors.New("error creating videoData")))
		//	return
		//}
		//result, err := json.Marshal(repository.GetCategoryObject(response))
		//if err != nil {
		//	fmt.Println("couldn't marshal")
		//	render.Render(w, r, util.ErrInvalidRequest(errors.New("error marshalling")))
		//	return
		//}
		//_, err = w.Write([]byte(result))
		//if err != nil {
		//	render.Render(w, r, util.ErrInvalidRequest(errors.New("writing bytes")))
		//	return
		//}
	}
}

func update(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &domain.VideoDate{}
		err := render.Bind(r, data)
		if err != nil {
			render.Render(w, r, util.ErrInvalidRequest(err))
			return
		}
		object := repository.GetCategoryObject(data)
		response := repository.UpdateVideoDate(object)
		if response.Id == "" {
			fmt.Println("error creating VideoData")
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error creating Video")))
			return
		}
		result, err := json.Marshal(repository.GetCategoryObject(response))
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

func readVideoFile(id string, extension string) ([]byte, error) {
	//file, err := os.ReadFile("files/test/"+id+""+extension)
	file, err := os.ReadFile("files/test/upload-196722650.mp4")
	if err != nil {
		fmt.Println(err, " error reading file!")
		return nil, err
	}
	return file, nil
}

func get(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id != "" {
			//object := repository.GetVideoDate(id)
			file, err := readVideoFile(id, "")
			result, err := json.Marshal(file)
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
