package main

import (
	"coding-challenge-go/cmd/api/config"
	"coding-challenge-go/pkg/api"
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs

	db, err := sql.Open("mysql", "user:password@tcp(db:3306)/product")

	if err != nil {
		log.Error().Err(err).Msg("Fail to create server")
		return
	}

	defer db.Close()
	appConfig := getConfig()
	engine, err := api.CreateAPIEngine(db, appConfig)

	if err != nil {
		log.Error().Err(err).Msg("Fail to create server")
		return
	}

	log.Info().Msg("Start server")
	log.Fatal().Err(engine.Run(os.Getenv("LISTEN"))).Msg("Fail to listen and serve")
}

func getConfig() config.Config {
	f, err := os.Open("config/config.yml")
	if err != nil {
		log.Fatal().Err(err).Msg("Fail to open configurations file. Please check permissions")
	}
	defer f.Close()

	var cfg config.Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("Fail to decode configurations file")
	}
	return cfg
}
