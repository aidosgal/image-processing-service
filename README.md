# Image Proccesing Service

This document outlines the structure of the **Image Processing Service** project.

## Project Overview

The Image Processing Service is a Go-based application designed for uploading, processing, and managing images.
It provides a set of gRPC APIs for clients to interact with the service. The service stores image metadata in PostgreSQL and handles images (original and processed) on the file system or in cloud storage.
The core functionality includes uploading images, retrieving them, and managing metadata.

This service demonstrates effective use of concurrency in Go and is built following clean architecture principles, making the system modular, testable, and maintainable. The application is containerized using Docker to provide a consistent environment for deployment and execution.

## Key Components

The project consists of the following main components:

- gRPC Server: Exposes the service’s API for clients to interact with.
- Image Processor: Handles image processing (e.g., creating thumbnails).
- Database (PostgreSQL): Stores metadata related to the images, such as file names, sizes, and processing status.
- File Storage: Manages the storage of original and processed images (either on disk or in cloud storage).
- Docker: Used for containerization, making it easier to deploy and manage the service.

## Directory Structure

Here’s a breakdown of the main directories and files in the project:
```
/cmd
  /image_service
    main.go                # Entry point for the image service
  /migrate
    main.go                # Entry point for running migrations

/config
  local.yaml              # Configuration for the image service

/internal
  app/
    grpc/
      app.go              # Entry point for the grpc
    app.go
  config/
    config.app            # Parsing the config from config/*.yaml
  delivary/
    image/                # Handles the grpc connection

  domain/                 # Common models for delivery
  lib/                    # Helpers
  service/                # Business Logic
  repository/             # Work with database

/migrations
  0001_initial_schema.sql # SQL migration file to create the database schema

/pkg
  /image
    processor.go           # Logic for processing images (e.g., creating thumbnails)
    storage.go             # File storage management (e.g., saving and retrieving images)
/proto
  image_service.proto      # gRPC service definitions for image upload, retrieval, etc.

/bin
  migrate                 # Binary for running migrations
  image_service           # Binary for running the image service

Dockerfile                # Dockerfile for building the image service container
docker-compose.yml        # Docker Compose configuration to run service and database together
go.mod                    # Go modules file for dependencies
go.sum                    # Go checksum file for dependencies
README.md                 # Project documentation (this file)
```

## Service Flow
The service interacts with the following components in a typical request flow:

### Image Upload:

- The client sends an UploadImage request via gRPC.
- The service receives the image and processes it asynchronously (e.g., creating a thumbnail).
- Metadata related to the image (e.g., filename, size, processing status) is saved to PostgreSQL.
- The original image and processed images (e.g., thumbnail) are stored on the file system.

### List Images:

The client sends a ListImages request to retrieve a list of image metadata.
The service queries PostgreSQL to retrieve metadata for all stored images.

### Get Image:

The client sends a GetImage request with the image ID.
The service retrieves the image (either original or processed) from storage and returns it.

### Delete Image:

The client sends a DeleteImage request with the image ID.
The service deletes the image and its associated metadata from file system storage and PostgreSQL.

## Database Schema: Images Table

The images table stores metadata about images uploaded to the system.
The schema includes fields for the image's filename, file size, MIME type, dimensions, upload and modification times, file paths,
and associated metadata such as tags and image format.
Below is the SQL definition for the images table:

```
CREATE TABLE IF NOT EXISTS images (
    id SERIAL PRIMARY KEY,              -- Unique identifier for each image (auto-incremented)
    filename VARCHAR(255) NOT NULL,      -- The original filename of the image
    file_size BIGINT NOT NULL,          -- The size of the image file in bytes
    mime_type VARCHAR(50) NOT NULL,     -- MIME type of the image (e.g., image/jpeg, image/png)
    width INTEGER NOT NULL,             -- Width of the image in pixels
    height INTEGER NOT NULL,            -- Height of the image in pixels
    uploaded_at TIMESTAMP DEFAULT NOW(),-- Timestamp for when the image was uploaded
    updated_at TIMESTAMP DEFAULT NOW(), -- Timestamp for when the image metadata was last updated
    file_path TEXT NOT NULL,            -- File path for storing the original image on disk or cloud
    thumbnail_path TEXT,                -- File path for storing the thumbnail of the image (nullable)
    image_format VARCHAR(50) NOT NULL,  -- Format of the image (e.g., jpeg, png)
    tags JSONB                          -- JSONB column to store image tags or other metadata
);
```

### Field Descriptions:
- id: A unique identifier for each image, automatically incremented by the database.
- filename: The name of the image file. This is important for identifying and retrieving the image from storage.
- file_size: The size of the image file in bytes. Useful for managing storage and validating file uploads.
- mime_type: Specifies the MIME type of the image (e.g., image/jpeg, image/png). This helps the application understand the type of the file for processing.
- width: The width of the image in pixels. This is useful for image resizing and processing.
- height: The height of the image in pixels. Similar to width, this is used for image manipulation and metadata storage.
- uploaded_at: The timestamp when the image was first uploaded to the system. This can be used for managing and querying uploaded images.
- updated_at: The timestamp when the image metadata was last updated (e.g., after processing). Automatically set to the current timestamp.
- file_path: The storage path for the original image. It can point to the file system or a cloud storage location.
- thumbnail_path: The file path to the thumbnail version of the image, if applicable. This field is nullable because not all images may have a thumbnail.
- image_format: The format of the image (e.g., jpeg, png). This helps in processing and managing images in different formats.
- tags: A JSONB column used to store tags or other metadata in JSON format. This allows for flexible, structured storage of additional image-related information, such as categories, keywords, or custom metadata.
