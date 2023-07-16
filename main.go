package main

import (
	"flag"
	"github.com/ESPOIR-DITE/tim-api/config/server.config"
	env2 "github.com/ESPOIR-DITE/tim-api/config/tim_api/env"
	"github.com/ESPOIR-DITE/tim-api/controller"
	"github.com/ESPOIR-DITE/tim-api/logger"
	"github.com/ESPOIR-DITE/tim-api/s3"
	"github.com/ESPOIR-DITE/tim-api/security"
	"github.com/ESPOIR-DITE/tim-api/storage/databases/postgres/config/connector"
	migration2 "github.com/ESPOIR-DITE/tim-api/storage/databases/postgres/config/migration"
	"github.com/alexedwards/scs/v2"
	"github.com/golang-migrate/migrate/v4"
	log "github.com/sirupsen/logrus"
	_ "github.com/swaggo/http-swagger/example/go-chi/docs"
	"gorm.io/gorm"
	"net/http"
	"os"
	"time"
)

var sessionManager *scs.SessionManager

func Environment(gormDb *gorm.DB, bucket *s3.S3Bucket) *server_config.Env {
	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.IdleTimeout = 20 * time.Minute
	var jwtkey = []byte("supesecretKey")
	securityService := security.NewJWTService(jwtkey)
	env := &server_config.Env{
		Path:            "./view/html/",
		Session:         sessionManager,
		GormDB:          gormDb,
		SecurityService: securityService,
		PhotoS3Bucket:   bucket,
	}

	return env
}

func main() {
	env := env2.NewTimApiConfigurationManagerImpl()

	config, err := env.Load()
	if err != nil {
		log.Fatal("Failed to load App config: %w", err.Error())
	}
	if err := logger.LogInit(config); err != nil {
		log.Fatal("Failed to configure logger: %w", err.Error())
		os.Exit(2)
	}
	gormDb, err := connector.NewPostgresDBConnector(config.DBConfig()).Connect()
	if err != nil {
		logger.Log.Fatal(err.Error())
		os.Exit(2)
	}

	migration, err := migration2.NewPostgresMigrator(config.DBConfig()).NewMigrator()
	if err != nil {
		logger.Log.Fatal(err.Error())
		os.Exit(2)
	}
	err = migration.Up()
	if err != migrate.ErrNoChange && err != nil {
		logger.Log.Fatalf("Failed to run db migration:%s", err.Error())
		os.Exit(2)
	}

	photoS3Bucket := s3.NewS3Bucket(config.S3Config().Region(), config.S3Config())
	if err := photoS3Bucket.Init(); err != nil {
		os.Exit(2)
	}

	addr := flag.String("addr", ":8081", "HTTP network address")
	flag.Parse()
	srv := &http.Server{
		Addr: *addr,
		//ErrorLog: Environment(nil, nil).ErrorLog,
		Handler: controller.Controllers(Environment(gormDb, photoS3Bucket)),
	}

	log.Infof("Starting server on %s", *addr)
	//Environment().InfoLog.Printf("Starting server on %s", *addr)
	// Call the ListenAndServe() method on our new http.Server struct.
	error := srv.ListenAndServe()
	Environment(nil, nil).ErrorLog.Fatal(error)

}
