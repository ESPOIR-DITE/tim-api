package server_config

import (
	"fmt"
	"github.com/ESPOIR-DITE/tim-api/s3"
	"github.com/ESPOIR-DITE/tim-api/security"
	"github.com/alexedwards/scs/v2"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"runtime/debug"
)

type Env struct {
	WarningLog      *log.Logger
	ErrorLog        *log.Logger
	InfoLog         *log.Logger
	Path            string
	Session         *scs.SessionManager
	GormDB          *gorm.DB
	SecurityService *security.JWTService
	PhotoS3Bucket   *s3.S3Bucket
}

type templateData struct {
	Title string
	Data  map[string]string
}

func (app *Env) ClientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// For consistency, we'll also implement a notFound helper. This is simply a
// convenience wrapper around clientError which sends a 404 Not Found response to
// the user.home.controller.domain.controller.
func (app *Env) NotFound(w http.ResponseWriter) {
	app.ClientError(w, http.StatusNotFound)
}

func (app *Env) ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Fatal(2, trace)
	//app.ErrorLog.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
