package migration

type IDBMigrator interface {
	Migrate() error
}
