package main

import (
	"database/sql"
	"eMobile/internal/config"
	"eMobile/internal/db"
	"errors"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

func main() {
	cfg := config.LoadConfig()
	db, err := db.ConnectToBD(cfg)
	if err != nil {
		panic(err)
	}

	defer func(db *sql.DB) {
		db.Close()
	}(db)

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://cmd/migrator/migrations", cfg.DB.DatabaseName, driver)
	if err != nil {
		panic(err)
	}

	cmd := os.Args[len(os.Args)-1]

	switch cmd {
	case "up":
		if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			log.Fatal(err)
		}
	case "down":
		if err := m.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			log.Fatal(err)
		}
	default:
		if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			log.Fatal(err)
		}
	}
	logrus.Info("Migrations complete")
}
