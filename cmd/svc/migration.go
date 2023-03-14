package main

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/imamponco/v-gin-boilerplate/src/pkg/vtype"
	"github.com/imamponco/v-gin-boilerplate/src/svc/contract"
	"path/filepath"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func bootMigration(workDir string, config *contract.Config) error {
	// If not enabled, then skip
	isEnabled := vtype.ParseBooleanFallback(config.DatabaseBootMigration, false)
	if !isEnabled {
		return nil
	}

	// Load database url
	dbUri := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable",
		config.DatabaseDriver,
		config.DatabaseUsername, config.DatabasePassword,
		config.DatabaseHost, config.DatabasePort,
		config.DatabaseName,
	)

	// Source dir
	sourcePath := filepath.Join(workDir, "/migrations/sql")
	sourceUri := "file://" + sourcePath

	// Init database
	log.Debug("migration: Connecting to database...")
	m, err := migrate.New(sourceUri, dbUri)
	if err != nil {
		log.Errorf("migration: Failed to connect database %s", err)
		return err
	}

	err = m.Up()
	if err != nil {
		if err.Error() == "no change" {
			log.Infof("migration: No changes")
			return nil
		}
		log.Error("migration: Failed to run up migration scripts", err)
		return err
	}

	// Get status
	version, dirty, err := m.Version()
	if err != nil {
		log.Error("migration: Failed to get database migration version")
		return err
	}
	log.Infof("migration: Database version = %d, Forced = %v", version, dirty)

	return nil
}
