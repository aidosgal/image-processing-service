package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/aidosgal/image-processing-service/internal/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	var migrationPath, migrationTable string
	flag.StringVar(&migrationPath, "migrations-path", "", "path to migrations")
	flag.StringVar(&migrationTable, "migration-table", "migrations", "name of migration table")
	cfg := config.MustLoad()
	flag.Parse()

	if migrationPath == "" {
		panic("migration path no defined")
	}

	postgresURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
		cfg.Database.SSLMode,
	)

	m, err := migrate.New("file://"+migrationPath, postgresURL)
	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no migration to apply")
			return
		}
		panic(err)
	}

	fmt.Println("migration apllied successfully")
}
