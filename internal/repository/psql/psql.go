package psql

import (
	"database/sql"
	"fmt"

	"github.com/aidosgal/image-processing-service/internal/config"
	_ "github.com/lib/pq"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(cfg config.DatabaseConfig) (*Repository, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/postgres?sslmode=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.SSLMode)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}
	defer db.Close()

	err = createDatabaseIfNotExists(db, cfg.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to create the database: %w", err)
	}

	connStr = fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name, cfg.SSLMode)
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the newly created database: %w", err)
	}

	return &Repository{db: db}, nil
}

func createDatabaseIfNotExists(db *sql.DB, dbName string) error {
	var exists bool
	err := db.QueryRow(`SELECT EXISTS (
		SELECT FROM pg_catalog.pg_database
		WHERE datname = $1
	);`, dbName).Scan(&exists)
	if err != nil {
		return fmt.Errorf("failed to check if database exists: %w", err)
	}

	if exists {
		return nil
	}

	_, err = db.Exec(fmt.Sprintf(`CREATE DATABASE %s`, dbName))
	if err != nil {
		return fmt.Errorf("failed to create the database: %w", err)
	}

	return nil
}
