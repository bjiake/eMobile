package config

import (
	"eMobile/docs"
	configDb "eMobile/internal/config/db"
	configSrv "eMobile/internal/config/server"
	"errors"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

type Environment string

const (
	StageEnv Environment = "stage"
	LocalEnv Environment = "local"
	DevEnv   Environment = "dev"
	ProdEnv  Environment = "prod"
)

var GlobalEnv Environment

type Config struct {
	DB     configDb.Database
	Server configSrv.Server
}

func LoadConfig() *Config {
	err := loadFlags()
	if err != nil {
		log.Fatal(err)
	}

	err = loadDotEnv()
	if err != nil {
		log.Fatal(err)
	}

	srv := configSrv.InitServerConfig()

	db := configDb.InitDbConfig()

	return &Config{
		DB:     db,
		Server: srv,
	}
}

func loadDotEnv() error {
	filePath := fmt.Sprintf(".env.%s", GlobalEnv)

	err := godotenv.Load(filePath)
	return err
}

func loadFlags() error {
	envFlag := flag.String("env", string(LocalEnv), "Environment type")
	flag.Parse()

	switch Environment(*envFlag) {
	case StageEnv, LocalEnv, DevEnv, ProdEnv:
		GlobalEnv = Environment(*envFlag)
	default:
		return errors.New("invalid environment type")
	}

	return nil
}

func SetSwaggerDefaultInfo(cfg *Config) {
	docs.SwaggerInfo.Title = "IOD App API"
	docs.SwaggerInfo.Description = "API Server for IOD application."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", cfg.Server.Addr, cfg.Server.Port)
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}
