package storage

type IDBMigrator interface {
	Migrate() error
}
