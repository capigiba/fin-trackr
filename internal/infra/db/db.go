package db

import (
	"database/sql"
	"fintrack/internal/infra/env"
	"fintrack/internal/infra/zap-logging/log"
	"io/ioutil"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

// InitializeDB initializes the SQLite3 database and returns a database connection.
func InitializeDB() (*sql.DB, error) {
	dbPath := env.EnvConfig.DBConnection
	dbDriver := env.EnvConfig.DBDriver

	// Open a connection to the database file (creates the file if it doesn't exist)
	db, err := sql.Open(dbDriver, dbPath)
	if err != nil {
		return nil, err
	}

	// Check if the connection is successful
	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	log.Info("Database connection established successfully!")
	return db, nil
}

// RunMigration runs the given SQL file.
func RunMigration(db *sql.DB, filePath string) error {
	query, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	_, err = db.Exec(string(query))
	if err != nil {
		return err
	}

	log.Info("Migration %s executed successfully!", filePath)
	return nil
}

// RunMigrations runs all migrations in the migrations directory.
func RunMigrations(db *sql.DB) error {
	migrationFiles := []string{
		"001_create_users_table.up.sql",
		"002_create_transactions_table.up.sql",
		"003_insert_sample_users.up.sql",
	}

	for _, migrationFile := range migrationFiles {
		filePath := filepath.Join("..", "..", "internal", "infra", "migrations", migrationFile)
		err := RunMigration(db, filePath)
		if err != nil {
			return err
		}
	}

	return nil
}

// RunDownMigration runs the given SQL file.
func RunDownMigration(db *sql.DB, filePath string) error {
	query, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	_, err = db.Exec(string(query))
	if err != nil {
		return err
	}

	log.Info("Successfully reverted %s", filePath)
	return nil
}

// RunDownMigrations runs all down migrations in reverse order to undo changes.
func RunDownMigrations(db *sql.DB) error {
	// Notes: reverse order of application
	migrationFiles := []string{
		"003_insert_sample_users.down.sql",
		"002_create_transactions_table.down.sql",
		"001_create_users_table.down.sql",
	}

	for _, file := range migrationFiles {
		filePath := filepath.Join("migrations", file)
		query, err := ioutil.ReadFile(filePath)
		if err != nil {
			return err
		}

		_, err = db.Exec(string(query))
		if err != nil {
			return err
		}

		log.Info("Successfully reverted %s", file)
	}

	return nil
}
