package main

import (
	"flag"
	//"log"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))
	app := &application{
		logger: logger,
	}

	addr := flag.String("addr", ":9999", "http port value")
	flag.Parse()
	logger.Info("starting server", "post", *addr)
	err := http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
