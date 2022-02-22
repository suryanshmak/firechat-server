package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v4"
)

type application struct {
	infoLog  *log.Logger
	errorLog *log.Logger
}

func main() {
	port := flag.String("p", ":8080", "Which port should the server run?")
	dsn := flag.String("dsn", "postgres://web:secretpassword@locahost:5432/firechat", "PostgreSQL data source name")
	flag.Parse()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {

	}

	app := &application{
		infoLog:  infoLog,
		errorLog: errorLog,
	}

	srv := &http.Server{
		Addr:     *port,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}
	srv.ListenAndServe()
}

func openDB(dsn string) (*pgx.Conn, error) {
	db, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(context.Background()); err != nil {
		return nil, err
	}
	return db, nil
}
