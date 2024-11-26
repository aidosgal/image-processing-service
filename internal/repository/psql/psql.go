package psql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/aidosgal/image-processing-service/internal/config"
	imagev1 "github.com/aidosgal/image-processing-service/pkg/gen/go/image"
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

	connStr = fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name, cfg.SSLMode)
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the newly created database: %w", err)
	}

	return &Repository{db: db}, nil
}

func (r *Repository) StoreImage(ctx context.Context, metadata *imagev1.ImageMetadata) (int64, error) {
	const op = "psql.StoreImage"
	result, err := r.db.Exec(`
		INSERT INTO images (
			filename,
			file_size,
			mime_type,
			width,
			height,
			uploaded_at,
			updated_at,
			file_path,
			thumbnail_path,
			image_format,
			tags
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, metadata.GetFilename(), metadata)
	if err != nil {
		return -1, fmt.Errorf("%s: %w", op, err)
	}

	image_id, err := result.LastInsertId()
	if err != nil {
		return -1, fmt.Errorf("%s: %w", op, err)
	}
	return image_id, nil
}

func (r *Repository) GetAllImages(ctx context.Context) ([]*imagev1.ImageMetadata, error) {
	return nil, nil
}

func (r *Repository) GetImageById(ctx context.Context, image_id int64) (*imagev1.ImageMetadata, error) {
	return nil, nil
}

func (r *Repository) DeleteImageById(ctx context.Context, image_id int64) (bool, error) {
	return false, nil
}
