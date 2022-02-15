package main

import (
	"flag"
	"log"
	"math/rand"
	"time"

	"github.com/bhankey/BD_lab/backend/internal/app"
	"github.com/bhankey/BD_lab/backend/pkg/logger"
	"github.com/joho/godotenv"
)

// nolint: gochecknoinits
func init() {
	rand.Seed(time.Now().UnixNano())
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	var configPath string
	flag.StringVar(&configPath, "config-path", "config/config.yaml", "path to configuration file")
	logger.Init()

	log := logger.GetLogger()
	log.Infoln("logger initialized")

	flag.Parse()

	a, err := app.NewApp(configPath)
	if err != nil {
		log.Fatal(err)
	}

	a.Start()
}
