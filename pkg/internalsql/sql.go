package internalsql

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Connect(dataSourceName string) (*sql.DB, error) {
	return sql.Open("postgres", dataSourceName)
}
