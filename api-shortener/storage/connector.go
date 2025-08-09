package storage

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConnectionSettings struct {
	Host     string `env:"DB_HOST,notEmpty" envDefault:"localhost"`
	Port     uint   `env:"DB_PORT,notEmpty" envDefault:"5432"`
	User     string `env:"DB_USER,required,notEmpty"`
	Password string `env:"DB_PASSWORD,required,notEmpty"`
	DBName   string `env:"DB_NAME,required,notEmpty"`
}

func (settings *DBConnectionSettings) String() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d",
		settings.Host,
		settings.User,
		settings.Password,
		settings.DBName,
		settings.Port,
	)
}

func CreateDBConnection(settings *DBConnectionSettings) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(settings.String()), &gorm.Config{})
}

func NewDBConnectionSettings() *DBConnectionSettings {
	var cfg DBConnectionSettings
	if err := env.Parse(&cfg); err != nil {
		logrus.Fatalf("Couldn't parse DBCreationSettings from env: %s", err)
	}
	return &cfg
}

func NewDB(settings *DBConnectionSettings) *gorm.DB {
	db, err := CreateDBConnection(settings)
	if err != nil {
		logrus.Fatalf("Couldn't connect to the DB: %s", err)
	}
	return db
}
