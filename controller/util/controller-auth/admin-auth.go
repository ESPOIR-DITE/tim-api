package controller_auth

import (
	"errors"
	"github.com/go-chi/render"
	"net/http"
	"tim-api/controller/util"
	"tim-api/security"
)

func IsAllowed(token string, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := security.ValidateToken(token)
		if err != nil {
			render.Render(w, r, util.ErrRecourseNotAllowed(errors.New("Please try to read again.")))
			return
		}

	}
}
