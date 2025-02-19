//go:build linux && cgo && !agent

package cluster

import (
	"database/sql"
	"time"
)

// Code generation directives.
//
//go:generate -command mapper generate-database db mapper -t images.mapper.go
//go:generate mapper generate -i -b "//go:build linux && cgo && !agent"
//
//generate-database:mapper stmt image objects
//generate-database:mapper stmt image objects-by-ID
//generate-database:mapper stmt image objects-by-Project
//generate-database:mapper stmt image objects-by-Project-and-Cached
//generate-database:mapper stmt image objects-by-Project-and-Public
//generate-database:mapper stmt image objects-by-Fingerprint
//generate-database:mapper stmt image objects-by-Cached
//generate-database:mapper stmt image objects-by-AutoUpdate
//
//generate-database:mapper method image GetMany
//generate-database:mapper method image GetOne

// Image is a value object holding db-related details about an image.
type Image struct {
	ID           int
	Project      string `db:"primary=yes&join=projects.name"`
	Fingerprint  string `db:"primary=yes"`
	Type         int
	Filename     string
	Size         int64
	Public       bool
	Architecture int
	CreationDate sql.NullTime
	ExpiryDate   sql.NullTime
	UploadDate   time.Time
	Cached       bool
	LastUseDate  sql.NullTime
	AutoUpdate   bool
}

// ImageFilter can be used to filter results yielded by GetImages.
type ImageFilter struct {
	ID          *int
	Project     *string
	Fingerprint *string
	Public      *bool
	Cached      *bool
	AutoUpdate  *bool
}
