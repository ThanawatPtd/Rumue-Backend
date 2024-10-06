package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ThanawatPtd/SAProject/config"
	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/pressly/goose/v3"
)

func main() {
	// Command-line flags for migration direction and version
	var direction string
	var version int64
	flag.StringVar(&direction, "direction", "up", "Specify migration direction: up or down")
	flag.Int64Var(&version, "version", 0, "Specify migration version (use 0 for latest/full migration)")
	flag.Parse()

	config := config.ProvideConfig()
	// Database connection string (PostgreSQL)
	postgresURI := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.PostgresHost, config.PostgresPort, config.PostgresUser, config.PostgresPassword, config.PostgresDB)
	fmt.Println(postgresURI)
	// dbURL := "postgres://user:password@localhost:5432/mydb?sslmode=disable"

	// Open a database connection
	db, err := sql.Open("postgres", postgresURI)
	if err != nil {
		log.Fatalf("Failed to open database connection: %v\n", err)
	}
	defer db.Close()

	// Check if the connection is valid
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to connect to the database: %v\n", err)
	}

	fmt.Println("Connected to the database successfully!")

	// Get the current working directory (for locating migration files)
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get working directory: %v\n", err)
	}

	// Path to migrations directory
	migrationsDir := fmt.Sprintf("%s/cmd/migrations", dir)

	// Handle migration based on direction (up/down) and version
	switch direction {
	case "up":
		if version > 0 {
			// Migrate up to a specific version
			fmt.Printf("Migrating up to version: %d\n", version)
			if err := goose.UpTo(db, migrationsDir, version); err != nil {
				log.Fatalf("Failed to migrate up to version %d: %v\n", version, err)
			}
		} else {
			// Migrate to the latest version
			fmt.Println("Migrating to the latest version...")
			if err := goose.Up(db, migrationsDir); err != nil {
				log.Fatalf("Failed to apply latest migrations: %v\n", err)
			}
		}

	case "down":
		if version > 0 {
			// Migrate down to a specific version
			fmt.Printf("Migrating down to version: %d\n", version)
			if err := goose.DownTo(db, migrationsDir, version); err != nil {
				log.Fatalf("Failed to migrate down to version %d: %v\n", version, err)
			}
		} else {
			// Migrate down one step
			fmt.Println("Migrating down one step...")
			if err := goose.Down(db, migrationsDir); err != nil {
				log.Fatalf("Failed to apply migration down: %v\n", err)
			}
		}

	case "reset":
		// Reset all migrations (down all tables)
		fmt.Println("Resetting all migrations (down all tables)...")
		if err := goose.Reset(db, migrationsDir); err != nil {
			log.Fatalf("Failed to reset migrations: %v\n", err)
		}

	default:
		log.Fatalf("Invalid direction: %s. Use 'up' or 'down'.\n", direction)
	}

	fmt.Println("Migrations applied successfully!")
}
