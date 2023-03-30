package database

import (
	"database/sql"
	"fmt"

	"github.com/gobuffalo/packr/v2"
	migrate "github.com/rubenv/sql-migrate"
)

var (
	DB  *sql.DB
)

func DbMigrate(dbParam *sql.DB) {
	migration := &migrate.PackrMigrationSource{
		Box: packr.New("migrations", "./sql_migrations"),
	}
	n, err := migrate.Exec(dbParam, "postgres", migration, migrate.Up)
	if err != nil {
		panic(err)
	}
	DB = dbParam
	fmt.Println("Applied", n, " migration.")
}
