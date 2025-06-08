package main

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
    m, err := migrate.New(
        "",
        "",
    )

    if err != nil {
        log.Fatalln("[migration/New]: Error:", err)
    }

    err = m.Up()

    if err != nil {
        log.Fatalln("[migration/Down]: Error:", err)
    }
}
