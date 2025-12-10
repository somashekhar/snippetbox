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

	// Initialize a new HTTP server mux (router)
	mux := http.NewServeMux()

	// Serve static files from ./ui/static/ directory under the /static/ URL path prefix
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	srv := &http.Server{
		Addr:     cfg.addr,
		Handler:  mux,
		ErrorLog: errorlog,
	}

	// Start the HTTP server
	infolog.Printf("Starting server on %s", cfg.addr)
	err := srv.ListenAndServe()
	errorlog.Fatal(err)
}
