package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/somashekhar/snippetbox/internal/models"
)

type application struct {
	infoLog  *log.Logger
	errorLog *log.Logger
	snippets *models.SnippetModel
}

func main() {
	// Parse command-line flags
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL data source name")
	flag.Parse()

	// Create loggers for writing info messages to stdout and error messages to stderr
	infolog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorlog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Database
	db, err := openDB(*dsn)
	if err != nil {
		errorlog.Fatal(err)
	}
	defer db.Close()
	infolog.Println("Database connection pool established")

	app := &application{
		infoLog:  infolog,
		errorLog: errorlog,
		snippets: &models.SnippetModel{DB: db},
	}

	srv := &http.Server{
		Addr:     *addr,
		Handler:  app.routes(),
		ErrorLog: errorlog,
	}

	// Start the HTTP server
	infolog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorlog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
