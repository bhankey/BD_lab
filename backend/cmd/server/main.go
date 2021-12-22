package main

import (
	"finance/internal/app"
	"finance/pkg/logger"
	"flag"
	"github.com/joho/godotenv"
	"log"
	"math/rand"
	"time"
)

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

	l := logger.GetLogger()
	l.Infoln("logger initialized")

	flag.Parse()

	a, err := app.NewApp(configPath)
	if err != nil {
		l.Fatal(err)
	}

	a.Start()
}
