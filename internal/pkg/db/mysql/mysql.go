package mysql

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

var Db *sql.DB

func InitDB() {
	// Use root:qwer@tcp(172.17.0.2)/hackernews, if you're using Windows.
	db, err := sql.Open("mysql", "root:qwer@tcp(localhost)/hackernews")
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	Db = db
}

func Migrate() {
	if err := Db.Ping(); err != nil {
		log.Fatal(err)
	}
	driver, _ := mysql.WithInstance(Db, &mysql.Config{})
	m, merr := migrate.NewWithDatabaseInstance(
		"file://internal/pkg/db/migrations/mysql",
		"hackernews",
		driver,
	)
	if merr != nil {
		log.Fatal(merr)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

}
