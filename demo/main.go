package main

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"

	"github.com/artemijspavlovs/snippets/demo/internal/api/rest/server"
)

func runMigrations(dburl string) {
	m, err := migrate.New(
		"file://db/migrations",
		dburl,
	)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("failed to finish the migrations: %v ", err)
	}
}

func connectToDatabase(dburl string) *sqlx.DB {
	db, err := sqlx.Connect("pgx", dburl)
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	return db
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dburl := os.Getenv("DATABASE_URL")
	db := connectToDatabase(dburl)
	runMigrations(dburl)

	srv := server.New(db)
	srv.Router.Logger.Fatal(srv.Router.Start(":3100"))
}
