package storage

import (
	api_dao "github.com/HidenYT/api-shortener/internal/storage/dao"
	"gorm.io/gorm"
)

type IDBMigrator interface {
	Migrate()
}

type DBMigrator struct {
	DB *gorm.DB
}

func (m *DBMigrator) Migrate() {
	m.DB.AutoMigrate(
		&api_dao.ShortenedAPI{},
		&api_dao.OutgoingRequestConfig{},
		&api_dao.OutgoingRequestHeader{},
		&api_dao.OutgoingRequestParam{},
		&api_dao.ShorteningRule{},
	)
}

func NewMigrator(conn *gorm.DB) IDBMigrator {
	return &DBMigrator{DB: conn}
}
