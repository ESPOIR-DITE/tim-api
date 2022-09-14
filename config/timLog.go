package config

import (
	"fmt"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"log"
	"time"
)

func InitLogger() {
	path := "./old.log"
	writer, err := rotatelogs.New(
		fmt.Sprintf("%s.%s", path, "%Y-%m-%d.%H:%M:%S"),
		rotatelogs.WithLinkName("./current.log"),
		rotatelogs.WithMaxAge(time.Second*10),
		rotatelogs.WithRotationTime(time.Second*1),
	)
	if err != nil {
		log.Fatalf("Failed to Initialize Log File %s", err)
	}
	log.SetOutput(writer)
	return
}
