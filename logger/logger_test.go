package logger

import (
	"fmt"
	env2 "github.com/ESPOIR-DITE/tim-api/config/tim_api/env"

	"testing"
)

func TestInit(t *testing.T) {
	env := env2.NewTimApiConfigurationManagerImpl()

	config, err := env.Load()
	if err != nil {
		fmt.Println(err)
	}
	LogInit(config)
	Log.Fatalf("testing")
}
