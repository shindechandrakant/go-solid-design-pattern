package design_pattern

import "sync"

type DatabaseConnection struct {
}

func (d *DatabaseConnection) Query(sql string) {
	_ = sql
}

var (
	dbInstance *DatabaseConnection
	once       sync.Once
)

func GetDatabaseConnection() *DatabaseConnection {
	once.Do(func() {
		dbInstance = &DatabaseConnection{}
	})
	return dbInstance
}
