package storage

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBCreationSettings struct {
	Host     string `env:"DB_HOST,notEmpty" envDefault:"localhost"`
	Port     uint   `env:"DB_PORT,notEmpty" envDefault:"5432"`
	User     string `env:"DB_USER,required,notEmpty"`
	Password string `env:"DB_PASSWORD,required,notEmpty"`
	DBName   string `env:"DB_NAME,required,notEmpty"`
}

func (settings *DBCreationSettings) GetConnectionString() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d",
		settings.Host,
		settings.User,
		settings.Password,
		settings.DBName,
		settings.Port,
	)
}

func CreateDB(settings *DBCreationSettings) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(settings.GetConnectionString()), &gorm.Config{})
}

func NewDBConnectionSettings() *DBCreationSettings {
	var cfg DBCreationSettings
	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}
	return &cfg
}

func NewDB(settings *DBCreationSettings) *gorm.DB {
	db, err := CreateDB(settings)
	if err != nil {
		panic(err.Error())
	}
	return db
}
