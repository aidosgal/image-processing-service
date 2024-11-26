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

	var imageID int64
	err := r.db.QueryRow(`
		INSERT INTO images (
			filename,
			file_size,
			mime_type,
			width,
			height,
			file_path,
			thumbnail_path,
			image_format
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id
	`, metadata.GetFilename(),
		metadata.GetFileSize(),
		metadata.GetMimeType(),
		metadata.GetWidth(),
		metadata.GetHeight(),
		metadata.GetFilePath(),
		metadata.GetThumbnailPath(),
		metadata.GetImageFormat(),
	).Scan(&imageID)
	if err != nil {
		return -1, fmt.Errorf("%s: %w", op, err)
	}

	return imageID, nil
}

func (r *Repository) GetAllImages(ctx context.Context) ([]*imagev1.ImageMetadata, error) {
	const op = "psql.GetAllImages"

	rows, err := r.db.QueryContext(ctx, `
		SELECT
			id,
			filename,
			file_size,
			mime_type,
			width,
			height,
			file_path,
			thumbnail_path,
			image_format
		FROM images
		ORDER BY uploaded_at DESC
	`)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to query images: %w", op, err)
	}
	defer rows.Close()

	var images []*imagev1.ImageMetadata
	for rows.Next() {
		var img imagev1.ImageMetadata

		err := rows.Scan(
			&img.ImageId,
			&img.Filename,
			&img.FileSize,
			&img.MimeType,
			&img.Width,
			&img.Height,
			&img.FilePath,
			&img.ThumbnailPath,
			&img.ImageFormat,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: failed to scan image row: %w", op, err)
		}

		images = append(images, &img)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: row iteration error: %w", op, err)
	}

	return images, nil
}

func (r *Repository) GetImageById(ctx context.Context, imageID int64) (*imagev1.ImageMetadata, error) {
	const op = "psql.GetImageById"

	var img imagev1.ImageMetadata

	err := r.db.QueryRowContext(ctx, `
		SELECT
			id,
			filename,
			file_size,
			mime_type,
			width,
			height,
			file_path,
			thumbnail_path,
			image_format
		FROM images
		WHERE id = $1
	`, imageID).Scan(
		&img.ImageId,
		&img.Filename,
		&img.FileSize,
		&img.MimeType,
		&img.Width,
		&img.Height,
		&img.FilePath,
		&img.ThumbnailPath,
		&img.ImageFormat,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("%s: image not found: %w", op, err)
	}
	if err != nil {
		return nil, fmt.Errorf("%s: failed to retrieve image: %w", op, err)
	}

	return &img, nil
}

func (r *Repository) DeleteImageById(ctx context.Context, imageID int64) (bool, error) {
	const op = "psql.DeleteImageById"

	var exists bool
	err := r.db.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM images WHERE id = $1)", imageID).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("%s: failed to check image existence: %w", op, err)
	}

	if !exists {
		return false, fmt.Errorf("%s: image not found", op)
	}

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return false, fmt.Errorf("%s: failed to begin transaction: %w", op, err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	result, err := tx.ExecContext(ctx, "DELETE FROM images WHERE id = $1", imageID)
	if err != nil {
		return false, fmt.Errorf("%s: failed to delete image record: %w", op, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("%s: failed to verify deletion: %w", op, err)
	}

	return rowsAffected > 0, nil
}
