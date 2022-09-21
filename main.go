package main

import (
	"flag"
	"github.com/alexedwards/scs/v2"
	_ "github.com/swaggo/http-swagger/example/go-chi/docs"
	"log"
	"net/http"
	"os"
	"tim-api/config"
	"tim-api/controller"
	"time"
)

var sessionManager *scs.SessionManager

func Environment() *config.Env {
	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.IdleTimeout = 20 * time.Minute
	warningFileName := time.Now().Format("2006-01-02") + "-WARNING-LOGS.txt"
	warningFile, err := os.OpenFile(warningFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	errorFileName := time.Now().Format("2006-01-02") + "-ERROR-LOGS.txt"
	errorFile, err := os.OpenFile(errorFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	env := &config.Env{
		WarningLog: log.New(warningFile, "WARNING\t", log.Ldate|log.Ltime|log.Lshortfile),
		ErrorLog:   log.New(errorFile, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		InfoLog:    log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile),
		Path:       "./view/html/",
		Session:    sessionManager,
	}
	return env
}

func main() {
	addr := flag.String("addr", ":8081", "HTTP network address")
	flag.Parse()
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: Environment().ErrorLog,
		Handler:  controller.Controllers(Environment()),
	}

	Environment().InfoLog.Printf("Starting server on %s", *addr)
	// Call the ListenAndServe() method on our new http.Server struct.
	error := srv.ListenAndServe()
	Environment().ErrorLog.Fatal(error)

}
