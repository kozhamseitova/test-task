package main

import (
	"log"

	"github.com/kozhamseitova/test-task/internal/app"
	"github.com/kozhamseitova/test-task/internal/config"
)

func main() {

	cfg, err := config.InitConfig("config.yaml")
	if err != nil {
		panic(err)
	}
	
	err = app.Run(cfg)

	if err != nil {
		log.Fatalf("failed to run app: %v", err)
	}

}