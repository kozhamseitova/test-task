package main

import (
	"context"
	"log"

	"github.com/kozhamseitova/test-task/internal/app"
	"github.com/kozhamseitova/test-task/internal/config"
)

func main() {
	ctx := context.Background()
	cfg, err := config.InitConfig("config.yaml")
	if err != nil {
		panic(err)
	}
	
	err = app.Run(ctx, cfg)
	if err != nil {
		log.Fatalf("failed to run app: %v", err)
	}

}