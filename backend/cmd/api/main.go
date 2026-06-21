package main

import (
	"flag"
	"fmt"
	"log/slog"	
	"net/http"
	"os"
	
)

const version = "1.0.0"

type config struct {
	port int
	env string
}


type application struct {
	config config
	logger *slog.Logger
}

func main() {
	var cfg config 

	flag.IntVar(&cfg.port, "port", 4000,"API Server Port")
	flag.StringVar(&cfg.env,"env","development","Environment (dev,stage,prod)")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout,nil))

	app := &application{
		config: cfg,
		logger: logger,
	}
	mux := http.NewServeMux()

	srv := NewServer(fmt.Sprintf("%s", cfg.port),mux)

	app.logger.Info("starting server","address", srv.Addr, "env", cfg.env)
	err := srv.ListenAndServe()
	app.logger.Error(err.Error())
	os.Exit(1)

}
