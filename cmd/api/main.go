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

	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("PLACEHOLDER"))
		if err != nil {
			log.Error().Err(err).Msg("")
		}
	})

	err := http.ListenAndServe(*addr, router)
	if err != nil {
		log.Err(err)
	}
}
