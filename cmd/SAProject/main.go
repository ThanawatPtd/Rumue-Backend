package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/ThanawatPtd/SAProject/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
	"github.com/pressly/goose/v3"
)

func main() {
	ctx := context.Background()
	config := config.ProvideConfig()
	postgresURI := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		config.PostgresHost, config.PostgresPort, config.PostgresUser, config.PostgresPassword, config.PostgresDB)

	// Connect to the database
	db, err := sql.Open("postgres", postgresURI)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	connConfig, err := pgxpool.ParseConfig(postgresURI)
	if err != nil {
		panic(err)
	}

	pgxPool, err := pgxpool.NewWithConfig(context.Background(), connConfig)
	if err != nil {
		panic(err)
	}

	// Connection Teardown
	conn, err := pgxPool.Acquire(ctx)
	if err != nil {
		panic(err)
	}
	defer conn.Release()

	if err := conn.Conn().Ping(ctx); err != nil {
		panic(err)
	}

	fmt.Println("🫙 Connected to Postgres")

	if err := goose.Up(db, "../migrations/"); err != nil {
		panic(err)
	}

	app := fiber.New(fiber.Config{
		AppName: "SEProject",
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000", // Replace with your frontend URL
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))
	app.Listen(":3001")

}
