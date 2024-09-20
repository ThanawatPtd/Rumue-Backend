package main

import (
	"context"
	"fmt"

	"github.com/ThanawatPtd/SAProject/domain/usecases"
	"github.com/ThanawatPtd/SAProject/internal/adapters/rest"
	"github.com/ThanawatPtd/SAProject/internal/adapters/sqlc"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jackc/pgx/v5/pgxpool"
)

const(
	PostgresHost = "localhost"
	PostgresPort = 5432
	PostgresDB = "mydatabase"
	PostgresUser = "myuser"
	PostgresPassword = "mypassword"
)	

func main() {
	app := fiber.New(fiber.Config{
		AppName: "SEProject",
	})

	ctx := context.Background()

	postgresURI := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
        PostgresHost, PostgresPort, PostgresUser, PostgresPassword, PostgresDB)

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

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000", // Replace with your frontend URL
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))

	userRepo := sqlc.NewSqlcUserRepository(pgxPool)
	userService := usecases.NewUserService(userRepo, &ctx)

	userOuterSpaceHandler := rest.NewUserRestHandler(userService)

	app.Post("/user/create", userOuterSpaceHandler.CreateUser)

	app.Listen(":3001")

}
