package main

import (
	"errors"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/KonnorFrik/ChatServer/pkg/db"
)

func main() {
    m, err := migrate.New(
        "file://db/migrations",
        db.DefaultConfig.ToURL(),
    )

    if err != nil {
        log.Fatalln("[migration/New]: Error:", err)
    }

    err = m.Up()

    if errors.Is(err, migrate.ErrNoChange) {
        log.Println("[migration]: Successfull")
        return
    }

    if err != nil {
        log.Fatalln("[migration/Up]: Error:", err)
    }
}
