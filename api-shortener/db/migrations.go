package db

import (
	"api-shortener/shortreq"

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
		&shortreq.ShortenedAPI{},
		&shortreq.OutgoingRequestConfig{},
		&shortreq.OutgoingRequestHeader{},
		&shortreq.OutgoingRequestParam{},
		&shortreq.ShorteningRule{},
	)
}

func NewMigrator(conn *gorm.DB) IDBMigrator {
	return &DBMigrator{DB: conn}
}
