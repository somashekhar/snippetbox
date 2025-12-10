package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	infoLog  *log.Logger
	errorLog *log.Logger
}

type config struct {
	addr      string
	staticDir string
}

func main() {
	// Parse command-line flags
	var cfg config
	flag.StringVar(&cfg.addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&cfg.staticDir, "static-dir", "./ui/static/", "Path to static files directory")
	flag.Parse()

	// Create loggers for writing info messages to stdout and error messages to stderr
	infolog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorlog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		infoLog:  infolog,
		errorLog: errorlog,
	}

	srv := &http.Server{
		Addr:     cfg.addr,
		Handler:  app.routes(),
		ErrorLog: errorlog,
	}

	// Start the HTTP server
	infolog.Printf("Starting server on %s", cfg.addr)
	err := srv.ListenAndServe()
	errorlog.Fatal(err)
}
