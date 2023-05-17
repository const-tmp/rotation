package main

import (
	"context"
	"fmt"
	"github.com/const-tmp/rotation"
	"github.com/const-tmp/rotation/db"
	"github.com/const-tmp/rotation/db/pgx"
	"github.com/const-tmp/rotation/extractor"
	"github.com/const-tmp/rotation/rotator"
	"github.com/const-tmp/rotation/trigger"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
	"syscall"
)

func main() {
	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_SSL"),
	))
	if err != nil {
		log.Fatal(err)
	}

	// file rotator
	file := rotator.NewFile(os.Getenv("DB_CREDS_FILE"))

	// database rotator
	dbRotator := db.DB{
		Host:     extractor.NewEnv(file, "POSTGRES_HOST"),
		Port:     extractor.NewEnv(file, "POSTGRES_PORT"),
		User:     extractor.NewEnv(file, "POSTGRES_USER"),
		Password: extractor.NewEnv(file, "POSTGRES_PASSWORD"),
	}

	trig := make(chan struct{})

	// create SIGHUP signal trigger
	go trigger.Signal(trig, syscall.SIGHUP)

	// start rotation
	go rotation.Rotate(trig, file)

	// configure creds rotation in pgx
	pgx.Configure(dbRotator, config)

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatal(err)
	}

	if err = pool.Ping(context.Background()); err != nil {
		log.Fatal(err)
	}
}
