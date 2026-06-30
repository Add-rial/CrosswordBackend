# Crossword Backend

A backend service for a multiplayer crossword puzzle platform developed as part of the BITS Pilani Coding Club. The project provides authentication, crossword management, leaderboard tracking, and administrative APIs.

## Features

- User authentication
- Google OAuth login
- JWT authorization
- Crossword generation and retrieval
- Leaderboard management
- Score calculation
- Administrative APIs
- Swagger API documentation
- PostgreSQL database integration

## Tech Stack

- Go
- Gin
- PostgreSQL
- JWT

## Project Structure

```
config/
controllers/
middleware/
models/
routes/
utils/
```

## Getting Started

### Clone the repository

```bash
git clone https://github.com/Add-rial/CrosswordBackend.git
```

### Install dependencies

```bash
go mod tidy
```

### Configure environment variables

Create a `.env` file containing the required database credentials and JWT secret.

### Run the project

```bash
go run main.go
```

## Documentation

Swagger documentation is available within the project for testing and exploring the available APIs.
