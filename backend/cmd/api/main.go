package main

import (
	"flag"
	"fmt"
	"log/slog"	
	"net/http"
	"os"
	"time"
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

	app := application{
		config: cfg,
		logger: logger,
	}

	
}
