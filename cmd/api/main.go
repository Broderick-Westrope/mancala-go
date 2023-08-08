package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
	flag "github.com/spf13/pflag"
)

func main() {
	addr := flag.StringP("addr", "a", ":4000", "HTTP network address")

	flag.Parse()

	err := run(addr)
	if err != nil {
		log.Fatal().Err(err).Send()
	}
}

func run(address *string) error {
	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("PLACEHOLDER"))
		if err != nil {
			log.Err(err).Send()
		}
	})

	err := http.ListenAndServe(*address, router)
	if err != nil {
		return err
	}

	return nil
}
