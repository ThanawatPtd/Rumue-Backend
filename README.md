
# SAProject

## Overview
SAProject is a Go-based application that utilizes Docker for environment setup and PostgreSQL as the database. This README provides a step-by-step guide for setting up the project, running database migrations, and starting the application.

## Prerequisites
Before starting, ensure that you have the following installed:
- [Go](https://golang.org/doc/install)
- [Docker](https://docs.docker.com/get-docker/)
- [Goose](https://github.com/pressly/goose) (for database migrations)
- PostgreSQL

## Project Setup

### 1. Clone the Repository
Clone the SAProject repository to your local environment:
```bash
git clone https://github.com/ThanawatPtd/SAProject.git
```

### 2. Fetch the Latest Changes
Navigate to the project directory and fetch the latest changes:
```bash
git fetch
```

### 3. Checkout the Desired Branch
List all available branches:
```bash
git branch -a
```

### 4. Create and Switch to the `develop` Branch
Create and switch to the `develop` branch:
```bash
git branch develop
git checkout develop
```

### 5. Start the Application Using Docker
Ensure that Docker is installed and running. Then, use the following command to bring up the necessary containers:
```bash
docker-compose up
```

### 6. Run Database Migrations
Use [Goose](https://github.com/pressly/goose) to run the database migrations. Replace the database credentials as needed:
```bash
goose -dir ./cmd/migrations postgres "host=localhost user=myuser dbname=mydatabase password=mypassword sslmode=disable" up
```

### 7. Run the Project
Run the application by executing the following command:
```bash
go run ./cmd/SAProject/.
```

## Attribute Type Mapping (Caution)

When working with Go and PostgreSQL, ensure that attribute types are correctly mapped:

| Go Type        | PostgreSQL Type |
| -------------- | --------------- |
| `string`       | `varchar/text`  |
| `int32`        | `int`           |
| `float64`      | `float`         |
| `time.Time`    | `timestamptz`   |
| `pgtype.UUID`  | `UUID`          |

Ensure proper type mapping to avoid conflicts during database operations.

## Additional Resources
- [Go](https://golang.org/doc/)
- [Docker](https://docs.docker.com/get-docker/)
- [Goose - Database Migrations](https://github.com/pressly/goose)
