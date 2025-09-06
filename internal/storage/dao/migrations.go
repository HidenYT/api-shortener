package api_dao

import (
	"github.com/HidenYT/api-shortener/internal/storage"
	"gorm.io/gorm"
)

type DBMigrator struct {
	DB *gorm.DB
}

func (m *DBMigrator) Migrate() error {
	return m.DB.AutoMigrate(
		&ShortenedAPI{},
		&OutgoingRequestConfig{},
		&OutgoingRequestHeader{},
		&OutgoingRequestParam{},
		&ShorteningRule{},
	)
}

func NewMigrator(conn *gorm.DB) storage.IDBMigrator {
	return &DBMigrator{DB: conn}
}
