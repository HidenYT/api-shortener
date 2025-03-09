package db

import (
	"api-shortener/restapi"

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
		&restapi.ShortenedAPI{},
		&restapi.OutgoingRequestConfig{},
		&restapi.OutgoingRequestHeader{},
		&restapi.OutgoingRequestParam{},
		&restapi.ShorteningRule{},
	)
}

func NewMigrator(conn *gorm.DB) IDBMigrator {
	return &DBMigrator{DB: conn}
}
