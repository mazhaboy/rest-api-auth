package configs

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // revive:disable:blank-imports
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

func Migrate(path, dbURI string) error {
	if m, err := migrate.New("file://"+path, dbURI); err != nil {
		return fmt.Errorf("couldn't init the migrate object: %w", err)
	} else if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("couldn't migrate the database: %w", err)
	}

	log.Println("migrations are applied")

	return nil
}
