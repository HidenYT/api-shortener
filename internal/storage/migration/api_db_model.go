package migration

import (
	db_model "github.com/HidenYT/api-shortener/internal/storage/db-model/api"
	"gorm.io/gorm"
)

type DBMigrator struct {
	DB *gorm.DB
}

func (m *DBMigrator) Migrate() error {
	return m.DB.AutoMigrate(
		&db_model.ShortenedAPI{},
		&db_model.OutgoingRequestConfig{},
		&db_model.OutgoingRequestHeader{},
		&db_model.OutgoingRequestParam{},
		&db_model.ShorteningRule{},
	)
}

func NewAPIDBMigrator(conn *gorm.DB) IDBMigrator {
	return &DBMigrator{DB: conn}
}
