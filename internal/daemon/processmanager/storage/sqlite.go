package processmanagerstorage

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/olexnzarov/gofu/pkg/gofu"
)

func WithSQLite(directories gofu.Directories) (*SQLDB, error) {
	db, err := sql.Open(
		"sqlite3",
		fmt.Sprintf("%s/daemon.db", directories.DataDirectory),
	)
	if err != nil {
		return nil, err
	}
	return &SQLDB{db}, nil
}
