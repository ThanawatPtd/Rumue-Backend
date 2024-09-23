
# Project Setup Guide

Follow these steps to set up the project on your local environment:

### 1. Clone the Repository
```bash
git clone https://github.com/ThanawatPtd/SAProject.git
```

### 2. Fetch the Latest Changes
```bash
git fetch
```

### 3. Checkout the Desired Branch
```bash
git branch -a
```

### 4. Create and Switch to the `develop` Branch
```bash
git branch develop
git checkout develop
```

### 5. Start the Application Using Docker
Make sure you have Docker installed and running. Then use the following command to bring up the necessary containers:
```bash
docker-compose up
```

### 6. Run Database Migrations
Use [Goose](https://github.com/pressly/goose) to run the database migrations. Replace the credentials as needed:
```bash
goose -dir ./cmd/migrations postgres "host=localhost user=myuser dbname=mydatabase password=mypassword sslmode=disable" up
```

### 7. Run the Project
Start the project by running the following command:
```bash
go run ./cmd/SAProject/.
```
