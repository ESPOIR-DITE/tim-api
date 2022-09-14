package config

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestInitLogger(t *testing.T) {
	InitLogger()
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 1)
		log.Printf("Hello, World!")
	}
	fmt.Scanln()
}
