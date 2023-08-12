package main

import (
	"flag"
	"net/http"

	log "golang.org/x/exp/slog"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()

	err := run(addr)
	if err != nil {
		log.Error(err.Error())
	}
}

func run(address *string) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("PLACEHOLDER"))
		if err != nil {
			log.Error(err.Error())
		}
	})

	log.Info("Starting server", log.String("addr", *address))
	err := http.ListenAndServe(*address, mux)
	if err != nil {
		return err
	}

	return nil
}
