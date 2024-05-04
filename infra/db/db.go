package db

import (
	"database/sql"
	"fmt"
	"log"

	"expert_systems_api/infra/config"

	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	err error
)

func handleDatabaseConnection() {

	appConfig := config.AppConfig()

	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		appConfig.DBHost,
		appConfig.DBPort,
		appConfig.DBUser,
		appConfig.DBPassword,
		appConfig.DBName,
	)

	db, err = sql.Open(appConfig.DBDialect, dataSourceName)

	if err != nil {
		log.Fatalf("error occured while trying to validate database arguments: %s\n", err.Error())
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("error occured while trying to connect to database: %s\n", err)
	}
}

func handleRequestTables() {
	const (
		createTableUserQuery = `
			CREATE TABLE IF NOT EXISTS "user" (
				id serial primary key,
				full_name varchar(60) not null,
				email varchar(60) not null unique,
				password text not null,
				role varchar(10) not null,
				created_at timestamptz default now(),
				updated_at timestamptz default now(),
				deleted_at timestamptz
			);
		`

		createTableSymtompQuery = `
			CREATE TABLE IF NOT EXISTS "symtomp" (
				id serial primary key,
				name varchar(60) not null,
				description text not null,
				created_at timestamptz default now(),
				updated_at timestamptz default now(),
				deleted_at timestamptz
			)
		`
	)

	if _, err := db.Exec(createTableUserQuery); err != nil {
		log.Fatalf("error while create table user: %s\n", err.Error())
	}

	if _, err := db.Exec(createTableSymtompQuery); err != nil {
		log.Fatalf("error while create table symtomp: %s\n", err.Error())
	}
}

func InitializeDatabase() {
	handleDatabaseConnection()
	handleRequestTables()
}

func GetDbInstance() *sql.DB {
	return db
}
