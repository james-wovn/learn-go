package main

import (
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jbwovn/learn-go/l10n/api/db"
)

func main() {
	db.Connect()
	db.Migrate()
}
